// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ReceiptLibReceipt is an auto generated low-level Go binding around an user-defined struct.
type ReceiptLibReceipt struct {
	Id              *big.Int
	User            common.Address
	GroupId         *big.Int
	AmountInSatoshi *big.Int
	CreateTimestamp *big.Int
	TxId            [32]byte
	Height          *big.Int
	Status          uint8
	BtcAddress      string
}

// ReceiptControllerABI is the input ABI used to generate the binding from.
const ReceiptControllerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"DepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"}],\"name\":\"DepositRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"WithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REQUEST_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECEIPT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_receiptId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"depositReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInSatoshi\",\"type\":\"uint256\"}],\"name\":\"depositRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getAmountInSatoshi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getCreateTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getGroupId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getReceiptInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"enumReceiptLib.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"internalType\":\"structReceiptLib.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getReceiptStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"getWorkingReceiptId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"isGroupAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"isStale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_receiptId\",\"type\":\"uint256\"}],\"name\":\"revokeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_receiptId\",\"type\":\"uint256\"}],\"name\":\"withdrawCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"withdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReceiptController is an auto generated Go binding around an Ethereum contract.
type ReceiptController struct {
	ReceiptControllerCaller     // Read-only binding to the contract
	ReceiptControllerTransactor // Write-only binding to the contract
	ReceiptControllerFilterer   // Log filterer for contract events
}

// ReceiptControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptControllerSession struct {
	Contract     *ReceiptController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReceiptControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptControllerCallerSession struct {
	Contract *ReceiptControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReceiptControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptControllerTransactorSession struct {
	Contract     *ReceiptControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReceiptControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptControllerRaw struct {
	Contract *ReceiptController // Generic contract binding to access the raw methods on
}

// ReceiptControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptControllerCallerRaw struct {
	Contract *ReceiptControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptControllerTransactorRaw struct {
	Contract *ReceiptControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptController creates a new instance of ReceiptController, bound to a specific deployed contract.
func NewReceiptController(address common.Address, backend bind.ContractBackend) (*ReceiptController, error) {
	contract, err := bindReceiptController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptController{ReceiptControllerCaller: ReceiptControllerCaller{contract: contract}, ReceiptControllerTransactor: ReceiptControllerTransactor{contract: contract}, ReceiptControllerFilterer: ReceiptControllerFilterer{contract: contract}}, nil
}

// NewReceiptControllerCaller creates a new read-only instance of ReceiptController, bound to a specific deployed contract.
func NewReceiptControllerCaller(address common.Address, caller bind.ContractCaller) (*ReceiptControllerCaller, error) {
	contract, err := bindReceiptController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerCaller{contract: contract}, nil
}

// NewReceiptControllerTransactor creates a new write-only instance of ReceiptController, bound to a specific deployed contract.
func NewReceiptControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptControllerTransactor, error) {
	contract, err := bindReceiptController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerTransactor{contract: contract}, nil
}

// NewReceiptControllerFilterer creates a new log filterer instance of ReceiptController, bound to a specific deployed contract.
func NewReceiptControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptControllerFilterer, error) {
	contract, err := bindReceiptController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerFilterer{contract: contract}, nil
}

// bindReceiptController binds a generic wrapper to an already deployed contract.
func bindReceiptController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiptControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptController *ReceiptControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptController.Contract.ReceiptControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptController *ReceiptControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptController.Contract.ReceiptControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptController *ReceiptControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptController.Contract.ReceiptControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptController *ReceiptControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptController *ReceiptControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptController *ReceiptControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptController.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ReceiptController.Contract.DEFAULTADMINROLE(&_ReceiptController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ReceiptController.Contract.DEFAULTADMINROLE(&_ReceiptController.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) MINTREQUESTGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "MINT_REQUEST_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) MINTREQUESTGRACEPERIOD() (*big.Int, error) {
	return _ReceiptController.Contract.MINTREQUESTGRACEPERIOD(&_ReceiptController.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) MINTREQUESTGRACEPERIOD() (*big.Int, error) {
	return _ReceiptController.Contract.MINTREQUESTGRACEPERIOD(&_ReceiptController.CallOpts)
}

// RECEIPTADMINROLE is a free data retrieval call binding the contract method 0xa8c3b5e6.
//
// Solidity: function RECEIPT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerCaller) RECEIPTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "RECEIPT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECEIPTADMINROLE is a free data retrieval call binding the contract method 0xa8c3b5e6.
//
// Solidity: function RECEIPT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerSession) RECEIPTADMINROLE() ([32]byte, error) {
	return _ReceiptController.Contract.RECEIPTADMINROLE(&_ReceiptController.CallOpts)
}

// RECEIPTADMINROLE is a free data retrieval call binding the contract method 0xa8c3b5e6.
//
// Solidity: function RECEIPT_ADMIN_ROLE() view returns(bytes32)
func (_ReceiptController *ReceiptControllerCallerSession) RECEIPTADMINROLE() ([32]byte, error) {
	return _ReceiptController.Contract.RECEIPTADMINROLE(&_ReceiptController.CallOpts)
}

// GetAmountInSatoshi is a free data retrieval call binding the contract method 0xe6927c3b.
//
// Solidity: function getAmountInSatoshi(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetAmountInSatoshi(opts *bind.CallOpts, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getAmountInSatoshi", receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountInSatoshi is a free data retrieval call binding the contract method 0xe6927c3b.
//
// Solidity: function getAmountInSatoshi(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetAmountInSatoshi(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetAmountInSatoshi(&_ReceiptController.CallOpts, receiptId)
}

// GetAmountInSatoshi is a free data retrieval call binding the contract method 0xe6927c3b.
//
// Solidity: function getAmountInSatoshi(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetAmountInSatoshi(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetAmountInSatoshi(&_ReceiptController.CallOpts, receiptId)
}

// GetCreateTimestamp is a free data retrieval call binding the contract method 0x83f93962.
//
// Solidity: function getCreateTimestamp(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetCreateTimestamp(opts *bind.CallOpts, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getCreateTimestamp", receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreateTimestamp is a free data retrieval call binding the contract method 0x83f93962.
//
// Solidity: function getCreateTimestamp(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetCreateTimestamp(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetCreateTimestamp(&_ReceiptController.CallOpts, receiptId)
}

// GetCreateTimestamp is a free data retrieval call binding the contract method 0x83f93962.
//
// Solidity: function getCreateTimestamp(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetCreateTimestamp(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetCreateTimestamp(&_ReceiptController.CallOpts, receiptId)
}

// GetGroupId is a free data retrieval call binding the contract method 0x99a29d74.
//
// Solidity: function getGroupId(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetGroupId(opts *bind.CallOpts, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getGroupId", receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupId is a free data retrieval call binding the contract method 0x99a29d74.
//
// Solidity: function getGroupId(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetGroupId(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetGroupId(&_ReceiptController.CallOpts, receiptId)
}

// GetGroupId is a free data retrieval call binding the contract method 0x99a29d74.
//
// Solidity: function getGroupId(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetGroupId(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetGroupId(&_ReceiptController.CallOpts, receiptId)
}

// GetReceiptInfo is a free data retrieval call binding the contract method 0xb97912ff.
//
// Solidity: function getReceiptInfo(uint256 receiptId) view returns((uint256,address,uint256,uint256,uint256,bytes32,uint256,uint8,string))
func (_ReceiptController *ReceiptControllerCaller) GetReceiptInfo(opts *bind.CallOpts, receiptId *big.Int) (ReceiptLibReceipt, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getReceiptInfo", receiptId)

	if err != nil {
		return *new(ReceiptLibReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(ReceiptLibReceipt)).(*ReceiptLibReceipt)

	return out0, err

}

// GetReceiptInfo is a free data retrieval call binding the contract method 0xb97912ff.
//
// Solidity: function getReceiptInfo(uint256 receiptId) view returns((uint256,address,uint256,uint256,uint256,bytes32,uint256,uint8,string))
func (_ReceiptController *ReceiptControllerSession) GetReceiptInfo(receiptId *big.Int) (ReceiptLibReceipt, error) {
	return _ReceiptController.Contract.GetReceiptInfo(&_ReceiptController.CallOpts, receiptId)
}

// GetReceiptInfo is a free data retrieval call binding the contract method 0xb97912ff.
//
// Solidity: function getReceiptInfo(uint256 receiptId) view returns((uint256,address,uint256,uint256,uint256,bytes32,uint256,uint8,string))
func (_ReceiptController *ReceiptControllerCallerSession) GetReceiptInfo(receiptId *big.Int) (ReceiptLibReceipt, error) {
	return _ReceiptController.Contract.GetReceiptInfo(&_ReceiptController.CallOpts, receiptId)
}

// GetReceiptStatus is a free data retrieval call binding the contract method 0x5fd7c24a.
//
// Solidity: function getReceiptStatus(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetReceiptStatus(opts *bind.CallOpts, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getReceiptStatus", receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptStatus is a free data retrieval call binding the contract method 0x5fd7c24a.
//
// Solidity: function getReceiptStatus(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetReceiptStatus(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetReceiptStatus(&_ReceiptController.CallOpts, receiptId)
}

// GetReceiptStatus is a free data retrieval call binding the contract method 0x5fd7c24a.
//
// Solidity: function getReceiptStatus(uint256 receiptId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetReceiptStatus(receiptId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetReceiptStatus(&_ReceiptController.CallOpts, receiptId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ReceiptController *ReceiptControllerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ReceiptController *ReceiptControllerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ReceiptController.Contract.GetRoleAdmin(&_ReceiptController.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ReceiptController *ReceiptControllerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ReceiptController.Contract.GetRoleAdmin(&_ReceiptController.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ReceiptController *ReceiptControllerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ReceiptController *ReceiptControllerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ReceiptController.Contract.GetRoleMember(&_ReceiptController.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ReceiptController *ReceiptControllerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ReceiptController.Contract.GetRoleMember(&_ReceiptController.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ReceiptController.Contract.GetRoleMemberCount(&_ReceiptController.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ReceiptController.Contract.GetRoleMemberCount(&_ReceiptController.CallOpts, role)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x90d976e2.
//
// Solidity: function getUserAddress(uint256 receiptId) view returns(address)
func (_ReceiptController *ReceiptControllerCaller) GetUserAddress(opts *bind.CallOpts, receiptId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getUserAddress", receiptId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAddress is a free data retrieval call binding the contract method 0x90d976e2.
//
// Solidity: function getUserAddress(uint256 receiptId) view returns(address)
func (_ReceiptController *ReceiptControllerSession) GetUserAddress(receiptId *big.Int) (common.Address, error) {
	return _ReceiptController.Contract.GetUserAddress(&_ReceiptController.CallOpts, receiptId)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x90d976e2.
//
// Solidity: function getUserAddress(uint256 receiptId) view returns(address)
func (_ReceiptController *ReceiptControllerCallerSession) GetUserAddress(receiptId *big.Int) (common.Address, error) {
	return _ReceiptController.Contract.GetUserAddress(&_ReceiptController.CallOpts, receiptId)
}

// GetWorkingReceiptId is a free data retrieval call binding the contract method 0x10eb6b13.
//
// Solidity: function getWorkingReceiptId(uint256 groupId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCaller) GetWorkingReceiptId(opts *bind.CallOpts, groupId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "getWorkingReceiptId", groupId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkingReceiptId is a free data retrieval call binding the contract method 0x10eb6b13.
//
// Solidity: function getWorkingReceiptId(uint256 groupId) view returns(uint256)
func (_ReceiptController *ReceiptControllerSession) GetWorkingReceiptId(groupId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetWorkingReceiptId(&_ReceiptController.CallOpts, groupId)
}

// GetWorkingReceiptId is a free data retrieval call binding the contract method 0x10eb6b13.
//
// Solidity: function getWorkingReceiptId(uint256 groupId) view returns(uint256)
func (_ReceiptController *ReceiptControllerCallerSession) GetWorkingReceiptId(groupId *big.Int) (*big.Int, error) {
	return _ReceiptController.Contract.GetWorkingReceiptId(&_ReceiptController.CallOpts, groupId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ReceiptController *ReceiptControllerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ReceiptController *ReceiptControllerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ReceiptController.Contract.HasRole(&_ReceiptController.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ReceiptController *ReceiptControllerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ReceiptController.Contract.HasRole(&_ReceiptController.CallOpts, role, account)
}

// IsGroupAvailable is a free data retrieval call binding the contract method 0xf5d52663.
//
// Solidity: function isGroupAvailable(uint256 groupId) view returns(bool)
func (_ReceiptController *ReceiptControllerCaller) IsGroupAvailable(opts *bind.CallOpts, groupId *big.Int) (bool, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "isGroupAvailable", groupId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGroupAvailable is a free data retrieval call binding the contract method 0xf5d52663.
//
// Solidity: function isGroupAvailable(uint256 groupId) view returns(bool)
func (_ReceiptController *ReceiptControllerSession) IsGroupAvailable(groupId *big.Int) (bool, error) {
	return _ReceiptController.Contract.IsGroupAvailable(&_ReceiptController.CallOpts, groupId)
}

// IsGroupAvailable is a free data retrieval call binding the contract method 0xf5d52663.
//
// Solidity: function isGroupAvailable(uint256 groupId) view returns(bool)
func (_ReceiptController *ReceiptControllerCallerSession) IsGroupAvailable(groupId *big.Int) (bool, error) {
	return _ReceiptController.Contract.IsGroupAvailable(&_ReceiptController.CallOpts, groupId)
}

// IsStale is a free data retrieval call binding the contract method 0xa624a1df.
//
// Solidity: function isStale(uint256 receiptId) view returns(bool)
func (_ReceiptController *ReceiptControllerCaller) IsStale(opts *bind.CallOpts, receiptId *big.Int) (bool, error) {
	var out []interface{}
	err := _ReceiptController.contract.Call(opts, &out, "isStale", receiptId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStale is a free data retrieval call binding the contract method 0xa624a1df.
//
// Solidity: function isStale(uint256 receiptId) view returns(bool)
func (_ReceiptController *ReceiptControllerSession) IsStale(receiptId *big.Int) (bool, error) {
	return _ReceiptController.Contract.IsStale(&_ReceiptController.CallOpts, receiptId)
}

// IsStale is a free data retrieval call binding the contract method 0xa624a1df.
//
// Solidity: function isStale(uint256 receiptId) view returns(bool)
func (_ReceiptController *ReceiptControllerCallerSession) IsStale(receiptId *big.Int) (bool, error) {
	return _ReceiptController.Contract.IsStale(&_ReceiptController.CallOpts, receiptId)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x15069a2f.
//
// Solidity: function depositReceived(uint256 _receiptId, bytes32 _txId, uint256 _height) returns()
func (_ReceiptController *ReceiptControllerTransactor) DepositReceived(opts *bind.TransactOpts, _receiptId *big.Int, _txId [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "depositReceived", _receiptId, _txId, _height)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x15069a2f.
//
// Solidity: function depositReceived(uint256 _receiptId, bytes32 _txId, uint256 _height) returns()
func (_ReceiptController *ReceiptControllerSession) DepositReceived(_receiptId *big.Int, _txId [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.DepositReceived(&_ReceiptController.TransactOpts, _receiptId, _txId, _height)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x15069a2f.
//
// Solidity: function depositReceived(uint256 _receiptId, bytes32 _txId, uint256 _height) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) DepositReceived(_receiptId *big.Int, _txId [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.DepositReceived(&_ReceiptController.TransactOpts, _receiptId, _txId, _height)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x54783d77.
//
// Solidity: function depositRequest(address _user, uint256 _groupId, uint256 _amountInSatoshi) returns(uint256)
func (_ReceiptController *ReceiptControllerTransactor) DepositRequest(opts *bind.TransactOpts, _user common.Address, _groupId *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "depositRequest", _user, _groupId, _amountInSatoshi)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x54783d77.
//
// Solidity: function depositRequest(address _user, uint256 _groupId, uint256 _amountInSatoshi) returns(uint256)
func (_ReceiptController *ReceiptControllerSession) DepositRequest(_user common.Address, _groupId *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.DepositRequest(&_ReceiptController.TransactOpts, _user, _groupId, _amountInSatoshi)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x54783d77.
//
// Solidity: function depositRequest(address _user, uint256 _groupId, uint256 _amountInSatoshi) returns(uint256)
func (_ReceiptController *ReceiptControllerTransactorSession) DepositRequest(_user common.Address, _groupId *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.DepositRequest(&_ReceiptController.TransactOpts, _user, _groupId, _amountInSatoshi)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.GrantRole(&_ReceiptController.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.GrantRole(&_ReceiptController.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.RenounceRole(&_ReceiptController.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.RenounceRole(&_ReceiptController.TransactOpts, role, account)
}

// RevokeRequest is a paid mutator transaction binding the contract method 0xe7d915cf.
//
// Solidity: function revokeRequest(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerTransactor) RevokeRequest(opts *bind.TransactOpts, _receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "revokeRequest", _receiptId)
}

// RevokeRequest is a paid mutator transaction binding the contract method 0xe7d915cf.
//
// Solidity: function revokeRequest(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerSession) RevokeRequest(_receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.RevokeRequest(&_ReceiptController.TransactOpts, _receiptId)
}

// RevokeRequest is a paid mutator transaction binding the contract method 0xe7d915cf.
//
// Solidity: function revokeRequest(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) RevokeRequest(_receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.RevokeRequest(&_ReceiptController.TransactOpts, _receiptId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.RevokeRole(&_ReceiptController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ReceiptController.Contract.RevokeRole(&_ReceiptController.TransactOpts, role, account)
}

// WithdrawCompleted is a paid mutator transaction binding the contract method 0xa88b124a.
//
// Solidity: function withdrawCompleted(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerTransactor) WithdrawCompleted(opts *bind.TransactOpts, _receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "withdrawCompleted", _receiptId)
}

// WithdrawCompleted is a paid mutator transaction binding the contract method 0xa88b124a.
//
// Solidity: function withdrawCompleted(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerSession) WithdrawCompleted(_receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.WithdrawCompleted(&_ReceiptController.TransactOpts, _receiptId)
}

// WithdrawCompleted is a paid mutator transaction binding the contract method 0xa88b124a.
//
// Solidity: function withdrawCompleted(uint256 _receiptId) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) WithdrawCompleted(_receiptId *big.Int) (*types.Transaction, error) {
	return _ReceiptController.Contract.WithdrawCompleted(&_ReceiptController.TransactOpts, _receiptId)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x395faadb.
//
// Solidity: function withdrawRequest(uint256 receiptId, string btcAddress) returns()
func (_ReceiptController *ReceiptControllerTransactor) WithdrawRequest(opts *bind.TransactOpts, receiptId *big.Int, btcAddress string) (*types.Transaction, error) {
	return _ReceiptController.contract.Transact(opts, "withdrawRequest", receiptId, btcAddress)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x395faadb.
//
// Solidity: function withdrawRequest(uint256 receiptId, string btcAddress) returns()
func (_ReceiptController *ReceiptControllerSession) WithdrawRequest(receiptId *big.Int, btcAddress string) (*types.Transaction, error) {
	return _ReceiptController.Contract.WithdrawRequest(&_ReceiptController.TransactOpts, receiptId, btcAddress)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x395faadb.
//
// Solidity: function withdrawRequest(uint256 receiptId, string btcAddress) returns()
func (_ReceiptController *ReceiptControllerTransactorSession) WithdrawRequest(receiptId *big.Int, btcAddress string) (*types.Transaction, error) {
	return _ReceiptController.Contract.WithdrawRequest(&_ReceiptController.TransactOpts, receiptId, btcAddress)
}

// ReceiptControllerDepositReceivedIterator is returned from FilterDepositReceived and is used to iterate over the raw logs and unpacked data for DepositReceived events raised by the ReceiptController contract.
type ReceiptControllerDepositReceivedIterator struct {
	Event *ReceiptControllerDepositReceived // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerDepositReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerDepositReceived)
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
		it.Event = new(ReceiptControllerDepositReceived)
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
func (it *ReceiptControllerDepositReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerDepositReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerDepositReceived represents a DepositReceived event raised by the ReceiptController contract.
type ReceiptControllerDepositReceived struct {
	ReceiptId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositReceived is a free log retrieval operation binding the contract event 0xa1ab46d0e38f33a7b1f03853088a305927deb55c906370ff5d6d2bc732f15095.
//
// Solidity: event DepositReceived(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) FilterDepositReceived(opts *bind.FilterOpts, receiptId []*big.Int) (*ReceiptControllerDepositReceivedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "DepositReceived", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerDepositReceivedIterator{contract: _ReceiptController.contract, event: "DepositReceived", logs: logs, sub: sub}, nil
}

// WatchDepositReceived is a free log subscription operation binding the contract event 0xa1ab46d0e38f33a7b1f03853088a305927deb55c906370ff5d6d2bc732f15095.
//
// Solidity: event DepositReceived(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) WatchDepositReceived(opts *bind.WatchOpts, sink chan<- *ReceiptControllerDepositReceived, receiptId []*big.Int) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "DepositReceived", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerDepositReceived)
				if err := _ReceiptController.contract.UnpackLog(event, "DepositReceived", log); err != nil {
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

// ParseDepositReceived is a log parse operation binding the contract event 0xa1ab46d0e38f33a7b1f03853088a305927deb55c906370ff5d6d2bc732f15095.
//
// Solidity: event DepositReceived(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) ParseDepositReceived(log types.Log) (*ReceiptControllerDepositReceived, error) {
	event := new(ReceiptControllerDepositReceived)
	if err := _ReceiptController.contract.UnpackLog(event, "DepositReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerDepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the ReceiptController contract.
type ReceiptControllerDepositRequestedIterator struct {
	Event *ReceiptControllerDepositRequested // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerDepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerDepositRequested)
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
		it.Event = new(ReceiptControllerDepositRequested)
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
func (it *ReceiptControllerDepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerDepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerDepositRequested represents a DepositRequested event raised by the ReceiptController contract.
type ReceiptControllerDepositRequested struct {
	ReceiptId *big.Int
	GroupId   *big.Int
	User      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0x1ef85c319750d68749b20f896d54f19ca53269aeb8980a66be107b54a3d510ab.
//
// Solidity: event DepositRequested(uint256 indexed receiptId, uint256 indexed groupId, address indexed user, uint256 amount)
func (_ReceiptController *ReceiptControllerFilterer) FilterDepositRequested(opts *bind.FilterOpts, receiptId []*big.Int, groupId []*big.Int, user []common.Address) (*ReceiptControllerDepositRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "DepositRequested", receiptIdRule, groupIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerDepositRequestedIterator{contract: _ReceiptController.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0x1ef85c319750d68749b20f896d54f19ca53269aeb8980a66be107b54a3d510ab.
//
// Solidity: event DepositRequested(uint256 indexed receiptId, uint256 indexed groupId, address indexed user, uint256 amount)
func (_ReceiptController *ReceiptControllerFilterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *ReceiptControllerDepositRequested, receiptId []*big.Int, groupId []*big.Int, user []common.Address) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "DepositRequested", receiptIdRule, groupIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerDepositRequested)
				if err := _ReceiptController.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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

// ParseDepositRequested is a log parse operation binding the contract event 0x1ef85c319750d68749b20f896d54f19ca53269aeb8980a66be107b54a3d510ab.
//
// Solidity: event DepositRequested(uint256 indexed receiptId, uint256 indexed groupId, address indexed user, uint256 amount)
func (_ReceiptController *ReceiptControllerFilterer) ParseDepositRequested(log types.Log) (*ReceiptControllerDepositRequested, error) {
	event := new(ReceiptControllerDepositRequested)
	if err := _ReceiptController.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerDepositRevokedIterator is returned from FilterDepositRevoked and is used to iterate over the raw logs and unpacked data for DepositRevoked events raised by the ReceiptController contract.
type ReceiptControllerDepositRevokedIterator struct {
	Event *ReceiptControllerDepositRevoked // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerDepositRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerDepositRevoked)
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
		it.Event = new(ReceiptControllerDepositRevoked)
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
func (it *ReceiptControllerDepositRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerDepositRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerDepositRevoked represents a DepositRevoked event raised by the ReceiptController contract.
type ReceiptControllerDepositRevoked struct {
	ReceiptId *big.Int
	Revoker   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositRevoked is a free log retrieval operation binding the contract event 0x5e6a6da2b39d1f3dee492fe933902f81da290ef69453e0dae0e33d84ebd29600.
//
// Solidity: event DepositRevoked(uint256 indexed receiptId, address indexed revoker)
func (_ReceiptController *ReceiptControllerFilterer) FilterDepositRevoked(opts *bind.FilterOpts, receiptId []*big.Int, revoker []common.Address) (*ReceiptControllerDepositRevokedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "DepositRevoked", receiptIdRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerDepositRevokedIterator{contract: _ReceiptController.contract, event: "DepositRevoked", logs: logs, sub: sub}, nil
}

// WatchDepositRevoked is a free log subscription operation binding the contract event 0x5e6a6da2b39d1f3dee492fe933902f81da290ef69453e0dae0e33d84ebd29600.
//
// Solidity: event DepositRevoked(uint256 indexed receiptId, address indexed revoker)
func (_ReceiptController *ReceiptControllerFilterer) WatchDepositRevoked(opts *bind.WatchOpts, sink chan<- *ReceiptControllerDepositRevoked, receiptId []*big.Int, revoker []common.Address) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "DepositRevoked", receiptIdRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerDepositRevoked)
				if err := _ReceiptController.contract.UnpackLog(event, "DepositRevoked", log); err != nil {
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

// ParseDepositRevoked is a log parse operation binding the contract event 0x5e6a6da2b39d1f3dee492fe933902f81da290ef69453e0dae0e33d84ebd29600.
//
// Solidity: event DepositRevoked(uint256 indexed receiptId, address indexed revoker)
func (_ReceiptController *ReceiptControllerFilterer) ParseDepositRevoked(log types.Log) (*ReceiptControllerDepositRevoked, error) {
	event := new(ReceiptControllerDepositRevoked)
	if err := _ReceiptController.contract.UnpackLog(event, "DepositRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ReceiptController contract.
type ReceiptControllerRoleAdminChangedIterator struct {
	Event *ReceiptControllerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerRoleAdminChanged)
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
		it.Event = new(ReceiptControllerRoleAdminChanged)
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
func (it *ReceiptControllerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerRoleAdminChanged represents a RoleAdminChanged event raised by the ReceiptController contract.
type ReceiptControllerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ReceiptController *ReceiptControllerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ReceiptControllerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerRoleAdminChangedIterator{contract: _ReceiptController.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ReceiptController *ReceiptControllerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ReceiptControllerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerRoleAdminChanged)
				if err := _ReceiptController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ReceiptController *ReceiptControllerFilterer) ParseRoleAdminChanged(log types.Log) (*ReceiptControllerRoleAdminChanged, error) {
	event := new(ReceiptControllerRoleAdminChanged)
	if err := _ReceiptController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ReceiptController contract.
type ReceiptControllerRoleGrantedIterator struct {
	Event *ReceiptControllerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerRoleGranted)
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
		it.Event = new(ReceiptControllerRoleGranted)
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
func (it *ReceiptControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerRoleGranted represents a RoleGranted event raised by the ReceiptController contract.
type ReceiptControllerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ReceiptController *ReceiptControllerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ReceiptControllerRoleGrantedIterator, error) {

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

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerRoleGrantedIterator{contract: _ReceiptController.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ReceiptController *ReceiptControllerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ReceiptControllerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerRoleGranted)
				if err := _ReceiptController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ReceiptController *ReceiptControllerFilterer) ParseRoleGranted(log types.Log) (*ReceiptControllerRoleGranted, error) {
	event := new(ReceiptControllerRoleGranted)
	if err := _ReceiptController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ReceiptController contract.
type ReceiptControllerRoleRevokedIterator struct {
	Event *ReceiptControllerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerRoleRevoked)
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
		it.Event = new(ReceiptControllerRoleRevoked)
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
func (it *ReceiptControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerRoleRevoked represents a RoleRevoked event raised by the ReceiptController contract.
type ReceiptControllerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ReceiptController *ReceiptControllerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ReceiptControllerRoleRevokedIterator, error) {

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

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerRoleRevokedIterator{contract: _ReceiptController.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ReceiptController *ReceiptControllerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ReceiptControllerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerRoleRevoked)
				if err := _ReceiptController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ReceiptController *ReceiptControllerFilterer) ParseRoleRevoked(log types.Log) (*ReceiptControllerRoleRevoked, error) {
	event := new(ReceiptControllerRoleRevoked)
	if err := _ReceiptController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerWithdrawCompletedIterator is returned from FilterWithdrawCompleted and is used to iterate over the raw logs and unpacked data for WithdrawCompleted events raised by the ReceiptController contract.
type ReceiptControllerWithdrawCompletedIterator struct {
	Event *ReceiptControllerWithdrawCompleted // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerWithdrawCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerWithdrawCompleted)
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
		it.Event = new(ReceiptControllerWithdrawCompleted)
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
func (it *ReceiptControllerWithdrawCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerWithdrawCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerWithdrawCompleted represents a WithdrawCompleted event raised by the ReceiptController contract.
type ReceiptControllerWithdrawCompleted struct {
	ReceiptId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCompleted is a free log retrieval operation binding the contract event 0xf5c5432de34b3d8107ca8b5be5379cdfbb805a494446d6f3a0b6d1714a3e2dc7.
//
// Solidity: event WithdrawCompleted(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) FilterWithdrawCompleted(opts *bind.FilterOpts, receiptId []*big.Int) (*ReceiptControllerWithdrawCompletedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "WithdrawCompleted", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerWithdrawCompletedIterator{contract: _ReceiptController.contract, event: "WithdrawCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawCompleted is a free log subscription operation binding the contract event 0xf5c5432de34b3d8107ca8b5be5379cdfbb805a494446d6f3a0b6d1714a3e2dc7.
//
// Solidity: event WithdrawCompleted(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) WatchWithdrawCompleted(opts *bind.WatchOpts, sink chan<- *ReceiptControllerWithdrawCompleted, receiptId []*big.Int) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "WithdrawCompleted", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerWithdrawCompleted)
				if err := _ReceiptController.contract.UnpackLog(event, "WithdrawCompleted", log); err != nil {
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

// ParseWithdrawCompleted is a log parse operation binding the contract event 0xf5c5432de34b3d8107ca8b5be5379cdfbb805a494446d6f3a0b6d1714a3e2dc7.
//
// Solidity: event WithdrawCompleted(uint256 indexed receiptId)
func (_ReceiptController *ReceiptControllerFilterer) ParseWithdrawCompleted(log types.Log) (*ReceiptControllerWithdrawCompleted, error) {
	event := new(ReceiptControllerWithdrawCompleted)
	if err := _ReceiptController.contract.UnpackLog(event, "WithdrawCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptControllerWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the ReceiptController contract.
type ReceiptControllerWithdrawRequestedIterator struct {
	Event *ReceiptControllerWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ReceiptControllerWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiptControllerWithdrawRequested)
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
		it.Event = new(ReceiptControllerWithdrawRequested)
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
func (it *ReceiptControllerWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiptControllerWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiptControllerWithdrawRequested represents a WithdrawRequested event raised by the ReceiptController contract.
type ReceiptControllerWithdrawRequested struct {
	ReceiptId  *big.Int
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0x16164b38adc054b584fc6ddb4a7aecbb14802344b2881b11d728d39200bbaff2.
//
// Solidity: event WithdrawRequested(uint256 indexed receiptId, string btcAddress)
func (_ReceiptController *ReceiptControllerFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, receiptId []*big.Int) (*ReceiptControllerWithdrawRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.FilterLogs(opts, "WithdrawRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &ReceiptControllerWithdrawRequestedIterator{contract: _ReceiptController.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0x16164b38adc054b584fc6ddb4a7aecbb14802344b2881b11d728d39200bbaff2.
//
// Solidity: event WithdrawRequested(uint256 indexed receiptId, string btcAddress)
func (_ReceiptController *ReceiptControllerFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ReceiptControllerWithdrawRequested, receiptId []*big.Int) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _ReceiptController.contract.WatchLogs(opts, "WithdrawRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiptControllerWithdrawRequested)
				if err := _ReceiptController.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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

// ParseWithdrawRequested is a log parse operation binding the contract event 0x16164b38adc054b584fc6ddb4a7aecbb14802344b2881b11d728d39200bbaff2.
//
// Solidity: event WithdrawRequested(uint256 indexed receiptId, string btcAddress)
func (_ReceiptController *ReceiptControllerFilterer) ParseWithdrawRequested(log types.Log) (*ReceiptControllerWithdrawRequested, error) {
	event := new(ReceiptControllerWithdrawRequested)
	if err := _ReceiptController.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
