package account

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/model/account"
	"log"
	"net/http"
	"strconv"
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
	log.Printf("POST account request\n")
}
