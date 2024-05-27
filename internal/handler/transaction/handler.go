package transaction

import (
	"encoding/json"
	"github.com/harssRajput/go_crud_sql/internal/model/transaction"
	"log"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var trx transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&trx)
	if err != nil {
		log.Printf("Error decoding request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("CreateTransaction request %v\n", trx)

	//validation
	if trx.Amount == 0 {
		http.Error(w, "Amount cannot be zero", http.StatusBadRequest)
		return
	}

	// Create the transaction
	trxResponse, err := transaction.CreateTransaction(&trx)
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
