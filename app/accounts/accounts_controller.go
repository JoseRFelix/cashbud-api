package accounts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/JoseRFelix/cashbud-api/app/common"

	"github.com/go-playground/validator/v10"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {
	result, err := GetAllAccounts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type CreateAccountDTO struct {
	Name  string `json:"name" validate:"required"`
	Owner string `json:"owner" validate:"required"`
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var model CreateAccountDTO
	err := common.DecodeJSONBody(w, r, &model)

	if err != nil {
		var malformedRequest *common.MalformedRequest
		if errors.As(err, &malformedRequest) {
			http.Error(w, malformedRequest.Msg, malformedRequest.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	validate := validator.New()
	err = validate.Struct(model)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.TranslateError(err))
		return
	}

	err = CreateAccount(model)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model)
}
