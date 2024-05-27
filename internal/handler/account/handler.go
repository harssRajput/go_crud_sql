package account

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/service"
	"github.com/harssRajput/go_crud_sql/internal/service/account"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type accountHandler struct {
	accountService account.AccountService
	logger         *log.Logger
}

func InitAccountHandler(r *mux.Router, logger *log.Logger, serviceStore *service.ServiceStore) error {

	ah := &accountHandler{
		accountService: serviceStore.AccountService,
		logger:         logger,
	}

	addRoutes(r, ah)
	return nil
}

func addRoutes(r *mux.Router, ah *accountHandler) {
	accountsRouter := r.PathPrefix("/accounts").Subrouter().StrictSlash(true)
	accountsRouter.HandleFunc("/", ah.CreateAccount).Methods("POST")
	accountsRouter.HandleFunc("/{id}", ah.GetAccount).Methods("GET")
}

func (ah *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]
	ah.logger.Printf("GetAccount request for accountId %s\n", accountId)

	accountIdInt, err := strconv.Atoi(accountId)
	if err != nil {
		ah.logger.Println("Invalid account ID", accountId)
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	// Retrieve the account information
	acc, err := ah.accountService.GetAccountByID(accountIdInt)
	if err != nil {
		ah.logger.Printf("[accountHandler] Error fetching account for Id %d : %v\n", accountIdInt, err)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Account not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if acc == nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	//send response
	ah.logger.Printf("Account: %v\n", acc)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(*acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc account.Account
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		ah.logger.Printf("Error decoding request: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ah.logger.Printf("CreateAccount request %v\n", acc)

	//validation
	err = ValidateDocumentNumber(acc.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//sanitize input: trim spaces
	acc.DocumentNumber = strings.Trim(acc.DocumentNumber, " ")

	// Create the account
	accResponse, err := ah.accountService.CreateAccount(&acc)
	if err != nil {
		ah.logger.Printf("[accountHandler] Error creating account for account %v : %v\n", acc, err)
		if strings.Contains(err.Error(), "Duplicate entry") {
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
