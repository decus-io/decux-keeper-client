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

// IKeeperRegistryKeeperData is an auto generated low-level Go binding around an user-defined struct.
type IKeeperRegistryKeeperData struct {
	Amount        *big.Int
	Asset         common.Address
	RefCount      uint32
	JoinTimestamp uint32
}

// KeeperRegistryABI is the input ABI used to generate the binding from.
const KeeperRegistryABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_sats\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btcRater\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Confiscated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"bps\",\"type\":\"uint8\"}],\"name\":\"EarlyExitFeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeeperAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeeperAssetSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cAmount\",\"type\":\"uint256\"}],\"name\":\"KeeperDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeeperImported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"KeeperPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"KeeperRefCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousLiquidation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newLiquidation\",\"type\":\"address\"}],\"name\":\"LiquidationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"satsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingOverissueAmount\",\"type\":\"uint256\"}],\"name\":\"OffsetOverissued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deduction\",\"type\":\"uint256\"}],\"name\":\"OverissueAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSystem\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSystem\",\"type\":\"address\"}],\"name\":\"SystemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_KEEPER_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"overissuedAmount\",\"type\":\"uint256\"}],\"name\":\"addOverissue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcRater\",\"outputs\":[{\"internalType\":\"contractIBtcRater\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"confiscate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confiscations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"decrementRefCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cAmount\",\"type\":\"uint256\"}],\"name\":\"deleteKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyExitFeeBps\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"getCollateralWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"getKeeper\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"refCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"joinTimestamp\",\"type\":\"uint32\"}],\"internalType\":\"structIKeeperRegistry.KeeperData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"importKeepers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"incrementRefCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"isKeeperQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keeperData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"refCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"joinTimestamp\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minKeeperCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"satsAmount\",\"type\":\"uint256\"}],\"name\":\"offsetOverissue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overissuedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"punishKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sats\",\"outputs\":[{\"internalType\":\"contractSATS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_system\",\"type\":\"address\"}],\"name\":\"setSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"system\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bps\",\"type\":\"uint8\"}],\"name\":\"updateEarlyExitFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLiquidation\",\"type\":\"address\"}],\"name\":\"updateLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinKeeperCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// MINKEEPERPERIOD is a free data retrieval call binding the contract method 0xaa6e080e.
//
// Solidity: function MIN_KEEPER_PERIOD() view returns(uint32)
func (_KeeperRegistry *KeeperRegistryCaller) MINKEEPERPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "MIN_KEEPER_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINKEEPERPERIOD is a free data retrieval call binding the contract method 0xaa6e080e.
//
// Solidity: function MIN_KEEPER_PERIOD() view returns(uint32)
func (_KeeperRegistry *KeeperRegistrySession) MINKEEPERPERIOD() (uint32, error) {
	return _KeeperRegistry.Contract.MINKEEPERPERIOD(&_KeeperRegistry.CallOpts)
}

// MINKEEPERPERIOD is a free data retrieval call binding the contract method 0xaa6e080e.
//
// Solidity: function MIN_KEEPER_PERIOD() view returns(uint32)
func (_KeeperRegistry *KeeperRegistryCallerSession) MINKEEPERPERIOD() (uint32, error) {
	return _KeeperRegistry.Contract.MINKEEPERPERIOD(&_KeeperRegistry.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Allowance(&_KeeperRegistry.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.Allowance(&_KeeperRegistry.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.BalanceOf(&_KeeperRegistry.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.BalanceOf(&_KeeperRegistry.CallOpts, account)
}

// BtcRater is a free data retrieval call binding the contract method 0x122e2f16.
//
// Solidity: function btcRater() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) BtcRater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "btcRater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcRater is a free data retrieval call binding the contract method 0x122e2f16.
//
// Solidity: function btcRater() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) BtcRater() (common.Address, error) {
	return _KeeperRegistry.Contract.BtcRater(&_KeeperRegistry.CallOpts)
}

// BtcRater is a free data retrieval call binding the contract method 0x122e2f16.
//
// Solidity: function btcRater() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) BtcRater() (common.Address, error) {
	return _KeeperRegistry.Contract.BtcRater(&_KeeperRegistry.CallOpts)
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

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KeeperRegistry *KeeperRegistryCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KeeperRegistry *KeeperRegistrySession) Decimals() (uint8, error) {
	return _KeeperRegistry.Contract.Decimals(&_KeeperRegistry.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KeeperRegistry *KeeperRegistryCallerSession) Decimals() (uint8, error) {
	return _KeeperRegistry.Contract.Decimals(&_KeeperRegistry.CallOpts)
}

// EarlyExitFeeBps is a free data retrieval call binding the contract method 0x5b560a86.
//
// Solidity: function earlyExitFeeBps() view returns(uint8)
func (_KeeperRegistry *KeeperRegistryCaller) EarlyExitFeeBps(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "earlyExitFeeBps")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EarlyExitFeeBps is a free data retrieval call binding the contract method 0x5b560a86.
//
// Solidity: function earlyExitFeeBps() view returns(uint8)
func (_KeeperRegistry *KeeperRegistrySession) EarlyExitFeeBps() (uint8, error) {
	return _KeeperRegistry.Contract.EarlyExitFeeBps(&_KeeperRegistry.CallOpts)
}

// EarlyExitFeeBps is a free data retrieval call binding the contract method 0x5b560a86.
//
// Solidity: function earlyExitFeeBps() view returns(uint8)
func (_KeeperRegistry *KeeperRegistryCallerSession) EarlyExitFeeBps() (uint8, error) {
	return _KeeperRegistry.Contract.EarlyExitFeeBps(&_KeeperRegistry.CallOpts)
}

// GetCollateralWei is a free data retrieval call binding the contract method 0xbd599554.
//
// Solidity: function getCollateralWei(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) GetCollateralWei(opts *bind.CallOpts, keeper common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getCollateralWei", keeper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralWei is a free data retrieval call binding the contract method 0xbd599554.
//
// Solidity: function getCollateralWei(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) GetCollateralWei(keeper common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetCollateralWei(&_KeeperRegistry.CallOpts, keeper)
}

// GetCollateralWei is a free data retrieval call binding the contract method 0xbd599554.
//
// Solidity: function getCollateralWei(address keeper) view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) GetCollateralWei(keeper common.Address) (*big.Int, error) {
	return _KeeperRegistry.Contract.GetCollateralWei(&_KeeperRegistry.CallOpts, keeper)
}

// GetKeeper is a free data retrieval call binding the contract method 0x471ae522.
//
// Solidity: function getKeeper(address keeper) view returns((uint256,address,uint32,uint32))
func (_KeeperRegistry *KeeperRegistryCaller) GetKeeper(opts *bind.CallOpts, keeper common.Address) (IKeeperRegistryKeeperData, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "getKeeper", keeper)

	if err != nil {
		return *new(IKeeperRegistryKeeperData), err
	}

	out0 := *abi.ConvertType(out[0], new(IKeeperRegistryKeeperData)).(*IKeeperRegistryKeeperData)

	return out0, err

}

// GetKeeper is a free data retrieval call binding the contract method 0x471ae522.
//
// Solidity: function getKeeper(address keeper) view returns((uint256,address,uint32,uint32))
func (_KeeperRegistry *KeeperRegistrySession) GetKeeper(keeper common.Address) (IKeeperRegistryKeeperData, error) {
	return _KeeperRegistry.Contract.GetKeeper(&_KeeperRegistry.CallOpts, keeper)
}

// GetKeeper is a free data retrieval call binding the contract method 0x471ae522.
//
// Solidity: function getKeeper(address keeper) view returns((uint256,address,uint32,uint32))
func (_KeeperRegistry *KeeperRegistryCallerSession) GetKeeper(keeper common.Address) (IKeeperRegistryKeeperData, error) {
	return _KeeperRegistry.Contract.GetKeeper(&_KeeperRegistry.CallOpts, keeper)
}

// IsKeeperQualified is a free data retrieval call binding the contract method 0xbdaabf82.
//
// Solidity: function isKeeperQualified(address keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCaller) IsKeeperQualified(opts *bind.CallOpts, keeper common.Address) (bool, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "isKeeperQualified", keeper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeperQualified is a free data retrieval call binding the contract method 0xbdaabf82.
//
// Solidity: function isKeeperQualified(address keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) IsKeeperQualified(keeper common.Address) (bool, error) {
	return _KeeperRegistry.Contract.IsKeeperQualified(&_KeeperRegistry.CallOpts, keeper)
}

// IsKeeperQualified is a free data retrieval call binding the contract method 0xbdaabf82.
//
// Solidity: function isKeeperQualified(address keeper) view returns(bool)
func (_KeeperRegistry *KeeperRegistryCallerSession) IsKeeperQualified(keeper common.Address) (bool, error) {
	return _KeeperRegistry.Contract.IsKeeperQualified(&_KeeperRegistry.CallOpts, keeper)
}

// KeeperData is a free data retrieval call binding the contract method 0x6f0817b9.
//
// Solidity: function keeperData(address ) view returns(uint256 amount, address asset, uint32 refCount, uint32 joinTimestamp)
func (_KeeperRegistry *KeeperRegistryCaller) KeeperData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount        *big.Int
	Asset         common.Address
	RefCount      uint32
	JoinTimestamp uint32
}, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "keeperData", arg0)

	outstruct := new(struct {
		Amount        *big.Int
		Asset         common.Address
		RefCount      uint32
		JoinTimestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Asset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RefCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.JoinTimestamp = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// KeeperData is a free data retrieval call binding the contract method 0x6f0817b9.
//
// Solidity: function keeperData(address ) view returns(uint256 amount, address asset, uint32 refCount, uint32 joinTimestamp)
func (_KeeperRegistry *KeeperRegistrySession) KeeperData(arg0 common.Address) (struct {
	Amount        *big.Int
	Asset         common.Address
	RefCount      uint32
	JoinTimestamp uint32
}, error) {
	return _KeeperRegistry.Contract.KeeperData(&_KeeperRegistry.CallOpts, arg0)
}

// KeeperData is a free data retrieval call binding the contract method 0x6f0817b9.
//
// Solidity: function keeperData(address ) view returns(uint256 amount, address asset, uint32 refCount, uint32 joinTimestamp)
func (_KeeperRegistry *KeeperRegistryCallerSession) KeeperData(arg0 common.Address) (struct {
	Amount        *big.Int
	Asset         common.Address
	RefCount      uint32
	JoinTimestamp uint32
}, error) {
	return _KeeperRegistry.Contract.KeeperData(&_KeeperRegistry.CallOpts, arg0)
}

// Liquidation is a free data retrieval call binding the contract method 0xf2dfbf66.
//
// Solidity: function liquidation() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) Liquidation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "liquidation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Liquidation is a free data retrieval call binding the contract method 0xf2dfbf66.
//
// Solidity: function liquidation() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) Liquidation() (common.Address, error) {
	return _KeeperRegistry.Contract.Liquidation(&_KeeperRegistry.CallOpts)
}

// Liquidation is a free data retrieval call binding the contract method 0xf2dfbf66.
//
// Solidity: function liquidation() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) Liquidation() (common.Address, error) {
	return _KeeperRegistry.Contract.Liquidation(&_KeeperRegistry.CallOpts)
}

// MinKeeperCollateral is a free data retrieval call binding the contract method 0xf4f8c025.
//
// Solidity: function minKeeperCollateral() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) MinKeeperCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "minKeeperCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinKeeperCollateral is a free data retrieval call binding the contract method 0xf4f8c025.
//
// Solidity: function minKeeperCollateral() view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) MinKeeperCollateral() (*big.Int, error) {
	return _KeeperRegistry.Contract.MinKeeperCollateral(&_KeeperRegistry.CallOpts)
}

// MinKeeperCollateral is a free data retrieval call binding the contract method 0xf4f8c025.
//
// Solidity: function minKeeperCollateral() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) MinKeeperCollateral() (*big.Int, error) {
	return _KeeperRegistry.Contract.MinKeeperCollateral(&_KeeperRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KeeperRegistry *KeeperRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KeeperRegistry *KeeperRegistrySession) Name() (string, error) {
	return _KeeperRegistry.Contract.Name(&_KeeperRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KeeperRegistry *KeeperRegistryCallerSession) Name() (string, error) {
	return _KeeperRegistry.Contract.Name(&_KeeperRegistry.CallOpts)
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

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) Sats(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "sats")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) Sats() (common.Address, error) {
	return _KeeperRegistry.Contract.Sats(&_KeeperRegistry.CallOpts)
}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) Sats() (common.Address, error) {
	return _KeeperRegistry.Contract.Sats(&_KeeperRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KeeperRegistry *KeeperRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KeeperRegistry *KeeperRegistrySession) Symbol() (string, error) {
	return _KeeperRegistry.Contract.Symbol(&_KeeperRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KeeperRegistry *KeeperRegistryCallerSession) Symbol() (string, error) {
	return _KeeperRegistry.Contract.Symbol(&_KeeperRegistry.CallOpts)
}

// System is a free data retrieval call binding the contract method 0x95bf75fd.
//
// Solidity: function system() view returns(address)
func (_KeeperRegistry *KeeperRegistryCaller) System(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "system")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// System is a free data retrieval call binding the contract method 0x95bf75fd.
//
// Solidity: function system() view returns(address)
func (_KeeperRegistry *KeeperRegistrySession) System() (common.Address, error) {
	return _KeeperRegistry.Contract.System(&_KeeperRegistry.CallOpts)
}

// System is a free data retrieval call binding the contract method 0x95bf75fd.
//
// Solidity: function system() view returns(address)
func (_KeeperRegistry *KeeperRegistryCallerSession) System() (common.Address, error) {
	return _KeeperRegistry.Contract.System(&_KeeperRegistry.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KeeperRegistry.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KeeperRegistry *KeeperRegistrySession) TotalSupply() (*big.Int, error) {
	return _KeeperRegistry.Contract.TotalSupply(&_KeeperRegistry.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KeeperRegistry *KeeperRegistryCallerSession) TotalSupply() (*big.Int, error) {
	return _KeeperRegistry.Contract.TotalSupply(&_KeeperRegistry.CallOpts)
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

// AddOverissue is a paid mutator transaction binding the contract method 0x0c3055f5.
//
// Solidity: function addOverissue(uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) AddOverissue(opts *bind.TransactOpts, overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "addOverissue", overissuedAmount)
}

// AddOverissue is a paid mutator transaction binding the contract method 0x0c3055f5.
//
// Solidity: function addOverissue(uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistrySession) AddOverissue(overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddOverissue(&_KeeperRegistry.TransactOpts, overissuedAmount)
}

// AddOverissue is a paid mutator transaction binding the contract method 0x0c3055f5.
//
// Solidity: function addOverissue(uint256 overissuedAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) AddOverissue(overissuedAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.AddOverissue(&_KeeperRegistry.TransactOpts, overissuedAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Approve(&_KeeperRegistry.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Approve(&_KeeperRegistry.TransactOpts, spender, amount)
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

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DecreaseAllowance(&_KeeperRegistry.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DecreaseAllowance(&_KeeperRegistry.TransactOpts, spender, subtractedValue)
}

// DecrementRefCount is a paid mutator transaction binding the contract method 0x1323bad4.
//
// Solidity: function decrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) DecrementRefCount(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "decrementRefCount", keeper)
}

// DecrementRefCount is a paid mutator transaction binding the contract method 0x1323bad4.
//
// Solidity: function decrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistrySession) DecrementRefCount(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DecrementRefCount(&_KeeperRegistry.TransactOpts, keeper)
}

// DecrementRefCount is a paid mutator transaction binding the contract method 0x1323bad4.
//
// Solidity: function decrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) DecrementRefCount(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DecrementRefCount(&_KeeperRegistry.TransactOpts, keeper)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0xcc2be1e4.
//
// Solidity: function deleteKeeper(uint256 cAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) DeleteKeeper(opts *bind.TransactOpts, cAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "deleteKeeper", cAmount)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0xcc2be1e4.
//
// Solidity: function deleteKeeper(uint256 cAmount) returns()
func (_KeeperRegistry *KeeperRegistrySession) DeleteKeeper(cAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, cAmount)
}

// DeleteKeeper is a paid mutator transaction binding the contract method 0xcc2be1e4.
//
// Solidity: function deleteKeeper(uint256 cAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) DeleteKeeper(cAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.DeleteKeeper(&_KeeperRegistry.TransactOpts, cAmount)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0xd67f23a9.
//
// Solidity: function importKeepers(uint256 amount, address asset, address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) ImportKeepers(opts *bind.TransactOpts, amount *big.Int, asset common.Address, keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "importKeepers", amount, asset, keepers)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0xd67f23a9.
//
// Solidity: function importKeepers(uint256 amount, address asset, address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistrySession) ImportKeepers(amount *big.Int, asset common.Address, keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.ImportKeepers(&_KeeperRegistry.TransactOpts, amount, asset, keepers)
}

// ImportKeepers is a paid mutator transaction binding the contract method 0xd67f23a9.
//
// Solidity: function importKeepers(uint256 amount, address asset, address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) ImportKeepers(amount *big.Int, asset common.Address, keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.ImportKeepers(&_KeeperRegistry.TransactOpts, amount, asset, keepers)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.IncreaseAllowance(&_KeeperRegistry.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.IncreaseAllowance(&_KeeperRegistry.TransactOpts, spender, addedValue)
}

// IncrementRefCount is a paid mutator transaction binding the contract method 0x00103431.
//
// Solidity: function incrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) IncrementRefCount(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "incrementRefCount", keeper)
}

// IncrementRefCount is a paid mutator transaction binding the contract method 0x00103431.
//
// Solidity: function incrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistrySession) IncrementRefCount(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.IncrementRefCount(&_KeeperRegistry.TransactOpts, keeper)
}

// IncrementRefCount is a paid mutator transaction binding the contract method 0x00103431.
//
// Solidity: function incrementRefCount(address keeper) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) IncrementRefCount(keeper common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.IncrementRefCount(&_KeeperRegistry.TransactOpts, keeper)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 satsAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) OffsetOverissue(opts *bind.TransactOpts, satsAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "offsetOverissue", satsAmount)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 satsAmount) returns()
func (_KeeperRegistry *KeeperRegistrySession) OffsetOverissue(satsAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.OffsetOverissue(&_KeeperRegistry.TransactOpts, satsAmount)
}

// OffsetOverissue is a paid mutator transaction binding the contract method 0xee529bff.
//
// Solidity: function offsetOverissue(uint256 satsAmount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) OffsetOverissue(satsAmount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.OffsetOverissue(&_KeeperRegistry.TransactOpts, satsAmount)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x834b33b4.
//
// Solidity: function punishKeeper(address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) PunishKeeper(opts *bind.TransactOpts, keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "punishKeeper", keepers)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x834b33b4.
//
// Solidity: function punishKeeper(address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistrySession) PunishKeeper(keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.PunishKeeper(&_KeeperRegistry.TransactOpts, keepers)
}

// PunishKeeper is a paid mutator transaction binding the contract method 0x834b33b4.
//
// Solidity: function punishKeeper(address[] keepers) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) PunishKeeper(keepers []common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.PunishKeeper(&_KeeperRegistry.TransactOpts, keepers)
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

// SetSystem is a paid mutator transaction binding the contract method 0x55837757.
//
// Solidity: function setSystem(address _system) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) SetSystem(opts *bind.TransactOpts, _system common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "setSystem", _system)
}

// SetSystem is a paid mutator transaction binding the contract method 0x55837757.
//
// Solidity: function setSystem(address _system) returns()
func (_KeeperRegistry *KeeperRegistrySession) SetSystem(_system common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SetSystem(&_KeeperRegistry.TransactOpts, _system)
}

// SetSystem is a paid mutator transaction binding the contract method 0x55837757.
//
// Solidity: function setSystem(address _system) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) SetSystem(_system common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SetSystem(&_KeeperRegistry.TransactOpts, _system)
}

// SwapAsset is a paid mutator transaction binding the contract method 0xdde13c46.
//
// Solidity: function swapAsset(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) SwapAsset(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "swapAsset", asset, amount)
}

// SwapAsset is a paid mutator transaction binding the contract method 0xdde13c46.
//
// Solidity: function swapAsset(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistrySession) SwapAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SwapAsset(&_KeeperRegistry.TransactOpts, asset, amount)
}

// SwapAsset is a paid mutator transaction binding the contract method 0xdde13c46.
//
// Solidity: function swapAsset(address asset, uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) SwapAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.SwapAsset(&_KeeperRegistry.TransactOpts, asset, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Transfer(&_KeeperRegistry.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.Transfer(&_KeeperRegistry.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistrySession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.TransferFrom(&_KeeperRegistry.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KeeperRegistry *KeeperRegistryTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.TransferFrom(&_KeeperRegistry.TransactOpts, sender, recipient, amount)
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

// UpdateEarlyExitFeeBps is a paid mutator transaction binding the contract method 0xfee48a50.
//
// Solidity: function updateEarlyExitFeeBps(uint8 bps) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) UpdateEarlyExitFeeBps(opts *bind.TransactOpts, bps uint8) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "updateEarlyExitFeeBps", bps)
}

// UpdateEarlyExitFeeBps is a paid mutator transaction binding the contract method 0xfee48a50.
//
// Solidity: function updateEarlyExitFeeBps(uint8 bps) returns()
func (_KeeperRegistry *KeeperRegistrySession) UpdateEarlyExitFeeBps(bps uint8) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateEarlyExitFeeBps(&_KeeperRegistry.TransactOpts, bps)
}

// UpdateEarlyExitFeeBps is a paid mutator transaction binding the contract method 0xfee48a50.
//
// Solidity: function updateEarlyExitFeeBps(uint8 bps) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) UpdateEarlyExitFeeBps(bps uint8) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateEarlyExitFeeBps(&_KeeperRegistry.TransactOpts, bps)
}

// UpdateLiquidation is a paid mutator transaction binding the contract method 0xcb303fbd.
//
// Solidity: function updateLiquidation(address newLiquidation) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) UpdateLiquidation(opts *bind.TransactOpts, newLiquidation common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "updateLiquidation", newLiquidation)
}

// UpdateLiquidation is a paid mutator transaction binding the contract method 0xcb303fbd.
//
// Solidity: function updateLiquidation(address newLiquidation) returns()
func (_KeeperRegistry *KeeperRegistrySession) UpdateLiquidation(newLiquidation common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateLiquidation(&_KeeperRegistry.TransactOpts, newLiquidation)
}

// UpdateLiquidation is a paid mutator transaction binding the contract method 0xcb303fbd.
//
// Solidity: function updateLiquidation(address newLiquidation) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) UpdateLiquidation(newLiquidation common.Address) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateLiquidation(&_KeeperRegistry.TransactOpts, newLiquidation)
}

// UpdateMinKeeperCollateral is a paid mutator transaction binding the contract method 0xeef36af0.
//
// Solidity: function updateMinKeeperCollateral(uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactor) UpdateMinKeeperCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.contract.Transact(opts, "updateMinKeeperCollateral", amount)
}

// UpdateMinKeeperCollateral is a paid mutator transaction binding the contract method 0xeef36af0.
//
// Solidity: function updateMinKeeperCollateral(uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistrySession) UpdateMinKeeperCollateral(amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateMinKeeperCollateral(&_KeeperRegistry.TransactOpts, amount)
}

// UpdateMinKeeperCollateral is a paid mutator transaction binding the contract method 0xeef36af0.
//
// Solidity: function updateMinKeeperCollateral(uint256 amount) returns()
func (_KeeperRegistry *KeeperRegistryTransactorSession) UpdateMinKeeperCollateral(amount *big.Int) (*types.Transaction, error) {
	return _KeeperRegistry.Contract.UpdateMinKeeperCollateral(&_KeeperRegistry.TransactOpts, amount)
}

// KeeperRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KeeperRegistry contract.
type KeeperRegistryApprovalIterator struct {
	Event *KeeperRegistryApproval // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryApproval)
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
		it.Event = new(KeeperRegistryApproval)
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
func (it *KeeperRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryApproval represents a Approval event raised by the KeeperRegistry contract.
type KeeperRegistryApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KeeperRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryApprovalIterator{contract: _KeeperRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KeeperRegistryApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryApproval)
				if err := _KeeperRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseApproval(log types.Log) (*KeeperRegistryApproval, error) {
	event := new(KeeperRegistryApproval)
	if err := _KeeperRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Liquidation common.Address
	Asset       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfiscated is a free log retrieval operation binding the contract event 0x1d47c187095aefccfdebc8a872d270e98511c7814959be6a804c25428d7fff79.
//
// Solidity: event Confiscated(address indexed liquidation, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterConfiscated(opts *bind.FilterOpts, liquidation []common.Address) (*KeeperRegistryConfiscatedIterator, error) {

	var liquidationRule []interface{}
	for _, liquidationItem := range liquidation {
		liquidationRule = append(liquidationRule, liquidationItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "Confiscated", liquidationRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryConfiscatedIterator{contract: _KeeperRegistry.contract, event: "Confiscated", logs: logs, sub: sub}, nil
}

// WatchConfiscated is a free log subscription operation binding the contract event 0x1d47c187095aefccfdebc8a872d270e98511c7814959be6a804c25428d7fff79.
//
// Solidity: event Confiscated(address indexed liquidation, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchConfiscated(opts *bind.WatchOpts, sink chan<- *KeeperRegistryConfiscated, liquidation []common.Address) (event.Subscription, error) {

	var liquidationRule []interface{}
	for _, liquidationItem := range liquidation {
		liquidationRule = append(liquidationRule, liquidationItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "Confiscated", liquidationRule)
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
// Solidity: event Confiscated(address indexed liquidation, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseConfiscated(log types.Log) (*KeeperRegistryConfiscated, error) {
	event := new(KeeperRegistryConfiscated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "Confiscated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryEarlyExitFeeBpsUpdatedIterator is returned from FilterEarlyExitFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for EarlyExitFeeBpsUpdated events raised by the KeeperRegistry contract.
type KeeperRegistryEarlyExitFeeBpsUpdatedIterator struct {
	Event *KeeperRegistryEarlyExitFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryEarlyExitFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryEarlyExitFeeBpsUpdated)
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
		it.Event = new(KeeperRegistryEarlyExitFeeBpsUpdated)
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
func (it *KeeperRegistryEarlyExitFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryEarlyExitFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryEarlyExitFeeBpsUpdated represents a EarlyExitFeeBpsUpdated event raised by the KeeperRegistry contract.
type KeeperRegistryEarlyExitFeeBpsUpdated struct {
	Bps uint8
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEarlyExitFeeBpsUpdated is a free log retrieval operation binding the contract event 0x5859a024dcec66c2037180ba1326d01a41e33d4ef7b94b17cdc3c017b0498b9a.
//
// Solidity: event EarlyExitFeeBpsUpdated(uint8 bps)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterEarlyExitFeeBpsUpdated(opts *bind.FilterOpts) (*KeeperRegistryEarlyExitFeeBpsUpdatedIterator, error) {

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "EarlyExitFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryEarlyExitFeeBpsUpdatedIterator{contract: _KeeperRegistry.contract, event: "EarlyExitFeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchEarlyExitFeeBpsUpdated is a free log subscription operation binding the contract event 0x5859a024dcec66c2037180ba1326d01a41e33d4ef7b94b17cdc3c017b0498b9a.
//
// Solidity: event EarlyExitFeeBpsUpdated(uint8 bps)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchEarlyExitFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *KeeperRegistryEarlyExitFeeBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "EarlyExitFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryEarlyExitFeeBpsUpdated)
				if err := _KeeperRegistry.contract.UnpackLog(event, "EarlyExitFeeBpsUpdated", log); err != nil {
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

// ParseEarlyExitFeeBpsUpdated is a log parse operation binding the contract event 0x5859a024dcec66c2037180ba1326d01a41e33d4ef7b94b17cdc3c017b0498b9a.
//
// Solidity: event EarlyExitFeeBpsUpdated(uint8 bps)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseEarlyExitFeeBpsUpdated(log types.Log) (*KeeperRegistryEarlyExitFeeBpsUpdated, error) {
	event := new(KeeperRegistryEarlyExitFeeBpsUpdated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "EarlyExitFeeBpsUpdated", log); err != nil {
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

// KeeperRegistryKeeperAssetSwappedIterator is returned from FilterKeeperAssetSwapped and is used to iterate over the raw logs and unpacked data for KeeperAssetSwapped events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperAssetSwappedIterator struct {
	Event *KeeperRegistryKeeperAssetSwapped // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperAssetSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperAssetSwapped)
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
		it.Event = new(KeeperRegistryKeeperAssetSwapped)
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
func (it *KeeperRegistryKeeperAssetSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperAssetSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperAssetSwapped represents a KeeperAssetSwapped event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperAssetSwapped struct {
	Keeper common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperAssetSwapped is a free log retrieval operation binding the contract event 0x64a4f8323ef4b922c8fc418aed8017443f62c3aa093ec29f7a62e8cbf06d25cf.
//
// Solidity: event KeeperAssetSwapped(address indexed keeper, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperAssetSwapped(opts *bind.FilterOpts, keeper []common.Address) (*KeeperRegistryKeeperAssetSwappedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperAssetSwapped", keeperRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperAssetSwappedIterator{contract: _KeeperRegistry.contract, event: "KeeperAssetSwapped", logs: logs, sub: sub}, nil
}

// WatchKeeperAssetSwapped is a free log subscription operation binding the contract event 0x64a4f8323ef4b922c8fc418aed8017443f62c3aa093ec29f7a62e8cbf06d25cf.
//
// Solidity: event KeeperAssetSwapped(address indexed keeper, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperAssetSwapped(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperAssetSwapped, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperAssetSwapped", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperAssetSwapped)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperAssetSwapped", log); err != nil {
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

// ParseKeeperAssetSwapped is a log parse operation binding the contract event 0x64a4f8323ef4b922c8fc418aed8017443f62c3aa093ec29f7a62e8cbf06d25cf.
//
// Solidity: event KeeperAssetSwapped(address indexed keeper, address asset, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperAssetSwapped(log types.Log) (*KeeperRegistryKeeperAssetSwapped, error) {
	event := new(KeeperRegistryKeeperAssetSwapped)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperAssetSwapped", log); err != nil {
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
	Keeper  common.Address
	Asset   common.Address
	Amount  *big.Int
	CAmount *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeeperDeleted is a free log retrieval operation binding the contract event 0x06b7631533b900b7ead1e432d65a3527f75c28916090fc83ea68796a8150cea0.
//
// Solidity: event KeeperDeleted(address indexed keeper, address asset, uint256 amount, uint256 cAmount)
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

// WatchKeeperDeleted is a free log subscription operation binding the contract event 0x06b7631533b900b7ead1e432d65a3527f75c28916090fc83ea68796a8150cea0.
//
// Solidity: event KeeperDeleted(address indexed keeper, address asset, uint256 amount, uint256 cAmount)
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

// ParseKeeperDeleted is a log parse operation binding the contract event 0x06b7631533b900b7ead1e432d65a3527f75c28916090fc83ea68796a8150cea0.
//
// Solidity: event KeeperDeleted(address indexed keeper, address asset, uint256 amount, uint256 cAmount)
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
	From    common.Address
	Asset   common.Address
	Keepers []common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeeperImported is a free log retrieval operation binding the contract event 0x2f63a3a45a4cfd9e30f44e24284095f439a7e04af40b0b4ca19b51e3adc0de86.
//
// Solidity: event KeeperImported(address indexed from, address asset, address[] keepers, uint256 amount)
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

// WatchKeeperImported is a free log subscription operation binding the contract event 0x2f63a3a45a4cfd9e30f44e24284095f439a7e04af40b0b4ca19b51e3adc0de86.
//
// Solidity: event KeeperImported(address indexed from, address asset, address[] keepers, uint256 amount)
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

// ParseKeeperImported is a log parse operation binding the contract event 0x2f63a3a45a4cfd9e30f44e24284095f439a7e04af40b0b4ca19b51e3adc0de86.
//
// Solidity: event KeeperImported(address indexed from, address asset, address[] keepers, uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperImported(log types.Log) (*KeeperRegistryKeeperImported, error) {
	event := new(KeeperRegistryKeeperImported)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryKeeperPunishedIterator is returned from FilterKeeperPunished and is used to iterate over the raw logs and unpacked data for KeeperPunished events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperPunishedIterator struct {
	Event *KeeperRegistryKeeperPunished // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperPunished)
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
		it.Event = new(KeeperRegistryKeeperPunished)
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
func (it *KeeperRegistryKeeperPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperPunished represents a KeeperPunished event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperPunished struct {
	Keeper     common.Address
	Asset      common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKeeperPunished is a free log retrieval operation binding the contract event 0x198a160a2fb6fe31d6b12fc94563783570f80c88e2de58a05d849d5b61dbe9c4.
//
// Solidity: event KeeperPunished(address indexed keeper, address asset, uint256 collateral)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperPunished(opts *bind.FilterOpts, keeper []common.Address) (*KeeperRegistryKeeperPunishedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperPunished", keeperRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperPunishedIterator{contract: _KeeperRegistry.contract, event: "KeeperPunished", logs: logs, sub: sub}, nil
}

// WatchKeeperPunished is a free log subscription operation binding the contract event 0x198a160a2fb6fe31d6b12fc94563783570f80c88e2de58a05d849d5b61dbe9c4.
//
// Solidity: event KeeperPunished(address indexed keeper, address asset, uint256 collateral)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperPunished(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperPunished, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperPunished", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperPunished)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperPunished", log); err != nil {
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

// ParseKeeperPunished is a log parse operation binding the contract event 0x198a160a2fb6fe31d6b12fc94563783570f80c88e2de58a05d849d5b61dbe9c4.
//
// Solidity: event KeeperPunished(address indexed keeper, address asset, uint256 collateral)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperPunished(log types.Log) (*KeeperRegistryKeeperPunished, error) {
	event := new(KeeperRegistryKeeperPunished)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryKeeperRefCountIterator is returned from FilterKeeperRefCount and is used to iterate over the raw logs and unpacked data for KeeperRefCount events raised by the KeeperRegistry contract.
type KeeperRegistryKeeperRefCountIterator struct {
	Event *KeeperRegistryKeeperRefCount // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryKeeperRefCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryKeeperRefCount)
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
		it.Event = new(KeeperRegistryKeeperRefCount)
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
func (it *KeeperRegistryKeeperRefCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryKeeperRefCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryKeeperRefCount represents a KeeperRefCount event raised by the KeeperRegistry contract.
type KeeperRegistryKeeperRefCount struct {
	Keeper common.Address
	Count  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperRefCount is a free log retrieval operation binding the contract event 0x27937aff960b9928a1a9719d1837f973fd249911354d1e79ecd56c4746f02eeb.
//
// Solidity: event KeeperRefCount(address indexed keeper, uint256 count)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterKeeperRefCount(opts *bind.FilterOpts, keeper []common.Address) (*KeeperRegistryKeeperRefCountIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "KeeperRefCount", keeperRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryKeeperRefCountIterator{contract: _KeeperRegistry.contract, event: "KeeperRefCount", logs: logs, sub: sub}, nil
}

// WatchKeeperRefCount is a free log subscription operation binding the contract event 0x27937aff960b9928a1a9719d1837f973fd249911354d1e79ecd56c4746f02eeb.
//
// Solidity: event KeeperRefCount(address indexed keeper, uint256 count)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchKeeperRefCount(opts *bind.WatchOpts, sink chan<- *KeeperRegistryKeeperRefCount, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "KeeperRefCount", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryKeeperRefCount)
				if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperRefCount", log); err != nil {
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

// ParseKeeperRefCount is a log parse operation binding the contract event 0x27937aff960b9928a1a9719d1837f973fd249911354d1e79ecd56c4746f02eeb.
//
// Solidity: event KeeperRefCount(address indexed keeper, uint256 count)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseKeeperRefCount(log types.Log) (*KeeperRegistryKeeperRefCount, error) {
	event := new(KeeperRegistryKeeperRefCount)
	if err := _KeeperRegistry.contract.UnpackLog(event, "KeeperRefCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryLiquidationUpdatedIterator is returned from FilterLiquidationUpdated and is used to iterate over the raw logs and unpacked data for LiquidationUpdated events raised by the KeeperRegistry contract.
type KeeperRegistryLiquidationUpdatedIterator struct {
	Event *KeeperRegistryLiquidationUpdated // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryLiquidationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryLiquidationUpdated)
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
		it.Event = new(KeeperRegistryLiquidationUpdated)
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
func (it *KeeperRegistryLiquidationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryLiquidationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryLiquidationUpdated represents a LiquidationUpdated event raised by the KeeperRegistry contract.
type KeeperRegistryLiquidationUpdated struct {
	PreviousLiquidation common.Address
	NewLiquidation      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLiquidationUpdated is a free log retrieval operation binding the contract event 0x960ce90a19dfa3c809fb677d4929ae04bfacc9138539566f398b02ac16b6b31e.
//
// Solidity: event LiquidationUpdated(address indexed previousLiquidation, address indexed newLiquidation)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterLiquidationUpdated(opts *bind.FilterOpts, previousLiquidation []common.Address, newLiquidation []common.Address) (*KeeperRegistryLiquidationUpdatedIterator, error) {

	var previousLiquidationRule []interface{}
	for _, previousLiquidationItem := range previousLiquidation {
		previousLiquidationRule = append(previousLiquidationRule, previousLiquidationItem)
	}
	var newLiquidationRule []interface{}
	for _, newLiquidationItem := range newLiquidation {
		newLiquidationRule = append(newLiquidationRule, newLiquidationItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "LiquidationUpdated", previousLiquidationRule, newLiquidationRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryLiquidationUpdatedIterator{contract: _KeeperRegistry.contract, event: "LiquidationUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidationUpdated is a free log subscription operation binding the contract event 0x960ce90a19dfa3c809fb677d4929ae04bfacc9138539566f398b02ac16b6b31e.
//
// Solidity: event LiquidationUpdated(address indexed previousLiquidation, address indexed newLiquidation)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchLiquidationUpdated(opts *bind.WatchOpts, sink chan<- *KeeperRegistryLiquidationUpdated, previousLiquidation []common.Address, newLiquidation []common.Address) (event.Subscription, error) {

	var previousLiquidationRule []interface{}
	for _, previousLiquidationItem := range previousLiquidation {
		previousLiquidationRule = append(previousLiquidationRule, previousLiquidationItem)
	}
	var newLiquidationRule []interface{}
	for _, newLiquidationItem := range newLiquidation {
		newLiquidationRule = append(newLiquidationRule, newLiquidationItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "LiquidationUpdated", previousLiquidationRule, newLiquidationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryLiquidationUpdated)
				if err := _KeeperRegistry.contract.UnpackLog(event, "LiquidationUpdated", log); err != nil {
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

// ParseLiquidationUpdated is a log parse operation binding the contract event 0x960ce90a19dfa3c809fb677d4929ae04bfacc9138539566f398b02ac16b6b31e.
//
// Solidity: event LiquidationUpdated(address indexed previousLiquidation, address indexed newLiquidation)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseLiquidationUpdated(log types.Log) (*KeeperRegistryLiquidationUpdated, error) {
	event := new(KeeperRegistryLiquidationUpdated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "LiquidationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryMinCollateralUpdatedIterator is returned from FilterMinCollateralUpdated and is used to iterate over the raw logs and unpacked data for MinCollateralUpdated events raised by the KeeperRegistry contract.
type KeeperRegistryMinCollateralUpdatedIterator struct {
	Event *KeeperRegistryMinCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryMinCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryMinCollateralUpdated)
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
		it.Event = new(KeeperRegistryMinCollateralUpdated)
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
func (it *KeeperRegistryMinCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryMinCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryMinCollateralUpdated represents a MinCollateralUpdated event raised by the KeeperRegistry contract.
type KeeperRegistryMinCollateralUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinCollateralUpdated is a free log retrieval operation binding the contract event 0xd19fe8ad9152af12b174a60210fb798db0767d63973ebb97298dc44d67a5c82d.
//
// Solidity: event MinCollateralUpdated(uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterMinCollateralUpdated(opts *bind.FilterOpts) (*KeeperRegistryMinCollateralUpdatedIterator, error) {

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "MinCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryMinCollateralUpdatedIterator{contract: _KeeperRegistry.contract, event: "MinCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchMinCollateralUpdated is a free log subscription operation binding the contract event 0xd19fe8ad9152af12b174a60210fb798db0767d63973ebb97298dc44d67a5c82d.
//
// Solidity: event MinCollateralUpdated(uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchMinCollateralUpdated(opts *bind.WatchOpts, sink chan<- *KeeperRegistryMinCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "MinCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryMinCollateralUpdated)
				if err := _KeeperRegistry.contract.UnpackLog(event, "MinCollateralUpdated", log); err != nil {
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

// ParseMinCollateralUpdated is a log parse operation binding the contract event 0xd19fe8ad9152af12b174a60210fb798db0767d63973ebb97298dc44d67a5c82d.
//
// Solidity: event MinCollateralUpdated(uint256 amount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseMinCollateralUpdated(log types.Log) (*KeeperRegistryMinCollateralUpdated, error) {
	event := new(KeeperRegistryMinCollateralUpdated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "MinCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryOffsetOverissuedIterator is returned from FilterOffsetOverissued and is used to iterate over the raw logs and unpacked data for OffsetOverissued events raised by the KeeperRegistry contract.
type KeeperRegistryOffsetOverissuedIterator struct {
	Event *KeeperRegistryOffsetOverissued // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryOffsetOverissuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryOffsetOverissued)
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
		it.Event = new(KeeperRegistryOffsetOverissued)
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
func (it *KeeperRegistryOffsetOverissuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryOffsetOverissuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryOffsetOverissued represents a OffsetOverissued event raised by the KeeperRegistry contract.
type KeeperRegistryOffsetOverissued struct {
	Operator                 common.Address
	SatsAmount               *big.Int
	RemainingOverissueAmount *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterOffsetOverissued is a free log retrieval operation binding the contract event 0xbe43989ace032e7eb39cb9d6d98d37d4b3ef4644f8a190af00046a9176faea47.
//
// Solidity: event OffsetOverissued(address indexed operator, uint256 satsAmount, uint256 remainingOverissueAmount)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterOffsetOverissued(opts *bind.FilterOpts, operator []common.Address) (*KeeperRegistryOffsetOverissuedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "OffsetOverissued", operatorRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryOffsetOverissuedIterator{contract: _KeeperRegistry.contract, event: "OffsetOverissued", logs: logs, sub: sub}, nil
}

// WatchOffsetOverissued is a free log subscription operation binding the contract event 0xbe43989ace032e7eb39cb9d6d98d37d4b3ef4644f8a190af00046a9176faea47.
//
// Solidity: event OffsetOverissued(address indexed operator, uint256 satsAmount, uint256 remainingOverissueAmount)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchOffsetOverissued(opts *bind.WatchOpts, sink chan<- *KeeperRegistryOffsetOverissued, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "OffsetOverissued", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryOffsetOverissued)
				if err := _KeeperRegistry.contract.UnpackLog(event, "OffsetOverissued", log); err != nil {
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

// ParseOffsetOverissued is a log parse operation binding the contract event 0xbe43989ace032e7eb39cb9d6d98d37d4b3ef4644f8a190af00046a9176faea47.
//
// Solidity: event OffsetOverissued(address indexed operator, uint256 satsAmount, uint256 remainingOverissueAmount)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseOffsetOverissued(log types.Log) (*KeeperRegistryOffsetOverissued, error) {
	event := new(KeeperRegistryOffsetOverissued)
	if err := _KeeperRegistry.contract.UnpackLog(event, "OffsetOverissued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryOverissueAddedIterator is returned from FilterOverissueAdded and is used to iterate over the raw logs and unpacked data for OverissueAdded events raised by the KeeperRegistry contract.
type KeeperRegistryOverissueAddedIterator struct {
	Event *KeeperRegistryOverissueAdded // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryOverissueAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryOverissueAdded)
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
		it.Event = new(KeeperRegistryOverissueAdded)
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
func (it *KeeperRegistryOverissueAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryOverissueAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryOverissueAdded represents a OverissueAdded event raised by the KeeperRegistry contract.
type KeeperRegistryOverissueAdded struct {
	Total     *big.Int
	Added     *big.Int
	Deduction *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOverissueAdded is a free log retrieval operation binding the contract event 0x6550740cc389a2fafbbf95e0714109700fa92f1f775d3257527a15107faf4585.
//
// Solidity: event OverissueAdded(uint256 total, uint256 added, uint256 deduction)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterOverissueAdded(opts *bind.FilterOpts) (*KeeperRegistryOverissueAddedIterator, error) {

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "OverissueAdded")
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryOverissueAddedIterator{contract: _KeeperRegistry.contract, event: "OverissueAdded", logs: logs, sub: sub}, nil
}

// WatchOverissueAdded is a free log subscription operation binding the contract event 0x6550740cc389a2fafbbf95e0714109700fa92f1f775d3257527a15107faf4585.
//
// Solidity: event OverissueAdded(uint256 total, uint256 added, uint256 deduction)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchOverissueAdded(opts *bind.WatchOpts, sink chan<- *KeeperRegistryOverissueAdded) (event.Subscription, error) {

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "OverissueAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryOverissueAdded)
				if err := _KeeperRegistry.contract.UnpackLog(event, "OverissueAdded", log); err != nil {
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

// ParseOverissueAdded is a log parse operation binding the contract event 0x6550740cc389a2fafbbf95e0714109700fa92f1f775d3257527a15107faf4585.
//
// Solidity: event OverissueAdded(uint256 total, uint256 added, uint256 deduction)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseOverissueAdded(log types.Log) (*KeeperRegistryOverissueAdded, error) {
	event := new(KeeperRegistryOverissueAdded)
	if err := _KeeperRegistry.contract.UnpackLog(event, "OverissueAdded", log); err != nil {
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

// KeeperRegistrySystemUpdatedIterator is returned from FilterSystemUpdated and is used to iterate over the raw logs and unpacked data for SystemUpdated events raised by the KeeperRegistry contract.
type KeeperRegistrySystemUpdatedIterator struct {
	Event *KeeperRegistrySystemUpdated // Event containing the contract specifics and raw log

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
func (it *KeeperRegistrySystemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistrySystemUpdated)
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
		it.Event = new(KeeperRegistrySystemUpdated)
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
func (it *KeeperRegistrySystemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistrySystemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistrySystemUpdated represents a SystemUpdated event raised by the KeeperRegistry contract.
type KeeperRegistrySystemUpdated struct {
	OldSystem common.Address
	NewSystem common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSystemUpdated is a free log retrieval operation binding the contract event 0xa6e2b13449f1d6c7eeb3cac7c7457ca2c712fcc0b9962fde10e367ff72f10e74.
//
// Solidity: event SystemUpdated(address oldSystem, address newSystem)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterSystemUpdated(opts *bind.FilterOpts) (*KeeperRegistrySystemUpdatedIterator, error) {

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "SystemUpdated")
	if err != nil {
		return nil, err
	}
	return &KeeperRegistrySystemUpdatedIterator{contract: _KeeperRegistry.contract, event: "SystemUpdated", logs: logs, sub: sub}, nil
}

// WatchSystemUpdated is a free log subscription operation binding the contract event 0xa6e2b13449f1d6c7eeb3cac7c7457ca2c712fcc0b9962fde10e367ff72f10e74.
//
// Solidity: event SystemUpdated(address oldSystem, address newSystem)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchSystemUpdated(opts *bind.WatchOpts, sink chan<- *KeeperRegistrySystemUpdated) (event.Subscription, error) {

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "SystemUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistrySystemUpdated)
				if err := _KeeperRegistry.contract.UnpackLog(event, "SystemUpdated", log); err != nil {
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

// ParseSystemUpdated is a log parse operation binding the contract event 0xa6e2b13449f1d6c7eeb3cac7c7457ca2c712fcc0b9962fde10e367ff72f10e74.
//
// Solidity: event SystemUpdated(address oldSystem, address newSystem)
func (_KeeperRegistry *KeeperRegistryFilterer) ParseSystemUpdated(log types.Log) (*KeeperRegistrySystemUpdated, error) {
	event := new(KeeperRegistrySystemUpdated)
	if err := _KeeperRegistry.contract.UnpackLog(event, "SystemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeeperRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KeeperRegistry contract.
type KeeperRegistryTransferIterator struct {
	Event *KeeperRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *KeeperRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeeperRegistryTransfer)
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
		it.Event = new(KeeperRegistryTransfer)
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
func (it *KeeperRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeeperRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeeperRegistryTransfer represents a Transfer event raised by the KeeperRegistry contract.
type KeeperRegistryTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KeeperRegistry *KeeperRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeeperRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeeperRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KeeperRegistryTransferIterator{contract: _KeeperRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KeeperRegistry *KeeperRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KeeperRegistryTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeeperRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeeperRegistryTransfer)
				if err := _KeeperRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KeeperRegistry *KeeperRegistryFilterer) ParseTransfer(log types.Log) (*KeeperRegistryTransfer, error) {
	event := new(KeeperRegistryTransfer)
	if err := _KeeperRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
