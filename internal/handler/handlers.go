package handler

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/handler/account"
	"github.com/harssRajput/go_crud_sql/internal/handler/transaction"
	"github.com/harssRajput/go_crud_sql/internal/service"
	"io"
	"log"
	"net/http"
)

func InitHandlers(r *mux.Router, sqldb *sql.DB, logger *log.Logger) error {
	serviceStore, err := service.InitServiceStore(sqldb, logger)
	if err != nil {
		logger.Printf("Error initializing services: %v\n", err)
		return err
	}

	err = account.InitAccountHandler(r, logger, serviceStore)
	if err != nil {
		logger.Printf("Error initializing account handler: %v\n", err)
		return err
	}

	err = transaction.InitTransactionHandler(r, logger, serviceStore)
	if err != nil {
		logger.Printf("Error initializing transaction handler: %v\n", err)
		return err
	}

	//catch-all router. must be last in the list
	r.PathPrefix("/").HandlerFunc(handleOthers)
	return nil
}

func handleOthers(w http.ResponseWriter, r *http.Request) {
	log.Printf("request %s %s not matched any route\n", r.Method, r.URL.Path)
	io.WriteString(w, "Whoops! That place seems to be off the map. How about trying a new spot?\n")
}
