package account

import (
	"database/sql"
	"log"
)

type Account struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type AccountService interface {
	GetAccountByID(accountID int) (*Account, error)
	CreateAccount(account *Account) (*Account, error)
}

type accountService struct {
	sqldb  *sql.DB
	logger *log.Logger
}

func InitAccountService(sqldb *sql.DB, log *log.Logger) (AccountService, error) {
	return &accountService{sqldb: sqldb, logger: log}, nil
}

func (as *accountService) GetAccountByID(accountID int) (*Account, error) {
	//get account from db by id
	query := `SELECT account_id, document_number FROM Account WHERE account_id = ?`
	row := as.sqldb.QueryRow(query, accountID)
	account := &Account{}
	err := row.Scan(&account.AccountId, &account.DocumentNumber)
	if err != nil {
		as.logger.Printf("Error getting account: %v\n", err)
		return nil, err
	}
	return account, nil
}

func (as *accountService) CreateAccount(account *Account) (*Account, error) {
	query := `INSERT INTO Account (document_number) VALUES (?)`
	response, err := as.sqldb.Exec(query, account.DocumentNumber)
	if err != nil {
		as.logger.Printf("Error inserting account: %v\n", err)
		return nil, err
	}
	accountId, err := response.LastInsertId()
	as.logger.Println("account created with accountId: ", accountId)
	account.AccountId = int(accountId)

	return account, nil
}
