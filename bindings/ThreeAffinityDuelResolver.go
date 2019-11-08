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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ThreeAffinityDuelResolverABI is the input ABI used to generate the binding from.
const ThreeAffinityDuelResolverABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MOVE_DELTA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MOVE_MASK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WEIGHT_SUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isValidAffinity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moveSet\",\"type\":\"bytes32\"}],\"name\":\"isValidMoveSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moveSet1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"moveSet2\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"power1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affinity1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"affinity2\",\"type\":\"uint256\"}],\"name\":\"resolveDuel\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"power\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ThreeAffinityDuelResolverFuncSigs maps the 4-byte function signature to its string representation.
var ThreeAffinityDuelResolverFuncSigs = map[string]string{
	"cdcf1e1c": "MOVE_DELTA()",
	"a8007fb5": "MOVE_MASK()",
	"f92d1960": "WEIGHT_SUM()",
	"e9563dee": "isValidAffinity(uint256)",
	"1823fbbc": "isValidMoveSet(bytes32)",
	"b089894c": "resolveDuel(bytes32,bytes32,uint256,uint256,uint256,uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ThreeAffinityDuelResolverBin is the compiled bytecode used for deploying new contracts.
var ThreeAffinityDuelResolverBin = "0x608060405234801561001057600080fd5b5061060c806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b089894c1161005b578063b089894c146100f4578063cdcf1e1c1461012f578063e9563dee14610137578063f92d1960146101545761007d565b806301ffc9a7146100825780631823fbbc146100bd578063a8007fb5146100da575b600080fd5b6100a96004803603602081101561009857600080fd5b50356001600160e01b03191661015c565b604080519115158252519081900360200190f35b6100a9600480360360208110156100d357600080fd5b5035610195565b6100e26101d6565b60408051918252519081900360200190f35b6100e2600480360360c081101561010a57600080fd5b5080359060208101359060408101359060608101359060808101359060a001356101e2565b6100e2610267565b6100a96004803603602081101561014d57600080fd5b5035610273565b6100e2610279565b60006001600160e01b031982166301ffc9a760e01b148061018d57506001600160e01b031982166320fe278f60e11b145b90505b919050565b647efefefeff60d91b01600064030303030360d81b821682146101ba57506000610190565b600182901b8216156101ce57506000610190565b506001919050565b64030303030360d81b81565b60006101ed87610195565b80156101fd57506101fd86610195565b610240576040805162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a59081b5bdd995cd95d608a1b604482015290519081900360640190fd5b600061024e8888868661027f565b905061025b81878761037b565b98975050505050505050565b64010101010160d91b81565b50600190565b6101a881565b6000610289610592565b506040805160a081018252604e8152604f60208201526051918101919091526056606082015260646080820152647efefefeff60d91b95860195949094019360011993840193929092019160005b600581101561036d5760008782602081106102ee57fe5b1a905060008783602081106102ff57fe5b1a90508082038061031257505050610365565b808102600414156103245760011d6000035b60640282881415610339576064608282020590505b8682141561034b576064608282020590505b84846005811061035757fe5b602002015102949094019350505b6001016102d7565b505060649005949350505050565b60008084121561039d5761039384600003838561037b565b60000390506104d6565b600160f51b831080156103b35750600160f51b82105b6103fa576040805162461bcd60e51b8152602060048201526013602482015272496e76616c696420706f7765722076616c756560681b604482015290519081900360640190fd5b60008311801561040a5750600082115b610451576040805162461bcd60e51b8152602060048201526013602482015272496e76616c696420706f7765722076616c756560681b604482015290519081900360640190fd5b8361045e575060006104d6565b60006101a8610400860204905061040081111561047a57506104005b6000610488826102006104dd565b9050846007028411156104a0578460070293506104b2565b836007028511156104b2578360070294505b60006104cb82878761040002816104c557fe5b046104dd565b8502600a1c93505050505b9392505050565b60006008820160041c6104f48482604060056104fc565b949350505050565b60006104008184868161050b57fe5b0490505b6016811061052c576016870a9190910260dc1c906015190161050f565b80156105405780870a91909102600a82021c905b600085878161054b57fe5b069050841580610559575080155b15610569578293505050506104f4565b61057c88610400038783600189036104fc565b6104000392909202600a1c979650505050505050565b6040518060a00160405280600590602082028038833950919291505056fea265627a7a72315820c18c9ea67e0fe5d5ce139c642d62cfa8097fbfd8462b489f01eb9dd8e70d153d64736f6c637829302e352e31332d646576656c6f702e323031392e31302e31392b636f6d6d69742e64356232663334370059"

// DeployThreeAffinityDuelResolver deploys a new Ethereum contract, binding an instance of ThreeAffinityDuelResolver to it.
func DeployThreeAffinityDuelResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ThreeAffinityDuelResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeAffinityDuelResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ThreeAffinityDuelResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ThreeAffinityDuelResolver{ThreeAffinityDuelResolverCaller: ThreeAffinityDuelResolverCaller{contract: contract}, ThreeAffinityDuelResolverTransactor: ThreeAffinityDuelResolverTransactor{contract: contract}, ThreeAffinityDuelResolverFilterer: ThreeAffinityDuelResolverFilterer{contract: contract}}, nil
}

// ThreeAffinityDuelResolver is an auto generated Go binding around an Ethereum contract.
type ThreeAffinityDuelResolver struct {
	ThreeAffinityDuelResolverCaller     // Read-only binding to the contract
	ThreeAffinityDuelResolverTransactor // Write-only binding to the contract
	ThreeAffinityDuelResolverFilterer   // Log filterer for contract events
}

// ThreeAffinityDuelResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThreeAffinityDuelResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeAffinityDuelResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThreeAffinityDuelResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeAffinityDuelResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThreeAffinityDuelResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThreeAffinityDuelResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThreeAffinityDuelResolverSession struct {
	Contract     *ThreeAffinityDuelResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ThreeAffinityDuelResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThreeAffinityDuelResolverCallerSession struct {
	Contract *ThreeAffinityDuelResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ThreeAffinityDuelResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThreeAffinityDuelResolverTransactorSession struct {
	Contract     *ThreeAffinityDuelResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ThreeAffinityDuelResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThreeAffinityDuelResolverRaw struct {
	Contract *ThreeAffinityDuelResolver // Generic contract binding to access the raw methods on
}

// ThreeAffinityDuelResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThreeAffinityDuelResolverCallerRaw struct {
	Contract *ThreeAffinityDuelResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ThreeAffinityDuelResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThreeAffinityDuelResolverTransactorRaw struct {
	Contract *ThreeAffinityDuelResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThreeAffinityDuelResolver creates a new instance of ThreeAffinityDuelResolver, bound to a specific deployed contract.
func NewThreeAffinityDuelResolver(address common.Address, backend bind.ContractBackend) (*ThreeAffinityDuelResolver, error) {
	contract, err := bindThreeAffinityDuelResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThreeAffinityDuelResolver{ThreeAffinityDuelResolverCaller: ThreeAffinityDuelResolverCaller{contract: contract}, ThreeAffinityDuelResolverTransactor: ThreeAffinityDuelResolverTransactor{contract: contract}, ThreeAffinityDuelResolverFilterer: ThreeAffinityDuelResolverFilterer{contract: contract}}, nil
}

// NewThreeAffinityDuelResolverCaller creates a new read-only instance of ThreeAffinityDuelResolver, bound to a specific deployed contract.
func NewThreeAffinityDuelResolverCaller(address common.Address, caller bind.ContractCaller) (*ThreeAffinityDuelResolverCaller, error) {
	contract, err := bindThreeAffinityDuelResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeAffinityDuelResolverCaller{contract: contract}, nil
}

// NewThreeAffinityDuelResolverTransactor creates a new write-only instance of ThreeAffinityDuelResolver, bound to a specific deployed contract.
func NewThreeAffinityDuelResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ThreeAffinityDuelResolverTransactor, error) {
	contract, err := bindThreeAffinityDuelResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThreeAffinityDuelResolverTransactor{contract: contract}, nil
}

// NewThreeAffinityDuelResolverFilterer creates a new log filterer instance of ThreeAffinityDuelResolver, bound to a specific deployed contract.
func NewThreeAffinityDuelResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ThreeAffinityDuelResolverFilterer, error) {
	contract, err := bindThreeAffinityDuelResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThreeAffinityDuelResolverFilterer{contract: contract}, nil
}

// bindThreeAffinityDuelResolver binds a generic wrapper to an already deployed contract.
func bindThreeAffinityDuelResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ThreeAffinityDuelResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ThreeAffinityDuelResolver.Contract.ThreeAffinityDuelResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeAffinityDuelResolver.Contract.ThreeAffinityDuelResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeAffinityDuelResolver.Contract.ThreeAffinityDuelResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ThreeAffinityDuelResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThreeAffinityDuelResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThreeAffinityDuelResolver.Contract.contract.Transact(opts, method, params...)
}

// MOVEDELTA is a free data retrieval call binding the contract method 0xcdcf1e1c.
//
// Solidity: function MOVE_DELTA() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) MOVEDELTA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "MOVE_DELTA")
	return *ret0, err
}

// MOVEDELTA is a free data retrieval call binding the contract method 0xcdcf1e1c.
//
// Solidity: function MOVE_DELTA() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) MOVEDELTA() (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.MOVEDELTA(&_ThreeAffinityDuelResolver.CallOpts)
}

// MOVEDELTA is a free data retrieval call binding the contract method 0xcdcf1e1c.
//
// Solidity: function MOVE_DELTA() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) MOVEDELTA() (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.MOVEDELTA(&_ThreeAffinityDuelResolver.CallOpts)
}

// MOVEMASK is a free data retrieval call binding the contract method 0xa8007fb5.
//
// Solidity: function MOVE_MASK() constant returns(bytes32)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) MOVEMASK(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "MOVE_MASK")
	return *ret0, err
}

// MOVEMASK is a free data retrieval call binding the contract method 0xa8007fb5.
//
// Solidity: function MOVE_MASK() constant returns(bytes32)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) MOVEMASK() ([32]byte, error) {
	return _ThreeAffinityDuelResolver.Contract.MOVEMASK(&_ThreeAffinityDuelResolver.CallOpts)
}

// MOVEMASK is a free data retrieval call binding the contract method 0xa8007fb5.
//
// Solidity: function MOVE_MASK() constant returns(bytes32)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) MOVEMASK() ([32]byte, error) {
	return _ThreeAffinityDuelResolver.Contract.MOVEMASK(&_ThreeAffinityDuelResolver.CallOpts)
}

// WEIGHTSUM is a free data retrieval call binding the contract method 0xf92d1960.
//
// Solidity: function WEIGHT_SUM() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) WEIGHTSUM(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "WEIGHT_SUM")
	return *ret0, err
}

// WEIGHTSUM is a free data retrieval call binding the contract method 0xf92d1960.
//
// Solidity: function WEIGHT_SUM() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) WEIGHTSUM() (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.WEIGHTSUM(&_ThreeAffinityDuelResolver.CallOpts)
}

// WEIGHTSUM is a free data retrieval call binding the contract method 0xf92d1960.
//
// Solidity: function WEIGHT_SUM() constant returns(uint256)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) WEIGHTSUM() (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.WEIGHTSUM(&_ThreeAffinityDuelResolver.CallOpts)
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 ) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) IsValidAffinity(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "isValidAffinity", arg0)
	return *ret0, err
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 ) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) IsValidAffinity(arg0 *big.Int) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.IsValidAffinity(&_ThreeAffinityDuelResolver.CallOpts, arg0)
}

// IsValidAffinity is a free data retrieval call binding the contract method 0xe9563dee.
//
// Solidity: function isValidAffinity(uint256 ) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) IsValidAffinity(arg0 *big.Int) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.IsValidAffinity(&_ThreeAffinityDuelResolver.CallOpts, arg0)
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) IsValidMoveSet(opts *bind.CallOpts, moveSet [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "isValidMoveSet", moveSet)
	return *ret0, err
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) IsValidMoveSet(moveSet [32]byte) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.IsValidMoveSet(&_ThreeAffinityDuelResolver.CallOpts, moveSet)
}

// IsValidMoveSet is a free data retrieval call binding the contract method 0x1823fbbc.
//
// Solidity: function isValidMoveSet(bytes32 moveSet) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) IsValidMoveSet(moveSet [32]byte) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.IsValidMoveSet(&_ThreeAffinityDuelResolver.CallOpts, moveSet)
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256 power)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) ResolveDuel(opts *bind.CallOpts, moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "resolveDuel", moveSet1, moveSet2, power1, power2, affinity1, affinity2)
	return *ret0, err
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256 power)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) ResolveDuel(moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.ResolveDuel(&_ThreeAffinityDuelResolver.CallOpts, moveSet1, moveSet2, power1, power2, affinity1, affinity2)
}

// ResolveDuel is a free data retrieval call binding the contract method 0xb089894c.
//
// Solidity: function resolveDuel(bytes32 moveSet1, bytes32 moveSet2, uint256 power1, uint256 power2, uint256 affinity1, uint256 affinity2) constant returns(int256 power)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) ResolveDuel(moveSet1 [32]byte, moveSet2 [32]byte, power1 *big.Int, power2 *big.Int, affinity1 *big.Int, affinity2 *big.Int) (*big.Int, error) {
	return _ThreeAffinityDuelResolver.Contract.ResolveDuel(&_ThreeAffinityDuelResolver.CallOpts, moveSet1, moveSet2, power1, power2, affinity1, affinity2)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ThreeAffinityDuelResolver.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.SupportsInterface(&_ThreeAffinityDuelResolver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ThreeAffinityDuelResolver *ThreeAffinityDuelResolverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ThreeAffinityDuelResolver.Contract.SupportsInterface(&_ThreeAffinityDuelResolver.CallOpts, interfaceId)
}
