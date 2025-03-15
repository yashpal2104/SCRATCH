package main

import (
	"fmt"
	"net/http"

	"github.com/yashpal2104/RSS-Aggregator/internal/auth"
	"github.com/yashpal2104/RSS-Aggregator/internal/database"
)

// This middleware is made to get api_key from authorization and header and fetch the user and then
// use that user in the handler rather than copy/pasting the func handlerGetUser() in handled_user.go
// in every handler we are going to build some middleware to DRY the code
// Note: This does not match the func signature i.e used in HTTP.Handler func (database.User is extra)
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// Closure: To also use User from db
// To convert it to a standard HTTP.{handler request func}
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user:%v", err))
			return
		}
		respondWithJSON(w, 201, databaseUserToUser(user))
	}
	
}
