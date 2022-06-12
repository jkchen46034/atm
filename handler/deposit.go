package handler

func (handler *Handler) Deposit(amount int) (error, int) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, 0
	}
	account := handler.accountMap.GetAccount(accountID)
	balance := account.Deposit(amount)

	return nil, balance
}
