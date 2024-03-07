package main

import "net/http"

func (c *apiConfig) feedsGetHandler(w http.ResponseWriter, r *http.Request) {
	dbFeeds, err := c.DB.GetFeeds(r.Context())	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch feeds")
		return
	}

	feeds := []Feed{}
	for _, f := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(f))
	}

	respondWithJson(w, http.StatusOK, feeds)
}

