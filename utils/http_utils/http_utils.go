package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
)

// RespondJSON func
func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// RespondError func
func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJSON(w, err.Status, err)
}
