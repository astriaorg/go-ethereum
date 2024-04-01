package vm

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
)

type methodID [4]byte

type statefulMethod struct {
	abiMethod     abi.Method
	reflectMethod reflect.Method
}

type precompileMethods map[methodID]*statefulMethod

type precompileManager struct {
	evm         *EVM
	precompiles map[common.Address]precompile.StatefulPrecompiledContract
	pMethods    map[common.Address]precompileMethods
}

func NewPrecompileManager(evm *EVM) PrecompileManager {
	precompiles := make(map[common.Address]precompile.StatefulPrecompiledContract)
	pMethods := make(map[common.Address]precompileMethods)
	return &precompileManager{
		evm:         evm,
		precompiles: precompiles,
		pMethods:    pMethods,
	}
}

func (pm *precompileManager) IsPrecompile(addr common.Address) bool {
	_, isEvmPrecompile := pm.evm.precompile(addr)
	if isEvmPrecompile {
		return true
	}

	_, isStatefulPrecompile := pm.precompiles[addr]
	return isStatefulPrecompile
}

func (pm *precompileManager) Run(
	addr common.Address,
	input []byte,
	caller common.Address,
	value *big.Int,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {

	// run core evm precompile
	p, isEvmPrecompile := pm.evm.precompile(addr)
	if isEvmPrecompile {
		return RunPrecompiledContract(p, input, suppliedGas)
	}

	contract, ok := pm.precompiles[addr]
	if !ok {
		return nil, 0, fmt.Errorf("no precompiled contract at address %v", addr.Hex())
	}

	// Extract the method ID from the input
	methodId := methodID(input)
	// Try to get the method from the precompiled contracts using the method ID
	method, exists := pm.pMethods[addr][methodId]
	if !exists {
		return nil, 0, fmt.Errorf("no method with id %v in precompiled contract at address %v", methodId, addr.Hex())
	}

	gasCost := contract.RequiredGas(input)
	if gasCost > suppliedGas {
		return nil, 0, ErrOutOfGas
	}

	// Unpack the input arguments using the ABI method's inputs
	unpackedArgs, err := method.abiMethod.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, err
	}

	// Convert the unpacked args to reflect values.
	reflectedUnpackedArgs := make([]reflect.Value, 0, len(unpackedArgs))
	for _, unpacked := range unpackedArgs {
		reflectedUnpackedArgs = append(reflectedUnpackedArgs, reflect.ValueOf(unpacked))
	}

	ctx := precompile.NewStatefulContext(pm.evm.StateDB, addr, caller, value)

	// Make sure the readOnly is only set if we aren't in readOnly yet.
	// This also makes sure that the readOnly flag isn't removed for child calls.
	if readOnly && !ctx.IsReadOnly() {
		ctx.SetReadOnly(true)
		defer func() { ctx.SetReadOnly(false) }()
	}

	results := method.reflectMethod.Func.Call(append(
		[]reflect.Value{
			reflect.ValueOf(contract),
			reflect.ValueOf(ctx),
		},
		reflectedUnpackedArgs...,
	))

	// check if precompile returned an error
	if len(results) > 0 {
		if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
			return nil, 0, err
		}
	}

	// Pack the result
	var output []byte
	if len(results) > 1 {
		interfaceArgs := make([]interface{}, len(results)-1)
		for i, v := range results[:len(results)-1] {
			interfaceArgs[i] = v.Interface()
		}
		output, err = method.abiMethod.Outputs.Pack(interfaceArgs...)
		if err != nil {
			return nil, 0, err
		}
	}

	suppliedGas -= gasCost
	return output, suppliedGas, nil
}

func (pm *precompileManager) Register(addr common.Address, p precompile.StatefulPrecompiledContract) error {
	if _, isEvmPrecompile := pm.evm.precompile(addr); isEvmPrecompile {
		return fmt.Errorf("precompiled contract already exists at address %v", addr.Hex())
	}

	if _, exists := pm.precompiles[addr]; exists {
		return fmt.Errorf("precompiled contract already exists at address %v", addr.Hex())
	}

	// niaeve implementation; parsed abi method names must match precompile method names 1:1
	//
	// Note on method naming:
	// Method name is the abi method name used for internal representation. It's derived from
	// the abi raw name and a suffix will be added in the case of a function overload.
	//
	// e.g.
	// These are two functions that have the same name:
	// * foo(int,int)
	// * foo(uint,uint)
	// The method name of the first one will be resolved as Foo while the second one
	// will be resolved as Foo0.
	//
	// Alternatively could require each precompile to define the func mapping instead of doing this magic
	abiMethods := p.GetABI().Methods
	contractType := reflect.ValueOf(p).Type()
	precompileMethods := make(precompileMethods)
	for _, abiMethod := range abiMethods {
		mName := strings.ToUpper(string(abiMethod.Name[0])) + abiMethod.Name[1:]
		reflectMethod, exists := contractType.MethodByName(mName)
		if !exists {
			return fmt.Errorf("precompiled contract does not implement abi method %s with signature %s", abiMethod.Name, abiMethod.RawName)
		}
		mID := methodID(abiMethod.ID)
		precompileMethods[mID] = &statefulMethod{
			abiMethod:     abiMethod,
			reflectMethod: reflectMethod,
		}
	}

	pm.precompiles[addr] = p
	pm.pMethods[addr] = precompileMethods
	return nil
}
