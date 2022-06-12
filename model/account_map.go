package model

type AccountMap struct {
	m map[string]*Account
}

func NewAccountMap(accounts []*Account) *AccountMap {
	accountMap := &AccountMap{m: make(map[string]*Account)}
	accountMap.Add(accounts)
	return accountMap
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

func (am *AccountMap) GetAccount(accountID string) *Account {
	return am.m[accountID]
}
