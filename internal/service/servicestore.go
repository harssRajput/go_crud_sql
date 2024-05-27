package service

import (
	"database/sql"
	"github.com/harssRajput/go_crud_sql/internal/service/account"
	"github.com/harssRajput/go_crud_sql/internal/service/transaction"
	"log"
)

type ServiceStore struct {
	AccountService     account.AccountService
	TransactionService transaction.TransactionService
}

func InitServiceStore(sqldb *sql.DB, logger *log.Logger) (*ServiceStore, error) {
	//init account service
	accountService, err := account.InitAccountService(sqldb, logger)
	if err != nil {
		logger.Printf("Error initializing account service: %v\n", err)
		return nil, err
	}

	//init transaction service
	transactionService, err := transaction.InitTransactionService(sqldb, logger)
	if err != nil {
		logger.Printf("Error initializing transaction service: %v\n", err)
		return nil, err
	}
	//create services object
	serviceObj := &ServiceStore{
		AccountService:     accountService,
		TransactionService: transactionService,
	}
	return serviceObj, nil
}
