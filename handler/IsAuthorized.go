package handler

import (
	"errors"
	"time"

	"takeoff.com/atm/model"
)

func (handler *Handler) IsAuthorized() (authorized error, accountID string) {
	authorized, accountID = model.Log.IsAuthorized(time.Now())
	if authorized != nil {
		return errors.New("Authorization required"), ""
	} else {
		return nil, accountID
	}
}
