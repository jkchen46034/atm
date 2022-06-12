package handler

import (
	"takeoff.com/atm/model"
)

func (handler *Handler) Logout() (err error, accountID string) {
	authorized, accountID := handler.IsAuthorized()
	if authorized != nil {
		return authorized, ""
	}
	model.Log.Logout()
	return nil, accountID
}
