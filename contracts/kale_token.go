// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// KaleTokenMetaData contains all meta data concerning the KaleToken contract.
var KaleTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"blacklist_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBlacklist\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklist\",\"type\":\"address\"}],\"name\":\"BlacklistChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"BlacklistStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blacklist_\",\"type\":\"address\"}],\"name\":\"setBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setBlacklistState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KaleTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use KaleTokenMetaData.ABI instead.
var KaleTokenABI = KaleTokenMetaData.ABI

// KaleToken is an auto generated Go binding around an Ethereum contract.
type KaleToken struct {
	KaleTokenCaller     // Read-only binding to the contract
	KaleTokenTransactor // Write-only binding to the contract
	KaleTokenFilterer   // Log filterer for contract events
}

// KaleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type KaleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KaleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KaleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KaleTokenSession struct {
	Contract     *KaleToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KaleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KaleTokenCallerSession struct {
	Contract *KaleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KaleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KaleTokenTransactorSession struct {
	Contract     *KaleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KaleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type KaleTokenRaw struct {
	Contract *KaleToken // Generic contract binding to access the raw methods on
}

// KaleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KaleTokenCallerRaw struct {
	Contract *KaleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// KaleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KaleTokenTransactorRaw struct {
	Contract *KaleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKaleToken creates a new instance of KaleToken, bound to a specific deployed contract.
func NewKaleToken(address common.Address, backend bind.ContractBackend) (*KaleToken, error) {
	contract, err := bindKaleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KaleToken{KaleTokenCaller: KaleTokenCaller{contract: contract}, KaleTokenTransactor: KaleTokenTransactor{contract: contract}, KaleTokenFilterer: KaleTokenFilterer{contract: contract}}, nil
}

// NewKaleTokenCaller creates a new read-only instance of KaleToken, bound to a specific deployed contract.
func NewKaleTokenCaller(address common.Address, caller bind.ContractCaller) (*KaleTokenCaller, error) {
	contract, err := bindKaleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KaleTokenCaller{contract: contract}, nil
}

// NewKaleTokenTransactor creates a new write-only instance of KaleToken, bound to a specific deployed contract.
func NewKaleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*KaleTokenTransactor, error) {
	contract, err := bindKaleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KaleTokenTransactor{contract: contract}, nil
}

// NewKaleTokenFilterer creates a new log filterer instance of KaleToken, bound to a specific deployed contract.
func NewKaleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*KaleTokenFilterer, error) {
	contract, err := bindKaleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KaleTokenFilterer{contract: contract}, nil
}

// bindKaleToken binds a generic wrapper to an already deployed contract.
func bindKaleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KaleTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KaleToken *KaleTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KaleToken.Contract.KaleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KaleToken *KaleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KaleToken.Contract.KaleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KaleToken *KaleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KaleToken.Contract.KaleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KaleToken *KaleTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KaleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KaleToken *KaleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KaleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KaleToken *KaleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KaleToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KaleToken *KaleTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KaleToken *KaleTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _KaleToken.Contract.DOMAINSEPARATOR(&_KaleToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KaleToken *KaleTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _KaleToken.Contract.DOMAINSEPARATOR(&_KaleToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KaleToken *KaleTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KaleToken *KaleTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KaleToken.Contract.Allowance(&_KaleToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KaleToken *KaleTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KaleToken.Contract.Allowance(&_KaleToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KaleToken *KaleTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KaleToken *KaleTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KaleToken.Contract.BalanceOf(&_KaleToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KaleToken *KaleTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KaleToken.Contract.BalanceOf(&_KaleToken.CallOpts, account)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_KaleToken *KaleTokenCaller) Blacklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "blacklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_KaleToken *KaleTokenSession) Blacklist() (common.Address, error) {
	return _KaleToken.Contract.Blacklist(&_KaleToken.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_KaleToken *KaleTokenCallerSession) Blacklist() (common.Address, error) {
	return _KaleToken.Contract.Blacklist(&_KaleToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KaleToken *KaleTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KaleToken *KaleTokenSession) Decimals() (uint8, error) {
	return _KaleToken.Contract.Decimals(&_KaleToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KaleToken *KaleTokenCallerSession) Decimals() (uint8, error) {
	return _KaleToken.Contract.Decimals(&_KaleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KaleToken *KaleTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KaleToken *KaleTokenSession) Name() (string, error) {
	return _KaleToken.Contract.Name(&_KaleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KaleToken *KaleTokenCallerSession) Name() (string, error) {
	return _KaleToken.Contract.Name(&_KaleToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KaleToken *KaleTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KaleToken *KaleTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _KaleToken.Contract.Nonces(&_KaleToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_KaleToken *KaleTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _KaleToken.Contract.Nonces(&_KaleToken.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KaleToken *KaleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KaleToken *KaleTokenSession) Owner() (common.Address, error) {
	return _KaleToken.Contract.Owner(&_KaleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KaleToken *KaleTokenCallerSession) Owner() (common.Address, error) {
	return _KaleToken.Contract.Owner(&_KaleToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KaleToken *KaleTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KaleToken *KaleTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KaleToken.Contract.SupportsInterface(&_KaleToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KaleToken *KaleTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KaleToken.Contract.SupportsInterface(&_KaleToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KaleToken *KaleTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KaleToken *KaleTokenSession) Symbol() (string, error) {
	return _KaleToken.Contract.Symbol(&_KaleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KaleToken *KaleTokenCallerSession) Symbol() (string, error) {
	return _KaleToken.Contract.Symbol(&_KaleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KaleToken *KaleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KaleToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KaleToken *KaleTokenSession) TotalSupply() (*big.Int, error) {
	return _KaleToken.Contract.TotalSupply(&_KaleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KaleToken *KaleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _KaleToken.Contract.TotalSupply(&_KaleToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.Approve(&_KaleToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.Approve(&_KaleToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KaleToken *KaleTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KaleToken *KaleTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.DecreaseAllowance(&_KaleToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KaleToken *KaleTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.DecreaseAllowance(&_KaleToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KaleToken *KaleTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KaleToken *KaleTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.IncreaseAllowance(&_KaleToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KaleToken *KaleTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.IncreaseAllowance(&_KaleToken.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KaleToken *KaleTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KaleToken *KaleTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KaleToken.Contract.Permit(&_KaleToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_KaleToken *KaleTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KaleToken.Contract.Permit(&_KaleToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KaleToken *KaleTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KaleToken *KaleTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _KaleToken.Contract.RenounceOwnership(&_KaleToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KaleToken *KaleTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KaleToken.Contract.RenounceOwnership(&_KaleToken.TransactOpts)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_KaleToken *KaleTokenTransactor) SetBlacklist(opts *bind.TransactOpts, blacklist_ common.Address) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "setBlacklist", blacklist_)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_KaleToken *KaleTokenSession) SetBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _KaleToken.Contract.SetBlacklist(&_KaleToken.TransactOpts, blacklist_)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_KaleToken *KaleTokenTransactorSession) SetBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _KaleToken.Contract.SetBlacklist(&_KaleToken.TransactOpts, blacklist_)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdb1954a.
//
// Solidity: function setBlacklistState(bool state) returns()
func (_KaleToken *KaleTokenTransactor) SetBlacklistState(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "setBlacklistState", state)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdb1954a.
//
// Solidity: function setBlacklistState(bool state) returns()
func (_KaleToken *KaleTokenSession) SetBlacklistState(state bool) (*types.Transaction, error) {
	return _KaleToken.Contract.SetBlacklistState(&_KaleToken.TransactOpts, state)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdb1954a.
//
// Solidity: function setBlacklistState(bool state) returns()
func (_KaleToken *KaleTokenTransactorSession) SetBlacklistState(state bool) (*types.Transaction, error) {
	return _KaleToken.Contract.SetBlacklistState(&_KaleToken.TransactOpts, state)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.Transfer(&_KaleToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.Transfer(&_KaleToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.TransferFrom(&_KaleToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KaleToken *KaleTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KaleToken.Contract.TransferFrom(&_KaleToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KaleToken *KaleTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KaleToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KaleToken *KaleTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KaleToken.Contract.TransferOwnership(&_KaleToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KaleToken *KaleTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KaleToken.Contract.TransferOwnership(&_KaleToken.TransactOpts, newOwner)
}

// KaleTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KaleToken contract.
type KaleTokenApprovalIterator struct {
	Event *KaleTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KaleTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaleTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KaleTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KaleTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaleTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaleTokenApproval represents a Approval event raised by the KaleToken contract.
type KaleTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KaleToken *KaleTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KaleTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KaleToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KaleTokenApprovalIterator{contract: _KaleToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KaleToken *KaleTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KaleTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KaleToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaleTokenApproval)
				if err := _KaleToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KaleToken *KaleTokenFilterer) ParseApproval(log types.Log) (*KaleTokenApproval, error) {
	event := new(KaleTokenApproval)
	if err := _KaleToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaleTokenBlacklistChangedIterator is returned from FilterBlacklistChanged and is used to iterate over the raw logs and unpacked data for BlacklistChanged events raised by the KaleToken contract.
type KaleTokenBlacklistChangedIterator struct {
	Event *KaleTokenBlacklistChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KaleTokenBlacklistChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaleTokenBlacklistChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KaleTokenBlacklistChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KaleTokenBlacklistChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaleTokenBlacklistChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaleTokenBlacklistChanged represents a BlacklistChanged event raised by the KaleToken contract.
type KaleTokenBlacklistChanged struct {
	OldBlacklist common.Address
	NewBlacklist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlacklistChanged is a free log retrieval operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_KaleToken *KaleTokenFilterer) FilterBlacklistChanged(opts *bind.FilterOpts, oldBlacklist []common.Address, newBlacklist []common.Address) (*KaleTokenBlacklistChangedIterator, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _KaleToken.contract.FilterLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return &KaleTokenBlacklistChangedIterator{contract: _KaleToken.contract, event: "BlacklistChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistChanged is a free log subscription operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_KaleToken *KaleTokenFilterer) WatchBlacklistChanged(opts *bind.WatchOpts, sink chan<- *KaleTokenBlacklistChanged, oldBlacklist []common.Address, newBlacklist []common.Address) (event.Subscription, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _KaleToken.contract.WatchLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaleTokenBlacklistChanged)
				if err := _KaleToken.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklistChanged is a log parse operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_KaleToken *KaleTokenFilterer) ParseBlacklistChanged(log types.Log) (*KaleTokenBlacklistChanged, error) {
	event := new(KaleTokenBlacklistChanged)
	if err := _KaleToken.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaleTokenBlacklistStatusChangedIterator is returned from FilterBlacklistStatusChanged and is used to iterate over the raw logs and unpacked data for BlacklistStatusChanged events raised by the KaleToken contract.
type KaleTokenBlacklistStatusChangedIterator struct {
	Event *KaleTokenBlacklistStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KaleTokenBlacklistStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaleTokenBlacklistStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KaleTokenBlacklistStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KaleTokenBlacklistStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaleTokenBlacklistStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaleTokenBlacklistStatusChanged represents a BlacklistStatusChanged event raised by the KaleToken contract.
type KaleTokenBlacklistStatusChanged struct {
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlacklistStatusChanged is a free log retrieval operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_KaleToken *KaleTokenFilterer) FilterBlacklistStatusChanged(opts *bind.FilterOpts, newStatus []bool) (*KaleTokenBlacklistStatusChangedIterator, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _KaleToken.contract.FilterLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return &KaleTokenBlacklistStatusChangedIterator{contract: _KaleToken.contract, event: "BlacklistStatusChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistStatusChanged is a free log subscription operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_KaleToken *KaleTokenFilterer) WatchBlacklistStatusChanged(opts *bind.WatchOpts, sink chan<- *KaleTokenBlacklistStatusChanged, newStatus []bool) (event.Subscription, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _KaleToken.contract.WatchLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaleTokenBlacklistStatusChanged)
				if err := _KaleToken.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklistStatusChanged is a log parse operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_KaleToken *KaleTokenFilterer) ParseBlacklistStatusChanged(log types.Log) (*KaleTokenBlacklistStatusChanged, error) {
	event := new(KaleTokenBlacklistStatusChanged)
	if err := _KaleToken.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaleTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KaleToken contract.
type KaleTokenOwnershipTransferredIterator struct {
	Event *KaleTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KaleTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaleTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KaleTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KaleTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaleTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaleTokenOwnershipTransferred represents a OwnershipTransferred event raised by the KaleToken contract.
type KaleTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KaleToken *KaleTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KaleTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KaleToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KaleTokenOwnershipTransferredIterator{contract: _KaleToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KaleToken *KaleTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KaleTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KaleToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaleTokenOwnershipTransferred)
				if err := _KaleToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KaleToken *KaleTokenFilterer) ParseOwnershipTransferred(log types.Log) (*KaleTokenOwnershipTransferred, error) {
	event := new(KaleTokenOwnershipTransferred)
	if err := _KaleToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KaleToken contract.
type KaleTokenTransferIterator struct {
	Event *KaleTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KaleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaleTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KaleTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KaleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaleTokenTransfer represents a Transfer event raised by the KaleToken contract.
type KaleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KaleToken *KaleTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KaleTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KaleToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KaleTokenTransferIterator{contract: _KaleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KaleToken *KaleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KaleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KaleToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaleTokenTransfer)
				if err := _KaleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KaleToken *KaleTokenFilterer) ParseTransfer(log types.Log) (*KaleTokenTransfer, error) {
	event := new(KaleTokenTransfer)
	if err := _KaleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
