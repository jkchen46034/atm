package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"takeoff.com/atm/model"
)

func Test_Withdraw(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(43000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, amount_dispensed, balance, message := handler.Withdraw(20000)
	assert.Nil(t, err)
	assert.Equal(t, amount_dispensed, 20000)
	assert.Equal(t, balance, 10000)
	assert.Equal(t, handler.machine.GetInventory(), 23000)
	assert.Equal(t, *message, "Amount dispensed: $20000. Current balance: $10000.")
}

func Test_Withdraw_OnlyIn20(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(40000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, amount_dispensed, balance, message := handler.Withdraw(3000)
	assert.Nil(t, err)
	assert.Equal(t, amount_dispensed, 2000)
	assert.Equal(t, balance, 8000)
	assert.Equal(t, handler.machine.GetInventory(), 38000)
	assert.Equal(t, *message, "Amount dispensed: $2000. Current balance: $8000.")
}

func Test_Withdraw_ATM_Inventory_Not_Enough(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(10000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, amount_dispensed, balance, message := handler.Withdraw(20000)
	assert.Nil(t, err)
	assert.Equal(t, amount_dispensed, 10000)
	assert.Equal(t, balance, 20000)
	assert.Equal(t, handler.machine.GetInventory(), 0)
	assert.Equal(t, *message, "Amount dispensed: $10000. Current balance: $20000.")
}

func Test_Withdraw_ATM_Inventory_is_0(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(0), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, _, _, _ = handler.Withdraw(20000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to process your withdrawal at this time.")
}

func Test_Withdraw_Account_Overdraft(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(20000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, withdraw_amount, balance, message := handler.Withdraw(12000)
	assert.Nil(t, err)
	assert.Equal(t, withdraw_amount, 12000)
	assert.Equal(t, balance, -2500)
	assert.Equal(t, *message, "Amount dispensed: $12000. You have been charged overdraft fee of $500. Current balance: $-2500")
}

func Test_Withdraw_Account_No_Fund(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", -10000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(10000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, _, _, _ = handler.Withdraw(20000)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Your account is overdrawn! You may not make withdrawals at this time.")
}
