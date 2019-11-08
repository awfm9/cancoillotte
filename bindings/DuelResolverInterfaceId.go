// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DuelResolverInterfaceIdABI is the input ABI used to generate the binding from.
const DuelResolverInterfaceIdABI = "[]"

// DuelResolverInterfaceIdBin is the compiled bytecode used for deploying new contracts.
var DuelResolverInterfaceIdBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a723158209c06883b2329cc0523ba73258ab334a2e4473a7f46fc1d11d3d7c1b3dc32ebfb64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployDuelResolverInterfaceId deploys a new Ethereum contract, binding an instance of DuelResolverInterfaceId to it.
func DeployDuelResolverInterfaceId(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DuelResolverInterfaceId, error) {
	parsed, err := abi.JSON(strings.NewReader(DuelResolverInterfaceIdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DuelResolverInterfaceIdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DuelResolverInterfaceId{DuelResolverInterfaceIdCaller: DuelResolverInterfaceIdCaller{contract: contract}, DuelResolverInterfaceIdTransactor: DuelResolverInterfaceIdTransactor{contract: contract}, DuelResolverInterfaceIdFilterer: DuelResolverInterfaceIdFilterer{contract: contract}}, nil
}

// DuelResolverInterfaceId is an auto generated Go binding around an Ethereum contract.
type DuelResolverInterfaceId struct {
	DuelResolverInterfaceIdCaller     // Read-only binding to the contract
	DuelResolverInterfaceIdTransactor // Write-only binding to the contract
	DuelResolverInterfaceIdFilterer   // Log filterer for contract events
}

// DuelResolverInterfaceIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type DuelResolverInterfaceIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DuelResolverInterfaceIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DuelResolverInterfaceIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DuelResolverInterfaceIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DuelResolverInterfaceIdSession struct {
	Contract     *DuelResolverInterfaceId // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DuelResolverInterfaceIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DuelResolverInterfaceIdCallerSession struct {
	Contract *DuelResolverInterfaceIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// DuelResolverInterfaceIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DuelResolverInterfaceIdTransactorSession struct {
	Contract     *DuelResolverInterfaceIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// DuelResolverInterfaceIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type DuelResolverInterfaceIdRaw struct {
	Contract *DuelResolverInterfaceId // Generic contract binding to access the raw methods on
}

// DuelResolverInterfaceIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DuelResolverInterfaceIdCallerRaw struct {
	Contract *DuelResolverInterfaceIdCaller // Generic read-only contract binding to access the raw methods on
}

// DuelResolverInterfaceIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DuelResolverInterfaceIdTransactorRaw struct {
	Contract *DuelResolverInterfaceIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDuelResolverInterfaceId creates a new instance of DuelResolverInterfaceId, bound to a specific deployed contract.
func NewDuelResolverInterfaceId(address common.Address, backend bind.ContractBackend) (*DuelResolverInterfaceId, error) {
	contract, err := bindDuelResolverInterfaceId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceId{DuelResolverInterfaceIdCaller: DuelResolverInterfaceIdCaller{contract: contract}, DuelResolverInterfaceIdTransactor: DuelResolverInterfaceIdTransactor{contract: contract}, DuelResolverInterfaceIdFilterer: DuelResolverInterfaceIdFilterer{contract: contract}}, nil
}

// NewDuelResolverInterfaceIdCaller creates a new read-only instance of DuelResolverInterfaceId, bound to a specific deployed contract.
func NewDuelResolverInterfaceIdCaller(address common.Address, caller bind.ContractCaller) (*DuelResolverInterfaceIdCaller, error) {
	contract, err := bindDuelResolverInterfaceId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceIdCaller{contract: contract}, nil
}

// NewDuelResolverInterfaceIdTransactor creates a new write-only instance of DuelResolverInterfaceId, bound to a specific deployed contract.
func NewDuelResolverInterfaceIdTransactor(address common.Address, transactor bind.ContractTransactor) (*DuelResolverInterfaceIdTransactor, error) {
	contract, err := bindDuelResolverInterfaceId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceIdTransactor{contract: contract}, nil
}

// NewDuelResolverInterfaceIdFilterer creates a new log filterer instance of DuelResolverInterfaceId, bound to a specific deployed contract.
func NewDuelResolverInterfaceIdFilterer(address common.Address, filterer bind.ContractFilterer) (*DuelResolverInterfaceIdFilterer, error) {
	contract, err := bindDuelResolverInterfaceId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DuelResolverInterfaceIdFilterer{contract: contract}, nil
}

// bindDuelResolverInterfaceId binds a generic wrapper to an already deployed contract.
func bindDuelResolverInterfaceId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DuelResolverInterfaceIdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DuelResolverInterfaceId.Contract.DuelResolverInterfaceIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DuelResolverInterfaceId.Contract.DuelResolverInterfaceIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DuelResolverInterfaceId.Contract.DuelResolverInterfaceIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DuelResolverInterfaceId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DuelResolverInterfaceId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DuelResolverInterfaceId *DuelResolverInterfaceIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DuelResolverInterfaceId.Contract.contract.Transact(opts, method, params...)
}
