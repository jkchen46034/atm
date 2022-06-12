package handler

import ()

func (handler *Handler) Withdraw(amount int) (error, int, int) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, 0, 0
	}
	account := handler.accountMap.GetAccount(accountID)

	withdraw_this := amount - amount%2000
	balance := account.Withdraw(withdraw_this)
	return nil, withdraw_this, balance
}
