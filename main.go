package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseRFelix/cashbud-api/app/accounts"
	"github.com/JoseRFelix/cashbud-api/app/common"

	"github.com/gorilla/mux"
)

func Migrate() {
	accounts.AutoMigrate()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping Pong!")
}

func main() {
	common.ConnectDatabase()
	Migrate()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", healthCheck)

	accounts.AccountRoutes(router)

	port := "8080"
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
