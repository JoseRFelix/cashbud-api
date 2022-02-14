package users

import (
	"github.com/gorilla/mux"
)

func UsersRoutes(router *mux.Router) {
	router.HandleFunc("/users", getUsers).Methods("GET")
}
