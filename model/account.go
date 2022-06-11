package model

type Account struct {
	accountID string
	pin       string
	balance   int
}

func NewAccount(accountID string, pin string, initialBalance int) *Account {
	return &Account{accountID, pin, initialBalance}
}

func (a *Account) GetAccountID() string {
	return a.accountID
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) SetBalance(balance int) bool {
	a.balance = balance
	return true
}
