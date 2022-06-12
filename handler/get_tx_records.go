package handler

import (
	"takeoff.com/atm/model"
)

func (handler *Handler) GetTxHistory() (error, []model.TxRecord) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, nil
	}
	return nil, handler.accountMap.GetAccount(accountID).GetTxHistory()
}
