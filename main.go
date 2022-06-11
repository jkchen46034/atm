package main

import (
	"fmt"
	"takeoff.com/atm/handler"
	"takeoff.com/atm/model"
)

func main() {
	fmt.Println("Hello Takeoff")

	accountMap := model.NewAccountMap()

	acc1 := model.NewAccount("2859459814", "7386", 1024)
	acc2 := model.NewAccount("1434597300", "4557", 9000055)
	acc3 := model.NewAccount("7089382418", "0075", 0)
	acc4 := model.NewAccount("2001377812", "5950", 0)

	accountMap.Add([]*model.Account{acc1, acc2, acc3, acc4})

	for accountID, account := range accountMap.GetMap() {
		fmt.Println(accountID, account)
	}
	machine := model.NewMachine(1000000)
	fmt.Println(machine)

	handler := handler.NewHandler(machine, accountMap)
	err := handler.Login("2859459814", "7386")
	fmt.Println(err)

}
