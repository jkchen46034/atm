package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"takeoff.com/atm/handler"
	"takeoff.com/atm/model"
)

func Test_Integration_Test1(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := handler.NewHandler(model.NewMachine(1000000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, balance := handler.GetBalance()
	assert.Nil(t, err)
	assert.Equal(t, balance, 10024)

	err, balance = handler.Deposit(3000)
	assert.Nil(t, err)
	assert.Equal(t, balance, 13024)

	err, amount_dispensed, balance := handler.Withdraw(3000)
	assert.Nil(t, err)
	assert.Equal(t, amount_dispensed, 2000)
	assert.Equal(t, balance, 11024)

	err, amount_dispensed, balance = handler.Withdraw(2000)
	assert.Equal(t, amount_dispensed, 2000)
	assert.Equal(t, balance, 9024)

	err, balance = handler.Deposit(10000)
	assert.Equal(t, err, nil)
	assert.Equal(t, balance, 19024)

	err, txHistory := handler.GetTxHistory()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(txHistory), 4)
	assert.Equal(t, txHistory[3].Amount, 3000)
	assert.Equal(t, txHistory[3].BalanceAfter, 13024)
	assert.Equal(t, txHistory[2].Amount, -2000)
	assert.Equal(t, txHistory[2].BalanceAfter, 11024)
	assert.Equal(t, txHistory[1].Amount, -2000)
	assert.Equal(t, txHistory[1].BalanceAfter, 9024)
	assert.Equal(t, txHistory[0].Amount, 10000)
	assert.Equal(t, txHistory[0].BalanceAfter, 19024)

	err, accountID := handler.Logout()
	assert.Equal(t, err, nil)
	assert.Equal(t, accountID, "2859459814")

	err, amount_dispensed, balance = handler.Withdraw(3000)
	assert.NotNil(t, err)

	err, balance = handler.Deposit(3000)
	assert.NotNil(t, err)

	err, balance = handler.GetBalance()
	assert.NotNil(t, err)

	err, txHistory = handler.GetTxHistory()
	assert.NotNil(t, err)
}
