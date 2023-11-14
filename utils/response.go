package utils

import (
	"encoding/json"
	"net/http"
)

func RespondError(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := make(map[string]string)
	response["error"] = message

	json.NewEncoder(w).Encode(response)
}
func Respond(w http.ResponseWriter, code int, payload interface{}) {

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(payload)
}
