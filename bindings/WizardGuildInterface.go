// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// WizardGuildInterfaceABI is the input ABI used to generate the binding from.
const WizardGuildInterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getWizard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint88\",\"name\":\"innatePower\",\"type\":\"uint88\"},{\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint88[]\",\"name\":\"powers\",\"type\":\"uint88[]\"},{\"internalType\":\"uint8[]\",\"name\":\"affinities\",\"type\":\"uint8[]\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mintReservedWizards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint88[]\",\"name\":\"powers\",\"type\":\"uint88[]\"},{\"internalType\":\"uint8[]\",\"name\":\"affinities\",\"type\":\"uint8[]\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mintWizards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"newAffinity\",\"type\":\"uint8\"}],\"name\":\"setAffinity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"wizardIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"metadata\",\"type\":\"bytes32[]\"}],\"name\":\"setMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wizardId2\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig2\",\"type\":\"bytes\"}],\"name\":\"verifySignatures\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WizardGuildInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var WizardGuildInterfaceFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"fac8eafc": "getWizard(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"430c2081": "isApprovedOrOwner(address,uint256)",
	"9d158023": "mintReservedWizards(uint256[],uint88[],uint8[],address)",
	"55fdbeec": "mintWizards(uint88[],uint8[],address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"98d7a414": "setAffinity(uint256,uint8)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"2738ec3c": "setMetadata(uint256[],bytes32[])",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2f81b15d": "verifySignature(uint256,bytes32,bytes)",
	"a096d9f0": "verifySignatures(uint256,uint256,bytes32,bytes32,bytes,bytes)",
}

// WizardGuildInterface is an auto generated Go binding around an Ethereum contract.
type WizardGuildInterface struct {
	WizardGuildInterfaceCaller     // Read-only binding to the contract
	WizardGuildInterfaceTransactor // Write-only binding to the contract
	WizardGuildInterfaceFilterer   // Log filterer for contract events
}

// WizardGuildInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardGuildInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardGuildInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardGuildInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardGuildInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardGuildInterfaceSession struct {
	Contract     *WizardGuildInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WizardGuildInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardGuildInterfaceCallerSession struct {
	Contract *WizardGuildInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// WizardGuildInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardGuildInterfaceTransactorSession struct {
	Contract     *WizardGuildInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// WizardGuildInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardGuildInterfaceRaw struct {
	Contract *WizardGuildInterface // Generic contract binding to access the raw methods on
}

// WizardGuildInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardGuildInterfaceCallerRaw struct {
	Contract *WizardGuildInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// WizardGuildInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardGuildInterfaceTransactorRaw struct {
	Contract *WizardGuildInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardGuildInterface creates a new instance of WizardGuildInterface, bound to a specific deployed contract.
func NewWizardGuildInterface(address common.Address, backend bind.ContractBackend) (*WizardGuildInterface, error) {
	contract, err := bindWizardGuildInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterface{WizardGuildInterfaceCaller: WizardGuildInterfaceCaller{contract: contract}, WizardGuildInterfaceTransactor: WizardGuildInterfaceTransactor{contract: contract}, WizardGuildInterfaceFilterer: WizardGuildInterfaceFilterer{contract: contract}}, nil
}

// NewWizardGuildInterfaceCaller creates a new read-only instance of WizardGuildInterface, bound to a specific deployed contract.
func NewWizardGuildInterfaceCaller(address common.Address, caller bind.ContractCaller) (*WizardGuildInterfaceCaller, error) {
	contract, err := bindWizardGuildInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceCaller{contract: contract}, nil
}

// NewWizardGuildInterfaceTransactor creates a new write-only instance of WizardGuildInterface, bound to a specific deployed contract.
func NewWizardGuildInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardGuildInterfaceTransactor, error) {
	contract, err := bindWizardGuildInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceTransactor{contract: contract}, nil
}

// NewWizardGuildInterfaceFilterer creates a new log filterer instance of WizardGuildInterface, bound to a specific deployed contract.
func NewWizardGuildInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardGuildInterfaceFilterer, error) {
	contract, err := bindWizardGuildInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceFilterer{contract: contract}, nil
}

// bindWizardGuildInterface binds a generic wrapper to an already deployed contract.
func bindWizardGuildInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardGuildInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardGuildInterface *WizardGuildInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardGuildInterface.Contract.WizardGuildInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardGuildInterface *WizardGuildInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.WizardGuildInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardGuildInterface *WizardGuildInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.WizardGuildInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardGuildInterface *WizardGuildInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardGuildInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardGuildInterface *WizardGuildInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardGuildInterface *WizardGuildInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_WizardGuildInterface *WizardGuildInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardGuildInterface.Contract.BalanceOf(&_WizardGuildInterface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardGuildInterface.Contract.BalanceOf(&_WizardGuildInterface.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_WizardGuildInterface *WizardGuildInterfaceSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WizardGuildInterface.Contract.GetApproved(&_WizardGuildInterface.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WizardGuildInterface.Contract.GetApproved(&_WizardGuildInterface.CallOpts, tokenId)
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 id) constant returns(address owner, uint88 innatePower, uint8 affinity, bytes32 metadata)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) GetWizard(opts *bind.CallOpts, id *big.Int) (struct {
	Owner       common.Address
	InnatePower *big.Int
	Affinity    uint8
	Metadata    [32]byte
}, error) {
	ret := new(struct {
		Owner       common.Address
		InnatePower *big.Int
		Affinity    uint8
		Metadata    [32]byte
	})
	out := ret
	err := _WizardGuildInterface.contract.Call(opts, out, "getWizard", id)
	return *ret, err
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 id) constant returns(address owner, uint88 innatePower, uint8 affinity, bytes32 metadata)
func (_WizardGuildInterface *WizardGuildInterfaceSession) GetWizard(id *big.Int) (struct {
	Owner       common.Address
	InnatePower *big.Int
	Affinity    uint8
	Metadata    [32]byte
}, error) {
	return _WizardGuildInterface.Contract.GetWizard(&_WizardGuildInterface.CallOpts, id)
}

// GetWizard is a free data retrieval call binding the contract method 0xfac8eafc.
//
// Solidity: function getWizard(uint256 id) constant returns(address owner, uint88 innatePower, uint8 affinity, bytes32 metadata)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) GetWizard(id *big.Int) (struct {
	Owner       common.Address
	InnatePower *big.Int
	Affinity    uint8
	Metadata    [32]byte
}, error) {
	return _WizardGuildInterface.Contract.GetWizard(&_WizardGuildInterface.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardGuildInterface.Contract.IsApprovedForAll(&_WizardGuildInterface.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardGuildInterface.Contract.IsApprovedForAll(&_WizardGuildInterface.CallOpts, owner, operator)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) IsApprovedOrOwner(opts *bind.CallOpts, spender common.Address, tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "isApprovedOrOwner", spender, tokenId)
	return *ret0, err
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _WizardGuildInterface.Contract.IsApprovedOrOwner(&_WizardGuildInterface.CallOpts, spender, tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _WizardGuildInterface.Contract.IsApprovedOrOwner(&_WizardGuildInterface.CallOpts, spender, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_WizardGuildInterface *WizardGuildInterfaceSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WizardGuildInterface.Contract.OwnerOf(&_WizardGuildInterface.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WizardGuildInterface.Contract.OwnerOf(&_WizardGuildInterface.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardGuildInterface.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardGuildInterface.Contract.SupportsInterface(&_WizardGuildInterface.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardGuildInterface.Contract.SupportsInterface(&_WizardGuildInterface.CallOpts, interfaceId)
}

// VerifySignature is a free data retrieval call binding the contract method 0x2f81b15d.
//
// Solidity: function verifySignature(uint256 wizardId, bytes32 hash, bytes sig) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceCaller) VerifySignature(opts *bind.CallOpts, wizardId *big.Int, hash [32]byte, sig []byte) error {
	var ()
	out := &[]interface{}{}
	err := _WizardGuildInterface.contract.Call(opts, out, "verifySignature", wizardId, hash, sig)
	return err
}

// VerifySignature is a free data retrieval call binding the contract method 0x2f81b15d.
//
// Solidity: function verifySignature(uint256 wizardId, bytes32 hash, bytes sig) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) VerifySignature(wizardId *big.Int, hash [32]byte, sig []byte) error {
	return _WizardGuildInterface.Contract.VerifySignature(&_WizardGuildInterface.CallOpts, wizardId, hash, sig)
}

// VerifySignature is a free data retrieval call binding the contract method 0x2f81b15d.
//
// Solidity: function verifySignature(uint256 wizardId, bytes32 hash, bytes sig) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) VerifySignature(wizardId *big.Int, hash [32]byte, sig []byte) error {
	return _WizardGuildInterface.Contract.VerifySignature(&_WizardGuildInterface.CallOpts, wizardId, hash, sig)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa096d9f0.
//
// Solidity: function verifySignatures(uint256 wizardId1, uint256 wizardId2, bytes32 hash1, bytes32 hash2, bytes sig1, bytes sig2) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceCaller) VerifySignatures(opts *bind.CallOpts, wizardId1 *big.Int, wizardId2 *big.Int, hash1 [32]byte, hash2 [32]byte, sig1 []byte, sig2 []byte) error {
	var ()
	out := &[]interface{}{}
	err := _WizardGuildInterface.contract.Call(opts, out, "verifySignatures", wizardId1, wizardId2, hash1, hash2, sig1, sig2)
	return err
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa096d9f0.
//
// Solidity: function verifySignatures(uint256 wizardId1, uint256 wizardId2, bytes32 hash1, bytes32 hash2, bytes sig1, bytes sig2) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) VerifySignatures(wizardId1 *big.Int, wizardId2 *big.Int, hash1 [32]byte, hash2 [32]byte, sig1 []byte, sig2 []byte) error {
	return _WizardGuildInterface.Contract.VerifySignatures(&_WizardGuildInterface.CallOpts, wizardId1, wizardId2, hash1, hash2, sig1, sig2)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xa096d9f0.
//
// Solidity: function verifySignatures(uint256 wizardId1, uint256 wizardId2, bytes32 hash1, bytes32 hash2, bytes sig1, bytes sig2) constant returns()
func (_WizardGuildInterface *WizardGuildInterfaceCallerSession) VerifySignatures(wizardId1 *big.Int, wizardId2 *big.Int, hash1 [32]byte, hash2 [32]byte, sig1 []byte, sig2 []byte) error {
	return _WizardGuildInterface.Contract.VerifySignatures(&_WizardGuildInterface.CallOpts, wizardId1, wizardId2, hash1, hash2, sig1, sig2)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.Approve(&_WizardGuildInterface.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.Approve(&_WizardGuildInterface.TransactOpts, to, tokenId)
}

// MintReservedWizards is a paid mutator transaction binding the contract method 0x9d158023.
//
// Solidity: function mintReservedWizards(uint256[] wizardIds, uint88[] powers, uint8[] affinities, address owner) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) MintReservedWizards(opts *bind.TransactOpts, wizardIds []*big.Int, powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "mintReservedWizards", wizardIds, powers, affinities, owner)
}

// MintReservedWizards is a paid mutator transaction binding the contract method 0x9d158023.
//
// Solidity: function mintReservedWizards(uint256[] wizardIds, uint88[] powers, uint8[] affinities, address owner) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) MintReservedWizards(wizardIds []*big.Int, powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.MintReservedWizards(&_WizardGuildInterface.TransactOpts, wizardIds, powers, affinities, owner)
}

// MintReservedWizards is a paid mutator transaction binding the contract method 0x9d158023.
//
// Solidity: function mintReservedWizards(uint256[] wizardIds, uint88[] powers, uint8[] affinities, address owner) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) MintReservedWizards(wizardIds []*big.Int, powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.MintReservedWizards(&_WizardGuildInterface.TransactOpts, wizardIds, powers, affinities, owner)
}

// MintWizards is a paid mutator transaction binding the contract method 0x55fdbeec.
//
// Solidity: function mintWizards(uint88[] powers, uint8[] affinities, address owner) returns(uint256[] wizardIds)
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) MintWizards(opts *bind.TransactOpts, powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "mintWizards", powers, affinities, owner)
}

// MintWizards is a paid mutator transaction binding the contract method 0x55fdbeec.
//
// Solidity: function mintWizards(uint88[] powers, uint8[] affinities, address owner) returns(uint256[] wizardIds)
func (_WizardGuildInterface *WizardGuildInterfaceSession) MintWizards(powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.MintWizards(&_WizardGuildInterface.TransactOpts, powers, affinities, owner)
}

// MintWizards is a paid mutator transaction binding the contract method 0x55fdbeec.
//
// Solidity: function mintWizards(uint88[] powers, uint8[] affinities, address owner) returns(uint256[] wizardIds)
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) MintWizards(powers []*big.Int, affinities []uint8, owner common.Address) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.MintWizards(&_WizardGuildInterface.TransactOpts, powers, affinities, owner)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SafeTransferFrom(&_WizardGuildInterface.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SafeTransferFrom(&_WizardGuildInterface.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SafeTransferFrom0(&_WizardGuildInterface.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SafeTransferFrom0(&_WizardGuildInterface.TransactOpts, from, to, tokenId, data)
}

// SetAffinity is a paid mutator transaction binding the contract method 0x98d7a414.
//
// Solidity: function setAffinity(uint256 wizardId, uint8 newAffinity) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) SetAffinity(opts *bind.TransactOpts, wizardId *big.Int, newAffinity uint8) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "setAffinity", wizardId, newAffinity)
}

// SetAffinity is a paid mutator transaction binding the contract method 0x98d7a414.
//
// Solidity: function setAffinity(uint256 wizardId, uint8 newAffinity) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) SetAffinity(wizardId *big.Int, newAffinity uint8) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetAffinity(&_WizardGuildInterface.TransactOpts, wizardId, newAffinity)
}

// SetAffinity is a paid mutator transaction binding the contract method 0x98d7a414.
//
// Solidity: function setAffinity(uint256 wizardId, uint8 newAffinity) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) SetAffinity(wizardId *big.Int, newAffinity uint8) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetAffinity(&_WizardGuildInterface.TransactOpts, wizardId, newAffinity)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetApprovalForAll(&_WizardGuildInterface.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetApprovalForAll(&_WizardGuildInterface.TransactOpts, operator, _approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x2738ec3c.
//
// Solidity: function setMetadata(uint256[] wizardIds, bytes32[] metadata) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) SetMetadata(opts *bind.TransactOpts, wizardIds []*big.Int, metadata [][32]byte) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "setMetadata", wizardIds, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x2738ec3c.
//
// Solidity: function setMetadata(uint256[] wizardIds, bytes32[] metadata) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) SetMetadata(wizardIds []*big.Int, metadata [][32]byte) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetMetadata(&_WizardGuildInterface.TransactOpts, wizardIds, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x2738ec3c.
//
// Solidity: function setMetadata(uint256[] wizardIds, bytes32[] metadata) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) SetMetadata(wizardIds []*big.Int, metadata [][32]byte) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.SetMetadata(&_WizardGuildInterface.TransactOpts, wizardIds, metadata)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.TransferFrom(&_WizardGuildInterface.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardGuildInterface *WizardGuildInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardGuildInterface.Contract.TransferFrom(&_WizardGuildInterface.TransactOpts, from, to, tokenId)
}

// WizardGuildInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WizardGuildInterface contract.
type WizardGuildInterfaceApprovalIterator struct {
	Event *WizardGuildInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *WizardGuildInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardGuildInterfaceApproval)
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
		it.Event = new(WizardGuildInterfaceApproval)
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
func (it *WizardGuildInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardGuildInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardGuildInterfaceApproval represents a Approval event raised by the WizardGuildInterface contract.
type WizardGuildInterfaceApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WizardGuildInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceApprovalIterator{contract: _WizardGuildInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WizardGuildInterfaceApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardGuildInterfaceApproval)
				if err := _WizardGuildInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) ParseApproval(log types.Log) (*WizardGuildInterfaceApproval, error) {
	event := new(WizardGuildInterfaceApproval)
	if err := _WizardGuildInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardGuildInterfaceApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the WizardGuildInterface contract.
type WizardGuildInterfaceApprovalForAllIterator struct {
	Event *WizardGuildInterfaceApprovalForAll // Event containing the contract specifics and raw log

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
func (it *WizardGuildInterfaceApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardGuildInterfaceApprovalForAll)
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
		it.Event = new(WizardGuildInterfaceApprovalForAll)
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
func (it *WizardGuildInterfaceApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardGuildInterfaceApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardGuildInterfaceApprovalForAll represents a ApprovalForAll event raised by the WizardGuildInterface contract.
type WizardGuildInterfaceApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WizardGuildInterfaceApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceApprovalForAllIterator{contract: _WizardGuildInterface.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WizardGuildInterfaceApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardGuildInterfaceApprovalForAll)
				if err := _WizardGuildInterface.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) ParseApprovalForAll(log types.Log) (*WizardGuildInterfaceApprovalForAll, error) {
	event := new(WizardGuildInterfaceApprovalForAll)
	if err := _WizardGuildInterface.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardGuildInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WizardGuildInterface contract.
type WizardGuildInterfaceTransferIterator struct {
	Event *WizardGuildInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *WizardGuildInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardGuildInterfaceTransfer)
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
		it.Event = new(WizardGuildInterfaceTransfer)
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
func (it *WizardGuildInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardGuildInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardGuildInterfaceTransfer represents a Transfer event raised by the WizardGuildInterface contract.
type WizardGuildInterfaceTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WizardGuildInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardGuildInterfaceTransferIterator{contract: _WizardGuildInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WizardGuildInterfaceTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardGuildInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardGuildInterfaceTransfer)
				if err := _WizardGuildInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardGuildInterface *WizardGuildInterfaceFilterer) ParseTransfer(log types.Log) (*WizardGuildInterfaceTransfer, error) {
	event := new(WizardGuildInterfaceTransfer)
	if err := _WizardGuildInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
