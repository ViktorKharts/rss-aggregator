package main

import (
	"database/sql"
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
	type User struct {
		ID string 
		Name string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error: request body parsing failed")
		return
	}

	user := User{
		ID: uuid.New().String(),
		Name: reqBody.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err = c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: sql.NullTime{
			Time: user.CreatedAt,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time: user.CreatedAt,
			Valid: true,
		},
	}); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create User")
		return
	}

	respondWithJson(w, http.StatusOK, user)
}

