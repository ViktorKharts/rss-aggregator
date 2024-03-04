package main

import (
	"net/http"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
)

func (c *apiConfig) feedFollowsGetHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	ff, err := c.DB.GetFeedFollowsByUser(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, ff)
} 

