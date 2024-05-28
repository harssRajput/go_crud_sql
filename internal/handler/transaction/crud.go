package transaction

import (
	"encoding/json"
	ent "github.com/harssRajput/go_crud_sql/internal/entity"
	"net/http"
	"strings"
)

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
		if strings.Contains(err.Error(), "foreign key constraint fails") || strings.Contains(err.Error(), "invalid amount") {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
