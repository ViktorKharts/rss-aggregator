package main

import (
	"net/http"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
)

func (c *apiConfig) usersGetHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(u))	
}

