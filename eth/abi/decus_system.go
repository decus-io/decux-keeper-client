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

// IDeCusSystemBtcRefundData is an auto generated low-level Go binding around an user-defined struct.
type IDeCusSystemBtcRefundData struct {
	TxId            [32]byte
	GroupBtcAddress string
	ExpiryTimestamp uint32
}

// IDeCusSystemMintRequest is an auto generated low-level Go binding around an user-defined struct.
type IDeCusSystemMintRequest struct {
	ReceiptId [32]byte
	TxId      [32]byte
	Height    uint32
}

// IDeCusSystemReceipt is an auto generated low-level Go binding around an user-defined struct.
type IDeCusSystemReceipt struct {
	GroupBtcAddress    string
	WithdrawBtcAddress string
	TxId               [32]byte
	AmountInSatoshi    uint32
	UpdateTimestamp    uint32
	Height             uint32
	Recipient          common.Address
	Status             uint8
}

// DeCusSystemABI is the input ABI used to generate the binding from.
const DeCusSystemABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"expiryTimestamp\",\"type\":\"uint32\"}],\"name\":\"BtcRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"GroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"GroupDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"}],\"name\":\"MintRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"MintRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"btcTxId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"btcTxHeight\",\"type\":\"uint32\"}],\"name\":\"MintVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_REUSING_GAP\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_COOLDOWN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REQUEST_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_GAP\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_VERIFICATION_END\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"addGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chillTime\",\"type\":\"uint32\"}],\"name\":\"chill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"deleteGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"contractISwapFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"forceRequestMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"getCooldownTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"currSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"workingReceiptId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"updateTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIDeCusSystem.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIDeCusSystem.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getReceiptId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRefundData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expiryTimestamp\",\"type\":\"uint32\"}],\"internalType\":\"structIDeCusSystem.BtcRefundData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIToken\",\"name\":\"_sats\",\"type\":\"address\"},{\"internalType\":\"contractIKeeperRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractISwapRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"contractISwapFee\",\"name\":\"_fee\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperRegistry\",\"outputs\":[{\"internalType\":\"contractIKeeperRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"recoverBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"}],\"name\":\"refundBtc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"}],\"name\":\"requestBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"requestMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"revokeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewarder\",\"outputs\":[{\"internalType\":\"contractISwapRewarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sats\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"verifyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"}],\"internalType\":\"structIDeCusSystem.MintRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"packedV\",\"type\":\"uint256\"}],\"name\":\"verifyMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DeCusSystem.Contract.DEFAULTADMINROLE(&_DeCusSystem.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DeCusSystem.Contract.DEFAULTADMINROLE(&_DeCusSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) GROUPREUSINGGAP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "GROUP_REUSING_GAP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) GROUPREUSINGGAP() (uint32, error) {
	return _DeCusSystem.Contract.GROUPREUSINGGAP(&_DeCusSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) GROUPREUSINGGAP() (uint32, error) {
	return _DeCusSystem.Contract.GROUPREUSINGGAP(&_DeCusSystem.CallOpts)
}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCaller) GROUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "GROUP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemSession) GROUPROLE() ([32]byte, error) {
	return _DeCusSystem.Contract.GROUPROLE(&_DeCusSystem.CallOpts)
}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DeCusSystem *DeCusSystemCallerSession) GROUPROLE() ([32]byte, error) {
	return _DeCusSystem.Contract.GROUPROLE(&_DeCusSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) KEEPERCOOLDOWN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "KEEPER_COOLDOWN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) KEEPERCOOLDOWN() (uint32, error) {
	return _DeCusSystem.Contract.KEEPERCOOLDOWN(&_DeCusSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) KEEPERCOOLDOWN() (uint32, error) {
	return _DeCusSystem.Contract.KEEPERCOOLDOWN(&_DeCusSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) MINTREQUESTGRACEPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "MINT_REQUEST_GRACE_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) MINTREQUESTGRACEPERIOD() (uint32, error) {
	return _DeCusSystem.Contract.MINTREQUESTGRACEPERIOD(&_DeCusSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) MINTREQUESTGRACEPERIOD() (uint32, error) {
	return _DeCusSystem.Contract.MINTREQUESTGRACEPERIOD(&_DeCusSystem.CallOpts)
}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) REFUNDGAP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "REFUND_GAP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) REFUNDGAP() (uint32, error) {
	return _DeCusSystem.Contract.REFUNDGAP(&_DeCusSystem.CallOpts)
}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) REFUNDGAP() (uint32, error) {
	return _DeCusSystem.Contract.REFUNDGAP(&_DeCusSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) WITHDRAWVERIFICATIONEND(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "WITHDRAW_VERIFICATION_END")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) WITHDRAWVERIFICATIONEND() (uint32, error) {
	return _DeCusSystem.Contract.WITHDRAWVERIFICATIONEND(&_DeCusSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) WITHDRAWVERIFICATIONEND() (uint32, error) {
	return _DeCusSystem.Contract.WITHDRAWVERIFICATIONEND(&_DeCusSystem.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) Fee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DeCusSystem *DeCusSystemSession) Fee() (common.Address, error) {
	return _DeCusSystem.Contract.Fee(&_DeCusSystem.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) Fee() (common.Address, error) {
	return _DeCusSystem.Contract.Fee(&_DeCusSystem.CallOpts)
}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint32)
func (_DeCusSystem *DeCusSystemCaller) GetCooldownTime(opts *bind.CallOpts, keeper common.Address) (uint32, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getCooldownTime", keeper)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint32)
func (_DeCusSystem *DeCusSystemSession) GetCooldownTime(keeper common.Address) (uint32, error) {
	return _DeCusSystem.Contract.GetCooldownTime(&_DeCusSystem.CallOpts, keeper)
}

// GetCooldownTime is a free data retrieval call binding the contract method 0x85c95a1e.
//
// Solidity: function getCooldownTime(address keeper) view returns(uint32)
func (_DeCusSystem *DeCusSystemCallerSession) GetCooldownTime(keeper common.Address) (uint32, error) {
	return _DeCusSystem.Contract.GetCooldownTime(&_DeCusSystem.CallOpts, keeper)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint32 required, uint32 maxSatoshi, uint32 currSatoshi, uint32 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemCaller) GetGroup(opts *bind.CallOpts, btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getGroup", btcAddress)

	outstruct := new(struct {
		Required         uint32
		MaxSatoshi       uint32
		CurrSatoshi      uint32
		Nonce            uint32
		Keepers          []common.Address
		WorkingReceiptId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Required = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaxSatoshi = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CurrSatoshi = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Nonce = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Keepers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)
	outstruct.WorkingReceiptId = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint32 required, uint32 maxSatoshi, uint32 currSatoshi, uint32 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemSession) GetGroup(btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DeCusSystem.Contract.GetGroup(&_DeCusSystem.CallOpts, btcAddress)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint32 required, uint32 maxSatoshi, uint32 currSatoshi, uint32 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DeCusSystem *DeCusSystemCallerSession) GetGroup(btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DeCusSystem.Contract.GetGroup(&_DeCusSystem.CallOpts, btcAddress)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
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
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
func (_DeCusSystem *DeCusSystemSession) GetReceipt(receiptId [32]byte) (IDeCusSystemReceipt, error) {
	return _DeCusSystem.Contract.GetReceipt(&_DeCusSystem.CallOpts, receiptId)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
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

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DeCusSystem *DeCusSystemCaller) GetRefundData(opts *bind.CallOpts) (IDeCusSystemBtcRefundData, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getRefundData")

	if err != nil {
		return *new(IDeCusSystemBtcRefundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDeCusSystemBtcRefundData)).(*IDeCusSystemBtcRefundData)

	return out0, err

}

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DeCusSystem *DeCusSystemSession) GetRefundData() (IDeCusSystemBtcRefundData, error) {
	return _DeCusSystem.Contract.GetRefundData(&_DeCusSystem.CallOpts)
}

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DeCusSystem *DeCusSystemCallerSession) GetRefundData() (IDeCusSystemBtcRefundData, error) {
	return _DeCusSystem.Contract.GetRefundData(&_DeCusSystem.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DeCusSystem *DeCusSystemCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DeCusSystem *DeCusSystemSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DeCusSystem.Contract.GetRoleAdmin(&_DeCusSystem.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DeCusSystem *DeCusSystemCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DeCusSystem.Contract.GetRoleAdmin(&_DeCusSystem.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DeCusSystem *DeCusSystemCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DeCusSystem *DeCusSystemSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DeCusSystem.Contract.GetRoleMember(&_DeCusSystem.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DeCusSystem.Contract.GetRoleMember(&_DeCusSystem.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DeCusSystem *DeCusSystemCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DeCusSystem *DeCusSystemSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DeCusSystem.Contract.GetRoleMemberCount(&_DeCusSystem.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DeCusSystem *DeCusSystemCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DeCusSystem.Contract.GetRoleMemberCount(&_DeCusSystem.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DeCusSystem *DeCusSystemCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DeCusSystem *DeCusSystemSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DeCusSystem.Contract.HasRole(&_DeCusSystem.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DeCusSystem *DeCusSystemCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DeCusSystem.Contract.HasRole(&_DeCusSystem.CallOpts, role, account)
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

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) Rewarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "rewarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DeCusSystem *DeCusSystemSession) Rewarder() (common.Address, error) {
	return _DeCusSystem.Contract.Rewarder(&_DeCusSystem.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) Rewarder() (common.Address, error) {
	return _DeCusSystem.Contract.Rewarder(&_DeCusSystem.CallOpts)
}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DeCusSystem *DeCusSystemCaller) Sats(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeCusSystem.contract.Call(opts, &out, "sats")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DeCusSystem *DeCusSystemSession) Sats() (common.Address, error) {
	return _DeCusSystem.Contract.Sats(&_DeCusSystem.CallOpts)
}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DeCusSystem *DeCusSystemCallerSession) Sats() (common.Address, error) {
	return _DeCusSystem.Contract.Sats(&_DeCusSystem.CallOpts)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemTransactor) AddGroup(opts *bind.TransactOpts, btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "addGroup", btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemSession) AddGroup(btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.AddGroup(&_DeCusSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) AddGroup(btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.AddGroup(&_DeCusSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DeCusSystem *DeCusSystemTransactor) Chill(opts *bind.TransactOpts, keeper common.Address, chillTime uint32) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "chill", keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DeCusSystem *DeCusSystemSession) Chill(keeper common.Address, chillTime uint32) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Chill(&_DeCusSystem.TransactOpts, keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Chill(keeper common.Address, chillTime uint32) (*types.Transaction, error) {
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

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemTransactor) ForceRequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "forceRequestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DeCusSystem.Contract.ForceRequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemTransactorSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DeCusSystem.Contract.ForceRequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.GrantRole(&_DeCusSystem.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.GrantRole(&_DeCusSystem.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DeCusSystem *DeCusSystemTransactor) Initialize(opts *bind.TransactOpts, _sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "initialize", _sats, _registry, _rewarder, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DeCusSystem *DeCusSystemSession) Initialize(_sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Initialize(&_DeCusSystem.TransactOpts, _sats, _registry, _rewarder, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Initialize(_sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.Initialize(&_DeCusSystem.TransactOpts, _sats, _registry, _rewarder, _fee)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DeCusSystem *DeCusSystemTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DeCusSystem *DeCusSystemSession) Pause() (*types.Transaction, error) {
	return _DeCusSystem.Contract.Pause(&_DeCusSystem.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Pause() (*types.Transaction, error) {
	return _DeCusSystem.Contract.Pause(&_DeCusSystem.TransactOpts)
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

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DeCusSystem *DeCusSystemTransactor) RefundBtc(opts *bind.TransactOpts, groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "refundBtc", groupBtcAddress, txId)
}

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DeCusSystem *DeCusSystemSession) RefundBtc(groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RefundBtc(&_DeCusSystem.TransactOpts, groupBtcAddress, txId)
}

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RefundBtc(groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RefundBtc(&_DeCusSystem.TransactOpts, groupBtcAddress, txId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RenounceRole(&_DeCusSystem.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RenounceRole(&_DeCusSystem.TransactOpts, role, account)
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

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemTransactor) RequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "requestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemSession) RequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RequestMint(&_DeCusSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
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

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RevokeRole(&_DeCusSystem.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DeCusSystem.Contract.RevokeRole(&_DeCusSystem.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DeCusSystem *DeCusSystemTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DeCusSystem *DeCusSystemSession) Unpause() (*types.Transaction, error) {
	return _DeCusSystem.Contract.Unpause(&_DeCusSystem.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DeCusSystem *DeCusSystemTransactorSession) Unpause() (*types.Transaction, error) {
	return _DeCusSystem.Contract.Unpause(&_DeCusSystem.TransactOpts)
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

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemTransactor) VerifyMint(opts *bind.TransactOpts, request IDeCusSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.contract.Transact(opts, "verifyMint", request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemSession) VerifyMint(request IDeCusSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyMint(&_DeCusSystem.TransactOpts, request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DeCusSystem *DeCusSystemTransactorSession) VerifyMint(request IDeCusSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DeCusSystem.Contract.VerifyMint(&_DeCusSystem.TransactOpts, request, keepers, r, s, packedV)
}

// DeCusSystemBtcRefundedIterator is returned from FilterBtcRefunded and is used to iterate over the raw logs and unpacked data for BtcRefunded events raised by the DeCusSystem contract.
type DeCusSystemBtcRefundedIterator struct {
	Event *DeCusSystemBtcRefunded // Event containing the contract specifics and raw log

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
func (it *DeCusSystemBtcRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemBtcRefunded)
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
		it.Event = new(DeCusSystemBtcRefunded)
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
func (it *DeCusSystemBtcRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemBtcRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemBtcRefunded represents a BtcRefunded event raised by the DeCusSystem contract.
type DeCusSystemBtcRefunded struct {
	GroupBtcAddress string
	TxId            [32]byte
	ExpiryTimestamp uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBtcRefunded is a free log retrieval operation binding the contract event 0x8588e607efd130a12b8af9ae3c9f09eb4e81b7976778146783cfa8cdc3799594.
//
// Solidity: event BtcRefunded(string groupBtcAddress, bytes32 txId, uint32 expiryTimestamp)
func (_DeCusSystem *DeCusSystemFilterer) FilterBtcRefunded(opts *bind.FilterOpts) (*DeCusSystemBtcRefundedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "BtcRefunded")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemBtcRefundedIterator{contract: _DeCusSystem.contract, event: "BtcRefunded", logs: logs, sub: sub}, nil
}

// WatchBtcRefunded is a free log subscription operation binding the contract event 0x8588e607efd130a12b8af9ae3c9f09eb4e81b7976778146783cfa8cdc3799594.
//
// Solidity: event BtcRefunded(string groupBtcAddress, bytes32 txId, uint32 expiryTimestamp)
func (_DeCusSystem *DeCusSystemFilterer) WatchBtcRefunded(opts *bind.WatchOpts, sink chan<- *DeCusSystemBtcRefunded) (event.Subscription, error) {

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "BtcRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemBtcRefunded)
				if err := _DeCusSystem.contract.UnpackLog(event, "BtcRefunded", log); err != nil {
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

// ParseBtcRefunded is a log parse operation binding the contract event 0x8588e607efd130a12b8af9ae3c9f09eb4e81b7976778146783cfa8cdc3799594.
//
// Solidity: event BtcRefunded(string groupBtcAddress, bytes32 txId, uint32 expiryTimestamp)
func (_DeCusSystem *DeCusSystemFilterer) ParseBtcRefunded(log types.Log) (*DeCusSystemBtcRefunded, error) {
	event := new(DeCusSystemBtcRefunded)
	if err := _DeCusSystem.contract.UnpackLog(event, "BtcRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Recipient       common.Address
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurnRevoked is a free log retrieval operation binding the contract event 0x1870528c09977fe75036f8586518e5853addf6419747cbb81294106f734fe2e1.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address recipient, address operator)
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

// WatchBurnRevoked is a free log subscription operation binding the contract event 0x1870528c09977fe75036f8586518e5853addf6419747cbb81294106f734fe2e1.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address recipient, address operator)
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

// ParseBurnRevoked is a log parse operation binding the contract event 0x1870528c09977fe75036f8586518e5853addf6419747cbb81294106f734fe2e1.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address recipient, address operator)
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
	EndTime uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0xc606d47187d0dea01a7ca1a0d8abf6027f8841bf0f82fab363549489e139e2ce.
//
// Solidity: event Cooldown(address indexed keeper, uint32 endTime)
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

// WatchCooldown is a free log subscription operation binding the contract event 0xc606d47187d0dea01a7ca1a0d8abf6027f8841bf0f82fab363549489e139e2ce.
//
// Solidity: event Cooldown(address indexed keeper, uint32 endTime)
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

// ParseCooldown is a log parse operation binding the contract event 0xc606d47187d0dea01a7ca1a0d8abf6027f8841bf0f82fab363549489e139e2ce.
//
// Solidity: event Cooldown(address indexed keeper, uint32 endTime)
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
	Required   uint32
	MaxSatoshi uint32
	Keepers    []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0xb14273dec0b1f9ed3657f0f047101d6250bbbb76a6c3e7b808dcca46bec0db4c.
//
// Solidity: event GroupAdded(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers)
func (_DeCusSystem *DeCusSystemFilterer) FilterGroupAdded(opts *bind.FilterOpts) (*DeCusSystemGroupAddedIterator, error) {

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return &DeCusSystemGroupAddedIterator{contract: _DeCusSystem.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0xb14273dec0b1f9ed3657f0f047101d6250bbbb76a6c3e7b808dcca46bec0db4c.
//
// Solidity: event GroupAdded(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers)
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

// ParseGroupAdded is a log parse operation binding the contract event 0xb14273dec0b1f9ed3657f0f047101d6250bbbb76a6c3e7b808dcca46bec0db4c.
//
// Solidity: event GroupAdded(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers)
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
	AmountInSatoshi uint32
	GroupBtcAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintRequested is a free log retrieval operation binding the contract event 0x93e399f7caab186cd9c637476e9d46c57d7e68ecb5392570123578f392d5e22c.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint32 amountInSatoshi, string groupBtcAddress)
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

// WatchMintRequested is a free log subscription operation binding the contract event 0x93e399f7caab186cd9c637476e9d46c57d7e68ecb5392570123578f392d5e22c.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint32 amountInSatoshi, string groupBtcAddress)
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

// ParseMintRequested is a log parse operation binding the contract event 0x93e399f7caab186cd9c637476e9d46c57d7e68ecb5392570123578f392d5e22c.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint32 amountInSatoshi, string groupBtcAddress)
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
	BtcTxId         [32]byte
	BtcTxHeight     uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintVerified is a free log retrieval operation binding the contract event 0x19d0f59e71cb6f748028c86f27cb22d05c57315a9c850320ecda2cd13f5a43d5.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers, bytes32 btcTxId, uint32 btcTxHeight)
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

// WatchMintVerified is a free log subscription operation binding the contract event 0x19d0f59e71cb6f748028c86f27cb22d05c57315a9c850320ecda2cd13f5a43d5.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers, bytes32 btcTxId, uint32 btcTxHeight)
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

// ParseMintVerified is a log parse operation binding the contract event 0x19d0f59e71cb6f748028c86f27cb22d05c57315a9c850320ecda2cd13f5a43d5.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers, bytes32 btcTxId, uint32 btcTxHeight)
func (_DeCusSystem *DeCusSystemFilterer) ParseMintVerified(log types.Log) (*DeCusSystemMintVerified, error) {
	event := new(DeCusSystemMintVerified)
	if err := _DeCusSystem.contract.UnpackLog(event, "MintVerified", log); err != nil {
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

// DeCusSystemRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DeCusSystem contract.
type DeCusSystemRoleAdminChangedIterator struct {
	Event *DeCusSystemRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DeCusSystemRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemRoleAdminChanged)
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
		it.Event = new(DeCusSystemRoleAdminChanged)
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
func (it *DeCusSystemRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemRoleAdminChanged represents a RoleAdminChanged event raised by the DeCusSystem contract.
type DeCusSystemRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DeCusSystem *DeCusSystemFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DeCusSystemRoleAdminChangedIterator, error) {

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

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemRoleAdminChangedIterator{contract: _DeCusSystem.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DeCusSystem *DeCusSystemFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DeCusSystemRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemRoleAdminChanged)
				if err := _DeCusSystem.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParseRoleAdminChanged(log types.Log) (*DeCusSystemRoleAdminChanged, error) {
	event := new(DeCusSystemRoleAdminChanged)
	if err := _DeCusSystem.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DeCusSystem contract.
type DeCusSystemRoleGrantedIterator struct {
	Event *DeCusSystemRoleGranted // Event containing the contract specifics and raw log

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
func (it *DeCusSystemRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemRoleGranted)
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
		it.Event = new(DeCusSystemRoleGranted)
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
func (it *DeCusSystemRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemRoleGranted represents a RoleGranted event raised by the DeCusSystem contract.
type DeCusSystemRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DeCusSystem *DeCusSystemFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeCusSystemRoleGrantedIterator, error) {

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

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemRoleGrantedIterator{contract: _DeCusSystem.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DeCusSystem *DeCusSystemFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DeCusSystemRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemRoleGranted)
				if err := _DeCusSystem.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParseRoleGranted(log types.Log) (*DeCusSystemRoleGranted, error) {
	event := new(DeCusSystemRoleGranted)
	if err := _DeCusSystem.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeCusSystemRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DeCusSystem contract.
type DeCusSystemRoleRevokedIterator struct {
	Event *DeCusSystemRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DeCusSystemRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeCusSystemRoleRevoked)
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
		it.Event = new(DeCusSystemRoleRevoked)
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
func (it *DeCusSystemRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeCusSystemRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeCusSystemRoleRevoked represents a RoleRevoked event raised by the DeCusSystem contract.
type DeCusSystemRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DeCusSystem *DeCusSystemFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeCusSystemRoleRevokedIterator, error) {

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

	logs, sub, err := _DeCusSystem.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeCusSystemRoleRevokedIterator{contract: _DeCusSystem.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DeCusSystem *DeCusSystemFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DeCusSystemRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DeCusSystem.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeCusSystemRoleRevoked)
				if err := _DeCusSystem.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DeCusSystem *DeCusSystemFilterer) ParseRoleRevoked(log types.Log) (*DeCusSystemRoleRevoked, error) {
	event := new(DeCusSystemRoleRevoked)
	if err := _DeCusSystem.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
