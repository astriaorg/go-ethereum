// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NativeMinterMetaData contains all meta data concerning the NativeMinter contract.
var NativeMinterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"newMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"}]",
}

// NativeMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeMinterMetaData.ABI instead.
var NativeMinterABI = NativeMinterMetaData.ABI

// NativeMinter is an auto generated Go binding around an Ethereum contract.
type NativeMinter struct {
	NativeMinterCaller     // Read-only binding to the contract
	NativeMinterTransactor // Write-only binding to the contract
	NativeMinterFilterer   // Log filterer for contract events
}

// NativeMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeMinterSession struct {
	Contract     *NativeMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeMinterCallerSession struct {
	Contract *NativeMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NativeMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeMinterTransactorSession struct {
	Contract     *NativeMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeMinterRaw struct {
	Contract *NativeMinter // Generic contract binding to access the raw methods on
}

// NativeMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeMinterCallerRaw struct {
	Contract *NativeMinterCaller // Generic read-only contract binding to access the raw methods on
}

// NativeMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeMinterTransactorRaw struct {
	Contract *NativeMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeMinter creates a new instance of NativeMinter, bound to a specific deployed contract.
func NewNativeMinter(address common.Address, backend bind.ContractBackend) (*NativeMinter, error) {
	contract, err := bindNativeMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeMinter{NativeMinterCaller: NativeMinterCaller{contract: contract}, NativeMinterTransactor: NativeMinterTransactor{contract: contract}, NativeMinterFilterer: NativeMinterFilterer{contract: contract}}, nil
}

// NewNativeMinterCaller creates a new read-only instance of NativeMinter, bound to a specific deployed contract.
func NewNativeMinterCaller(address common.Address, caller bind.ContractCaller) (*NativeMinterCaller, error) {
	contract, err := bindNativeMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMinterCaller{contract: contract}, nil
}

// NewNativeMinterTransactor creates a new write-only instance of NativeMinter, bound to a specific deployed contract.
func NewNativeMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeMinterTransactor, error) {
	contract, err := bindNativeMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMinterTransactor{contract: contract}, nil
}

// NewNativeMinterFilterer creates a new log filterer instance of NativeMinter, bound to a specific deployed contract.
func NewNativeMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeMinterFilterer, error) {
	contract, err := bindNativeMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeMinterFilterer{contract: contract}, nil
}

// bindNativeMinter binds a generic wrapper to an already deployed contract.
func bindNativeMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMinter *NativeMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMinter.Contract.NativeMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMinter *NativeMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMinter.Contract.NativeMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMinter *NativeMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMinter.Contract.NativeMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMinter *NativeMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMinter *NativeMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMinter *NativeMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMinter.Contract.contract.Transact(opts, method, params...)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NativeMinter *NativeMinterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeMinter.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NativeMinter *NativeMinterSession) Minter() (common.Address, error) {
	return _NativeMinter.Contract.Minter(&_NativeMinter.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NativeMinter *NativeMinterCallerSession) Minter() (common.Address, error) {
	return _NativeMinter.Contract.Minter(&_NativeMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeMinter *NativeMinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeMinter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeMinter *NativeMinterSession) Owner() (common.Address, error) {
	return _NativeMinter.Contract.Owner(&_NativeMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeMinter *NativeMinterCallerSession) Owner() (common.Address, error) {
	return _NativeMinter.Contract.Owner(&_NativeMinter.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterTransactor) Burn(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.contract.Transact(opts, "burn", addr, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterSession) Burn(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.Contract.Burn(&_NativeMinter.TransactOpts, addr, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterTransactorSession) Burn(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.Contract.Burn(&_NativeMinter.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.Contract.Mint(&_NativeMinter.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns(bool)
func (_NativeMinter *NativeMinterTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeMinter.Contract.Mint(&_NativeMinter.TransactOpts, addr, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns(bool)
func (_NativeMinter *NativeMinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMinter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns(bool)
func (_NativeMinter *NativeMinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeMinter.Contract.RenounceOwnership(&_NativeMinter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns(bool)
func (_NativeMinter *NativeMinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeMinter.Contract.RenounceOwnership(&_NativeMinter.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns(bool)
func (_NativeMinter *NativeMinterTransactor) SetMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _NativeMinter.contract.Transact(opts, "setMinter", newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns(bool)
func (_NativeMinter *NativeMinterSession) SetMinter(newMinter common.Address) (*types.Transaction, error) {
	return _NativeMinter.Contract.SetMinter(&_NativeMinter.TransactOpts, newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns(bool)
func (_NativeMinter *NativeMinterTransactorSession) SetMinter(newMinter common.Address) (*types.Transaction, error) {
	return _NativeMinter.Contract.SetMinter(&_NativeMinter.TransactOpts, newMinter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_NativeMinter *NativeMinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeMinter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_NativeMinter *NativeMinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeMinter.Contract.TransferOwnership(&_NativeMinter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_NativeMinter *NativeMinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeMinter.Contract.TransferOwnership(&_NativeMinter.TransactOpts, newOwner)
}
