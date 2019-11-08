// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SigToolsABI is the input ABI used to generate the binding from.
const SigToolsABI = "[]"

// SigToolsBin is the compiled bytecode used for deploying new contracts.
var SigToolsBin = "0x607c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582017c9abedd8f8e8f2221684718f623ebb4cf45040831769c7bbf891937fa3d7e464736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeploySigTools deploys a new Ethereum contract, binding an instance of SigTools to it.
func DeploySigTools(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigTools, error) {
	parsed, err := abi.JSON(strings.NewReader(SigToolsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigToolsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigTools{SigToolsCaller: SigToolsCaller{contract: contract}, SigToolsTransactor: SigToolsTransactor{contract: contract}, SigToolsFilterer: SigToolsFilterer{contract: contract}}, nil
}

// SigTools is an auto generated Go binding around an Ethereum contract.
type SigTools struct {
	SigToolsCaller     // Read-only binding to the contract
	SigToolsTransactor // Write-only binding to the contract
	SigToolsFilterer   // Log filterer for contract events
}

// SigToolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigToolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigToolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigToolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigToolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigToolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigToolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigToolsSession struct {
	Contract     *SigTools         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigToolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigToolsCallerSession struct {
	Contract *SigToolsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigToolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigToolsTransactorSession struct {
	Contract     *SigToolsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigToolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigToolsRaw struct {
	Contract *SigTools // Generic contract binding to access the raw methods on
}

// SigToolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigToolsCallerRaw struct {
	Contract *SigToolsCaller // Generic read-only contract binding to access the raw methods on
}

// SigToolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigToolsTransactorRaw struct {
	Contract *SigToolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigTools creates a new instance of SigTools, bound to a specific deployed contract.
func NewSigTools(address common.Address, backend bind.ContractBackend) (*SigTools, error) {
	contract, err := bindSigTools(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigTools{SigToolsCaller: SigToolsCaller{contract: contract}, SigToolsTransactor: SigToolsTransactor{contract: contract}, SigToolsFilterer: SigToolsFilterer{contract: contract}}, nil
}

// NewSigToolsCaller creates a new read-only instance of SigTools, bound to a specific deployed contract.
func NewSigToolsCaller(address common.Address, caller bind.ContractCaller) (*SigToolsCaller, error) {
	contract, err := bindSigTools(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigToolsCaller{contract: contract}, nil
}

// NewSigToolsTransactor creates a new write-only instance of SigTools, bound to a specific deployed contract.
func NewSigToolsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigToolsTransactor, error) {
	contract, err := bindSigTools(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigToolsTransactor{contract: contract}, nil
}

// NewSigToolsFilterer creates a new log filterer instance of SigTools, bound to a specific deployed contract.
func NewSigToolsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigToolsFilterer, error) {
	contract, err := bindSigTools(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigToolsFilterer{contract: contract}, nil
}

// bindSigTools binds a generic wrapper to an already deployed contract.
func bindSigTools(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigToolsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigTools *SigToolsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigTools.Contract.SigToolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigTools *SigToolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigTools.Contract.SigToolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigTools *SigToolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigTools.Contract.SigToolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigTools *SigToolsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigTools.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigTools *SigToolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigTools.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigTools *SigToolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigTools.Contract.contract.Transact(opts, method, params...)
}
