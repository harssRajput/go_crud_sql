package transaction

import (
	"encoding/json"
	"github.com/gorilla/mux"
	ent "github.com/harssRajput/go_crud_sql/internal/entity"
	"github.com/harssRajput/go_crud_sql/internal/service/account"
	"github.com/harssRajput/go_crud_sql/internal/service/transaction"
	us "github.com/harssRajput/go_crud_sql/internal/utilityStore"
	"log"
	"net/http"
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

func (th *transactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var trx ent.Transaction
	err := json.NewDecoder(r.Body).Decode(&trx)
	if err != nil {
		th.logger.Printf("Error decoding request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	th.logger.Printf("CreateTransaction request %v\n", trx)

	//validation
	if trx.Amount == 0 {
		http.Error(w, "Amount cannot be zero", http.StatusBadRequest)
		return
	}

	// Create the transaction
	trxResponse, err := th.transactionService.CreateTransaction(&trx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(trxResponse); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
