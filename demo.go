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
	fmt.Printf("Current balance $%d.%d\n", balance/100, balance%100)

	fmt.Println("Withdrawing $90.00")
	err, _, _, msg := handler.Withdraw(9000)
	fmt.Println(msg)

	fmt.Println("Depositing $30.00")
	err, balance, msg = handler.Deposit(3000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $40.00")
	err, _, _, msg = handler.Withdraw(4000)
	fmt.Println(msg)

	fmt.Println("Depositing $40.00")
	err, _, msg = handler.Deposit(4000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $30.00")
	err, _, _, msg = handler.Withdraw(3000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $200.00")
	err, _, _, msg = handler.Withdraw(20000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $400.00")
	err, _, _, msg = handler.Withdraw(40000)
	fmt.Println(err)

	fmt.Println("Depositing $400.00")
	err, _, msg = handler.Deposit(40000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $90.00")
	err, _, _, msg = handler.Withdraw(9000)
	fmt.Println(msg)

	fmt.Println("Getting transaction history")
	err, _, msg = handler.GetHistory()
	fmt.Println(msg)

	err, accountID = handler.Logout()
	if err == nil {
		fmt.Println("Account", accountID, "logged out.")
	} else {
		fmt.Println(err)
	}

	fmt.Println("Withdrawing $20.00")
	err, _, _, _ = handler.Withdraw(2000)
	fmt.Println(err)

	fmt.Println("Getting balance")
	err, balance = handler.GetBalance()
	fmt.Println(err)

	err = handler.Login("1434597300", "4557")
	if err == nil {
		fmt.Println("Account 1434597300 successfully authorized")
	}
	err, balance = handler.GetBalance()
	fmt.Printf("Current balance $%d.%d\n", balance/100, balance%100)
	fmt.Println("Withdrawing $20000.00")
	err, _, _, msg = handler.Withdraw(2000000)
	fmt.Println(msg)

	fmt.Println("Withdrawing $10000.00")
	err, _, _, msg = handler.Withdraw(1000000)
	fmt.Println(err, msg)

	fmt.Println("Getting transaction history")
	err, _, msg = handler.GetHistory()
	fmt.Println(msg)

	err, accountID = handler.Logout()
	if err == nil {
		fmt.Println("Account", accountID, "logged out.")
	} else {
		fmt.Println(err)
	}

	accountID = "2859459814"
	err = handler.Login(accountID, "7386")
	if err == nil {
		fmt.Println("Account", accountID, "successfully authorized")
	} else {
		fmt.Println("Authorized failed.")
	}

	fmt.Println("Withdrawing $20.00")
	err, _, _, msg = handler.Withdraw(2000)
	fmt.Println(err, msg)
}

/*

jk@jk-HP:~/dev/golang/atm$ go build -o atm
jk@jk-HP:~/dev/golang/atm$ ./atm
Account 2859459814 successfully authorized
Current balance $100.24
Withdrawing $90.00
Amount dispensed: $80.00. Current balance: $20.24.
Depositing $30.00
Current balance: $50.24.
Withdrawing $40.00
Amount dispensed: $40.00. Current balance: $10.24.
Depositing $40.00
Current balance: $50.24.
Withdrawing $30.00
Amount dispensed: $20.00. Current balance: $30.24.
Withdrawing $200.00
Amount dispensed: $200.00. You have been charged overdraft fee of $5.00. Current balance: $-174.-76
Withdrawing $400.00
Your account is overdrawn! You may not make withdrawals at this time.
Depositing $400.00
Current balance: $225.24.
Withdrawing $90.00
Amount dispensed: $80.00. Current balance: $145.24.
Getting transaction history
2022-06-12 16:22:47 -80.00 145.24
2022-06-12 16:22:47 400.00 225.24
2022-06-12 16:22:47 -200.00 -174.-76
2022-06-12 16:22:47 -20.00 30.24
2022-06-12 16:22:47 40.00 50.24
2022-06-12 16:22:47 -40.00 10.24
2022-06-12 16:22:47 30.00 50.24
2022-06-12 16:22:47 -80.00 20.24

Account 2859459814 logged out.
Withdrawing $20.00
Authorization required
Getting balance
Authorization required
Account 1434597300 successfully authorized
Current balance $90000.55
Withdrawing $20000.00
Amount dispensed: $10050.00. Current balance: $79950.55.
Withdrawing $10000.00
Unable to process your withdrawal at this time.
Getting transaction history
2022-06-12 16:22:47 -10050.00 79950.55

Account 1434597300 logged out.
Account 2859459814 successfully authorized
Withdrawing $20.00
Unable to process your withdrawal at this time.
jk@jk-HP:~/dev/golang/atm$
*/
