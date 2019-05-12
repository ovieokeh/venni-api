package venni

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse - structure of an HTTP response object
type HTTPResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    *struct{} `json:"data"`
}

// RespondWithError - send an error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := HTTPResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	RespondWithJSON(w, code, response)
}

// RespondWithJSON - send a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload HTTPResponse) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
