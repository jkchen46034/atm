package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"takeoff.com/atm/handler"
	"takeoff.com/atm/model"
)

func shell() {
	accounts := []*model.Account{
		model.NewAccount("2859459814", "7386", 10024),
		model.NewAccount("1434597300", "4557", 9000055),
		model.NewAccount("7089382418", "0075", 0),
		model.NewAccount("2001377812", "5950", 0),
	}

	handler := handler.NewHandler(model.NewMachine(1000000), model.NewAccountMap(accounts))

	// iterate to read commands
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = exec(input, handler); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func exec(input string, handler *handler.Handler) error {

	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "authorize":
		return exec_authorize_command(args, handler)

	case "withdraw":
		return exec_withdraw_command(args, handler)

	case "deposit":
		return exec_deposit_command(args, handler)

	case "balance":
		return exec_balance_command(args, handler)

	case "history":
		return exec_history_command(args, handler)

	case "logout":
		return exec_logout_command(args, handler)

	case "demo":
		demo()
		return nil

	case "end", "exit":
		os.Exit(0)
	}
	return nil
}

func exec_authorize_command(args []string, handler *handler.Handler) error {
	if len(args) != 3 {
		return errors.New("authorize <accout_id> <pin>")
	}
	err := handler.Login(args[1], args[2])
	if err == nil {
		fmt.Println("Account", args[1], "successfully authorized")
	}
	return err
}

// amount is in unit of cent
func exec_withdraw_command(args []string, handler *handler.Handler) error {

	if len(args) != 2 {
		return errors.New("withdraw <value>")
	}

	amount, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	amount = amount * 100

	err, _, _, msg := handler.Withdraw(amount)
	if err == nil {
		fmt.Println(msg)
	}
	return err
}

func exec_deposit_command(args []string, handler *handler.Handler) error {

	if len(args) != 2 {
		return errors.New("deposit <value>")
	}

	amount, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	amount = amount * 100

	err, _, msg := handler.Deposit(amount)
	if err == nil {
		fmt.Println(msg)
	}
	return err
}

func exec_balance_command(args []string, handler *handler.Handler) error {

	if len(args) != 1 {
		return errors.New("balance")
	}

	err, balance := handler.GetBalance()
	if err == nil {
		fmt.Printf("Current balance $%d.%02d\n", balance/100, balance%100)
	}
	return err
}

func exec_history_command(args []string, handler *handler.Handler) error {

	if len(args) != 1 {
		return errors.New("history")
	}

	err, _, msg := handler.GetHistory()
	if err == nil {
		fmt.Println(msg)
	}
	return err
}

func exec_logout_command(args []string, handler *handler.Handler) error {

	if len(args) != 1 {
		return errors.New("logout")
	}

	err, accountID := handler.Logout()
	if err == nil {
		fmt.Println("Account", accountID, "logged out.")
	}
	return err
}
