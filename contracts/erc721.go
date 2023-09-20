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

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"CannotRevokeSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"FunctionDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotMetadataUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter_\",\"type\":\"address\"}],\"name\":\"NotMinterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotRoyaltySetterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotURIUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBlacklist\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklist\",\"type\":\"address\"}],\"name\":\"BlacklistChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"BlacklistStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METADATA_UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROYALTY_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"URI_UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMetadataUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addRoyaltySetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addUriUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blacklist_\",\"type\":\"address\"}],\"name\":\"changeBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"changeBlacklistState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"mintBatchTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"_tokenIds\",\"type\":\"uint256[][]\"}],\"name\":\"mintBatchToMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"notifyBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"notifyMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"notifyMetadataUpdateForMultipleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"resetTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"retrieveMultipleBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"retriveBatchOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeMetadataUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"revokeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"revokeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRoyaltySetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeUriUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royalty\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royalty\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721.Contract.DEFAULTADMINROLE(&_ERC721.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721.Contract.DEFAULTADMINROLE(&_ERC721.CallOpts)
}

// METADATAUPDATERROLE is a free data retrieval call binding the contract method 0x89025750.
//
// Solidity: function METADATA_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) METADATAUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "METADATA_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// METADATAUPDATERROLE is a free data retrieval call binding the contract method 0x89025750.
//
// Solidity: function METADATA_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) METADATAUPDATERROLE() ([32]byte, error) {
	return _ERC721.Contract.METADATAUPDATERROLE(&_ERC721.CallOpts)
}

// METADATAUPDATERROLE is a free data retrieval call binding the contract method 0x89025750.
//
// Solidity: function METADATA_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) METADATAUPDATERROLE() ([32]byte, error) {
	return _ERC721.Contract.METADATAUPDATERROLE(&_ERC721.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) MINTERROLE() ([32]byte, error) {
	return _ERC721.Contract.MINTERROLE(&_ERC721.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC721.Contract.MINTERROLE(&_ERC721.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) OWNERROLE() ([32]byte, error) {
	return _ERC721.Contract.OWNERROLE(&_ERC721.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) OWNERROLE() ([32]byte, error) {
	return _ERC721.Contract.OWNERROLE(&_ERC721.CallOpts)
}

// ROYALTYSETTERROLE is a free data retrieval call binding the contract method 0xf44acff9.
//
// Solidity: function ROYALTY_SETTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) ROYALTYSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ROYALTY_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROYALTYSETTERROLE is a free data retrieval call binding the contract method 0xf44acff9.
//
// Solidity: function ROYALTY_SETTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) ROYALTYSETTERROLE() ([32]byte, error) {
	return _ERC721.Contract.ROYALTYSETTERROLE(&_ERC721.CallOpts)
}

// ROYALTYSETTERROLE is a free data retrieval call binding the contract method 0xf44acff9.
//
// Solidity: function ROYALTY_SETTER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) ROYALTYSETTERROLE() ([32]byte, error) {
	return _ERC721.Contract.ROYALTYSETTERROLE(&_ERC721.CallOpts)
}

// URIUPDATERROLE is a free data retrieval call binding the contract method 0xb79c2e56.
//
// Solidity: function URI_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Caller) URIUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "URI_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// URIUPDATERROLE is a free data retrieval call binding the contract method 0xb79c2e56.
//
// Solidity: function URI_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721Session) URIUPDATERROLE() ([32]byte, error) {
	return _ERC721.Contract.URIUPDATERROLE(&_ERC721.CallOpts)
}

// URIUPDATERROLE is a free data retrieval call binding the contract method 0xb79c2e56.
//
// Solidity: function URI_UPDATER_ROLE() view returns(bytes32)
func (_ERC721 *ERC721CallerSession) URIUPDATERROLE() ([32]byte, error) {
	return _ERC721.Contract.URIUPDATERROLE(&_ERC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_ERC721 *ERC721Caller) Blacklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "blacklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_ERC721 *ERC721Session) Blacklist() (common.Address, error) {
	return _ERC721.Contract.Blacklist(&_ERC721.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_ERC721 *ERC721CallerSession) Blacklist() (common.Address, error) {
	return _ERC721.Contract.Blacklist(&_ERC721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721 *ERC721Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721 *ERC721Session) ContractURI() (string, error) {
	return _ERC721.Contract.ContractURI(&_ERC721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721 *ERC721CallerSession) ContractURI() (string, error) {
	return _ERC721.Contract.ContractURI(&_ERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721 *ERC721Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721 *ERC721Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721.Contract.GetRoleAdmin(&_ERC721.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721 *ERC721CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721.Contract.GetRoleAdmin(&_ERC721.CallOpts, role)
}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Caller) GrantRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "grantRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Session) GrantRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.GrantRole(&_ERC721.CallOpts, arg0, arg1)
}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721CallerSession) GrantRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.GrantRole(&_ERC721.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721 *ERC721Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721 *ERC721Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721.Contract.HasRole(&_ERC721.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721 *ERC721CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721.Contract.HasRole(&_ERC721.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Caller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Session) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.RenounceRole(&_ERC721.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721CallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.RenounceRole(&_ERC721.CallOpts, arg0, arg1)
}

// RetrieveMultipleBalances is a free data retrieval call binding the contract method 0xc77449f6.
//
// Solidity: function retrieveMultipleBalances(address[] accounts) view returns(uint256[] balances)
func (_ERC721 *ERC721Caller) RetrieveMultipleBalances(opts *bind.CallOpts, accounts []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "retrieveMultipleBalances", accounts)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RetrieveMultipleBalances is a free data retrieval call binding the contract method 0xc77449f6.
//
// Solidity: function retrieveMultipleBalances(address[] accounts) view returns(uint256[] balances)
func (_ERC721 *ERC721Session) RetrieveMultipleBalances(accounts []common.Address) ([]*big.Int, error) {
	return _ERC721.Contract.RetrieveMultipleBalances(&_ERC721.CallOpts, accounts)
}

// RetrieveMultipleBalances is a free data retrieval call binding the contract method 0xc77449f6.
//
// Solidity: function retrieveMultipleBalances(address[] accounts) view returns(uint256[] balances)
func (_ERC721 *ERC721CallerSession) RetrieveMultipleBalances(accounts []common.Address) ([]*big.Int, error) {
	return _ERC721.Contract.RetrieveMultipleBalances(&_ERC721.CallOpts, accounts)
}

// RetriveBatchOwners is a free data retrieval call binding the contract method 0x0dc9c4fb.
//
// Solidity: function retriveBatchOwners(uint256[] tokenIds) view returns(address[] owners)
func (_ERC721 *ERC721Caller) RetriveBatchOwners(opts *bind.CallOpts, tokenIds []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "retriveBatchOwners", tokenIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RetriveBatchOwners is a free data retrieval call binding the contract method 0x0dc9c4fb.
//
// Solidity: function retriveBatchOwners(uint256[] tokenIds) view returns(address[] owners)
func (_ERC721 *ERC721Session) RetriveBatchOwners(tokenIds []*big.Int) ([]common.Address, error) {
	return _ERC721.Contract.RetriveBatchOwners(&_ERC721.CallOpts, tokenIds)
}

// RetriveBatchOwners is a free data retrieval call binding the contract method 0x0dc9c4fb.
//
// Solidity: function retriveBatchOwners(uint256[] tokenIds) view returns(address[] owners)
func (_ERC721 *ERC721CallerSession) RetriveBatchOwners(tokenIds []*big.Int) ([]common.Address, error) {
	return _ERC721.Contract.RetriveBatchOwners(&_ERC721.CallOpts, tokenIds)
}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Caller) RevokeRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "revokeRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721Session) RevokeRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.RevokeRole(&_ERC721.CallOpts, arg0, arg1)
}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_ERC721 *ERC721CallerSession) RevokeRole(arg0 [32]byte, arg1 common.Address) error {
	return _ERC721.Contract.RevokeRole(&_ERC721.CallOpts, arg0, arg1)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ERC721 *ERC721Caller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ERC721 *ERC721Session) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC721.Contract.RoyaltyInfo(&_ERC721.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ERC721 *ERC721CallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC721.Contract.RoyaltyInfo(&_ERC721.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// AddMetadataUpdater is a paid mutator transaction binding the contract method 0x18049434.
//
// Solidity: function addMetadataUpdater(address account) returns()
func (_ERC721 *ERC721Transactor) AddMetadataUpdater(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "addMetadataUpdater", account)
}

// AddMetadataUpdater is a paid mutator transaction binding the contract method 0x18049434.
//
// Solidity: function addMetadataUpdater(address account) returns()
func (_ERC721 *ERC721Session) AddMetadataUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddMetadataUpdater(&_ERC721.TransactOpts, account)
}

// AddMetadataUpdater is a paid mutator transaction binding the contract method 0x18049434.
//
// Solidity: function addMetadataUpdater(address account) returns()
func (_ERC721 *ERC721TransactorSession) AddMetadataUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddMetadataUpdater(&_ERC721.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ERC721 *ERC721Transactor) AddMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "addMinter", _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ERC721 *ERC721Session) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddMinter(&_ERC721.TransactOpts, _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ERC721 *ERC721TransactorSession) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddMinter(&_ERC721.TransactOpts, _minter)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner_) returns()
func (_ERC721 *ERC721Transactor) AddOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "addOwner", owner_)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner_) returns()
func (_ERC721 *ERC721Session) AddOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddOwner(&_ERC721.TransactOpts, owner_)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner_) returns()
func (_ERC721 *ERC721TransactorSession) AddOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddOwner(&_ERC721.TransactOpts, owner_)
}

// AddRoyaltySetter is a paid mutator transaction binding the contract method 0x5418cd14.
//
// Solidity: function addRoyaltySetter(address account) returns()
func (_ERC721 *ERC721Transactor) AddRoyaltySetter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "addRoyaltySetter", account)
}

// AddRoyaltySetter is a paid mutator transaction binding the contract method 0x5418cd14.
//
// Solidity: function addRoyaltySetter(address account) returns()
func (_ERC721 *ERC721Session) AddRoyaltySetter(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddRoyaltySetter(&_ERC721.TransactOpts, account)
}

// AddRoyaltySetter is a paid mutator transaction binding the contract method 0x5418cd14.
//
// Solidity: function addRoyaltySetter(address account) returns()
func (_ERC721 *ERC721TransactorSession) AddRoyaltySetter(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddRoyaltySetter(&_ERC721.TransactOpts, account)
}

// AddUriUpdater is a paid mutator transaction binding the contract method 0xab86bfc9.
//
// Solidity: function addUriUpdater(address account) returns()
func (_ERC721 *ERC721Transactor) AddUriUpdater(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "addUriUpdater", account)
}

// AddUriUpdater is a paid mutator transaction binding the contract method 0xab86bfc9.
//
// Solidity: function addUriUpdater(address account) returns()
func (_ERC721 *ERC721Session) AddUriUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddUriUpdater(&_ERC721.TransactOpts, account)
}

// AddUriUpdater is a paid mutator transaction binding the contract method 0xab86bfc9.
//
// Solidity: function addUriUpdater(address account) returns()
func (_ERC721 *ERC721TransactorSession) AddUriUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.AddUriUpdater(&_ERC721.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// ChangeBlacklist is a paid mutator transaction binding the contract method 0x53735dae.
//
// Solidity: function changeBlacklist(address blacklist_) returns()
func (_ERC721 *ERC721Transactor) ChangeBlacklist(opts *bind.TransactOpts, blacklist_ common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "changeBlacklist", blacklist_)
}

// ChangeBlacklist is a paid mutator transaction binding the contract method 0x53735dae.
//
// Solidity: function changeBlacklist(address blacklist_) returns()
func (_ERC721 *ERC721Session) ChangeBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.ChangeBlacklist(&_ERC721.TransactOpts, blacklist_)
}

// ChangeBlacklist is a paid mutator transaction binding the contract method 0x53735dae.
//
// Solidity: function changeBlacklist(address blacklist_) returns()
func (_ERC721 *ERC721TransactorSession) ChangeBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.ChangeBlacklist(&_ERC721.TransactOpts, blacklist_)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_ERC721 *ERC721Transactor) ChangeBlacklistState(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "changeBlacklistState", state)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_ERC721 *ERC721Session) ChangeBlacklistState(state bool) (*types.Transaction, error) {
	return _ERC721.Contract.ChangeBlacklistState(&_ERC721.TransactOpts, state)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_ERC721 *ERC721TransactorSession) ChangeBlacklistState(state bool) (*types.Transaction, error) {
	return _ERC721.Contract.ChangeBlacklistState(&_ERC721.TransactOpts, state)
}

// MintBatchTo is a paid mutator transaction binding the contract method 0xe5b0c237.
//
// Solidity: function mintBatchTo(address _to, uint256[] _tokenIds) returns()
func (_ERC721 *ERC721Transactor) MintBatchTo(opts *bind.TransactOpts, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "mintBatchTo", _to, _tokenIds)
}

// MintBatchTo is a paid mutator transaction binding the contract method 0xe5b0c237.
//
// Solidity: function mintBatchTo(address _to, uint256[] _tokenIds) returns()
func (_ERC721 *ERC721Session) MintBatchTo(_to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintBatchTo(&_ERC721.TransactOpts, _to, _tokenIds)
}

// MintBatchTo is a paid mutator transaction binding the contract method 0xe5b0c237.
//
// Solidity: function mintBatchTo(address _to, uint256[] _tokenIds) returns()
func (_ERC721 *ERC721TransactorSession) MintBatchTo(_to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintBatchTo(&_ERC721.TransactOpts, _to, _tokenIds)
}

// MintBatchToMultiple is a paid mutator transaction binding the contract method 0x03f8e34a.
//
// Solidity: function mintBatchToMultiple(address[] _to, uint256[][] _tokenIds) returns()
func (_ERC721 *ERC721Transactor) MintBatchToMultiple(opts *bind.TransactOpts, _to []common.Address, _tokenIds [][]*big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "mintBatchToMultiple", _to, _tokenIds)
}

// MintBatchToMultiple is a paid mutator transaction binding the contract method 0x03f8e34a.
//
// Solidity: function mintBatchToMultiple(address[] _to, uint256[][] _tokenIds) returns()
func (_ERC721 *ERC721Session) MintBatchToMultiple(_to []common.Address, _tokenIds [][]*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintBatchToMultiple(&_ERC721.TransactOpts, _to, _tokenIds)
}

// MintBatchToMultiple is a paid mutator transaction binding the contract method 0x03f8e34a.
//
// Solidity: function mintBatchToMultiple(address[] _to, uint256[][] _tokenIds) returns()
func (_ERC721 *ERC721TransactorSession) MintBatchToMultiple(_to []common.Address, _tokenIds [][]*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintBatchToMultiple(&_ERC721.TransactOpts, _to, _tokenIds)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) MintTo(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "mintTo", _to, _tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) MintTo(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintTo(&_ERC721.TransactOpts, _to, _tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) MintTo(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.MintTo(&_ERC721.TransactOpts, _to, _tokenId)
}

// NotifyBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x73320902.
//
// Solidity: function notifyBatchMetadataUpdate(uint256 from, uint256 to) returns()
func (_ERC721 *ERC721Transactor) NotifyBatchMetadataUpdate(opts *bind.TransactOpts, from *big.Int, to *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "notifyBatchMetadataUpdate", from, to)
}

// NotifyBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x73320902.
//
// Solidity: function notifyBatchMetadataUpdate(uint256 from, uint256 to) returns()
func (_ERC721 *ERC721Session) NotifyBatchMetadataUpdate(from *big.Int, to *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyBatchMetadataUpdate(&_ERC721.TransactOpts, from, to)
}

// NotifyBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x73320902.
//
// Solidity: function notifyBatchMetadataUpdate(uint256 from, uint256 to) returns()
func (_ERC721 *ERC721TransactorSession) NotifyBatchMetadataUpdate(from *big.Int, to *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyBatchMetadataUpdate(&_ERC721.TransactOpts, from, to)
}

// NotifyMetadataUpdate is a paid mutator transaction binding the contract method 0x34555d7e.
//
// Solidity: function notifyMetadataUpdate(uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) NotifyMetadataUpdate(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "notifyMetadataUpdate", tokenId)
}

// NotifyMetadataUpdate is a paid mutator transaction binding the contract method 0x34555d7e.
//
// Solidity: function notifyMetadataUpdate(uint256 tokenId) returns()
func (_ERC721 *ERC721Session) NotifyMetadataUpdate(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyMetadataUpdate(&_ERC721.TransactOpts, tokenId)
}

// NotifyMetadataUpdate is a paid mutator transaction binding the contract method 0x34555d7e.
//
// Solidity: function notifyMetadataUpdate(uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) NotifyMetadataUpdate(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyMetadataUpdate(&_ERC721.TransactOpts, tokenId)
}

// NotifyMetadataUpdateForMultipleTokens is a paid mutator transaction binding the contract method 0x33f53073.
//
// Solidity: function notifyMetadataUpdateForMultipleTokens(uint256[] tokenIds) returns()
func (_ERC721 *ERC721Transactor) NotifyMetadataUpdateForMultipleTokens(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "notifyMetadataUpdateForMultipleTokens", tokenIds)
}

// NotifyMetadataUpdateForMultipleTokens is a paid mutator transaction binding the contract method 0x33f53073.
//
// Solidity: function notifyMetadataUpdateForMultipleTokens(uint256[] tokenIds) returns()
func (_ERC721 *ERC721Session) NotifyMetadataUpdateForMultipleTokens(tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyMetadataUpdateForMultipleTokens(&_ERC721.TransactOpts, tokenIds)
}

// NotifyMetadataUpdateForMultipleTokens is a paid mutator transaction binding the contract method 0x33f53073.
//
// Solidity: function notifyMetadataUpdateForMultipleTokens(uint256[] tokenIds) returns()
func (_ERC721 *ERC721TransactorSession) NotifyMetadataUpdateForMultipleTokens(tokenIds []*big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.NotifyMetadataUpdateForMultipleTokens(&_ERC721.TransactOpts, tokenIds)
}

// ResetDefaultRoyalty is a paid mutator transaction binding the contract method 0xee437dae.
//
// Solidity: function resetDefaultRoyalty() returns()
func (_ERC721 *ERC721Transactor) ResetDefaultRoyalty(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "resetDefaultRoyalty")
}

// ResetDefaultRoyalty is a paid mutator transaction binding the contract method 0xee437dae.
//
// Solidity: function resetDefaultRoyalty() returns()
func (_ERC721 *ERC721Session) ResetDefaultRoyalty() (*types.Transaction, error) {
	return _ERC721.Contract.ResetDefaultRoyalty(&_ERC721.TransactOpts)
}

// ResetDefaultRoyalty is a paid mutator transaction binding the contract method 0xee437dae.
//
// Solidity: function resetDefaultRoyalty() returns()
func (_ERC721 *ERC721TransactorSession) ResetDefaultRoyalty() (*types.Transaction, error) {
	return _ERC721.Contract.ResetDefaultRoyalty(&_ERC721.TransactOpts)
}

// ResetTokenRoyalty is a paid mutator transaction binding the contract method 0x8a616bc0.
//
// Solidity: function resetTokenRoyalty(uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) ResetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "resetTokenRoyalty", tokenId)
}

// ResetTokenRoyalty is a paid mutator transaction binding the contract method 0x8a616bc0.
//
// Solidity: function resetTokenRoyalty(uint256 tokenId) returns()
func (_ERC721 *ERC721Session) ResetTokenRoyalty(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.ResetTokenRoyalty(&_ERC721.TransactOpts, tokenId)
}

// ResetTokenRoyalty is a paid mutator transaction binding the contract method 0x8a616bc0.
//
// Solidity: function resetTokenRoyalty(uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) ResetTokenRoyalty(tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.ResetTokenRoyalty(&_ERC721.TransactOpts, tokenId)
}

// RevokeMetadataUpdater is a paid mutator transaction binding the contract method 0x75cde755.
//
// Solidity: function revokeMetadataUpdater(address account) returns()
func (_ERC721 *ERC721Transactor) RevokeMetadataUpdater(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "revokeMetadataUpdater", account)
}

// RevokeMetadataUpdater is a paid mutator transaction binding the contract method 0x75cde755.
//
// Solidity: function revokeMetadataUpdater(address account) returns()
func (_ERC721 *ERC721Session) RevokeMetadataUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeMetadataUpdater(&_ERC721.TransactOpts, account)
}

// RevokeMetadataUpdater is a paid mutator transaction binding the contract method 0x75cde755.
//
// Solidity: function revokeMetadataUpdater(address account) returns()
func (_ERC721 *ERC721TransactorSession) RevokeMetadataUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeMetadataUpdater(&_ERC721.TransactOpts, account)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _minter) returns()
func (_ERC721 *ERC721Transactor) RevokeMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "revokeMinter", _minter)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _minter) returns()
func (_ERC721 *ERC721Session) RevokeMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeMinter(&_ERC721.TransactOpts, _minter)
}

// RevokeMinter is a paid mutator transaction binding the contract method 0xcfbd4885.
//
// Solidity: function revokeMinter(address _minter) returns()
func (_ERC721 *ERC721TransactorSession) RevokeMinter(_minter common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeMinter(&_ERC721.TransactOpts, _minter)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address owner_) returns()
func (_ERC721 *ERC721Transactor) RevokeOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "revokeOwner", owner_)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address owner_) returns()
func (_ERC721 *ERC721Session) RevokeOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeOwner(&_ERC721.TransactOpts, owner_)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address owner_) returns()
func (_ERC721 *ERC721TransactorSession) RevokeOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeOwner(&_ERC721.TransactOpts, owner_)
}

// RevokeRoyaltySetter is a paid mutator transaction binding the contract method 0xf470b0c3.
//
// Solidity: function revokeRoyaltySetter(address account) returns()
func (_ERC721 *ERC721Transactor) RevokeRoyaltySetter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "revokeRoyaltySetter", account)
}

// RevokeRoyaltySetter is a paid mutator transaction binding the contract method 0xf470b0c3.
//
// Solidity: function revokeRoyaltySetter(address account) returns()
func (_ERC721 *ERC721Session) RevokeRoyaltySetter(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeRoyaltySetter(&_ERC721.TransactOpts, account)
}

// RevokeRoyaltySetter is a paid mutator transaction binding the contract method 0xf470b0c3.
//
// Solidity: function revokeRoyaltySetter(address account) returns()
func (_ERC721 *ERC721TransactorSession) RevokeRoyaltySetter(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeRoyaltySetter(&_ERC721.TransactOpts, account)
}

// RevokeUriUpdater is a paid mutator transaction binding the contract method 0xba528fa9.
//
// Solidity: function revokeUriUpdater(address account) returns()
func (_ERC721 *ERC721Transactor) RevokeUriUpdater(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "revokeUriUpdater", account)
}

// RevokeUriUpdater is a paid mutator transaction binding the contract method 0xba528fa9.
//
// Solidity: function revokeUriUpdater(address account) returns()
func (_ERC721 *ERC721Session) RevokeUriUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeUriUpdater(&_ERC721.TransactOpts, account)
}

// RevokeUriUpdater is a paid mutator transaction binding the contract method 0xba528fa9.
//
// Solidity: function revokeUriUpdater(address account) returns()
func (_ERC721 *ERC721TransactorSession) RevokeUriUpdater(account common.Address) (*types.Transaction, error) {
	return _ERC721.Contract.RevokeUriUpdater(&_ERC721.TransactOpts, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri_) returns()
func (_ERC721 *ERC721Transactor) SetContractURI(opts *bind.TransactOpts, uri_ string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setContractURI", uri_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri_) returns()
func (_ERC721 *ERC721Session) SetContractURI(uri_ string) (*types.Transaction, error) {
	return _ERC721.Contract.SetContractURI(&_ERC721.TransactOpts, uri_)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri_) returns()
func (_ERC721 *ERC721TransactorSession) SetContractURI(uri_ string) (*types.Transaction, error) {
	return _ERC721.Contract.SetContractURI(&_ERC721.TransactOpts, uri_)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721Transactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setDefaultRoyalty", receiver, royalty)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721Session) SetDefaultRoyalty(receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetDefaultRoyalty(&_ERC721.TransactOpts, receiver, royalty)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721TransactorSession) SetDefaultRoyalty(receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetDefaultRoyalty(&_ERC721.TransactOpts, receiver, royalty)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721Transactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, royalty)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721Session) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetTokenRoyalty(&_ERC721.TransactOpts, tokenId, receiver, royalty)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 royalty) returns()
func (_ERC721 *ERC721TransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, royalty *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetTokenRoyalty(&_ERC721.TransactOpts, tokenId, receiver, royalty)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x09a3beef.
//
// Solidity: function setTokenURI(string uri_, uint256 totalTokens) returns()
func (_ERC721 *ERC721Transactor) SetTokenURI(opts *bind.TransactOpts, uri_ string, totalTokens *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setTokenURI", uri_, totalTokens)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x09a3beef.
//
// Solidity: function setTokenURI(string uri_, uint256 totalTokens) returns()
func (_ERC721 *ERC721Session) SetTokenURI(uri_ string, totalTokens *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetTokenURI(&_ERC721.TransactOpts, uri_, totalTokens)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x09a3beef.
//
// Solidity: function setTokenURI(string uri_, uint256 totalTokens) returns()
func (_ERC721 *ERC721TransactorSession) SetTokenURI(uri_ string, totalTokens *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SetTokenURI(&_ERC721.TransactOpts, uri_, totalTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ERC721 contract.
type ERC721BatchMetadataUpdateIterator struct {
	Event *ERC721BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC721BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BatchMetadataUpdate)
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
		it.Event = new(ERC721BatchMetadataUpdate)
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
func (it *ERC721BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ERC721 contract.
type ERC721BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721 *ERC721Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ERC721BatchMetadataUpdateIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC721BatchMetadataUpdateIterator{contract: _ERC721.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721 *ERC721Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC721BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BatchMetadataUpdate)
				if err := _ERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721 *ERC721Filterer) ParseBatchMetadataUpdate(log types.Log) (*ERC721BatchMetadataUpdate, error) {
	event := new(ERC721BatchMetadataUpdate)
	if err := _ERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721BlacklistChangedIterator is returned from FilterBlacklistChanged and is used to iterate over the raw logs and unpacked data for BlacklistChanged events raised by the ERC721 contract.
type ERC721BlacklistChangedIterator struct {
	Event *ERC721BlacklistChanged // Event containing the contract specifics and raw log

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
func (it *ERC721BlacklistChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BlacklistChanged)
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
		it.Event = new(ERC721BlacklistChanged)
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
func (it *ERC721BlacklistChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BlacklistChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BlacklistChanged represents a BlacklistChanged event raised by the ERC721 contract.
type ERC721BlacklistChanged struct {
	OldBlacklist common.Address
	NewBlacklist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlacklistChanged is a free log retrieval operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_ERC721 *ERC721Filterer) FilterBlacklistChanged(opts *bind.FilterOpts, oldBlacklist []common.Address, newBlacklist []common.Address) (*ERC721BlacklistChangedIterator, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BlacklistChangedIterator{contract: _ERC721.contract, event: "BlacklistChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistChanged is a free log subscription operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_ERC721 *ERC721Filterer) WatchBlacklistChanged(opts *bind.WatchOpts, sink chan<- *ERC721BlacklistChanged, oldBlacklist []common.Address, newBlacklist []common.Address) (event.Subscription, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BlacklistChanged)
				if err := _ERC721.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseBlacklistChanged(log types.Log) (*ERC721BlacklistChanged, error) {
	event := new(ERC721BlacklistChanged)
	if err := _ERC721.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721BlacklistStatusChangedIterator is returned from FilterBlacklistStatusChanged and is used to iterate over the raw logs and unpacked data for BlacklistStatusChanged events raised by the ERC721 contract.
type ERC721BlacklistStatusChangedIterator struct {
	Event *ERC721BlacklistStatusChanged // Event containing the contract specifics and raw log

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
func (it *ERC721BlacklistStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BlacklistStatusChanged)
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
		it.Event = new(ERC721BlacklistStatusChanged)
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
func (it *ERC721BlacklistStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BlacklistStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BlacklistStatusChanged represents a BlacklistStatusChanged event raised by the ERC721 contract.
type ERC721BlacklistStatusChanged struct {
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlacklistStatusChanged is a free log retrieval operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_ERC721 *ERC721Filterer) FilterBlacklistStatusChanged(opts *bind.FilterOpts, newStatus []bool) (*ERC721BlacklistStatusChangedIterator, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BlacklistStatusChangedIterator{contract: _ERC721.contract, event: "BlacklistStatusChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistStatusChanged is a free log subscription operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_ERC721 *ERC721Filterer) WatchBlacklistStatusChanged(opts *bind.WatchOpts, sink chan<- *ERC721BlacklistStatusChanged, newStatus []bool) (event.Subscription, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BlacklistStatusChanged)
				if err := _ERC721.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseBlacklistStatusChanged(log types.Log) (*ERC721BlacklistStatusChanged, error) {
	event := new(ERC721BlacklistStatusChanged)
	if err := _ERC721.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ContractURIUpdateIterator is returned from FilterContractURIUpdate and is used to iterate over the raw logs and unpacked data for ContractURIUpdate events raised by the ERC721 contract.
type ERC721ContractURIUpdateIterator struct {
	Event *ERC721ContractURIUpdate // Event containing the contract specifics and raw log

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
func (it *ERC721ContractURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractURIUpdate)
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
		it.Event = new(ERC721ContractURIUpdate)
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
func (it *ERC721ContractURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractURIUpdate represents a ContractURIUpdate event raised by the ERC721 contract.
type ERC721ContractURIUpdate struct {
	NewContractURI string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdate is a free log retrieval operation binding the contract event 0x5a381feb44d0e107829cd2fcde8ec87ca91ae94b43e505b14c36b31a06025135.
//
// Solidity: event ContractURIUpdate(string newContractURI)
func (_ERC721 *ERC721Filterer) FilterContractURIUpdate(opts *bind.FilterOpts) (*ERC721ContractURIUpdateIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ContractURIUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC721ContractURIUpdateIterator{contract: _ERC721.contract, event: "ContractURIUpdate", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdate is a free log subscription operation binding the contract event 0x5a381feb44d0e107829cd2fcde8ec87ca91ae94b43e505b14c36b31a06025135.
//
// Solidity: event ContractURIUpdate(string newContractURI)
func (_ERC721 *ERC721Filterer) WatchContractURIUpdate(opts *bind.WatchOpts, sink chan<- *ERC721ContractURIUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ContractURIUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractURIUpdate)
				if err := _ERC721.contract.UnpackLog(event, "ContractURIUpdate", log); err != nil {
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

// ParseContractURIUpdate is a log parse operation binding the contract event 0x5a381feb44d0e107829cd2fcde8ec87ca91ae94b43e505b14c36b31a06025135.
//
// Solidity: event ContractURIUpdate(string newContractURI)
func (_ERC721 *ERC721Filterer) ParseContractURIUpdate(log types.Log) (*ERC721ContractURIUpdate, error) {
	event := new(ERC721ContractURIUpdate)
	if err := _ERC721.contract.UnpackLog(event, "ContractURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ERC721 contract.
type ERC721MetadataUpdateIterator struct {
	Event *ERC721MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataUpdate)
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
		it.Event = new(ERC721MetadataUpdate)
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
func (it *ERC721MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataUpdate represents a MetadataUpdate event raised by the ERC721 contract.
type ERC721MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721 *ERC721Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ERC721MetadataUpdateIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataUpdateIterator{contract: _ERC721.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721 *ERC721Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC721MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataUpdate)
				if err := _ERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721 *ERC721Filterer) ParseMetadataUpdate(log types.Log) (*ERC721MetadataUpdate, error) {
	event := new(ERC721MetadataUpdate)
	if err := _ERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC721 contract.
type ERC721RoleAdminChangedIterator struct {
	Event *ERC721RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC721RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721RoleAdminChanged)
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
		it.Event = new(ERC721RoleAdminChanged)
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
func (it *ERC721RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721RoleAdminChanged represents a RoleAdminChanged event raised by the ERC721 contract.
type ERC721RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721 *ERC721Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC721RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC721RoleAdminChangedIterator{contract: _ERC721.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721 *ERC721Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC721RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721RoleAdminChanged)
				if err := _ERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721 *ERC721Filterer) ParseRoleAdminChanged(log types.Log) (*ERC721RoleAdminChanged, error) {
	event := new(ERC721RoleAdminChanged)
	if err := _ERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC721 contract.
type ERC721RoleGrantedIterator struct {
	Event *ERC721RoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC721RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721RoleGranted)
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
		it.Event = new(ERC721RoleGranted)
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
func (it *ERC721RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721RoleGranted represents a RoleGranted event raised by the ERC721 contract.
type ERC721RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721RoleGrantedIterator{contract: _ERC721.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC721RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721RoleGranted)
				if err := _ERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) ParseRoleGranted(log types.Log) (*ERC721RoleGranted, error) {
	event := new(ERC721RoleGranted)
	if err := _ERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC721 contract.
type ERC721RoleRevokedIterator struct {
	Event *ERC721RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC721RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721RoleRevoked)
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
		it.Event = new(ERC721RoleRevoked)
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
func (it *ERC721RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721RoleRevoked represents a RoleRevoked event raised by the ERC721 contract.
type ERC721RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721RoleRevokedIterator{contract: _ERC721.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC721RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721RoleRevoked)
				if err := _ERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721 *ERC721Filterer) ParseRoleRevoked(log types.Log) (*ERC721RoleRevoked, error) {
	event := new(ERC721RoleRevoked)
	if err := _ERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
