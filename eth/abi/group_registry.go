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

// GroupRegistryABI is the input ABI used to generate the binding from.
const GroupRegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"group_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSatoshi\",\"type\":\"uint256\"}],\"name\":\"GroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GroupDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_keepers\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"_btcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxSatoshi\",\"type\":\"uint256\"}],\"name\":\"addGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInSatoshi\",\"type\":\"uint256\"}],\"name\":\"depositReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getGroupAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_btcAddress\",\"type\":\"string\"}],\"name\":\"getGroupId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getGroupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastWithdrawTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getGroupKeepers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"}],\"name\":\"getKeeperGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"isGroupKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInSatoshi\",\"type\":\"uint256\"}],\"name\":\"withdrawRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GroupRegistry is an auto generated Go binding around an Ethereum contract.
type GroupRegistry struct {
	GroupRegistryCaller     // Read-only binding to the contract
	GroupRegistryTransactor // Write-only binding to the contract
	GroupRegistryFilterer   // Log filterer for contract events
}

// GroupRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GroupRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GroupRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GroupRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GroupRegistrySession struct {
	Contract     *GroupRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GroupRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GroupRegistryCallerSession struct {
	Contract *GroupRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GroupRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GroupRegistryTransactorSession struct {
	Contract     *GroupRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GroupRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GroupRegistryRaw struct {
	Contract *GroupRegistry // Generic contract binding to access the raw methods on
}

// GroupRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GroupRegistryCallerRaw struct {
	Contract *GroupRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// GroupRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GroupRegistryTransactorRaw struct {
	Contract *GroupRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGroupRegistry creates a new instance of GroupRegistry, bound to a specific deployed contract.
func NewGroupRegistry(address common.Address, backend bind.ContractBackend) (*GroupRegistry, error) {
	contract, err := bindGroupRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GroupRegistry{GroupRegistryCaller: GroupRegistryCaller{contract: contract}, GroupRegistryTransactor: GroupRegistryTransactor{contract: contract}, GroupRegistryFilterer: GroupRegistryFilterer{contract: contract}}, nil
}

// NewGroupRegistryCaller creates a new read-only instance of GroupRegistry, bound to a specific deployed contract.
func NewGroupRegistryCaller(address common.Address, caller bind.ContractCaller) (*GroupRegistryCaller, error) {
	contract, err := bindGroupRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryCaller{contract: contract}, nil
}

// NewGroupRegistryTransactor creates a new write-only instance of GroupRegistry, bound to a specific deployed contract.
func NewGroupRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*GroupRegistryTransactor, error) {
	contract, err := bindGroupRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryTransactor{contract: contract}, nil
}

// NewGroupRegistryFilterer creates a new log filterer instance of GroupRegistry, bound to a specific deployed contract.
func NewGroupRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*GroupRegistryFilterer, error) {
	contract, err := bindGroupRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryFilterer{contract: contract}, nil
}

// bindGroupRegistry binds a generic wrapper to an already deployed contract.
func bindGroupRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GroupRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GroupRegistry *GroupRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GroupRegistry.Contract.GroupRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GroupRegistry *GroupRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GroupRegistry.Contract.GroupRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GroupRegistry *GroupRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GroupRegistry.Contract.GroupRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GroupRegistry *GroupRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GroupRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GroupRegistry *GroupRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GroupRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GroupRegistry *GroupRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GroupRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GroupRegistry.Contract.DEFAULTADMINROLE(&_GroupRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GroupRegistry.Contract.DEFAULTADMINROLE(&_GroupRegistry.CallOpts)
}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistryCaller) GROUPADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "GROUP_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistrySession) GROUPADMINROLE() ([32]byte, error) {
	return _GroupRegistry.Contract.GROUPADMINROLE(&_GroupRegistry.CallOpts)
}

// GROUPADMINROLE is a free data retrieval call binding the contract method 0x0cec8e7a.
//
// Solidity: function GROUP_ADMIN_ROLE() view returns(bytes32)
func (_GroupRegistry *GroupRegistryCallerSession) GROUPADMINROLE() ([32]byte, error) {
	return _GroupRegistry.Contract.GROUPADMINROLE(&_GroupRegistry.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x4ebbc92a.
//
// Solidity: function exist(uint256 _id) view returns(bool)
func (_GroupRegistry *GroupRegistryCaller) Exist(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "exist", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x4ebbc92a.
//
// Solidity: function exist(uint256 _id) view returns(bool)
func (_GroupRegistry *GroupRegistrySession) Exist(_id *big.Int) (bool, error) {
	return _GroupRegistry.Contract.Exist(&_GroupRegistry.CallOpts, _id)
}

// Exist is a free data retrieval call binding the contract method 0x4ebbc92a.
//
// Solidity: function exist(uint256 _id) view returns(bool)
func (_GroupRegistry *GroupRegistryCallerSession) Exist(_id *big.Int) (bool, error) {
	return _GroupRegistry.Contract.Exist(&_GroupRegistry.CallOpts, _id)
}

// GetGroupAllowance is a free data retrieval call binding the contract method 0xb5878059.
//
// Solidity: function getGroupAllowance(uint256 _id) view returns(uint256)
func (_GroupRegistry *GroupRegistryCaller) GetGroupAllowance(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getGroupAllowance", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupAllowance is a free data retrieval call binding the contract method 0xb5878059.
//
// Solidity: function getGroupAllowance(uint256 _id) view returns(uint256)
func (_GroupRegistry *GroupRegistrySession) GetGroupAllowance(_id *big.Int) (*big.Int, error) {
	return _GroupRegistry.Contract.GetGroupAllowance(&_GroupRegistry.CallOpts, _id)
}

// GetGroupAllowance is a free data retrieval call binding the contract method 0xb5878059.
//
// Solidity: function getGroupAllowance(uint256 _id) view returns(uint256)
func (_GroupRegistry *GroupRegistryCallerSession) GetGroupAllowance(_id *big.Int) (*big.Int, error) {
	return _GroupRegistry.Contract.GetGroupAllowance(&_GroupRegistry.CallOpts, _id)
}

// GetGroupId is a free data retrieval call binding the contract method 0x52127000.
//
// Solidity: function getGroupId(string _btcAddress) view returns(uint256)
func (_GroupRegistry *GroupRegistryCaller) GetGroupId(opts *bind.CallOpts, _btcAddress string) (*big.Int, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getGroupId", _btcAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupId is a free data retrieval call binding the contract method 0x52127000.
//
// Solidity: function getGroupId(string _btcAddress) view returns(uint256)
func (_GroupRegistry *GroupRegistrySession) GetGroupId(_btcAddress string) (*big.Int, error) {
	return _GroupRegistry.Contract.GetGroupId(&_GroupRegistry.CallOpts, _btcAddress)
}

// GetGroupId is a free data retrieval call binding the contract method 0x52127000.
//
// Solidity: function getGroupId(string _btcAddress) view returns(uint256)
func (_GroupRegistry *GroupRegistryCallerSession) GetGroupId(_btcAddress string) (*big.Int, error) {
	return _GroupRegistry.Contract.GetGroupId(&_GroupRegistry.CallOpts, _btcAddress)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x99d4360f.
//
// Solidity: function getGroupInfo(uint256 _id) view returns(uint256 maxSatoshi, uint256 currSatoshi, uint256 lastWithdrawTimestamp, string btcAddress, address[] keepers)
func (_GroupRegistry *GroupRegistryCaller) GetGroupInfo(opts *bind.CallOpts, _id *big.Int) (struct {
	MaxSatoshi            *big.Int
	CurrSatoshi           *big.Int
	LastWithdrawTimestamp *big.Int
	BtcAddress            string
	Keepers               []common.Address
}, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getGroupInfo", _id)

	outstruct := new(struct {
		MaxSatoshi            *big.Int
		CurrSatoshi           *big.Int
		LastWithdrawTimestamp *big.Int
		BtcAddress            string
		Keepers               []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxSatoshi = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrSatoshi = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastWithdrawTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BtcAddress = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Keepers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetGroupInfo is a free data retrieval call binding the contract method 0x99d4360f.
//
// Solidity: function getGroupInfo(uint256 _id) view returns(uint256 maxSatoshi, uint256 currSatoshi, uint256 lastWithdrawTimestamp, string btcAddress, address[] keepers)
func (_GroupRegistry *GroupRegistrySession) GetGroupInfo(_id *big.Int) (struct {
	MaxSatoshi            *big.Int
	CurrSatoshi           *big.Int
	LastWithdrawTimestamp *big.Int
	BtcAddress            string
	Keepers               []common.Address
}, error) {
	return _GroupRegistry.Contract.GetGroupInfo(&_GroupRegistry.CallOpts, _id)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x99d4360f.
//
// Solidity: function getGroupInfo(uint256 _id) view returns(uint256 maxSatoshi, uint256 currSatoshi, uint256 lastWithdrawTimestamp, string btcAddress, address[] keepers)
func (_GroupRegistry *GroupRegistryCallerSession) GetGroupInfo(_id *big.Int) (struct {
	MaxSatoshi            *big.Int
	CurrSatoshi           *big.Int
	LastWithdrawTimestamp *big.Int
	BtcAddress            string
	Keepers               []common.Address
}, error) {
	return _GroupRegistry.Contract.GetGroupInfo(&_GroupRegistry.CallOpts, _id)
}

// GetGroupKeepers is a free data retrieval call binding the contract method 0x00854ace.
//
// Solidity: function getGroupKeepers(uint256 _id) view returns(address[])
func (_GroupRegistry *GroupRegistryCaller) GetGroupKeepers(opts *bind.CallOpts, _id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getGroupKeepers", _id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGroupKeepers is a free data retrieval call binding the contract method 0x00854ace.
//
// Solidity: function getGroupKeepers(uint256 _id) view returns(address[])
func (_GroupRegistry *GroupRegistrySession) GetGroupKeepers(_id *big.Int) ([]common.Address, error) {
	return _GroupRegistry.Contract.GetGroupKeepers(&_GroupRegistry.CallOpts, _id)
}

// GetGroupKeepers is a free data retrieval call binding the contract method 0x00854ace.
//
// Solidity: function getGroupKeepers(uint256 _id) view returns(address[])
func (_GroupRegistry *GroupRegistryCallerSession) GetGroupKeepers(_id *big.Int) ([]common.Address, error) {
	return _GroupRegistry.Contract.GetGroupKeepers(&_GroupRegistry.CallOpts, _id)
}

// GetKeeperGroups is a free data retrieval call binding the contract method 0x297dd2c8.
//
// Solidity: function getKeeperGroups(address _keeper, uint256 _start) view returns(uint256)
func (_GroupRegistry *GroupRegistryCaller) GetKeeperGroups(opts *bind.CallOpts, _keeper common.Address, _start *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getKeeperGroups", _keeper, _start)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeeperGroups is a free data retrieval call binding the contract method 0x297dd2c8.
//
// Solidity: function getKeeperGroups(address _keeper, uint256 _start) view returns(uint256)
func (_GroupRegistry *GroupRegistrySession) GetKeeperGroups(_keeper common.Address, _start *big.Int) (*big.Int, error) {
	return _GroupRegistry.Contract.GetKeeperGroups(&_GroupRegistry.CallOpts, _keeper, _start)
}

// GetKeeperGroups is a free data retrieval call binding the contract method 0x297dd2c8.
//
// Solidity: function getKeeperGroups(address _keeper, uint256 _start) view returns(uint256)
func (_GroupRegistry *GroupRegistryCallerSession) GetKeeperGroups(_keeper common.Address, _start *big.Int) (*big.Int, error) {
	return _GroupRegistry.Contract.GetKeeperGroups(&_GroupRegistry.CallOpts, _keeper, _start)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GroupRegistry *GroupRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GroupRegistry *GroupRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GroupRegistry.Contract.GetRoleAdmin(&_GroupRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GroupRegistry *GroupRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GroupRegistry.Contract.GetRoleAdmin(&_GroupRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GroupRegistry *GroupRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GroupRegistry *GroupRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _GroupRegistry.Contract.GetRoleMember(&_GroupRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GroupRegistry *GroupRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _GroupRegistry.Contract.GetRoleMember(&_GroupRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GroupRegistry *GroupRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GroupRegistry *GroupRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _GroupRegistry.Contract.GetRoleMemberCount(&_GroupRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GroupRegistry *GroupRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _GroupRegistry.Contract.GetRoleMemberCount(&_GroupRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GroupRegistry *GroupRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GroupRegistry *GroupRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GroupRegistry.Contract.HasRole(&_GroupRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GroupRegistry *GroupRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GroupRegistry.Contract.HasRole(&_GroupRegistry.CallOpts, role, account)
}

// IsGroupKeeper is a free data retrieval call binding the contract method 0xcb2022f6.
//
// Solidity: function isGroupKeeper(uint256 _id, address _keeper) view returns(bool)
func (_GroupRegistry *GroupRegistryCaller) IsGroupKeeper(opts *bind.CallOpts, _id *big.Int, _keeper common.Address) (bool, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "isGroupKeeper", _id, _keeper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGroupKeeper is a free data retrieval call binding the contract method 0xcb2022f6.
//
// Solidity: function isGroupKeeper(uint256 _id, address _keeper) view returns(bool)
func (_GroupRegistry *GroupRegistrySession) IsGroupKeeper(_id *big.Int, _keeper common.Address) (bool, error) {
	return _GroupRegistry.Contract.IsGroupKeeper(&_GroupRegistry.CallOpts, _id, _keeper)
}

// IsGroupKeeper is a free data retrieval call binding the contract method 0xcb2022f6.
//
// Solidity: function isGroupKeeper(uint256 _id, address _keeper) view returns(bool)
func (_GroupRegistry *GroupRegistryCallerSession) IsGroupKeeper(_id *big.Int, _keeper common.Address) (bool, error) {
	return _GroupRegistry.Contract.IsGroupKeeper(&_GroupRegistry.CallOpts, _id, _keeper)
}

// NGroups is a free data retrieval call binding the contract method 0xea9ebf9e.
//
// Solidity: function nGroups() view returns(uint256)
func (_GroupRegistry *GroupRegistryCaller) NGroups(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GroupRegistry.contract.Call(opts, &out, "nGroups")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NGroups is a free data retrieval call binding the contract method 0xea9ebf9e.
//
// Solidity: function nGroups() view returns(uint256)
func (_GroupRegistry *GroupRegistrySession) NGroups() (*big.Int, error) {
	return _GroupRegistry.Contract.NGroups(&_GroupRegistry.CallOpts)
}

// NGroups is a free data retrieval call binding the contract method 0xea9ebf9e.
//
// Solidity: function nGroups() view returns(uint256)
func (_GroupRegistry *GroupRegistryCallerSession) NGroups() (*big.Int, error) {
	return _GroupRegistry.Contract.NGroups(&_GroupRegistry.CallOpts)
}

// AddGroup is a paid mutator transaction binding the contract method 0x2504b714.
//
// Solidity: function addGroup(address[] _keepers, string _btcAddress, uint256 _maxSatoshi) returns(uint256)
func (_GroupRegistry *GroupRegistryTransactor) AddGroup(opts *bind.TransactOpts, _keepers []common.Address, _btcAddress string, _maxSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "addGroup", _keepers, _btcAddress, _maxSatoshi)
}

// AddGroup is a paid mutator transaction binding the contract method 0x2504b714.
//
// Solidity: function addGroup(address[] _keepers, string _btcAddress, uint256 _maxSatoshi) returns(uint256)
func (_GroupRegistry *GroupRegistrySession) AddGroup(_keepers []common.Address, _btcAddress string, _maxSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.AddGroup(&_GroupRegistry.TransactOpts, _keepers, _btcAddress, _maxSatoshi)
}

// AddGroup is a paid mutator transaction binding the contract method 0x2504b714.
//
// Solidity: function addGroup(address[] _keepers, string _btcAddress, uint256 _maxSatoshi) returns(uint256)
func (_GroupRegistry *GroupRegistryTransactorSession) AddGroup(_keepers []common.Address, _btcAddress string, _maxSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.AddGroup(&_GroupRegistry.TransactOpts, _keepers, _btcAddress, _maxSatoshi)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x45d88d2d.
//
// Solidity: function deleteGroup(uint256 _id) returns()
func (_GroupRegistry *GroupRegistryTransactor) DeleteGroup(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "deleteGroup", _id)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x45d88d2d.
//
// Solidity: function deleteGroup(uint256 _id) returns()
func (_GroupRegistry *GroupRegistrySession) DeleteGroup(_id *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.DeleteGroup(&_GroupRegistry.TransactOpts, _id)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x45d88d2d.
//
// Solidity: function deleteGroup(uint256 _id) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) DeleteGroup(_id *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.DeleteGroup(&_GroupRegistry.TransactOpts, _id)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x58b1ba5b.
//
// Solidity: function depositReceived(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistryTransactor) DepositReceived(opts *bind.TransactOpts, _id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "depositReceived", _id, _amountInSatoshi)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x58b1ba5b.
//
// Solidity: function depositReceived(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistrySession) DepositReceived(_id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.DepositReceived(&_GroupRegistry.TransactOpts, _id, _amountInSatoshi)
}

// DepositReceived is a paid mutator transaction binding the contract method 0x58b1ba5b.
//
// Solidity: function depositReceived(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) DepositReceived(_id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.DepositReceived(&_GroupRegistry.TransactOpts, _id, _amountInSatoshi)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.GrantRole(&_GroupRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.GrantRole(&_GroupRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.RenounceRole(&_GroupRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.RenounceRole(&_GroupRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.RevokeRole(&_GroupRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GroupRegistry.Contract.RevokeRole(&_GroupRegistry.TransactOpts, role, account)
}

// WithdrawRequested is a paid mutator transaction binding the contract method 0xe437d5b3.
//
// Solidity: function withdrawRequested(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistryTransactor) WithdrawRequested(opts *bind.TransactOpts, _id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.contract.Transact(opts, "withdrawRequested", _id, _amountInSatoshi)
}

// WithdrawRequested is a paid mutator transaction binding the contract method 0xe437d5b3.
//
// Solidity: function withdrawRequested(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistrySession) WithdrawRequested(_id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.WithdrawRequested(&_GroupRegistry.TransactOpts, _id, _amountInSatoshi)
}

// WithdrawRequested is a paid mutator transaction binding the contract method 0xe437d5b3.
//
// Solidity: function withdrawRequested(uint256 _id, uint256 _amountInSatoshi) returns()
func (_GroupRegistry *GroupRegistryTransactorSession) WithdrawRequested(_id *big.Int, _amountInSatoshi *big.Int) (*types.Transaction, error) {
	return _GroupRegistry.Contract.WithdrawRequested(&_GroupRegistry.TransactOpts, _id, _amountInSatoshi)
}

// GroupRegistryGroupAddedIterator is returned from FilterGroupAdded and is used to iterate over the raw logs and unpacked data for GroupAdded events raised by the GroupRegistry contract.
type GroupRegistryGroupAddedIterator struct {
	Event *GroupRegistryGroupAdded // Event containing the contract specifics and raw log

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
func (it *GroupRegistryGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupRegistryGroupAdded)
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
		it.Event = new(GroupRegistryGroupAdded)
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
func (it *GroupRegistryGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupRegistryGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupRegistryGroupAdded represents a GroupAdded event raised by the GroupRegistry contract.
type GroupRegistryGroupAdded struct {
	Id         *big.Int
	Keepers    []common.Address
	BtcAddress string
	MaxSatoshi *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0x2b5b732b9ed4435d138142fcaf77d4f418c7ff1d03c7860899b26365ea62c392.
//
// Solidity: event GroupAdded(uint256 indexed id, address[] keepers, string btcAddress, uint256 maxSatoshi)
func (_GroupRegistry *GroupRegistryFilterer) FilterGroupAdded(opts *bind.FilterOpts, id []*big.Int) (*GroupRegistryGroupAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GroupRegistry.contract.FilterLogs(opts, "GroupAdded", idRule)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryGroupAddedIterator{contract: _GroupRegistry.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0x2b5b732b9ed4435d138142fcaf77d4f418c7ff1d03c7860899b26365ea62c392.
//
// Solidity: event GroupAdded(uint256 indexed id, address[] keepers, string btcAddress, uint256 maxSatoshi)
func (_GroupRegistry *GroupRegistryFilterer) WatchGroupAdded(opts *bind.WatchOpts, sink chan<- *GroupRegistryGroupAdded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GroupRegistry.contract.WatchLogs(opts, "GroupAdded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupRegistryGroupAdded)
				if err := _GroupRegistry.contract.UnpackLog(event, "GroupAdded", log); err != nil {
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

// ParseGroupAdded is a log parse operation binding the contract event 0x2b5b732b9ed4435d138142fcaf77d4f418c7ff1d03c7860899b26365ea62c392.
//
// Solidity: event GroupAdded(uint256 indexed id, address[] keepers, string btcAddress, uint256 maxSatoshi)
func (_GroupRegistry *GroupRegistryFilterer) ParseGroupAdded(log types.Log) (*GroupRegistryGroupAdded, error) {
	event := new(GroupRegistryGroupAdded)
	if err := _GroupRegistry.contract.UnpackLog(event, "GroupAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupRegistryGroupDeletedIterator is returned from FilterGroupDeleted and is used to iterate over the raw logs and unpacked data for GroupDeleted events raised by the GroupRegistry contract.
type GroupRegistryGroupDeletedIterator struct {
	Event *GroupRegistryGroupDeleted // Event containing the contract specifics and raw log

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
func (it *GroupRegistryGroupDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupRegistryGroupDeleted)
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
		it.Event = new(GroupRegistryGroupDeleted)
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
func (it *GroupRegistryGroupDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupRegistryGroupDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupRegistryGroupDeleted represents a GroupDeleted event raised by the GroupRegistry contract.
type GroupRegistryGroupDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGroupDeleted is a free log retrieval operation binding the contract event 0x05f39787bb37d43683a578f1d720308a1e4cd944934f11f73c9564925aca900d.
//
// Solidity: event GroupDeleted(uint256 indexed id)
func (_GroupRegistry *GroupRegistryFilterer) FilterGroupDeleted(opts *bind.FilterOpts, id []*big.Int) (*GroupRegistryGroupDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GroupRegistry.contract.FilterLogs(opts, "GroupDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryGroupDeletedIterator{contract: _GroupRegistry.contract, event: "GroupDeleted", logs: logs, sub: sub}, nil
}

// WatchGroupDeleted is a free log subscription operation binding the contract event 0x05f39787bb37d43683a578f1d720308a1e4cd944934f11f73c9564925aca900d.
//
// Solidity: event GroupDeleted(uint256 indexed id)
func (_GroupRegistry *GroupRegistryFilterer) WatchGroupDeleted(opts *bind.WatchOpts, sink chan<- *GroupRegistryGroupDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GroupRegistry.contract.WatchLogs(opts, "GroupDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupRegistryGroupDeleted)
				if err := _GroupRegistry.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
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

// ParseGroupDeleted is a log parse operation binding the contract event 0x05f39787bb37d43683a578f1d720308a1e4cd944934f11f73c9564925aca900d.
//
// Solidity: event GroupDeleted(uint256 indexed id)
func (_GroupRegistry *GroupRegistryFilterer) ParseGroupDeleted(log types.Log) (*GroupRegistryGroupDeleted, error) {
	event := new(GroupRegistryGroupDeleted)
	if err := _GroupRegistry.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GroupRegistry contract.
type GroupRegistryRoleAdminChangedIterator struct {
	Event *GroupRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GroupRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupRegistryRoleAdminChanged)
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
		it.Event = new(GroupRegistryRoleAdminChanged)
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
func (it *GroupRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the GroupRegistry contract.
type GroupRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GroupRegistry *GroupRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GroupRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GroupRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryRoleAdminChangedIterator{contract: _GroupRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GroupRegistry *GroupRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GroupRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GroupRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupRegistryRoleAdminChanged)
				if err := _GroupRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GroupRegistry *GroupRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*GroupRegistryRoleAdminChanged, error) {
	event := new(GroupRegistryRoleAdminChanged)
	if err := _GroupRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GroupRegistry contract.
type GroupRegistryRoleGrantedIterator struct {
	Event *GroupRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *GroupRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupRegistryRoleGranted)
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
		it.Event = new(GroupRegistryRoleGranted)
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
func (it *GroupRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupRegistryRoleGranted represents a RoleGranted event raised by the GroupRegistry contract.
type GroupRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GroupRegistry *GroupRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GroupRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _GroupRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryRoleGrantedIterator{contract: _GroupRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GroupRegistry *GroupRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GroupRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GroupRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupRegistryRoleGranted)
				if err := _GroupRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GroupRegistry *GroupRegistryFilterer) ParseRoleGranted(log types.Log) (*GroupRegistryRoleGranted, error) {
	event := new(GroupRegistryRoleGranted)
	if err := _GroupRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GroupRegistry contract.
type GroupRegistryRoleRevokedIterator struct {
	Event *GroupRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GroupRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupRegistryRoleRevoked)
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
		it.Event = new(GroupRegistryRoleRevoked)
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
func (it *GroupRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupRegistryRoleRevoked represents a RoleRevoked event raised by the GroupRegistry contract.
type GroupRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GroupRegistry *GroupRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GroupRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _GroupRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GroupRegistryRoleRevokedIterator{contract: _GroupRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GroupRegistry *GroupRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GroupRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GroupRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupRegistryRoleRevoked)
				if err := _GroupRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GroupRegistry *GroupRegistryFilterer) ParseRoleRevoked(log types.Log) (*GroupRegistryRoleRevoked, error) {
	event := new(GroupRegistryRoleRevoked)
	if err := _GroupRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
