package precompile

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type StateDB interface {
	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	GetBalance(common.Address) *big.Int
	GetState(common.Address, common.Hash) common.Hash
	SetState(common.Address, common.Hash, common.Hash)
}

type StatefulContext interface {
	SetState(common.Hash, common.Hash) error
	GetState(common.Hash) common.Hash
	SubBalance(common.Address, *big.Int) error
	AddBalance(common.Address, *big.Int) error
	GetBalance(common.Address) *big.Int
	Address() common.Address
	MsgSender() common.Address
	MsgValue() *big.Int
	IsReadOnly() bool
	SetReadOnly(bool)
}

type StatefulPrecompiledContract interface {
	GetABI() abi.ABI
	RequiredGas(input []byte) uint64 // RequiredPrice calculates the contract gas use
}
