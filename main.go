package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseRFelix/cashbud-api/app/accounts"
	"github.com/JoseRFelix/cashbud-api/app/common"
	"github.com/JoseRFelix/cashbud-api/app/users"
	"github.com/urfave/negroni"

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
	users.UsersRoutes(router)

	negroni := negroni.Classic()
	negroni.UseHandler(router)

	port := common.GetEnvironmentVariable("PORT")
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, negroni))
}
