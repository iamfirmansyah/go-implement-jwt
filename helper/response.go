package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func Response(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if response != nil {
		err := json.NewEncoder(w).Encode(response)
		PanicIfError(err)
	}
}

func FormatValidationError(err error) []string {
	var errors []string

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
	}

	return errors

}
