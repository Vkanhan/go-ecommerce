package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError sends a JSON error response with the specified HTTP status code.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg,})
}

// respondWithJSON sends a JSON response with the specified HTTP status code.
func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(code)
		return
	} 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}