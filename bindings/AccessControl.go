// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package bindings

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AccessControlABI is the input ABI used to generate the binding from.
const AccessControlABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCooAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"newCfoAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCeo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"CEOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCfo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"CFOTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousCoo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"COOTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCeo\",\"type\":\"address\"}],\"name\":\"setCeo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newCfo\",\"type\":\"address\"}],\"name\":\"setCfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCoo\",\"type\":\"address\"}],\"name\":\"setCoo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = map[string]string{
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"88975198": "setCeo(address)",
	"2d46ed56": "setCfo(address)",
	"9986a0c6": "setCoo(address)",
}

// AccessControlBin is the compiled bytecode used for deploying new contracts.
var AccessControlBin = "0x608060405234801561001057600080fd5b506040516107c73803806107c78339818101604052604081101561003357600080fd5b50805160209091015161004e336001600160e01b0361008816565b610060826001600160e01b036100f216565b6001600160a01b0381161561008157610081816001600160e01b036101e716565b5050610370565b600054604080516001600160a01b039283168152918316602083015280517f9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc69281900390910190a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331461016b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f4f6e6c792043454f000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61017d816001600160e01b036102dc16565b600154604080516001600160a01b039283168152918316602083015280517f1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe11849281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331461026057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f4f6e6c792043454f000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610272816001600160e01b036102dc16565b600254604080516001600160a01b039283168152918316602083015280517fe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d69281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381161580159061030257506000546001600160a01b03828116911614155b61036d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e76616c69642043454f206164647265737300000000000000000000000000604482015290519081900360640190fd5b50565b6104488061037f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630519ce79146100675780630a0f81681461008b5780632d46ed561461009357806388975198146100bb5780639986a0c6146100e1578063b047fb5014610107575b600080fd5b61006f61010f565b604080516001600160a01b039092168252519081900360200190f35b61006f61011e565b6100b9600480360360208110156100a957600080fd5b50356001600160a01b031661012d565b005b6100b9600480360360208110156100d157600080fd5b50356001600160a01b03166101ea565b6100b9600480360360208110156100f757600080fd5b50356001600160a01b0316610249565b61006f610306565b6002546001600160a01b031681565b6000546001600160a01b031681565b6000546001600160a01b03163314610177576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b61018081610315565b600254604080516001600160a01b039283168152918316602083015280517fe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d69281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314610234576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b61023d81610315565b61024681610382565b50565b6000546001600160a01b03163314610293576040805162461bcd60e51b81526020600482015260086024820152674f6e6c792043454f60c01b604482015290519081900360640190fd5b61029c81610315565b600154604080516001600160a01b039283168152918316602083015280517f1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe11849281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b6001600160a01b0381161580159061033b57506000546001600160a01b03828116911614155b610246576040805162461bcd60e51b8152602060048201526013602482015272496e76616c69642043454f206164647265737360681b604482015290519081900360640190fd5b600054604080516001600160a01b039283168152918316602083015280517f9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc69281900390910190a1600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820d46fe110ab8e95965bc56eb73c8a5556e5f691904c75904e11f9ef9b3eae9dff64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployAccessControl deploys a new Ethereum contract, binding an instance of AccessControl to it.
func DeployAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend, newCooAddress common.Address, newCfoAddress common.Address) (common.Address, *types.Transaction, *AccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlBin), backend, newCooAddress, newCfoAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_AccessControl *AccessControlCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "ceoAddress")
	return *ret0, err
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_AccessControl *AccessControlSession) CeoAddress() (common.Address, error) {
	return _AccessControl.Contract.CeoAddress(&_AccessControl.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() constant returns(address)
func (_AccessControl *AccessControlCallerSession) CeoAddress() (common.Address, error) {
	return _AccessControl.Contract.CeoAddress(&_AccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_AccessControl *AccessControlCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "cfoAddress")
	return *ret0, err
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_AccessControl *AccessControlSession) CfoAddress() (common.Address, error) {
	return _AccessControl.Contract.CfoAddress(&_AccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() constant returns(address)
func (_AccessControl *AccessControlCallerSession) CfoAddress() (common.Address, error) {
	return _AccessControl.Contract.CfoAddress(&_AccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_AccessControl *AccessControlCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControl.contract.Call(opts, out, "cooAddress")
	return *ret0, err
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_AccessControl *AccessControlSession) CooAddress() (common.Address, error) {
	return _AccessControl.Contract.CooAddress(&_AccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() constant returns(address)
func (_AccessControl *AccessControlCallerSession) CooAddress() (common.Address, error) {
	return _AccessControl.Contract.CooAddress(&_AccessControl.CallOpts)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_AccessControl *AccessControlTransactor) SetCeo(opts *bind.TransactOpts, newCeo common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "setCeo", newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_AccessControl *AccessControlSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCeo(&_AccessControl.TransactOpts, newCeo)
}

// SetCeo is a paid mutator transaction binding the contract method 0x88975198.
//
// Solidity: function setCeo(address newCeo) returns()
func (_AccessControl *AccessControlTransactorSession) SetCeo(newCeo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCeo(&_AccessControl.TransactOpts, newCeo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_AccessControl *AccessControlTransactor) SetCfo(opts *bind.TransactOpts, newCfo common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "setCfo", newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_AccessControl *AccessControlSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCfo(&_AccessControl.TransactOpts, newCfo)
}

// SetCfo is a paid mutator transaction binding the contract method 0x2d46ed56.
//
// Solidity: function setCfo(address newCfo) returns()
func (_AccessControl *AccessControlTransactorSession) SetCfo(newCfo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCfo(&_AccessControl.TransactOpts, newCfo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_AccessControl *AccessControlTransactor) SetCoo(opts *bind.TransactOpts, newCoo common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "setCoo", newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_AccessControl *AccessControlSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCoo(&_AccessControl.TransactOpts, newCoo)
}

// SetCoo is a paid mutator transaction binding the contract method 0x9986a0c6.
//
// Solidity: function setCoo(address newCoo) returns()
func (_AccessControl *AccessControlTransactorSession) SetCoo(newCoo common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetCoo(&_AccessControl.TransactOpts, newCoo)
}

// AccessControlCEOTransferredIterator is returned from FilterCEOTransferred and is used to iterate over the raw logs and unpacked data for CEOTransferred events raised by the AccessControl contract.
type AccessControlCEOTransferredIterator struct {
	Event *AccessControlCEOTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlCEOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlCEOTransferred)
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
		it.Event = new(AccessControlCEOTransferred)
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
func (it *AccessControlCEOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlCEOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlCEOTransferred represents a CEOTransferred event raised by the AccessControl contract.
type AccessControlCEOTransferred struct {
	PreviousCeo common.Address
	NewCeo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCEOTransferred is a free log retrieval operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_AccessControl *AccessControlFilterer) FilterCEOTransferred(opts *bind.FilterOpts) (*AccessControlCEOTransferredIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return &AccessControlCEOTransferredIterator{contract: _AccessControl.contract, event: "CEOTransferred", logs: logs, sub: sub}, nil
}

// WatchCEOTransferred is a free log subscription operation binding the contract event 0x9d05f170f1d545b1aa21c4a4f79f17ff737f5f020ea1b333d88f29f0bbfa9fc6.
//
// Solidity: event CEOTransferred(address previousCeo, address newCeo)
func (_AccessControl *AccessControlFilterer) WatchCEOTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlCEOTransferred) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "CEOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlCEOTransferred)
				if err := _AccessControl.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseCEOTransferred(log types.Log) (*AccessControlCEOTransferred, error) {
	event := new(AccessControlCEOTransferred)
	if err := _AccessControl.contract.UnpackLog(event, "CEOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlCFOTransferredIterator is returned from FilterCFOTransferred and is used to iterate over the raw logs and unpacked data for CFOTransferred events raised by the AccessControl contract.
type AccessControlCFOTransferredIterator struct {
	Event *AccessControlCFOTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlCFOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlCFOTransferred)
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
		it.Event = new(AccessControlCFOTransferred)
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
func (it *AccessControlCFOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlCFOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlCFOTransferred represents a CFOTransferred event raised by the AccessControl contract.
type AccessControlCFOTransferred struct {
	PreviousCfo common.Address
	NewCfo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCFOTransferred is a free log retrieval operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_AccessControl *AccessControlFilterer) FilterCFOTransferred(opts *bind.FilterOpts) (*AccessControlCFOTransferredIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return &AccessControlCFOTransferredIterator{contract: _AccessControl.contract, event: "CFOTransferred", logs: logs, sub: sub}, nil
}

// WatchCFOTransferred is a free log subscription operation binding the contract event 0xe1033d3cc535efc343c53636bdc05c52a44d9e70b089d4ad6e974379f2c651d6.
//
// Solidity: event CFOTransferred(address previousCfo, address newCfo)
func (_AccessControl *AccessControlFilterer) WatchCFOTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlCFOTransferred) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "CFOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlCFOTransferred)
				if err := _AccessControl.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseCFOTransferred(log types.Log) (*AccessControlCFOTransferred, error) {
	event := new(AccessControlCFOTransferred)
	if err := _AccessControl.contract.UnpackLog(event, "CFOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlCOOTransferredIterator is returned from FilterCOOTransferred and is used to iterate over the raw logs and unpacked data for COOTransferred events raised by the AccessControl contract.
type AccessControlCOOTransferredIterator struct {
	Event *AccessControlCOOTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlCOOTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlCOOTransferred)
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
		it.Event = new(AccessControlCOOTransferred)
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
func (it *AccessControlCOOTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlCOOTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlCOOTransferred represents a COOTransferred event raised by the AccessControl contract.
type AccessControlCOOTransferred struct {
	PreviousCoo common.Address
	NewCoo      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCOOTransferred is a free log retrieval operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_AccessControl *AccessControlFilterer) FilterCOOTransferred(opts *bind.FilterOpts) (*AccessControlCOOTransferredIterator, error) {

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return &AccessControlCOOTransferredIterator{contract: _AccessControl.contract, event: "COOTransferred", logs: logs, sub: sub}, nil
}

// WatchCOOTransferred is a free log subscription operation binding the contract event 0x1cd3afc04e6ae479d2b9f74533351b52218c5b2ae4f847f681a5eac514fe1184.
//
// Solidity: event COOTransferred(address previousCoo, address newCoo)
func (_AccessControl *AccessControlFilterer) WatchCOOTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlCOOTransferred) (event.Subscription, error) {

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "COOTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlCOOTransferred)
				if err := _AccessControl.contract.UnpackLog(event, "COOTransferred", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseCOOTransferred(log types.Log) (*AccessControlCOOTransferred, error) {
	event := new(AccessControlCOOTransferred)
	if err := _AccessControl.contract.UnpackLog(event, "COOTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
