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

// PixassetABI is the input ABI used to generate the binding from.
const PixassetABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contenthash\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"newToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"split\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contenthash\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"findTokenId\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"string\"},{\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Pixasset is an auto generated Go binding around an Ethereum contract.
type Pixasset struct {
	PixassetCaller     // Read-only binding to the contract
	PixassetTransactor // Write-only binding to the contract
	PixassetFilterer   // Log filterer for contract events
}

// PixassetCaller is an auto generated read-only Go binding around an Ethereum contract.
type PixassetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PixassetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PixassetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PixassetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PixassetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PixassetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PixassetSession struct {
	Contract     *Pixasset         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PixassetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PixassetCallerSession struct {
	Contract *PixassetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PixassetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PixassetTransactorSession struct {
	Contract     *PixassetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PixassetRaw is an auto generated low-level Go binding around an Ethereum contract.
type PixassetRaw struct {
	Contract *Pixasset // Generic contract binding to access the raw methods on
}

// PixassetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PixassetCallerRaw struct {
	Contract *PixassetCaller // Generic read-only contract binding to access the raw methods on
}

// PixassetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PixassetTransactorRaw struct {
	Contract *PixassetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPixasset creates a new instance of Pixasset, bound to a specific deployed contract.
func NewPixasset(address common.Address, backend bind.ContractBackend) (*Pixasset, error) {
	contract, err := bindPixasset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pixasset{PixassetCaller: PixassetCaller{contract: contract}, PixassetTransactor: PixassetTransactor{contract: contract}, PixassetFilterer: PixassetFilterer{contract: contract}}, nil
}

// NewPixassetCaller creates a new read-only instance of Pixasset, bound to a specific deployed contract.
func NewPixassetCaller(address common.Address, caller bind.ContractCaller) (*PixassetCaller, error) {
	contract, err := bindPixasset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PixassetCaller{contract: contract}, nil
}

// NewPixassetTransactor creates a new write-only instance of Pixasset, bound to a specific deployed contract.
func NewPixassetTransactor(address common.Address, transactor bind.ContractTransactor) (*PixassetTransactor, error) {
	contract, err := bindPixasset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PixassetTransactor{contract: contract}, nil
}

// NewPixassetFilterer creates a new log filterer instance of Pixasset, bound to a specific deployed contract.
func NewPixassetFilterer(address common.Address, filterer bind.ContractFilterer) (*PixassetFilterer, error) {
	contract, err := bindPixasset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PixassetFilterer{contract: contract}, nil
}

// bindPixasset binds a generic wrapper to an already deployed contract.
func bindPixasset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PixassetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pixasset *PixassetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pixasset.Contract.PixassetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pixasset *PixassetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pixasset.Contract.PixassetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pixasset *PixassetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pixasset.Contract.PixassetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pixasset *PixassetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pixasset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pixasset *PixassetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pixasset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pixasset *PixassetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pixasset.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pixasset *PixassetCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pixasset *PixassetSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pixasset.Contract.BalanceOf(&_Pixasset.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pixasset *PixassetCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pixasset.Contract.BalanceOf(&_Pixasset.CallOpts, _owner)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(a string, b string) constant returns(bool)
func (_Pixasset *PixassetCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "compareStrings", a, b)
	return *ret0, err
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(a string, b string) constant returns(bool)
func (_Pixasset *PixassetSession) CompareStrings(a string, b string) (bool, error) {
	return _Pixasset.Contract.CompareStrings(&_Pixasset.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(a string, b string) constant returns(bool)
func (_Pixasset *PixassetCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _Pixasset.Contract.CompareStrings(&_Pixasset.CallOpts, a, b)
}

// FindTokenId is a free data retrieval call binding the contract method 0xa14d9320.
//
// Solidity: function findTokenId(contenthash string, addr address) constant returns(tokenId uint256)
func (_Pixasset *PixassetCaller) FindTokenId(opts *bind.CallOpts, contenthash string, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "findTokenId", contenthash, addr)
	return *ret0, err
}

// FindTokenId is a free data retrieval call binding the contract method 0xa14d9320.
//
// Solidity: function findTokenId(contenthash string, addr address) constant returns(tokenId uint256)
func (_Pixasset *PixassetSession) FindTokenId(contenthash string, addr common.Address) (*big.Int, error) {
	return _Pixasset.Contract.FindTokenId(&_Pixasset.CallOpts, contenthash, addr)
}

// FindTokenId is a free data retrieval call binding the contract method 0xa14d9320.
//
// Solidity: function findTokenId(contenthash string, addr address) constant returns(tokenId uint256)
func (_Pixasset *PixassetCallerSession) FindTokenId(contenthash string, addr common.Address) (*big.Int, error) {
	return _Pixasset.Contract.FindTokenId(&_Pixasset.CallOpts, contenthash, addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pixasset *PixassetCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pixasset *PixassetSession) Name() (string, error) {
	return _Pixasset.Contract.Name(&_Pixasset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pixasset *PixassetCallerSession) Name() (string, error) {
	return _Pixasset.Contract.Name(&_Pixasset.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_Pixasset *PixassetCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_Pixasset *PixassetSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Pixasset.Contract.OwnerOf(&_Pixasset.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_Pixasset *PixassetCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Pixasset.Contract.OwnerOf(&_Pixasset.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pixasset *PixassetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pixasset *PixassetSession) Symbol() (string, error) {
	return _Pixasset.Contract.Symbol(&_Pixasset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pixasset *PixassetCallerSession) Symbol() (string, error) {
	return _Pixasset.Contract.Symbol(&_Pixasset.CallOpts)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owners address) constant returns(ownerTokens uint256[])
func (_Pixasset *PixassetCaller) TokensOfOwner(opts *bind.CallOpts, _owners common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "tokensOfOwner", _owners)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owners address) constant returns(ownerTokens uint256[])
func (_Pixasset *PixassetSession) TokensOfOwner(_owners common.Address) ([]*big.Int, error) {
	return _Pixasset.Contract.TokensOfOwner(&_Pixasset.CallOpts, _owners)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owners address) constant returns(ownerTokens uint256[])
func (_Pixasset *PixassetCallerSession) TokensOfOwner(_owners common.Address) ([]*big.Int, error) {
	return _Pixasset.Contract.TokensOfOwner(&_Pixasset.CallOpts, _owners)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(total uint256)
func (_Pixasset *PixassetCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pixasset.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(total uint256)
func (_Pixasset *PixassetSession) TotalSupply() (*big.Int, error) {
	return _Pixasset.Contract.TotalSupply(&_Pixasset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(total uint256)
func (_Pixasset *PixassetCallerSession) TotalSupply() (*big.Int, error) {
	return _Pixasset.Contract.TotalSupply(&_Pixasset.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Approve(&_Pixasset.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Approve(&_Pixasset.TransactOpts, _to, _tokenId)
}

// NewToken is a paid mutator transaction binding the contract method 0x4629ffea.
//
// Solidity: function newToken(contenthash string, metadata string) returns(uint256)
func (_Pixasset *PixassetTransactor) NewToken(opts *bind.TransactOpts, contenthash string, metadata string) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "newToken", contenthash, metadata)
}

// NewToken is a paid mutator transaction binding the contract method 0x4629ffea.
//
// Solidity: function newToken(contenthash string, metadata string) returns(uint256)
func (_Pixasset *PixassetSession) NewToken(contenthash string, metadata string) (*types.Transaction, error) {
	return _Pixasset.Contract.NewToken(&_Pixasset.TransactOpts, contenthash, metadata)
}

// NewToken is a paid mutator transaction binding the contract method 0x4629ffea.
//
// Solidity: function newToken(contenthash string, metadata string) returns(uint256)
func (_Pixasset *PixassetTransactorSession) NewToken(contenthash string, metadata string) (*types.Transaction, error) {
	return _Pixasset.Contract.NewToken(&_Pixasset.TransactOpts, contenthash, metadata)
}

// Split is a paid mutator transaction binding the contract method 0x4b19becc.
//
// Solidity: function split(tokenId uint256, weight uint256) returns()
func (_Pixasset *PixassetTransactor) Split(opts *bind.TransactOpts, tokenId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "split", tokenId, weight)
}

// Split is a paid mutator transaction binding the contract method 0x4b19becc.
//
// Solidity: function split(tokenId uint256, weight uint256) returns()
func (_Pixasset *PixassetSession) Split(tokenId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Split(&_Pixasset.TransactOpts, tokenId, weight)
}

// Split is a paid mutator transaction binding the contract method 0x4b19becc.
//
// Solidity: function split(tokenId uint256, weight uint256) returns()
func (_Pixasset *PixassetTransactorSession) Split(tokenId *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Split(&_Pixasset.TransactOpts, tokenId, weight)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Transfer(&_Pixasset.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.Transfer(&_Pixasset.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pixasset *PixassetSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.TransferFrom(&_Pixasset.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pixasset *PixassetTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pixasset.Contract.TransferFrom(&_Pixasset.TransactOpts, _from, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pixasset *PixassetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pixasset.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pixasset *PixassetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pixasset.Contract.TransferOwnership(&_Pixasset.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pixasset *PixassetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pixasset.Contract.TransferOwnership(&_Pixasset.TransactOpts, newOwner)
}

// PixassetApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pixasset contract.
type PixassetApprovalIterator struct {
	Event *PixassetApproval // Event containing the contract specifics and raw log

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
func (it *PixassetApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PixassetApproval)
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
		it.Event = new(PixassetApproval)
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
func (it *PixassetApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PixassetApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PixassetApproval represents a Approval event raised by the Pixasset contract.
type PixassetApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner address, approved address, tokenId uint256)
func (_Pixasset *PixassetFilterer) FilterApproval(opts *bind.FilterOpts) (*PixassetApprovalIterator, error) {

	logs, sub, err := _Pixasset.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &PixassetApprovalIterator{contract: _Pixasset.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner address, approved address, tokenId uint256)
func (_Pixasset *PixassetFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PixassetApproval) (event.Subscription, error) {

	logs, sub, err := _Pixasset.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PixassetApproval)
				if err := _Pixasset.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PixassetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pixasset contract.
type PixassetTransferIterator struct {
	Event *PixassetTransfer // Event containing the contract specifics and raw log

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
func (it *PixassetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PixassetTransfer)
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
		it.Event = new(PixassetTransfer)
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
func (it *PixassetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PixassetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PixassetTransfer represents a Transfer event raised by the Pixasset contract.
type PixassetTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenId uint256)
func (_Pixasset *PixassetFilterer) FilterTransfer(opts *bind.FilterOpts) (*PixassetTransferIterator, error) {

	logs, sub, err := _Pixasset.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &PixassetTransferIterator{contract: _Pixasset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenId uint256)
func (_Pixasset *PixassetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PixassetTransfer) (event.Subscription, error) {

	logs, sub, err := _Pixasset.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PixassetTransfer)
				if err := _Pixasset.contract.UnpackLog(event, "Transfer", log); err != nil {
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
