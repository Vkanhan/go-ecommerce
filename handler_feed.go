package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vkanhan/go-aggregator/internal/database"
	"github.com/google/uuid"
)

// handlerCreateFeed handles the creation of a new feed for a user
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON request body")
		return 
	}

	// Create the feed
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Coundn't create user: %v", err))
		return
	}

	// Respond with the created feed
	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Coundn't get the feeds: %v", err))
		return
	}

	// Respond with the created feed
	respondWithJSON(w, http.StatusCreated, databaseFeedstoReturn(feeds))
}
