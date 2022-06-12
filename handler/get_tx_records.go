package handler

import (
	"fmt"

	"takeoff.com/atm/model"
)

func (handler *Handler) GetTxHistory() (err error, records []model.TxRecord, message *string) {
	authorized, accountID := handler.IsAuthorized()

	if authorized != nil {
		return authorized, nil, nil
	}

	records = handler.accountMap.GetAccount(accountID).GetTxHistory()
	if len(records) == 0 {
		msg := "No history found."
		return nil, []model.TxRecord{}, &msg
	}

	var msg string
	for _, record := range records {
		msg += fmt.Sprintf("%s %d %d\n", record.TxTime.Format("2006-01-02 15:04:05"), record.Amount, record.BalanceAfter)
	}
	return nil, records, &msg
}
