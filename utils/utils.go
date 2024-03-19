package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, i interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		return
	}
}
