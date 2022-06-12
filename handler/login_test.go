package handler

import (
	"github.com/stretchr/testify/assert"
	"takeoff.com/atm/model"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 1024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(1000000), accountMap)

	var tests = []struct {
		accountID string
		pid       string
	}{
		{"2859459814", "7386"},
		{"1434597300", "4557"},
		{"7089382418", "0075"},
		{"2001377812", "5950"},
	}

	assert := assert.New(t)

	for _, test := range tests {
		assert.Nil(handler.Login(test.accountID, test.pid))
	}
}

func TestLogin_BadPIN(t *testing.T) {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 1024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(1000000), accountMap)

	var tests = []struct {
		accountID string
		pid       string
	}{
		{"2859459814", "77386"},
		{"1434597300", "5557"},
		{"7089382418", "9075"},
		{"2001377812", "4950"},
	}

	assert := assert.New(t)

	for _, test := range tests {
		assert.NotNil(handler.Login(test.accountID, test.pid))
	}
}

func TestLogin_BadAccountID(t *testing.T) {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 1024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	accountMap := model.NewAccountMap(accounts)

	handler := NewHandler(model.NewMachine(1000000), accountMap)

	var tests = []struct {
		accountID string
		pid       string
	}{
		{"3859459814", "7386"},
		{"2434597300", "4557"},
		{"9089382418", "0075"},
		{"4001377812", "5950"},
	}

	assert := assert.New(t)

	for _, test := range tests {
		assert.NotNil(handler.Login(test.accountID, test.pid))
	}
}
