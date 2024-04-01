package nativeminter

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/mocks"
)

var (
	zero = big.NewInt(0)

	SelfAddrForTest  = common.HexToAddress("0x1000")
	OwnerForTest     = common.BytesToAddress([]byte("0xOwner"))
	MinterForTest    = common.BytesToAddress([]byte("0xMinter"))
	RecipientForTest = common.BytesToAddress([]byte("0xRecipient"))
)

func TestInvalidMintAndBurn(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	minter := NewNativeMinter()

	recipient := RecipientForTest
	amount := big.NewInt(100)

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MinterForTest, amount)

	_, err := minter.Mint(ctx, recipient, amount)
	if err == nil {
		t.Fatalf("Expected error when minting, got nil")
	}

	ctx.AddBalance(recipient, amount)

	_, err = minter.Burn(ctx, recipient, amount)
	if err == nil {
		t.Fatalf("Expected error when minting, got nil")
	}
}

func TestValidMintAndBurn(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	minter := NewNativeMinter()

	recipient := RecipientForTest
	amount := big.NewInt(100)

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MinterForTest, amount)

	// set minter in state db for testing
	ctx.SetState(Slots.Minter, common.BytesToHash(MinterForTest.Bytes()))

	_, err := minter.Mint(ctx, recipient, amount)
	if err != nil {
		t.Fatalf("Minting failed with error: %v", err)
	}

	balance := stateDB.GetBalance(recipient)
	if balance.Cmp(amount) != 0 {
		t.Fatalf("Expected balance of %v, got %v", amount, balance)
	}

	_, err = minter.Burn(ctx, recipient, amount)
	if err != nil {
		t.Fatalf("Burning failed with error: %v", err)
	}

	balance = stateDB.GetBalance(recipient)
	if balance.Cmp(zero) != 0 {
		t.Fatalf("Expected balance of %v, got %v", zero, balance)
	}
}

func TestTransferOwnership(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	minter := NewNativeMinter()

	owner := OwnerForTest
	newOwner := common.BytesToAddress([]byte("0xNewOwner"))

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, owner, big.NewInt(0))

	_, err := minter.TransferOwnership(ctx, newOwner)
	if err == nil {
		t.Fatalf("Expected error when transferring ownership, got nil")
	}

	// set owner in state db for testing
	ctx.SetState(Slots.Owner, common.BytesToHash(owner.Bytes()))

	// check owner is set
	setOwner, err := minter.Owner(ctx)
	if err != nil {
		t.Fatalf("Owner failed with error: %v", err)
	}
	if setOwner.Cmp(owner) != 0 {
		t.Fatalf("Expected owner to be %v, got %v", owner, setOwner)
	}

	_, err = minter.TransferOwnership(ctx, newOwner)
	if err != nil {
		t.Fatalf("TransferOwnership failed with error: %v", err)
	}

	// Check if the new owner is set correctly
	setOwner, err = minter.Owner(ctx)
	if err != nil {
		t.Fatalf("Owner failed with error: %v", err)
	}
	if setOwner.Cmp(newOwner) != 0 {
		t.Fatalf("Expected new owner to be %v, got %v", newOwner, setOwner)
	}
}

func TestRenounceOwnership(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	minter := NewNativeMinter()

	owner := OwnerForTest

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, owner, big.NewInt(0))

	_, err := minter.RenounceOwnership(ctx)
	if err == nil {
		t.Fatalf("Expected error when renouncing ownership, got nil")
	}

	// set owner in state db for testing
	ctx.SetState(Slots.Owner, common.BytesToHash(owner.Bytes()))

	_, err = minter.RenounceOwnership(ctx)
	if err != nil {
		t.Fatalf("RenounceOwnership failed with error: %v", err)
	}

	// Check if the owner is set to zero address after renouncing ownership
	setOwner, err := minter.Owner(ctx)
	if err != nil {
		t.Fatalf("Owner failed with error: %v", err)
	}
	if setOwner.Cmp(ZeroAddress) != 0 {
		t.Fatalf("Expected owner to be zero address after renouncing ownership, got %v", setOwner)
	}
}

func TestSetMinter(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	minter := NewNativeMinter()

	owner := OwnerForTest
	newMinter := MinterForTest

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, owner, big.NewInt(0))

	// check current minter is zero address (not set)
	setMinter, err := minter.Minter(ctx)
	if err != nil {
		t.Fatalf("Owner failed with error: %v", err)
	}
	if setMinter.Cmp(ZeroAddress) != 0 {
		t.Fatalf("Expected minter to be zero address, got %v", setMinter)
	}

	// should fail as owner has not been set
	_, err = minter.SetMinter(ctx, newMinter)
	if err == nil {
		t.Fatalf("Expected error when renouncing ownership, got nil")
	}

	// set owner in state db for testing
	ctx.SetState(Slots.Owner, common.BytesToHash(owner.Bytes()))

	_, err = minter.SetMinter(ctx, newMinter)
	if err != nil {
		t.Fatalf("SetMinter failed with error: %v", err)
	}

	// Check if the minter has been properly set
	setMinter, err = minter.Minter(ctx)
	if err != nil {
		t.Fatalf("Owner failed with error: %v", err)
	}
	if setMinter.Cmp(newMinter) != 0 {
		t.Fatalf("Expected minter to be %v, got %v", newMinter, setMinter)
	}
}
