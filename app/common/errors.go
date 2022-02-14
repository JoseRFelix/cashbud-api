package common

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}

func TranslateError(err error) (errs []ApiError) {
	if err == nil {
		return nil
	}

	var validatorErrors validator.ValidationErrors

	if errors.As(err, &validatorErrors) {
		out := make([]ApiError, len(validatorErrors))
		for i, fe := range validatorErrors {
			out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
		}
		return out
	}

	return make([]ApiError, 0)
}
