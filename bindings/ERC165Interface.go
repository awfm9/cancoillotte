// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC165InterfaceABI is the input ABI used to generate the binding from.
const ERC165InterfaceABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165InterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ERC165InterfaceFuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165Interface is an auto generated Go binding around an Ethereum contract.
type ERC165Interface struct {
	ERC165InterfaceCaller     // Read-only binding to the contract
	ERC165InterfaceTransactor // Write-only binding to the contract
	ERC165InterfaceFilterer   // Log filterer for contract events
}

// ERC165InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165InterfaceSession struct {
	Contract     *ERC165Interface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165InterfaceCallerSession struct {
	Contract *ERC165InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC165InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165InterfaceTransactorSession struct {
	Contract     *ERC165InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC165InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165InterfaceRaw struct {
	Contract *ERC165Interface // Generic contract binding to access the raw methods on
}

// ERC165InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165InterfaceCallerRaw struct {
	Contract *ERC165InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165InterfaceTransactorRaw struct {
	Contract *ERC165InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Interface creates a new instance of ERC165Interface, bound to a specific deployed contract.
func NewERC165Interface(address common.Address, backend bind.ContractBackend) (*ERC165Interface, error) {
	contract, err := bindERC165Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Interface{ERC165InterfaceCaller: ERC165InterfaceCaller{contract: contract}, ERC165InterfaceTransactor: ERC165InterfaceTransactor{contract: contract}, ERC165InterfaceFilterer: ERC165InterfaceFilterer{contract: contract}}, nil
}

// NewERC165InterfaceCaller creates a new read-only instance of ERC165Interface, bound to a specific deployed contract.
func NewERC165InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC165InterfaceCaller, error) {
	contract, err := bindERC165Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165InterfaceCaller{contract: contract}, nil
}

// NewERC165InterfaceTransactor creates a new write-only instance of ERC165Interface, bound to a specific deployed contract.
func NewERC165InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165InterfaceTransactor, error) {
	contract, err := bindERC165Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165InterfaceTransactor{contract: contract}, nil
}

// NewERC165InterfaceFilterer creates a new log filterer instance of ERC165Interface, bound to a specific deployed contract.
func NewERC165InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165InterfaceFilterer, error) {
	contract, err := bindERC165Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165InterfaceFilterer{contract: contract}, nil
}

// bindERC165Interface binds a generic wrapper to an already deployed contract.
func bindERC165Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Interface *ERC165InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Interface.Contract.ERC165InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Interface *ERC165InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Interface.Contract.ERC165InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Interface *ERC165InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Interface.Contract.ERC165InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Interface *ERC165InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Interface *ERC165InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Interface *ERC165InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Interface.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165Interface *ERC165InterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC165Interface.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165Interface *ERC165InterfaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Interface.Contract.SupportsInterface(&_ERC165Interface.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165Interface *ERC165InterfaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Interface.Contract.SupportsInterface(&_ERC165Interface.CallOpts, interfaceId)
}
