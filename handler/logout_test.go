package handler

import (
	"github.com/stretchr/testify/assert"
	"takeoff.com/atm/model"
	"testing"
)

func TestLogout_0(t *testing.T) {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 1024),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(1000000), accountMap)

	_ = handler.Login("2859459814", "7386")

	err, accountID := handler.Logout()

	assert.Nil(t, err)
	assert.Equal(t, accountID, "2859459814")
}

func TestLogout_No_one_login(t *testing.T) {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 1024),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(1000000), accountMap)

	err, accountID := handler.Logout()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "No account is currently authorized")
	assert.Equal(t, accountID, "")
}
