// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

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

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	Issuer    string
	Worker    string
	Desc      string
	Bonus     *big.Int
	Status    uint8
	Timestamp *big.Int
	Comment   string
}

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_owner\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qryAllTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"internalType\":\"structTaskInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"qryOneTask\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, _owner string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _owner) constant returns(uint256 balance)
func (_Task *TaskCallerSession) BalanceOf(_owner string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, _owner)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "login", username, passwd)
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) constant returns(bool)
func (_Task *TaskCallerSession) Login(username string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, username, passwd)
}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() constant returns([]TaskInfo)
func (_Task *TaskCaller) QryAllTasks(opts *bind.CallOpts) ([]TaskInfo, error) {
	var (
		ret0 = new([]TaskInfo)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "qryAllTasks")
	return *ret0, err
}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() constant returns([]TaskInfo)
func (_Task *TaskSession) QryAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.QryAllTasks(&_Task.CallOpts)
}

// QryAllTasks is a free data retrieval call binding the contract method 0x8a3f7bb3.
//
// Solidity: function qryAllTasks() constant returns([]TaskInfo)
func (_Task *TaskCallerSession) QryAllTasks() ([]TaskInfo, error) {
	return _Task.Contract.QryAllTasks(&_Task.CallOpts)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskCaller) QryOneTask(opts *bind.CallOpts, taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Task.contract.Call(opts, out, "qryOneTask", taskID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// QryOneTask is a free data retrieval call binding the contract method 0x9b3add75.
//
// Solidity: function qryOneTask(uint256 taskID) constant returns(string, string, uint256, string, uint8)
func (_Task *TaskCallerSession) QryOneTask(taskID *big.Int) (string, string, *big.Int, string, uint8, error) {
	return _Task.Contract.QryOneTask(&_Task.CallOpts, taskID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Task *TaskCallerSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "commit", worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Commit(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, worker, passwd, taskID)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "confirm", issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string passwd, uint256 taskID, string comment, uint8 status) returns()
func (_Task *TaskTransactorSession) Confirm(issuer string, passwd string, taskID *big.Int, comment string, status uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, passwd, taskID, comment, status)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "issue", issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string passwd, string desc, uint256 bonus) returns()
func (_Task *TaskTransactorSession) Issue(issuer string, passwd string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, passwd, desc, bonus)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _to, uint256 _value) returns()
func (_Task *TaskTransactorSession) Mint(_to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, _to, _value)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Task *TaskTransactorSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, username, passwd)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "take", worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string worker, string passwd, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Take(worker string, passwd string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, worker, passwd, taskID)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactor) Transfer(opts *bind.TransactOpts, _from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transfer", _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string _from, string _to, uint256 _value) returns(bool success)
func (_Task *TaskTransactorSession) Transfer(_from string, _to string, _value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, _from, _to, _value)
}
