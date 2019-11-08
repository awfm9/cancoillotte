// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WizardPresaleInterfaceABI is the input ABI used to generate the binding from.
const WizardPresaleInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_INTERFACE_ID_WIZARDPRESALE\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"absorbWizard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"absorbWizardMulti\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"affinities\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"costToPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"powerToCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// WizardPresaleInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var WizardPresaleInterfaceFuncSigs = map[string]string{
	"8ef0da6f": "_INTERFACE_ID_WIZARDPRESALE()",
	"a7847c3a": "absorbWizard(uint256)",
	"476c4a5e": "absorbWizardMulti(uint256[])",
	"e5a604bf": "costToPower(uint256)",
	"48b92c20": "powerToCost(uint256)",
}

// WizardPresaleInterface is an auto generated Go binding around an Ethereum contract.
type WizardPresaleInterface struct {
	WizardPresaleInterfaceCaller     // Read-only binding to the contract
	WizardPresaleInterfaceTransactor // Write-only binding to the contract
	WizardPresaleInterfaceFilterer   // Log filterer for contract events
}

// WizardPresaleInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardPresaleInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardPresaleInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardPresaleInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardPresaleInterfaceSession struct {
	Contract     *WizardPresaleInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WizardPresaleInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardPresaleInterfaceCallerSession struct {
	Contract *WizardPresaleInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// WizardPresaleInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardPresaleInterfaceTransactorSession struct {
	Contract     *WizardPresaleInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// WizardPresaleInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardPresaleInterfaceRaw struct {
	Contract *WizardPresaleInterface // Generic contract binding to access the raw methods on
}

// WizardPresaleInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardPresaleInterfaceCallerRaw struct {
	Contract *WizardPresaleInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// WizardPresaleInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardPresaleInterfaceTransactorRaw struct {
	Contract *WizardPresaleInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardPresaleInterface creates a new instance of WizardPresaleInterface, bound to a specific deployed contract.
func NewWizardPresaleInterface(address common.Address, backend bind.ContractBackend) (*WizardPresaleInterface, error) {
	contract, err := bindWizardPresaleInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleInterface{WizardPresaleInterfaceCaller: WizardPresaleInterfaceCaller{contract: contract}, WizardPresaleInterfaceTransactor: WizardPresaleInterfaceTransactor{contract: contract}, WizardPresaleInterfaceFilterer: WizardPresaleInterfaceFilterer{contract: contract}}, nil
}

// NewWizardPresaleInterfaceCaller creates a new read-only instance of WizardPresaleInterface, bound to a specific deployed contract.
func NewWizardPresaleInterfaceCaller(address common.Address, caller bind.ContractCaller) (*WizardPresaleInterfaceCaller, error) {
	contract, err := bindWizardPresaleInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleInterfaceCaller{contract: contract}, nil
}

// NewWizardPresaleInterfaceTransactor creates a new write-only instance of WizardPresaleInterface, bound to a specific deployed contract.
func NewWizardPresaleInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardPresaleInterfaceTransactor, error) {
	contract, err := bindWizardPresaleInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleInterfaceTransactor{contract: contract}, nil
}

// NewWizardPresaleInterfaceFilterer creates a new log filterer instance of WizardPresaleInterface, bound to a specific deployed contract.
func NewWizardPresaleInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardPresaleInterfaceFilterer, error) {
	contract, err := bindWizardPresaleInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleInterfaceFilterer{contract: contract}, nil
}

// bindWizardPresaleInterface binds a generic wrapper to an already deployed contract.
func bindWizardPresaleInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardPresaleInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardPresaleInterface *WizardPresaleInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardPresaleInterface.Contract.WizardPresaleInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardPresaleInterface *WizardPresaleInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.WizardPresaleInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardPresaleInterface *WizardPresaleInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.WizardPresaleInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardPresaleInterface *WizardPresaleInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardPresaleInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.contract.Transact(opts, method, params...)
}

// INTERFACEIDWIZARDPRESALE is a free data retrieval call binding the contract method 0x8ef0da6f.
//
// Solidity: function _INTERFACE_ID_WIZARDPRESALE() constant returns(bytes4)
func (_WizardPresaleInterface *WizardPresaleInterfaceCaller) INTERFACEIDWIZARDPRESALE(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _WizardPresaleInterface.contract.Call(opts, out, "_INTERFACE_ID_WIZARDPRESALE")
	return *ret0, err
}

// INTERFACEIDWIZARDPRESALE is a free data retrieval call binding the contract method 0x8ef0da6f.
//
// Solidity: function _INTERFACE_ID_WIZARDPRESALE() constant returns(bytes4)
func (_WizardPresaleInterface *WizardPresaleInterfaceSession) INTERFACEIDWIZARDPRESALE() ([4]byte, error) {
	return _WizardPresaleInterface.Contract.INTERFACEIDWIZARDPRESALE(&_WizardPresaleInterface.CallOpts)
}

// INTERFACEIDWIZARDPRESALE is a free data retrieval call binding the contract method 0x8ef0da6f.
//
// Solidity: function _INTERFACE_ID_WIZARDPRESALE() constant returns(bytes4)
func (_WizardPresaleInterface *WizardPresaleInterfaceCallerSession) INTERFACEIDWIZARDPRESALE() ([4]byte, error) {
	return _WizardPresaleInterface.Contract.INTERFACEIDWIZARDPRESALE(&_WizardPresaleInterface.CallOpts)
}

// CostToPower is a free data retrieval call binding the contract method 0xe5a604bf.
//
// Solidity: function costToPower(uint256 cost) constant returns(uint256 power)
func (_WizardPresaleInterface *WizardPresaleInterfaceCaller) CostToPower(opts *bind.CallOpts, cost *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WizardPresaleInterface.contract.Call(opts, out, "costToPower", cost)
	return *ret0, err
}

// CostToPower is a free data retrieval call binding the contract method 0xe5a604bf.
//
// Solidity: function costToPower(uint256 cost) constant returns(uint256 power)
func (_WizardPresaleInterface *WizardPresaleInterfaceSession) CostToPower(cost *big.Int) (*big.Int, error) {
	return _WizardPresaleInterface.Contract.CostToPower(&_WizardPresaleInterface.CallOpts, cost)
}

// CostToPower is a free data retrieval call binding the contract method 0xe5a604bf.
//
// Solidity: function costToPower(uint256 cost) constant returns(uint256 power)
func (_WizardPresaleInterface *WizardPresaleInterfaceCallerSession) CostToPower(cost *big.Int) (*big.Int, error) {
	return _WizardPresaleInterface.Contract.CostToPower(&_WizardPresaleInterface.CallOpts, cost)
}

// PowerToCost is a free data retrieval call binding the contract method 0x48b92c20.
//
// Solidity: function powerToCost(uint256 power) constant returns(uint256 cost)
func (_WizardPresaleInterface *WizardPresaleInterfaceCaller) PowerToCost(opts *bind.CallOpts, power *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WizardPresaleInterface.contract.Call(opts, out, "powerToCost", power)
	return *ret0, err
}

// PowerToCost is a free data retrieval call binding the contract method 0x48b92c20.
//
// Solidity: function powerToCost(uint256 power) constant returns(uint256 cost)
func (_WizardPresaleInterface *WizardPresaleInterfaceSession) PowerToCost(power *big.Int) (*big.Int, error) {
	return _WizardPresaleInterface.Contract.PowerToCost(&_WizardPresaleInterface.CallOpts, power)
}

// PowerToCost is a free data retrieval call binding the contract method 0x48b92c20.
//
// Solidity: function powerToCost(uint256 power) constant returns(uint256 cost)
func (_WizardPresaleInterface *WizardPresaleInterfaceCallerSession) PowerToCost(power *big.Int) (*big.Int, error) {
	return _WizardPresaleInterface.Contract.PowerToCost(&_WizardPresaleInterface.CallOpts, power)
}

// AbsorbWizard is a paid mutator transaction binding the contract method 0xa7847c3a.
//
// Solidity: function absorbWizard(uint256 id) returns(address owner, uint256 power, uint8 affinity)
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactor) AbsorbWizard(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.contract.Transact(opts, "absorbWizard", id)
}

// AbsorbWizard is a paid mutator transaction binding the contract method 0xa7847c3a.
//
// Solidity: function absorbWizard(uint256 id) returns(address owner, uint256 power, uint8 affinity)
func (_WizardPresaleInterface *WizardPresaleInterfaceSession) AbsorbWizard(id *big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.AbsorbWizard(&_WizardPresaleInterface.TransactOpts, id)
}

// AbsorbWizard is a paid mutator transaction binding the contract method 0xa7847c3a.
//
// Solidity: function absorbWizard(uint256 id) returns(address owner, uint256 power, uint8 affinity)
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactorSession) AbsorbWizard(id *big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.AbsorbWizard(&_WizardPresaleInterface.TransactOpts, id)
}

// AbsorbWizardMulti is a paid mutator transaction binding the contract method 0x476c4a5e.
//
// Solidity: function absorbWizardMulti(uint256[] ids) returns(address[] owners, uint256[] powers, uint8[] affinities)
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactor) AbsorbWizardMulti(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.contract.Transact(opts, "absorbWizardMulti", ids)
}

// AbsorbWizardMulti is a paid mutator transaction binding the contract method 0x476c4a5e.
//
// Solidity: function absorbWizardMulti(uint256[] ids) returns(address[] owners, uint256[] powers, uint8[] affinities)
func (_WizardPresaleInterface *WizardPresaleInterfaceSession) AbsorbWizardMulti(ids []*big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.AbsorbWizardMulti(&_WizardPresaleInterface.TransactOpts, ids)
}

// AbsorbWizardMulti is a paid mutator transaction binding the contract method 0x476c4a5e.
//
// Solidity: function absorbWizardMulti(uint256[] ids) returns(address[] owners, uint256[] powers, uint8[] affinities)
func (_WizardPresaleInterface *WizardPresaleInterfaceTransactorSession) AbsorbWizardMulti(ids []*big.Int) (*types.Transaction, error) {
	return _WizardPresaleInterface.Contract.AbsorbWizardMulti(&_WizardPresaleInterface.TransactOpts, ids)
}
