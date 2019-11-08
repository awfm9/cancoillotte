// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ERC165QueryABI is the input ABI used to generate the binding from.
const ERC165QueryABI = "[]"

// ERC165QueryBin is the compiled bytecode used for deploying new contracts.
var ERC165QueryBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a72315820e48805ed8c84eee93bb680523de0b1f84055fbbda775bd96ce3fc33f4162f45564736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployERC165Query deploys a new Ethereum contract, binding an instance of ERC165Query to it.
func DeployERC165Query(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165Query, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165QueryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165QueryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165Query{ERC165QueryCaller: ERC165QueryCaller{contract: contract}, ERC165QueryTransactor: ERC165QueryTransactor{contract: contract}, ERC165QueryFilterer: ERC165QueryFilterer{contract: contract}}, nil
}

// ERC165Query is an auto generated Go binding around an Ethereum contract.
type ERC165Query struct {
	ERC165QueryCaller     // Read-only binding to the contract
	ERC165QueryTransactor // Write-only binding to the contract
	ERC165QueryFilterer   // Log filterer for contract events
}

// ERC165QueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165QueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165QueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165QueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165QueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165QueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165QuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165QuerySession struct {
	Contract     *ERC165Query      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165QueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165QueryCallerSession struct {
	Contract *ERC165QueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC165QueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165QueryTransactorSession struct {
	Contract     *ERC165QueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC165QueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165QueryRaw struct {
	Contract *ERC165Query // Generic contract binding to access the raw methods on
}

// ERC165QueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165QueryCallerRaw struct {
	Contract *ERC165QueryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165QueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165QueryTransactorRaw struct {
	Contract *ERC165QueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Query creates a new instance of ERC165Query, bound to a specific deployed contract.
func NewERC165Query(address common.Address, backend bind.ContractBackend) (*ERC165Query, error) {
	contract, err := bindERC165Query(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Query{ERC165QueryCaller: ERC165QueryCaller{contract: contract}, ERC165QueryTransactor: ERC165QueryTransactor{contract: contract}, ERC165QueryFilterer: ERC165QueryFilterer{contract: contract}}, nil
}

// NewERC165QueryCaller creates a new read-only instance of ERC165Query, bound to a specific deployed contract.
func NewERC165QueryCaller(address common.Address, caller bind.ContractCaller) (*ERC165QueryCaller, error) {
	contract, err := bindERC165Query(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165QueryCaller{contract: contract}, nil
}

// NewERC165QueryTransactor creates a new write-only instance of ERC165Query, bound to a specific deployed contract.
func NewERC165QueryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165QueryTransactor, error) {
	contract, err := bindERC165Query(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165QueryTransactor{contract: contract}, nil
}

// NewERC165QueryFilterer creates a new log filterer instance of ERC165Query, bound to a specific deployed contract.
func NewERC165QueryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165QueryFilterer, error) {
	contract, err := bindERC165Query(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165QueryFilterer{contract: contract}, nil
}

// bindERC165Query binds a generic wrapper to an already deployed contract.
func bindERC165Query(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165QueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Query *ERC165QueryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Query.Contract.ERC165QueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Query *ERC165QueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Query.Contract.ERC165QueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Query *ERC165QueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Query.Contract.ERC165QueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Query *ERC165QueryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165Query.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Query *ERC165QueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Query.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Query *ERC165QueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Query.Contract.contract.Transact(opts, method, params...)
}
