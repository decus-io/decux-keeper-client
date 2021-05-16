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

// IDeCusSystemReceipt is an auto generated low-level Go binding around an user-defined struct.
type IDeCusSystemReceipt struct {
	AmountInSatoshi    *big.Int
	UpdateTimestamp    *big.Int
	TxId               [32]byte
	Height             *big.Int
	Status             uint8
	Recipient          common.Address
	GroupBtcAddress    string
	WithdrawBtcAddress string
}

// LibRequestMintRequest is an auto generated low-level Go binding around an user-defined struct.
type LibRequestMintRequest struct {
	ReceiptId [32]byte
	TxId      [32]byte
	Height    *big.Int
}

// DeCusSystemABI is the input ABI used to generate the binding from.
const DeCusSystemABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSatoshi\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"GroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"GroupDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInSatoshi\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"}],\"name\":\"MintRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"MintRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"MintVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_REUSING_GAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_COOLDOWN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REQUEST_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_VERIFICATION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"addGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chillTime\",\"type\":\"uint256\"}],\"name\":\"chill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"deleteGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eBTC\",\"outputs\":[{\"internalType\":\"contractIEBTC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountInSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"forceRequestMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"getCooldownTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"workingReceiptId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountInSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"enumIDeCusSystem.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"}],\"internalType\":\"structIDeCusSystem.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getReceiptId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperRegistry\",\"outputs\":[{\"internalType\":\"contractIKeeperRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minKeeperSatoshi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"recoverBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"}],\"name\":\"requestBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountInSatoshi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"requestMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"revokeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"verifyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"internalType\":\"structLibRequest.MintRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"packedV\",\"type\":\"uint256\"}],\"name\":\"verifyMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeCusSystem is an auto generated Go binding around an Ethereum contract.
type DeCusSystem struct {
	DeCusSystemCaller     // Read-only binding to the contract
	DeCusSystemTransactor // Write-only binding to the contract
	DeCusSystemFilterer   // Log filterer for contract events
}

// DeCusSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeCusSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeCusSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeCusSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeCusSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeCusSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeCusSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeCusSystemSession struct {
	Contract     *DeCusSystem      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeCusSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeCusSystemCallerSession struct {
	Contract *DeCusSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DeCusSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeCusSystemTransactorSession struct {
	Contract     *DeCusSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DeCusSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeCusSystemRaw struct {
	Contract *DeCusSystem // Generic contract binding to access the raw methods on
}

// DeCusSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeCusSystemCallerRaw struct {
	Contract *DeCusSystemCaller // Generic read-only contract binding to access the raw methods on
}

// DeCusSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeCusSystemTransactorRaw struct {
	Contract *DeCusSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeCusSystem creates a new instance of DeCusSystem, bound to a specific deployed contract.
func NewDeCusSystem(address common.Address, backend bind.ContractBackend) (*DeCusSystem, error) {
	contract, err := bindDeCusSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeCusSystem{DeCusSystemCaller: DeCusSystemCaller{contract: contract}, DeCusSystemTransactor: DeCusSystemTransactor{contract: contract}, DeCusSystemFilterer: DeCusSystemFilterer{contract: contract}}, nil
}

// NewDeCusSystemCaller creates a new read-only instance of DeCusSystem, bound to a specific deployed contract.
func NewDeCusSystemCaller(address common.Address, caller bind.ContractCaller) (*DeCusSystemCaller, error) {
	contract, err := bindDeCusSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemCaller{contract: contract}, nil
}

// NewDeCusSystemTransactor creates a new write-only instance of DeCusSystem, bound to a specific deployed contract.
func NewDeCusSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*DeCusSystemTransactor, error) {
	contract, err := bindDeCusSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemTransactor{contract: contract}, nil
}

// NewDeCusSystemFilterer creates a new log filterer instance of DeCusSystem, bound to a specific deployed contract.
func NewDeCusSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*DeCusSystemFilterer, error) {
	contract, err := bindDeCusSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemFilterer{contract: contract}, nil
}

// bindDeCusSystem binds a generic wrapper to an already deployed contract.
func bindDeCusSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeCusSystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeCusSystem *DeCusSystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeCusSystem.Contract.DeCusSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeCusSystem *DeCusSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeCusSystem.Contract.DeCusSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeCusSystem *DeCusSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeCusSystem.Contract.DeCusSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeCusSystem *DeCusSystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeCusSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeCusSystem *DeCusSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeCusSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeCusSystem *DeCusSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeCusSystem.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "EIP712_DOMAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_DeCusSystem *DeCusSystemSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _DeCusSystem.Contract.EIP712DOMAINHASH(&_DeCusSystem.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _DeCusSystem.Contract.EIP712DOMAINHASH(&_DeCusSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) GROUPREUSINGGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "GROUP_REUSING_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) GROUPREUSINGGAP() (*big.Int, error) {
	return _DeCusSystem.Contract.GROUPREUSINGGAP(&_DeCusSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) GROUPREUSINGGAP() (*big.Int, error) {
	return _DeCusSystem.Contract.GROUPREUSINGGAP(&_DeCusSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) KEEPERCOOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "KEEPER_COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) KEEPERCOOLDOWN() (*big.Int, error) {
	return _DeCusSystem.Contract.KEEPERCOOLDOWN(&_DeCusSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) KEEPERCOOLDOWN() (*big.Int, error) {
	return _DeCusSystem.Contract.KEEPERCOOLDOWN(&_DeCusSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) MINTREQUESTGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "MINT_REQUEST_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) MINTREQUESTGRACEPERIOD() (*big.Int, error) {
	return _DeCusSystem.Contract.MINTREQUESTGRACEPERIOD(&_DeCusSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) MINTREQUESTGRACEPERIOD() (*big.Int, error) {
	return _DeCusSystem.Contract.MINTREQUESTGRACEPERIOD(&_DeCusSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) WITHDRAWVERIFICATIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "WITHDRAW_VERIFICATION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) WITHDRAWVERIFICATIONEND() (*big.Int, error) {
	return _DeCusSystem.Contract.WITHDRAWVERIFICATIONEND(&_DeCusSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) WITHDRAWVERIFICATIONEND() (*big.Int, error) {
	return _DeCusSystem.Contract.WITHDRAWVERIFICATIONEND(&_DeCusSystem.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) ChainId() (*big.Int, error) {
	return _DeCusSystem.Contract.ChainId(&_DeCusSystem.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) ChainId() (*big.Int, error) {
	return _DeCusSystem.Contract.ChainId(&_DeCusSystem.CallOpts)
}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) EBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "eBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_DeCusSystem *DeCusSystemSession) EBTC() (common.Address, error) {
	return _DeCusSystem.Contract.EBTC(&_DeCusSystem.CallOpts)
}

// EBTC is a free data retrieval call binding the contract method 0x5f447e04.
//
// Solidity: function eBTC() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) EBTC() (common.Address, error) {
	return _DeCusSystem.Contract.EBTC(&_DeCusSystem.CallOpts)
}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) GetCooldownTime(opts *bind.CallOpts, keeper common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getCooldownTime", keeper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) GetCooldownTime(keeper common.Address) (*big.Int, error) {
	return _DeCusSystem.Contract.GetCooldownTime(&_DeCusSystem.CallOpts, keeper)
}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) GetCooldownTime(keeper common.Address) (*big.Int, error) {
	return _DeCusSystem.Contract.GetCooldownTime(&_DeCusSystem.CallOpts, keeper)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint256 required, uint256 maxSatoshi, uint256 currSatoshi, uint256 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemCaller) GetGroup(opts *bind.CallOpts, btcAddress string) (struct {
	Required         *big.Int
	MaxSatoshi       *big.Int
	CurrSatoshi      *big.Int
	Nonce            *big.Int
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getGroup", btcAddress)

	outstruct := new(struct {
		Required         *big.Int
		MaxSatoshi       *big.Int
		CurrSatoshi      *big.Int
		Nonce            *big.Int
		Keepers          []common.Address
		WorkingReceiptId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Required = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxSatoshi = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrSatoshi = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Keepers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)
	outstruct.WorkingReceiptId = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint256 required, uint256 maxSatoshi, uint256 currSatoshi, uint256 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemSession) GetGroup(btcAddress string) (struct {
	Required         *big.Int
	MaxSatoshi       *big.Int
	CurrSatoshi      *big.Int
	Nonce            *big.Int
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DeCusSystem.Contract.GetGroup(&_DeCusSystem.CallOpts, btcAddress)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint256 required, uint256 maxSatoshi, uint256 currSatoshi, uint256 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemCallerSession) GetGroup(btcAddress string) (struct {
	Required         *big.Int
	MaxSatoshi       *big.Int
	CurrSatoshi      *big.Int
	Nonce            *big.Int
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DeCusSystem.Contract.GetGroup(&_DeCusSystem.CallOpts, btcAddress)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((uint256,uint256,bytes32,uint256,uint8,address,string,string))
func (_DeCusSystem *DeCusSystemCaller) GetReceipt(opts *bind.CallOpts, receiptId [32]byte) (IDeCusSystemReceipt, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getReceipt", receiptId)

	if err != nil {
		return *new(IDeCusSystemReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(IDeCusSystemReceipt)).(*IDeCusSystemReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((uint256,uint256,bytes32,uint256,uint8,address,string,string))
func (_DeCusSystem *DeCusSystemSession) GetReceipt(receiptId [32]byte) (IDeCusSystemReceipt, error) {
	return _DeCusSystem.Contract.GetReceipt(&_DeCusSystem.CallOpts, receiptId)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((uint256,uint256,bytes32,uint256,uint8,address,string,string))
func (_DeCusSystem *DeCusSystemCallerSession) GetReceipt(receiptId [32]byte) (IDeCusSystemReceipt, error) {
	return _DeCusSystem.Contract.GetReceipt(&_DeCusSystem.CallOpts, receiptId)
}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DeCusSystem *DeCusSystemCaller) GetReceiptId(opts *bind.CallOpts, groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getReceiptId", groupBtcAddress, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DeCusSystem *DeCusSystemSession) GetReceiptId(groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	return _DeCusSystem.Contract.GetReceiptId(&_DeCusSystem.CallOpts, groupBtcAddress, nonce)
}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DeCusSystem *DeCusSystemCallerSession) GetReceiptId(groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	return _DeCusSystem.Contract.GetReceiptId(&_DeCusSystem.CallOpts, groupBtcAddress, nonce)
}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) KeeperRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "keeperRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DeCusSystem *DeCusSystemSession) KeeperRegistry() (common.Address, error) {
	return _DeCusSystem.Contract.KeeperRegistry(&_DeCusSystem.CallOpts)
}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) KeeperRegistry() (common.Address, error) {
	return _DeCusSystem.Contract.KeeperRegistry(&_DeCusSystem.CallOpts)
}

// MinKeeperSatoshi is a free data retrieval call binding the contract method 0xad95033a.
//
// Solidity: function minKeeperSatoshi() view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) MinKeeperSatoshi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "minKeeperSatoshi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinKeeperSatoshi is a free data retrieval call binding the contract method 0xad95033a.
//
// Solidity: function minKeeperSatoshi() view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) MinKeeperSatoshi() (*big.Int, error) {
	return _DeCusSystem.Contract.MinKeeperSatoshi(&_DeCusSystem.CallOpts)
}

// MinKeeperSatoshi is a free data retrieval call binding the contract method 0xad95033a.
//
// Solidity: function minKeeperSatoshi() view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) MinKeeperSatoshi() (*big.Int, error) {
	return _DeCusSystem.Contract.MinKeeperSatoshi(&_DeCusSystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeCusSystem *DeCusSystemSession) Owner() (common.Address, error) {
	return _DeCusSystem.Contract.Owner(&_DeCusSystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) Owner() (common.Address, error) {
	return _DeCusSystem.Contract.Owner(&_DeCusSystem.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DeCusSystem *DeCusSystemCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DeCusSystem *DeCusSystemSession) Paused() (bool, error) {
	return _DeCusSystem.Contract.Paused(&_DeCusSystem.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DeCusSystem *DeCusSystemCallerSession) Paused() (bool, error) {
	return _DeCusSystem.Contract.Paused(&_DeCusSystem.CallOpts)
}

// AddGroup is a paid mutator transaction binding the contract method 0xc5a746bf.
//
// Solidity: function addGroup(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemTransactor) AddGroup(opts *bind.TransactOpts, btcAddress string, required *big.Int, maxSatoshi *big.Int, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "addGroup", btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xc5a746bf.
//
// Solidity: function addGroup(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemSession) AddGroup(btcAddress string, required *big.Int, maxSatoshi *big.Int, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.AddGroup(&_DeCusSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xc5a746bf.
//
// Solidity: function addGroup(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) AddGroup(btcAddress string, required *big.Int, maxSatoshi *big.Int, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.AddGroup(&_DeCusSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// Chill is a paid mutator transaction binding the contract method 0xc976966a.
//
// Solidity: function chill(address keeper, uint256 chillTime) returns()
func (_DeCusSystem *DeCusSystemTransactor) Chill(opts *bind.TransactOpts, keeper common.Address, chillTime *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "chill", keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0xc976966a.
//
// Solidity: function chill(address keeper, uint256 chillTime) returns()
func (_DeCusSystem *DeCusSystemSession) Chill(keeper common.Address, chillTime *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Chill(&_DeCusSystem.TransactOpts, keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0xc976966a.
//
// Solidity: function chill(address keeper, uint256 chillTime) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Chill(keeper common.Address, chillTime *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Chill(&_DeCusSystem.TransactOpts, keeper, chillTime)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DeCusSystem *DeCusSystemTransactor) DeleteGroup(opts *bind.TransactOpts, btcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "deleteGroup", btcAddress)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DeCusSystem *DeCusSystemSession) DeleteGroup(btcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.Contract.DeleteGroup(&_DeCusSystem.TransactOpts, btcAddress)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) DeleteGroup(btcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.Contract.DeleteGroup(&_DeCusSystem.TransactOpts, btcAddress)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x6d3c5abc.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemTransactor) ForceRequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "forceRequestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x6d3c5abc.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.ForceRequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x6d3c5abc.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.ForceRequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _eBTC, address _registry) returns()
func (_DeCusSystem *DeCusSystemTransactor) Initialize(opts *bind.TransactOpts, _eBTC common.Address, _registry common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "initialize", _eBTC, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _eBTC, address _registry) returns()
func (_DeCusSystem *DeCusSystemSession) Initialize(_eBTC common.Address, _registry common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Initialize(&_DeCusSystem.TransactOpts, _eBTC, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _eBTC, address _registry) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Initialize(_eBTC common.Address, _registry common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Initialize(&_DeCusSystem.TransactOpts, _eBTC, _registry)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactor) RecoverBurn(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "recoverBurn", receiptId)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemSession) RecoverBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RecoverBurn(&_DeCusSystem.TransactOpts, receiptId)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RecoverBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RecoverBurn(&_DeCusSystem.TransactOpts, receiptId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeCusSystem *DeCusSystemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeCusSystem *DeCusSystemSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeCusSystem.Contract.RenounceOwnership(&_DeCusSystem.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeCusSystem.Contract.RenounceOwnership(&_DeCusSystem.TransactOpts)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DeCusSystem *DeCusSystemTransactor) RequestBurn(opts *bind.TransactOpts, receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "requestBurn", receiptId, withdrawBtcAddress)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DeCusSystem *DeCusSystemSession) RequestBurn(receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RequestBurn(&_DeCusSystem.TransactOpts, receiptId, withdrawBtcAddress)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RequestBurn(receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RequestBurn(&_DeCusSystem.TransactOpts, receiptId, withdrawBtcAddress)
}

// RequestMint is a paid mutator transaction binding the contract method 0x9123cdfc.
//
// Solidity: function requestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemTransactor) RequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "requestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0x9123cdfc.
//
// Solidity: function requestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemSession) RequestMint(groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0x9123cdfc.
//
// Solidity: function requestMint(string groupBtcAddress, uint256 amountInSatoshi, uint256 nonce) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RequestMint(groupBtcAddress string, amountInSatoshi *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactor) RevokeMint(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "revokeMint", receiptId)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemSession) RevokeMint(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RevokeMint(&_DeCusSystem.TransactOpts, receiptId)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RevokeMint(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RevokeMint(&_DeCusSystem.TransactOpts, receiptId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeCusSystem *DeCusSystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeCusSystem *DeCusSystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.TransferOwnership(&_DeCusSystem.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.TransferOwnership(&_DeCusSystem.TransactOpts, newOwner)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactor) VerifyBurn(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "verifyBurn", receiptId)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemSession) VerifyBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyBurn(&_DeCusSystem.TransactOpts, receiptId)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) VerifyBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyBurn(&_DeCusSystem.TransactOpts, receiptId)
}

// VerifyMint is a paid mutator transaction binding the contract method 0xc764654c.
//
// Solidity: function verifyMint((bytes32,bytes32,uint256) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemTransactor) VerifyMint(opts *bind.TransactOpts, request LibRequestMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "verifyMint", request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0xc764654c.
//
// Solidity: function verifyMint((bytes32,bytes32,uint256) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemSession) VerifyMint(request LibRequestMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyMint(&_DeCusSystem.TransactOpts, request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0xc764654c.
//
// Solidity: function verifyMint((bytes32,bytes32,uint256) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) VerifyMint(request LibRequestMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyMint(&_DeCusSystem.TransactOpts, request, keepers, r, s, packedV)
}

// DeCusSystemBurnRequestedIterator is returned from FilterBurnRequested and is used to iterate over the raw logs and unpacked data for BurnRequested events raised by the DeCusSystem contract.
type DeCusSystemBurnRequestedIterator struct {
	Event *DeCusSystemBurnRequested // Event containing the contract specifics and raw log

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
func (it *DeCusSystemBurnRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemBurnRequested)
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
		it.Event = new(DeCusSystemBurnRequested)
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
func (it *DeCusSystemBurnRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemBurnRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemBurnRequested represents a BurnRequested event raised by the DeCusSystem contract.
type DeCusSystemBurnRequested struct {
	ReceiptId          [32]byte
	GroupBtcAddress    string
	WithdrawBtcAddress string
	Operator           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBurnRequested is a free log retrieval operation binding the contract event 0xf0e638df6d296aaddeb18409852c4fa659a87b113a1fdb9cf7535668bd49497d.
//
// Solidity: event BurnRequested(bytes32 indexed receiptId, string groupBtcAddress, string withdrawBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) FilterBurnRequested(opts *bind.FilterOpts, receiptId [][32]byte) (*DeCusSystemBurnRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "BurnRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemBurnRequestedIterator{contract: _DeCusSystem.contract, event: "BurnRequested", logs: logs, sub: sub}, nil
}

// WatchBurnRequested is a free log subscription operation binding the contract event 0xf0e638df6d296aaddeb18409852c4fa659a87b113a1fdb9cf7535668bd49497d.
//
// Solidity: event BurnRequested(bytes32 indexed receiptId, string groupBtcAddress, string withdrawBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) WatchBurnRequested(opts *bind.WatchOpts, sink chan<- *DeCusSystemBurnRequested, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "BurnRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemBurnRequested)
				if err := _DeCusSystem.contract.UnpackLog(event, "BurnRequested", log); err != nil {
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

// ParseBurnRequested is a log parse operation binding the contract event 0xf0e638df6d296aaddeb18409852c4fa659a87b113a1fdb9cf7535668bd49497d.
//
// Solidity: event BurnRequested(bytes32 indexed receiptId, string groupBtcAddress, string withdrawBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) ParseBurnRequested(log types.Log) (*DeCusSystemBurnRequested, error) {
	event := new(DeCusSystemBurnRequested)
	if err := _DeCusSystem.contract.UnpackLog(event, "BurnRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemBurnRevokedIterator is returned from FilterBurnRevoked and is used to iterate over the raw logs and unpacked data for BurnRevoked events raised by the DeCusSystem contract.
type DeCusSystemBurnRevokedIterator struct {
	Event *DeCusSystemBurnRevoked // Event containing the contract specifics and raw log

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
func (it *DeCusSystemBurnRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemBurnRevoked)
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
		it.Event = new(DeCusSystemBurnRevoked)
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
func (it *DeCusSystemBurnRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemBurnRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemBurnRevoked represents a BurnRevoked event raised by the DeCusSystem contract.
type DeCusSystemBurnRevoked struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurnRevoked is a free log retrieval operation binding the contract event 0x8e24f8539f9ec3a6069ecf3348ec641da0bf8d223c4e5a69611809462087a9e8.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) FilterBurnRevoked(opts *bind.FilterOpts, receiptId [][32]byte) (*DeCusSystemBurnRevokedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "BurnRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemBurnRevokedIterator{contract: _DeCusSystem.contract, event: "BurnRevoked", logs: logs, sub: sub}, nil
}

// WatchBurnRevoked is a free log subscription operation binding the contract event 0x8e24f8539f9ec3a6069ecf3348ec641da0bf8d223c4e5a69611809462087a9e8.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) WatchBurnRevoked(opts *bind.WatchOpts, sink chan<- *DeCusSystemBurnRevoked, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "BurnRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemBurnRevoked)
				if err := _DeCusSystem.contract.UnpackLog(event, "BurnRevoked", log); err != nil {
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

// ParseBurnRevoked is a log parse operation binding the contract event 0x8e24f8539f9ec3a6069ecf3348ec641da0bf8d223c4e5a69611809462087a9e8.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) ParseBurnRevoked(log types.Log) (*DeCusSystemBurnRevoked, error) {
	event := new(DeCusSystemBurnRevoked)
	if err := _DeCusSystem.contract.UnpackLog(event, "BurnRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemBurnVerifiedIterator is returned from FilterBurnVerified and is used to iterate over the raw logs and unpacked data for BurnVerified events raised by the DeCusSystem contract.
type DeCusSystemBurnVerifiedIterator struct {
	Event *DeCusSystemBurnVerified // Event containing the contract specifics and raw log

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
func (it *DeCusSystemBurnVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemBurnVerified)
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
		it.Event = new(DeCusSystemBurnVerified)
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
func (it *DeCusSystemBurnVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemBurnVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemBurnVerified represents a BurnVerified event raised by the DeCusSystem contract.
type DeCusSystemBurnVerified struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurnVerified is a free log retrieval operation binding the contract event 0xd90232357c1ad1e207410849ce0a0374bac4c22771453838f285ceca745547e3.
//
// Solidity: event BurnVerified(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) FilterBurnVerified(opts *bind.FilterOpts, receiptId [][32]byte) (*DeCusSystemBurnVerifiedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "BurnVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemBurnVerifiedIterator{contract: _DeCusSystem.contract, event: "BurnVerified", logs: logs, sub: sub}, nil
}

// WatchBurnVerified is a free log subscription operation binding the contract event 0xd90232357c1ad1e207410849ce0a0374bac4c22771453838f285ceca745547e3.
//
// Solidity: event BurnVerified(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) WatchBurnVerified(opts *bind.WatchOpts, sink chan<- *DeCusSystemBurnVerified, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "BurnVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemBurnVerified)
				if err := _DeCusSystem.contract.UnpackLog(event, "BurnVerified", log); err != nil {
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

// ParseBurnVerified is a log parse operation binding the contract event 0xd90232357c1ad1e207410849ce0a0374bac4c22771453838f285ceca745547e3.
//
// Solidity: event BurnVerified(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) ParseBurnVerified(log types.Log) (*DeCusSystemBurnVerified, error) {
	event := new(DeCusSystemBurnVerified)
	if err := _DeCusSystem.contract.UnpackLog(event, "BurnVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the DeCusSystem contract.
type DeCusSystemCooldownIterator struct {
	Event *DeCusSystemCooldown // Event containing the contract specifics and raw log

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
func (it *DeCusSystemCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemCooldown)
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
		it.Event = new(DeCusSystemCooldown)
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
func (it *DeCusSystemCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemCooldown represents a Cooldown event raised by the DeCusSystem contract.
type DeCusSystemCooldown struct {
	Keeper  common.Address
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed keeper, uint256 endTime)
func (_DeCusSystem *DeCusSystemFilterer) FilterCooldown(opts *bind.FilterOpts, keeper []common.Address) (*DeCusSystemCooldownIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "Cooldown", keeperRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemCooldownIterator{contract: _DeCusSystem.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed keeper, uint256 endTime)
func (_DeCusSystem *DeCusSystemFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *DeCusSystemCooldown, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "Cooldown", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemCooldown)
				if err := _DeCusSystem.contract.UnpackLog(event, "Cooldown", log); err != nil {
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

// ParseCooldown is a log parse operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed keeper, uint256 endTime)
func (_DeCusSystem *DeCusSystemFilterer) ParseCooldown(log types.Log) (*DeCusSystemCooldown, error) {
	event := new(DeCusSystemCooldown)
	if err := _DeCusSystem.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemGroupAddedIterator is returned from FilterGroupAdded and is used to iterate over the raw logs and unpacked data for GroupAdded events raised by the DeCusSystem contract.
type DeCusSystemGroupAddedIterator struct {
	Event *DeCusSystemGroupAdded // Event containing the contract specifics and raw log

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
func (it *DeCusSystemGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemGroupAdded)
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
		it.Event = new(DeCusSystemGroupAdded)
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
func (it *DeCusSystemGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemGroupAdded represents a GroupAdded event raised by the DeCusSystem contract.
type DeCusSystemGroupAdded struct {
	BtcAddress string
	Required   *big.Int
	MaxSatoshi *big.Int
	Keepers    []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0x318cdd783f12822d86ff4e2e147ae395b85e9e163d0f2f36eefe022364c06213.
//
// Solidity: event GroupAdded(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) FilterGroupAdded(opts *bind.FilterOpts) (*DeCusSystemGroupAddedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemGroupAddedIterator{contract: _DeCusSystem.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0x318cdd783f12822d86ff4e2e147ae395b85e9e163d0f2f36eefe022364c06213.
//
// Solidity: event GroupAdded(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) WatchGroupAdded(opts *bind.WatchOpts, sink chan<- *DeCusSystemGroupAdded) (event.Subscription, error) {

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemGroupAdded)
				if err := _DeCusSystem.contract.UnpackLog(event, "GroupAdded", log); err != nil {
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

// ParseGroupAdded is a log parse operation binding the contract event 0x318cdd783f12822d86ff4e2e147ae395b85e9e163d0f2f36eefe022364c06213.
//
// Solidity: event GroupAdded(string btcAddress, uint256 required, uint256 maxSatoshi, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) ParseGroupAdded(log types.Log) (*DeCusSystemGroupAdded, error) {
	event := new(DeCusSystemGroupAdded)
	if err := _DeCusSystem.contract.UnpackLog(event, "GroupAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemGroupDeletedIterator is returned from FilterGroupDeleted and is used to iterate over the raw logs and unpacked data for GroupDeleted events raised by the DeCusSystem contract.
type DeCusSystemGroupDeletedIterator struct {
	Event *DeCusSystemGroupDeleted // Event containing the contract specifics and raw log

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
func (it *DeCusSystemGroupDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemGroupDeleted)
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
		it.Event = new(DeCusSystemGroupDeleted)
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
func (it *DeCusSystemGroupDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemGroupDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemGroupDeleted represents a GroupDeleted event raised by the DeCusSystem contract.
type DeCusSystemGroupDeleted struct {
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupDeleted is a free log retrieval operation binding the contract event 0xeacdb5801d8160882398b78a544379ff28cbe8bd4de76aca39013c89311e00d5.
//
// Solidity: event GroupDeleted(string btcAddress)
func (_DeCusSystem *DeCusSystemFilterer) FilterGroupDeleted(opts *bind.FilterOpts) (*DeCusSystemGroupDeletedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "GroupDeleted")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemGroupDeletedIterator{contract: _DeCusSystem.contract, event: "GroupDeleted", logs: logs, sub: sub}, nil
}

// WatchGroupDeleted is a free log subscription operation binding the contract event 0xeacdb5801d8160882398b78a544379ff28cbe8bd4de76aca39013c89311e00d5.
//
// Solidity: event GroupDeleted(string btcAddress)
func (_DeCusSystem *DeCusSystemFilterer) WatchGroupDeleted(opts *bind.WatchOpts, sink chan<- *DeCusSystemGroupDeleted) (event.Subscription, error) {

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "GroupDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemGroupDeleted)
				if err := _DeCusSystem.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
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

// ParseGroupDeleted is a log parse operation binding the contract event 0xeacdb5801d8160882398b78a544379ff28cbe8bd4de76aca39013c89311e00d5.
//
// Solidity: event GroupDeleted(string btcAddress)
func (_DeCusSystem *DeCusSystemFilterer) ParseGroupDeleted(log types.Log) (*DeCusSystemGroupDeleted, error) {
	event := new(DeCusSystemGroupDeleted)
	if err := _DeCusSystem.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemMintRequestedIterator is returned from FilterMintRequested and is used to iterate over the raw logs and unpacked data for MintRequested events raised by the DeCusSystem contract.
type DeCusSystemMintRequestedIterator struct {
	Event *DeCusSystemMintRequested // Event containing the contract specifics and raw log

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
func (it *DeCusSystemMintRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemMintRequested)
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
		it.Event = new(DeCusSystemMintRequested)
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
func (it *DeCusSystemMintRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemMintRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemMintRequested represents a MintRequested event raised by the DeCusSystem contract.
type DeCusSystemMintRequested struct {
	ReceiptId       [32]byte
	Recipient       common.Address
	AmountInSatoshi *big.Int
	GroupBtcAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintRequested is a free log retrieval operation binding the contract event 0x5e67e1245326e4daf00735ab1689559d49d833312858659221cf9a1d50f20d48.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint256 amountInSatoshi, string groupBtcAddress)
func (_DeCusSystem *DeCusSystemFilterer) FilterMintRequested(opts *bind.FilterOpts, receiptId [][32]byte, recipient []common.Address) (*DeCusSystemMintRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "MintRequested", receiptIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemMintRequestedIterator{contract: _DeCusSystem.contract, event: "MintRequested", logs: logs, sub: sub}, nil
}

// WatchMintRequested is a free log subscription operation binding the contract event 0x5e67e1245326e4daf00735ab1689559d49d833312858659221cf9a1d50f20d48.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint256 amountInSatoshi, string groupBtcAddress)
func (_DeCusSystem *DeCusSystemFilterer) WatchMintRequested(opts *bind.WatchOpts, sink chan<- *DeCusSystemMintRequested, receiptId [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "MintRequested", receiptIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemMintRequested)
				if err := _DeCusSystem.contract.UnpackLog(event, "MintRequested", log); err != nil {
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

// ParseMintRequested is a log parse operation binding the contract event 0x5e67e1245326e4daf00735ab1689559d49d833312858659221cf9a1d50f20d48.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint256 amountInSatoshi, string groupBtcAddress)
func (_DeCusSystem *DeCusSystemFilterer) ParseMintRequested(log types.Log) (*DeCusSystemMintRequested, error) {
	event := new(DeCusSystemMintRequested)
	if err := _DeCusSystem.contract.UnpackLog(event, "MintRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemMintRevokedIterator is returned from FilterMintRevoked and is used to iterate over the raw logs and unpacked data for MintRevoked events raised by the DeCusSystem contract.
type DeCusSystemMintRevokedIterator struct {
	Event *DeCusSystemMintRevoked // Event containing the contract specifics and raw log

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
func (it *DeCusSystemMintRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemMintRevoked)
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
		it.Event = new(DeCusSystemMintRevoked)
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
func (it *DeCusSystemMintRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemMintRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemMintRevoked represents a MintRevoked event raised by the DeCusSystem contract.
type DeCusSystemMintRevoked struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintRevoked is a free log retrieval operation binding the contract event 0x24da21178c24a520845b59d6d351ecf3033e3ceefb8c66f286f667caa74cd2c3.
//
// Solidity: event MintRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) FilterMintRevoked(opts *bind.FilterOpts, receiptId [][32]byte) (*DeCusSystemMintRevokedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "MintRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemMintRevokedIterator{contract: _DeCusSystem.contract, event: "MintRevoked", logs: logs, sub: sub}, nil
}

// WatchMintRevoked is a free log subscription operation binding the contract event 0x24da21178c24a520845b59d6d351ecf3033e3ceefb8c66f286f667caa74cd2c3.
//
// Solidity: event MintRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) WatchMintRevoked(opts *bind.WatchOpts, sink chan<- *DeCusSystemMintRevoked, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "MintRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemMintRevoked)
				if err := _DeCusSystem.contract.UnpackLog(event, "MintRevoked", log); err != nil {
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

// ParseMintRevoked is a log parse operation binding the contract event 0x24da21178c24a520845b59d6d351ecf3033e3ceefb8c66f286f667caa74cd2c3.
//
// Solidity: event MintRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DeCusSystem *DeCusSystemFilterer) ParseMintRevoked(log types.Log) (*DeCusSystemMintRevoked, error) {
	event := new(DeCusSystemMintRevoked)
	if err := _DeCusSystem.contract.UnpackLog(event, "MintRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemMintVerifiedIterator is returned from FilterMintVerified and is used to iterate over the raw logs and unpacked data for MintVerified events raised by the DeCusSystem contract.
type DeCusSystemMintVerifiedIterator struct {
	Event *DeCusSystemMintVerified // Event containing the contract specifics and raw log

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
func (it *DeCusSystemMintVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemMintVerified)
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
		it.Event = new(DeCusSystemMintVerified)
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
func (it *DeCusSystemMintVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemMintVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemMintVerified represents a MintVerified event raised by the DeCusSystem contract.
type DeCusSystemMintVerified struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Keepers         []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintVerified is a free log retrieval operation binding the contract event 0xf2b0b070a7c8b8436fc8c16f75c3f8d2ae027ca97adc891f7c450f7b46cce531.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) FilterMintVerified(opts *bind.FilterOpts, receiptId [][32]byte) (*DeCusSystemMintVerifiedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "MintVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemMintVerifiedIterator{contract: _DeCusSystem.contract, event: "MintVerified", logs: logs, sub: sub}, nil
}

// WatchMintVerified is a free log subscription operation binding the contract event 0xf2b0b070a7c8b8436fc8c16f75c3f8d2ae027ca97adc891f7c450f7b46cce531.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) WatchMintVerified(opts *bind.WatchOpts, sink chan<- *DeCusSystemMintVerified, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "MintVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemMintVerified)
				if err := _DeCusSystem.contract.UnpackLog(event, "MintVerified", log); err != nil {
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

// ParseMintVerified is a log parse operation binding the contract event 0xf2b0b070a7c8b8436fc8c16f75c3f8d2ae027ca97adc891f7c450f7b46cce531.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) ParseMintVerified(log types.Log) (*DeCusSystemMintVerified, error) {
	event := new(DeCusSystemMintVerified)
	if err := _DeCusSystem.contract.UnpackLog(event, "MintVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DeCusSystem contract.
type DeCusSystemOwnershipTransferredIterator struct {
	Event *DeCusSystemOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeCusSystemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemOwnershipTransferred)
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
		it.Event = new(DeCusSystemOwnershipTransferred)
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
func (it *DeCusSystemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemOwnershipTransferred represents a OwnershipTransferred event raised by the DeCusSystem contract.
type DeCusSystemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeCusSystem *DeCusSystemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeCusSystemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemOwnershipTransferredIterator{contract: _DeCusSystem.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeCusSystem *DeCusSystemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeCusSystemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemOwnershipTransferred)
				if err := _DeCusSystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParseOwnershipTransferred(log types.Log) (*DeCusSystemOwnershipTransferred, error) {
	event := new(DeCusSystemOwnershipTransferred)
	if err := _DeCusSystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DeCusSystem contract.
type DeCusSystemPausedIterator struct {
	Event *DeCusSystemPaused // Event containing the contract specifics and raw log

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
func (it *DeCusSystemPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemPaused)
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
		it.Event = new(DeCusSystemPaused)
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
func (it *DeCusSystemPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemPaused represents a Paused event raised by the DeCusSystem contract.
type DeCusSystemPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DeCusSystem *DeCusSystemFilterer) FilterPaused(opts *bind.FilterOpts) (*DeCusSystemPausedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemPausedIterator{contract: _DeCusSystem.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DeCusSystem *DeCusSystemFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DeCusSystemPaused) (event.Subscription, error) {

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemPaused)
				if err := _DeCusSystem.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParsePaused(log types.Log) (*DeCusSystemPaused, error) {
	event := new(DeCusSystemPaused)
	if err := _DeCusSystem.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DeCusSystem contract.
type DeCusSystemUnpausedIterator struct {
	Event *DeCusSystemUnpaused // Event containing the contract specifics and raw log

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
func (it *DeCusSystemUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemUnpaused)
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
		it.Event = new(DeCusSystemUnpaused)
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
func (it *DeCusSystemUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemUnpaused represents a Unpaused event raised by the DeCusSystem contract.
type DeCusSystemUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DeCusSystem *DeCusSystemFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DeCusSystemUnpausedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemUnpausedIterator{contract: _DeCusSystem.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DeCusSystem *DeCusSystemFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DeCusSystemUnpaused) (event.Subscription, error) {

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemUnpaused)
				if err := _DeCusSystem.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParseUnpaused(log types.Log) (*DeCusSystemUnpaused, error) {
	event := new(DeCusSystemUnpaused)
	if err := _DeCusSystem.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
