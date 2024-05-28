package transaction

import (
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/service/account"
	"github.com/harssRajput/go_crud_sql/internal/service/transaction"
	us "github.com/harssRajput/go_crud_sql/internal/utilityStore"
	"log"
)

type transactionHandler struct {
	transactionService transaction.TransactionService
	accountService     account.AccountService
	logger             *log.Logger
}

func InitTransactionHandler(utilityStore *us.UtilityStore) error {
	th := &transactionHandler{
		transactionService: utilityStore.ServiceStore.TransactionService,
		accountService:     utilityStore.ServiceStore.AccountService,
		logger:             utilityStore.Logger,
	}

	addRoutes(utilityStore.HttpRouter, th)
	return nil
}

func addRoutes(r *mux.Router, th *transactionHandler) {
	transactionsRouter := r.PathPrefix("/transactions").Subrouter().StrictSlash(true)
	transactionsRouter.HandleFunc("/", th.CreateTransaction).Methods("POST")
}
