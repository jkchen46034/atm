package main

import (
	"fmt"
	"takeoff.com/atm/handler"
	"takeoff.com/atm/model"
)

func demo() {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	// accounts are stored as hash map; keyed by account ID
	accountMap := model.NewAccountMap(accounts)

	// money is represented in unit of cent
	handler := handler.NewHandler(model.NewMachine(1000000), accountMap)

	accountID := "2859459814"
	err := handler.Login(accountID, "7386")
	if err == nil {
		fmt.Println("Account", accountID, "successfully authorized")
	} else {
		fmt.Println("Authorized failed.")
	}
	err, balance := handler.GetBalance()
	fmt.Println("Current balance", balance)

	fmt.Println("Depositing 3000")
	amount := 3000
	err, balance = handler.Deposit(amount)
	fmt.Println("Current balance", balance)

	amount = 3000
	fmt.Println("Withdrawing 3000")
	err, amount_dispensed, balance := handler.Withdraw(amount)
	fmt.Println("Amount dispensed:", amount_dispensed)
	fmt.Println("Current balance:", balance)

	amount = 4000
	fmt.Println("Depositing", amount)
	err, balance = handler.Deposit(amount)
	fmt.Println("Current balance", balance)

	amount = 2000
	fmt.Println("Withdrawing", amount)
	err, amount_dispensed, balance = handler.Withdraw(amount)
	fmt.Println("Amount dispensed:", amount_dispensed)
	fmt.Println("Current balance:", balance)

	fmt.Println("Getting transaction history")
	err, txHistory := handler.GetTxHistory()
	for _, tx := range txHistory {
		fmt.Println(tx.TxTime.Format("2006-01-02 15:04:05"), tx.Amount, tx.BalanceAfter)
	}

	err, accountID = handler.Logout()
	if err == nil {
		fmt.Println("Account", accountID, "logged out.")
	} else {
		fmt.Println(err)
	}

	amount = 2000
	fmt.Println("Withdrawing", amount)
	err, amount_dispensed, balance = handler.Withdraw(amount)
	fmt.Println(err)

	fmt.Println("Getting balance")
	err, balance = handler.GetBalance()
	fmt.Println(err)
}

/*
jk@jk-HP:~/dev/golang/atm$ go build -o atm
jk@jk-HP:~/dev/golang/atm$ ./atm
Account 2859459814 successfully authorized
Current balance 10024
Depositing 3000
Current balance 13024
Withdrawing 3000
Amount dispensed: 2000
Current balance: 11024
Depositing 4000
Current balance 15024
Withdrawing 2000
Amount dispensed: 2000
Current balance: 13024
Getting transaction history
2022-06-11 23:54:06 -2000 13024
2022-06-11 23:54:06 4000 15024
2022-06-11 23:54:06 -2000 11024
2022-06-11 23:54:06 3000 13024
Account 2859459814 logged out.
Withdrawing 2000
Authorization required
Getting balance
Authorization required
jk@jk-HP:~/dev/golang/atm$

*/
