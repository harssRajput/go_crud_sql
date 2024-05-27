package account

type Account struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

// create seed data in memory
var accounts = []Account{
	{AccountId: 1, DocumentNumber: "0001"},
	{AccountId: 2, DocumentNumber: "0002"},
	{AccountId: 3, DocumentNumber: "0003"},
}

func GetAccountByID(accountID int) (*Account, error) {
	for _, account := range accounts {
		if account.AccountId == accountID {
			return &account, nil
		}
	}
	return nil, nil
}
