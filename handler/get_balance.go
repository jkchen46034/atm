package handler

import ()

func (handler *Handler) GetBalance() (error, int) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, 0
	}
	balance := handler.accountMap.GetAccount(accountID).GetBalance()
	return nil, balance
}
