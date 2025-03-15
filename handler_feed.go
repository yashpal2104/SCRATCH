package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/yashpal2104/RSS-Aggregator/internal/database"
)

// Specific function signature this is the func signature that go expects to be written if you want to define an HTTP Handler in the way that GO standard library expects
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	// CreateUser -> 'sqlc' generated a func for us to create a User for us and created a params(i.e CreateuserParams) as a struct
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}

// Authenticated endpoint in order to Get your one user information you require authentication key(api_key)
// otherwise to register your account you don't need api_key

// Refactored the code to include User and more clean code
// func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request, user database.User) {

// 	respondWithJSON(w, 200, databaseFeedToFeed(feed))
// }
