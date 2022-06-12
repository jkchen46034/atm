package main

import (
	"fmt"
	"takeoff.com/atm/handler"
	"takeoff.com/atm/model"
)

func main() {

	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10000),
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
    fmt.Println(accountID, "successfully autorized")
  } else {
    fmt.Println("Authorized failed.")
  }
  err, balance := handler.GetBalance()
  fmt.Println("Current balance", balance)

  fmt.Println("Depositing 3000")
  amount := 3000
  err, balance = handler.Deposit(amount)
  fmt.Println("Current balance", balance)

  fmt.Println("Withdrawing 3000")
  amount = 3000
  err, amount_dispensed, balance := handler.Withdraw(3000)
  fmt.Println("Amount dispensed:", amount_dispensed)
  fmt.Println("Current balance:", balance)
  
  fmt.Println("Getting transaction history")
  err, txHistory := handler.GetTxHistory()
  for _, tx := range txHistory {
    fmt.Println(tx.TxTime.Format("2006-01-02 15:04:05"), tx.Amount, tx.Balance_after)
  }
  
  /*
  err = handler.Logout()
  /*
  err = hadnler.End()
  */
}
