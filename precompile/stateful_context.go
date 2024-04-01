package precompile

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type statefulContext struct {
	state     StateDB
	address   common.Address
	msgSender common.Address
	msgValue  *big.Int
	readOnly  bool
}

func NewStatefulContext(
	state StateDB,
	address common.Address,
	msgSender common.Address,
	msgValue *big.Int,
) StatefulContext {
	return &statefulContext{
		state:     state,
		address:   address,
		msgSender: msgSender,
		msgValue:  msgValue,
		readOnly:  false,
	}
}

func (sc *statefulContext) SetState(key common.Hash, value common.Hash) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.SetState(sc.address, key, value)
	return nil
}

func (sc *statefulContext) GetState(key common.Hash) common.Hash {
	return sc.state.GetState(sc.address, key)
}

func (sc *statefulContext) SubBalance(address common.Address, amount *big.Int) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.SubBalance(address, amount)
	return nil
}

func (sc *statefulContext) AddBalance(address common.Address, amount *big.Int) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.AddBalance(address, amount)
	return nil
}

func (sc *statefulContext) GetBalance(address common.Address) *big.Int {
	return sc.state.GetBalance(address)
}

func (sc *statefulContext) Address() common.Address {
	return sc.address
}

func (sc *statefulContext) MsgSender() common.Address {
	return sc.msgSender
}

func (sc *statefulContext) MsgValue() *big.Int {
	return sc.msgValue
}

func (sc *statefulContext) IsReadOnly() bool {
	return sc.readOnly
}

func (sc *statefulContext) SetReadOnly(readOnly bool) {
	sc.readOnly = readOnly
}
