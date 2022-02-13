package accounts

import (
	"encoding/json"
	"net/http"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {
	// var accounts []AccountModel
	result, err := GetAllAccounts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
