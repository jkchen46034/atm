package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"takeoff.com/atm/model"
)

func Test_Deposit(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(43000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, balance, message := handler.Deposit(20000)
	assert.Nil(t, err)
	assert.Equal(t, balance, 50000)
	assert.Equal(t, handler.machine.GetInventory(), 63000)
	assert.Equal(t, *message, "Current balance: $50000.")
}

func Test_Deposit_NegativeAccount(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", -10000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(43000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, balance, message := handler.Deposit(30000)
	assert.Nil(t, err)
	assert.Equal(t, balance, 20000)
	assert.Equal(t, handler.machine.GetInventory(), 73000)
	assert.Equal(t, *message, "Current balance: $20000.")
}
