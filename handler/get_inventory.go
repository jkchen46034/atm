package handler

import ()

func (handler *Handler) GetInventory() int {
	return handler.machine.GetInventory()
}
