package account

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/model/account"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]
	log.Printf("GetAccount request for accountId %s\n", accountId)

	//validation
	err := validateAccountId(accountId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the account information
	accountIdInt, _ := strconv.Atoi(accountId)
	acc, err := account.GetAccountByID(accountIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if acc == nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	//send response
	log.Printf("Account: %v\n", acc)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc account.Account
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		log.Printf("Error decoding request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("CreateAccount request %v\n", acc)

	//validation
	err = validateDocumentNumber(acc.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//sanitize input: trim spaces
	acc.DocumentNumber = strings.Trim(acc.DocumentNumber, " ")

	// Create the account
	accResponse, err := account.CreateAccount(&acc)
	if err != nil {
		if strings.Contains(err.Error(), "account already exists") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(accResponse); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
