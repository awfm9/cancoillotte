// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WizardConstantsABI is the input ABI used to generate the binding from.
const WizardConstantsABI = "[]"

// WizardConstantsBin is the compiled bytecode used for deploying new contracts.
var WizardConstantsBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a723158205f76db54f5d3285209ac191ebb1cd5f3cf78d73e80d2155f566b2f221ba735b364736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployWizardConstants deploys a new Ethereum contract, binding an instance of WizardConstants to it.
func DeployWizardConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WizardConstants, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardConstantsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WizardConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WizardConstants{WizardConstantsCaller: WizardConstantsCaller{contract: contract}, WizardConstantsTransactor: WizardConstantsTransactor{contract: contract}, WizardConstantsFilterer: WizardConstantsFilterer{contract: contract}}, nil
}

// WizardConstants is an auto generated Go binding around an Ethereum contract.
type WizardConstants struct {
	WizardConstantsCaller     // Read-only binding to the contract
	WizardConstantsTransactor // Write-only binding to the contract
	WizardConstantsFilterer   // Log filterer for contract events
}

// WizardConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardConstantsSession struct {
	Contract     *WizardConstants  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WizardConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardConstantsCallerSession struct {
	Contract *WizardConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WizardConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardConstantsTransactorSession struct {
	Contract     *WizardConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WizardConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardConstantsRaw struct {
	Contract *WizardConstants // Generic contract binding to access the raw methods on
}

// WizardConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardConstantsCallerRaw struct {
	Contract *WizardConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// WizardConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardConstantsTransactorRaw struct {
	Contract *WizardConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardConstants creates a new instance of WizardConstants, bound to a specific deployed contract.
func NewWizardConstants(address common.Address, backend bind.ContractBackend) (*WizardConstants, error) {
	contract, err := bindWizardConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardConstants{WizardConstantsCaller: WizardConstantsCaller{contract: contract}, WizardConstantsTransactor: WizardConstantsTransactor{contract: contract}, WizardConstantsFilterer: WizardConstantsFilterer{contract: contract}}, nil
}

// NewWizardConstantsCaller creates a new read-only instance of WizardConstants, bound to a specific deployed contract.
func NewWizardConstantsCaller(address common.Address, caller bind.ContractCaller) (*WizardConstantsCaller, error) {
	contract, err := bindWizardConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardConstantsCaller{contract: contract}, nil
}

// NewWizardConstantsTransactor creates a new write-only instance of WizardConstants, bound to a specific deployed contract.
func NewWizardConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardConstantsTransactor, error) {
	contract, err := bindWizardConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardConstantsTransactor{contract: contract}, nil
}

// NewWizardConstantsFilterer creates a new log filterer instance of WizardConstants, bound to a specific deployed contract.
func NewWizardConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardConstantsFilterer, error) {
	contract, err := bindWizardConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardConstantsFilterer{contract: contract}, nil
}

// bindWizardConstants binds a generic wrapper to an already deployed contract.
func bindWizardConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardConstants *WizardConstantsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardConstants.Contract.WizardConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardConstants *WizardConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardConstants.Contract.WizardConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardConstants *WizardConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardConstants.Contract.WizardConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardConstants *WizardConstantsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardConstants *WizardConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardConstants *WizardConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardConstants.Contract.contract.Transact(opts, method, params...)
}
