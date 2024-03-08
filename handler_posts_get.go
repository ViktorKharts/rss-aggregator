package main

import (
	"net/http"
	"strconv"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
)

func (c *apiConfig) postsGetHandler(w http.ResponseWriter, r *http.Request, u database.User) {
	l := chi.URLParam(r, "limit")	
	lInt, err := strconv.Atoi(l)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	posts, err := c.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: u.ID,
		Limit: int32(lInt),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	respondWithJson(w, http.StatusOK, posts)
}
