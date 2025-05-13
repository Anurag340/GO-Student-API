package response

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func ValidationErrors(errs validator.ValidationErrors) *http.Response {
	var errors []string
	for _, err := range errs {
		switch err.ActualTag(){
			case "required":
				errors = append(errors, "Field "+err.Field() + " is required")
			default:
				errors = append(errors, "Field "+err.Field() + " is invalid")
		}
	}


	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(strings.Join(errors, ","))),
	}

}