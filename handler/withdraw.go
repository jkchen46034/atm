package handler

import (
	"errors"
	"fmt"
)

func (handler *Handler) Withdraw(amount int) (err error, withdraw_amount int, balance int, message string) {
	authorized, accountID := handler.IsAuthorized()

	if authorized != nil {
		return authorized, 0, 0, ""
	}

	atm := handler.machine
	inventory := atm.GetInventory()

	if inventory <= 0 {
		return errors.New("Unable to process your withdrawal at this time."), 0, 0, ""
	}

	account := handler.accountMap.GetAccount(accountID)

	if account.GetBalance() < 0 {
		return errors.New("Your account is overdrawn! You may not make withdrawals at this time."), 0, 0, ""
	}

	withdraw_amount = amount - amount%2000
	if withdraw_amount > inventory {
		withdraw_amount = inventory
	}

	balance, overdraft_fee := account.Withdraw(withdraw_amount)

	handler.machine.Add(-withdraw_amount)

	var msg string
	if overdraft_fee == 0 {
		msg = fmt.Sprintf("Amount dispensed: $%d. Current balance: $%d.", withdraw_amount, balance)
	} else {
		msg = fmt.Sprintf("Amount dispensed: $%d. You have been charged overdraft fee of $%d. Current balance: $%d", withdraw_amount, overdraft_fee, balance)
	}

	return nil, withdraw_amount, balance, msg
}
