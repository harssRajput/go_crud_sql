package server

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harssRajput/go_crud_sql/internal/handler/account"
	"io"
	"net/http"
	"os"
)

// TODO: put it in a config file. (optional: put separate config file based on envTYpe)
const (
	HTTP_PORT = "3333"
)

func getOthers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request %s not matched any route\n", r.URL.Path)
	io.WriteString(w, "You damn stumbled! unknown URL\n")
}

func RunServer() {
	initHttpServer()
}

func initHttpServer() {
	r := mux.NewRouter()

	accountsRouter := r.PathPrefix("/accounts").Subrouter().StrictSlash(true)
	accountsRouter.HandleFunc("/", account.CreateAccount).Methods("POST")
	accountsRouter.HandleFunc("/{id}", account.GetAccount).Methods("GET")

	//catch-all router
	r.PathPrefix("/").HandlerFunc(getOthers)
	//start listening
	err := http.ListenAndServe(":"+HTTP_PORT, r)
	if errors.Is(err, http.ErrServerClosed) {
		//TODO: put graceful shutdon if possible. (optional: not part of test)
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
