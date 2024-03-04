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
	type response struct {
		Feed database.Feed `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
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
		return
	}

	ff, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: u.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, response{
		Feed: feed,
		FeedFollow: ff,
	})
}
