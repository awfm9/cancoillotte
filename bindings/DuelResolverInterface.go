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

// DuelResolverInterfaceABI is the input ABI used to generate the binding from.
const DuelResolverInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"affinity\",\"type\":\"uint256\"}],\"name\":\"isValidAffinity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moveSet\",\"type\":\"bytes32\"}],\"name\":\"isValidMoveSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moveSet1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"moveSet2\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"power1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affinity1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affinity2\",\"type\":\"uint256\"}],\"name\":\"resolveDuel\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DuelResolverInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var DuelResolverInterfaceFuncSigs = map[string]string{
	"e9563dee": "isValidAffinity(uint256)",
	"1823fbbc": "isValidMoveSet(bytes32)",
	"b089894c": "resolveDuel(bytes32,bytes32,uint256,uint256,uint256,uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// DuelResolverInterface is an auto generated Go binding around an Ethereum contract.
type DuelResolverInterface struct {
	DuelResolverInterfaceCaller     // Read-only binding to the contract
	DuelResolverInterfaceTransactor // Write-only binding to the contract
	DuelResolverInterfaceFilterer   // Log filterer for contract events
}

// DuelResolverInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DuelResolverInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DuelResolverInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DuelResolverInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DuelResolverInterfaceSession struct {
	Contract     *DuelResolverInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DuelResolverInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DuelResolverInterfaceCallerSession struct {
	Contract *DuelResolverInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DuelResolverInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DuelResolverInterfaceTransactorSession struct {
	Contract     *DuelResolverInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DuelResolverInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DuelResolverInterfaceRaw struct {
	Contract *DuelResolverInterface // Generic contract binding to access the raw methods on
}

// DuelResolverInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DuelResolverInterfaceCallerRaw struct {
	Contract *DuelResolverInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DuelResolverInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DuelResolverInterfaceTransactorRaw struct {
	Contract *DuelResolverInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDuelResolverInterface creates a new instance of DuelResolverInterface, bound to a specific deployed contract.
func NewDuelResolverInterface(address common.Address, backend bind.ContractBackend) (*DuelResolverInterface, error) {
	contract, err := bindDuelResolverInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterface{DuelResolverInterfaceCaller: DuelResolverInterfaceCaller{contract: contract}, DuelResolverInterfaceTransactor: DuelResolverInterfaceTransactor{contract: contract}, DuelResolverInterfaceFilterer: DuelResolverInterfaceFilterer{contract: contract}}, nil
}

// NewDuelResolverInterfaceCaller creates a new read-only instance of DuelResolverInterface, bound to a specific deployed contract.
func NewDuelResolverInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DuelResolverInterfaceCaller, error) {
	contract, err := bindDuelResolverInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceCaller{contract: contract}, nil
}

// NewDuelResolverInterfaceTransactor creates a new write-only instance of DuelResolverInterface, bound to a specific deployed contract.
func NewDuelResolverInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DuelResolverInterfaceTransactor, error) {
	contract, err := bindDuelResolverInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceTransactor{contract: contract}, nil
}

// NewDuelResolverInterfaceFilterer creates a new log filterer instance of DuelResolverInterface, bound to a specific deployed contract.
func NewDuelResolverInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DuelResolverInterfaceFilterer, error) {
	contract, err := bindDuelResolverInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceFilterer{contract: contract}, nil
}

// bindDuelResolverInterface binds a generic wrapper to an already deployed contract.
func bindDuelResolverInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DuelResolverInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DuelResolverInterface *DuelResolverInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DuelResolverInterface.Contract.DuelResolverInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DuelResolverInterface *DuelResolverInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DuelResolverInterface.Contract.DuelResolverInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DuelResolverInterface *DuelResolverInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DuelResolverInterface.Contract.DuelResolverInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DuelResolverInterface *DuelResolverInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DuelResolverInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DuelResolverInterface *DuelResolverInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DuelResolverInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DuelResolverInterface *DuelResolverInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DuelResolverInterface.Contract.contract.Transact(opts, method, params...)
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 affinity) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCaller) IsValidAffinity(opts *bind.CallOpts, affinity *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DuelResolverInterface.contract.Call(opts, out, "isValidAffinity", affinity)
	return *ret0, err
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 affinity) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceSession) IsValidAffinity(affinity *big.Int) (bool, error) {
	return _DuelResolverInterface.Contract.IsValidAffinity(&_DuelResolverInterface.CallOpts, affinity)
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 affinity) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCallerSession) IsValidAffinity(affinity *big.Int) (bool, error) {
	return _DuelResolverInterface.Contract.IsValidAffinity(&_DuelResolverInterface.CallOpts, affinity)
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCaller) IsValidMoveSet(opts *bind.CallOpts, moveSet [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DuelResolverInterface.contract.Call(opts, out, "isValidMoveSet", moveSet)
	return *ret0, err
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceSession) IsValidMoveSet(moveSet [32]byte) (bool, error) {
	return _DuelResolverInterface.Contract.IsValidMoveSet(&_DuelResolverInterface.CallOpts, moveSet)
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCallerSession) IsValidMoveSet(moveSet [32]byte) (bool, error) {
	return _DuelResolverInterface.Contract.IsValidMoveSet(&_DuelResolverInterface.CallOpts, moveSet)
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256)
func (_DuelResolverInterface *DuelResolverInterfaceCaller) ResolveDuel(opts *bind.CallOpts, moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DuelResolverInterface.contract.Call(opts, out, "resolveDuel", moveSet1, moveSet2, power1, power2, affinity1, affinity2)
	return *ret0, err
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256)
func (_DuelResolverInterface *DuelResolverInterfaceSession) ResolveDuel(moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	return _DuelResolverInterface.Contract.ResolveDuel(&_DuelResolverInterface.CallOpts, moveSet1, moveSet2, power1, power2, affinity1, affinity2)
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256)
func (_DuelResolverInterface *DuelResolverInterfaceCallerSession) ResolveDuel(moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	return _DuelResolverInterface.Contract.ResolveDuel(&_DuelResolverInterface.CallOpts, moveSet1, moveSet2, power1, power2, affinity1, affinity2)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DuelResolverInterface.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DuelResolverInterface.Contract.SupportsInterface(&_DuelResolverInterface.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_DuelResolverInterface *DuelResolverInterfaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DuelResolverInterface.Contract.SupportsInterface(&_DuelResolverInterface.CallOpts, interfaceId)
}
