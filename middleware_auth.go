package main

import (
	"fmt"
	"net/http"

	"github.com/Vkanhan/go-aggregator/internal/auth"
	"github.com/Vkanhan/go-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get API key from headers
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid API Key: %v", err))
			return
		}

		// Fetch the user associated with the API key
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("Invalid user credentials: %v", err))
			return
		}
		// Pass user to the handler
		handler(w, r, user)
	}
}
