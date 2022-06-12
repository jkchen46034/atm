package handler

import (
	"errors"

	"takeoff.com/atm/model"
)

func (handler *Handler) Logout() (err error, accountID string) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return errors.New("No account is currently authorized"), ""
	}
	model.Log.Logout()
	return nil, accountID
}
