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

// KeeperRegistryABI is the input ABI used to generate the binding from.
const KeeperRegistryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"meta\",\"type\":\"address\"}],\"name\":\"DependenciesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"btc\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"KeeperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"KeeperDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"keeper_amounts\",\"type\":\"uint256[]\"}],\"name\":\"KeeperImported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset_meta\",\"outputs\":[{\"internalType\":\"contractAssetMeta\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"deleteKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_keepers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_keeper_amounts\",\"type\":\"uint256[]\"}],\"name\":\"importKeepers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAssetMeta\",\"name\":\"_meta\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KeeperRegistry is an auto generated Go binding around an Ethereum contract.
type KeeperRegistry struct {
	KeeperRegistryCaller     // Read-only binding to the contract
	KeeperRegistryTransactor // Write-only binding to the contract
	KeeperRegistryFilterer   // Log filterer for contract events
}

// KeeperRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeeperRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeeperRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeeperRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeeperRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeeperRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeeperRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeeperRegistrySession struct {
	Contract     *KeeperRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeeperRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeeperRegistryCallerSession struct {
	Contract *KeeperRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KeeperRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeeperRegistryTransactorSession struct {
	Contract     *KeeperRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KeeperRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeeperRegistryRaw struct {
	Contract *KeeperRegistry // Generic contract binding to access the raw methods on
}

// KeeperRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeeperRegistryCallerRaw struct {
	Contract *KeeperRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// KeeperRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeeperRegistryTransactorRaw struct {
	Contract *KeeperRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeeperRegistry creates a new instance of KeeperRegistry, bound to a specific deployed contract.
func NewKeeperRegistry(address common.Address, backend bind.ContractBackend) (*KeeperRegistry, error) {
	contract, err := bindKeeperRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistry{KeeperRegistryCaller: KeeperRegistryCaller{contract: contract}, KeeperRegistryTransactor: KeeperRegistryTransactor{contract: contract}, KeeperRegistryFilterer: KeeperRegistryFilterer{contract: contract}}, nil
}

// NewKeeperRegistryCaller creates a new read-only instance of KeeperRegistry, bound to a specific deployed contract.
func NewKeeperRegistryCaller(address common.Address, caller bind.ContractCaller) (*KeeperRegistryCaller, error) {
	contract, err := bindKeeperRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryCaller{contract: contract}, nil
}

// NewKeeperRegistryTransactor creates a new write-only instance of KeeperRegistry, bound to a specific deployed contract.
func NewKeeperRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*KeeperRegistryTransactor, error) {
	contract, err := bindKeeperRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryTransactor{contract: contract}, nil
}

// NewKeeperRegistryFilterer creates a new log filterer instance of KeeperRegistry, bound to a specific deployed contract.
func NewKeeperRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*KeeperRegistryFilterer, error) {
	contract, err := bindKeeperRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryFilterer{contract: contract}, nil
}

// bindKeeperRegistry binds a generic wrapper to an already deployed contract.
func bindKeeperRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeeperRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeeperRegistry *KeeperRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeeperRegistry.Contract.KeeperRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeeperRegistry *KeeperRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.KeeperRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeeperRegistry *KeeperRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.KeeperRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeeperRegistry *KeeperRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeeperRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeeperRegistry *KeeperRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeeperRegistry *KeeperRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KeeperRegistry.Contract.DEFAULTADMINROLE(&_KeeperRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KeeperRegistry.Contract.DEFAULTADMINROLE(&_KeeperRegistry.CallOpts)
}

// KEEPERADMINROLE is a free data retrieval call binding the contract method 0xf419c26d.
//
// Solidity: function KEEPER_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCaller) KEEPERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "KEEPER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEEPERADMINROLE is a free data retrieval call binding the contract method 0xf419c26d.
//
// Solidity: function KEEPER_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistrySession) KEEPERADMINROLE() ([32]byte, error) {
	return _KeeperRegistry.Contract.KEEPERADMINROLE(&_KeeperRegistry.CallOpts)
}

// KEEPERADMINROLE is a free data retrieval call binding the contract method 0xf419c26d.
//
// Solidity: function KEEPER_ADMIN_ROLE() view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCallerSession) KEEPERADMINROLE() ([32]byte, error) {
	return _KeeperRegistry.Contract.KEEPERADMINROLE(&_KeeperRegistry.CallOpts)
}

// AssetMeta is a free data retrieval call binding the contract method 0xfe2b3de1.
//
// Solidity: function asset_meta() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) AssetMeta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "asset_meta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetMeta is a free data retrieval call binding the contract method 0xfe2b3de1.
//
// Solidity: function asset_meta() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) AssetMeta() (common.Address, error) {
	return _KeeperRegistry.Contract.AssetMeta(&_KeeperRegistry.CallOpts)
}

// AssetMeta is a free data retrieval call binding the contract method 0xfe2b3de1.
//
// Solidity: function asset_meta() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) AssetMeta() (common.Address, error) {
	return _KeeperRegistry.Contract.AssetMeta(&_KeeperRegistry.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCaller) Exist(opts *bind.CallOpts, _keeper common.Address) (bool, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "exist", _keeper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) Exist(_keeper common.Address) (bool, error) {
	return _KeeperRegistry.Contract.Exist(&_KeeperRegistry.CallOpts, _keeper)
}

// Exist is a free data retrieval call binding the contract method 0x4dfefc4b.
//
// Solidity: function exist(address _keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCallerSession) Exist(_keeper common.Address) (bool, error) {
	return _KeeperRegistry.Contract.Exist(&_KeeperRegistry.CallOpts, _keeper)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeeperRegistry *KeeperRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KeeperRegistry.Contract.GetRoleAdmin(&_KeeperRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KeeperRegistry *KeeperRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KeeperRegistry.Contract.GetRoleAdmin(&_KeeperRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KeeperRegistry.Contract.GetRoleMember(&_KeeperRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KeeperRegistry.Contract.GetRoleMember(&_KeeperRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetRoleMemberCount(&_KeeperRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetRoleMemberCount(&_KeeperRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KeeperRegistry.Contract.HasRole(&_KeeperRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KeeperRegistry.Contract.HasRole(&_KeeperRegistry.CallOpts, role, account)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4842c7a1.
//
// Solidity: function addKeeper(address _keeper, address[] _assets, uint256[] _amounts) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) AddKeeper(opts *bind.TransactOpts, _keeper common.Address, _assets []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "addKeeper", _keeper, _assets, _amounts)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4842c7a1.
//
// Solidity: function addKeeper(address _keeper, address[] _assets, uint256[] _amounts) returns()
func (_KeeperRegistry *KeeperRegistrySession) AddKeeper(_keeper common.Address, _assets []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddKeeper(&_KeeperRegistry.TransactOpts, _keeper, _assets, _amounts)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4842c7a1.
//
// Solidity: function addKeeper(address _keeper, address[] _assets, uint256[] _amounts) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) AddKeeper(_keeper common.Address, _assets []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddKeeper(&_KeeperRegistry.TransactOpts, _keeper, _assets, _amounts)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address _keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) DeleteKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "deleteKeeper", _keeper)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address _keeper) returns()
func (_KeeperRegistry *KeeperRegistrySession) DeleteKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, _keeper)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address _keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) DeleteKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, _keeper)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.GrantRole(&_KeeperRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.GrantRole(&_KeeperRegistry.TransactOpts, role, account)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0x4bffed06.
//
// Solidity: function importKeepers(address _from, address[] _assets, address[] _keepers, uint256[] _keeper_amounts) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) ImportKeepers(opts *bind.TransactOpts, _from common.Address, _assets []common.Address, _keepers []common.Address, _keeper_amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "importKeepers", _from, _assets, _keepers, _keeper_amounts)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0x4bffed06.
//
// Solidity: function importKeepers(address _from, address[] _assets, address[] _keepers, uint256[] _keeper_amounts) returns()
func (_KeeperRegistry *KeeperRegistrySession) ImportKeepers(_from common.Address, _assets []common.Address, _keepers []common.Address, _keeper_amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.ImportKeepers(&_KeeperRegistry.TransactOpts, _from, _assets, _keepers, _keeper_amounts)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0x4bffed06.
//
// Solidity: function importKeepers(address _from, address[] _assets, address[] _keepers, uint256[] _keeper_amounts) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) ImportKeepers(_from common.Address, _assets []common.Address, _keepers []common.Address, _keeper_amounts []*big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.ImportKeepers(&_KeeperRegistry.TransactOpts, _from, _assets, _keepers, _keeper_amounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RenounceRole(&_KeeperRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RenounceRole(&_KeeperRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RevokeRole(&_KeeperRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RevokeRole(&_KeeperRegistry.TransactOpts, role, account)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address _meta) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) SetDependencies(opts *bind.TransactOpts, _meta common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "setDependencies", _meta)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address _meta) returns()
func (_KeeperRegistry *KeeperRegistrySession) SetDependencies(_meta common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SetDependencies(&_KeeperRegistry.TransactOpts, _meta)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address _meta) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) SetDependencies(_meta common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SetDependencies(&_KeeperRegistry.TransactOpts, _meta)
}

// KeeperRegistryDependenciesSetIterator is returned from FilterDependenciesSet and is used to iterate over the raw logs and unpacked data for DependenciesSet events raised by the KeeperRegistry contract.
type KeeperRegistryDependenciesSetIterator struct {
	Event *KeeperRegistryDependenciesSet // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryDependenciesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryDependenciesSet)
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
		it.Event = new(KeeperRegistryDependenciesSet)
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
func (it *KeeperRegistryDependenciesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryDependenciesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryDependenciesSet represents a DependenciesSet event raised by the KeeperRegistry contract.
type KeeperRegistryDependenciesSet struct {
	Meta common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDependenciesSet is a free log retrieval operation binding the contract event 0x08752c0c302e35153e80e86bf52b2744b55f0d92384008e0ce53f7390cc2cd18.
//
// Solidity: event DependenciesSet(address indexed meta)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterDependenciesSet(opts *bind.FilterOpts, meta []common.Address) (*KeeperRegistryDependenciesSetIterator, error) {

	var metaRule []interface{}
	for _, metaItem := range meta {
		metaRule = append(metaRule, metaItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "DependenciesSet", metaRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryDependenciesSetIterator{contract: _KeeperRegistry.contract, event: "DependenciesSet", logs: logs, sub: sub}, nil
}

// WatchDependenciesSet is a free log subscription operation binding the contract event 0x08752c0c302e35153e80e86bf52b2744b55f0d92384008e0ce53f7390cc2cd18.
//
// Solidity: event DependenciesSet(address indexed meta)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchDependenciesSet(opts *bind.WatchOpts, sink chan<- *KeeperRegistryDependenciesSet, meta []common.Address) (event.Subscription, error) {

	var metaRule []interface{}
	for _, metaItem := range meta {
		metaRule = append(metaRule, metaItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "DependenciesSet", metaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryDependenciesSet)
				if err := _KeeperRegistry.contract.UnpackLog(event, "DependenciesSet", log); err != nil {
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

// ParseDependenciesSet is a log parse operation binding the contract event 0x08752c0c302e35153e80e86bf52b2744b55f0d92384008e0ce53f7390cc2cd18.
//
// Solidity: event DependenciesSet(address indexed meta)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseDependenciesSet(log types.Log) (*KeeperRegistryDependenciesSet, error) {
	event := new(KeeperRegistryDependenciesSet)
	if err := _KeeperRegistry.contract.UnpackLog(event, "DependenciesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryKeeperAddedIterator is returned from FilterKeeperAdded and is used to iterate over the raw logs and unpacked data for KeeperAdded events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperAddedIterator struct {
	Event *KeeperRegistryKeeperAdded // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperAdded)
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
		it.Event = new(KeeperRegistryKeeperAdded)
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
func (it *KeeperRegistryKeeperAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperAdded represents a KeeperAdded event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperAdded struct {
	Keeper common.Address
	Btc    []common.Address
	Amount []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperAdded is a free log retrieval operation binding the contract event 0x252451adfc7fb43897515e5893ed5bb7b4be14d82d690bbacf315f275e419048.
//
// Solidity: event KeeperAdded(address indexed keeper, address[] btc, uint256[] amount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperAdded(opts *bind.FilterOpts, keeper []common.Address) (*KeeperRegistryKeeperAddedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperAdded", keeperRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperAddedIterator{contract: _KeeperRegistry.contract, event: "KeeperAdded", logs: logs, sub: sub}, nil
}

// WatchKeeperAdded is a free log subscription operation binding the contract event 0x252451adfc7fb43897515e5893ed5bb7b4be14d82d690bbacf315f275e419048.
//
// Solidity: event KeeperAdded(address indexed keeper, address[] btc, uint256[] amount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperAdded(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperAdded, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperAdded", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperAdded)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperAdded", log); err != nil {
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

// ParseKeeperAdded is a log parse operation binding the contract event 0x252451adfc7fb43897515e5893ed5bb7b4be14d82d690bbacf315f275e419048.
//
// Solidity: event KeeperAdded(address indexed keeper, address[] btc, uint256[] amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperAdded(log types.Log) (*KeeperRegistryKeeperAdded, error) {
	event := new(KeeperRegistryKeeperAdded)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryKeeperDeletedIterator is returned from FilterKeeperDeleted and is used to iterate over the raw logs and unpacked data for KeeperDeleted events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperDeletedIterator struct {
	Event *KeeperRegistryKeeperDeleted // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperDeleted)
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
		it.Event = new(KeeperRegistryKeeperDeleted)
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
func (it *KeeperRegistryKeeperDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperDeleted represents a KeeperDeleted event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperDeleted struct {
	Keeper common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperDeleted is a free log retrieval operation binding the contract event 0x5d27a2180685c533af64d224ad0f39b644017f0f0a22497a0fe29f7d7450674a.
//
// Solidity: event KeeperDeleted(address indexed keeper)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperDeleted(opts *bind.FilterOpts, keeper []common.Address) (*KeeperRegistryKeeperDeletedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperDeleted", keeperRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperDeletedIterator{contract: _KeeperRegistry.contract, event: "KeeperDeleted", logs: logs, sub: sub}, nil
}

// WatchKeeperDeleted is a free log subscription operation binding the contract event 0x5d27a2180685c533af64d224ad0f39b644017f0f0a22497a0fe29f7d7450674a.
//
// Solidity: event KeeperDeleted(address indexed keeper)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperDeleted(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperDeleted, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperDeleted", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperDeleted)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperDeleted", log); err != nil {
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

// ParseKeeperDeleted is a log parse operation binding the contract event 0x5d27a2180685c533af64d224ad0f39b644017f0f0a22497a0fe29f7d7450674a.
//
// Solidity: event KeeperDeleted(address indexed keeper)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperDeleted(log types.Log) (*KeeperRegistryKeeperDeleted, error) {
	event := new(KeeperRegistryKeeperDeleted)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryKeeperImportedIterator is returned from FilterKeeperImported and is used to iterate over the raw logs and unpacked data for KeeperImported events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperImportedIterator struct {
	Event *KeeperRegistryKeeperImported // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperImported)
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
		it.Event = new(KeeperRegistryKeeperImported)
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
func (it *KeeperRegistryKeeperImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperImported represents a KeeperImported event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperImported struct {
	From          common.Address
	Assets        []common.Address
	Amounts       []*big.Int
	Keepers       []common.Address
	KeeperAmounts []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterKeeperImported is a free log retrieval operation binding the contract event 0x2a88002dff9b0894e3b41ae1f9b3248bbbc422ad67edfe308806ca948ee49964.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, uint256[] amounts, address[] keepers, uint256[] keeper_amounts)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperImported(opts *bind.FilterOpts, from []common.Address) (*KeeperRegistryKeeperImportedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperImported", fromRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperImportedIterator{contract: _KeeperRegistry.contract, event: "KeeperImported", logs: logs, sub: sub}, nil
}

// WatchKeeperImported is a free log subscription operation binding the contract event 0x2a88002dff9b0894e3b41ae1f9b3248bbbc422ad67edfe308806ca948ee49964.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, uint256[] amounts, address[] keepers, uint256[] keeper_amounts)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperImported(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperImported, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperImported", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperImported)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperImported", log); err != nil {
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

// ParseKeeperImported is a log parse operation binding the contract event 0x2a88002dff9b0894e3b41ae1f9b3248bbbc422ad67edfe308806ca948ee49964.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, uint256[] amounts, address[] keepers, uint256[] keeper_amounts)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperImported(log types.Log) (*KeeperRegistryKeeperImported, error) {
	event := new(KeeperRegistryKeeperImported)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the KeeperRegistry contract.
type KeeperRegistryRoleAdminChangedIterator struct {
	Event *KeeperRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryRoleAdminChanged)
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
		it.Event = new(KeeperRegistryRoleAdminChanged)
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
func (it *KeeperRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the KeeperRegistry contract.
type KeeperRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KeeperRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryRoleAdminChangedIterator{contract: _KeeperRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KeeperRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryRoleAdminChanged)
				if err := _KeeperRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*KeeperRegistryRoleAdminChanged, error) {
	event := new(KeeperRegistryRoleAdminChanged)
	if err := _KeeperRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KeeperRegistry contract.
type KeeperRegistryRoleGrantedIterator struct {
	Event *KeeperRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryRoleGranted)
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
		it.Event = new(KeeperRegistryRoleGranted)
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
func (it *KeeperRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryRoleGranted represents a RoleGranted event raised by the KeeperRegistry contract.
type KeeperRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeeperRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryRoleGrantedIterator{contract: _KeeperRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KeeperRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryRoleGranted)
				if err := _KeeperRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseRoleGranted(log types.Log) (*KeeperRegistryRoleGranted, error) {
	event := new(KeeperRegistryRoleGranted)
	if err := _KeeperRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KeeperRegistry contract.
type KeeperRegistryRoleRevokedIterator struct {
	Event *KeeperRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryRoleRevoked)
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
		it.Event = new(KeeperRegistryRoleRevoked)
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
func (it *KeeperRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryRoleRevoked represents a RoleRevoked event raised by the KeeperRegistry contract.
type KeeperRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeeperRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryRoleRevokedIterator{contract: _KeeperRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KeeperRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryRoleRevoked)
				if err := _KeeperRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseRoleRevoked(log types.Log) (*KeeperRegistryRoleRevoked, error) {
	event := new(KeeperRegistryRoleRevoked)
	if err := _KeeperRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
