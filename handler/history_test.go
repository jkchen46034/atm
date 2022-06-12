package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"takeoff.com/atm/model"
)

func Test_GetHistory_0(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(43000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	err, records, message := handler.GetTxHistory()
	assert.Nil(t, err)
	assert.Equal(t, len(records), 0)
	assert.Equal(t, *message, "No history found.")
}

func Test_GetHistory_1D(t *testing.T) {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 30000),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(43000), accountMap)

	err := handler.Login("2859459814", "7386")
	assert.Nil(t, err)

	handler.Deposit(20000)
	err, records, message := handler.GetTxHistory()
	assert.Nil(t, err)
	assert.Equal(t, len(records), 1)
	assert.NotEqual(t, len(*message), 0)
}
