package account

import (
	"errors"
)

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

func CreateAccount(account *Account) (*Account, error) {
	//generate account id. In real scenario, it's auto-generated DB field
	account.AccountId = len(accounts) + 1

	// check if account already exists
	for _, acc := range accounts {
		if acc.DocumentNumber == account.DocumentNumber {
			return nil, errors.New("account already exists")
		}
	}

	//accounts array is not thread-safe as eventually it will be replaced by DB.
	accounts = append(accounts, *account)
	return account, nil
}
