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

// IDecuxSystemBtcRefundData is an auto generated low-level Go binding around an user-defined struct.
type IDecuxSystemBtcRefundData struct {
	TxId            [32]byte
	GroupBtcAddress string
	ExpiryTimestamp uint32
}

// IDecuxSystemMintRequest is an auto generated low-level Go binding around an user-defined struct.
type IDecuxSystemMintRequest struct {
	ReceiptId [32]byte
	TxId      [32]byte
	Height    uint32
}

// IDecuxSystemReceipt is an auto generated low-level Go binding around an user-defined struct.
type IDecuxSystemReceipt struct {
	GroupBtcAddress    string
	WithdrawBtcAddress string
	TxId               [32]byte
	AmountInSatoshi    uint32
	UpdateTimestamp    uint32
	Height             uint32
	Recipient          common.Address
	Status             uint8
}

// DecuxSystemABI is the input ABI used to generate the binding from.
const DecuxSystemABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"AllowKeeperExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"expiryTimestamp\",\"type\":\"uint32\"}],\"name\":\"BtcRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"BurnVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"GroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"GroupDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"}],\"name\":\"MintRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"MintRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"btcTxId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"btcTxHeight\",\"type\":\"uint32\"}],\"name\":\"MintVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"}],\"name\":\"ToggleExitKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_REUSING_GAP\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_COOLDOWN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_REQUEST_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_GAP\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_VERIFICATION_END\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"}],\"name\":\"addGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowKeeperExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chillTime\",\"type\":\"uint32\"}],\"name\":\"chill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cooldownUntil\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"deleteGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"btcAddresses\",\"type\":\"string[]\"}],\"name\":\"deleteGroups\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"contractISwapFee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"forceRequestMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"required\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"currSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"workingReceiptId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"updateTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIDecuxSystem.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIDecuxSystem.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getReceiptId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRefundData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expiryTimestamp\",\"type\":\"uint32\"}],\"internalType\":\"structIDecuxSystem.BtcRefundData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSATS\",\"name\":\"_sats\",\"type\":\"address\"},{\"internalType\":\"contractIKeeperRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractISwapRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"contractISwapFee\",\"name\":\"_fee\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperExitAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keeperExiting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperRegistry\",\"outputs\":[{\"internalType\":\"contractIKeeperRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"recoverBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"}],\"name\":\"refundBtc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"withdrawBtcAddress\",\"type\":\"string\"}],\"name\":\"requestBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupBtcAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"amountInSatoshi\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"requestMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"revokeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewarder\",\"outputs\":[{\"internalType\":\"contractISwapRewarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sats\",\"outputs\":[{\"internalType\":\"contractSATS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleExitKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"verifyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"}],\"internalType\":\"structIDecuxSystem.MintRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"keepers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"packedV\",\"type\":\"uint256\"}],\"name\":\"verifyMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DecuxSystem is an auto generated Go binding around an Ethereum contract.
type DecuxSystem struct {
	DecuxSystemCaller     // Read-only binding to the contract
	DecuxSystemTransactor // Write-only binding to the contract
	DecuxSystemFilterer   // Log filterer for contract events
}

// DecuxSystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecuxSystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecuxSystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecuxSystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecuxSystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecuxSystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecuxSystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecuxSystemSession struct {
	Contract     *DecuxSystem      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecuxSystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecuxSystemCallerSession struct {
	Contract *DecuxSystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DecuxSystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecuxSystemTransactorSession struct {
	Contract     *DecuxSystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DecuxSystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecuxSystemRaw struct {
	Contract *DecuxSystem // Generic contract binding to access the raw methods on
}

// DecuxSystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecuxSystemCallerRaw struct {
	Contract *DecuxSystemCaller // Generic read-only contract binding to access the raw methods on
}

// DecuxSystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecuxSystemTransactorRaw struct {
	Contract *DecuxSystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecuxSystem creates a new instance of DecuxSystem, bound to a specific deployed contract.
func NewDecuxSystem(address common.Address, backend bind.ContractBackend) (*DecuxSystem, error) {
	contract, err := bindDecuxSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecuxSystem{DecuxSystemCaller: DecuxSystemCaller{contract: contract}, DecuxSystemTransactor: DecuxSystemTransactor{contract: contract}, DecuxSystemFilterer: DecuxSystemFilterer{contract: contract}}, nil
}

// NewDecuxSystemCaller creates a new read-only instance of DecuxSystem, bound to a specific deployed contract.
func NewDecuxSystemCaller(address common.Address, caller bind.ContractCaller) (*DecuxSystemCaller, error) {
	contract, err := bindDecuxSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemCaller{contract: contract}, nil
}

// NewDecuxSystemTransactor creates a new write-only instance of DecuxSystem, bound to a specific deployed contract.
func NewDecuxSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*DecuxSystemTransactor, error) {
	contract, err := bindDecuxSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemTransactor{contract: contract}, nil
}

// NewDecuxSystemFilterer creates a new log filterer instance of DecuxSystem, bound to a specific deployed contract.
func NewDecuxSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*DecuxSystemFilterer, error) {
	contract, err := bindDecuxSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemFilterer{contract: contract}, nil
}

// bindDecuxSystem binds a generic wrapper to an already deployed contract.
func bindDecuxSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DecuxSystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecuxSystem *DecuxSystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecuxSystem.Contract.DecuxSystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecuxSystem *DecuxSystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DecuxSystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecuxSystem *DecuxSystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DecuxSystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecuxSystem *DecuxSystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecuxSystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecuxSystem *DecuxSystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecuxSystem *DecuxSystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecuxSystem.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.DEFAULTADMINROLE(&_DecuxSystem.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.DEFAULTADMINROLE(&_DecuxSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) GROUPREUSINGGAP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "GROUP_REUSING_GAP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) GROUPREUSINGGAP() (uint32, error) {
	return _DecuxSystem.Contract.GROUPREUSINGGAP(&_DecuxSystem.CallOpts)
}

// GROUPREUSINGGAP is a free data retrieval call binding the contract method 0x6e157869.
//
// Solidity: function GROUP_REUSING_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) GROUPREUSINGGAP() (uint32, error) {
	return _DecuxSystem.Contract.GROUPREUSINGGAP(&_DecuxSystem.CallOpts)
}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCaller) GROUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "GROUP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemSession) GROUPROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.GROUPROLE(&_DecuxSystem.CallOpts)
}

// GROUPROLE is a free data retrieval call binding the contract method 0xc812a215.
//
// Solidity: function GROUP_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCallerSession) GROUPROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.GROUPROLE(&_DecuxSystem.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCaller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemSession) GUARDROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.GUARDROLE(&_DecuxSystem.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_DecuxSystem *DecuxSystemCallerSession) GUARDROLE() ([32]byte, error) {
	return _DecuxSystem.Contract.GUARDROLE(&_DecuxSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) KEEPERCOOLDOWN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "KEEPER_COOLDOWN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) KEEPERCOOLDOWN() (uint32, error) {
	return _DecuxSystem.Contract.KEEPERCOOLDOWN(&_DecuxSystem.CallOpts)
}

// KEEPERCOOLDOWN is a free data retrieval call binding the contract method 0x7f7bcb62.
//
// Solidity: function KEEPER_COOLDOWN() view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) KEEPERCOOLDOWN() (uint32, error) {
	return _DecuxSystem.Contract.KEEPERCOOLDOWN(&_DecuxSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) MINTREQUESTGRACEPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "MINT_REQUEST_GRACE_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) MINTREQUESTGRACEPERIOD() (uint32, error) {
	return _DecuxSystem.Contract.MINTREQUESTGRACEPERIOD(&_DecuxSystem.CallOpts)
}

// MINTREQUESTGRACEPERIOD is a free data retrieval call binding the contract method 0x48a4a489.
//
// Solidity: function MINT_REQUEST_GRACE_PERIOD() view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) MINTREQUESTGRACEPERIOD() (uint32, error) {
	return _DecuxSystem.Contract.MINTREQUESTGRACEPERIOD(&_DecuxSystem.CallOpts)
}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) REFUNDGAP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "REFUND_GAP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) REFUNDGAP() (uint32, error) {
	return _DecuxSystem.Contract.REFUNDGAP(&_DecuxSystem.CallOpts)
}

// REFUNDGAP is a free data retrieval call binding the contract method 0x833eeb57.
//
// Solidity: function REFUND_GAP() view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) REFUNDGAP() (uint32, error) {
	return _DecuxSystem.Contract.REFUNDGAP(&_DecuxSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) WITHDRAWVERIFICATIONEND(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "WITHDRAW_VERIFICATION_END")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) WITHDRAWVERIFICATIONEND() (uint32, error) {
	return _DecuxSystem.Contract.WITHDRAWVERIFICATIONEND(&_DecuxSystem.CallOpts)
}

// WITHDRAWVERIFICATIONEND is a free data retrieval call binding the contract method 0x6557f366.
//
// Solidity: function WITHDRAW_VERIFICATION_END() view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) WITHDRAWVERIFICATIONEND() (uint32, error) {
	return _DecuxSystem.Contract.WITHDRAWVERIFICATIONEND(&_DecuxSystem.CallOpts)
}

// CooldownUntil is a free data retrieval call binding the contract method 0x25d40e3c.
//
// Solidity: function cooldownUntil(address ) view returns(uint32)
func (_DecuxSystem *DecuxSystemCaller) CooldownUntil(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "cooldownUntil", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CooldownUntil is a free data retrieval call binding the contract method 0x25d40e3c.
//
// Solidity: function cooldownUntil(address ) view returns(uint32)
func (_DecuxSystem *DecuxSystemSession) CooldownUntil(arg0 common.Address) (uint32, error) {
	return _DecuxSystem.Contract.CooldownUntil(&_DecuxSystem.CallOpts, arg0)
}

// CooldownUntil is a free data retrieval call binding the contract method 0x25d40e3c.
//
// Solidity: function cooldownUntil(address ) view returns(uint32)
func (_DecuxSystem *DecuxSystemCallerSession) CooldownUntil(arg0 common.Address) (uint32, error) {
	return _DecuxSystem.Contract.CooldownUntil(&_DecuxSystem.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DecuxSystem *DecuxSystemCaller) Fee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DecuxSystem *DecuxSystemSession) Fee() (common.Address, error) {
	return _DecuxSystem.Contract.Fee(&_DecuxSystem.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(address)
func (_DecuxSystem *DecuxSystemCallerSession) Fee() (common.Address, error) {
	return _DecuxSystem.Contract.Fee(&_DecuxSystem.CallOpts)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint32 required, uint32 maxSatoshi, uint32 currSatoshi, uint32 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DecuxSystem *DecuxSystemCaller) GetGroup(opts *bind.CallOpts, btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getGroup", btcAddress)

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
func (_DecuxSystem *DecuxSystemSession) GetGroup(btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DecuxSystem.Contract.GetGroup(&_DecuxSystem.CallOpts, btcAddress)
}

// GetGroup is a free data retrieval call binding the contract method 0xabef281e.
//
// Solidity: function getGroup(string btcAddress) view returns(uint32 required, uint32 maxSatoshi, uint32 currSatoshi, uint32 nonce, address[] keepers, bytes32 workingReceiptId)
func (_DecuxSystem *DecuxSystemCallerSession) GetGroup(btcAddress string) (struct {
	Required         uint32
	MaxSatoshi       uint32
	CurrSatoshi      uint32
	Nonce            uint32
	Keepers          []common.Address
	WorkingReceiptId [32]byte
}, error) {
	return _DecuxSystem.Contract.GetGroup(&_DecuxSystem.CallOpts, btcAddress)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
func (_DecuxSystem *DecuxSystemCaller) GetReceipt(opts *bind.CallOpts, receiptId [32]byte) (IDecuxSystemReceipt, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getReceipt", receiptId)

	if err != nil {
		return *new(IDecuxSystemReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(IDecuxSystemReceipt)).(*IDecuxSystemReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
func (_DecuxSystem *DecuxSystemSession) GetReceipt(receiptId [32]byte) (IDecuxSystemReceipt, error) {
	return _DecuxSystem.Contract.GetReceipt(&_DecuxSystem.CallOpts, receiptId)
}

// GetReceipt is a free data retrieval call binding the contract method 0xfcecbb61.
//
// Solidity: function getReceipt(bytes32 receiptId) view returns((string,string,bytes32,uint32,uint32,uint32,address,uint8))
func (_DecuxSystem *DecuxSystemCallerSession) GetReceipt(receiptId [32]byte) (IDecuxSystemReceipt, error) {
	return _DecuxSystem.Contract.GetReceipt(&_DecuxSystem.CallOpts, receiptId)
}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DecuxSystem *DecuxSystemCaller) GetReceiptId(opts *bind.CallOpts, groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getReceiptId", groupBtcAddress, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DecuxSystem *DecuxSystemSession) GetReceiptId(groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	return _DecuxSystem.Contract.GetReceiptId(&_DecuxSystem.CallOpts, groupBtcAddress, nonce)
}

// GetReceiptId is a free data retrieval call binding the contract method 0xa2ad4d9b.
//
// Solidity: function getReceiptId(string groupBtcAddress, uint256 nonce) pure returns(bytes32)
func (_DecuxSystem *DecuxSystemCallerSession) GetReceiptId(groupBtcAddress string, nonce *big.Int) ([32]byte, error) {
	return _DecuxSystem.Contract.GetReceiptId(&_DecuxSystem.CallOpts, groupBtcAddress, nonce)
}

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DecuxSystem *DecuxSystemCaller) GetRefundData(opts *bind.CallOpts) (IDecuxSystemBtcRefundData, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getRefundData")

	if err != nil {
		return *new(IDecuxSystemBtcRefundData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDecuxSystemBtcRefundData)).(*IDecuxSystemBtcRefundData)

	return out0, err

}

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DecuxSystem *DecuxSystemSession) GetRefundData() (IDecuxSystemBtcRefundData, error) {
	return _DecuxSystem.Contract.GetRefundData(&_DecuxSystem.CallOpts)
}

// GetRefundData is a free data retrieval call binding the contract method 0x38f39b48.
//
// Solidity: function getRefundData() view returns((bytes32,string,uint32))
func (_DecuxSystem *DecuxSystemCallerSession) GetRefundData() (IDecuxSystemBtcRefundData, error) {
	return _DecuxSystem.Contract.GetRefundData(&_DecuxSystem.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DecuxSystem *DecuxSystemCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DecuxSystem *DecuxSystemSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DecuxSystem.Contract.GetRoleAdmin(&_DecuxSystem.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DecuxSystem *DecuxSystemCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DecuxSystem.Contract.GetRoleAdmin(&_DecuxSystem.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DecuxSystem *DecuxSystemCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DecuxSystem *DecuxSystemSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DecuxSystem.Contract.GetRoleMember(&_DecuxSystem.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DecuxSystem *DecuxSystemCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DecuxSystem.Contract.GetRoleMember(&_DecuxSystem.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DecuxSystem *DecuxSystemCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DecuxSystem *DecuxSystemSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DecuxSystem.Contract.GetRoleMemberCount(&_DecuxSystem.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DecuxSystem *DecuxSystemCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DecuxSystem.Contract.GetRoleMemberCount(&_DecuxSystem.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DecuxSystem *DecuxSystemCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DecuxSystem *DecuxSystemSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DecuxSystem.Contract.HasRole(&_DecuxSystem.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DecuxSystem *DecuxSystemCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DecuxSystem.Contract.HasRole(&_DecuxSystem.CallOpts, role, account)
}

// KeeperExitAllowed is a free data retrieval call binding the contract method 0x7188ac00.
//
// Solidity: function keeperExitAllowed() view returns(bool)
func (_DecuxSystem *DecuxSystemCaller) KeeperExitAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "keeperExitAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeeperExitAllowed is a free data retrieval call binding the contract method 0x7188ac00.
//
// Solidity: function keeperExitAllowed() view returns(bool)
func (_DecuxSystem *DecuxSystemSession) KeeperExitAllowed() (bool, error) {
	return _DecuxSystem.Contract.KeeperExitAllowed(&_DecuxSystem.CallOpts)
}

// KeeperExitAllowed is a free data retrieval call binding the contract method 0x7188ac00.
//
// Solidity: function keeperExitAllowed() view returns(bool)
func (_DecuxSystem *DecuxSystemCallerSession) KeeperExitAllowed() (bool, error) {
	return _DecuxSystem.Contract.KeeperExitAllowed(&_DecuxSystem.CallOpts)
}

// KeeperExiting is a free data retrieval call binding the contract method 0x2f7f636a.
//
// Solidity: function keeperExiting(address ) view returns(bool)
func (_DecuxSystem *DecuxSystemCaller) KeeperExiting(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "keeperExiting", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeeperExiting is a free data retrieval call binding the contract method 0x2f7f636a.
//
// Solidity: function keeperExiting(address ) view returns(bool)
func (_DecuxSystem *DecuxSystemSession) KeeperExiting(arg0 common.Address) (bool, error) {
	return _DecuxSystem.Contract.KeeperExiting(&_DecuxSystem.CallOpts, arg0)
}

// KeeperExiting is a free data retrieval call binding the contract method 0x2f7f636a.
//
// Solidity: function keeperExiting(address ) view returns(bool)
func (_DecuxSystem *DecuxSystemCallerSession) KeeperExiting(arg0 common.Address) (bool, error) {
	return _DecuxSystem.Contract.KeeperExiting(&_DecuxSystem.CallOpts, arg0)
}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DecuxSystem *DecuxSystemCaller) KeeperRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "keeperRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DecuxSystem *DecuxSystemSession) KeeperRegistry() (common.Address, error) {
	return _DecuxSystem.Contract.KeeperRegistry(&_DecuxSystem.CallOpts)
}

// KeeperRegistry is a free data retrieval call binding the contract method 0x83e22774.
//
// Solidity: function keeperRegistry() view returns(address)
func (_DecuxSystem *DecuxSystemCallerSession) KeeperRegistry() (common.Address, error) {
	return _DecuxSystem.Contract.KeeperRegistry(&_DecuxSystem.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DecuxSystem *DecuxSystemCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DecuxSystem *DecuxSystemSession) Paused() (bool, error) {
	return _DecuxSystem.Contract.Paused(&_DecuxSystem.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DecuxSystem *DecuxSystemCallerSession) Paused() (bool, error) {
	return _DecuxSystem.Contract.Paused(&_DecuxSystem.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DecuxSystem *DecuxSystemCaller) Rewarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "rewarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DecuxSystem *DecuxSystemSession) Rewarder() (common.Address, error) {
	return _DecuxSystem.Contract.Rewarder(&_DecuxSystem.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xdcc3e06e.
//
// Solidity: function rewarder() view returns(address)
func (_DecuxSystem *DecuxSystemCallerSession) Rewarder() (common.Address, error) {
	return _DecuxSystem.Contract.Rewarder(&_DecuxSystem.CallOpts)
}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DecuxSystem *DecuxSystemCaller) Sats(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DecuxSystem.contract.Call(opts, &out, "sats")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DecuxSystem *DecuxSystemSession) Sats() (common.Address, error) {
	return _DecuxSystem.Contract.Sats(&_DecuxSystem.CallOpts)
}

// Sats is a free data retrieval call binding the contract method 0x7a5c8f47.
//
// Solidity: function sats() view returns(address)
func (_DecuxSystem *DecuxSystemCallerSession) Sats() (common.Address, error) {
	return _DecuxSystem.Contract.Sats(&_DecuxSystem.CallOpts)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DecuxSystem *DecuxSystemTransactor) AddGroup(opts *bind.TransactOpts, btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "addGroup", btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DecuxSystem *DecuxSystemSession) AddGroup(btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.AddGroup(&_DecuxSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// AddGroup is a paid mutator transaction binding the contract method 0xb94cd07f.
//
// Solidity: function addGroup(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) AddGroup(btcAddress string, required uint32, maxSatoshi uint32, keepers []common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.AddGroup(&_DecuxSystem.TransactOpts, btcAddress, required, maxSatoshi, keepers)
}

// AllowKeeperExit is a paid mutator transaction binding the contract method 0xa965476d.
//
// Solidity: function allowKeeperExit() returns()
func (_DecuxSystem *DecuxSystemTransactor) AllowKeeperExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "allowKeeperExit")
}

// AllowKeeperExit is a paid mutator transaction binding the contract method 0xa965476d.
//
// Solidity: function allowKeeperExit() returns()
func (_DecuxSystem *DecuxSystemSession) AllowKeeperExit() (*types.Transaction, error) {
	return _DecuxSystem.Contract.AllowKeeperExit(&_DecuxSystem.TransactOpts)
}

// AllowKeeperExit is a paid mutator transaction binding the contract method 0xa965476d.
//
// Solidity: function allowKeeperExit() returns()
func (_DecuxSystem *DecuxSystemTransactorSession) AllowKeeperExit() (*types.Transaction, error) {
	return _DecuxSystem.Contract.AllowKeeperExit(&_DecuxSystem.TransactOpts)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DecuxSystem *DecuxSystemTransactor) Chill(opts *bind.TransactOpts, keeper common.Address, chillTime uint32) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "chill", keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DecuxSystem *DecuxSystemSession) Chill(keeper common.Address, chillTime uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.Chill(&_DecuxSystem.TransactOpts, keeper, chillTime)
}

// Chill is a paid mutator transaction binding the contract method 0x3d701d00.
//
// Solidity: function chill(address keeper, uint32 chillTime) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) Chill(keeper common.Address, chillTime uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.Chill(&_DecuxSystem.TransactOpts, keeper, chillTime)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DecuxSystem *DecuxSystemTransactor) DeleteGroup(opts *bind.TransactOpts, btcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "deleteGroup", btcAddress)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DecuxSystem *DecuxSystemSession) DeleteGroup(btcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DeleteGroup(&_DecuxSystem.TransactOpts, btcAddress)
}

// DeleteGroup is a paid mutator transaction binding the contract method 0x2e8b92a9.
//
// Solidity: function deleteGroup(string btcAddress) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) DeleteGroup(btcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DeleteGroup(&_DecuxSystem.TransactOpts, btcAddress)
}

// DeleteGroups is a paid mutator transaction binding the contract method 0x5cabd667.
//
// Solidity: function deleteGroups(string[] btcAddresses) returns()
func (_DecuxSystem *DecuxSystemTransactor) DeleteGroups(opts *bind.TransactOpts, btcAddresses []string) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "deleteGroups", btcAddresses)
}

// DeleteGroups is a paid mutator transaction binding the contract method 0x5cabd667.
//
// Solidity: function deleteGroups(string[] btcAddresses) returns()
func (_DecuxSystem *DecuxSystemSession) DeleteGroups(btcAddresses []string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DeleteGroups(&_DecuxSystem.TransactOpts, btcAddresses)
}

// DeleteGroups is a paid mutator transaction binding the contract method 0x5cabd667.
//
// Solidity: function deleteGroups(string[] btcAddresses) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) DeleteGroups(btcAddresses []string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.DeleteGroups(&_DecuxSystem.TransactOpts, btcAddresses)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemTransactor) ForceRequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "forceRequestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.ForceRequestMint(&_DecuxSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// ForceRequestMint is a paid mutator transaction binding the contract method 0x72df0e38.
//
// Solidity: function forceRequestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemTransactorSession) ForceRequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.ForceRequestMint(&_DecuxSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.GrantRole(&_DecuxSystem.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.GrantRole(&_DecuxSystem.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DecuxSystem *DecuxSystemTransactor) Initialize(opts *bind.TransactOpts, _sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "initialize", _sats, _registry, _rewarder, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DecuxSystem *DecuxSystemSession) Initialize(_sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.Initialize(&_DecuxSystem.TransactOpts, _sats, _registry, _rewarder, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _sats, address _registry, address _rewarder, address _fee) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) Initialize(_sats common.Address, _registry common.Address, _rewarder common.Address, _fee common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.Initialize(&_DecuxSystem.TransactOpts, _sats, _registry, _rewarder, _fee)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DecuxSystem *DecuxSystemTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DecuxSystem *DecuxSystemSession) Pause() (*types.Transaction, error) {
	return _DecuxSystem.Contract.Pause(&_DecuxSystem.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DecuxSystem *DecuxSystemTransactorSession) Pause() (*types.Transaction, error) {
	return _DecuxSystem.Contract.Pause(&_DecuxSystem.TransactOpts)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactor) RecoverBurn(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "recoverBurn", receiptId)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemSession) RecoverBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RecoverBurn(&_DecuxSystem.TransactOpts, receiptId)
}

// RecoverBurn is a paid mutator transaction binding the contract method 0xe8c86b03.
//
// Solidity: function recoverBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RecoverBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RecoverBurn(&_DecuxSystem.TransactOpts, receiptId)
}

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DecuxSystem *DecuxSystemTransactor) RefundBtc(opts *bind.TransactOpts, groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "refundBtc", groupBtcAddress, txId)
}

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DecuxSystem *DecuxSystemSession) RefundBtc(groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RefundBtc(&_DecuxSystem.TransactOpts, groupBtcAddress, txId)
}

// RefundBtc is a paid mutator transaction binding the contract method 0x9fd72189.
//
// Solidity: function refundBtc(string groupBtcAddress, bytes32 txId) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RefundBtc(groupBtcAddress string, txId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RefundBtc(&_DecuxSystem.TransactOpts, groupBtcAddress, txId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RenounceRole(&_DecuxSystem.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RenounceRole(&_DecuxSystem.TransactOpts, role, account)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DecuxSystem *DecuxSystemTransactor) RequestBurn(opts *bind.TransactOpts, receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "requestBurn", receiptId, withdrawBtcAddress)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DecuxSystem *DecuxSystemSession) RequestBurn(receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RequestBurn(&_DecuxSystem.TransactOpts, receiptId, withdrawBtcAddress)
}

// RequestBurn is a paid mutator transaction binding the contract method 0x5f890657.
//
// Solidity: function requestBurn(bytes32 receiptId, string withdrawBtcAddress) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RequestBurn(receiptId [32]byte, withdrawBtcAddress string) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RequestBurn(&_DecuxSystem.TransactOpts, receiptId, withdrawBtcAddress)
}

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemTransactor) RequestMint(opts *bind.TransactOpts, groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "requestMint", groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemSession) RequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RequestMint(&_DecuxSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// RequestMint is a paid mutator transaction binding the contract method 0xaa57665b.
//
// Solidity: function requestMint(string groupBtcAddress, uint32 amountInSatoshi, uint32 nonce) payable returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RequestMint(groupBtcAddress string, amountInSatoshi uint32, nonce uint32) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RequestMint(&_DecuxSystem.TransactOpts, groupBtcAddress, amountInSatoshi, nonce)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactor) RevokeMint(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "revokeMint", receiptId)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemSession) RevokeMint(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RevokeMint(&_DecuxSystem.TransactOpts, receiptId)
}

// RevokeMint is a paid mutator transaction binding the contract method 0x507d87e4.
//
// Solidity: function revokeMint(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RevokeMint(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RevokeMint(&_DecuxSystem.TransactOpts, receiptId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RevokeRole(&_DecuxSystem.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DecuxSystem.Contract.RevokeRole(&_DecuxSystem.TransactOpts, role, account)
}

// ToggleExitKeeper is a paid mutator transaction binding the contract method 0x1aa8517c.
//
// Solidity: function toggleExitKeeper() returns()
func (_DecuxSystem *DecuxSystemTransactor) ToggleExitKeeper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "toggleExitKeeper")
}

// ToggleExitKeeper is a paid mutator transaction binding the contract method 0x1aa8517c.
//
// Solidity: function toggleExitKeeper() returns()
func (_DecuxSystem *DecuxSystemSession) ToggleExitKeeper() (*types.Transaction, error) {
	return _DecuxSystem.Contract.ToggleExitKeeper(&_DecuxSystem.TransactOpts)
}

// ToggleExitKeeper is a paid mutator transaction binding the contract method 0x1aa8517c.
//
// Solidity: function toggleExitKeeper() returns()
func (_DecuxSystem *DecuxSystemTransactorSession) ToggleExitKeeper() (*types.Transaction, error) {
	return _DecuxSystem.Contract.ToggleExitKeeper(&_DecuxSystem.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DecuxSystem *DecuxSystemTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DecuxSystem *DecuxSystemSession) Unpause() (*types.Transaction, error) {
	return _DecuxSystem.Contract.Unpause(&_DecuxSystem.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DecuxSystem *DecuxSystemTransactorSession) Unpause() (*types.Transaction, error) {
	return _DecuxSystem.Contract.Unpause(&_DecuxSystem.TransactOpts)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactor) VerifyBurn(opts *bind.TransactOpts, receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "verifyBurn", receiptId)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemSession) VerifyBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.VerifyBurn(&_DecuxSystem.TransactOpts, receiptId)
}

// VerifyBurn is a paid mutator transaction binding the contract method 0x372fedf1.
//
// Solidity: function verifyBurn(bytes32 receiptId) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) VerifyBurn(receiptId [32]byte) (*types.Transaction, error) {
	return _DecuxSystem.Contract.VerifyBurn(&_DecuxSystem.TransactOpts, receiptId)
}

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DecuxSystem *DecuxSystemTransactor) VerifyMint(opts *bind.TransactOpts, request IDecuxSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DecuxSystem.contract.Transact(opts, "verifyMint", request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DecuxSystem *DecuxSystemSession) VerifyMint(request IDecuxSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DecuxSystem.Contract.VerifyMint(&_DecuxSystem.TransactOpts, request, keepers, r, s, packedV)
}

// VerifyMint is a paid mutator transaction binding the contract method 0x4327c083.
//
// Solidity: function verifyMint((bytes32,bytes32,uint32) request, address[] keepers, bytes32[] r, bytes32[] s, uint256 packedV) returns()
func (_DecuxSystem *DecuxSystemTransactorSession) VerifyMint(request IDecuxSystemMintRequest, keepers []common.Address, r [][32]byte, s [][32]byte, packedV *big.Int) (*types.Transaction, error) {
	return _DecuxSystem.Contract.VerifyMint(&_DecuxSystem.TransactOpts, request, keepers, r, s, packedV)
}

// DecuxSystemAllowKeeperExitIterator is returned from FilterAllowKeeperExit and is used to iterate over the raw logs and unpacked data for AllowKeeperExit events raised by the DecuxSystem contract.
type DecuxSystemAllowKeeperExitIterator struct {
	Event *DecuxSystemAllowKeeperExit // Event containing the contract specifics and raw log

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
func (it *DecuxSystemAllowKeeperExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemAllowKeeperExit)
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
		it.Event = new(DecuxSystemAllowKeeperExit)
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
func (it *DecuxSystemAllowKeeperExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemAllowKeeperExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemAllowKeeperExit represents a AllowKeeperExit event raised by the DecuxSystem contract.
type DecuxSystemAllowKeeperExit struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowKeeperExit is a free log retrieval operation binding the contract event 0x0ff3bf87a808983b14400cade429c3e8e3ee8459a32a33de6f3f8b1cd2845dea.
//
// Solidity: event AllowKeeperExit(address indexed operator)
func (_DecuxSystem *DecuxSystemFilterer) FilterAllowKeeperExit(opts *bind.FilterOpts, operator []common.Address) (*DecuxSystemAllowKeeperExitIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "AllowKeeperExit", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemAllowKeeperExitIterator{contract: _DecuxSystem.contract, event: "AllowKeeperExit", logs: logs, sub: sub}, nil
}

// WatchAllowKeeperExit is a free log subscription operation binding the contract event 0x0ff3bf87a808983b14400cade429c3e8e3ee8459a32a33de6f3f8b1cd2845dea.
//
// Solidity: event AllowKeeperExit(address indexed operator)
func (_DecuxSystem *DecuxSystemFilterer) WatchAllowKeeperExit(opts *bind.WatchOpts, sink chan<- *DecuxSystemAllowKeeperExit, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "AllowKeeperExit", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemAllowKeeperExit)
				if err := _DecuxSystem.contract.UnpackLog(event, "AllowKeeperExit", log); err != nil {
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

// ParseAllowKeeperExit is a log parse operation binding the contract event 0x0ff3bf87a808983b14400cade429c3e8e3ee8459a32a33de6f3f8b1cd2845dea.
//
// Solidity: event AllowKeeperExit(address indexed operator)
func (_DecuxSystem *DecuxSystemFilterer) ParseAllowKeeperExit(log types.Log) (*DecuxSystemAllowKeeperExit, error) {
	event := new(DecuxSystemAllowKeeperExit)
	if err := _DecuxSystem.contract.UnpackLog(event, "AllowKeeperExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemBtcRefundedIterator is returned from FilterBtcRefunded and is used to iterate over the raw logs and unpacked data for BtcRefunded events raised by the DecuxSystem contract.
type DecuxSystemBtcRefundedIterator struct {
	Event *DecuxSystemBtcRefunded // Event containing the contract specifics and raw log

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
func (it *DecuxSystemBtcRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemBtcRefunded)
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
		it.Event = new(DecuxSystemBtcRefunded)
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
func (it *DecuxSystemBtcRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemBtcRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemBtcRefunded represents a BtcRefunded event raised by the DecuxSystem contract.
type DecuxSystemBtcRefunded struct {
	GroupBtcAddress string
	TxId            [32]byte
	ExpiryTimestamp uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBtcRefunded is a free log retrieval operation binding the contract event 0x8588e607efd130a12b8af9ae3c9f09eb4e81b7976778146783cfa8cdc3799594.
//
// Solidity: event BtcRefunded(string groupBtcAddress, bytes32 txId, uint32 expiryTimestamp)
func (_DecuxSystem *DecuxSystemFilterer) FilterBtcRefunded(opts *bind.FilterOpts) (*DecuxSystemBtcRefundedIterator, error) {

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "BtcRefunded")
	if err != nil {
		return nil, err
	}
	return &DecuxSystemBtcRefundedIterator{contract: _DecuxSystem.contract, event: "BtcRefunded", logs: logs, sub: sub}, nil
}

// WatchBtcRefunded is a free log subscription operation binding the contract event 0x8588e607efd130a12b8af9ae3c9f09eb4e81b7976778146783cfa8cdc3799594.
//
// Solidity: event BtcRefunded(string groupBtcAddress, bytes32 txId, uint32 expiryTimestamp)
func (_DecuxSystem *DecuxSystemFilterer) WatchBtcRefunded(opts *bind.WatchOpts, sink chan<- *DecuxSystemBtcRefunded) (event.Subscription, error) {

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "BtcRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemBtcRefunded)
				if err := _DecuxSystem.contract.UnpackLog(event, "BtcRefunded", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseBtcRefunded(log types.Log) (*DecuxSystemBtcRefunded, error) {
	event := new(DecuxSystemBtcRefunded)
	if err := _DecuxSystem.contract.UnpackLog(event, "BtcRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemBurnRequestedIterator is returned from FilterBurnRequested and is used to iterate over the raw logs and unpacked data for BurnRequested events raised by the DecuxSystem contract.
type DecuxSystemBurnRequestedIterator struct {
	Event *DecuxSystemBurnRequested // Event containing the contract specifics and raw log

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
func (it *DecuxSystemBurnRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemBurnRequested)
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
		it.Event = new(DecuxSystemBurnRequested)
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
func (it *DecuxSystemBurnRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemBurnRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemBurnRequested represents a BurnRequested event raised by the DecuxSystem contract.
type DecuxSystemBurnRequested struct {
	ReceiptId          [32]byte
	GroupBtcAddress    string
	WithdrawBtcAddress string
	Operator           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBurnRequested is a free log retrieval operation binding the contract event 0xf0e638df6d296aaddeb18409852c4fa659a87b113a1fdb9cf7535668bd49497d.
//
// Solidity: event BurnRequested(bytes32 indexed receiptId, string groupBtcAddress, string withdrawBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) FilterBurnRequested(opts *bind.FilterOpts, receiptId [][32]byte) (*DecuxSystemBurnRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "BurnRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemBurnRequestedIterator{contract: _DecuxSystem.contract, event: "BurnRequested", logs: logs, sub: sub}, nil
}

// WatchBurnRequested is a free log subscription operation binding the contract event 0xf0e638df6d296aaddeb18409852c4fa659a87b113a1fdb9cf7535668bd49497d.
//
// Solidity: event BurnRequested(bytes32 indexed receiptId, string groupBtcAddress, string withdrawBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) WatchBurnRequested(opts *bind.WatchOpts, sink chan<- *DecuxSystemBurnRequested, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "BurnRequested", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemBurnRequested)
				if err := _DecuxSystem.contract.UnpackLog(event, "BurnRequested", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseBurnRequested(log types.Log) (*DecuxSystemBurnRequested, error) {
	event := new(DecuxSystemBurnRequested)
	if err := _DecuxSystem.contract.UnpackLog(event, "BurnRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemBurnRevokedIterator is returned from FilterBurnRevoked and is used to iterate over the raw logs and unpacked data for BurnRevoked events raised by the DecuxSystem contract.
type DecuxSystemBurnRevokedIterator struct {
	Event *DecuxSystemBurnRevoked // Event containing the contract specifics and raw log

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
func (it *DecuxSystemBurnRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemBurnRevoked)
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
		it.Event = new(DecuxSystemBurnRevoked)
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
func (it *DecuxSystemBurnRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemBurnRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemBurnRevoked represents a BurnRevoked event raised by the DecuxSystem contract.
type DecuxSystemBurnRevoked struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Recipient       common.Address
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurnRevoked is a free log retrieval operation binding the contract event 0x1870528c09977fe75036f8586518e5853addf6419747cbb81294106f734fe2e1.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address recipient, address operator)
func (_DecuxSystem *DecuxSystemFilterer) FilterBurnRevoked(opts *bind.FilterOpts, receiptId [][32]byte) (*DecuxSystemBurnRevokedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "BurnRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemBurnRevokedIterator{contract: _DecuxSystem.contract, event: "BurnRevoked", logs: logs, sub: sub}, nil
}

// WatchBurnRevoked is a free log subscription operation binding the contract event 0x1870528c09977fe75036f8586518e5853addf6419747cbb81294106f734fe2e1.
//
// Solidity: event BurnRevoked(bytes32 indexed receiptId, string groupBtcAddress, address recipient, address operator)
func (_DecuxSystem *DecuxSystemFilterer) WatchBurnRevoked(opts *bind.WatchOpts, sink chan<- *DecuxSystemBurnRevoked, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "BurnRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemBurnRevoked)
				if err := _DecuxSystem.contract.UnpackLog(event, "BurnRevoked", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseBurnRevoked(log types.Log) (*DecuxSystemBurnRevoked, error) {
	event := new(DecuxSystemBurnRevoked)
	if err := _DecuxSystem.contract.UnpackLog(event, "BurnRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemBurnVerifiedIterator is returned from FilterBurnVerified and is used to iterate over the raw logs and unpacked data for BurnVerified events raised by the DecuxSystem contract.
type DecuxSystemBurnVerifiedIterator struct {
	Event *DecuxSystemBurnVerified // Event containing the contract specifics and raw log

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
func (it *DecuxSystemBurnVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemBurnVerified)
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
		it.Event = new(DecuxSystemBurnVerified)
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
func (it *DecuxSystemBurnVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemBurnVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemBurnVerified represents a BurnVerified event raised by the DecuxSystem contract.
type DecuxSystemBurnVerified struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurnVerified is a free log retrieval operation binding the contract event 0xd90232357c1ad1e207410849ce0a0374bac4c22771453838f285ceca745547e3.
//
// Solidity: event BurnVerified(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) FilterBurnVerified(opts *bind.FilterOpts, receiptId [][32]byte) (*DecuxSystemBurnVerifiedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "BurnVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemBurnVerifiedIterator{contract: _DecuxSystem.contract, event: "BurnVerified", logs: logs, sub: sub}, nil
}

// WatchBurnVerified is a free log subscription operation binding the contract event 0xd90232357c1ad1e207410849ce0a0374bac4c22771453838f285ceca745547e3.
//
// Solidity: event BurnVerified(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) WatchBurnVerified(opts *bind.WatchOpts, sink chan<- *DecuxSystemBurnVerified, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "BurnVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemBurnVerified)
				if err := _DecuxSystem.contract.UnpackLog(event, "BurnVerified", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseBurnVerified(log types.Log) (*DecuxSystemBurnVerified, error) {
	event := new(DecuxSystemBurnVerified)
	if err := _DecuxSystem.contract.UnpackLog(event, "BurnVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the DecuxSystem contract.
type DecuxSystemCooldownIterator struct {
	Event *DecuxSystemCooldown // Event containing the contract specifics and raw log

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
func (it *DecuxSystemCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemCooldown)
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
		it.Event = new(DecuxSystemCooldown)
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
func (it *DecuxSystemCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemCooldown represents a Cooldown event raised by the DecuxSystem contract.
type DecuxSystemCooldown struct {
	Keeper  common.Address
	EndTime uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0xc606d47187d0dea01a7ca1a0d8abf6027f8841bf0f82fab363549489e139e2ce.
//
// Solidity: event Cooldown(address indexed keeper, uint32 endTime)
func (_DecuxSystem *DecuxSystemFilterer) FilterCooldown(opts *bind.FilterOpts, keeper []common.Address) (*DecuxSystemCooldownIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "Cooldown", keeperRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemCooldownIterator{contract: _DecuxSystem.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0xc606d47187d0dea01a7ca1a0d8abf6027f8841bf0f82fab363549489e139e2ce.
//
// Solidity: event Cooldown(address indexed keeper, uint32 endTime)
func (_DecuxSystem *DecuxSystemFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *DecuxSystemCooldown, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "Cooldown", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemCooldown)
				if err := _DecuxSystem.contract.UnpackLog(event, "Cooldown", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseCooldown(log types.Log) (*DecuxSystemCooldown, error) {
	event := new(DecuxSystemCooldown)
	if err := _DecuxSystem.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemGroupAddedIterator is returned from FilterGroupAdded and is used to iterate over the raw logs and unpacked data for GroupAdded events raised by the DecuxSystem contract.
type DecuxSystemGroupAddedIterator struct {
	Event *DecuxSystemGroupAdded // Event containing the contract specifics and raw log

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
func (it *DecuxSystemGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemGroupAdded)
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
		it.Event = new(DecuxSystemGroupAdded)
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
func (it *DecuxSystemGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemGroupAdded represents a GroupAdded event raised by the DecuxSystem contract.
type DecuxSystemGroupAdded struct {
	BtcAddress string
	Required   uint32
	MaxSatoshi uint32
	Keepers    []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0xb14273dec0b1f9ed3657f0f047101d6250bbbb76a6c3e7b808dcca46bec0db4c.
//
// Solidity: event GroupAdded(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers)
func (_DecuxSystem *DecuxSystemFilterer) FilterGroupAdded(opts *bind.FilterOpts) (*DecuxSystemGroupAddedIterator, error) {

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return &DecuxSystemGroupAddedIterator{contract: _DecuxSystem.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0xb14273dec0b1f9ed3657f0f047101d6250bbbb76a6c3e7b808dcca46bec0db4c.
//
// Solidity: event GroupAdded(string btcAddress, uint32 required, uint32 maxSatoshi, address[] keepers)
func (_DecuxSystem *DecuxSystemFilterer) WatchGroupAdded(opts *bind.WatchOpts, sink chan<- *DecuxSystemGroupAdded) (event.Subscription, error) {

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemGroupAdded)
				if err := _DecuxSystem.contract.UnpackLog(event, "GroupAdded", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseGroupAdded(log types.Log) (*DecuxSystemGroupAdded, error) {
	event := new(DecuxSystemGroupAdded)
	if err := _DecuxSystem.contract.UnpackLog(event, "GroupAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemGroupDeletedIterator is returned from FilterGroupDeleted and is used to iterate over the raw logs and unpacked data for GroupDeleted events raised by the DecuxSystem contract.
type DecuxSystemGroupDeletedIterator struct {
	Event *DecuxSystemGroupDeleted // Event containing the contract specifics and raw log

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
func (it *DecuxSystemGroupDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemGroupDeleted)
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
		it.Event = new(DecuxSystemGroupDeleted)
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
func (it *DecuxSystemGroupDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemGroupDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemGroupDeleted represents a GroupDeleted event raised by the DecuxSystem contract.
type DecuxSystemGroupDeleted struct {
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGroupDeleted is a free log retrieval operation binding the contract event 0xeacdb5801d8160882398b78a544379ff28cbe8bd4de76aca39013c89311e00d5.
//
// Solidity: event GroupDeleted(string btcAddress)
func (_DecuxSystem *DecuxSystemFilterer) FilterGroupDeleted(opts *bind.FilterOpts) (*DecuxSystemGroupDeletedIterator, error) {

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "GroupDeleted")
	if err != nil {
		return nil, err
	}
	return &DecuxSystemGroupDeletedIterator{contract: _DecuxSystem.contract, event: "GroupDeleted", logs: logs, sub: sub}, nil
}

// WatchGroupDeleted is a free log subscription operation binding the contract event 0xeacdb5801d8160882398b78a544379ff28cbe8bd4de76aca39013c89311e00d5.
//
// Solidity: event GroupDeleted(string btcAddress)
func (_DecuxSystem *DecuxSystemFilterer) WatchGroupDeleted(opts *bind.WatchOpts, sink chan<- *DecuxSystemGroupDeleted) (event.Subscription, error) {

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "GroupDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemGroupDeleted)
				if err := _DecuxSystem.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseGroupDeleted(log types.Log) (*DecuxSystemGroupDeleted, error) {
	event := new(DecuxSystemGroupDeleted)
	if err := _DecuxSystem.contract.UnpackLog(event, "GroupDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemMintRequestedIterator is returned from FilterMintRequested and is used to iterate over the raw logs and unpacked data for MintRequested events raised by the DecuxSystem contract.
type DecuxSystemMintRequestedIterator struct {
	Event *DecuxSystemMintRequested // Event containing the contract specifics and raw log

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
func (it *DecuxSystemMintRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemMintRequested)
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
		it.Event = new(DecuxSystemMintRequested)
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
func (it *DecuxSystemMintRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemMintRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemMintRequested represents a MintRequested event raised by the DecuxSystem contract.
type DecuxSystemMintRequested struct {
	ReceiptId       [32]byte
	Recipient       common.Address
	AmountInSatoshi uint32
	GroupBtcAddress string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintRequested is a free log retrieval operation binding the contract event 0x93e399f7caab186cd9c637476e9d46c57d7e68ecb5392570123578f392d5e22c.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint32 amountInSatoshi, string groupBtcAddress)
func (_DecuxSystem *DecuxSystemFilterer) FilterMintRequested(opts *bind.FilterOpts, receiptId [][32]byte, recipient []common.Address) (*DecuxSystemMintRequestedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "MintRequested", receiptIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemMintRequestedIterator{contract: _DecuxSystem.contract, event: "MintRequested", logs: logs, sub: sub}, nil
}

// WatchMintRequested is a free log subscription operation binding the contract event 0x93e399f7caab186cd9c637476e9d46c57d7e68ecb5392570123578f392d5e22c.
//
// Solidity: event MintRequested(bytes32 indexed receiptId, address indexed recipient, uint32 amountInSatoshi, string groupBtcAddress)
func (_DecuxSystem *DecuxSystemFilterer) WatchMintRequested(opts *bind.WatchOpts, sink chan<- *DecuxSystemMintRequested, receiptId [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "MintRequested", receiptIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemMintRequested)
				if err := _DecuxSystem.contract.UnpackLog(event, "MintRequested", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseMintRequested(log types.Log) (*DecuxSystemMintRequested, error) {
	event := new(DecuxSystemMintRequested)
	if err := _DecuxSystem.contract.UnpackLog(event, "MintRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemMintRevokedIterator is returned from FilterMintRevoked and is used to iterate over the raw logs and unpacked data for MintRevoked events raised by the DecuxSystem contract.
type DecuxSystemMintRevokedIterator struct {
	Event *DecuxSystemMintRevoked // Event containing the contract specifics and raw log

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
func (it *DecuxSystemMintRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemMintRevoked)
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
		it.Event = new(DecuxSystemMintRevoked)
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
func (it *DecuxSystemMintRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemMintRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemMintRevoked represents a MintRevoked event raised by the DecuxSystem contract.
type DecuxSystemMintRevoked struct {
	ReceiptId       [32]byte
	GroupBtcAddress string
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintRevoked is a free log retrieval operation binding the contract event 0x24da21178c24a520845b59d6d351ecf3033e3ceefb8c66f286f667caa74cd2c3.
//
// Solidity: event MintRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) FilterMintRevoked(opts *bind.FilterOpts, receiptId [][32]byte) (*DecuxSystemMintRevokedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "MintRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemMintRevokedIterator{contract: _DecuxSystem.contract, event: "MintRevoked", logs: logs, sub: sub}, nil
}

// WatchMintRevoked is a free log subscription operation binding the contract event 0x24da21178c24a520845b59d6d351ecf3033e3ceefb8c66f286f667caa74cd2c3.
//
// Solidity: event MintRevoked(bytes32 indexed receiptId, string groupBtcAddress, address operator)
func (_DecuxSystem *DecuxSystemFilterer) WatchMintRevoked(opts *bind.WatchOpts, sink chan<- *DecuxSystemMintRevoked, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "MintRevoked", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemMintRevoked)
				if err := _DecuxSystem.contract.UnpackLog(event, "MintRevoked", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseMintRevoked(log types.Log) (*DecuxSystemMintRevoked, error) {
	event := new(DecuxSystemMintRevoked)
	if err := _DecuxSystem.contract.UnpackLog(event, "MintRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemMintVerifiedIterator is returned from FilterMintVerified and is used to iterate over the raw logs and unpacked data for MintVerified events raised by the DecuxSystem contract.
type DecuxSystemMintVerifiedIterator struct {
	Event *DecuxSystemMintVerified // Event containing the contract specifics and raw log

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
func (it *DecuxSystemMintVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemMintVerified)
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
		it.Event = new(DecuxSystemMintVerified)
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
func (it *DecuxSystemMintVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemMintVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemMintVerified represents a MintVerified event raised by the DecuxSystem contract.
type DecuxSystemMintVerified struct {
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
func (_DecuxSystem *DecuxSystemFilterer) FilterMintVerified(opts *bind.FilterOpts, receiptId [][32]byte) (*DecuxSystemMintVerifiedIterator, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "MintVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemMintVerifiedIterator{contract: _DecuxSystem.contract, event: "MintVerified", logs: logs, sub: sub}, nil
}

// WatchMintVerified is a free log subscription operation binding the contract event 0x19d0f59e71cb6f748028c86f27cb22d05c57315a9c850320ecda2cd13f5a43d5.
//
// Solidity: event MintVerified(bytes32 indexed receiptId, string groupBtcAddress, address[] keepers, bytes32 btcTxId, uint32 btcTxHeight)
func (_DecuxSystem *DecuxSystemFilterer) WatchMintVerified(opts *bind.WatchOpts, sink chan<- *DecuxSystemMintVerified, receiptId [][32]byte) (event.Subscription, error) {

	var receiptIdRule []interface{}
	for _, receiptIdItem := range receiptId {
		receiptIdRule = append(receiptIdRule, receiptIdItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "MintVerified", receiptIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemMintVerified)
				if err := _DecuxSystem.contract.UnpackLog(event, "MintVerified", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseMintVerified(log types.Log) (*DecuxSystemMintVerified, error) {
	event := new(DecuxSystemMintVerified)
	if err := _DecuxSystem.contract.UnpackLog(event, "MintVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DecuxSystem contract.
type DecuxSystemPausedIterator struct {
	Event *DecuxSystemPaused // Event containing the contract specifics and raw log

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
func (it *DecuxSystemPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemPaused)
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
		it.Event = new(DecuxSystemPaused)
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
func (it *DecuxSystemPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemPaused represents a Paused event raised by the DecuxSystem contract.
type DecuxSystemPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DecuxSystem *DecuxSystemFilterer) FilterPaused(opts *bind.FilterOpts) (*DecuxSystemPausedIterator, error) {

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DecuxSystemPausedIterator{contract: _DecuxSystem.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DecuxSystem *DecuxSystemFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DecuxSystemPaused) (event.Subscription, error) {

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemPaused)
				if err := _DecuxSystem.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParsePaused(log types.Log) (*DecuxSystemPaused, error) {
	event := new(DecuxSystemPaused)
	if err := _DecuxSystem.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DecuxSystem contract.
type DecuxSystemRoleAdminChangedIterator struct {
	Event *DecuxSystemRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DecuxSystemRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemRoleAdminChanged)
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
		it.Event = new(DecuxSystemRoleAdminChanged)
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
func (it *DecuxSystemRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemRoleAdminChanged represents a RoleAdminChanged event raised by the DecuxSystem contract.
type DecuxSystemRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DecuxSystem *DecuxSystemFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DecuxSystemRoleAdminChangedIterator, error) {

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

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemRoleAdminChangedIterator{contract: _DecuxSystem.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DecuxSystem *DecuxSystemFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DecuxSystemRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemRoleAdminChanged)
				if err := _DecuxSystem.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseRoleAdminChanged(log types.Log) (*DecuxSystemRoleAdminChanged, error) {
	event := new(DecuxSystemRoleAdminChanged)
	if err := _DecuxSystem.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DecuxSystem contract.
type DecuxSystemRoleGrantedIterator struct {
	Event *DecuxSystemRoleGranted // Event containing the contract specifics and raw log

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
func (it *DecuxSystemRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemRoleGranted)
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
		it.Event = new(DecuxSystemRoleGranted)
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
func (it *DecuxSystemRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemRoleGranted represents a RoleGranted event raised by the DecuxSystem contract.
type DecuxSystemRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DecuxSystem *DecuxSystemFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DecuxSystemRoleGrantedIterator, error) {

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

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemRoleGrantedIterator{contract: _DecuxSystem.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DecuxSystem *DecuxSystemFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DecuxSystemRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemRoleGranted)
				if err := _DecuxSystem.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseRoleGranted(log types.Log) (*DecuxSystemRoleGranted, error) {
	event := new(DecuxSystemRoleGranted)
	if err := _DecuxSystem.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DecuxSystem contract.
type DecuxSystemRoleRevokedIterator struct {
	Event *DecuxSystemRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DecuxSystemRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemRoleRevoked)
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
		it.Event = new(DecuxSystemRoleRevoked)
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
func (it *DecuxSystemRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemRoleRevoked represents a RoleRevoked event raised by the DecuxSystem contract.
type DecuxSystemRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DecuxSystem *DecuxSystemFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DecuxSystemRoleRevokedIterator, error) {

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

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemRoleRevokedIterator{contract: _DecuxSystem.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DecuxSystem *DecuxSystemFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DecuxSystemRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemRoleRevoked)
				if err := _DecuxSystem.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseRoleRevoked(log types.Log) (*DecuxSystemRoleRevoked, error) {
	event := new(DecuxSystemRoleRevoked)
	if err := _DecuxSystem.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemToggleExitKeeperIterator is returned from FilterToggleExitKeeper and is used to iterate over the raw logs and unpacked data for ToggleExitKeeper events raised by the DecuxSystem contract.
type DecuxSystemToggleExitKeeperIterator struct {
	Event *DecuxSystemToggleExitKeeper // Event containing the contract specifics and raw log

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
func (it *DecuxSystemToggleExitKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemToggleExitKeeper)
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
		it.Event = new(DecuxSystemToggleExitKeeper)
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
func (it *DecuxSystemToggleExitKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemToggleExitKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemToggleExitKeeper represents a ToggleExitKeeper event raised by the DecuxSystem contract.
type DecuxSystemToggleExitKeeper struct {
	Keeper common.Address
	IsExit bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterToggleExitKeeper is a free log retrieval operation binding the contract event 0x83cc6c1fe93786e75c247528820fb8e5def97df72d6233ff0afaad2b46e54a21.
//
// Solidity: event ToggleExitKeeper(address indexed keeper, bool isExit)
func (_DecuxSystem *DecuxSystemFilterer) FilterToggleExitKeeper(opts *bind.FilterOpts, keeper []common.Address) (*DecuxSystemToggleExitKeeperIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "ToggleExitKeeper", keeperRule)
	if err != nil {
		return nil, err
	}
	return &DecuxSystemToggleExitKeeperIterator{contract: _DecuxSystem.contract, event: "ToggleExitKeeper", logs: logs, sub: sub}, nil
}

// WatchToggleExitKeeper is a free log subscription operation binding the contract event 0x83cc6c1fe93786e75c247528820fb8e5def97df72d6233ff0afaad2b46e54a21.
//
// Solidity: event ToggleExitKeeper(address indexed keeper, bool isExit)
func (_DecuxSystem *DecuxSystemFilterer) WatchToggleExitKeeper(opts *bind.WatchOpts, sink chan<- *DecuxSystemToggleExitKeeper, keeper []common.Address) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "ToggleExitKeeper", keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemToggleExitKeeper)
				if err := _DecuxSystem.contract.UnpackLog(event, "ToggleExitKeeper", log); err != nil {
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

// ParseToggleExitKeeper is a log parse operation binding the contract event 0x83cc6c1fe93786e75c247528820fb8e5def97df72d6233ff0afaad2b46e54a21.
//
// Solidity: event ToggleExitKeeper(address indexed keeper, bool isExit)
func (_DecuxSystem *DecuxSystemFilterer) ParseToggleExitKeeper(log types.Log) (*DecuxSystemToggleExitKeeper, error) {
	event := new(DecuxSystemToggleExitKeeper)
	if err := _DecuxSystem.contract.UnpackLog(event, "ToggleExitKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecuxSystemUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DecuxSystem contract.
type DecuxSystemUnpausedIterator struct {
	Event *DecuxSystemUnpaused // Event containing the contract specifics and raw log

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
func (it *DecuxSystemUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecuxSystemUnpaused)
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
		it.Event = new(DecuxSystemUnpaused)
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
func (it *DecuxSystemUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecuxSystemUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecuxSystemUnpaused represents a Unpaused event raised by the DecuxSystem contract.
type DecuxSystemUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DecuxSystem *DecuxSystemFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DecuxSystemUnpausedIterator, error) {

	logs, sub, err := _DecuxSystem.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DecuxSystemUnpausedIterator{contract: _DecuxSystem.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DecuxSystem *DecuxSystemFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DecuxSystemUnpaused) (event.Subscription, error) {

	logs, sub, err := _DecuxSystem.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecuxSystemUnpaused)
				if err := _DecuxSystem.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DecuxSystem *DecuxSystemFilterer) ParseUnpaused(log types.Log) (*DecuxSystemUnpaused, error) {
	event := new(DecuxSystemUnpaused)
	if err := _DecuxSystem.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
