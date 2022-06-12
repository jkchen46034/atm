package handler

import (
	"errors"
	"time"

	"takeoff.com/atm/model"
)

func (handler *Handler) Login(accountID, pin string) (err error) {
	account, exist := handler.accountMap.GetMap()[accountID]
	if exist && account.GetPIN() == pin {
		model.Log.Add(accountID, time.Now())
		return nil
	} else {
		return errors.New("Authorization failed")
	}
}
