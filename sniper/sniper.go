// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sniper

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
)

// SniperMetaData contains all meta data concerning the Sniper contract.
var SniperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeMultipleAccountsFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"SetAutomatedMarketMakerPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIntoLiqudity\",\"type\":\"uint256\"}],\"name\":\"SwapAndLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_earlyBuyer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_isExcludedMaxTransactionAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_mevBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"automatedMarketMakerPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bots_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"changeAccountStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"updAds\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEx\",\"type\":\"bool\"}],\"name\":\"excludeFromMaxTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeMultipleAccountsFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitsInEffect\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityActiveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manualSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingBuyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingSellFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTransactionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLimits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAutomatedMarketMakerPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTokensAtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenLockBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensForFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensForMarketing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuyFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSellFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingActiveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferDelayEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketingBuyFee\",\"type\":\"uint256\"}],\"name\":\"updateBuyFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extrablocks\",\"type\":\"uint256\"}],\"name\":\"updateEarlyBuyBlockNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNum\",\"type\":\"uint256\"}],\"name\":\"updateMaxTxnAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNum\",\"type\":\"uint256\"}],\"name\":\"updateMaxWalletAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketingSellFee\",\"type\":\"uint256\"}],\"name\":\"updateSellFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"updateSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStuckEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SniperABI is the input ABI used to generate the binding from.
// Deprecated: Use SniperMetaData.ABI instead.
var SniperABI = SniperMetaData.ABI

// Sniper is an auto generated Go binding around an Ethereum contract.
type Sniper struct {
	SniperCaller     // Read-only binding to the contract
	SniperTransactor // Write-only binding to the contract
	SniperFilterer   // Log filterer for contract events
}

// SniperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SniperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SniperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SniperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SniperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SniperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SniperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SniperSession struct {
	Contract     *Sniper           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SniperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SniperCallerSession struct {
	Contract *SniperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SniperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SniperTransactorSession struct {
	Contract     *SniperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SniperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SniperRaw struct {
	Contract *Sniper // Generic contract binding to access the raw methods on
}

// SniperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SniperCallerRaw struct {
	Contract *SniperCaller // Generic read-only contract binding to access the raw methods on
}

// SniperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SniperTransactorRaw struct {
	Contract *SniperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSniper creates a new instance of Sniper, bound to a specific deployed contract.
func NewSniper(address common.Address, backend bind.ContractBackend) (*Sniper, error) {
	contract, err := bindSniper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sniper{SniperCaller: SniperCaller{contract: contract}, SniperTransactor: SniperTransactor{contract: contract}, SniperFilterer: SniperFilterer{contract: contract}}, nil
}

// NewSniperCaller creates a new read-only instance of Sniper, bound to a specific deployed contract.
func NewSniperCaller(address common.Address, caller bind.ContractCaller) (*SniperCaller, error) {
	contract, err := bindSniper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SniperCaller{contract: contract}, nil
}

// NewSniperTransactor creates a new write-only instance of Sniper, bound to a specific deployed contract.
func NewSniperTransactor(address common.Address, transactor bind.ContractTransactor) (*SniperTransactor, error) {
	contract, err := bindSniper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SniperTransactor{contract: contract}, nil
}

// NewSniperFilterer creates a new log filterer instance of Sniper, bound to a specific deployed contract.
func NewSniperFilterer(address common.Address, filterer bind.ContractFilterer) (*SniperFilterer, error) {
	contract, err := bindSniper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SniperFilterer{contract: contract}, nil
}

// bindSniper binds a generic wrapper to an already deployed contract.
func bindSniper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SniperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sniper *SniperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sniper.Contract.SniperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sniper *SniperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.Contract.SniperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sniper *SniperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sniper.Contract.SniperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sniper *SniperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sniper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sniper *SniperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sniper *SniperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sniper.Contract.contract.Transact(opts, method, params...)
}

// EarlyBuyer is a free data retrieval call binding the contract method 0x2a3de72e.
//
// Solidity: function _earlyBuyer(address ) view returns(uint256)
func (_Sniper *SniperCaller) EarlyBuyer(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "_earlyBuyer", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarlyBuyer is a free data retrieval call binding the contract method 0x2a3de72e.
//
// Solidity: function _earlyBuyer(address ) view returns(uint256)
func (_Sniper *SniperSession) EarlyBuyer(arg0 common.Address) (*big.Int, error) {
	return _Sniper.Contract.EarlyBuyer(&_Sniper.CallOpts, arg0)
}

// EarlyBuyer is a free data retrieval call binding the contract method 0x2a3de72e.
//
// Solidity: function _earlyBuyer(address ) view returns(uint256)
func (_Sniper *SniperCallerSession) EarlyBuyer(arg0 common.Address) (*big.Int, error) {
	return _Sniper.Contract.EarlyBuyer(&_Sniper.CallOpts, arg0)
}

// IsExcludedMaxTransactionAmount is a free data retrieval call binding the contract method 0x10d5de53.
//
// Solidity: function _isExcludedMaxTransactionAmount(address ) view returns(bool)
func (_Sniper *SniperCaller) IsExcludedMaxTransactionAmount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "_isExcludedMaxTransactionAmount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedMaxTransactionAmount is a free data retrieval call binding the contract method 0x10d5de53.
//
// Solidity: function _isExcludedMaxTransactionAmount(address ) view returns(bool)
func (_Sniper *SniperSession) IsExcludedMaxTransactionAmount(arg0 common.Address) (bool, error) {
	return _Sniper.Contract.IsExcludedMaxTransactionAmount(&_Sniper.CallOpts, arg0)
}

// IsExcludedMaxTransactionAmount is a free data retrieval call binding the contract method 0x10d5de53.
//
// Solidity: function _isExcludedMaxTransactionAmount(address ) view returns(bool)
func (_Sniper *SniperCallerSession) IsExcludedMaxTransactionAmount(arg0 common.Address) (bool, error) {
	return _Sniper.Contract.IsExcludedMaxTransactionAmount(&_Sniper.CallOpts, arg0)
}

// MevBlock is a free data retrieval call binding the contract method 0x25fedfb5.
//
// Solidity: function _mevBlock(address ) view returns(uint256)
func (_Sniper *SniperCaller) MevBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "_mevBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MevBlock is a free data retrieval call binding the contract method 0x25fedfb5.
//
// Solidity: function _mevBlock(address ) view returns(uint256)
func (_Sniper *SniperSession) MevBlock(arg0 common.Address) (*big.Int, error) {
	return _Sniper.Contract.MevBlock(&_Sniper.CallOpts, arg0)
}

// MevBlock is a free data retrieval call binding the contract method 0x25fedfb5.
//
// Solidity: function _mevBlock(address ) view returns(uint256)
func (_Sniper *SniperCallerSession) MevBlock(arg0 common.Address) (*big.Int, error) {
	return _Sniper.Contract.MevBlock(&_Sniper.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sniper *SniperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sniper *SniperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sniper.Contract.Allowance(&_Sniper.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sniper *SniperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sniper.Contract.Allowance(&_Sniper.CallOpts, owner, spender)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Sniper *SniperCaller) AutomatedMarketMakerPairs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "automatedMarketMakerPairs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Sniper *SniperSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Sniper.Contract.AutomatedMarketMakerPairs(&_Sniper.CallOpts, arg0)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Sniper *SniperCallerSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Sniper.Contract.AutomatedMarketMakerPairs(&_Sniper.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sniper *SniperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sniper *SniperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Sniper.Contract.BalanceOf(&_Sniper.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sniper *SniperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Sniper.Contract.BalanceOf(&_Sniper.CallOpts, account)
}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_Sniper *SniperCaller) BurnWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "burnWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_Sniper *SniperSession) BurnWallet() (common.Address, error) {
	return _Sniper.Contract.BurnWallet(&_Sniper.CallOpts)
}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_Sniper *SniperCallerSession) BurnWallet() (common.Address, error) {
	return _Sniper.Contract.BurnWallet(&_Sniper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sniper *SniperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sniper *SniperSession) Decimals() (uint8, error) {
	return _Sniper.Contract.Decimals(&_Sniper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sniper *SniperCallerSession) Decimals() (uint8, error) {
	return _Sniper.Contract.Decimals(&_Sniper.CallOpts)
}

// FeeDivisor is a free data retrieval call binding the contract method 0x9a36f932.
//
// Solidity: function feeDivisor() view returns(uint256)
func (_Sniper *SniperCaller) FeeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "feeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDivisor is a free data retrieval call binding the contract method 0x9a36f932.
//
// Solidity: function feeDivisor() view returns(uint256)
func (_Sniper *SniperSession) FeeDivisor() (*big.Int, error) {
	return _Sniper.Contract.FeeDivisor(&_Sniper.CallOpts)
}

// FeeDivisor is a free data retrieval call binding the contract method 0x9a36f932.
//
// Solidity: function feeDivisor() view returns(uint256)
func (_Sniper *SniperCallerSession) FeeDivisor() (*big.Int, error) {
	return _Sniper.Contract.FeeDivisor(&_Sniper.CallOpts)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Sniper *SniperCaller) IsExcludedFromFees(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "isExcludedFromFees", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Sniper *SniperSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Sniper.Contract.IsExcludedFromFees(&_Sniper.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Sniper *SniperCallerSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Sniper.Contract.IsExcludedFromFees(&_Sniper.CallOpts, account)
}

// LimitsInEffect is a free data retrieval call binding the contract method 0x4a62bb65.
//
// Solidity: function limitsInEffect() view returns(bool)
func (_Sniper *SniperCaller) LimitsInEffect(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "limitsInEffect")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LimitsInEffect is a free data retrieval call binding the contract method 0x4a62bb65.
//
// Solidity: function limitsInEffect() view returns(bool)
func (_Sniper *SniperSession) LimitsInEffect() (bool, error) {
	return _Sniper.Contract.LimitsInEffect(&_Sniper.CallOpts)
}

// LimitsInEffect is a free data retrieval call binding the contract method 0x4a62bb65.
//
// Solidity: function limitsInEffect() view returns(bool)
func (_Sniper *SniperCallerSession) LimitsInEffect() (bool, error) {
	return _Sniper.Contract.LimitsInEffect(&_Sniper.CallOpts)
}

// LiquidityActiveBlock is a free data retrieval call binding the contract method 0x0f4432e3.
//
// Solidity: function liquidityActiveBlock() view returns(uint256)
func (_Sniper *SniperCaller) LiquidityActiveBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "liquidityActiveBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityActiveBlock is a free data retrieval call binding the contract method 0x0f4432e3.
//
// Solidity: function liquidityActiveBlock() view returns(uint256)
func (_Sniper *SniperSession) LiquidityActiveBlock() (*big.Int, error) {
	return _Sniper.Contract.LiquidityActiveBlock(&_Sniper.CallOpts)
}

// LiquidityActiveBlock is a free data retrieval call binding the contract method 0x0f4432e3.
//
// Solidity: function liquidityActiveBlock() view returns(uint256)
func (_Sniper *SniperCallerSession) LiquidityActiveBlock() (*big.Int, error) {
	return _Sniper.Contract.LiquidityActiveBlock(&_Sniper.CallOpts)
}

// MarketingBuyFee is a free data retrieval call binding the contract method 0x68078952.
//
// Solidity: function marketingBuyFee() view returns(uint256)
func (_Sniper *SniperCaller) MarketingBuyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "marketingBuyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingBuyFee is a free data retrieval call binding the contract method 0x68078952.
//
// Solidity: function marketingBuyFee() view returns(uint256)
func (_Sniper *SniperSession) MarketingBuyFee() (*big.Int, error) {
	return _Sniper.Contract.MarketingBuyFee(&_Sniper.CallOpts)
}

// MarketingBuyFee is a free data retrieval call binding the contract method 0x68078952.
//
// Solidity: function marketingBuyFee() view returns(uint256)
func (_Sniper *SniperCallerSession) MarketingBuyFee() (*big.Int, error) {
	return _Sniper.Contract.MarketingBuyFee(&_Sniper.CallOpts)
}

// MarketingSellFee is a free data retrieval call binding the contract method 0xe7f444b3.
//
// Solidity: function marketingSellFee() view returns(uint256)
func (_Sniper *SniperCaller) MarketingSellFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "marketingSellFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingSellFee is a free data retrieval call binding the contract method 0xe7f444b3.
//
// Solidity: function marketingSellFee() view returns(uint256)
func (_Sniper *SniperSession) MarketingSellFee() (*big.Int, error) {
	return _Sniper.Contract.MarketingSellFee(&_Sniper.CallOpts)
}

// MarketingSellFee is a free data retrieval call binding the contract method 0xe7f444b3.
//
// Solidity: function marketingSellFee() view returns(uint256)
func (_Sniper *SniperCallerSession) MarketingSellFee() (*big.Int, error) {
	return _Sniper.Contract.MarketingSellFee(&_Sniper.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_Sniper *SniperCaller) MarketingWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "marketingWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_Sniper *SniperSession) MarketingWallet() (common.Address, error) {
	return _Sniper.Contract.MarketingWallet(&_Sniper.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_Sniper *SniperCallerSession) MarketingWallet() (common.Address, error) {
	return _Sniper.Contract.MarketingWallet(&_Sniper.CallOpts)
}

// MaxTransactionAmount is a free data retrieval call binding the contract method 0xc8c8ebe4.
//
// Solidity: function maxTransactionAmount() view returns(uint256)
func (_Sniper *SniperCaller) MaxTransactionAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "maxTransactionAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTransactionAmount is a free data retrieval call binding the contract method 0xc8c8ebe4.
//
// Solidity: function maxTransactionAmount() view returns(uint256)
func (_Sniper *SniperSession) MaxTransactionAmount() (*big.Int, error) {
	return _Sniper.Contract.MaxTransactionAmount(&_Sniper.CallOpts)
}

// MaxTransactionAmount is a free data retrieval call binding the contract method 0xc8c8ebe4.
//
// Solidity: function maxTransactionAmount() view returns(uint256)
func (_Sniper *SniperCallerSession) MaxTransactionAmount() (*big.Int, error) {
	return _Sniper.Contract.MaxTransactionAmount(&_Sniper.CallOpts)
}

// MaxWallet is a free data retrieval call binding the contract method 0xf8b45b05.
//
// Solidity: function maxWallet() view returns(uint256)
func (_Sniper *SniperCaller) MaxWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "maxWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWallet is a free data retrieval call binding the contract method 0xf8b45b05.
//
// Solidity: function maxWallet() view returns(uint256)
func (_Sniper *SniperSession) MaxWallet() (*big.Int, error) {
	return _Sniper.Contract.MaxWallet(&_Sniper.CallOpts)
}

// MaxWallet is a free data retrieval call binding the contract method 0xf8b45b05.
//
// Solidity: function maxWallet() view returns(uint256)
func (_Sniper *SniperCallerSession) MaxWallet() (*big.Int, error) {
	return _Sniper.Contract.MaxWallet(&_Sniper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sniper *SniperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sniper *SniperSession) Name() (string, error) {
	return _Sniper.Contract.Name(&_Sniper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sniper *SniperCallerSession) Name() (string, error) {
	return _Sniper.Contract.Name(&_Sniper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sniper *SniperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sniper *SniperSession) Owner() (common.Address, error) {
	return _Sniper.Contract.Owner(&_Sniper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sniper *SniperCallerSession) Owner() (common.Address, error) {
	return _Sniper.Contract.Owner(&_Sniper.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Sniper *SniperCaller) SwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "swapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Sniper *SniperSession) SwapEnabled() (bool, error) {
	return _Sniper.Contract.SwapEnabled(&_Sniper.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Sniper *SniperCallerSession) SwapEnabled() (bool, error) {
	return _Sniper.Contract.SwapEnabled(&_Sniper.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Sniper *SniperCaller) SwapTokensAtAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "swapTokensAtAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Sniper *SniperSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Sniper.Contract.SwapTokensAtAmount(&_Sniper.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Sniper *SniperCallerSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Sniper.Contract.SwapTokensAtAmount(&_Sniper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sniper *SniperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sniper *SniperSession) Symbol() (string, error) {
	return _Sniper.Contract.Symbol(&_Sniper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sniper *SniperCallerSession) Symbol() (string, error) {
	return _Sniper.Contract.Symbol(&_Sniper.CallOpts)
}

// TokenLockBlockNum is a free data retrieval call binding the contract method 0xb19cca8f.
//
// Solidity: function tokenLockBlockNum() view returns(uint256)
func (_Sniper *SniperCaller) TokenLockBlockNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "tokenLockBlockNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenLockBlockNum is a free data retrieval call binding the contract method 0xb19cca8f.
//
// Solidity: function tokenLockBlockNum() view returns(uint256)
func (_Sniper *SniperSession) TokenLockBlockNum() (*big.Int, error) {
	return _Sniper.Contract.TokenLockBlockNum(&_Sniper.CallOpts)
}

// TokenLockBlockNum is a free data retrieval call binding the contract method 0xb19cca8f.
//
// Solidity: function tokenLockBlockNum() view returns(uint256)
func (_Sniper *SniperCallerSession) TokenLockBlockNum() (*big.Int, error) {
	return _Sniper.Contract.TokenLockBlockNum(&_Sniper.CallOpts)
}

// TokensForFees is a free data retrieval call binding the contract method 0x3f8a6204.
//
// Solidity: function tokensForFees() view returns(uint256)
func (_Sniper *SniperCaller) TokensForFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "tokensForFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensForFees is a free data retrieval call binding the contract method 0x3f8a6204.
//
// Solidity: function tokensForFees() view returns(uint256)
func (_Sniper *SniperSession) TokensForFees() (*big.Int, error) {
	return _Sniper.Contract.TokensForFees(&_Sniper.CallOpts)
}

// TokensForFees is a free data retrieval call binding the contract method 0x3f8a6204.
//
// Solidity: function tokensForFees() view returns(uint256)
func (_Sniper *SniperCallerSession) TokensForFees() (*big.Int, error) {
	return _Sniper.Contract.TokensForFees(&_Sniper.CallOpts)
}

// TokensForMarketing is a free data retrieval call binding the contract method 0x1f3fed8f.
//
// Solidity: function tokensForMarketing() view returns(uint256)
func (_Sniper *SniperCaller) TokensForMarketing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "tokensForMarketing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensForMarketing is a free data retrieval call binding the contract method 0x1f3fed8f.
//
// Solidity: function tokensForMarketing() view returns(uint256)
func (_Sniper *SniperSession) TokensForMarketing() (*big.Int, error) {
	return _Sniper.Contract.TokensForMarketing(&_Sniper.CallOpts)
}

// TokensForMarketing is a free data retrieval call binding the contract method 0x1f3fed8f.
//
// Solidity: function tokensForMarketing() view returns(uint256)
func (_Sniper *SniperCallerSession) TokensForMarketing() (*big.Int, error) {
	return _Sniper.Contract.TokensForMarketing(&_Sniper.CallOpts)
}

// TotalBuyFees is a free data retrieval call binding the contract method 0xb9e93700.
//
// Solidity: function totalBuyFees() view returns(uint256)
func (_Sniper *SniperCaller) TotalBuyFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "totalBuyFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuyFees is a free data retrieval call binding the contract method 0xb9e93700.
//
// Solidity: function totalBuyFees() view returns(uint256)
func (_Sniper *SniperSession) TotalBuyFees() (*big.Int, error) {
	return _Sniper.Contract.TotalBuyFees(&_Sniper.CallOpts)
}

// TotalBuyFees is a free data retrieval call binding the contract method 0xb9e93700.
//
// Solidity: function totalBuyFees() view returns(uint256)
func (_Sniper *SniperCallerSession) TotalBuyFees() (*big.Int, error) {
	return _Sniper.Contract.TotalBuyFees(&_Sniper.CallOpts)
}

// TotalSellFees is a free data retrieval call binding the contract method 0xd0a39814.
//
// Solidity: function totalSellFees() view returns(uint256)
func (_Sniper *SniperCaller) TotalSellFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "totalSellFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSellFees is a free data retrieval call binding the contract method 0xd0a39814.
//
// Solidity: function totalSellFees() view returns(uint256)
func (_Sniper *SniperSession) TotalSellFees() (*big.Int, error) {
	return _Sniper.Contract.TotalSellFees(&_Sniper.CallOpts)
}

// TotalSellFees is a free data retrieval call binding the contract method 0xd0a39814.
//
// Solidity: function totalSellFees() view returns(uint256)
func (_Sniper *SniperCallerSession) TotalSellFees() (*big.Int, error) {
	return _Sniper.Contract.TotalSellFees(&_Sniper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sniper *SniperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sniper *SniperSession) TotalSupply() (*big.Int, error) {
	return _Sniper.Contract.TotalSupply(&_Sniper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sniper *SniperCallerSession) TotalSupply() (*big.Int, error) {
	return _Sniper.Contract.TotalSupply(&_Sniper.CallOpts)
}

// TradingActive is a free data retrieval call binding the contract method 0xbbc0c742.
//
// Solidity: function tradingActive() view returns(bool)
func (_Sniper *SniperCaller) TradingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "tradingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradingActive is a free data retrieval call binding the contract method 0xbbc0c742.
//
// Solidity: function tradingActive() view returns(bool)
func (_Sniper *SniperSession) TradingActive() (bool, error) {
	return _Sniper.Contract.TradingActive(&_Sniper.CallOpts)
}

// TradingActive is a free data retrieval call binding the contract method 0xbbc0c742.
//
// Solidity: function tradingActive() view returns(bool)
func (_Sniper *SniperCallerSession) TradingActive() (bool, error) {
	return _Sniper.Contract.TradingActive(&_Sniper.CallOpts)
}

// TradingActiveBlock is a free data retrieval call binding the contract method 0xee40166e.
//
// Solidity: function tradingActiveBlock() view returns(uint256)
func (_Sniper *SniperCaller) TradingActiveBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "tradingActiveBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradingActiveBlock is a free data retrieval call binding the contract method 0xee40166e.
//
// Solidity: function tradingActiveBlock() view returns(uint256)
func (_Sniper *SniperSession) TradingActiveBlock() (*big.Int, error) {
	return _Sniper.Contract.TradingActiveBlock(&_Sniper.CallOpts)
}

// TradingActiveBlock is a free data retrieval call binding the contract method 0xee40166e.
//
// Solidity: function tradingActiveBlock() view returns(uint256)
func (_Sniper *SniperCallerSession) TradingActiveBlock() (*big.Int, error) {
	return _Sniper.Contract.TradingActiveBlock(&_Sniper.CallOpts)
}

// TransferDelayEnabled is a free data retrieval call binding the contract method 0xc876d0b9.
//
// Solidity: function transferDelayEnabled() view returns(bool)
func (_Sniper *SniperCaller) TransferDelayEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "transferDelayEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferDelayEnabled is a free data retrieval call binding the contract method 0xc876d0b9.
//
// Solidity: function transferDelayEnabled() view returns(bool)
func (_Sniper *SniperSession) TransferDelayEnabled() (bool, error) {
	return _Sniper.Contract.TransferDelayEnabled(&_Sniper.CallOpts)
}

// TransferDelayEnabled is a free data retrieval call binding the contract method 0xc876d0b9.
//
// Solidity: function transferDelayEnabled() view returns(bool)
func (_Sniper *SniperCallerSession) TransferDelayEnabled() (bool, error) {
	return _Sniper.Contract.TransferDelayEnabled(&_Sniper.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Sniper *SniperCaller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Sniper *SniperSession) UniswapV2Pair() (common.Address, error) {
	return _Sniper.Contract.UniswapV2Pair(&_Sniper.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Sniper *SniperCallerSession) UniswapV2Pair() (common.Address, error) {
	return _Sniper.Contract.UniswapV2Pair(&_Sniper.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Sniper *SniperCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sniper.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Sniper *SniperSession) UniswapV2Router() (common.Address, error) {
	return _Sniper.Contract.UniswapV2Router(&_Sniper.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Sniper *SniperCallerSession) UniswapV2Router() (common.Address, error) {
	return _Sniper.Contract.UniswapV2Router(&_Sniper.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sniper *SniperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sniper *SniperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.Approve(&_Sniper.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sniper *SniperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.Approve(&_Sniper.TransactOpts, spender, amount)
}

// ChangeAccountStatus is a paid mutator transaction binding the contract method 0x17533bca.
//
// Solidity: function changeAccountStatus(address[] bots_, bool status) returns()
func (_Sniper *SniperTransactor) ChangeAccountStatus(opts *bind.TransactOpts, bots_ []common.Address, status bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "changeAccountStatus", bots_, status)
}

// ChangeAccountStatus is a paid mutator transaction binding the contract method 0x17533bca.
//
// Solidity: function changeAccountStatus(address[] bots_, bool status) returns()
func (_Sniper *SniperSession) ChangeAccountStatus(bots_ []common.Address, status bool) (*types.Transaction, error) {
	return _Sniper.Contract.ChangeAccountStatus(&_Sniper.TransactOpts, bots_, status)
}

// ChangeAccountStatus is a paid mutator transaction binding the contract method 0x17533bca.
//
// Solidity: function changeAccountStatus(address[] bots_, bool status) returns()
func (_Sniper *SniperTransactorSession) ChangeAccountStatus(bots_ []common.Address, status bool) (*types.Transaction, error) {
	return _Sniper.Contract.ChangeAccountStatus(&_Sniper.TransactOpts, bots_, status)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sniper *SniperTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sniper *SniperSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.DecreaseAllowance(&_Sniper.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sniper *SniperTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.DecreaseAllowance(&_Sniper.TransactOpts, spender, subtractedValue)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Sniper *SniperTransactor) EnableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "enableTrading")
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Sniper *SniperSession) EnableTrading() (*types.Transaction, error) {
	return _Sniper.Contract.EnableTrading(&_Sniper.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Sniper *SniperTransactorSession) EnableTrading() (*types.Transaction, error) {
	return _Sniper.Contract.EnableTrading(&_Sniper.TransactOpts)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Sniper *SniperTransactor) ExcludeFromFees(opts *bind.TransactOpts, account common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "excludeFromFees", account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Sniper *SniperSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeFromFees(&_Sniper.TransactOpts, account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Sniper *SniperTransactorSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeFromFees(&_Sniper.TransactOpts, account, excluded)
}

// ExcludeFromMaxTransaction is a paid mutator transaction binding the contract method 0x7571336a.
//
// Solidity: function excludeFromMaxTransaction(address updAds, bool isEx) returns()
func (_Sniper *SniperTransactor) ExcludeFromMaxTransaction(opts *bind.TransactOpts, updAds common.Address, isEx bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "excludeFromMaxTransaction", updAds, isEx)
}

// ExcludeFromMaxTransaction is a paid mutator transaction binding the contract method 0x7571336a.
//
// Solidity: function excludeFromMaxTransaction(address updAds, bool isEx) returns()
func (_Sniper *SniperSession) ExcludeFromMaxTransaction(updAds common.Address, isEx bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeFromMaxTransaction(&_Sniper.TransactOpts, updAds, isEx)
}

// ExcludeFromMaxTransaction is a paid mutator transaction binding the contract method 0x7571336a.
//
// Solidity: function excludeFromMaxTransaction(address updAds, bool isEx) returns()
func (_Sniper *SniperTransactorSession) ExcludeFromMaxTransaction(updAds common.Address, isEx bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeFromMaxTransaction(&_Sniper.TransactOpts, updAds, isEx)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Sniper *SniperTransactor) ExcludeMultipleAccountsFromFees(opts *bind.TransactOpts, accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "excludeMultipleAccountsFromFees", accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Sniper *SniperSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeMultipleAccountsFromFees(&_Sniper.TransactOpts, accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Sniper *SniperTransactorSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Sniper.Contract.ExcludeMultipleAccountsFromFees(&_Sniper.TransactOpts, accounts, excluded)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sniper *SniperTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sniper *SniperSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.IncreaseAllowance(&_Sniper.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sniper *SniperTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.IncreaseAllowance(&_Sniper.TransactOpts, spender, addedValue)
}

// ManualSwap is a paid mutator transaction binding the contract method 0x51bc3c85.
//
// Solidity: function manualSwap() returns()
func (_Sniper *SniperTransactor) ManualSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "manualSwap")
}

// ManualSwap is a paid mutator transaction binding the contract method 0x51bc3c85.
//
// Solidity: function manualSwap() returns()
func (_Sniper *SniperSession) ManualSwap() (*types.Transaction, error) {
	return _Sniper.Contract.ManualSwap(&_Sniper.TransactOpts)
}

// ManualSwap is a paid mutator transaction binding the contract method 0x51bc3c85.
//
// Solidity: function manualSwap() returns()
func (_Sniper *SniperTransactorSession) ManualSwap() (*types.Transaction, error) {
	return _Sniper.Contract.ManualSwap(&_Sniper.TransactOpts)
}

// RemoveLimits is a paid mutator transaction binding the contract method 0x751039fc.
//
// Solidity: function removeLimits() returns(bool)
func (_Sniper *SniperTransactor) RemoveLimits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "removeLimits")
}

// RemoveLimits is a paid mutator transaction binding the contract method 0x751039fc.
//
// Solidity: function removeLimits() returns(bool)
func (_Sniper *SniperSession) RemoveLimits() (*types.Transaction, error) {
	return _Sniper.Contract.RemoveLimits(&_Sniper.TransactOpts)
}

// RemoveLimits is a paid mutator transaction binding the contract method 0x751039fc.
//
// Solidity: function removeLimits() returns(bool)
func (_Sniper *SniperTransactorSession) RemoveLimits() (*types.Transaction, error) {
	return _Sniper.Contract.RemoveLimits(&_Sniper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sniper *SniperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sniper *SniperSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sniper.Contract.RenounceOwnership(&_Sniper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sniper *SniperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sniper.Contract.RenounceOwnership(&_Sniper.TransactOpts)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Sniper *SniperTransactor) SetAutomatedMarketMakerPair(opts *bind.TransactOpts, pair common.Address, value bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "setAutomatedMarketMakerPair", pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Sniper *SniperSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Sniper.Contract.SetAutomatedMarketMakerPair(&_Sniper.TransactOpts, pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Sniper *SniperTransactorSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Sniper.Contract.SetAutomatedMarketMakerPair(&_Sniper.TransactOpts, pair, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.Transfer(&_Sniper.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.Transfer(&_Sniper.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.TransferFrom(&_Sniper.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sniper *SniperTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.TransferFrom(&_Sniper.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sniper *SniperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sniper *SniperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sniper.Contract.TransferOwnership(&_Sniper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sniper *SniperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sniper.Contract.TransferOwnership(&_Sniper.TransactOpts, newOwner)
}

// UpdateBuyFees is a paid mutator transaction binding the contract method 0x71fc4688.
//
// Solidity: function updateBuyFees(uint256 _marketingBuyFee) returns()
func (_Sniper *SniperTransactor) UpdateBuyFees(opts *bind.TransactOpts, _marketingBuyFee *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateBuyFees", _marketingBuyFee)
}

// UpdateBuyFees is a paid mutator transaction binding the contract method 0x71fc4688.
//
// Solidity: function updateBuyFees(uint256 _marketingBuyFee) returns()
func (_Sniper *SniperSession) UpdateBuyFees(_marketingBuyFee *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateBuyFees(&_Sniper.TransactOpts, _marketingBuyFee)
}

// UpdateBuyFees is a paid mutator transaction binding the contract method 0x71fc4688.
//
// Solidity: function updateBuyFees(uint256 _marketingBuyFee) returns()
func (_Sniper *SniperTransactorSession) UpdateBuyFees(_marketingBuyFee *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateBuyFees(&_Sniper.TransactOpts, _marketingBuyFee)
}

// UpdateEarlyBuyBlockNum is a paid mutator transaction binding the contract method 0x8d140b3d.
//
// Solidity: function updateEarlyBuyBlockNum(uint256 extrablocks) returns()
func (_Sniper *SniperTransactor) UpdateEarlyBuyBlockNum(opts *bind.TransactOpts, extrablocks *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateEarlyBuyBlockNum", extrablocks)
}

// UpdateEarlyBuyBlockNum is a paid mutator transaction binding the contract method 0x8d140b3d.
//
// Solidity: function updateEarlyBuyBlockNum(uint256 extrablocks) returns()
func (_Sniper *SniperSession) UpdateEarlyBuyBlockNum(extrablocks *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateEarlyBuyBlockNum(&_Sniper.TransactOpts, extrablocks)
}

// UpdateEarlyBuyBlockNum is a paid mutator transaction binding the contract method 0x8d140b3d.
//
// Solidity: function updateEarlyBuyBlockNum(uint256 extrablocks) returns()
func (_Sniper *SniperTransactorSession) UpdateEarlyBuyBlockNum(extrablocks *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateEarlyBuyBlockNum(&_Sniper.TransactOpts, extrablocks)
}

// UpdateMaxTxnAmount is a paid mutator transaction binding the contract method 0x203e727e.
//
// Solidity: function updateMaxTxnAmount(uint256 newNum) returns()
func (_Sniper *SniperTransactor) UpdateMaxTxnAmount(opts *bind.TransactOpts, newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateMaxTxnAmount", newNum)
}

// UpdateMaxTxnAmount is a paid mutator transaction binding the contract method 0x203e727e.
//
// Solidity: function updateMaxTxnAmount(uint256 newNum) returns()
func (_Sniper *SniperSession) UpdateMaxTxnAmount(newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateMaxTxnAmount(&_Sniper.TransactOpts, newNum)
}

// UpdateMaxTxnAmount is a paid mutator transaction binding the contract method 0x203e727e.
//
// Solidity: function updateMaxTxnAmount(uint256 newNum) returns()
func (_Sniper *SniperTransactorSession) UpdateMaxTxnAmount(newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateMaxTxnAmount(&_Sniper.TransactOpts, newNum)
}

// UpdateMaxWalletAmount is a paid mutator transaction binding the contract method 0xc18bc195.
//
// Solidity: function updateMaxWalletAmount(uint256 newNum) returns()
func (_Sniper *SniperTransactor) UpdateMaxWalletAmount(opts *bind.TransactOpts, newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateMaxWalletAmount", newNum)
}

// UpdateMaxWalletAmount is a paid mutator transaction binding the contract method 0xc18bc195.
//
// Solidity: function updateMaxWalletAmount(uint256 newNum) returns()
func (_Sniper *SniperSession) UpdateMaxWalletAmount(newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateMaxWalletAmount(&_Sniper.TransactOpts, newNum)
}

// UpdateMaxWalletAmount is a paid mutator transaction binding the contract method 0xc18bc195.
//
// Solidity: function updateMaxWalletAmount(uint256 newNum) returns()
func (_Sniper *SniperTransactorSession) UpdateMaxWalletAmount(newNum *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateMaxWalletAmount(&_Sniper.TransactOpts, newNum)
}

// UpdateSellFees is a paid mutator transaction binding the contract method 0xeba4c333.
//
// Solidity: function updateSellFees(uint256 _marketingSellFee) returns()
func (_Sniper *SniperTransactor) UpdateSellFees(opts *bind.TransactOpts, _marketingSellFee *big.Int) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateSellFees", _marketingSellFee)
}

// UpdateSellFees is a paid mutator transaction binding the contract method 0xeba4c333.
//
// Solidity: function updateSellFees(uint256 _marketingSellFee) returns()
func (_Sniper *SniperSession) UpdateSellFees(_marketingSellFee *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateSellFees(&_Sniper.TransactOpts, _marketingSellFee)
}

// UpdateSellFees is a paid mutator transaction binding the contract method 0xeba4c333.
//
// Solidity: function updateSellFees(uint256 _marketingSellFee) returns()
func (_Sniper *SniperTransactorSession) UpdateSellFees(_marketingSellFee *big.Int) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateSellFees(&_Sniper.TransactOpts, _marketingSellFee)
}

// UpdateSwapEnabled is a paid mutator transaction binding the contract method 0x924de9b7.
//
// Solidity: function updateSwapEnabled(bool enabled) returns()
func (_Sniper *SniperTransactor) UpdateSwapEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "updateSwapEnabled", enabled)
}

// UpdateSwapEnabled is a paid mutator transaction binding the contract method 0x924de9b7.
//
// Solidity: function updateSwapEnabled(bool enabled) returns()
func (_Sniper *SniperSession) UpdateSwapEnabled(enabled bool) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateSwapEnabled(&_Sniper.TransactOpts, enabled)
}

// UpdateSwapEnabled is a paid mutator transaction binding the contract method 0x924de9b7.
//
// Solidity: function updateSwapEnabled(bool enabled) returns()
func (_Sniper *SniperTransactorSession) UpdateSwapEnabled(enabled bool) (*types.Transaction, error) {
	return _Sniper.Contract.UpdateSwapEnabled(&_Sniper.TransactOpts, enabled)
}

// WithdrawStuckEth is a paid mutator transaction binding the contract method 0x7fa787ba.
//
// Solidity: function withdrawStuckEth() returns()
func (_Sniper *SniperTransactor) WithdrawStuckEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.Transact(opts, "withdrawStuckEth")
}

// WithdrawStuckEth is a paid mutator transaction binding the contract method 0x7fa787ba.
//
// Solidity: function withdrawStuckEth() returns()
func (_Sniper *SniperSession) WithdrawStuckEth() (*types.Transaction, error) {
	return _Sniper.Contract.WithdrawStuckEth(&_Sniper.TransactOpts)
}

// WithdrawStuckEth is a paid mutator transaction binding the contract method 0x7fa787ba.
//
// Solidity: function withdrawStuckEth() returns()
func (_Sniper *SniperTransactorSession) WithdrawStuckEth() (*types.Transaction, error) {
	return _Sniper.Contract.WithdrawStuckEth(&_Sniper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sniper *SniperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sniper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sniper *SniperSession) Receive() (*types.Transaction, error) {
	return _Sniper.Contract.Receive(&_Sniper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sniper *SniperTransactorSession) Receive() (*types.Transaction, error) {
	return _Sniper.Contract.Receive(&_Sniper.TransactOpts)
}

// SniperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sniper contract.
type SniperApprovalIterator struct {
	Event *SniperApproval // Event containing the contract specifics and raw log

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
func (it *SniperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperApproval)
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
		it.Event = new(SniperApproval)
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
func (it *SniperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperApproval represents a Approval event raised by the Sniper contract.
type SniperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sniper *SniperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SniperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SniperApprovalIterator{contract: _Sniper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sniper *SniperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SniperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperApproval)
				if err := _Sniper.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Sniper *SniperFilterer) ParseApproval(log types.Log) (*SniperApproval, error) {
	event := new(SniperApproval)
	if err := _Sniper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperExcludeFromFeesIterator is returned from FilterExcludeFromFees and is used to iterate over the raw logs and unpacked data for ExcludeFromFees events raised by the Sniper contract.
type SniperExcludeFromFeesIterator struct {
	Event *SniperExcludeFromFees // Event containing the contract specifics and raw log

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
func (it *SniperExcludeFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperExcludeFromFees)
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
		it.Event = new(SniperExcludeFromFees)
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
func (it *SniperExcludeFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperExcludeFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperExcludeFromFees represents a ExcludeFromFees event raised by the Sniper contract.
type SniperExcludeFromFees struct {
	Account    common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeFromFees is a free log retrieval operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Sniper *SniperFilterer) FilterExcludeFromFees(opts *bind.FilterOpts, account []common.Address) (*SniperExcludeFromFeesIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return &SniperExcludeFromFeesIterator{contract: _Sniper.contract, event: "ExcludeFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeFromFees is a free log subscription operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Sniper *SniperFilterer) WatchExcludeFromFees(opts *bind.WatchOpts, sink chan<- *SniperExcludeFromFees, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperExcludeFromFees)
				if err := _Sniper.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
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

// ParseExcludeFromFees is a log parse operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Sniper *SniperFilterer) ParseExcludeFromFees(log types.Log) (*SniperExcludeFromFees, error) {
	event := new(SniperExcludeFromFees)
	if err := _Sniper.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperExcludeMultipleAccountsFromFeesIterator is returned from FilterExcludeMultipleAccountsFromFees and is used to iterate over the raw logs and unpacked data for ExcludeMultipleAccountsFromFees events raised by the Sniper contract.
type SniperExcludeMultipleAccountsFromFeesIterator struct {
	Event *SniperExcludeMultipleAccountsFromFees // Event containing the contract specifics and raw log

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
func (it *SniperExcludeMultipleAccountsFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperExcludeMultipleAccountsFromFees)
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
		it.Event = new(SniperExcludeMultipleAccountsFromFees)
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
func (it *SniperExcludeMultipleAccountsFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperExcludeMultipleAccountsFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperExcludeMultipleAccountsFromFees represents a ExcludeMultipleAccountsFromFees event raised by the Sniper contract.
type SniperExcludeMultipleAccountsFromFees struct {
	Accounts   []common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeMultipleAccountsFromFees is a free log retrieval operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Sniper *SniperFilterer) FilterExcludeMultipleAccountsFromFees(opts *bind.FilterOpts) (*SniperExcludeMultipleAccountsFromFeesIterator, error) {

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return &SniperExcludeMultipleAccountsFromFeesIterator{contract: _Sniper.contract, event: "ExcludeMultipleAccountsFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeMultipleAccountsFromFees is a free log subscription operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Sniper *SniperFilterer) WatchExcludeMultipleAccountsFromFees(opts *bind.WatchOpts, sink chan<- *SniperExcludeMultipleAccountsFromFees) (event.Subscription, error) {

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperExcludeMultipleAccountsFromFees)
				if err := _Sniper.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
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

// ParseExcludeMultipleAccountsFromFees is a log parse operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Sniper *SniperFilterer) ParseExcludeMultipleAccountsFromFees(log types.Log) (*SniperExcludeMultipleAccountsFromFees, error) {
	event := new(SniperExcludeMultipleAccountsFromFees)
	if err := _Sniper.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sniper contract.
type SniperOwnershipTransferredIterator struct {
	Event *SniperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SniperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperOwnershipTransferred)
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
		it.Event = new(SniperOwnershipTransferred)
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
func (it *SniperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperOwnershipTransferred represents a OwnershipTransferred event raised by the Sniper contract.
type SniperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sniper *SniperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SniperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SniperOwnershipTransferredIterator{contract: _Sniper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sniper *SniperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SniperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperOwnershipTransferred)
				if err := _Sniper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Sniper *SniperFilterer) ParseOwnershipTransferred(log types.Log) (*SniperOwnershipTransferred, error) {
	event := new(SniperOwnershipTransferred)
	if err := _Sniper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperSetAutomatedMarketMakerPairIterator is returned from FilterSetAutomatedMarketMakerPair and is used to iterate over the raw logs and unpacked data for SetAutomatedMarketMakerPair events raised by the Sniper contract.
type SniperSetAutomatedMarketMakerPairIterator struct {
	Event *SniperSetAutomatedMarketMakerPair // Event containing the contract specifics and raw log

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
func (it *SniperSetAutomatedMarketMakerPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperSetAutomatedMarketMakerPair)
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
		it.Event = new(SniperSetAutomatedMarketMakerPair)
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
func (it *SniperSetAutomatedMarketMakerPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperSetAutomatedMarketMakerPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperSetAutomatedMarketMakerPair represents a SetAutomatedMarketMakerPair event raised by the Sniper contract.
type SniperSetAutomatedMarketMakerPair struct {
	Pair  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAutomatedMarketMakerPair is a free log retrieval operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Sniper *SniperFilterer) FilterSetAutomatedMarketMakerPair(opts *bind.FilterOpts, pair []common.Address, value []bool) (*SniperSetAutomatedMarketMakerPairIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &SniperSetAutomatedMarketMakerPairIterator{contract: _Sniper.contract, event: "SetAutomatedMarketMakerPair", logs: logs, sub: sub}, nil
}

// WatchSetAutomatedMarketMakerPair is a free log subscription operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Sniper *SniperFilterer) WatchSetAutomatedMarketMakerPair(opts *bind.WatchOpts, sink chan<- *SniperSetAutomatedMarketMakerPair, pair []common.Address, value []bool) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperSetAutomatedMarketMakerPair)
				if err := _Sniper.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
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

// ParseSetAutomatedMarketMakerPair is a log parse operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Sniper *SniperFilterer) ParseSetAutomatedMarketMakerPair(log types.Log) (*SniperSetAutomatedMarketMakerPair, error) {
	event := new(SniperSetAutomatedMarketMakerPair)
	if err := _Sniper.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperSwapAndLiquifyIterator is returned from FilterSwapAndLiquify and is used to iterate over the raw logs and unpacked data for SwapAndLiquify events raised by the Sniper contract.
type SniperSwapAndLiquifyIterator struct {
	Event *SniperSwapAndLiquify // Event containing the contract specifics and raw log

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
func (it *SniperSwapAndLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperSwapAndLiquify)
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
		it.Event = new(SniperSwapAndLiquify)
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
func (it *SniperSwapAndLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperSwapAndLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperSwapAndLiquify represents a SwapAndLiquify event raised by the Sniper contract.
type SniperSwapAndLiquify struct {
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquify is a free log retrieval operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Sniper *SniperFilterer) FilterSwapAndLiquify(opts *bind.FilterOpts) (*SniperSwapAndLiquifyIterator, error) {

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return &SniperSwapAndLiquifyIterator{contract: _Sniper.contract, event: "SwapAndLiquify", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquify is a free log subscription operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Sniper *SniperFilterer) WatchSwapAndLiquify(opts *bind.WatchOpts, sink chan<- *SniperSwapAndLiquify) (event.Subscription, error) {

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperSwapAndLiquify)
				if err := _Sniper.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
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

// ParseSwapAndLiquify is a log parse operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Sniper *SniperFilterer) ParseSwapAndLiquify(log types.Log) (*SniperSwapAndLiquify, error) {
	event := new(SniperSwapAndLiquify)
	if err := _Sniper.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SniperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sniper contract.
type SniperTransferIterator struct {
	Event *SniperTransfer // Event containing the contract specifics and raw log

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
func (it *SniperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SniperTransfer)
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
		it.Event = new(SniperTransfer)
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
func (it *SniperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SniperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SniperTransfer represents a Transfer event raised by the Sniper contract.
type SniperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sniper *SniperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SniperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sniper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SniperTransferIterator{contract: _Sniper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sniper *SniperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SniperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sniper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SniperTransfer)
				if err := _Sniper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Sniper *SniperFilterer) ParseTransfer(log types.Log) (*SniperTransfer, error) {
	event := new(SniperTransfer)
	if err := _Sniper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
