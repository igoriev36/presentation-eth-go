// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PowerABI is the input ABI used to generate the binding from.
const PowerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"payBill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"meter\",\"type\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\"}],\"name\":\"enableAccount\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"meter_id\",\"type\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"last_reading\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"newOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBalance\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"newReading\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"teller\",\"type\":\"address\"},{\"name\":\"tech\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"meter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"last\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"ErrorEvent\",\"type\":\"event\"}]"

// PowerBin is the compiled bytecode used for deploying new contracts.
const PowerBin = `0x6060604052341561000c57fe5b6040516040806105428339810160405280516020909101515b60048054600160a060020a03338116600160a060020a0319928316179092556002805485841690831617905560038054928416929091169190911790555b50505b6104cd806100756000396000f3006060604052361561005c5763ffffffff60e060020a60003504166322934487811461005e5780632433ba461461007f57806327e235e3146100a657806370a08231146100ea5780638595245414610118578063ee3e158214610136575bfe5b341561006657fe5b61007d600160a060020a0360043516602435610157565b005b341561008757fe5b61007d600160a060020a03600435811690602435166044356101ac565b005b34156100ae57fe5b6100c2600160a060020a0360043516610218565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34156100f257fe5b610106600160a060020a0360043516610243565b60408051918252519081900360200190f35b341561012057fe5b61007d600160a060020a036004351661026a565b005b341561013e57fe5b61007d600435600160a060020a03602435166102b3565b005b60025433600160a060020a039081169116146101735760006000fd5b6001548181151561018057fe5b600160a060020a038416600090815260208190526040902060010180549290910490910390555b5b5050565b60035433600160a060020a039081169116146101c85760006000fd5b600160a060020a0383811660009081526020819052604081206001810191909155805473ffffffffffffffffffffffffffffffffffffffff19169184169190911781556002018190555b5b505050565b600060208190529081526040902080546001820154600290920154600160a060020a03909116919083565b60018054600160a060020a038316600090815260208190526040902090910154025b919050565b60045433600160a060020a039081169116146102865760006000fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600160a060020a0381811660009081526020819052604090205433821691161461037b57600160a060020a03808216600081815260208181526040918290206002015482513395909516918501919091528382019290925260608301919091526080820184905260a0808352600b908301527f77726f6e67206d6574657200000000000000000000000000000000000000000060c0830152517fa12de0e8c86e6f779e48e150f79abf5639db02a729748411d2bc332d886878b19181900360e00190a16101a7565b600160a060020a0381166000908152602081905260409020600201548214156103a3576101a7565b600160a060020a0381166000908152602081905260409020600201548211156103fc57600160a060020a0381166000908152602081905260409020600281018054600190920180549285039092019091558290556101a7565b600160a060020a03808216600081815260208181526040918290206002015482513395909516918501919091528382019290925260608301919091526080820184905260a0808352600e908301527f6e6567617469766520757361676500000000000000000000000000000000000060c0830152517fa12de0e8c86e6f779e48e150f79abf5639db02a729748411d2bc332d886878b19181900360e00190a15b5b50505600a165627a7a7230582091c36bb099a2bd4c1ce30bc98f22a3bd779a267f917a6cb580f3f77053237e140029`

//start deploy OMIT

// DeployPower deploys a new Ethereum contract, binding an instance of Power to it.
func DeployPower(auth *bind.TransactOpts, backend bind.ContractBackend, teller common.Address, tech common.Address) (common.Address, *types.Transaction, *Power, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PowerBin), backend, teller, tech)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Power{PowerCaller: PowerCaller{contract: contract}, PowerTransactor: PowerTransactor{contract: contract}}, nil
}

//end deploy OMIT
// Power is an auto generated Go binding around an Ethereum contract.
type Power struct {
	PowerCaller     // Read-only binding to the contract
	PowerTransactor // Write-only binding to the contract
}

// PowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowerSession struct {
	Contract     *Power            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowerCallerSession struct {
	Contract *PowerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowerTransactorSession struct {
	Contract     *PowerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowerRaw struct {
	Contract *Power // Generic contract binding to access the raw methods on
}

// PowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowerCallerRaw struct {
	Contract *PowerCaller // Generic read-only contract binding to access the raw methods on
}

// PowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowerTransactorRaw struct {
	Contract *PowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPower creates a new instance of Power, bound to a specific deployed contract.
func NewPower(address common.Address, backend bind.ContractBackend) (*Power, error) {
	contract, err := bindPower(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Power{PowerCaller: PowerCaller{contract: contract}, PowerTransactor: PowerTransactor{contract: contract}}, nil
}

// NewPowerCaller creates a new read-only instance of Power, bound to a specific deployed contract.
func NewPowerCaller(address common.Address, caller bind.ContractCaller) (*PowerCaller, error) {
	contract, err := bindPower(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PowerCaller{contract: contract}, nil
}

// NewPowerTransactor creates a new write-only instance of Power, bound to a specific deployed contract.
func NewPowerTransactor(address common.Address, transactor bind.ContractTransactor) (*PowerTransactor, error) {
	contract, err := bindPower(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PowerTransactor{contract: contract}, nil
}

// bindPower binds a generic wrapper to an already deployed contract.
func bindPower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Power *PowerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Power.Contract.PowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Power *PowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Power.Contract.PowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Power *PowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Power.Contract.PowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Power *PowerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Power.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Power *PowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Power.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Power *PowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Power.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(user address) constant returns(uint256)
func (_Power *PowerCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Power.contract.Call(opts, out, "balanceOf", user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(user address) constant returns(uint256)
func (_Power *PowerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Power.Contract.BalanceOf(&_Power.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(user address) constant returns(uint256)
func (_Power *PowerCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _Power.Contract.BalanceOf(&_Power.CallOpts, user)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(meter_id address, balance uint256, last_reading uint256)
func (_Power *PowerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
	Meter_id     common.Address
	Balance      *big.Int
	Last_reading *big.Int
}, error) {
	ret := new(struct {
		Meter_id     common.Address
		Balance      *big.Int
		Last_reading *big.Int
	})
	out := ret
	err := _Power.contract.Call(opts, out, "balances", arg0)
	return *ret, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(meter_id address, balance uint256, last_reading uint256)
func (_Power *PowerSession) Balances(arg0 common.Address) (struct {
	Meter_id     common.Address
	Balance      *big.Int
	Last_reading *big.Int
}, error) {
	return _Power.Contract.Balances(&_Power.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(meter_id address, balance uint256, last_reading uint256)
func (_Power *PowerCallerSession) Balances(arg0 common.Address) (struct {
	Meter_id     common.Address
	Balance      *big.Int
	Last_reading *big.Int
}, error) {
	return _Power.Contract.Balances(&_Power.CallOpts, arg0)
}

// EnableAccount is a paid mutator transaction binding the contract method 0x2433ba46.
//
// Solidity: function enableAccount(user address, meter address, start uint256) returns()
func (_Power *PowerTransactor) EnableAccount(opts *bind.TransactOpts, user common.Address, meter common.Address, start *big.Int) (*types.Transaction, error) {
	return _Power.contract.Transact(opts, "enableAccount", user, meter, start)
}

// EnableAccount is a paid mutator transaction binding the contract method 0x2433ba46.
//
// Solidity: function enableAccount(user address, meter address, start uint256) returns()
func (_Power *PowerSession) EnableAccount(user common.Address, meter common.Address, start *big.Int) (*types.Transaction, error) {
	return _Power.Contract.EnableAccount(&_Power.TransactOpts, user, meter, start)
}

// EnableAccount is a paid mutator transaction binding the contract method 0x2433ba46.
//
// Solidity: function enableAccount(user address, meter address, start uint256) returns()
func (_Power *PowerTransactorSession) EnableAccount(user common.Address, meter common.Address, start *big.Int) (*types.Transaction, error) {
	return _Power.Contract.EnableAccount(&_Power.TransactOpts, user, meter, start)
}

// NewOwner is a paid mutator transaction binding the contract method 0x85952454.
//
// Solidity: function newOwner(newOwner address) returns()
func (_Power *PowerTransactor) NewOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Power.contract.Transact(opts, "newOwner", newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0x85952454.
//
// Solidity: function newOwner(newOwner address) returns()
func (_Power *PowerSession) NewOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Power.Contract.NewOwner(&_Power.TransactOpts, newOwner)
}

// NewOwner is a paid mutator transaction binding the contract method 0x85952454.
//
// Solidity: function newOwner(newOwner address) returns()
func (_Power *PowerTransactorSession) NewOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Power.Contract.NewOwner(&_Power.TransactOpts, newOwner)
}

// NewReading is a paid mutator transaction binding the contract method 0xee3e1582.
//
// Solidity: function newReading(newBalance uint256, user address) returns()
func (_Power *PowerTransactor) NewReading(opts *bind.TransactOpts, newBalance *big.Int, user common.Address) (*types.Transaction, error) {
	return _Power.contract.Transact(opts, "newReading", newBalance, user)
}

// NewReading is a paid mutator transaction binding the contract method 0xee3e1582.
//
// Solidity: function newReading(newBalance uint256, user address) returns()
func (_Power *PowerSession) NewReading(newBalance *big.Int, user common.Address) (*types.Transaction, error) {
	return _Power.Contract.NewReading(&_Power.TransactOpts, newBalance, user)
}

// NewReading is a paid mutator transaction binding the contract method 0xee3e1582.
//
// Solidity: function newReading(newBalance uint256, user address) returns()
func (_Power *PowerTransactorSession) NewReading(newBalance *big.Int, user common.Address) (*types.Transaction, error) {
	return _Power.Contract.NewReading(&_Power.TransactOpts, newBalance, user)
}

// PayBill is a paid mutator transaction binding the contract method 0x22934487.
//
// Solidity: function payBill(user address, value uint256) returns()
func (_Power *PowerTransactor) PayBill(opts *bind.TransactOpts, user common.Address, value *big.Int) (*types.Transaction, error) {
	return _Power.contract.Transact(opts, "payBill", user, value)
}

// PayBill is a paid mutator transaction binding the contract method 0x22934487.
//
// Solidity: function payBill(user address, value uint256) returns()
func (_Power *PowerSession) PayBill(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _Power.Contract.PayBill(&_Power.TransactOpts, user, value)
}

// PayBill is a paid mutator transaction binding the contract method 0x22934487.
//
// Solidity: function payBill(user address, value uint256) returns()
func (_Power *PowerTransactorSession) PayBill(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _Power.Contract.PayBill(&_Power.TransactOpts, user, value)
}
