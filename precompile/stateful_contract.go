package precompile

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type statefulPrecompiledContract struct {
	abi abi.ABI
}

func NewStatefulPrecompiledContract(abiStr string) StatefulPrecompiledContract {
	abi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(err)
	}
	return &statefulPrecompiledContract{
		abi: abi,
	}
}

func (spc *statefulPrecompiledContract) GetABI() abi.ABI {
	return spc.abi
}

func (spc *statefulPrecompiledContract) RequiredGas(input []byte) uint64 {
	// This is a placeholder implementation. The actual gas required would depend on the specific contract.
	// You should replace this with the actual implementation.
	return 0
}
