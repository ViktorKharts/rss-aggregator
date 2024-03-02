package main

import (
	"net/http"
	"strings"
)

func (c *apiConfig) usersGetHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	if header == "" {
		respondWithError(w, http.StatusUnauthorized, "No api key provided")
		return
	}
	apiKey := strings.Split(header, " ")[1]
	
	user, err := c.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))	
}

