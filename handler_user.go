package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vkanhan/go-aggregator/internal/auth"
	"github.com/Vkanhan/go-aggregator/internal/database"
	"github.com/google/uuid"
)

// handlerCreateUserreads the user's name from the request body, generates a new user ID (UUID) and inserts the user into the database.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	// Input validation
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}

	// Create the user with a UUID and current timestamp.
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUse(user))
}

// handlerGetUser retrieves a user based on the API key provided in the request header.
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Missing or invalid API key: %v", err))
		return
	}

	// Fetch the user from the database using the API key.
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUse(user))
}
