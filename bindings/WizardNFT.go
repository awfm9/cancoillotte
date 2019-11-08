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

// WizardNFTABI is the input ABI used to generate the binding from.
const WizardNFTABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"}],\"name\":\"WizardAffinityAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"innatePower\",\"type\":\"uint256\"}],\"name\":\"WizardConjured\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wizardId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wizardsById\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"},{\"internalType\":\"uint88\",\"name\":\"innatePower\",\"type\":\"uint88\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WizardNFTFuncSigs maps the 4-byte function signature to its string representation.
var WizardNFTFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"5938d97a": "wizardsById(uint256)",
}

// WizardNFTBin is the compiled bytecode used for deploying new contracts.
var WizardNFTBin = "0x608060405234801561001057600080fd5b50610d6c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635938d97a116100715780635938d97a146101bc5780636352211e1461021757806370a0823114610234578063a22cb4651461026c578063b88d4fde1461029a578063e985e9c514610360576100a9565b806301ffc9a7146100ae578063081812fc146100e9578063095ea7b31461012257806323b872dd1461015057806342842e0e14610186575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b03191661038e565b604080519115158252519081900360200190f35b610106600480360360208110156100ff57600080fd5b50356103c5565b604080516001600160a01b039092168252519081900360200190f35b61014e6004803603604081101561013857600080fd5b506001600160a01b038135169060200135610427565b005b61014e6004803603606081101561016657600080fd5b506001600160a01b03813581169160208101359091169060400135610538565b61014e6004803603606081101561019c57600080fd5b506001600160a01b0381358116916020810135909116906040013561058d565b6101d9600480360360208110156101d257600080fd5b50356105a8565b6040805160ff90951685526affffffffffffffffffffff90931660208501526001600160a01b03909116838301526060830152519081900360800190f35b6101066004803603602081101561022d57600080fd5b50356105eb565b61025a6004803603602081101561024a57600080fd5b50356001600160a01b0316610646565b60408051918252519081900360200190f35b61014e6004803603604081101561028257600080fd5b506001600160a01b03813516906020013515156106a9565b61014e600480360360808110156102b057600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102eb57600080fd5b8201836020820111156102fd57600080fd5b8035906020019184600183028401116401000000008311171561031f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610775945050505050565b6100d56004803603604081101561037657600080fd5b506001600160a01b03813581169160200135166107cd565b60006001600160e01b031982166301ffc9a760e01b14806103bf57506001600160e01b031982166380ac58cd60e01b145b92915050565b60006103d0826107fb565b61040b5760405162461bcd60e51b815260040180806020018281038252602c815260200180610c6a602c913960400191505060405180910390fd5b506000908152600160205260409020546001600160a01b031690565b6000610432826105eb565b9050806001600160a01b0316836001600160a01b031614156104855760405162461bcd60e51b8152600401808060200182810382526021815260200180610cbf6021913960400191505060405180910390fd5b336001600160a01b03821614806104a157506104a181336107cd565b6104dc5760405162461bcd60e51b8152600401808060200182810382526038815260200180610bdf6038913960400191505060405180910390fd5b60008281526001602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b610542338261081f565b61057d5760405162461bcd60e51b8152600401808060200182810382526031815260200180610ce06031913960400191505060405180910390fd5b6105888383836108c3565b505050565b61058883838360405180602001604052806000815250610775565b6000602081905290815260409020805460019091015460ff82169161010081046affffffffffffffffffffff1691600160601b9091046001600160a01b03169084565b600081815260208190526040812054600160601b90046001600160a01b0316806103bf5760405162461bcd60e51b8152600401808060200182810382526029815260200180610c416029913960400191505060405180910390fd5b60006001600160a01b03821661068d5760405162461bcd60e51b815260040180806020018281038252602a815260200180610c17602a913960400191505060405180910390fd5b506001600160a01b031660009081526002602052604090205490565b6001600160a01b038216331415610707576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b3360008181526003602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610780848484610538565b61078c848484846109ef565b6107c75760405162461bcd60e51b8152600401808060200182810382526032815260200180610b5d6032913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260036020908152604080832093909416825291909152205460ff1690565b600090815260208190526040902054600160601b90046001600160a01b0316151590565b600061082a826107fb565b6108655760405162461bcd60e51b815260040180806020018281038252602c815260200180610bb3602c913960400191505060405180910390fd5b6000610870836105eb565b9050806001600160a01b0316846001600160a01b031614806108ab5750836001600160a01b03166108a0846103c5565b6001600160a01b0316145b806108bb57506108bb81856107cd565b949350505050565b826001600160a01b03166108d6826105eb565b6001600160a01b03161461091b5760405162461bcd60e51b8152600401808060200182810382526029815260200180610c966029913960400191505060405180910390fd5b6001600160a01b0382166109605760405162461bcd60e51b8152600401808060200182810382526024815260200180610b8f6024913960400191505060405180910390fd5b61096981610b19565b6001600160a01b0383811660008181526002602090815260408083208054600019019055938616808352848320805460010190558583529082905283822080546bffffffffffffffffffffffff16600160601b83021790559251849392917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60006109fa84610b56565b610a06575060016108bb565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b83811015610a80578181015183820152602001610a68565b50505050905090810190601f168015610aad5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610acf57600080fd5b505af1158015610ae3573d6000803e3d6000fd5b505050506040513d6020811015610af957600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6000818152600160205260409020546001600160a01b031615610b5357600081815260016020526040902080546001600160a01b03191690555b50565b3b15159056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a723158203ff652a0c96178be1755870def59c2933c21ac84fe4bb68850f05a4212a3082e64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployWizardNFT deploys a new Ethereum contract, binding an instance of WizardNFT to it.
func DeployWizardNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WizardNFT, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardNFTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WizardNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WizardNFT{WizardNFTCaller: WizardNFTCaller{contract: contract}, WizardNFTTransactor: WizardNFTTransactor{contract: contract}, WizardNFTFilterer: WizardNFTFilterer{contract: contract}}, nil
}

// WizardNFT is an auto generated Go binding around an Ethereum contract.
type WizardNFT struct {
	WizardNFTCaller     // Read-only binding to the contract
	WizardNFTTransactor // Write-only binding to the contract
	WizardNFTFilterer   // Log filterer for contract events
}

// WizardNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardNFTSession struct {
	Contract     *WizardNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WizardNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardNFTCallerSession struct {
	Contract *WizardNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WizardNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardNFTTransactorSession struct {
	Contract     *WizardNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WizardNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardNFTRaw struct {
	Contract *WizardNFT // Generic contract binding to access the raw methods on
}

// WizardNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardNFTCallerRaw struct {
	Contract *WizardNFTCaller // Generic read-only contract binding to access the raw methods on
}

// WizardNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardNFTTransactorRaw struct {
	Contract *WizardNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardNFT creates a new instance of WizardNFT, bound to a specific deployed contract.
func NewWizardNFT(address common.Address, backend bind.ContractBackend) (*WizardNFT, error) {
	contract, err := bindWizardNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardNFT{WizardNFTCaller: WizardNFTCaller{contract: contract}, WizardNFTTransactor: WizardNFTTransactor{contract: contract}, WizardNFTFilterer: WizardNFTFilterer{contract: contract}}, nil
}

// NewWizardNFTCaller creates a new read-only instance of WizardNFT, bound to a specific deployed contract.
func NewWizardNFTCaller(address common.Address, caller bind.ContractCaller) (*WizardNFTCaller, error) {
	contract, err := bindWizardNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardNFTCaller{contract: contract}, nil
}

// NewWizardNFTTransactor creates a new write-only instance of WizardNFT, bound to a specific deployed contract.
func NewWizardNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardNFTTransactor, error) {
	contract, err := bindWizardNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardNFTTransactor{contract: contract}, nil
}

// NewWizardNFTFilterer creates a new log filterer instance of WizardNFT, bound to a specific deployed contract.
func NewWizardNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardNFTFilterer, error) {
	contract, err := bindWizardNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardNFTFilterer{contract: contract}, nil
}

// bindWizardNFT binds a generic wrapper to an already deployed contract.
func bindWizardNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardNFT *WizardNFTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardNFT.Contract.WizardNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardNFT *WizardNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardNFT.Contract.WizardNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardNFT *WizardNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardNFT.Contract.WizardNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardNFT *WizardNFTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardNFT *WizardNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardNFT *WizardNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardNFT *WizardNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WizardNFT.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardNFT *WizardNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardNFT.Contract.BalanceOf(&_WizardNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardNFT *WizardNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardNFT.Contract.BalanceOf(&_WizardNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTCaller) GetApproved(opts *bind.CallOpts, wizardId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardNFT.contract.Call(opts, out, "getApproved", wizardId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTSession) GetApproved(wizardId *big.Int) (common.Address, error) {
	return _WizardNFT.Contract.GetApproved(&_WizardNFT.CallOpts, wizardId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTCallerSession) GetApproved(wizardId *big.Int) (common.Address, error) {
	return _WizardNFT.Contract.GetApproved(&_WizardNFT.CallOpts, wizardId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardNFT *WizardNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardNFT.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardNFT *WizardNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardNFT.Contract.IsApprovedForAll(&_WizardNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardNFT *WizardNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardNFT.Contract.IsApprovedForAll(&_WizardNFT.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTCaller) OwnerOf(opts *bind.CallOpts, wizardId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardNFT.contract.Call(opts, out, "ownerOf", wizardId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTSession) OwnerOf(wizardId *big.Int) (common.Address, error) {
	return _WizardNFT.Contract.OwnerOf(&_WizardNFT.CallOpts, wizardId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 wizardId) constant returns(address)
func (_WizardNFT *WizardNFTCallerSession) OwnerOf(wizardId *big.Int) (common.Address, error) {
	return _WizardNFT.Contract.OwnerOf(&_WizardNFT.CallOpts, wizardId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardNFT *WizardNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardNFT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardNFT *WizardNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardNFT.Contract.SupportsInterface(&_WizardNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardNFT *WizardNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardNFT.Contract.SupportsInterface(&_WizardNFT.CallOpts, interfaceId)
}

// WizardsById is a free data retrieval call binding the contract method 0x5938d97a.
//
// Solidity: function wizardsById(uint256 ) constant returns(uint8 affinity, uint88 innatePower, address owner, bytes32 metadata)
func (_WizardNFT *WizardNFTCaller) WizardsById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Affinity    uint8
	InnatePower *big.Int
	Owner       common.Address
	Metadata    [32]byte
}, error) {
	ret := new(struct {
		Affinity    uint8
		InnatePower *big.Int
		Owner       common.Address
		Metadata    [32]byte
	})
	out := ret
	err := _WizardNFT.contract.Call(opts, out, "wizardsById", arg0)
	return *ret, err
}

// WizardsById is a free data retrieval call binding the contract method 0x5938d97a.
//
// Solidity: function wizardsById(uint256 ) constant returns(uint8 affinity, uint88 innatePower, address owner, bytes32 metadata)
func (_WizardNFT *WizardNFTSession) WizardsById(arg0 *big.Int) (struct {
	Affinity    uint8
	InnatePower *big.Int
	Owner       common.Address
	Metadata    [32]byte
}, error) {
	return _WizardNFT.Contract.WizardsById(&_WizardNFT.CallOpts, arg0)
}

// WizardsById is a free data retrieval call binding the contract method 0x5938d97a.
//
// Solidity: function wizardsById(uint256 ) constant returns(uint8 affinity, uint88 innatePower, address owner, bytes32 metadata)
func (_WizardNFT *WizardNFTCallerSession) WizardsById(arg0 *big.Int) (struct {
	Affinity    uint8
	InnatePower *big.Int
	Owner       common.Address
	Metadata    [32]byte
}, error) {
	return _WizardNFT.Contract.WizardsById(&_WizardNFT.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.contract.Transact(opts, "approve", to, wizardId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTSession) Approve(to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.Approve(&_WizardNFT.TransactOpts, to, wizardId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactorSession) Approve(to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.Approve(&_WizardNFT.TransactOpts, to, wizardId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.contract.Transact(opts, "safeTransferFrom", from, to, wizardId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTSession) SafeTransferFrom(from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.SafeTransferFrom(&_WizardNFT.TransactOpts, from, to, wizardId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.SafeTransferFrom(&_WizardNFT.TransactOpts, from, to, wizardId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId, bytes _data) returns()
func (_WizardNFT *WizardNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, wizardId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardNFT.contract.Transact(opts, "safeTransferFrom0", from, to, wizardId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId, bytes _data) returns()
func (_WizardNFT *WizardNFTSession) SafeTransferFrom0(from common.Address, to common.Address, wizardId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardNFT.Contract.SafeTransferFrom0(&_WizardNFT.TransactOpts, from, to, wizardId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 wizardId, bytes _data) returns()
func (_WizardNFT *WizardNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, wizardId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardNFT.Contract.SafeTransferFrom0(&_WizardNFT.TransactOpts, from, to, wizardId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardNFT *WizardNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardNFT.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardNFT *WizardNFTSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardNFT.Contract.SetApprovalForAll(&_WizardNFT.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardNFT *WizardNFTTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardNFT.Contract.SetApprovalForAll(&_WizardNFT.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.contract.Transact(opts, "transferFrom", from, to, wizardId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTSession) TransferFrom(from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.TransferFrom(&_WizardNFT.TransactOpts, from, to, wizardId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 wizardId) returns()
func (_WizardNFT *WizardNFTTransactorSession) TransferFrom(from common.Address, to common.Address, wizardId *big.Int) (*types.Transaction, error) {
	return _WizardNFT.Contract.TransferFrom(&_WizardNFT.TransactOpts, from, to, wizardId)
}

// WizardNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WizardNFT contract.
type WizardNFTApprovalIterator struct {
	Event *WizardNFTApproval // Event containing the contract specifics and raw log

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
func (it *WizardNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardNFTApproval)
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
		it.Event = new(WizardNFTApproval)
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
func (it *WizardNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardNFTApproval represents a Approval event raised by the WizardNFT contract.
type WizardNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardNFT *WizardNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WizardNFTApprovalIterator, error) {

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

	logs, sub, err := _WizardNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardNFTApprovalIterator{contract: _WizardNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardNFT *WizardNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WizardNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _WizardNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardNFTApproval)
				if err := _WizardNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WizardNFT *WizardNFTFilterer) ParseApproval(log types.Log) (*WizardNFTApproval, error) {
	event := new(WizardNFTApproval)
	if err := _WizardNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the WizardNFT contract.
type WizardNFTApprovalForAllIterator struct {
	Event *WizardNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *WizardNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardNFTApprovalForAll)
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
		it.Event = new(WizardNFTApprovalForAll)
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
func (it *WizardNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardNFTApprovalForAll represents a ApprovalForAll event raised by the WizardNFT contract.
type WizardNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardNFT *WizardNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WizardNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WizardNFTApprovalForAllIterator{contract: _WizardNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardNFT *WizardNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WizardNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardNFTApprovalForAll)
				if err := _WizardNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_WizardNFT *WizardNFTFilterer) ParseApprovalForAll(log types.Log) (*WizardNFTApprovalForAll, error) {
	event := new(WizardNFTApprovalForAll)
	if err := _WizardNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WizardNFT contract.
type WizardNFTTransferIterator struct {
	Event *WizardNFTTransfer // Event containing the contract specifics and raw log

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
func (it *WizardNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardNFTTransfer)
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
		it.Event = new(WizardNFTTransfer)
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
func (it *WizardNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardNFTTransfer represents a Transfer event raised by the WizardNFT contract.
type WizardNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardNFT *WizardNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WizardNFTTransferIterator, error) {

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

	logs, sub, err := _WizardNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardNFTTransferIterator{contract: _WizardNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardNFT *WizardNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WizardNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _WizardNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardNFTTransfer)
				if err := _WizardNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WizardNFT *WizardNFTFilterer) ParseTransfer(log types.Log) (*WizardNFTTransfer, error) {
	event := new(WizardNFTTransfer)
	if err := _WizardNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardNFTWizardAffinityAssignedIterator is returned from FilterWizardAffinityAssigned and is used to iterate over the raw logs and unpacked data for WizardAffinityAssigned events raised by the WizardNFT contract.
type WizardNFTWizardAffinityAssignedIterator struct {
	Event *WizardNFTWizardAffinityAssigned // Event containing the contract specifics and raw log

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
func (it *WizardNFTWizardAffinityAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardNFTWizardAffinityAssigned)
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
		it.Event = new(WizardNFTWizardAffinityAssigned)
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
func (it *WizardNFTWizardAffinityAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardNFTWizardAffinityAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardNFTWizardAffinityAssigned represents a WizardAffinityAssigned event raised by the WizardNFT contract.
type WizardNFTWizardAffinityAssigned struct {
	WizardId *big.Int
	Affinity uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWizardAffinityAssigned is a free log retrieval operation binding the contract event 0xef597ca22f25aec904866c3228f39dd59d3bc8345b8fd3cbea7ce568c5b1b22a.
//
// Solidity: event WizardAffinityAssigned(uint256 wizardId, uint8 affinity)
func (_WizardNFT *WizardNFTFilterer) FilterWizardAffinityAssigned(opts *bind.FilterOpts) (*WizardNFTWizardAffinityAssignedIterator, error) {

	logs, sub, err := _WizardNFT.contract.FilterLogs(opts, "WizardAffinityAssigned")
	if err != nil {
		return nil, err
	}
	return &WizardNFTWizardAffinityAssignedIterator{contract: _WizardNFT.contract, event: "WizardAffinityAssigned", logs: logs, sub: sub}, nil
}

// WatchWizardAffinityAssigned is a free log subscription operation binding the contract event 0xef597ca22f25aec904866c3228f39dd59d3bc8345b8fd3cbea7ce568c5b1b22a.
//
// Solidity: event WizardAffinityAssigned(uint256 wizardId, uint8 affinity)
func (_WizardNFT *WizardNFTFilterer) WatchWizardAffinityAssigned(opts *bind.WatchOpts, sink chan<- *WizardNFTWizardAffinityAssigned) (event.Subscription, error) {

	logs, sub, err := _WizardNFT.contract.WatchLogs(opts, "WizardAffinityAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardNFTWizardAffinityAssigned)
				if err := _WizardNFT.contract.UnpackLog(event, "WizardAffinityAssigned", log); err != nil {
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

// ParseWizardAffinityAssigned is a log parse operation binding the contract event 0xef597ca22f25aec904866c3228f39dd59d3bc8345b8fd3cbea7ce568c5b1b22a.
//
// Solidity: event WizardAffinityAssigned(uint256 wizardId, uint8 affinity)
func (_WizardNFT *WizardNFTFilterer) ParseWizardAffinityAssigned(log types.Log) (*WizardNFTWizardAffinityAssigned, error) {
	event := new(WizardNFTWizardAffinityAssigned)
	if err := _WizardNFT.contract.UnpackLog(event, "WizardAffinityAssigned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardNFTWizardConjuredIterator is returned from FilterWizardConjured and is used to iterate over the raw logs and unpacked data for WizardConjured events raised by the WizardNFT contract.
type WizardNFTWizardConjuredIterator struct {
	Event *WizardNFTWizardConjured // Event containing the contract specifics and raw log

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
func (it *WizardNFTWizardConjuredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardNFTWizardConjured)
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
		it.Event = new(WizardNFTWizardConjured)
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
func (it *WizardNFTWizardConjuredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardNFTWizardConjuredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardNFTWizardConjured represents a WizardConjured event raised by the WizardNFT contract.
type WizardNFTWizardConjured struct {
	WizardId    *big.Int
	Affinity    uint8
	InnatePower *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWizardConjured is a free log retrieval operation binding the contract event 0x3b7fcf44666972972487f1ac302bef03ee0d35027387ba8a897207466a081725.
//
// Solidity: event WizardConjured(uint256 wizardId, uint8 affinity, uint256 innatePower)
func (_WizardNFT *WizardNFTFilterer) FilterWizardConjured(opts *bind.FilterOpts) (*WizardNFTWizardConjuredIterator, error) {

	logs, sub, err := _WizardNFT.contract.FilterLogs(opts, "WizardConjured")
	if err != nil {
		return nil, err
	}
	return &WizardNFTWizardConjuredIterator{contract: _WizardNFT.contract, event: "WizardConjured", logs: logs, sub: sub}, nil
}

// WatchWizardConjured is a free log subscription operation binding the contract event 0x3b7fcf44666972972487f1ac302bef03ee0d35027387ba8a897207466a081725.
//
// Solidity: event WizardConjured(uint256 wizardId, uint8 affinity, uint256 innatePower)
func (_WizardNFT *WizardNFTFilterer) WatchWizardConjured(opts *bind.WatchOpts, sink chan<- *WizardNFTWizardConjured) (event.Subscription, error) {

	logs, sub, err := _WizardNFT.contract.WatchLogs(opts, "WizardConjured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardNFTWizardConjured)
				if err := _WizardNFT.contract.UnpackLog(event, "WizardConjured", log); err != nil {
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

// ParseWizardConjured is a log parse operation binding the contract event 0x3b7fcf44666972972487f1ac302bef03ee0d35027387ba8a897207466a081725.
//
// Solidity: event WizardConjured(uint256 wizardId, uint8 affinity, uint256 innatePower)
func (_WizardNFT *WizardNFTFilterer) ParseWizardConjured(log types.Log) (*WizardNFTWizardConjured, error) {
	event := new(WizardNFTWizardConjured)
	if err := _WizardNFT.contract.UnpackLog(event, "WizardConjured", log); err != nil {
		return nil, err
	}
	return event, nil
}
