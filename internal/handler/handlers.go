package handler

import (
	"github.com/harssRajput/go_crud_sql/internal/handler/account"
	"github.com/harssRajput/go_crud_sql/internal/handler/transaction"
	us "github.com/harssRajput/go_crud_sql/internal/utilityStore"
	"io"
	"log"
	"net/http"
)

func InitHandlers(utilityStore *us.UtilityStore) error {
	err := account.InitAccountHandler(utilityStore)
	if err != nil {
		utilityStore.Logger.Printf("Error initializing account handler: %v\n", err)
		return err
	}

	err = transaction.InitTransactionHandler(utilityStore)
	if err != nil {
		utilityStore.Logger.Printf("Error initializing transaction handler: %v\n", err)
		return err
	}

	//catch-all router. must be last in the list
	utilityStore.HttpRouter.PathPrefix("/").HandlerFunc(handleOthers)
	return nil
}

func handleOthers(w http.ResponseWriter, r *http.Request) {
	log.Printf("request %s %s not matched any route\n", r.Method, r.URL.Path)
	io.WriteString(w, "Whoops! That place seems to be off the map. How about trying a new spot?\n")
}
