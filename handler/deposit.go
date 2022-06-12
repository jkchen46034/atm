package handler

import (
	"fmt"
)

func (handler *Handler) Deposit(amount int) (err error, balance int, message string) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, 0, ""
	}
	account := handler.accountMap.GetAccount(accountID)
	balance = account.Deposit(amount)
	handler.machine.Add(amount)

	return nil, balance, fmt.Sprintf("Current balance: $%d.%02d.", balance/100, balance%100)
}
