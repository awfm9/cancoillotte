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

// WizardPresaleNFTABI is the input ABI used to generate the binding from.
const WizardPresaleNFTABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"}],\"name\":\"WizardAlignmentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"name\":\"WizardSummoned\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_wizardsById\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"affinity\",\"type\":\"uint8\"},{\"internalType\":\"uint88\",\"name\":\"power\",\"type\":\"uint88\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WizardPresaleNFTFuncSigs maps the 4-byte function signature to its string representation.
var WizardPresaleNFTFuncSigs = map[string]string{
	"63f3ce31": "_wizardsById(uint256)",
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
}

// WizardPresaleNFTBin is the compiled bytecode used for deploying new contracts.
var WizardPresaleNFTBin = "0x608060405234801561001057600080fd5b506100437f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b6100757f80ac58cd000000000000000000000000000000000000000000000000000000006001600160e01b0361007a16565b6100e6565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100a957600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610d4c806100f56000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636352211e116100715780636352211e146101bc57806363f3ce31146101d957806370a082311461022d578063a22cb46514610265578063b88d4fde14610293578063e985e9c514610359576100a9565b806301ffc9a7146100ae578063081812fc146100e9578063095ea7b31461012257806323b872dd1461015057806342842e0e14610186575b600080fd5b6100d5600480360360208110156100c457600080fd5b50356001600160e01b031916610387565b604080519115158252519081900360200190f35b610106600480360360208110156100ff57600080fd5b50356103a6565b604080516001600160a01b039092168252519081900360200190f35b61014e6004803603604081101561013857600080fd5b506001600160a01b038135169060200135610408565b005b61014e6004803603606081101561016657600080fd5b506001600160a01b03813581169160208101359091169060400135610519565b61014e6004803603606081101561019c57600080fd5b506001600160a01b0381358116916020810135909116906040013561056e565b610106600480360360208110156101d257600080fd5b5035610589565b6101f6600480360360208110156101ef57600080fd5b50356105ea565b6040805160ff90941684526affffffffffffffffffffff90921660208401526001600160a01b031682820152519081900360600190f35b6102536004803603602081101561024357600080fd5b50356001600160a01b0316610624565b60408051918252519081900360200190f35b61014e6004803603604081101561027b57600080fd5b506001600160a01b0381351690602001351515610687565b61014e600480360360808110156102a957600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156102e457600080fd5b8201836020820111156102f657600080fd5b8035906020019184600183028401116401000000008311171561031857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610753945050505050565b6100d56004803603604081101561036f57600080fd5b506001600160a01b03813581169160200135166107ab565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103b1826107d9565b6103ec5760405162461bcd60e51b815260040180806020018281038252602c815260200180610c4a602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b600061041382610589565b9050806001600160a01b0316836001600160a01b031614156104665760405162461bcd60e51b8152600401808060200182810382526021815260200180610c9f6021913960400191505060405180910390fd5b336001600160a01b0382161480610482575061048281336107ab565b6104bd5760405162461bcd60e51b8152600401808060200182810382526038815260200180610bbf6038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61052333826107fd565b61055e5760405162461bcd60e51b8152600401808060200182810382526031815260200180610cc06031913960400191505060405180910390fd5b6105698383836108a1565b505050565b61056983838360405180602001604052806000815250610753565b600081815260016020526040812054600160601b90046001600160a01b0316806105e45760405162461bcd60e51b8152600401808060200182810382526029815260200180610c216029913960400191505060405180910390fd5b92915050565b60016020526000908152604090205460ff81169061010081046affffffffffffffffffffff1690600160601b90046001600160a01b031683565b60006001600160a01b03821661066b5760405162461bcd60e51b815260040180806020018281038252602a815260200180610bf7602a913960400191505060405180910390fd5b506001600160a01b031660009081526003602052604090205490565b6001600160a01b0382163314156106e5576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61075e848484610519565b61076a848484846109cf565b6107a55760405162461bcd60e51b8152600401808060200182810382526032815260200180610b3d6032913960400191505060405180910390fd5b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160601b90046001600160a01b0316151590565b6000610808826107d9565b6108435760405162461bcd60e51b815260040180806020018281038252602c815260200180610b93602c913960400191505060405180910390fd5b600061084e83610589565b9050806001600160a01b0316846001600160a01b031614806108895750836001600160a01b031661087e846103a6565b6001600160a01b0316145b80610899575061089981856107ab565b949350505050565b826001600160a01b03166108b482610589565b6001600160a01b0316146108f95760405162461bcd60e51b8152600401808060200182810382526029815260200180610c766029913960400191505060405180910390fd5b6001600160a01b03821661093e5760405162461bcd60e51b8152600401808060200182810382526024815260200180610b6f6024913960400191505060405180910390fd5b61094781610af9565b6001600160a01b03838116600081815260036020908152604080832080546000190190559386168083528483208054600190810190915586845290915283822080546bffffffffffffffffffffffff16600160601b83021790559251849392917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60006109da84610b36565b6109e657506001610899565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b83811015610a60578181015183820152602001610a48565b50505050905090810190601f168015610a8d5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610aaf57600080fd5b505af1158015610ac3573d6000803e3d6000fd5b505050506040513d6020811015610ad957600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6000818152600260205260409020546001600160a01b031615610b3357600081815260026020526040902080546001600160a01b03191690555b50565b3b15159056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a723158204a2915fcafe5199e8afffde229a22f0c4265aca52556c263bcf34be522f8bb3c64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployWizardPresaleNFT deploys a new Ethereum contract, binding an instance of WizardPresaleNFT to it.
func DeployWizardPresaleNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WizardPresaleNFT, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardPresaleNFTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WizardPresaleNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WizardPresaleNFT{WizardPresaleNFTCaller: WizardPresaleNFTCaller{contract: contract}, WizardPresaleNFTTransactor: WizardPresaleNFTTransactor{contract: contract}, WizardPresaleNFTFilterer: WizardPresaleNFTFilterer{contract: contract}}, nil
}

// WizardPresaleNFT is an auto generated Go binding around an Ethereum contract.
type WizardPresaleNFT struct {
	WizardPresaleNFTCaller     // Read-only binding to the contract
	WizardPresaleNFTTransactor // Write-only binding to the contract
	WizardPresaleNFTFilterer   // Log filterer for contract events
}

// WizardPresaleNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type WizardPresaleNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WizardPresaleNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WizardPresaleNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WizardPresaleNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WizardPresaleNFTSession struct {
	Contract     *WizardPresaleNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WizardPresaleNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WizardPresaleNFTCallerSession struct {
	Contract *WizardPresaleNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// WizardPresaleNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WizardPresaleNFTTransactorSession struct {
	Contract     *WizardPresaleNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// WizardPresaleNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type WizardPresaleNFTRaw struct {
	Contract *WizardPresaleNFT // Generic contract binding to access the raw methods on
}

// WizardPresaleNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WizardPresaleNFTCallerRaw struct {
	Contract *WizardPresaleNFTCaller // Generic read-only contract binding to access the raw methods on
}

// WizardPresaleNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WizardPresaleNFTTransactorRaw struct {
	Contract *WizardPresaleNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWizardPresaleNFT creates a new instance of WizardPresaleNFT, bound to a specific deployed contract.
func NewWizardPresaleNFT(address common.Address, backend bind.ContractBackend) (*WizardPresaleNFT, error) {
	contract, err := bindWizardPresaleNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFT{WizardPresaleNFTCaller: WizardPresaleNFTCaller{contract: contract}, WizardPresaleNFTTransactor: WizardPresaleNFTTransactor{contract: contract}, WizardPresaleNFTFilterer: WizardPresaleNFTFilterer{contract: contract}}, nil
}

// NewWizardPresaleNFTCaller creates a new read-only instance of WizardPresaleNFT, bound to a specific deployed contract.
func NewWizardPresaleNFTCaller(address common.Address, caller bind.ContractCaller) (*WizardPresaleNFTCaller, error) {
	contract, err := bindWizardPresaleNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTCaller{contract: contract}, nil
}

// NewWizardPresaleNFTTransactor creates a new write-only instance of WizardPresaleNFT, bound to a specific deployed contract.
func NewWizardPresaleNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*WizardPresaleNFTTransactor, error) {
	contract, err := bindWizardPresaleNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTTransactor{contract: contract}, nil
}

// NewWizardPresaleNFTFilterer creates a new log filterer instance of WizardPresaleNFT, bound to a specific deployed contract.
func NewWizardPresaleNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*WizardPresaleNFTFilterer, error) {
	contract, err := bindWizardPresaleNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTFilterer{contract: contract}, nil
}

// bindWizardPresaleNFT binds a generic wrapper to an already deployed contract.
func bindWizardPresaleNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WizardPresaleNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardPresaleNFT *WizardPresaleNFTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardPresaleNFT.Contract.WizardPresaleNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardPresaleNFT *WizardPresaleNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.WizardPresaleNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardPresaleNFT *WizardPresaleNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.WizardPresaleNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WizardPresaleNFT *WizardPresaleNFTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WizardPresaleNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WizardPresaleNFT *WizardPresaleNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WizardPresaleNFT *WizardPresaleNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.contract.Transact(opts, method, params...)
}

// WizardsById is a free data retrieval call binding the contract method 0x63f3ce31.
//
// Solidity: function _wizardsById(uint256 ) constant returns(uint8 affinity, uint88 power, address owner)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) WizardsById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Affinity uint8
	Power    *big.Int
	Owner    common.Address
}, error) {
	ret := new(struct {
		Affinity uint8
		Power    *big.Int
		Owner    common.Address
	})
	out := ret
	err := _WizardPresaleNFT.contract.Call(opts, out, "_wizardsById", arg0)
	return *ret, err
}

// WizardsById is a free data retrieval call binding the contract method 0x63f3ce31.
//
// Solidity: function _wizardsById(uint256 ) constant returns(uint8 affinity, uint88 power, address owner)
func (_WizardPresaleNFT *WizardPresaleNFTSession) WizardsById(arg0 *big.Int) (struct {
	Affinity uint8
	Power    *big.Int
	Owner    common.Address
}, error) {
	return _WizardPresaleNFT.Contract.WizardsById(&_WizardPresaleNFT.CallOpts, arg0)
}

// WizardsById is a free data retrieval call binding the contract method 0x63f3ce31.
//
// Solidity: function _wizardsById(uint256 ) constant returns(uint8 affinity, uint88 power, address owner)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) WizardsById(arg0 *big.Int) (struct {
	Affinity uint8
	Power    *big.Int
	Owner    common.Address
}, error) {
	return _WizardPresaleNFT.Contract.WizardsById(&_WizardPresaleNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WizardPresaleNFT.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardPresaleNFT *WizardPresaleNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardPresaleNFT.Contract.BalanceOf(&_WizardPresaleNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WizardPresaleNFT.Contract.BalanceOf(&_WizardPresaleNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardPresaleNFT.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WizardPresaleNFT.Contract.GetApproved(&_WizardPresaleNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WizardPresaleNFT.Contract.GetApproved(&_WizardPresaleNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardPresaleNFT.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardPresaleNFT.Contract.IsApprovedForAll(&_WizardPresaleNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WizardPresaleNFT.Contract.IsApprovedForAll(&_WizardPresaleNFT.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WizardPresaleNFT.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WizardPresaleNFT.Contract.OwnerOf(&_WizardPresaleNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WizardPresaleNFT.Contract.OwnerOf(&_WizardPresaleNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WizardPresaleNFT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardPresaleNFT.Contract.SupportsInterface(&_WizardPresaleNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_WizardPresaleNFT *WizardPresaleNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WizardPresaleNFT.Contract.SupportsInterface(&_WizardPresaleNFT.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.Approve(&_WizardPresaleNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.Approve(&_WizardPresaleNFT.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SafeTransferFrom(&_WizardPresaleNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SafeTransferFrom(&_WizardPresaleNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardPresaleNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_WizardPresaleNFT *WizardPresaleNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SafeTransferFrom0(&_WizardPresaleNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SafeTransferFrom0(&_WizardPresaleNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardPresaleNFT.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardPresaleNFT *WizardPresaleNFTSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SetApprovalForAll(&_WizardPresaleNFT.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.SetApprovalForAll(&_WizardPresaleNFT.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.TransferFrom(&_WizardPresaleNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WizardPresaleNFT *WizardPresaleNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WizardPresaleNFT.Contract.TransferFrom(&_WizardPresaleNFT.TransactOpts, from, to, tokenId)
}

// WizardPresaleNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WizardPresaleNFT contract.
type WizardPresaleNFTApprovalIterator struct {
	Event *WizardPresaleNFTApproval // Event containing the contract specifics and raw log

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
func (it *WizardPresaleNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardPresaleNFTApproval)
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
		it.Event = new(WizardPresaleNFTApproval)
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
func (it *WizardPresaleNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardPresaleNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardPresaleNFTApproval represents a Approval event raised by the WizardPresaleNFT contract.
type WizardPresaleNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WizardPresaleNFTApprovalIterator, error) {

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

	logs, sub, err := _WizardPresaleNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTApprovalIterator{contract: _WizardPresaleNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WizardPresaleNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _WizardPresaleNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardPresaleNFTApproval)
				if err := _WizardPresaleNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) ParseApproval(log types.Log) (*WizardPresaleNFTApproval, error) {
	event := new(WizardPresaleNFTApproval)
	if err := _WizardPresaleNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardPresaleNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the WizardPresaleNFT contract.
type WizardPresaleNFTApprovalForAllIterator struct {
	Event *WizardPresaleNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *WizardPresaleNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardPresaleNFTApprovalForAll)
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
		it.Event = new(WizardPresaleNFTApprovalForAll)
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
func (it *WizardPresaleNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardPresaleNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardPresaleNFTApprovalForAll represents a ApprovalForAll event raised by the WizardPresaleNFT contract.
type WizardPresaleNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WizardPresaleNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTApprovalForAllIterator{contract: _WizardPresaleNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WizardPresaleNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardPresaleNFTApprovalForAll)
				if err := _WizardPresaleNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) ParseApprovalForAll(log types.Log) (*WizardPresaleNFTApprovalForAll, error) {
	event := new(WizardPresaleNFTApprovalForAll)
	if err := _WizardPresaleNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardPresaleNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WizardPresaleNFT contract.
type WizardPresaleNFTTransferIterator struct {
	Event *WizardPresaleNFTTransfer // Event containing the contract specifics and raw log

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
func (it *WizardPresaleNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardPresaleNFTTransfer)
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
		it.Event = new(WizardPresaleNFTTransfer)
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
func (it *WizardPresaleNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardPresaleNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardPresaleNFTTransfer represents a Transfer event raised by the WizardPresaleNFT contract.
type WizardPresaleNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WizardPresaleNFTTransferIterator, error) {

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

	logs, sub, err := _WizardPresaleNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTTransferIterator{contract: _WizardPresaleNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WizardPresaleNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _WizardPresaleNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardPresaleNFTTransfer)
				if err := _WizardPresaleNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) ParseTransfer(log types.Log) (*WizardPresaleNFTTransfer, error) {
	event := new(WizardPresaleNFTTransfer)
	if err := _WizardPresaleNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardPresaleNFTWizardAlignmentAssignedIterator is returned from FilterWizardAlignmentAssigned and is used to iterate over the raw logs and unpacked data for WizardAlignmentAssigned events raised by the WizardPresaleNFT contract.
type WizardPresaleNFTWizardAlignmentAssignedIterator struct {
	Event *WizardPresaleNFTWizardAlignmentAssigned // Event containing the contract specifics and raw log

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
func (it *WizardPresaleNFTWizardAlignmentAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardPresaleNFTWizardAlignmentAssigned)
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
		it.Event = new(WizardPresaleNFTWizardAlignmentAssigned)
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
func (it *WizardPresaleNFTWizardAlignmentAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardPresaleNFTWizardAlignmentAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardPresaleNFTWizardAlignmentAssigned represents a WizardAlignmentAssigned event raised by the WizardPresaleNFT contract.
type WizardPresaleNFTWizardAlignmentAssigned struct {
	TokenId *big.Int
	Element uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWizardAlignmentAssigned is a free log retrieval operation binding the contract event 0x0f0bc6913c78f45e890c2265a362fd955a202a954142552907cb83e4ef9f409a.
//
// Solidity: event WizardAlignmentAssigned(uint256 indexed tokenId, uint8 element)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) FilterWizardAlignmentAssigned(opts *bind.FilterOpts, tokenId []*big.Int) (*WizardPresaleNFTWizardAlignmentAssignedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.FilterLogs(opts, "WizardAlignmentAssigned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTWizardAlignmentAssignedIterator{contract: _WizardPresaleNFT.contract, event: "WizardAlignmentAssigned", logs: logs, sub: sub}, nil
}

// WatchWizardAlignmentAssigned is a free log subscription operation binding the contract event 0x0f0bc6913c78f45e890c2265a362fd955a202a954142552907cb83e4ef9f409a.
//
// Solidity: event WizardAlignmentAssigned(uint256 indexed tokenId, uint8 element)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) WatchWizardAlignmentAssigned(opts *bind.WatchOpts, sink chan<- *WizardPresaleNFTWizardAlignmentAssigned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.WatchLogs(opts, "WizardAlignmentAssigned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardPresaleNFTWizardAlignmentAssigned)
				if err := _WizardPresaleNFT.contract.UnpackLog(event, "WizardAlignmentAssigned", log); err != nil {
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

// ParseWizardAlignmentAssigned is a log parse operation binding the contract event 0x0f0bc6913c78f45e890c2265a362fd955a202a954142552907cb83e4ef9f409a.
//
// Solidity: event WizardAlignmentAssigned(uint256 indexed tokenId, uint8 element)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) ParseWizardAlignmentAssigned(log types.Log) (*WizardPresaleNFTWizardAlignmentAssigned, error) {
	event := new(WizardPresaleNFTWizardAlignmentAssigned)
	if err := _WizardPresaleNFT.contract.UnpackLog(event, "WizardAlignmentAssigned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WizardPresaleNFTWizardSummonedIterator is returned from FilterWizardSummoned and is used to iterate over the raw logs and unpacked data for WizardSummoned events raised by the WizardPresaleNFT contract.
type WizardPresaleNFTWizardSummonedIterator struct {
	Event *WizardPresaleNFTWizardSummoned // Event containing the contract specifics and raw log

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
func (it *WizardPresaleNFTWizardSummonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WizardPresaleNFTWizardSummoned)
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
		it.Event = new(WizardPresaleNFTWizardSummoned)
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
func (it *WizardPresaleNFTWizardSummonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WizardPresaleNFTWizardSummonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WizardPresaleNFTWizardSummoned represents a WizardSummoned event raised by the WizardPresaleNFT contract.
type WizardPresaleNFTWizardSummoned struct {
	TokenId *big.Int
	Element uint8
	Power   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWizardSummoned is a free log retrieval operation binding the contract event 0x7c0fbd69c04ea8ef6f62724eebd9c311d984e86457a801d81c0cb52ec9039170.
//
// Solidity: event WizardSummoned(uint256 indexed tokenId, uint8 element, uint256 power)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) FilterWizardSummoned(opts *bind.FilterOpts, tokenId []*big.Int) (*WizardPresaleNFTWizardSummonedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.FilterLogs(opts, "WizardSummoned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WizardPresaleNFTWizardSummonedIterator{contract: _WizardPresaleNFT.contract, event: "WizardSummoned", logs: logs, sub: sub}, nil
}

// WatchWizardSummoned is a free log subscription operation binding the contract event 0x7c0fbd69c04ea8ef6f62724eebd9c311d984e86457a801d81c0cb52ec9039170.
//
// Solidity: event WizardSummoned(uint256 indexed tokenId, uint8 element, uint256 power)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) WatchWizardSummoned(opts *bind.WatchOpts, sink chan<- *WizardPresaleNFTWizardSummoned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WizardPresaleNFT.contract.WatchLogs(opts, "WizardSummoned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WizardPresaleNFTWizardSummoned)
				if err := _WizardPresaleNFT.contract.UnpackLog(event, "WizardSummoned", log); err != nil {
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

// ParseWizardSummoned is a log parse operation binding the contract event 0x7c0fbd69c04ea8ef6f62724eebd9c311d984e86457a801d81c0cb52ec9039170.
//
// Solidity: event WizardSummoned(uint256 indexed tokenId, uint8 element, uint256 power)
func (_WizardPresaleNFT *WizardPresaleNFTFilterer) ParseWizardSummoned(log types.Log) (*WizardPresaleNFTWizardSummoned, error) {
	event := new(WizardPresaleNFTWizardSummoned)
	if err := _WizardPresaleNFT.contract.UnpackLog(event, "WizardSummoned", log); err != nil {
		return nil, err
	}
	return event, nil
}
