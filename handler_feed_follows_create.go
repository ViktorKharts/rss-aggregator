package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (c *apiConfig) feedFollowsCreateHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"` 
	}	
	decoder := json.NewDecoder(r.Body)
	params := parameters {}
	decoder.Decode(&params)
	
	ff, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: params.FeedID,
		UserID: u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, ff)
}
