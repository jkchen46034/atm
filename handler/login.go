package handler

import (
	"fmt"
)

func (handler *Handler) Login(accountID, pid string) error {
	account, exist := handler.accountMap.GetMap()[accountID]
	fmt.Println(exist, account)
	return nil
}
