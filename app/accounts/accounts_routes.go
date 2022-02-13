package accounts

import (
	"github.com/gorilla/mux"
)

func AccountRoutes(router *mux.Router) {
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
}
