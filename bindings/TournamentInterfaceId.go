// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TournamentInterfaceIdABI is the input ABI used to generate the binding from.
const TournamentInterfaceIdABI = "[]"

// TournamentInterfaceIdBin is the compiled bytecode used for deploying new contracts.
var TournamentInterfaceIdBin = "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052600080fdfea265627a7a7231582058f63ff194e4b0783f77029328907a72133009127ffe67e71eb9a4336f99ea1064736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployTournamentInterfaceId deploys a new Ethereum contract, binding an instance of TournamentInterfaceId to it.
func DeployTournamentInterfaceId(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TournamentInterfaceId, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentInterfaceIdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TournamentInterfaceIdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TournamentInterfaceId{TournamentInterfaceIdCaller: TournamentInterfaceIdCaller{contract: contract}, TournamentInterfaceIdTransactor: TournamentInterfaceIdTransactor{contract: contract}, TournamentInterfaceIdFilterer: TournamentInterfaceIdFilterer{contract: contract}}, nil
}

// TournamentInterfaceId is an auto generated Go binding around an Ethereum contract.
type TournamentInterfaceId struct {
	TournamentInterfaceIdCaller     // Read-only binding to the contract
	TournamentInterfaceIdTransactor // Write-only binding to the contract
	TournamentInterfaceIdFilterer   // Log filterer for contract events
}

// TournamentInterfaceIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type TournamentInterfaceIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TournamentInterfaceIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TournamentInterfaceIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentInterfaceIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TournamentInterfaceIdSession struct {
	Contract     *TournamentInterfaceId // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TournamentInterfaceIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TournamentInterfaceIdCallerSession struct {
	Contract *TournamentInterfaceIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TournamentInterfaceIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TournamentInterfaceIdTransactorSession struct {
	Contract     *TournamentInterfaceIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TournamentInterfaceIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type TournamentInterfaceIdRaw struct {
	Contract *TournamentInterfaceId // Generic contract binding to access the raw methods on
}

// TournamentInterfaceIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TournamentInterfaceIdCallerRaw struct {
	Contract *TournamentInterfaceIdCaller // Generic read-only contract binding to access the raw methods on
}

// TournamentInterfaceIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TournamentInterfaceIdTransactorRaw struct {
	Contract *TournamentInterfaceIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTournamentInterfaceId creates a new instance of TournamentInterfaceId, bound to a specific deployed contract.
func NewTournamentInterfaceId(address common.Address, backend bind.ContractBackend) (*TournamentInterfaceId, error) {
	contract, err := bindTournamentInterfaceId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceId{TournamentInterfaceIdCaller: TournamentInterfaceIdCaller{contract: contract}, TournamentInterfaceIdTransactor: TournamentInterfaceIdTransactor{contract: contract}, TournamentInterfaceIdFilterer: TournamentInterfaceIdFilterer{contract: contract}}, nil
}

// NewTournamentInterfaceIdCaller creates a new read-only instance of TournamentInterfaceId, bound to a specific deployed contract.
func NewTournamentInterfaceIdCaller(address common.Address, caller bind.ContractCaller) (*TournamentInterfaceIdCaller, error) {
	contract, err := bindTournamentInterfaceId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceIdCaller{contract: contract}, nil
}

// NewTournamentInterfaceIdTransactor creates a new write-only instance of TournamentInterfaceId, bound to a specific deployed contract.
func NewTournamentInterfaceIdTransactor(address common.Address, transactor bind.ContractTransactor) (*TournamentInterfaceIdTransactor, error) {
	contract, err := bindTournamentInterfaceId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceIdTransactor{contract: contract}, nil
}

// NewTournamentInterfaceIdFilterer creates a new log filterer instance of TournamentInterfaceId, bound to a specific deployed contract.
func NewTournamentInterfaceIdFilterer(address common.Address, filterer bind.ContractFilterer) (*TournamentInterfaceIdFilterer, error) {
	contract, err := bindTournamentInterfaceId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TournamentInterfaceIdFilterer{contract: contract}, nil
}

// bindTournamentInterfaceId binds a generic wrapper to an already deployed contract.
func bindTournamentInterfaceId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentInterfaceIdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentInterfaceId *TournamentInterfaceIdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentInterfaceId.Contract.TournamentInterfaceIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentInterfaceId *TournamentInterfaceIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentInterfaceId.Contract.TournamentInterfaceIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentInterfaceId *TournamentInterfaceIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentInterfaceId.Contract.TournamentInterfaceIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentInterfaceId *TournamentInterfaceIdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentInterfaceId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentInterfaceId *TournamentInterfaceIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentInterfaceId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentInterfaceId *TournamentInterfaceIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentInterfaceId.Contract.contract.Transact(opts, method, params...)
}
