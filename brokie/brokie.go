// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brokie

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

// BrokieMetaData contains all meta data concerning the Brokie contract.
var BrokieMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBOG\",\"type\":\"uint256\"}],\"name\":\"AutoLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_maxWalletAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearStuckBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIDEXRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketingFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPercent\",\"type\":\"uint256\"}],\"name\":\"setWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BrokieABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokieMetaData.ABI instead.
var BrokieABI = BrokieMetaData.ABI

// Brokie is an auto generated Go binding around an Ethereum contract.
type Brokie struct {
	BrokieCaller     // Read-only binding to the contract
	BrokieTransactor // Write-only binding to the contract
	BrokieFilterer   // Log filterer for contract events
}

// BrokieCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokieSession struct {
	Contract     *Brokie           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokieCallerSession struct {
	Contract *BrokieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokieTransactorSession struct {
	Contract     *BrokieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokieRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokieRaw struct {
	Contract *Brokie // Generic contract binding to access the raw methods on
}

// BrokieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokieCallerRaw struct {
	Contract *BrokieCaller // Generic read-only contract binding to access the raw methods on
}

// BrokieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokieTransactorRaw struct {
	Contract *BrokieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrokie creates a new instance of Brokie, bound to a specific deployed contract.
func NewBrokie(address common.Address, backend bind.ContractBackend) (*Brokie, error) {
	contract, err := bindBrokie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brokie{BrokieCaller: BrokieCaller{contract: contract}, BrokieTransactor: BrokieTransactor{contract: contract}, BrokieFilterer: BrokieFilterer{contract: contract}}, nil
}

// NewBrokieCaller creates a new read-only instance of Brokie, bound to a specific deployed contract.
func NewBrokieCaller(address common.Address, caller bind.ContractCaller) (*BrokieCaller, error) {
	contract, err := bindBrokie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokieCaller{contract: contract}, nil
}

// NewBrokieTransactor creates a new write-only instance of Brokie, bound to a specific deployed contract.
func NewBrokieTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokieTransactor, error) {
	contract, err := bindBrokie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokieTransactor{contract: contract}, nil
}

// NewBrokieFilterer creates a new log filterer instance of Brokie, bound to a specific deployed contract.
func NewBrokieFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokieFilterer, error) {
	contract, err := bindBrokie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokieFilterer{contract: contract}, nil
}

// bindBrokie binds a generic wrapper to an already deployed contract.
func bindBrokie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokieABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brokie *BrokieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brokie.Contract.BrokieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brokie *BrokieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brokie.Contract.BrokieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brokie *BrokieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brokie.Contract.BrokieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brokie *BrokieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brokie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brokie *BrokieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brokie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brokie *BrokieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brokie.Contract.contract.Transact(opts, method, params...)
}

// MaxWalletAmount is a free data retrieval call binding the contract method 0x6c0a24eb.
//
// Solidity: function _maxWalletAmount() view returns(uint256)
func (_Brokie *BrokieCaller) MaxWalletAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "_maxWalletAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWalletAmount is a free data retrieval call binding the contract method 0x6c0a24eb.
//
// Solidity: function _maxWalletAmount() view returns(uint256)
func (_Brokie *BrokieSession) MaxWalletAmount() (*big.Int, error) {
	return _Brokie.Contract.MaxWalletAmount(&_Brokie.CallOpts)
}

// MaxWalletAmount is a free data retrieval call binding the contract method 0x6c0a24eb.
//
// Solidity: function _maxWalletAmount() view returns(uint256)
func (_Brokie *BrokieCallerSession) MaxWalletAmount() (*big.Int, error) {
	return _Brokie.Contract.MaxWalletAmount(&_Brokie.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Brokie *BrokieCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Brokie *BrokieSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Brokie.Contract.Allowance(&_Brokie.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Brokie *BrokieCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Brokie.Contract.Allowance(&_Brokie.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brokie *BrokieCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brokie *BrokieSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Brokie.Contract.BalanceOf(&_Brokie.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brokie *BrokieCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Brokie.Contract.BalanceOf(&_Brokie.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Brokie *BrokieCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Brokie *BrokieSession) Decimals() (uint8, error) {
	return _Brokie.Contract.Decimals(&_Brokie.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Brokie *BrokieCallerSession) Decimals() (uint8, error) {
	return _Brokie.Contract.Decimals(&_Brokie.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Brokie *BrokieCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Brokie *BrokieSession) GetOwner() (common.Address, error) {
	return _Brokie.Contract.GetOwner(&_Brokie.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Brokie *BrokieCallerSession) GetOwner() (common.Address, error) {
	return _Brokie.Contract.GetOwner(&_Brokie.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Brokie *BrokieCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Brokie *BrokieSession) IsOwner(account common.Address) (bool, error) {
	return _Brokie.Contract.IsOwner(&_Brokie.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Brokie *BrokieCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Brokie.Contract.IsOwner(&_Brokie.CallOpts, account)
}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_Brokie *BrokieCaller) MarketingFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "marketingFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_Brokie *BrokieSession) MarketingFeeReceiver() (common.Address, error) {
	return _Brokie.Contract.MarketingFeeReceiver(&_Brokie.CallOpts)
}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_Brokie *BrokieCallerSession) MarketingFeeReceiver() (common.Address, error) {
	return _Brokie.Contract.MarketingFeeReceiver(&_Brokie.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Brokie *BrokieCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Brokie *BrokieSession) Name() (string, error) {
	return _Brokie.Contract.Name(&_Brokie.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Brokie *BrokieCallerSession) Name() (string, error) {
	return _Brokie.Contract.Name(&_Brokie.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Brokie *BrokieCaller) Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Brokie *BrokieSession) Pair() (common.Address, error) {
	return _Brokie.Contract.Pair(&_Brokie.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Brokie *BrokieCallerSession) Pair() (common.Address, error) {
	return _Brokie.Contract.Pair(&_Brokie.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Brokie *BrokieCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Brokie *BrokieSession) Router() (common.Address, error) {
	return _Brokie.Contract.Router(&_Brokie.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Brokie *BrokieCallerSession) Router() (common.Address, error) {
	return _Brokie.Contract.Router(&_Brokie.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Brokie *BrokieCaller) SwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "swapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Brokie *BrokieSession) SwapEnabled() (bool, error) {
	return _Brokie.Contract.SwapEnabled(&_Brokie.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_Brokie *BrokieCallerSession) SwapEnabled() (bool, error) {
	return _Brokie.Contract.SwapEnabled(&_Brokie.CallOpts)
}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Brokie *BrokieCaller) SwapThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "swapThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Brokie *BrokieSession) SwapThreshold() (*big.Int, error) {
	return _Brokie.Contract.SwapThreshold(&_Brokie.CallOpts)
}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Brokie *BrokieCallerSession) SwapThreshold() (*big.Int, error) {
	return _Brokie.Contract.SwapThreshold(&_Brokie.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Brokie *BrokieCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Brokie *BrokieSession) Symbol() (string, error) {
	return _Brokie.Contract.Symbol(&_Brokie.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Brokie *BrokieCallerSession) Symbol() (string, error) {
	return _Brokie.Contract.Symbol(&_Brokie.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brokie *BrokieCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brokie.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brokie *BrokieSession) TotalSupply() (*big.Int, error) {
	return _Brokie.Contract.TotalSupply(&_Brokie.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brokie *BrokieCallerSession) TotalSupply() (*big.Int, error) {
	return _Brokie.Contract.TotalSupply(&_Brokie.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Brokie *BrokieSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.Approve(&_Brokie.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.Approve(&_Brokie.TransactOpts, spender, amount)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_Brokie *BrokieTransactor) ApproveMax(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "approveMax", spender)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_Brokie *BrokieSession) ApproveMax(spender common.Address) (*types.Transaction, error) {
	return _Brokie.Contract.ApproveMax(&_Brokie.TransactOpts, spender)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address spender) returns(bool)
func (_Brokie *BrokieTransactorSession) ApproveMax(spender common.Address) (*types.Transaction, error) {
	return _Brokie.Contract.ApproveMax(&_Brokie.TransactOpts, spender)
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x364333f4.
//
// Solidity: function clearStuckBalance() returns()
func (_Brokie *BrokieTransactor) ClearStuckBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "clearStuckBalance")
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x364333f4.
//
// Solidity: function clearStuckBalance() returns()
func (_Brokie *BrokieSession) ClearStuckBalance() (*types.Transaction, error) {
	return _Brokie.Contract.ClearStuckBalance(&_Brokie.TransactOpts)
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x364333f4.
//
// Solidity: function clearStuckBalance() returns()
func (_Brokie *BrokieTransactorSession) ClearStuckBalance() (*types.Transaction, error) {
	return _Brokie.Contract.ClearStuckBalance(&_Brokie.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brokie *BrokieTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brokie *BrokieSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brokie.Contract.RenounceOwnership(&_Brokie.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brokie *BrokieTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brokie.Contract.RenounceOwnership(&_Brokie.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _liquidityFee, uint256 _marketingFee) returns()
func (_Brokie *BrokieTransactor) SetFee(opts *bind.TransactOpts, _liquidityFee *big.Int, _marketingFee *big.Int) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "setFee", _liquidityFee, _marketingFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _liquidityFee, uint256 _marketingFee) returns()
func (_Brokie *BrokieSession) SetFee(_liquidityFee *big.Int, _marketingFee *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.SetFee(&_Brokie.TransactOpts, _liquidityFee, _marketingFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _liquidityFee, uint256 _marketingFee) returns()
func (_Brokie *BrokieTransactorSession) SetFee(_liquidityFee *big.Int, _marketingFee *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.SetFee(&_Brokie.TransactOpts, _liquidityFee, _marketingFee)
}

// SetWalletLimit is a paid mutator transaction binding the contract method 0xf1d5f517.
//
// Solidity: function setWalletLimit(uint256 amountPercent) returns()
func (_Brokie *BrokieTransactor) SetWalletLimit(opts *bind.TransactOpts, amountPercent *big.Int) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "setWalletLimit", amountPercent)
}

// SetWalletLimit is a paid mutator transaction binding the contract method 0xf1d5f517.
//
// Solidity: function setWalletLimit(uint256 amountPercent) returns()
func (_Brokie *BrokieSession) SetWalletLimit(amountPercent *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.SetWalletLimit(&_Brokie.TransactOpts, amountPercent)
}

// SetWalletLimit is a paid mutator transaction binding the contract method 0xf1d5f517.
//
// Solidity: function setWalletLimit(uint256 amountPercent) returns()
func (_Brokie *BrokieTransactorSession) SetWalletLimit(amountPercent *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.SetWalletLimit(&_Brokie.TransactOpts, amountPercent)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.Transfer(&_Brokie.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.Transfer(&_Brokie.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.TransferFrom(&_Brokie.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Brokie *BrokieTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Brokie.Contract.TransferFrom(&_Brokie.TransactOpts, sender, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Brokie *BrokieTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brokie.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Brokie *BrokieSession) Receive() (*types.Transaction, error) {
	return _Brokie.Contract.Receive(&_Brokie.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Brokie *BrokieTransactorSession) Receive() (*types.Transaction, error) {
	return _Brokie.Contract.Receive(&_Brokie.TransactOpts)
}

// BrokieApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Brokie contract.
type BrokieApprovalIterator struct {
	Event *BrokieApproval // Event containing the contract specifics and raw log

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
func (it *BrokieApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokieApproval)
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
		it.Event = new(BrokieApproval)
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
func (it *BrokieApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokieApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokieApproval represents a Approval event raised by the Brokie contract.
type BrokieApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Brokie *BrokieFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BrokieApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Brokie.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BrokieApprovalIterator{contract: _Brokie.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Brokie *BrokieFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BrokieApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Brokie.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokieApproval)
				if err := _Brokie.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Brokie *BrokieFilterer) ParseApproval(log types.Log) (*BrokieApproval, error) {
	event := new(BrokieApproval)
	if err := _Brokie.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokieAutoLiquifyIterator is returned from FilterAutoLiquify and is used to iterate over the raw logs and unpacked data for AutoLiquify events raised by the Brokie contract.
type BrokieAutoLiquifyIterator struct {
	Event *BrokieAutoLiquify // Event containing the contract specifics and raw log

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
func (it *BrokieAutoLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokieAutoLiquify)
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
		it.Event = new(BrokieAutoLiquify)
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
func (it *BrokieAutoLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokieAutoLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokieAutoLiquify represents a AutoLiquify event raised by the Brokie contract.
type BrokieAutoLiquify struct {
	AmountETH *big.Int
	AmountBOG *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAutoLiquify is a free log retrieval operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountETH, uint256 amountBOG)
func (_Brokie *BrokieFilterer) FilterAutoLiquify(opts *bind.FilterOpts) (*BrokieAutoLiquifyIterator, error) {

	logs, sub, err := _Brokie.contract.FilterLogs(opts, "AutoLiquify")
	if err != nil {
		return nil, err
	}
	return &BrokieAutoLiquifyIterator{contract: _Brokie.contract, event: "AutoLiquify", logs: logs, sub: sub}, nil
}

// WatchAutoLiquify is a free log subscription operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountETH, uint256 amountBOG)
func (_Brokie *BrokieFilterer) WatchAutoLiquify(opts *bind.WatchOpts, sink chan<- *BrokieAutoLiquify) (event.Subscription, error) {

	logs, sub, err := _Brokie.contract.WatchLogs(opts, "AutoLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokieAutoLiquify)
				if err := _Brokie.contract.UnpackLog(event, "AutoLiquify", log); err != nil {
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

// ParseAutoLiquify is a log parse operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountETH, uint256 amountBOG)
func (_Brokie *BrokieFilterer) ParseAutoLiquify(log types.Log) (*BrokieAutoLiquify, error) {
	event := new(BrokieAutoLiquify)
	if err := _Brokie.contract.UnpackLog(event, "AutoLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokieOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Brokie contract.
type BrokieOwnershipTransferredIterator struct {
	Event *BrokieOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrokieOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokieOwnershipTransferred)
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
		it.Event = new(BrokieOwnershipTransferred)
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
func (it *BrokieOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokieOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokieOwnershipTransferred represents a OwnershipTransferred event raised by the Brokie contract.
type BrokieOwnershipTransferred struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address owner)
func (_Brokie *BrokieFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*BrokieOwnershipTransferredIterator, error) {

	logs, sub, err := _Brokie.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &BrokieOwnershipTransferredIterator{contract: _Brokie.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address owner)
func (_Brokie *BrokieFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrokieOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Brokie.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokieOwnershipTransferred)
				if err := _Brokie.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address owner)
func (_Brokie *BrokieFilterer) ParseOwnershipTransferred(log types.Log) (*BrokieOwnershipTransferred, error) {
	event := new(BrokieOwnershipTransferred)
	if err := _Brokie.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokieTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Brokie contract.
type BrokieTransferIterator struct {
	Event *BrokieTransfer // Event containing the contract specifics and raw log

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
func (it *BrokieTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokieTransfer)
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
		it.Event = new(BrokieTransfer)
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
func (it *BrokieTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokieTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokieTransfer represents a Transfer event raised by the Brokie contract.
type BrokieTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Brokie *BrokieFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BrokieTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brokie.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrokieTransferIterator{contract: _Brokie.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Brokie *BrokieFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BrokieTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brokie.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokieTransfer)
				if err := _Brokie.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Brokie *BrokieFilterer) ParseTransfer(log types.Log) (*BrokieTransfer, error) {
	event := new(BrokieTransfer)
	if err := _Brokie.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
