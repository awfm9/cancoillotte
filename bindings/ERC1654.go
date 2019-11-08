// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC1654ABI is the input ABI used to generate the binding from.
const ERC1654ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ERC1654_VALIDSIGNATURE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC1654FuncSigs maps the 4-byte function signature to its string representation.
var ERC1654FuncSigs = map[string]string{
	"50abdde2": "ERC1654_VALIDSIGNATURE()",
	"1626ba7e": "isValidSignature(bytes32,bytes)",
}

// ERC1654 is an auto generated Go binding around an Ethereum contract.
type ERC1654 struct {
	ERC1654Caller     // Read-only binding to the contract
	ERC1654Transactor // Write-only binding to the contract
	ERC1654Filterer   // Log filterer for contract events
}

// ERC1654Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1654Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1654Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1654Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1654Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1654Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1654Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1654Session struct {
	Contract     *ERC1654          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1654CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1654CallerSession struct {
	Contract *ERC1654Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1654TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1654TransactorSession struct {
	Contract     *ERC1654Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1654Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1654Raw struct {
	Contract *ERC1654 // Generic contract binding to access the raw methods on
}

// ERC1654CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1654CallerRaw struct {
	Contract *ERC1654Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1654TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1654TransactorRaw struct {
	Contract *ERC1654Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1654 creates a new instance of ERC1654, bound to a specific deployed contract.
func NewERC1654(address common.Address, backend bind.ContractBackend) (*ERC1654, error) {
	contract, err := bindERC1654(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1654{ERC1654Caller: ERC1654Caller{contract: contract}, ERC1654Transactor: ERC1654Transactor{contract: contract}, ERC1654Filterer: ERC1654Filterer{contract: contract}}, nil
}

// NewERC1654Caller creates a new read-only instance of ERC1654, bound to a specific deployed contract.
func NewERC1654Caller(address common.Address, caller bind.ContractCaller) (*ERC1654Caller, error) {
	contract, err := bindERC1654(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1654Caller{contract: contract}, nil
}

// NewERC1654Transactor creates a new write-only instance of ERC1654, bound to a specific deployed contract.
func NewERC1654Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1654Transactor, error) {
	contract, err := bindERC1654(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1654Transactor{contract: contract}, nil
}

// NewERC1654Filterer creates a new log filterer instance of ERC1654, bound to a specific deployed contract.
func NewERC1654Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1654Filterer, error) {
	contract, err := bindERC1654(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1654Filterer{contract: contract}, nil
}

// bindERC1654 binds a generic wrapper to an already deployed contract.
func bindERC1654(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1654ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1654 *ERC1654Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC1654.Contract.ERC1654Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1654 *ERC1654Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1654.Contract.ERC1654Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1654 *ERC1654Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1654.Contract.ERC1654Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1654 *ERC1654CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC1654.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1654 *ERC1654TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1654.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1654 *ERC1654TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1654.Contract.contract.Transact(opts, method, params...)
}

// ERC1654VALIDSIGNATURE is a free data retrieval call binding the contract method 0x50abdde2.
//
// Solidity: function ERC1654_VALIDSIGNATURE() constant returns(bytes4)
func (_ERC1654 *ERC1654Caller) ERC1654VALIDSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC1654.contract.Call(opts, out, "ERC1654_VALIDSIGNATURE")
	return *ret0, err
}

// ERC1654VALIDSIGNATURE is a free data retrieval call binding the contract method 0x50abdde2.
//
// Solidity: function ERC1654_VALIDSIGNATURE() constant returns(bytes4)
func (_ERC1654 *ERC1654Session) ERC1654VALIDSIGNATURE() ([4]byte, error) {
	return _ERC1654.Contract.ERC1654VALIDSIGNATURE(&_ERC1654.CallOpts)
}

// ERC1654VALIDSIGNATURE is a free data retrieval call binding the contract method 0x50abdde2.
//
// Solidity: function ERC1654_VALIDSIGNATURE() constant returns(bytes4)
func (_ERC1654 *ERC1654CallerSession) ERC1654VALIDSIGNATURE() ([4]byte, error) {
	return _ERC1654.Contract.ERC1654VALIDSIGNATURE(&_ERC1654.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) constant returns(bytes4)
func (_ERC1654 *ERC1654Caller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC1654.contract.Call(opts, out, "isValidSignature", hash, _signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) constant returns(bytes4)
func (_ERC1654 *ERC1654Session) IsValidSignature(hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ERC1654.Contract.IsValidSignature(&_ERC1654.CallOpts, hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) constant returns(bytes4)
func (_ERC1654 *ERC1654CallerSession) IsValidSignature(hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ERC1654.Contract.IsValidSignature(&_ERC1654.CallOpts, hash, _signature)
}
