package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_Create(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	assert.NotNil(t, account)
}

func TestAccount_GetAccountID(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	assert.Equal(t, account.GetAccountID(), "2859459814")
}

func TestAccount_GetAccountPIN(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	assert.Equal(t, account.GetPIN(), "7386")
}

func TestAccount_GetAccountBalance(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	assert.Equal(t, account.GetBalance(), 10024)
}

func TestAccount_Withdraw(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	balance := account.Withdraw(4000)
	assert.Equal(t, balance, 6024)
	assert.Equal(t, account.GetBalance(), 6024)
}

func TestAccount_Deposit(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	balance := account.Deposit(4000)
	assert.Equal(t, balance, 14024)
	assert.Equal(t, account.GetBalance(), 14024)
}

func TestAccount_GetTxHistory_0_Transaction(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	assert.Equal(t, account.GetTxHistory(), []TxRecord{})
}

func TestAccount_GetTxHistory_1_Transaction(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	_ = account.Withdraw(4000)
	history := account.GetTxHistory()
	assert.Equal(t, len(history), 1)
	assert.Equal(t, history[0].Amount, -4000)
	assert.Equal(t, history[0].BalanceAfter, 6024)
}

func TestAccount_GetTxHistory_Another_1_Transaction(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	_ = account.Deposit(3000)
	history := account.GetTxHistory()
	assert.Equal(t, len(history), 1)
	assert.Equal(t, history[0].Amount, 3000)
	assert.Equal(t, history[0].BalanceAfter, 13024)
}

func TestAccount_GetTxHistory_2_Transactions(t *testing.T) {
	account := NewAccount("2859459814", "7386", 10024)
	_ = account.Deposit(4000)
	_ = account.Withdraw(3000)
	history := account.GetTxHistory()
	assert.Equal(t, len(history), 2)
	assert.Equal(t, history[0].Amount, -3000)
	assert.Equal(t, history[0].BalanceAfter, 11024)
	assert.Equal(t, history[1].Amount, 4000)
	assert.Equal(t, history[1].BalanceAfter, 14024)
}
