// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WizardGuildInterfaceIdABI is the input ABI used to generate the binding from.
const WizardGuildInterfaceIdABI = "[]"

// WizardGuildInterfaceIdBin is the compiled bytecode used for deploying new contracts.
var WizardGuildInterfaceIdBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a723158206dff98d2e0a6da7153525223a8aec830647d35fff2d55fd971eadebddd640cbf64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployWizardGuildInterfaceId deploys a new Ethereum contract, binding an instance of WizardGuildInterfaceId to it.
func DeployWizardGuildInterfaceId(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WizardGuildInterfaceId, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardGuildInterfaceIdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WizardGuildInterfaceIdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WizardGuildInterfaceId{WizardGuildInterfaceIdCaller: WizardGuildInterfaceIdCaller{contract: contract}, WizardGuildInterfaceIdTransactor: WizardGuildInterfaceIdTransactor{contract: contract}, WizardGuildInterfaceIdFilterer: WizardGuildInterfaceIdFilterer{contract: contract}}, nil
}

// WizardGuildInterfaceId is an auto generated Go binding around an Ethereum contract.
type WizardGuildInterfaceId struct {
	WizardGuildInterfaceIdCaller     // Read-only binding to the contract
	WizardGuildInterfaceIdTransactor // Write-only binding to the contract
	WizardGuildInterfaceIdFilterer   // Log filterer for contract events
}

// WizardGuildInterfaceIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardGuildInterfaceIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardGuildInterfaceIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardGuildInterfaceIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardGuildInterfaceIdSession struct {
	Contract     *WizardGuildInterfaceId // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WizardGuildInterfaceIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardGuildInterfaceIdCallerSession struct {
	Contract *WizardGuildInterfaceIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// WizardGuildInterfaceIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardGuildInterfaceIdTransactorSession struct {
	Contract     *WizardGuildInterfaceIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// WizardGuildInterfaceIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardGuildInterfaceIdRaw struct {
	Contract *WizardGuildInterfaceId // Generic contract binding to access the raw methods on
}

// WizardGuildInterfaceIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardGuildInterfaceIdCallerRaw struct {
	Contract *WizardGuildInterfaceIdCaller // Generic read-only contract binding to access the raw methods on
}

// WizardGuildInterfaceIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardGuildInterfaceIdTransactorRaw struct {
	Contract *WizardGuildInterfaceIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardGuildInterfaceId creates a new instance of WizardGuildInterfaceId, bound to a specific deployed contract.
func NewWizardGuildInterfaceId(address common.Address, backend bind.ContractBackend) (*WizardGuildInterfaceId, error) {
	contract, err := bindWizardGuildInterfaceId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceId{WizardGuildInterfaceIdCaller: WizardGuildInterfaceIdCaller{contract: contract}, WizardGuildInterfaceIdTransactor: WizardGuildInterfaceIdTransactor{contract: contract}, WizardGuildInterfaceIdFilterer: WizardGuildInterfaceIdFilterer{contract: contract}}, nil
}

// NewWizardGuildInterfaceIdCaller creates a new read-only instance of WizardGuildInterfaceId, bound to a specific deployed contract.
func NewWizardGuildInterfaceIdCaller(address common.Address, caller bind.ContractCaller) (*WizardGuildInterfaceIdCaller, error) {
	contract, err := bindWizardGuildInterfaceId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceIdCaller{contract: contract}, nil
}

// NewWizardGuildInterfaceIdTransactor creates a new write-only instance of WizardGuildInterfaceId, bound to a specific deployed contract.
func NewWizardGuildInterfaceIdTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardGuildInterfaceIdTransactor, error) {
	contract, err := bindWizardGuildInterfaceId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceIdTransactor{contract: contract}, nil
}

// NewWizardGuildInterfaceIdFilterer creates a new log filterer instance of WizardGuildInterfaceId, bound to a specific deployed contract.
func NewWizardGuildInterfaceIdFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardGuildInterfaceIdFilterer, error) {
	contract, err := bindWizardGuildInterfaceId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceIdFilterer{contract: contract}, nil
}

// bindWizardGuildInterfaceId binds a generic wrapper to an already deployed contract.
func bindWizardGuildInterfaceId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardGuildInterfaceIdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardGuildInterfaceId.Contract.WizardGuildInterfaceIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardGuildInterfaceId.Contract.WizardGuildInterfaceIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardGuildInterfaceId.Contract.WizardGuildInterfaceIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardGuildInterfaceId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardGuildInterfaceId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardGuildInterfaceId *WizardGuildInterfaceIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardGuildInterfaceId.Contract.contract.Transact(opts, method, params...)
}
