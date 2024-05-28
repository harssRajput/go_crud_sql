package account

import (
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/service/account"
	us "github.com/harssRajput/go_crud_sql/internal/utilityStore"
	"log"
)

type accountHandler struct {
	accountService account.AccountService
	logger         *log.Logger
}

func InitAccountHandler(utilityStore *us.UtilityStore) error {

	ah := &accountHandler{
		accountService: utilityStore.ServiceStore.AccountService,
		logger:         utilityStore.Logger,
	}

	addRoutes(utilityStore.HttpRouter, ah)
	return nil
}

func addRoutes(r *mux.Router, ah *accountHandler) {
	accountsRouter := r.PathPrefix("/accounts").Subrouter().StrictSlash(true)
	accountsRouter.HandleFunc("/", ah.CreateAccount).Methods("POST")
	accountsRouter.HandleFunc("/{id}", ah.GetAccount).Methods("GET")
}
