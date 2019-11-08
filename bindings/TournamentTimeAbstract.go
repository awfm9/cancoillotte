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

// TournamentTimeAbstractABI is the input ABI used to generate the binding from.
const TournamentTimeAbstractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cooAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"tournamentStartBlock\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"admissionDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"revivalDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"ascensionDuration\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"fightDuration\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"cullingDuration\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"duelTimeoutDuration\",\"type\":\"uint24\"},{\"internalType\":\"uint88\",\"name\":\"blueMoldBasePower\",\"type\":\"uint88\"},{\"internalType\":\"uint24\",\"name\":\"sessionsBetweenMoldDoubling\",\"type\":\"uint24\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCeo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"CEOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCfo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"CFOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCoo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"COOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseEndedBlock\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlueMoldParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimeParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tournamentStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pauseEndedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"admissionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revivalDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duelTimeoutDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ascensionWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ascensionWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fightWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fightWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolutionWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resolutionWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cullingWindowStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cullingWindowDuration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseDuration\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"setCeo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"setCfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"setCoo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TournamentTimeAbstractFuncSigs maps the 4-byte function signature to its string representation.
var TournamentTimeAbstractFuncSigs = map[string]string{
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"079cfa79": "getBlueMoldParameters()",
	"92420c90": "getTimeParameters()",
	"b187bd26": "isPaused()",
	"136439dd": "pause(uint256)",
	"88975198": "setCeo(address)",
	"2d46ed56": "setCfo(address)",
	"9986a0c6": "setCoo(address)",
}

// TournamentTimeAbstract is an auto generated Go binding around an Ethereum contract.
type TournamentTimeAbstract struct {
	TournamentTimeAbstractCaller     // Read-only binding to the contract
	TournamentTimeAbstractTransactor // Write-only binding to the contract
	TournamentTimeAbstractFilterer   // Log filterer for contract events
}

// TournamentTimeAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TournamentTimeAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentTimeAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TournamentTimeAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentTimeAbstractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TournamentTimeAbstractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentTimeAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TournamentTimeAbstractSession struct {
	Contract     *TournamentTimeAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TournamentTimeAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TournamentTimeAbstractCallerSession struct {
	Contract *TournamentTimeAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TournamentTimeAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TournamentTimeAbstractTransactorSession struct {
	Contract     *TournamentTimeAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TournamentTimeAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TournamentTimeAbstractRaw struct {
	Contract *TournamentTimeAbstract // Generic contract binding to access the raw methods on
}

// TournamentTimeAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TournamentTimeAbstractCallerRaw struct {
	Contract *TournamentTimeAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// TournamentTimeAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TournamentTimeAbstractTransactorRaw struct {
	Contract *TournamentTimeAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTournamentTimeAbstract creates a new instance of TournamentTimeAbstract, bound to a specific deployed contract.
func NewTournamentTimeAbstract(address common.Address, backend bind.ContractBackend) (*TournamentTimeAbstract, error) {
	contract, err := bindTournamentTimeAbstract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstract{TournamentTimeAbstractCaller: TournamentTimeAbstractCaller{contract: contract}, TournamentTimeAbstractTransactor: TournamentTimeAbstractTransactor{contract: contract}, TournamentTimeAbstractFilterer: TournamentTimeAbstractFilterer{contract: contract}}, nil
}

// NewTournamentTimeAbstractCaller creates a new read-only instance of TournamentTimeAbstract, bound to a specific deployed contract.
func NewTournamentTimeAbstractCaller(address common.Address, caller bind.ContractCaller) (*TournamentTimeAbstractCaller, error) {
	contract, err := bindTournamentTimeAbstract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractCaller{contract: contract}, nil
}

// NewTournamentTimeAbstractTransactor creates a new write-only instance of TournamentTimeAbstract, bound to a specific deployed contract.
func NewTournamentTimeAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*TournamentTimeAbstractTransactor, error) {
	contract, err := bindTournamentTimeAbstract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractTransactor{contract: contract}, nil
}

// NewTournamentTimeAbstractFilterer creates a new log filterer instance of TournamentTimeAbstract, bound to a specific deployed contract.
func NewTournamentTimeAbstractFilterer(address common.Address, filterer bind.ContractFilterer) (*TournamentTimeAbstractFilterer, error) {
	contract, err := bindTournamentTimeAbstract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractFilterer{contract: contract}, nil
}

// bindTournamentTimeAbstract binds a generic wrapper to an already deployed contract.
func bindTournamentTimeAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentTimeAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentTimeAbstract *TournamentTimeAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentTimeAbstract.Contract.TournamentTimeAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentTimeAbstract *TournamentTimeAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.TournamentTimeAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentTimeAbstract *TournamentTimeAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.TournamentTimeAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentTimeAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentTimeAbstract.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) CeoAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CeoAddress(&_TournamentTimeAbstract.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) CeoAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CeoAddress(&_TournamentTimeAbstract.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentTimeAbstract.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) CfoAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CfoAddress(&_TournamentTimeAbstract.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) CfoAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CfoAddress(&_TournamentTimeAbstract.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentTimeAbstract.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) CooAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CooAddress(&_TournamentTimeAbstract.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) CooAddress() (common.Address, error) {
	return _TournamentTimeAbstract.Contract.CooAddress(&_TournamentTimeAbstract.CallOpts)
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) GetBlueMoldParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _TournamentTimeAbstract.contract.Call(opts, out, "getBlueMoldParameters")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) GetBlueMoldParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _TournamentTimeAbstract.Contract.GetBlueMoldParameters(&_TournamentTimeAbstract.CallOpts)
}

// GetBlueMoldParameters is a free data retrieval call binding the contract method 0x079cfa79.
//
// Solidity: function getBlueMoldParameters() constant returns(uint256, uint256, uint256, uint256)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) GetBlueMoldParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _TournamentTimeAbstract.Contract.GetBlueMoldParameters(&_TournamentTimeAbstract.CallOpts)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) GetTimeParameters(opts *bind.CallOpts) (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	ret := new(struct {
		TournamentStartBlock     *big.Int
		PauseEndedBlock          *big.Int
		AdmissionDuration        *big.Int
		RevivalDuration          *big.Int
		DuelTimeoutDuration      *big.Int
		AscensionWindowStart     *big.Int
		AscensionWindowDuration  *big.Int
		FightWindowStart         *big.Int
		FightWindowDuration      *big.Int
		ResolutionWindowStart    *big.Int
		ResolutionWindowDuration *big.Int
		CullingWindowStart       *big.Int
		CullingWindowDuration    *big.Int
	})
	out := ret
	err := _TournamentTimeAbstract.contract.Call(opts, out, "getTimeParameters")
	return *ret, err
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) GetTimeParameters() (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	return _TournamentTimeAbstract.Contract.GetTimeParameters(&_TournamentTimeAbstract.CallOpts)
}

// GetTimeParameters is a free data retrieval call binding the contract method 0x92420c90.
//
// Solidity: function getTimeParameters() constant returns(uint256 tournamentStartBlock, uint256 pauseEndedBlock, uint256 admissionDuration, uint256 revivalDuration, uint256 duelTimeoutDuration, uint256 ascensionWindowStart, uint256 ascensionWindowDuration, uint256 fightWindowStart, uint256 fightWindowDuration, uint256 resolutionWindowStart, uint256 resolutionWindowDuration, uint256 cullingWindowStart, uint256 cullingWindowDuration)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) GetTimeParameters() (struct {
	TournamentStartBlock     *big.Int
	PauseEndedBlock          *big.Int
	AdmissionDuration        *big.Int
	RevivalDuration          *big.Int
	DuelTimeoutDuration      *big.Int
	AscensionWindowStart     *big.Int
	AscensionWindowDuration  *big.Int
	FightWindowStart         *big.Int
	FightWindowDuration      *big.Int
	ResolutionWindowStart    *big.Int
	ResolutionWindowDuration *big.Int
	CullingWindowStart       *big.Int
	CullingWindowDuration    *big.Int
}, error) {
	return _TournamentTimeAbstract.Contract.GetTimeParameters(&_TournamentTimeAbstract.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_TournamentTimeAbstract *TournamentTimeAbstractCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentTimeAbstract.contract.Call(opts, out, "isPaused")
	return *ret0, err
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) IsPaused() (bool, error) {
	return _TournamentTimeAbstract.Contract.IsPaused(&_TournamentTimeAbstract.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() constant returns(bool)
func (_TournamentTimeAbstract *TournamentTimeAbstractCallerSession) IsPaused() (bool, error) {
	return _TournamentTimeAbstract.Contract.IsPaused(&_TournamentTimeAbstract.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactor) Pause(opts *bind.TransactOpts, pauseDuration *big.Int) (*types.Transaction, error) {
	return _TournamentTimeAbstract.contract.Transact(opts, "pause", pauseDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) Pause(pauseDuration *big.Int) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.Pause(&_TournamentTimeAbstract.TransactOpts, pauseDuration)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 pauseDuration) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorSession) Pause(pauseDuration *big.Int) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.Pause(&_TournamentTimeAbstract.TransactOpts, pauseDuration)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactor) SetCeo(opts *bind.TransactOpts, newCeo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.contract.Transact(opts, "setCeo", newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCeo(&_TournamentTimeAbstract.TransactOpts, newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCeo(&_TournamentTimeAbstract.TransactOpts, newCeo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactor) SetCfo(opts *bind.TransactOpts, newCfo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.contract.Transact(opts, "setCfo", newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCfo(&_TournamentTimeAbstract.TransactOpts, newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCfo(&_TournamentTimeAbstract.TransactOpts, newCfo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactor) SetCoo(opts *bind.TransactOpts, newCoo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.contract.Transact(opts, "setCoo", newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCoo(&_TournamentTimeAbstract.TransactOpts, newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_TournamentTimeAbstract *TournamentTimeAbstractTransactorSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _TournamentTimeAbstract.Contract.SetCoo(&_TournamentTimeAbstract.TransactOpts, newCoo)
}

// TournamentTimeAbstractCEOTransferredIterator is returned from FilterCEOTransferred and is used to iterate over the raw logs and unpacked data for CEOTransferred events raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCEOTransferredIterator struct {
	Event *TournamentTimeAbstractCEOTransferred // Event containing the contract specifics and raw log

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
func (it *TournamentTimeAbstractCEOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentTimeAbstractCEOTransferred)
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
		it.Event = new(TournamentTimeAbstractCEOTransferred)
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
func (it *TournamentTimeAbstractCEOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentTimeAbstractCEOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentTimeAbstractCEOTransferred represents a CEOTransferred event raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCEOTransferred struct {
	PreviousCeo common.Address
	NewCeo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCEOTransferred is a free log retrieval operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) FilterCEOTransferred(opts *bind.FilterOpts) (*TournamentTimeAbstractCEOTransferredIterator, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.FilterLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractCEOTransferredIterator{contract: _TournamentTimeAbstract.contract, event: "CEOTransferred", logs: logs, sub: sub}, nil
}

// WatchCEOTransferred is a free log subscription operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) WatchCEOTransferred(opts *bind.WatchOpts, sink chan<- *TournamentTimeAbstractCEOTransferred) (event.Subscription, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.WatchLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentTimeAbstractCEOTransferred)
				if err := _TournamentTimeAbstract.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
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

// ParseCEOTransferred is a log parse operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) ParseCEOTransferred(log types.Log) (*TournamentTimeAbstractCEOTransferred, error) {
	event := new(TournamentTimeAbstractCEOTransferred)
	if err := _TournamentTimeAbstract.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TournamentTimeAbstractCFOTransferredIterator is returned from FilterCFOTransferred and is used to iterate over the raw logs and unpacked data for CFOTransferred events raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCFOTransferredIterator struct {
	Event *TournamentTimeAbstractCFOTransferred // Event containing the contract specifics and raw log

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
func (it *TournamentTimeAbstractCFOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentTimeAbstractCFOTransferred)
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
		it.Event = new(TournamentTimeAbstractCFOTransferred)
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
func (it *TournamentTimeAbstractCFOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentTimeAbstractCFOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentTimeAbstractCFOTransferred represents a CFOTransferred event raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCFOTransferred struct {
	PreviousCfo common.Address
	NewCfo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCFOTransferred is a free log retrieval operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) FilterCFOTransferred(opts *bind.FilterOpts) (*TournamentTimeAbstractCFOTransferredIterator, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.FilterLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractCFOTransferredIterator{contract: _TournamentTimeAbstract.contract, event: "CFOTransferred", logs: logs, sub: sub}, nil
}

// WatchCFOTransferred is a free log subscription operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) WatchCFOTransferred(opts *bind.WatchOpts, sink chan<- *TournamentTimeAbstractCFOTransferred) (event.Subscription, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.WatchLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentTimeAbstractCFOTransferred)
				if err := _TournamentTimeAbstract.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
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

// ParseCFOTransferred is a log parse operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) ParseCFOTransferred(log types.Log) (*TournamentTimeAbstractCFOTransferred, error) {
	event := new(TournamentTimeAbstractCFOTransferred)
	if err := _TournamentTimeAbstract.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TournamentTimeAbstractCOOTransferredIterator is returned from FilterCOOTransferred and is used to iterate over the raw logs and unpacked data for COOTransferred events raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCOOTransferredIterator struct {
	Event *TournamentTimeAbstractCOOTransferred // Event containing the contract specifics and raw log

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
func (it *TournamentTimeAbstractCOOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentTimeAbstractCOOTransferred)
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
		it.Event = new(TournamentTimeAbstractCOOTransferred)
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
func (it *TournamentTimeAbstractCOOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentTimeAbstractCOOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentTimeAbstractCOOTransferred represents a COOTransferred event raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractCOOTransferred struct {
	PreviousCoo common.Address
	NewCoo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCOOTransferred is a free log retrieval operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) FilterCOOTransferred(opts *bind.FilterOpts) (*TournamentTimeAbstractCOOTransferredIterator, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.FilterLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractCOOTransferredIterator{contract: _TournamentTimeAbstract.contract, event: "COOTransferred", logs: logs, sub: sub}, nil
}

// WatchCOOTransferred is a free log subscription operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) WatchCOOTransferred(opts *bind.WatchOpts, sink chan<- *TournamentTimeAbstractCOOTransferred) (event.Subscription, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.WatchLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentTimeAbstractCOOTransferred)
				if err := _TournamentTimeAbstract.contract.UnpackLog(event, "COOTransferred", log); err != nil {
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

// ParseCOOTransferred is a log parse operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) ParseCOOTransferred(log types.Log) (*TournamentTimeAbstractCOOTransferred, error) {
	event := new(TournamentTimeAbstractCOOTransferred)
	if err := _TournamentTimeAbstract.contract.UnpackLog(event, "COOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TournamentTimeAbstractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractPausedIterator struct {
	Event *TournamentTimeAbstractPaused // Event containing the contract specifics and raw log

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
func (it *TournamentTimeAbstractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentTimeAbstractPaused)
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
		it.Event = new(TournamentTimeAbstractPaused)
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
func (it *TournamentTimeAbstractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentTimeAbstractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentTimeAbstractPaused represents a Paused event raised by the TournamentTimeAbstract contract.
type TournamentTimeAbstractPaused struct {
	PauseEndedBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) FilterPaused(opts *bind.FilterOpts) (*TournamentTimeAbstractPausedIterator, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TournamentTimeAbstractPausedIterator{contract: _TournamentTimeAbstract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TournamentTimeAbstractPaused) (event.Subscription, error) {

	logs, sub, err := _TournamentTimeAbstract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentTimeAbstractPaused)
				if err := _TournamentTimeAbstract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pauseEndedBlock)
func (_TournamentTimeAbstract *TournamentTimeAbstractFilterer) ParsePaused(log types.Log) (*TournamentTimeAbstractPaused, error) {
	event := new(TournamentTimeAbstractPaused)
	if err := _TournamentTimeAbstract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}
