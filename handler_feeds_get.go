package main

import "net/http"

func (c *apiConfig) feedsGetHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.DB.GetFeeds(r.Context())	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch feeds")
		return
	}

	respondWithJson(w, http.StatusOK, feeds)
}

