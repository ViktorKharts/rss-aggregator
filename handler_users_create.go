package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ViktorKharts/rss-aggregator/internal/database"

	"github.com/google/uuid"
)

func (c *apiConfig) usersCreateHandler(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error: request body parsing failed")
		return
	}

	user, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: reqBody.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error: create user failed")
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}

