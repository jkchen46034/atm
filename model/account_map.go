package model

type AccountMap struct {
	m map[string]*Account
}

func NewAccountMap() *AccountMap {
	return &AccountMap{m: make(map[string]*Account)}
}

func (am *AccountMap) GetMap() map[string]*Account {
	return am.m
}

func (am *AccountMap) Add(accounts []*Account) bool {
	for _, account := range accounts {
		am.m[account.GetAccountID()] = account
	}
	return true
}
