package users

import (
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	result, err := GetAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
