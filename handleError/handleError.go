package handleError

import (
	"encoding/json"
	"net/http"
)

func WriteHTTPStatus(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(HTTPError{
		Code:    statusCode,
		Message: message,
	})
}