package model

import (
	"time"
)

type TxRecord struct {
	Amount       int
	BalanceAfter int
	TxTime       time.Time
}

type Account struct {
	accountID string
	pin       string
	balance   int
	txRecords []TxRecord
}

const OVERDRAFT_FEE = 500

func NewAccount(accountID string, pin string, initialBalance int) *Account {
	return &Account{accountID, pin, initialBalance, []TxRecord{}}
}

func (a *Account) GetAccountID() string {
	return a.accountID
}

func (a *Account) GetPIN() string {
	return a.pin
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) GetHistory() []TxRecord {
	return a.txRecords
}

func (a *Account) Withdraw(amount int) (balance int, overdraft_fee int) {
	if a.balance-amount < 0 {
		overdraft_fee = OVERDRAFT_FEE
	}
	a.balance = a.balance - amount - overdraft_fee
	a.AddTx(-amount, a.balance, time.Now())
	return a.balance, overdraft_fee
}

func (a *Account) Deposit(amount int) int {
	a.balance = a.balance + amount
	a.AddTx(amount, a.balance, time.Now())
	return a.balance
}

func (a *Account) AddTx(amount int, balance_after int, t time.Time) {
	a.txRecords = append([]TxRecord{TxRecord{amount, balance_after, t}}, a.txRecords...)
}
