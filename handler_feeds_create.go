package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) feedsCreateHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	type params struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}	
	decoder := json.NewDecoder(r.Body)
	body := params{}
	decoder.Decode(&body)
	
	feed, err := c.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: body.Name,
		Url: body.Url,
		UserID: u.ID,
	})	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, feed)
}
