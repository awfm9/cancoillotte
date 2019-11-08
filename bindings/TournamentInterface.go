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

// TournamentInterfaceABI is the input ABI used to generate the binding from.
const TournamentInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint88[]\",\"name\":\"powers\",\"type\":\"uint88[]\"}],\"name\":\"enterWizards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"powerScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"revive\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TournamentInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var TournamentInterfaceFuncSigs = map[string]string{
	"83197ef0": "destroy()",
	"b9d95abb": "enterWizards(uint256[],uint88[])",
	"22f3e2d4": "isActive()",
	"ad81e4d6": "powerScale()",
	"8baecc21": "revive(uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// TournamentInterface is an auto generated Go binding around an Ethereum contract.
type TournamentInterface struct {
	TournamentInterfaceCaller     // Read-only binding to the contract
	TournamentInterfaceTransactor // Write-only binding to the contract
	TournamentInterfaceFilterer   // Log filterer for contract events
}

// TournamentInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TournamentInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TournamentInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TournamentInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TournamentInterfaceSession struct {
	Contract     *TournamentInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TournamentInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TournamentInterfaceCallerSession struct {
	Contract *TournamentInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TournamentInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TournamentInterfaceTransactorSession struct {
	Contract     *TournamentInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TournamentInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TournamentInterfaceRaw struct {
	Contract *TournamentInterface // Generic contract binding to access the raw methods on
}

// TournamentInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TournamentInterfaceCallerRaw struct {
	Contract *TournamentInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TournamentInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TournamentInterfaceTransactorRaw struct {
	Contract *TournamentInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTournamentInterface creates a new instance of TournamentInterface, bound to a specific deployed contract.
func NewTournamentInterface(address common.Address, backend bind.ContractBackend) (*TournamentInterface, error) {
	contract, err := bindTournamentInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TournamentInterface{TournamentInterfaceCaller: TournamentInterfaceCaller{contract: contract}, TournamentInterfaceTransactor: TournamentInterfaceTransactor{contract: contract}, TournamentInterfaceFilterer: TournamentInterfaceFilterer{contract: contract}}, nil
}

// NewTournamentInterfaceCaller creates a new read-only instance of TournamentInterface, bound to a specific deployed contract.
func NewTournamentInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TournamentInterfaceCaller, error) {
	contract, err := bindTournamentInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceCaller{contract: contract}, nil
}

// NewTournamentInterfaceTransactor creates a new write-only instance of TournamentInterface, bound to a specific deployed contract.
func NewTournamentInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TournamentInterfaceTransactor, error) {
	contract, err := bindTournamentInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceTransactor{contract: contract}, nil
}

// NewTournamentInterfaceFilterer creates a new log filterer instance of TournamentInterface, bound to a specific deployed contract.
func NewTournamentInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TournamentInterfaceFilterer, error) {
	contract, err := bindTournamentInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceFilterer{contract: contract}, nil
}

// bindTournamentInterface binds a generic wrapper to an already deployed contract.
func bindTournamentInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentInterface *TournamentInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentInterface.Contract.TournamentInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentInterface *TournamentInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentInterface.Contract.TournamentInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentInterface *TournamentInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentInterface.Contract.TournamentInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentInterface *TournamentInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentInterface *TournamentInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentInterface *TournamentInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentInterface.Contract.contract.Transact(opts, method, params...)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_TournamentInterface *TournamentInterfaceCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentInterface.contract.Call(opts, out, "isActive")
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_TournamentInterface *TournamentInterfaceSession) IsActive() (bool, error) {
	return _TournamentInterface.Contract.IsActive(&_TournamentInterface.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_TournamentInterface *TournamentInterfaceCallerSession) IsActive() (bool, error) {
	return _TournamentInterface.Contract.IsActive(&_TournamentInterface.CallOpts)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_TournamentInterface *TournamentInterfaceCaller) PowerScale(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentInterface.contract.Call(opts, out, "powerScale")
	return *ret0, err
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_TournamentInterface *TournamentInterfaceSession) PowerScale() (*big.Int, error) {
	return _TournamentInterface.Contract.PowerScale(&_TournamentInterface.CallOpts)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() constant returns(uint256)
func (_TournamentInterface *TournamentInterfaceCallerSession) PowerScale() (*big.Int, error) {
	return _TournamentInterface.Contract.PowerScale(&_TournamentInterface.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TournamentInterface *TournamentInterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentInterface.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TournamentInterface *TournamentInterfaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TournamentInterface.Contract.SupportsInterface(&_TournamentInterface.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TournamentInterface *TournamentInterfaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TournamentInterface.Contract.SupportsInterface(&_TournamentInterface.CallOpts, interfaceId)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_TournamentInterface *TournamentInterfaceTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentInterface.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_TournamentInterface *TournamentInterfaceSession) Destroy() (*types.Transaction, error) {
	return _TournamentInterface.Contract.Destroy(&_TournamentInterface.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_TournamentInterface *TournamentInterfaceTransactorSession) Destroy() (*types.Transaction, error) {
	return _TournamentInterface.Contract.Destroy(&_TournamentInterface.TransactOpts)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_TournamentInterface *TournamentInterfaceTransactor) EnterWizards(opts *bind.TransactOpts, wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _TournamentInterface.contract.Transact(opts, "enterWizards", wizardIds, powers)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_TournamentInterface *TournamentInterfaceSession) EnterWizards(wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _TournamentInterface.Contract.EnterWizards(&_TournamentInterface.TransactOpts, wizardIds, powers)
}

// EnterWizards is a paid mutator transaction binding the contract method 0xb9d95abb.
//
// Solidity: function enterWizards(uint256[] wizardIds, uint88[] powers) returns()
func (_TournamentInterface *TournamentInterfaceTransactorSession) EnterWizards(wizardIds []*big.Int, powers []*big.Int) (*types.Transaction, error) {
	return _TournamentInterface.Contract.EnterWizards(&_TournamentInterface.TransactOpts, wizardIds, powers)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_TournamentInterface *TournamentInterfaceTransactor) Revive(opts *bind.TransactOpts, wizardId *big.Int) (*types.Transaction, error) {
	return _TournamentInterface.contract.Transact(opts, "revive", wizardId)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_TournamentInterface *TournamentInterfaceSession) Revive(wizardId *big.Int) (*types.Transaction, error) {
	return _TournamentInterface.Contract.Revive(&_TournamentInterface.TransactOpts, wizardId)
}

// Revive is a paid mutator transaction binding the contract method 0x8baecc21.
//
// Solidity: function revive(uint256 wizardId) returns()
func (_TournamentInterface *TournamentInterfaceTransactorSession) Revive(wizardId *big.Int) (*types.Transaction, error) {
	return _TournamentInterface.Contract.Revive(&_TournamentInterface.TransactOpts, wizardId)
}
