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
const KeeperRegistryABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_eBTC\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Confiscated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeeperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"KeeperDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[][]\",\"name\":\"keeperAmounts\",\"type\":\"uint256[][]\"}],\"name\":\"KeeperImported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"confiscate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confiscations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"deleteKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eBTC\",\"outputs\":[{\"internalType\":\"contractIEBTC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"getCollateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ebtcAmount\",\"type\":\"uint256\"}],\"name\":\"offsetOverissue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overissuedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"perUserCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"overissuedAmount\",\"type\":\"uint256\"}],\"name\":\"punishKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"updateTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Collaterals is a free data retrieval call binding the contract method 0xe51e119e.
//
// Solidity: function collaterals(address , address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) Collaterals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "collaterals", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collaterals is a free data retrieval call binding the contract method 0xe51e119e.
//
// Solidity: function collaterals(address , address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) Collaterals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Collaterals(&_KeeperRegistry.CallOpts, arg0, arg1)
}

// Collaterals is a free data retrieval call binding the contract method 0xe51e119e.
//
// Solidity: function collaterals(address , address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) Collaterals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Collaterals(&_KeeperRegistry.CallOpts, arg0, arg1)
}

// Confiscations is a free data retrieval call binding the contract method 0x1df398a2.
//
// Solidity: function confiscations(address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) Confiscations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "confiscations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Confiscations is a free data retrieval call binding the contract method 0x1df398a2.
//
// Solidity: function confiscations(address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) Confiscations(arg0 common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Confiscations(&_KeeperRegistry.CallOpts, arg0)
}

// Confiscations is a free data retrieval call binding the contract method 0x1df398a2.
//
// Solidity: function confiscations(address ) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) Confiscations(arg0 common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Confiscations(&_KeeperRegistry.CallOpts, arg0)
}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) EBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "eBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) EBTC() (common.Address, error) {
	return _KeeperRegistry.Contract.EBTC(&_KeeperRegistry.CallOpts)
}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) EBTC() (common.Address, error) {
	return _KeeperRegistry.Contract.EBTC(&_KeeperRegistry.CallOpts)
}

// GetCollateralValue is a free data retrieval call binding the contract method 0x97904e42.
//
// Solidity: function getCollateralValue(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) GetCollateralValue(opts *bind.CallOpts, keeper common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getCollateralValue", keeper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralValue is a free data retrieval call binding the contract method 0x97904e42.
//
// Solidity: function getCollateralValue(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) GetCollateralValue(keeper common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetCollateralValue(&_KeeperRegistry.CallOpts, keeper)
}

// GetCollateralValue is a free data retrieval call binding the contract method 0x97904e42.
//
// Solidity: function getCollateralValue(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) GetCollateralValue(keeper common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetCollateralValue(&_KeeperRegistry.CallOpts, keeper)
}

// OverissuedTotal is a free data retrieval call binding the contract method 0xda3f6c24.
//
// Solidity: function overissuedTotal() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) OverissuedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "overissuedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OverissuedTotal is a free data retrieval call binding the contract method 0xda3f6c24.
//
// Solidity: function overissuedTotal() view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) OverissuedTotal() (*big.Int, error) {
	return _KeeperRegistry.Contract.OverissuedTotal(&_KeeperRegistry.CallOpts)
}

// OverissuedTotal is a free data retrieval call binding the contract method 0xda3f6c24.
//
// Solidity: function overissuedTotal() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) OverissuedTotal() (*big.Int, error) {
	return _KeeperRegistry.Contract.OverissuedTotal(&_KeeperRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) Owner() (common.Address, error) {
	return _KeeperRegistry.Contract.Owner(&_KeeperRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) Owner() (common.Address, error) {
	return _KeeperRegistry.Contract.Owner(&_KeeperRegistry.CallOpts)
}

// PerUserCollateral is a free data retrieval call binding the contract method 0xd209793b.
//
// Solidity: function perUserCollateral(address ) view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) PerUserCollateral(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "perUserCollateral", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerUserCollateral is a free data retrieval call binding the contract method 0xd209793b.
//
// Solidity: function perUserCollateral(address ) view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) PerUserCollateral(arg0 common.Address) (common.Address, error) {
	return _KeeperRegistry.Contract.PerUserCollateral(&_KeeperRegistry.CallOpts, arg0)
}

// PerUserCollateral is a free data retrieval call binding the contract method 0xd209793b.
//
// Solidity: function perUserCollateral(address ) view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) PerUserCollateral(arg0 common.Address) (common.Address, error) {
	return _KeeperRegistry.Contract.PerUserCollateral(&_KeeperRegistry.CallOpts, arg0)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) Treasury() (common.Address, error) {
	return _KeeperRegistry.Contract.Treasury(&_KeeperRegistry.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) Treasury() (common.Address, error) {
	return _KeeperRegistry.Contract.Treasury(&_KeeperRegistry.CallOpts)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address asset) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) AddAsset(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "addAsset", asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address asset) returns()
func (_KeeperRegistry *KeeperRegistrySession) AddAsset(asset common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddAsset(&_KeeperRegistry.TransactOpts, asset)
}

// AddAsset is a paid mutator transaction binding the contract method 0x298410e5.
//
// Solidity: function addAsset(address asset) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) AddAsset(asset common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddAsset(&_KeeperRegistry.TransactOpts, asset)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x5e63edde.
//
// Solidity: function addKeeper(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) AddKeeper(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "addKeeper", asset, amount)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x5e63edde.
//
// Solidity: function addKeeper(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistrySession) AddKeeper(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddKeeper(&_KeeperRegistry.TransactOpts, asset, amount)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x5e63edde.
//
// Solidity: function addKeeper(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) AddKeeper(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddKeeper(&_KeeperRegistry.TransactOpts, asset, amount)
}

// Confiscate is a paid mutator transaction binding the contract method 0x3e7a6823.
//
// Solidity: function confiscate(address[] assets) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) Confiscate(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "confiscate", assets)
}

// Confiscate is a paid mutator transaction binding the contract method 0x3e7a6823.
//
// Solidity: function confiscate(address[] assets) returns()
func (_KeeperRegistry *KeeperRegistrySession) Confiscate(assets []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Confiscate(&_KeeperRegistry.TransactOpts, assets)
}

// Confiscate is a paid mutator transaction binding the contract method 0x3e7a6823.
//
// Solidity: function confiscate(address[] assets) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) Confiscate(assets []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Confiscate(&_KeeperRegistry.TransactOpts, assets)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) DeleteKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "deleteKeeper", keeper)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address keeper) returns()
func (_KeeperRegistry *KeeperRegistrySession) DeleteKeeper(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, keeper)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0x71101871.
//
// Solidity: function deleteKeeper(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) DeleteKeeper(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, keeper)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 ebtcAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) OffsetOverissue(opts *bind.TransactOpts, ebtcAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "offsetOverissue", ebtcAmount)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 ebtcAmount) returns()
func (_KeeperRegistry *KeeperRegistrySession) OffsetOverissue(ebtcAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.OffsetOverissue(&_KeeperRegistry.TransactOpts, ebtcAmount)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 ebtcAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) OffsetOverissue(ebtcAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.OffsetOverissue(&_KeeperRegistry.TransactOpts, ebtcAmount)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x68dcc499.
//
// Solidity: function punishKeeper(address[] keepers, uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) PunishKeeper(opts *bind.TransactOpts, keepers []common.Address, overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "punishKeeper", keepers, overissuedAmount)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x68dcc499.
//
// Solidity: function punishKeeper(address[] keepers, uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistrySession) PunishKeeper(keepers []common.Address, overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.PunishKeeper(&_KeeperRegistry.TransactOpts, keepers, overissuedAmount)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x68dcc499.
//
// Solidity: function punishKeeper(address[] keepers, uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) PunishKeeper(keepers []common.Address, overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.PunishKeeper(&_KeeperRegistry.TransactOpts, keepers, overissuedAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeeperRegistry *KeeperRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeeperRegistry *KeeperRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RenounceOwnership(&_KeeperRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeeperRegistry.Contract.RenounceOwnership(&_KeeperRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeeperRegistry *KeeperRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.TransferOwnership(&_KeeperRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.TransferOwnership(&_KeeperRegistry.TransactOpts, newOwner)
}

// UpdateTreasury is a paid mutator transaction binding the contract method 0x7f51bb1f.
//
// Solidity: function updateTreasury(address newTreasury) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) UpdateTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "updateTreasury", newTreasury)
}

// UpdateTreasury is a paid mutator transaction binding the contract method 0x7f51bb1f.
//
// Solidity: function updateTreasury(address newTreasury) returns()
func (_KeeperRegistry *KeeperRegistrySession) UpdateTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateTreasury(&_KeeperRegistry.TransactOpts, newTreasury)
}

// UpdateTreasury is a paid mutator transaction binding the contract method 0x7f51bb1f.
//
// Solidity: function updateTreasury(address newTreasury) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) UpdateTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateTreasury(&_KeeperRegistry.TransactOpts, newTreasury)
}

// KeeperRegistryAssetAddedIterator is returned from FilterAssetAdded and is used to iterate over the raw logs and unpacked data for AssetAdded events raised by the KeeperRegistry contract.
type KeeperRegistryAssetAddedIterator struct {
	Event *KeeperRegistryAssetAdded // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryAssetAdded)
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
		it.Event = new(KeeperRegistryAssetAdded)
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
func (it *KeeperRegistryAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryAssetAdded represents a AssetAdded event raised by the KeeperRegistry contract.
type KeeperRegistryAssetAdded struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAdded is a free log retrieval operation binding the contract event 0x0e3c58ebfb2e7465fbb1c32e6b4f40c3c4f5ca77e8218a386aff8617831260d7.
//
// Solidity: event AssetAdded(address indexed asset)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterAssetAdded(opts *bind.FilterOpts, asset []common.Address) (*KeeperRegistryAssetAddedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "AssetAdded", assetRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryAssetAddedIterator{contract: _KeeperRegistry.contract, event: "AssetAdded", logs: logs, sub: sub}, nil
}

// WatchAssetAdded is a free log subscription operation binding the contract event 0x0e3c58ebfb2e7465fbb1c32e6b4f40c3c4f5ca77e8218a386aff8617831260d7.
//
// Solidity: event AssetAdded(address indexed asset)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchAssetAdded(opts *bind.WatchOpts, sink chan<- *KeeperRegistryAssetAdded, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "AssetAdded", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryAssetAdded)
				if err := _KeeperRegistry.contract.UnpackLog(event, "AssetAdded", log); err != nil {
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

// ParseAssetAdded is a log parse operation binding the contract event 0x0e3c58ebfb2e7465fbb1c32e6b4f40c3c4f5ca77e8218a386aff8617831260d7.
//
// Solidity: event AssetAdded(address indexed asset)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseAssetAdded(log types.Log) (*KeeperRegistryAssetAdded, error) {
	event := new(KeeperRegistryAssetAdded)
	if err := _KeeperRegistry.contract.UnpackLog(event, "AssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryConfiscatedIterator is returned from FilterConfiscated and is used to iterate over the raw logs and unpacked data for Confiscated events raised by the KeeperRegistry contract.
type KeeperRegistryConfiscatedIterator struct {
	Event *KeeperRegistryConfiscated // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryConfiscatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryConfiscated)
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
		it.Event = new(KeeperRegistryConfiscated)
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
func (it *KeeperRegistryConfiscatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryConfiscatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryConfiscated represents a Confiscated event raised by the KeeperRegistry contract.
type KeeperRegistryConfiscated struct {
	Treasury common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfiscated is a free log retrieval operation binding the contract event 0x1d47c187095aefccfdebc8a872d270e98511c7814959be6a804c25428d7fff79.
//
// Solidity: event Confiscated(address indexed treasury, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterConfiscated(opts *bind.FilterOpts, treasury []common.Address) (*KeeperRegistryConfiscatedIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "Confiscated", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryConfiscatedIterator{contract: _KeeperRegistry.contract, event: "Confiscated", logs: logs, sub: sub}, nil
}

// WatchConfiscated is a free log subscription operation binding the contract event 0x1d47c187095aefccfdebc8a872d270e98511c7814959be6a804c25428d7fff79.
//
// Solidity: event Confiscated(address indexed treasury, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchConfiscated(opts *bind.WatchOpts, sink chan<- *KeeperRegistryConfiscated, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "Confiscated", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryConfiscated)
				if err := _KeeperRegistry.contract.UnpackLog(event, "Confiscated", log); err != nil {
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

// ParseConfiscated is a log parse operation binding the contract event 0x1d47c187095aefccfdebc8a872d270e98511c7814959be6a804c25428d7fff79.
//
// Solidity: event Confiscated(address indexed treasury, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseConfiscated(log types.Log) (*KeeperRegistryConfiscated, error) {
	event := new(KeeperRegistryConfiscated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "Confiscated", log); err != nil {
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
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperAdded is a free log retrieval operation binding the contract event 0x9cb76b51144d46005efa8364e8e73c13bed396c03ba2a905e9817b4a1011614f.
//
// Solidity: event KeeperAdded(address indexed keeper, address asset, uint256 amount)
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

// WatchKeeperAdded is a free log subscription operation binding the contract event 0x9cb76b51144d46005efa8364e8e73c13bed396c03ba2a905e9817b4a1011614f.
//
// Solidity: event KeeperAdded(address indexed keeper, address asset, uint256 amount)
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

// ParseKeeperAdded is a log parse operation binding the contract event 0x9cb76b51144d46005efa8364e8e73c13bed396c03ba2a905e9817b4a1011614f.
//
// Solidity: event KeeperAdded(address indexed keeper, address asset, uint256 amount)
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
	Keepers       []common.Address
	KeeperAmounts [][]*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterKeeperImported is a free log retrieval operation binding the contract event 0xfede295b36d1e0bf215f4c6fe12408d45367e9af7f106ed43df3e495dd774bcb.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, address[] keepers, uint256[][] keeperAmounts)
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

// WatchKeeperImported is a free log subscription operation binding the contract event 0xfede295b36d1e0bf215f4c6fe12408d45367e9af7f106ed43df3e495dd774bcb.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, address[] keepers, uint256[][] keeperAmounts)
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

// ParseKeeperImported is a log parse operation binding the contract event 0xfede295b36d1e0bf215f4c6fe12408d45367e9af7f106ed43df3e495dd774bcb.
//
// Solidity: event KeeperImported(address indexed from, address[] assets, address[] keepers, uint256[][] keeperAmounts)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperImported(log types.Log) (*KeeperRegistryKeeperImported, error) {
	event := new(KeeperRegistryKeeperImported)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeeperRegistry contract.
type KeeperRegistryOwnershipTransferredIterator struct {
	Event *KeeperRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryOwnershipTransferred)
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
		it.Event = new(KeeperRegistryOwnershipTransferred)
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
func (it *KeeperRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the KeeperRegistry contract.
type KeeperRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeeperRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryOwnershipTransferredIterator{contract: _KeeperRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeeperRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryOwnershipTransferred)
				if err := _KeeperRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*KeeperRegistryOwnershipTransferred, error) {
	event := new(KeeperRegistryOwnershipTransferred)
	if err := _KeeperRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryTreasuryTransferredIterator is returned from FilterTreasuryTransferred and is used to iterate over the raw logs and unpacked data for TreasuryTransferred events raised by the KeeperRegistry contract.
type KeeperRegistryTreasuryTransferredIterator struct {
	Event *KeeperRegistryTreasuryTransferred // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryTreasuryTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryTreasuryTransferred)
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
		it.Event = new(KeeperRegistryTreasuryTransferred)
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
func (it *KeeperRegistryTreasuryTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryTreasuryTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryTreasuryTransferred represents a TreasuryTransferred event raised by the KeeperRegistry contract.
type KeeperRegistryTreasuryTransferred struct {
	PreviousTreasury common.Address
	NewTreasury      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTreasuryTransferred is a free log retrieval operation binding the contract event 0xebebcc22237fa047dcfebf54b185ec126c9b24d45f4bd2a18e4fa02e07308c40.
//
// Solidity: event TreasuryTransferred(address indexed previousTreasury, address indexed newTreasury)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterTreasuryTransferred(opts *bind.FilterOpts, previousTreasury []common.Address, newTreasury []common.Address) (*KeeperRegistryTreasuryTransferredIterator, error) {

	var previousTreasuryRule []interface{}
	for _, previousTreasuryItem := range previousTreasury {
		previousTreasuryRule = append(previousTreasuryRule, previousTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "TreasuryTransferred", previousTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryTreasuryTransferredIterator{contract: _KeeperRegistry.contract, event: "TreasuryTransferred", logs: logs, sub: sub}, nil
}

// WatchTreasuryTransferred is a free log subscription operation binding the contract event 0xebebcc22237fa047dcfebf54b185ec126c9b24d45f4bd2a18e4fa02e07308c40.
//
// Solidity: event TreasuryTransferred(address indexed previousTreasury, address indexed newTreasury)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchTreasuryTransferred(opts *bind.WatchOpts, sink chan<- *KeeperRegistryTreasuryTransferred, previousTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var previousTreasuryRule []interface{}
	for _, previousTreasuryItem := range previousTreasury {
		previousTreasuryRule = append(previousTreasuryRule, previousTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "TreasuryTransferred", previousTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryTreasuryTransferred)
				if err := _KeeperRegistry.contract.UnpackLog(event, "TreasuryTransferred", log); err != nil {
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

// ParseTreasuryTransferred is a log parse operation binding the contract event 0xebebcc22237fa047dcfebf54b185ec126c9b24d45f4bd2a18e4fa02e07308c40.
//
// Solidity: event TreasuryTransferred(address indexed previousTreasury, address indexed newTreasury)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseTreasuryTransferred(log types.Log) (*KeeperRegistryTreasuryTransferred, error) {
	event := new(KeeperRegistryTreasuryTransferred)
	if err := _KeeperRegistry.contract.UnpackLog(event, "TreasuryTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
