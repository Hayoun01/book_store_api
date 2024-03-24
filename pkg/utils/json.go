package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("%v: %v", message, code)
	}
	type resError struct {
		Error string `json:"error"`
	}
	ResponseWithJson(w, code, resError{
		Error: message,
	})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Failed to marshal JSON payload: ", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
