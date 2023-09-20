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

// BridgeSigTransferSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type BridgeSigTransferSignerInfo struct {
	Deadline *big.Int
	R        [32]byte
	S        [32]byte
	V        uint8
}

// BridgeBaseMetaData contains all meta data concerning the BridgeBase contract.
var BridgeBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"blacklist_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBlacklist\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklist\",\"type\":\"address\"}],\"name\":\"BlacklistChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"BlacklistStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyTokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBridgeSigTransfer.SigInvalidReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"InvalidSigDetected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinTeleportAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumBridgeRefundRequest.RefundStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RefundStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Teleport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_REQUEST_MIN_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"approveRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"areProcessedNonces\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"changeBlacklistState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"otherChainNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structBridgeSigTransfer.SignerInfo[]\",\"name\":\"signerInfo\",\"type\":\"tuple[]\"}],\"name\":\"claimSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"declineRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasSignerRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTeleportAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"approved\",\"type\":\"bool[]\"}],\"name\":\"processRefunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"processedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"refundRequests\",\"outputs\":[{\"internalType\":\"enumBridgeRefundRequest.RefundStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"reopenRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blacklist_\",\"type\":\"address\"}],\"name\":\"setBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinTeleportAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"teleport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"teleportSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"teleports\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBaseMetaData.ABI instead.
var BridgeBaseABI = BridgeBaseMetaData.ABI

// BridgeBase is an auto generated Go binding around an Ethereum contract.
type BridgeBase struct {
	BridgeBaseCaller     // Read-only binding to the contract
	BridgeBaseTransactor // Write-only binding to the contract
	BridgeBaseFilterer   // Log filterer for contract events
}

// BridgeBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBaseSession struct {
	Contract     *BridgeBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBaseCallerSession struct {
	Contract *BridgeBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBaseTransactorSession struct {
	Contract     *BridgeBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBaseRaw struct {
	Contract *BridgeBase // Generic contract binding to access the raw methods on
}

// BridgeBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBaseCallerRaw struct {
	Contract *BridgeBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBaseTransactorRaw struct {
	Contract *BridgeBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBase creates a new instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBase(address common.Address, backend bind.ContractBackend) (*BridgeBase, error) {
	contract, err := bindBridgeBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBase{BridgeBaseCaller: BridgeBaseCaller{contract: contract}, BridgeBaseTransactor: BridgeBaseTransactor{contract: contract}, BridgeBaseFilterer: BridgeBaseFilterer{contract: contract}}, nil
}

// NewBridgeBaseCaller creates a new read-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseCaller(address common.Address, caller bind.ContractCaller) (*BridgeBaseCaller, error) {
	contract, err := bindBridgeBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseCaller{contract: contract}, nil
}

// NewBridgeBaseTransactor creates a new write-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBaseTransactor, error) {
	contract, err := bindBridgeBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseTransactor{contract: contract}, nil
}

// NewBridgeBaseFilterer creates a new log filterer instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBaseFilterer, error) {
	contract, err := bindBridgeBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseFilterer{contract: contract}, nil
}

// bindBridgeBase binds a generic wrapper to an already deployed contract.
func bindBridgeBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.BridgeBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) ADMINROLE() ([32]byte, error) {
	return _BridgeBase.Contract.ADMINROLE(&_BridgeBase.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) ADMINROLE() ([32]byte, error) {
	return _BridgeBase.Contract.ADMINROLE(&_BridgeBase.CallOpts)
}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) CLAIMTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "CLAIM_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) CLAIMTYPEHASH() ([32]byte, error) {
	return _BridgeBase.Contract.CLAIMTYPEHASH(&_BridgeBase.CallOpts)
}

// CLAIMTYPEHASH is a free data retrieval call binding the contract method 0x6b0509b1.
//
// Solidity: function CLAIM_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) CLAIMTYPEHASH() ([32]byte, error) {
	return _BridgeBase.Contract.CLAIMTYPEHASH(&_BridgeBase.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeBase.Contract.DEFAULTADMINROLE(&_BridgeBase.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeBase.Contract.DEFAULTADMINROLE(&_BridgeBase.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) OWNERROLE() ([32]byte, error) {
	return _BridgeBase.Contract.OWNERROLE(&_BridgeBase.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) OWNERROLE() ([32]byte, error) {
	return _BridgeBase.Contract.OWNERROLE(&_BridgeBase.CallOpts)
}

// REFUNDREQUESTMINDELAY is a free data retrieval call binding the contract method 0x2646b7b5.
//
// Solidity: function REFUND_REQUEST_MIN_DELAY() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) REFUNDREQUESTMINDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "REFUND_REQUEST_MIN_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFUNDREQUESTMINDELAY is a free data retrieval call binding the contract method 0x2646b7b5.
//
// Solidity: function REFUND_REQUEST_MIN_DELAY() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) REFUNDREQUESTMINDELAY() (*big.Int, error) {
	return _BridgeBase.Contract.REFUNDREQUESTMINDELAY(&_BridgeBase.CallOpts)
}

// REFUNDREQUESTMINDELAY is a free data retrieval call binding the contract method 0x2646b7b5.
//
// Solidity: function REFUND_REQUEST_MIN_DELAY() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) REFUNDREQUESTMINDELAY() (*big.Int, error) {
	return _BridgeBase.Contract.REFUNDREQUESTMINDELAY(&_BridgeBase.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) SIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) SIGNERROLE() ([32]byte, error) {
	return _BridgeBase.Contract.SIGNERROLE(&_BridgeBase.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) SIGNERROLE() ([32]byte, error) {
	return _BridgeBase.Contract.SIGNERROLE(&_BridgeBase.CallOpts)
}

// TELEPORTTYPEHASH is a free data retrieval call binding the contract method 0xcf41a5fb.
//
// Solidity: function TELEPORT_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) TELEPORTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "TELEPORT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTTYPEHASH is a free data retrieval call binding the contract method 0xcf41a5fb.
//
// Solidity: function TELEPORT_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) TELEPORTTYPEHASH() ([32]byte, error) {
	return _BridgeBase.Contract.TELEPORTTYPEHASH(&_BridgeBase.CallOpts)
}

// TELEPORTTYPEHASH is a free data retrieval call binding the contract method 0xcf41a5fb.
//
// Solidity: function TELEPORT_TYPEHASH() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) TELEPORTTYPEHASH() ([32]byte, error) {
	return _BridgeBase.Contract.TELEPORTTYPEHASH(&_BridgeBase.CallOpts)
}

// AreProcessedNonces is a free data retrieval call binding the contract method 0x39653984.
//
// Solidity: function areProcessedNonces(uint256[] nonces) view returns(bool[])
func (_BridgeBase *BridgeBaseCaller) AreProcessedNonces(opts *bind.CallOpts, nonces []*big.Int) ([]bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "areProcessedNonces", nonces)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// AreProcessedNonces is a free data retrieval call binding the contract method 0x39653984.
//
// Solidity: function areProcessedNonces(uint256[] nonces) view returns(bool[])
func (_BridgeBase *BridgeBaseSession) AreProcessedNonces(nonces []*big.Int) ([]bool, error) {
	return _BridgeBase.Contract.AreProcessedNonces(&_BridgeBase.CallOpts, nonces)
}

// AreProcessedNonces is a free data retrieval call binding the contract method 0x39653984.
//
// Solidity: function areProcessedNonces(uint256[] nonces) view returns(bool[])
func (_BridgeBase *BridgeBaseCallerSession) AreProcessedNonces(nonces []*big.Int) ([]bool, error) {
	return _BridgeBase.Contract.AreProcessedNonces(&_BridgeBase.CallOpts, nonces)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_BridgeBase *BridgeBaseCaller) Blacklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "blacklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_BridgeBase *BridgeBaseSession) Blacklist() (common.Address, error) {
	return _BridgeBase.Contract.Blacklist(&_BridgeBase.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0xa4b5fa56.
//
// Solidity: function blacklist() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) Blacklist() (common.Address, error) {
	return _BridgeBase.Contract.Blacklist(&_BridgeBase.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeBase.Contract.GetRoleAdmin(&_BridgeBase.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeBase.Contract.GetRoleAdmin(&_BridgeBase.CallOpts, role)
}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCaller) GrantRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "grantRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseSession) GrantRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.GrantRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// GrantRole is a free data retrieval call binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCallerSession) GrantRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.GrantRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBase *BridgeBaseSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeBase.Contract.HasRole(&_BridgeBase.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeBase.Contract.HasRole(&_BridgeBase.CallOpts, role, account)
}

// HasSignerRole is a free data retrieval call binding the contract method 0xb7792945.
//
// Solidity: function hasSignerRole(address account) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) HasSignerRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "hasSignerRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSignerRole is a free data retrieval call binding the contract method 0xb7792945.
//
// Solidity: function hasSignerRole(address account) view returns(bool)
func (_BridgeBase *BridgeBaseSession) HasSignerRole(account common.Address) (bool, error) {
	return _BridgeBase.Contract.HasSignerRole(&_BridgeBase.CallOpts, account)
}

// HasSignerRole is a free data retrieval call binding the contract method 0xb7792945.
//
// Solidity: function hasSignerRole(address account) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) HasSignerRole(account common.Address) (bool, error) {
	return _BridgeBase.Contract.HasSignerRole(&_BridgeBase.CallOpts, account)
}

// MinTeleportAmount is a free data retrieval call binding the contract method 0x6ccb63a6.
//
// Solidity: function minTeleportAmount() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) MinTeleportAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "minTeleportAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTeleportAmount is a free data retrieval call binding the contract method 0x6ccb63a6.
//
// Solidity: function minTeleportAmount() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) MinTeleportAmount() (*big.Int, error) {
	return _BridgeBase.Contract.MinTeleportAmount(&_BridgeBase.CallOpts)
}

// MinTeleportAmount is a free data retrieval call binding the contract method 0x6ccb63a6.
//
// Solidity: function minTeleportAmount() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) MinTeleportAmount() (*big.Int, error) {
	return _BridgeBase.Contract.MinTeleportAmount(&_BridgeBase.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) Nonce() (*big.Int, error) {
	return _BridgeBase.Contract.Nonce(&_BridgeBase.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) Nonce() (*big.Int, error) {
	return _BridgeBase.Contract.Nonce(&_BridgeBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseSession) Paused() (bool, error) {
	return _BridgeBase.Contract.Paused(&_BridgeBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) Paused() (bool, error) {
	return _BridgeBase.Contract.Paused(&_BridgeBase.CallOpts)
}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) ProcessedNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "processedNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgeBase *BridgeBaseSession) ProcessedNonces(arg0 *big.Int) (bool, error) {
	return _BridgeBase.Contract.ProcessedNonces(&_BridgeBase.CallOpts, arg0)
}

// ProcessedNonces is a free data retrieval call binding the contract method 0x91dae519.
//
// Solidity: function processedNonces(uint256 ) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) ProcessedNonces(arg0 *big.Int) (bool, error) {
	return _BridgeBase.Contract.ProcessedNonces(&_BridgeBase.CallOpts, arg0)
}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint8 status, uint64 timestamp, address from, uint256 nonce)
func (_BridgeBase *BridgeBaseCaller) RefundRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status    uint8
	Timestamp uint64
	From      common.Address
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "refundRequests", arg0)

	outstruct := new(struct {
		Status    uint8
		Timestamp uint64
		From      common.Address
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.From = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint8 status, uint64 timestamp, address from, uint256 nonce)
func (_BridgeBase *BridgeBaseSession) RefundRequests(arg0 *big.Int) (struct {
	Status    uint8
	Timestamp uint64
	From      common.Address
	Nonce     *big.Int
}, error) {
	return _BridgeBase.Contract.RefundRequests(&_BridgeBase.CallOpts, arg0)
}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint8 status, uint64 timestamp, address from, uint256 nonce)
func (_BridgeBase *BridgeBaseCallerSession) RefundRequests(arg0 *big.Int) (struct {
	Status    uint8
	Timestamp uint64
	From      common.Address
	Nonce     *big.Int
}, error) {
	return _BridgeBase.Contract.RefundRequests(&_BridgeBase.CallOpts, arg0)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCaller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.RenounceRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.RenounceRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "requiredSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) RequiredSignatures() (*big.Int, error) {
	return _BridgeBase.Contract.RequiredSignatures(&_BridgeBase.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) RequiredSignatures() (*big.Int, error) {
	return _BridgeBase.Contract.RequiredSignatures(&_BridgeBase.CallOpts)
}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCaller) RevokeRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "revokeRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseSession) RevokeRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.RevokeRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// RevokeRole is a free data retrieval call binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 , address ) pure returns()
func (_BridgeBase *BridgeBaseCallerSession) RevokeRole(arg0 [32]byte, arg1 common.Address) error {
	return _BridgeBase.Contract.RevokeRole(&_BridgeBase.CallOpts, arg0, arg1)
}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) SignerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "signerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) SignerCount() (*big.Int, error) {
	return _BridgeBase.Contract.SignerCount(&_BridgeBase.CallOpts)
}

// SignerCount is a free data retrieval call binding the contract method 0x7ca548c6.
//
// Solidity: function signerCount() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) SignerCount() (*big.Int, error) {
	return _BridgeBase.Contract.SignerCount(&_BridgeBase.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBase *BridgeBaseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBase *BridgeBaseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeBase.Contract.SupportsInterface(&_BridgeBase.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBase *BridgeBaseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeBase.Contract.SupportsInterface(&_BridgeBase.CallOpts, interfaceId)
}

// Teleports is a free data retrieval call binding the contract method 0x78289788.
//
// Solidity: function teleports(uint256 ) view returns(uint64 timestamp, address from, uint64 nonce, address to, uint256 amount)
func (_BridgeBase *BridgeBaseCaller) Teleports(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp uint64
	From      common.Address
	Nonce     uint64
	To        common.Address
	Amount    *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "teleports", arg0)

	outstruct := new(struct {
		Timestamp uint64
		From      common.Address
		Nonce     uint64
		To        common.Address
		Amount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.From = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.To = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Teleports is a free data retrieval call binding the contract method 0x78289788.
//
// Solidity: function teleports(uint256 ) view returns(uint64 timestamp, address from, uint64 nonce, address to, uint256 amount)
func (_BridgeBase *BridgeBaseSession) Teleports(arg0 *big.Int) (struct {
	Timestamp uint64
	From      common.Address
	Nonce     uint64
	To        common.Address
	Amount    *big.Int
}, error) {
	return _BridgeBase.Contract.Teleports(&_BridgeBase.CallOpts, arg0)
}

// Teleports is a free data retrieval call binding the contract method 0x78289788.
//
// Solidity: function teleports(uint256 ) view returns(uint64 timestamp, address from, uint64 nonce, address to, uint256 amount)
func (_BridgeBase *BridgeBaseCallerSession) Teleports(arg0 *big.Int) (struct {
	Timestamp uint64
	From      common.Address
	Nonce     uint64
	To        common.Address
	Amount    *big.Int
}, error) {
	return _BridgeBase.Contract.Teleports(&_BridgeBase.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBase *BridgeBaseCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBase *BridgeBaseSession) Token() (common.Address, error) {
	return _BridgeBase.Contract.Token(&_BridgeBase.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) Token() (common.Address, error) {
	return _BridgeBase.Contract.Token(&_BridgeBase.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_BridgeBase *BridgeBaseTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_BridgeBase *BridgeBaseSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddAdmin(&_BridgeBase.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_BridgeBase *BridgeBaseTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddAdmin(&_BridgeBase.TransactOpts, account)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address account) returns()
func (_BridgeBase *BridgeBaseTransactor) AddOwner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "addOwner", account)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address account) returns()
func (_BridgeBase *BridgeBaseSession) AddOwner(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddOwner(&_BridgeBase.TransactOpts, account)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address account) returns()
func (_BridgeBase *BridgeBaseTransactorSession) AddOwner(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddOwner(&_BridgeBase.TransactOpts, account)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseTransactor) AddSigners(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "addSigners", accounts)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseSession) AddSigners(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddSigners(&_BridgeBase.TransactOpts, accounts)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseTransactorSession) AddSigners(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.AddSigners(&_BridgeBase.TransactOpts, accounts)
}

// ApproveRefund is a paid mutator transaction binding the contract method 0x348a71a6.
//
// Solidity: function approveRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactor) ApproveRefund(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "approveRefund", nonce)
}

// ApproveRefund is a paid mutator transaction binding the contract method 0x348a71a6.
//
// Solidity: function approveRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseSession) ApproveRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.ApproveRefund(&_BridgeBase.TransactOpts, nonce)
}

// ApproveRefund is a paid mutator transaction binding the contract method 0x348a71a6.
//
// Solidity: function approveRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ApproveRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.ApproveRefund(&_BridgeBase.TransactOpts, nonce)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_BridgeBase *BridgeBaseTransactor) ChangeBlacklistState(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "changeBlacklistState", state)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_BridgeBase *BridgeBaseSession) ChangeBlacklistState(state bool) (*types.Transaction, error) {
	return _BridgeBase.Contract.ChangeBlacklistState(&_BridgeBase.TransactOpts, state)
}

// ChangeBlacklistState is a paid mutator transaction binding the contract method 0x9f3eb02a.
//
// Solidity: function changeBlacklistState(bool state) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ChangeBlacklistState(state bool) (*types.Transaction, error) {
	return _BridgeBase.Contract.ChangeBlacklistState(&_BridgeBase.TransactOpts, state)
}

// ClaimSig is a paid mutator transaction binding the contract method 0x8f4ce922.
//
// Solidity: function claimSig(address to, uint256 amount, uint256 otherChainNonce, (uint256,bytes32,bytes32,uint8)[] signerInfo) returns()
func (_BridgeBase *BridgeBaseTransactor) ClaimSig(opts *bind.TransactOpts, to common.Address, amount *big.Int, otherChainNonce *big.Int, signerInfo []BridgeSigTransferSignerInfo) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "claimSig", to, amount, otherChainNonce, signerInfo)
}

// ClaimSig is a paid mutator transaction binding the contract method 0x8f4ce922.
//
// Solidity: function claimSig(address to, uint256 amount, uint256 otherChainNonce, (uint256,bytes32,bytes32,uint8)[] signerInfo) returns()
func (_BridgeBase *BridgeBaseSession) ClaimSig(to common.Address, amount *big.Int, otherChainNonce *big.Int, signerInfo []BridgeSigTransferSignerInfo) (*types.Transaction, error) {
	return _BridgeBase.Contract.ClaimSig(&_BridgeBase.TransactOpts, to, amount, otherChainNonce, signerInfo)
}

// ClaimSig is a paid mutator transaction binding the contract method 0x8f4ce922.
//
// Solidity: function claimSig(address to, uint256 amount, uint256 otherChainNonce, (uint256,bytes32,bytes32,uint8)[] signerInfo) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ClaimSig(to common.Address, amount *big.Int, otherChainNonce *big.Int, signerInfo []BridgeSigTransferSignerInfo) (*types.Transaction, error) {
	return _BridgeBase.Contract.ClaimSig(&_BridgeBase.TransactOpts, to, amount, otherChainNonce, signerInfo)
}

// DeclineRefund is a paid mutator transaction binding the contract method 0x4b117e32.
//
// Solidity: function declineRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactor) DeclineRefund(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "declineRefund", nonce)
}

// DeclineRefund is a paid mutator transaction binding the contract method 0x4b117e32.
//
// Solidity: function declineRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseSession) DeclineRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.DeclineRefund(&_BridgeBase.TransactOpts, nonce)
}

// DeclineRefund is a paid mutator transaction binding the contract method 0x4b117e32.
//
// Solidity: function declineRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactorSession) DeclineRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.DeclineRefund(&_BridgeBase.TransactOpts, nonce)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseSession) Pause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Pause(&_BridgeBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBase *BridgeBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Pause(&_BridgeBase.TransactOpts)
}

// ProcessRefunds is a paid mutator transaction binding the contract method 0x226e11ce.
//
// Solidity: function processRefunds(uint256[] nonces, bool[] approved) returns()
func (_BridgeBase *BridgeBaseTransactor) ProcessRefunds(opts *bind.TransactOpts, nonces []*big.Int, approved []bool) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "processRefunds", nonces, approved)
}

// ProcessRefunds is a paid mutator transaction binding the contract method 0x226e11ce.
//
// Solidity: function processRefunds(uint256[] nonces, bool[] approved) returns()
func (_BridgeBase *BridgeBaseSession) ProcessRefunds(nonces []*big.Int, approved []bool) (*types.Transaction, error) {
	return _BridgeBase.Contract.ProcessRefunds(&_BridgeBase.TransactOpts, nonces, approved)
}

// ProcessRefunds is a paid mutator transaction binding the contract method 0x226e11ce.
//
// Solidity: function processRefunds(uint256[] nonces, bool[] approved) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ProcessRefunds(nonces []*big.Int, approved []bool) (*types.Transaction, error) {
	return _BridgeBase.Contract.ProcessRefunds(&_BridgeBase.TransactOpts, nonces, approved)
}

// ReopenRefund is a paid mutator transaction binding the contract method 0x25774686.
//
// Solidity: function reopenRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactor) ReopenRefund(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "reopenRefund", nonce)
}

// ReopenRefund is a paid mutator transaction binding the contract method 0x25774686.
//
// Solidity: function reopenRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseSession) ReopenRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.ReopenRefund(&_BridgeBase.TransactOpts, nonce)
}

// ReopenRefund is a paid mutator transaction binding the contract method 0x25774686.
//
// Solidity: function reopenRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ReopenRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.ReopenRefund(&_BridgeBase.TransactOpts, nonce)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa4b2409e.
//
// Solidity: function requestRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactor) RequestRefund(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "requestRefund", nonce)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa4b2409e.
//
// Solidity: function requestRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseSession) RequestRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.RequestRefund(&_BridgeBase.TransactOpts, nonce)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa4b2409e.
//
// Solidity: function requestRefund(uint256 nonce) returns()
func (_BridgeBase *BridgeBaseTransactorSession) RequestRefund(nonce *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.RequestRefund(&_BridgeBase.TransactOpts, nonce)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address account) returns()
func (_BridgeBase *BridgeBaseTransactor) RevokeAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "revokeAdmin", account)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address account) returns()
func (_BridgeBase *BridgeBaseSession) RevokeAdmin(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeAdmin(&_BridgeBase.TransactOpts, account)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address account) returns()
func (_BridgeBase *BridgeBaseTransactorSession) RevokeAdmin(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeAdmin(&_BridgeBase.TransactOpts, account)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address account) returns()
func (_BridgeBase *BridgeBaseTransactor) RevokeOwner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "revokeOwner", account)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address account) returns()
func (_BridgeBase *BridgeBaseSession) RevokeOwner(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeOwner(&_BridgeBase.TransactOpts, account)
}

// RevokeOwner is a paid mutator transaction binding the contract method 0x6d151101.
//
// Solidity: function revokeOwner(address account) returns()
func (_BridgeBase *BridgeBaseTransactorSession) RevokeOwner(account common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeOwner(&_BridgeBase.TransactOpts, account)
}

// RevokeSigners is a paid mutator transaction binding the contract method 0xe9981490.
//
// Solidity: function revokeSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseTransactor) RevokeSigners(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "revokeSigners", accounts)
}

// RevokeSigners is a paid mutator transaction binding the contract method 0xe9981490.
//
// Solidity: function revokeSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseSession) RevokeSigners(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeSigners(&_BridgeBase.TransactOpts, accounts)
}

// RevokeSigners is a paid mutator transaction binding the contract method 0xe9981490.
//
// Solidity: function revokeSigners(address[] accounts) returns()
func (_BridgeBase *BridgeBaseTransactorSession) RevokeSigners(accounts []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RevokeSigners(&_BridgeBase.TransactOpts, accounts)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_BridgeBase *BridgeBaseTransactor) SetBlacklist(opts *bind.TransactOpts, blacklist_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "setBlacklist", blacklist_)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_BridgeBase *BridgeBaseSession) SetBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetBlacklist(&_BridgeBase.TransactOpts, blacklist_)
}

// SetBlacklist is a paid mutator transaction binding the contract method 0x4e054a67.
//
// Solidity: function setBlacklist(address blacklist_) returns()
func (_BridgeBase *BridgeBaseTransactorSession) SetBlacklist(blacklist_ common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetBlacklist(&_BridgeBase.TransactOpts, blacklist_)
}

// SetMinTeleportAmount is a paid mutator transaction binding the contract method 0xaf7dcc88.
//
// Solidity: function setMinTeleportAmount(uint256 amount) returns()
func (_BridgeBase *BridgeBaseTransactor) SetMinTeleportAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "setMinTeleportAmount", amount)
}

// SetMinTeleportAmount is a paid mutator transaction binding the contract method 0xaf7dcc88.
//
// Solidity: function setMinTeleportAmount(uint256 amount) returns()
func (_BridgeBase *BridgeBaseSession) SetMinTeleportAmount(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetMinTeleportAmount(&_BridgeBase.TransactOpts, amount)
}

// SetMinTeleportAmount is a paid mutator transaction binding the contract method 0xaf7dcc88.
//
// Solidity: function setMinTeleportAmount(uint256 amount) returns()
func (_BridgeBase *BridgeBaseTransactorSession) SetMinTeleportAmount(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetMinTeleportAmount(&_BridgeBase.TransactOpts, amount)
}

// Teleport is a paid mutator transaction binding the contract method 0x89832beb.
//
// Solidity: function teleport(uint256 amount) returns()
func (_BridgeBase *BridgeBaseTransactor) Teleport(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "teleport", amount)
}

// Teleport is a paid mutator transaction binding the contract method 0x89832beb.
//
// Solidity: function teleport(uint256 amount) returns()
func (_BridgeBase *BridgeBaseSession) Teleport(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.Teleport(&_BridgeBase.TransactOpts, amount)
}

// Teleport is a paid mutator transaction binding the contract method 0x89832beb.
//
// Solidity: function teleport(uint256 amount) returns()
func (_BridgeBase *BridgeBaseTransactorSession) Teleport(amount *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.Teleport(&_BridgeBase.TransactOpts, amount)
}

// TeleportSig is a paid mutator transaction binding the contract method 0x5e603adb.
//
// Solidity: function teleportSig(address from, uint256 amount, uint256 deadline, bytes32 r, bytes32 s, uint8 v) returns()
func (_BridgeBase *BridgeBaseTransactor) TeleportSig(opts *bind.TransactOpts, from common.Address, amount *big.Int, deadline *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "teleportSig", from, amount, deadline, r, s, v)
}

// TeleportSig is a paid mutator transaction binding the contract method 0x5e603adb.
//
// Solidity: function teleportSig(address from, uint256 amount, uint256 deadline, bytes32 r, bytes32 s, uint8 v) returns()
func (_BridgeBase *BridgeBaseSession) TeleportSig(from common.Address, amount *big.Int, deadline *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _BridgeBase.Contract.TeleportSig(&_BridgeBase.TransactOpts, from, amount, deadline, r, s, v)
}

// TeleportSig is a paid mutator transaction binding the contract method 0x5e603adb.
//
// Solidity: function teleportSig(address from, uint256 amount, uint256 deadline, bytes32 r, bytes32 s, uint8 v) returns()
func (_BridgeBase *BridgeBaseTransactorSession) TeleportSig(from common.Address, amount *big.Int, deadline *big.Int, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _BridgeBase.Contract.TeleportSig(&_BridgeBase.TransactOpts, from, amount, deadline, r, s, v)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseSession) Unpause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Unpause(&_BridgeBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBase *BridgeBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeBase.Contract.Unpause(&_BridgeBase.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_BridgeBase *BridgeBaseTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_BridgeBase *BridgeBaseSession) WithdrawEther() (*types.Transaction, error) {
	return _BridgeBase.Contract.WithdrawEther(&_BridgeBase.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_BridgeBase *BridgeBaseTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _BridgeBase.Contract.WithdrawEther(&_BridgeBase.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5ecb16cd.
//
// Solidity: function withdrawTokens(address[] tokens) returns()
func (_BridgeBase *BridgeBaseTransactor) WithdrawTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "withdrawTokens", tokens)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5ecb16cd.
//
// Solidity: function withdrawTokens(address[] tokens) returns()
func (_BridgeBase *BridgeBaseSession) WithdrawTokens(tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.WithdrawTokens(&_BridgeBase.TransactOpts, tokens)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5ecb16cd.
//
// Solidity: function withdrawTokens(address[] tokens) returns()
func (_BridgeBase *BridgeBaseTransactorSession) WithdrawTokens(tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.WithdrawTokens(&_BridgeBase.TransactOpts, tokens)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseSession) Receive() (*types.Transaction, error) {
	return _BridgeBase.Contract.Receive(&_BridgeBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBase *BridgeBaseTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBase.Contract.Receive(&_BridgeBase.TransactOpts)
}

// BridgeBaseBlacklistChangedIterator is returned from FilterBlacklistChanged and is used to iterate over the raw logs and unpacked data for BlacklistChanged events raised by the BridgeBase contract.
type BridgeBaseBlacklistChangedIterator struct {
	Event *BridgeBaseBlacklistChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBaseBlacklistChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseBlacklistChanged)
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
		it.Event = new(BridgeBaseBlacklistChanged)
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
func (it *BridgeBaseBlacklistChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseBlacklistChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseBlacklistChanged represents a BlacklistChanged event raised by the BridgeBase contract.
type BridgeBaseBlacklistChanged struct {
	OldBlacklist common.Address
	NewBlacklist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlacklistChanged is a free log retrieval operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_BridgeBase *BridgeBaseFilterer) FilterBlacklistChanged(opts *bind.FilterOpts, oldBlacklist []common.Address, newBlacklist []common.Address) (*BridgeBaseBlacklistChangedIterator, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseBlacklistChangedIterator{contract: _BridgeBase.contract, event: "BlacklistChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistChanged is a free log subscription operation binding the contract event 0x6bd4b770608a77cbad8c46801ad2580495c7ad5df7efebb95a940b27a7abb79f.
//
// Solidity: event BlacklistChanged(address indexed oldBlacklist, address indexed newBlacklist)
func (_BridgeBase *BridgeBaseFilterer) WatchBlacklistChanged(opts *bind.WatchOpts, sink chan<- *BridgeBaseBlacklistChanged, oldBlacklist []common.Address, newBlacklist []common.Address) (event.Subscription, error) {

	var oldBlacklistRule []interface{}
	for _, oldBlacklistItem := range oldBlacklist {
		oldBlacklistRule = append(oldBlacklistRule, oldBlacklistItem)
	}
	var newBlacklistRule []interface{}
	for _, newBlacklistItem := range newBlacklist {
		newBlacklistRule = append(newBlacklistRule, newBlacklistItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "BlacklistChanged", oldBlacklistRule, newBlacklistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseBlacklistChanged)
				if err := _BridgeBase.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseBlacklistChanged(log types.Log) (*BridgeBaseBlacklistChanged, error) {
	event := new(BridgeBaseBlacklistChanged)
	if err := _BridgeBase.contract.UnpackLog(event, "BlacklistChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseBlacklistStatusChangedIterator is returned from FilterBlacklistStatusChanged and is used to iterate over the raw logs and unpacked data for BlacklistStatusChanged events raised by the BridgeBase contract.
type BridgeBaseBlacklistStatusChangedIterator struct {
	Event *BridgeBaseBlacklistStatusChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBaseBlacklistStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseBlacklistStatusChanged)
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
		it.Event = new(BridgeBaseBlacklistStatusChanged)
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
func (it *BridgeBaseBlacklistStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseBlacklistStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseBlacklistStatusChanged represents a BlacklistStatusChanged event raised by the BridgeBase contract.
type BridgeBaseBlacklistStatusChanged struct {
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlacklistStatusChanged is a free log retrieval operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_BridgeBase *BridgeBaseFilterer) FilterBlacklistStatusChanged(opts *bind.FilterOpts, newStatus []bool) (*BridgeBaseBlacklistStatusChangedIterator, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseBlacklistStatusChangedIterator{contract: _BridgeBase.contract, event: "BlacklistStatusChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistStatusChanged is a free log subscription operation binding the contract event 0x288d46a73defdc9e7cba47d73bc52c0abbc2edbc62b67fb4a5d60bac7959aade.
//
// Solidity: event BlacklistStatusChanged(bool indexed newStatus)
func (_BridgeBase *BridgeBaseFilterer) WatchBlacklistStatusChanged(opts *bind.WatchOpts, sink chan<- *BridgeBaseBlacklistStatusChanged, newStatus []bool) (event.Subscription, error) {

	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "BlacklistStatusChanged", newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseBlacklistStatusChanged)
				if err := _BridgeBase.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseBlacklistStatusChanged(log types.Log) (*BridgeBaseBlacklistStatusChanged, error) {
	event := new(BridgeBaseBlacklistStatusChanged)
	if err := _BridgeBase.contract.UnpackLog(event, "BlacklistStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BridgeBase contract.
type BridgeBaseClaimedIterator struct {
	Event *BridgeBaseClaimed // Event containing the contract specifics and raw log

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
func (it *BridgeBaseClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseClaimed)
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
		it.Event = new(BridgeBaseClaimed)
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
func (it *BridgeBaseClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseClaimed represents a Claimed event raised by the BridgeBase contract.
type BridgeBaseClaimed struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) FilterClaimed(opts *bind.FilterOpts, from []common.Address, to []common.Address, nonce []*big.Int) (*BridgeBaseClaimedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Claimed", fromRule, toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseClaimedIterator{contract: _BridgeBase.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BridgeBaseClaimed, from []common.Address, to []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Claimed", fromRule, toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseClaimed)
				if err := _BridgeBase.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) ParseClaimed(log types.Log) (*BridgeBaseClaimed, error) {
	event := new(BridgeBaseClaimed)
	if err := _BridgeBase.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseEmergencyTokenWithdrawIterator is returned from FilterEmergencyTokenWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyTokenWithdraw events raised by the BridgeBase contract.
type BridgeBaseEmergencyTokenWithdrawIterator struct {
	Event *BridgeBaseEmergencyTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeBaseEmergencyTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseEmergencyTokenWithdraw)
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
		it.Event = new(BridgeBaseEmergencyTokenWithdraw)
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
func (it *BridgeBaseEmergencyTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseEmergencyTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseEmergencyTokenWithdraw represents a EmergencyTokenWithdraw event raised by the BridgeBase contract.
type BridgeBaseEmergencyTokenWithdraw struct {
	Token     common.Address
	To        common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmergencyTokenWithdraw is a free log retrieval operation binding the contract event 0x299ebbc94ff2c6eedc1376c962dca03a87feea4cbc7cf1af5966c5f47e158956.
//
// Solidity: event EmergencyTokenWithdraw(address indexed token, address indexed to, uint256 amount, uint256 timestamp)
func (_BridgeBase *BridgeBaseFilterer) FilterEmergencyTokenWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeBaseEmergencyTokenWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "EmergencyTokenWithdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseEmergencyTokenWithdrawIterator{contract: _BridgeBase.contract, event: "EmergencyTokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyTokenWithdraw is a free log subscription operation binding the contract event 0x299ebbc94ff2c6eedc1376c962dca03a87feea4cbc7cf1af5966c5f47e158956.
//
// Solidity: event EmergencyTokenWithdraw(address indexed token, address indexed to, uint256 amount, uint256 timestamp)
func (_BridgeBase *BridgeBaseFilterer) WatchEmergencyTokenWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeBaseEmergencyTokenWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "EmergencyTokenWithdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseEmergencyTokenWithdraw)
				if err := _BridgeBase.contract.UnpackLog(event, "EmergencyTokenWithdraw", log); err != nil {
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

// ParseEmergencyTokenWithdraw is a log parse operation binding the contract event 0x299ebbc94ff2c6eedc1376c962dca03a87feea4cbc7cf1af5966c5f47e158956.
//
// Solidity: event EmergencyTokenWithdraw(address indexed token, address indexed to, uint256 amount, uint256 timestamp)
func (_BridgeBase *BridgeBaseFilterer) ParseEmergencyTokenWithdraw(log types.Log) (*BridgeBaseEmergencyTokenWithdraw, error) {
	event := new(BridgeBaseEmergencyTokenWithdraw)
	if err := _BridgeBase.contract.UnpackLog(event, "EmergencyTokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseInvalidSigDetectedIterator is returned from FilterInvalidSigDetected and is used to iterate over the raw logs and unpacked data for InvalidSigDetected events raised by the BridgeBase contract.
type BridgeBaseInvalidSigDetectedIterator struct {
	Event *BridgeBaseInvalidSigDetected // Event containing the contract specifics and raw log

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
func (it *BridgeBaseInvalidSigDetectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseInvalidSigDetected)
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
		it.Event = new(BridgeBaseInvalidSigDetected)
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
func (it *BridgeBaseInvalidSigDetectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseInvalidSigDetectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseInvalidSigDetected represents a InvalidSigDetected event raised by the BridgeBase contract.
type BridgeBaseInvalidSigDetected struct {
	Nonce  *big.Int
	Id     *big.Int
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInvalidSigDetected is a free log retrieval operation binding the contract event 0x46883fb25d9127d8d911bc110e75d5016b90645693aacfd8158263c63955328b.
//
// Solidity: event InvalidSigDetected(uint256 indexed nonce, uint256 id, uint8 reason)
func (_BridgeBase *BridgeBaseFilterer) FilterInvalidSigDetected(opts *bind.FilterOpts, nonce []*big.Int) (*BridgeBaseInvalidSigDetectedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "InvalidSigDetected", nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseInvalidSigDetectedIterator{contract: _BridgeBase.contract, event: "InvalidSigDetected", logs: logs, sub: sub}, nil
}

// WatchInvalidSigDetected is a free log subscription operation binding the contract event 0x46883fb25d9127d8d911bc110e75d5016b90645693aacfd8158263c63955328b.
//
// Solidity: event InvalidSigDetected(uint256 indexed nonce, uint256 id, uint8 reason)
func (_BridgeBase *BridgeBaseFilterer) WatchInvalidSigDetected(opts *bind.WatchOpts, sink chan<- *BridgeBaseInvalidSigDetected, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "InvalidSigDetected", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseInvalidSigDetected)
				if err := _BridgeBase.contract.UnpackLog(event, "InvalidSigDetected", log); err != nil {
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

// ParseInvalidSigDetected is a log parse operation binding the contract event 0x46883fb25d9127d8d911bc110e75d5016b90645693aacfd8158263c63955328b.
//
// Solidity: event InvalidSigDetected(uint256 indexed nonce, uint256 id, uint8 reason)
func (_BridgeBase *BridgeBaseFilterer) ParseInvalidSigDetected(log types.Log) (*BridgeBaseInvalidSigDetected, error) {
	event := new(BridgeBaseInvalidSigDetected)
	if err := _BridgeBase.contract.UnpackLog(event, "InvalidSigDetected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseMinTeleportAmountChangedIterator is returned from FilterMinTeleportAmountChanged and is used to iterate over the raw logs and unpacked data for MinTeleportAmountChanged events raised by the BridgeBase contract.
type BridgeBaseMinTeleportAmountChangedIterator struct {
	Event *BridgeBaseMinTeleportAmountChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBaseMinTeleportAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseMinTeleportAmountChanged)
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
		it.Event = new(BridgeBaseMinTeleportAmountChanged)
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
func (it *BridgeBaseMinTeleportAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseMinTeleportAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseMinTeleportAmountChanged represents a MinTeleportAmountChanged event raised by the BridgeBase contract.
type BridgeBaseMinTeleportAmountChanged struct {
	Changer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleportAmountChanged is a free log retrieval operation binding the contract event 0xc97cb7f52f8eda365e1c51d75db5fd8d0a2e5a99a11f62af4ee85648d1bb0c5c.
//
// Solidity: event MinTeleportAmountChanged(address changer, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) FilterMinTeleportAmountChanged(opts *bind.FilterOpts) (*BridgeBaseMinTeleportAmountChangedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "MinTeleportAmountChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeBaseMinTeleportAmountChangedIterator{contract: _BridgeBase.contract, event: "MinTeleportAmountChanged", logs: logs, sub: sub}, nil
}

// WatchMinTeleportAmountChanged is a free log subscription operation binding the contract event 0xc97cb7f52f8eda365e1c51d75db5fd8d0a2e5a99a11f62af4ee85648d1bb0c5c.
//
// Solidity: event MinTeleportAmountChanged(address changer, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) WatchMinTeleportAmountChanged(opts *bind.WatchOpts, sink chan<- *BridgeBaseMinTeleportAmountChanged) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "MinTeleportAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseMinTeleportAmountChanged)
				if err := _BridgeBase.contract.UnpackLog(event, "MinTeleportAmountChanged", log); err != nil {
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

// ParseMinTeleportAmountChanged is a log parse operation binding the contract event 0xc97cb7f52f8eda365e1c51d75db5fd8d0a2e5a99a11f62af4ee85648d1bb0c5c.
//
// Solidity: event MinTeleportAmountChanged(address changer, uint256 amount)
func (_BridgeBase *BridgeBaseFilterer) ParseMinTeleportAmountChanged(log types.Log) (*BridgeBaseMinTeleportAmountChanged, error) {
	event := new(BridgeBaseMinTeleportAmountChanged)
	if err := _BridgeBase.contract.UnpackLog(event, "MinTeleportAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeBase contract.
type BridgeBasePausedIterator struct {
	Event *BridgeBasePaused // Event containing the contract specifics and raw log

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
func (it *BridgeBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBasePaused)
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
		it.Event = new(BridgeBasePaused)
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
func (it *BridgeBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBasePaused represents a Paused event raised by the BridgeBase contract.
type BridgeBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeBasePausedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeBasePausedIterator{contract: _BridgeBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeBasePaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBasePaused)
				if err := _BridgeBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBase *BridgeBaseFilterer) ParsePaused(log types.Log) (*BridgeBasePaused, error) {
	event := new(BridgeBasePaused)
	if err := _BridgeBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseRefundStatusChangedIterator is returned from FilterRefundStatusChanged and is used to iterate over the raw logs and unpacked data for RefundStatusChanged events raised by the BridgeBase contract.
type BridgeBaseRefundStatusChangedIterator struct {
	Event *BridgeBaseRefundStatusChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBaseRefundStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseRefundStatusChanged)
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
		it.Event = new(BridgeBaseRefundStatusChanged)
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
func (it *BridgeBaseRefundStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseRefundStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseRefundStatusChanged represents a RefundStatusChanged event raised by the BridgeBase contract.
type BridgeBaseRefundStatusChanged struct {
	From      common.Address
	Nonce     *big.Int
	Status    uint8
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundStatusChanged is a free log retrieval operation binding the contract event 0x80223ad6c02b5ccbc94a594100922a5de2c86d3a09718223d3a64a7a5effa2d7.
//
// Solidity: event RefundStatusChanged(address indexed from, uint256 indexed nonce, uint8 indexed status, uint64 timestamp)
func (_BridgeBase *BridgeBaseFilterer) FilterRefundStatusChanged(opts *bind.FilterOpts, from []common.Address, nonce []*big.Int, status []uint8) (*BridgeBaseRefundStatusChangedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "RefundStatusChanged", fromRule, nonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseRefundStatusChangedIterator{contract: _BridgeBase.contract, event: "RefundStatusChanged", logs: logs, sub: sub}, nil
}

// WatchRefundStatusChanged is a free log subscription operation binding the contract event 0x80223ad6c02b5ccbc94a594100922a5de2c86d3a09718223d3a64a7a5effa2d7.
//
// Solidity: event RefundStatusChanged(address indexed from, uint256 indexed nonce, uint8 indexed status, uint64 timestamp)
func (_BridgeBase *BridgeBaseFilterer) WatchRefundStatusChanged(opts *bind.WatchOpts, sink chan<- *BridgeBaseRefundStatusChanged, from []common.Address, nonce []*big.Int, status []uint8) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "RefundStatusChanged", fromRule, nonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseRefundStatusChanged)
				if err := _BridgeBase.contract.UnpackLog(event, "RefundStatusChanged", log); err != nil {
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

// ParseRefundStatusChanged is a log parse operation binding the contract event 0x80223ad6c02b5ccbc94a594100922a5de2c86d3a09718223d3a64a7a5effa2d7.
//
// Solidity: event RefundStatusChanged(address indexed from, uint256 indexed nonce, uint8 indexed status, uint64 timestamp)
func (_BridgeBase *BridgeBaseFilterer) ParseRefundStatusChanged(log types.Log) (*BridgeBaseRefundStatusChanged, error) {
	event := new(BridgeBaseRefundStatusChanged)
	if err := _BridgeBase.contract.UnpackLog(event, "RefundStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeBase contract.
type BridgeBaseRoleAdminChangedIterator struct {
	Event *BridgeBaseRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBaseRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseRoleAdminChanged)
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
		it.Event = new(BridgeBaseRoleAdminChanged)
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
func (it *BridgeBaseRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeBase contract.
type BridgeBaseRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeBase *BridgeBaseFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeBaseRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseRoleAdminChangedIterator{contract: _BridgeBase.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeBase *BridgeBaseFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeBaseRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseRoleAdminChanged)
				if err := _BridgeBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeBaseRoleAdminChanged, error) {
	event := new(BridgeBaseRoleAdminChanged)
	if err := _BridgeBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeBase contract.
type BridgeBaseRoleGrantedIterator struct {
	Event *BridgeBaseRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeBaseRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseRoleGranted)
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
		it.Event = new(BridgeBaseRoleGranted)
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
func (it *BridgeBaseRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseRoleGranted represents a RoleGranted event raised by the BridgeBase contract.
type BridgeBaseRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBase *BridgeBaseFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeBaseRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseRoleGrantedIterator{contract: _BridgeBase.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBase *BridgeBaseFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeBaseRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseRoleGranted)
				if err := _BridgeBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseRoleGranted(log types.Log) (*BridgeBaseRoleGranted, error) {
	event := new(BridgeBaseRoleGranted)
	if err := _BridgeBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeBase contract.
type BridgeBaseRoleRevokedIterator struct {
	Event *BridgeBaseRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeBaseRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseRoleRevoked)
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
		it.Event = new(BridgeBaseRoleRevoked)
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
func (it *BridgeBaseRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseRoleRevoked represents a RoleRevoked event raised by the BridgeBase contract.
type BridgeBaseRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBase *BridgeBaseFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeBaseRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseRoleRevokedIterator{contract: _BridgeBase.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBase *BridgeBaseFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeBaseRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseRoleRevoked)
				if err := _BridgeBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseRoleRevoked(log types.Log) (*BridgeBaseRoleRevoked, error) {
	event := new(BridgeBaseRoleRevoked)
	if err := _BridgeBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseTeleportIterator is returned from FilterTeleport and is used to iterate over the raw logs and unpacked data for Teleport events raised by the BridgeBase contract.
type BridgeBaseTeleportIterator struct {
	Event *BridgeBaseTeleport // Event containing the contract specifics and raw log

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
func (it *BridgeBaseTeleportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseTeleport)
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
		it.Event = new(BridgeBaseTeleport)
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
func (it *BridgeBaseTeleportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseTeleportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseTeleport represents a Teleport event raised by the BridgeBase contract.
type BridgeBaseTeleport struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTeleport is a free log retrieval operation binding the contract event 0x95dd69575b9e3de8b3ccfcc61a580c8b0123533a83ea8c43e12449837c80187b.
//
// Solidity: event Teleport(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) FilterTeleport(opts *bind.FilterOpts, from []common.Address, to []common.Address, nonce []*big.Int) (*BridgeBaseTeleportIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Teleport", fromRule, toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseTeleportIterator{contract: _BridgeBase.contract, event: "Teleport", logs: logs, sub: sub}, nil
}

// WatchTeleport is a free log subscription operation binding the contract event 0x95dd69575b9e3de8b3ccfcc61a580c8b0123533a83ea8c43e12449837c80187b.
//
// Solidity: event Teleport(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) WatchTeleport(opts *bind.WatchOpts, sink chan<- *BridgeBaseTeleport, from []common.Address, to []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Teleport", fromRule, toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseTeleport)
				if err := _BridgeBase.contract.UnpackLog(event, "Teleport", log); err != nil {
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

// ParseTeleport is a log parse operation binding the contract event 0x95dd69575b9e3de8b3ccfcc61a580c8b0123533a83ea8c43e12449837c80187b.
//
// Solidity: event Teleport(address indexed from, address indexed to, uint256 amount, uint256 indexed nonce)
func (_BridgeBase *BridgeBaseFilterer) ParseTeleport(log types.Log) (*BridgeBaseTeleport, error) {
	event := new(BridgeBaseTeleport)
	if err := _BridgeBase.contract.UnpackLog(event, "Teleport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeBase contract.
type BridgeBaseUnpausedIterator struct {
	Event *BridgeBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseUnpaused)
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
		it.Event = new(BridgeBaseUnpaused)
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
func (it *BridgeBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseUnpaused represents a Unpaused event raised by the BridgeBase contract.
type BridgeBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeBaseUnpausedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeBaseUnpausedIterator{contract: _BridgeBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseUnpaused)
				if err := _BridgeBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBase *BridgeBaseFilterer) ParseUnpaused(log types.Log) (*BridgeBaseUnpaused, error) {
	event := new(BridgeBaseUnpaused)
	if err := _BridgeBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
