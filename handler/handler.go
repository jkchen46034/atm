package handler

import (
	"takeoff.com/atm/model"
)

type Handler struct {
	machine    *model.Machine
	accountMap *model.AccountMap
}

func NewHandler(machine *model.Machine, accountMap *model.AccountMap) *Handler {
	return &Handler{machine, accountMap}
}
