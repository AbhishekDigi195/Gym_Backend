// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.ConvertType
)

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"membershipId\",\"type\":\"uint256\"}],\"name\":\"MembershipActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"membershipId\",\"type\":\"uint256\"}],\"name\":\"MembershipCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"membershipId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"membershipType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"MembershipPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"membershipId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEndDate\",\"type\":\"uint256\"}],\"name\":\"MembershipRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"UserDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contact\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"UserVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_membershipId\",\"type\":\"uint256\"}],\"name\":\"activateMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_membershipId\",\"type\":\"uint256\"}],\"name\":\"cancelMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"checkMembershipStatus\",\"outputs\":[{\"internalType\":\"enumGymMembership.MembershipStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"deleteUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_membershipId\",\"type\":\"uint256\"}],\"name\":\"getMembershipDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumGymMembership.MembershipStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membershipCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberships\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"membershipType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"enumGymMembership.MembershipStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_membershipType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endDate\",\"type\":\"uint256\"}],\"name\":\"purchaseMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contact\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_membershipId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newEndDate\",\"type\":\"uint256\"}],\"name\":\"renewMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contact\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"membershipId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// CheckMembershipStatus is a free data retrieval call binding the contract method 0xb211cb89.
//
// Solidity: function checkMembershipStatus(string _email) view returns(uint8)
func (_Storage *StorageCaller) CheckMembershipStatus(opts *bind.CallOpts, _email string) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "checkMembershipStatus", _email)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CheckMembershipStatus is a free data retrieval call binding the contract method 0xb211cb89.
//
// Solidity: function checkMembershipStatus(string _email) view returns(uint8)
func (_Storage *StorageSession) CheckMembershipStatus(_email string) (uint8, error) {
	return _Storage.Contract.CheckMembershipStatus(&_Storage.CallOpts, _email)
}

// CheckMembershipStatus is a free data retrieval call binding the contract method 0xb211cb89.
//
// Solidity: function checkMembershipStatus(string _email) view returns(uint8)
func (_Storage *StorageCallerSession) CheckMembershipStatus(_email string) (uint8, error) {
	return _Storage.Contract.CheckMembershipStatus(&_Storage.CallOpts, _email)
}

// GetMembershipDetails is a free data retrieval call binding the contract method 0xbba1568c.
//
// Solidity: function getMembershipDetails(uint256 _membershipId) view returns(string, string, uint256, uint256, uint8)
func (_Storage *StorageCaller) GetMembershipDetails(opts *bind.CallOpts, _membershipId *big.Int) (string, string, *big.Int, *big.Int, uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getMembershipDetails", _membershipId)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return out0, out1, out2, out3, out4, err

}

// GetMembershipDetails is a free data retrieval call binding the contract method 0xbba1568c.
//
// Solidity: function getMembershipDetails(uint256 _membershipId) view returns(string, string, uint256, uint256, uint8)
func (_Storage *StorageSession) GetMembershipDetails(_membershipId *big.Int) (string, string, *big.Int, *big.Int, uint8, error) {
	return _Storage.Contract.GetMembershipDetails(&_Storage.CallOpts, _membershipId)
}

// GetMembershipDetails is a free data retrieval call binding the contract method 0xbba1568c.
//
// Solidity: function getMembershipDetails(uint256 _membershipId) view returns(string, string, uint256, uint256, uint8)
func (_Storage *StorageCallerSession) GetMembershipDetails(_membershipId *big.Int) (string, string, *big.Int, *big.Int, uint8, error) {
	return _Storage.Contract.GetMembershipDetails(&_Storage.CallOpts, _membershipId)
}

// MembershipCounter is a free data retrieval call binding the contract method 0x20dae4cf.
//
// Solidity: function membershipCounter() view returns(uint256)
func (_Storage *StorageCaller) MembershipCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "membershipCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MembershipCounter is a free data retrieval call binding the contract method 0x20dae4cf.
//
// Solidity: function membershipCounter() view returns(uint256)
func (_Storage *StorageSession) MembershipCounter() (*big.Int, error) {
	return _Storage.Contract.MembershipCounter(&_Storage.CallOpts)
}

// MembershipCounter is a free data retrieval call binding the contract method 0x20dae4cf.
//
// Solidity: function membershipCounter() view returns(uint256)
func (_Storage *StorageCallerSession) MembershipCounter() (*big.Int, error) {
	return _Storage.Contract.MembershipCounter(&_Storage.CallOpts)
}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint256 id, string email, string membershipType, uint256 startDate, uint256 endDate, uint8 status)
func (_Storage *StorageCaller) Memberships(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id             *big.Int
	Email          string
	MembershipType string
	StartDate      *big.Int
	EndDate        *big.Int
	Status         uint8
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "memberships", arg0)

	outstruct := new(struct {
		Id             *big.Int
		Email          string
		MembershipType string
		StartDate      *big.Int
		EndDate        *big.Int
		Status         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Email = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MembershipType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.StartDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint256 id, string email, string membershipType, uint256 startDate, uint256 endDate, uint8 status)
func (_Storage *StorageSession) Memberships(arg0 *big.Int) (struct {
	Id             *big.Int
	Email          string
	MembershipType string
	StartDate      *big.Int
	EndDate        *big.Int
	Status         uint8
}, error) {
	return _Storage.Contract.Memberships(&_Storage.CallOpts, arg0)
}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint256 id, string email, string membershipType, uint256 startDate, uint256 endDate, uint8 status)
func (_Storage *StorageCallerSession) Memberships(arg0 *big.Int) (struct {
	Id             *big.Int
	Email          string
	MembershipType string
	StartDate      *big.Int
	EndDate        *big.Int
	Status         uint8
}, error) {
	return _Storage.Contract.Memberships(&_Storage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x689e1c03.
//
// Solidity: function users(string ) view returns(string name, string email, string contact, uint256 membershipId)
func (_Storage *StorageCaller) Users(opts *bind.CallOpts, arg0 string) (struct {
	Name         string
	Email        string
	Contact      string
	MembershipId *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Name         string
		Email        string
		Contact      string
		MembershipId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Contact = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.MembershipId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x689e1c03.
//
// Solidity: function users(string ) view returns(string name, string email, string contact, uint256 membershipId)
func (_Storage *StorageSession) Users(arg0 string) (struct {
	Name         string
	Email        string
	Contact      string
	MembershipId *big.Int
}, error) {
	return _Storage.Contract.Users(&_Storage.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x689e1c03.
//
// Solidity: function users(string ) view returns(string name, string email, string contact, uint256 membershipId)
func (_Storage *StorageCallerSession) Users(arg0 string) (struct {
	Name         string
	Email        string
	Contact      string
	MembershipId *big.Int
}, error) {
	return _Storage.Contract.Users(&_Storage.CallOpts, arg0)
}

// ActivateMembership is a paid mutator transaction binding the contract method 0x32ea2363.
//
// Solidity: function activateMembership(uint256 _membershipId) returns()
func (_Storage *StorageTransactor) ActivateMembership(opts *bind.TransactOpts, _membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "activateMembership", _membershipId)
}

// ActivateMembership is a paid mutator transaction binding the contract method 0x32ea2363.
//
// Solidity: function activateMembership(uint256 _membershipId) returns()
func (_Storage *StorageSession) ActivateMembership(_membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ActivateMembership(&_Storage.TransactOpts, _membershipId)
}

// ActivateMembership is a paid mutator transaction binding the contract method 0x32ea2363.
//
// Solidity: function activateMembership(uint256 _membershipId) returns()
func (_Storage *StorageTransactorSession) ActivateMembership(_membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ActivateMembership(&_Storage.TransactOpts, _membershipId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 _membershipId) returns()
func (_Storage *StorageTransactor) CancelMembership(opts *bind.TransactOpts, _membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "cancelMembership", _membershipId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 _membershipId) returns()
func (_Storage *StorageSession) CancelMembership(_membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.CancelMembership(&_Storage.TransactOpts, _membershipId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 _membershipId) returns()
func (_Storage *StorageTransactorSession) CancelMembership(_membershipId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.CancelMembership(&_Storage.TransactOpts, _membershipId)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xd3695161.
//
// Solidity: function deleteUser(string _email) returns()
func (_Storage *StorageTransactor) DeleteUser(opts *bind.TransactOpts, _email string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteUser", _email)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xd3695161.
//
// Solidity: function deleteUser(string _email) returns()
func (_Storage *StorageSession) DeleteUser(_email string) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUser(&_Storage.TransactOpts, _email)
}

// DeleteUser is a paid mutator transaction binding the contract method 0xd3695161.
//
// Solidity: function deleteUser(string _email) returns()
func (_Storage *StorageTransactorSession) DeleteUser(_email string) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUser(&_Storage.TransactOpts, _email)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// PurchaseMembership is a paid mutator transaction binding the contract method 0x31e8729f.
//
// Solidity: function purchaseMembership(string _email, string _membershipType, uint256 _startDate, uint256 _endDate) returns()
func (_Storage *StorageTransactor) PurchaseMembership(opts *bind.TransactOpts, _email string, _membershipType string, _startDate *big.Int, _endDate *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "purchaseMembership", _email, _membershipType, _startDate, _endDate)
}

// PurchaseMembership is a paid mutator transaction binding the contract method 0x31e8729f.
//
// Solidity: function purchaseMembership(string _email, string _membershipType, uint256 _startDate, uint256 _endDate) returns()
func (_Storage *StorageSession) PurchaseMembership(_email string, _membershipType string, _startDate *big.Int, _endDate *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.PurchaseMembership(&_Storage.TransactOpts, _email, _membershipType, _startDate, _endDate)
}

// PurchaseMembership is a paid mutator transaction binding the contract method 0x31e8729f.
//
// Solidity: function purchaseMembership(string _email, string _membershipType, uint256 _startDate, uint256 _endDate) returns()
func (_Storage *StorageTransactorSession) PurchaseMembership(_email string, _membershipType string, _startDate *big.Int, _endDate *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.PurchaseMembership(&_Storage.TransactOpts, _email, _membershipType, _startDate, _endDate)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xd637dcfa.
//
// Solidity: function registerUser(string _name, string _email, string _contact) returns()
func (_Storage *StorageTransactor) RegisterUser(opts *bind.TransactOpts, _name string, _email string, _contact string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "registerUser", _name, _email, _contact)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xd637dcfa.
//
// Solidity: function registerUser(string _name, string _email, string _contact) returns()
func (_Storage *StorageSession) RegisterUser(_name string, _email string, _contact string) (*types.Transaction, error) {
	return _Storage.Contract.RegisterUser(&_Storage.TransactOpts, _name, _email, _contact)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xd637dcfa.
//
// Solidity: function registerUser(string _name, string _email, string _contact) returns()
func (_Storage *StorageTransactorSession) RegisterUser(_name string, _email string, _contact string) (*types.Transaction, error) {
	return _Storage.Contract.RegisterUser(&_Storage.TransactOpts, _name, _email, _contact)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x8c99f69a.
//
// Solidity: function renewMembership(uint256 _membershipId, uint256 _newEndDate) returns()
func (_Storage *StorageTransactor) RenewMembership(opts *bind.TransactOpts, _membershipId *big.Int, _newEndDate *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renewMembership", _membershipId, _newEndDate)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x8c99f69a.
//
// Solidity: function renewMembership(uint256 _membershipId, uint256 _newEndDate) returns()
func (_Storage *StorageSession) RenewMembership(_membershipId *big.Int, _newEndDate *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RenewMembership(&_Storage.TransactOpts, _membershipId, _newEndDate)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x8c99f69a.
//
// Solidity: function renewMembership(uint256 _membershipId, uint256 _newEndDate) returns()
func (_Storage *StorageTransactorSession) RenewMembership(_membershipId *big.Int, _newEndDate *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RenewMembership(&_Storage.TransactOpts, _membershipId, _newEndDate)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xe600c817.
//
// Solidity: function verifyUser(string _email) returns(bool)
func (_Storage *StorageTransactor) VerifyUser(opts *bind.TransactOpts, _email string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "verifyUser", _email)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xe600c817.
//
// Solidity: function verifyUser(string _email) returns(bool)
func (_Storage *StorageSession) VerifyUser(_email string) (*types.Transaction, error) {
	return _Storage.Contract.VerifyUser(&_Storage.TransactOpts, _email)
}

// VerifyUser is a paid mutator transaction binding the contract method 0xe600c817.
//
// Solidity: function verifyUser(string _email) returns(bool)
func (_Storage *StorageTransactorSession) VerifyUser(_email string) (*types.Transaction, error) {
	return _Storage.Contract.VerifyUser(&_Storage.TransactOpts, _email)
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMembershipActivatedIterator is returned from FilterMembershipActivated and is used to iterate over the raw logs and unpacked data for MembershipActivated events raised by the Storage contract.
type StorageMembershipActivatedIterator struct {
	Event *StorageMembershipActivated // Event containing the contract specifics and raw log

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
func (it *StorageMembershipActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMembershipActivated)
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
		it.Event = new(StorageMembershipActivated)
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
func (it *StorageMembershipActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMembershipActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMembershipActivated represents a MembershipActivated event raised by the Storage contract.
type StorageMembershipActivated struct {
	Email        string
	MembershipId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMembershipActivated is a free log retrieval operation binding the contract event 0xd497ec759e28fc740b28b2e984199dec07fa10bcfd7928d3454990cef6d5d0ff.
//
// Solidity: event MembershipActivated(string email, uint256 membershipId)
func (_Storage *StorageFilterer) FilterMembershipActivated(opts *bind.FilterOpts) (*StorageMembershipActivatedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MembershipActivated")
	if err != nil {
		return nil, err
	}
	return &StorageMembershipActivatedIterator{contract: _Storage.contract, event: "MembershipActivated", logs: logs, sub: sub}, nil
}

// WatchMembershipActivated is a free log subscription operation binding the contract event 0xd497ec759e28fc740b28b2e984199dec07fa10bcfd7928d3454990cef6d5d0ff.
//
// Solidity: event MembershipActivated(string email, uint256 membershipId)
func (_Storage *StorageFilterer) WatchMembershipActivated(opts *bind.WatchOpts, sink chan<- *StorageMembershipActivated) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MembershipActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMembershipActivated)
				if err := _Storage.contract.UnpackLog(event, "MembershipActivated", log); err != nil {
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

// ParseMembershipActivated is a log parse operation binding the contract event 0xd497ec759e28fc740b28b2e984199dec07fa10bcfd7928d3454990cef6d5d0ff.
//
// Solidity: event MembershipActivated(string email, uint256 membershipId)
func (_Storage *StorageFilterer) ParseMembershipActivated(log types.Log) (*StorageMembershipActivated, error) {
	event := new(StorageMembershipActivated)
	if err := _Storage.contract.UnpackLog(event, "MembershipActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMembershipCancelledIterator is returned from FilterMembershipCancelled and is used to iterate over the raw logs and unpacked data for MembershipCancelled events raised by the Storage contract.
type StorageMembershipCancelledIterator struct {
	Event *StorageMembershipCancelled // Event containing the contract specifics and raw log

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
func (it *StorageMembershipCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMembershipCancelled)
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
		it.Event = new(StorageMembershipCancelled)
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
func (it *StorageMembershipCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMembershipCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMembershipCancelled represents a MembershipCancelled event raised by the Storage contract.
type StorageMembershipCancelled struct {
	Email        string
	MembershipId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMembershipCancelled is a free log retrieval operation binding the contract event 0xc13913f730cfeddf618cd1246dfdb08f7df5c77ee946af5febf0c166310bd2d4.
//
// Solidity: event MembershipCancelled(string email, uint256 membershipId)
func (_Storage *StorageFilterer) FilterMembershipCancelled(opts *bind.FilterOpts) (*StorageMembershipCancelledIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MembershipCancelled")
	if err != nil {
		return nil, err
	}
	return &StorageMembershipCancelledIterator{contract: _Storage.contract, event: "MembershipCancelled", logs: logs, sub: sub}, nil
}

// WatchMembershipCancelled is a free log subscription operation binding the contract event 0xc13913f730cfeddf618cd1246dfdb08f7df5c77ee946af5febf0c166310bd2d4.
//
// Solidity: event MembershipCancelled(string email, uint256 membershipId)
func (_Storage *StorageFilterer) WatchMembershipCancelled(opts *bind.WatchOpts, sink chan<- *StorageMembershipCancelled) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MembershipCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMembershipCancelled)
				if err := _Storage.contract.UnpackLog(event, "MembershipCancelled", log); err != nil {
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

// ParseMembershipCancelled is a log parse operation binding the contract event 0xc13913f730cfeddf618cd1246dfdb08f7df5c77ee946af5febf0c166310bd2d4.
//
// Solidity: event MembershipCancelled(string email, uint256 membershipId)
func (_Storage *StorageFilterer) ParseMembershipCancelled(log types.Log) (*StorageMembershipCancelled, error) {
	event := new(StorageMembershipCancelled)
	if err := _Storage.contract.UnpackLog(event, "MembershipCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMembershipPurchasedIterator is returned from FilterMembershipPurchased and is used to iterate over the raw logs and unpacked data for MembershipPurchased events raised by the Storage contract.
type StorageMembershipPurchasedIterator struct {
	Event *StorageMembershipPurchased // Event containing the contract specifics and raw log

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
func (it *StorageMembershipPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMembershipPurchased)
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
		it.Event = new(StorageMembershipPurchased)
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
func (it *StorageMembershipPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMembershipPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMembershipPurchased represents a MembershipPurchased event raised by the Storage contract.
type StorageMembershipPurchased struct {
	Email          string
	MembershipId   *big.Int
	MembershipType string
	StartDate      *big.Int
	EndDate        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMembershipPurchased is a free log retrieval operation binding the contract event 0x8d8b7df8430bfd03c03180dcea3afe40fc7fd02888f9b2f0920c7e8950928f8b.
//
// Solidity: event MembershipPurchased(string email, uint256 membershipId, string membershipType, uint256 startDate, uint256 endDate)
func (_Storage *StorageFilterer) FilterMembershipPurchased(opts *bind.FilterOpts) (*StorageMembershipPurchasedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MembershipPurchased")
	if err != nil {
		return nil, err
	}
	return &StorageMembershipPurchasedIterator{contract: _Storage.contract, event: "MembershipPurchased", logs: logs, sub: sub}, nil
}

// WatchMembershipPurchased is a free log subscription operation binding the contract event 0x8d8b7df8430bfd03c03180dcea3afe40fc7fd02888f9b2f0920c7e8950928f8b.
//
// Solidity: event MembershipPurchased(string email, uint256 membershipId, string membershipType, uint256 startDate, uint256 endDate)
func (_Storage *StorageFilterer) WatchMembershipPurchased(opts *bind.WatchOpts, sink chan<- *StorageMembershipPurchased) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MembershipPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMembershipPurchased)
				if err := _Storage.contract.UnpackLog(event, "MembershipPurchased", log); err != nil {
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

// ParseMembershipPurchased is a log parse operation binding the contract event 0x8d8b7df8430bfd03c03180dcea3afe40fc7fd02888f9b2f0920c7e8950928f8b.
//
// Solidity: event MembershipPurchased(string email, uint256 membershipId, string membershipType, uint256 startDate, uint256 endDate)
func (_Storage *StorageFilterer) ParseMembershipPurchased(log types.Log) (*StorageMembershipPurchased, error) {
	event := new(StorageMembershipPurchased)
	if err := _Storage.contract.UnpackLog(event, "MembershipPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMembershipRenewedIterator is returned from FilterMembershipRenewed and is used to iterate over the raw logs and unpacked data for MembershipRenewed events raised by the Storage contract.
type StorageMembershipRenewedIterator struct {
	Event *StorageMembershipRenewed // Event containing the contract specifics and raw log

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
func (it *StorageMembershipRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMembershipRenewed)
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
		it.Event = new(StorageMembershipRenewed)
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
func (it *StorageMembershipRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMembershipRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMembershipRenewed represents a MembershipRenewed event raised by the Storage contract.
type StorageMembershipRenewed struct {
	Email        string
	MembershipId *big.Int
	NewEndDate   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMembershipRenewed is a free log retrieval operation binding the contract event 0x22df7f3093777cf15962bceb24f52d3838275deda87e305d1ede5b9a46c15acf.
//
// Solidity: event MembershipRenewed(string email, uint256 membershipId, uint256 newEndDate)
func (_Storage *StorageFilterer) FilterMembershipRenewed(opts *bind.FilterOpts) (*StorageMembershipRenewedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MembershipRenewed")
	if err != nil {
		return nil, err
	}
	return &StorageMembershipRenewedIterator{contract: _Storage.contract, event: "MembershipRenewed", logs: logs, sub: sub}, nil
}

// WatchMembershipRenewed is a free log subscription operation binding the contract event 0x22df7f3093777cf15962bceb24f52d3838275deda87e305d1ede5b9a46c15acf.
//
// Solidity: event MembershipRenewed(string email, uint256 membershipId, uint256 newEndDate)
func (_Storage *StorageFilterer) WatchMembershipRenewed(opts *bind.WatchOpts, sink chan<- *StorageMembershipRenewed) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MembershipRenewed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMembershipRenewed)
				if err := _Storage.contract.UnpackLog(event, "MembershipRenewed", log); err != nil {
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

// ParseMembershipRenewed is a log parse operation binding the contract event 0x22df7f3093777cf15962bceb24f52d3838275deda87e305d1ede5b9a46c15acf.
//
// Solidity: event MembershipRenewed(string email, uint256 membershipId, uint256 newEndDate)
func (_Storage *StorageFilterer) ParseMembershipRenewed(log types.Log) (*StorageMembershipRenewed, error) {
	event := new(StorageMembershipRenewed)
	if err := _Storage.contract.UnpackLog(event, "MembershipRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUserDeletedIterator is returned from FilterUserDeleted and is used to iterate over the raw logs and unpacked data for UserDeleted events raised by the Storage contract.
type StorageUserDeletedIterator struct {
	Event *StorageUserDeleted // Event containing the contract specifics and raw log

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
func (it *StorageUserDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUserDeleted)
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
		it.Event = new(StorageUserDeleted)
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
func (it *StorageUserDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUserDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUserDeleted represents a UserDeleted event raised by the Storage contract.
type StorageUserDeleted struct {
	Email string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserDeleted is a free log retrieval operation binding the contract event 0x28f7ad4961343bc792aec8e7886fc38c6847f9cd253ec14a633ff3bed8370883.
//
// Solidity: event UserDeleted(string email)
func (_Storage *StorageFilterer) FilterUserDeleted(opts *bind.FilterOpts) (*StorageUserDeletedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UserDeleted")
	if err != nil {
		return nil, err
	}
	return &StorageUserDeletedIterator{contract: _Storage.contract, event: "UserDeleted", logs: logs, sub: sub}, nil
}

// WatchUserDeleted is a free log subscription operation binding the contract event 0x28f7ad4961343bc792aec8e7886fc38c6847f9cd253ec14a633ff3bed8370883.
//
// Solidity: event UserDeleted(string email)
func (_Storage *StorageFilterer) WatchUserDeleted(opts *bind.WatchOpts, sink chan<- *StorageUserDeleted) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UserDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUserDeleted)
				if err := _Storage.contract.UnpackLog(event, "UserDeleted", log); err != nil {
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

// ParseUserDeleted is a log parse operation binding the contract event 0x28f7ad4961343bc792aec8e7886fc38c6847f9cd253ec14a633ff3bed8370883.
//
// Solidity: event UserDeleted(string email)
func (_Storage *StorageFilterer) ParseUserDeleted(log types.Log) (*StorageUserDeleted, error) {
	event := new(StorageUserDeleted)
	if err := _Storage.contract.UnpackLog(event, "UserDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the Storage contract.
type StorageUserRegisteredIterator struct {
	Event *StorageUserRegistered // Event containing the contract specifics and raw log

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
func (it *StorageUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUserRegistered)
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
		it.Event = new(StorageUserRegistered)
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
func (it *StorageUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUserRegistered represents a UserRegistered event raised by the Storage contract.
type StorageUserRegistered struct {
	Email   string
	Name    string
	Contact string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0x0a012201ed6ad03a46b4a2be566d40057141543493eb16b35966ad83539abe4f.
//
// Solidity: event UserRegistered(string email, string name, string contact)
func (_Storage *StorageFilterer) FilterUserRegistered(opts *bind.FilterOpts) (*StorageUserRegisteredIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return &StorageUserRegisteredIterator{contract: _Storage.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0x0a012201ed6ad03a46b4a2be566d40057141543493eb16b35966ad83539abe4f.
//
// Solidity: event UserRegistered(string email, string name, string contact)
func (_Storage *StorageFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *StorageUserRegistered) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUserRegistered)
				if err := _Storage.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// ParseUserRegistered is a log parse operation binding the contract event 0x0a012201ed6ad03a46b4a2be566d40057141543493eb16b35966ad83539abe4f.
//
// Solidity: event UserRegistered(string email, string name, string contact)
func (_Storage *StorageFilterer) ParseUserRegistered(log types.Log) (*StorageUserRegistered, error) {
	event := new(StorageUserRegistered)
	if err := _Storage.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUserVerifiedIterator is returned from FilterUserVerified and is used to iterate over the raw logs and unpacked data for UserVerified events raised by the Storage contract.
type StorageUserVerifiedIterator struct {
	Event *StorageUserVerified // Event containing the contract specifics and raw log

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
func (it *StorageUserVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUserVerified)
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
		it.Event = new(StorageUserVerified)
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
func (it *StorageUserVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUserVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUserVerified represents a UserVerified event raised by the Storage contract.
type StorageUserVerified struct {
	Email string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserVerified is a free log retrieval operation binding the contract event 0x820df4d36e99ae1d3d17c6a1bc5028e4e1c9aefe2532ee67f936de82d3aecde9.
//
// Solidity: event UserVerified(string email)
func (_Storage *StorageFilterer) FilterUserVerified(opts *bind.FilterOpts) (*StorageUserVerifiedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UserVerified")
	if err != nil {
		return nil, err
	}
	return &StorageUserVerifiedIterator{contract: _Storage.contract, event: "UserVerified", logs: logs, sub: sub}, nil
}

// WatchUserVerified is a free log subscription operation binding the contract event 0x820df4d36e99ae1d3d17c6a1bc5028e4e1c9aefe2532ee67f936de82d3aecde9.
//
// Solidity: event UserVerified(string email)
func (_Storage *StorageFilterer) WatchUserVerified(opts *bind.WatchOpts, sink chan<- *StorageUserVerified) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UserVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUserVerified)
				if err := _Storage.contract.UnpackLog(event, "UserVerified", log); err != nil {
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

// ParseUserVerified is a log parse operation binding the contract event 0x820df4d36e99ae1d3d17c6a1bc5028e4e1c9aefe2532ee67f936de82d3aecde9.
//
// Solidity: event UserVerified(string email)
func (_Storage *StorageFilterer) ParseUserVerified(log types.Log) (*StorageUserVerified, error) {
	event := new(StorageUserVerified)
	if err := _Storage.contract.UnpackLog(event, "UserVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
