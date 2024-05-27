package account

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]

	fmt.Fprintf(w, "GET account request for accountId %s\n", accountId)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST account request\n")

}
