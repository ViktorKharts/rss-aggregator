package main

import (
	"net/http"
	"strings"
)

func (c *apiConfig) usersGetHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))	
}

