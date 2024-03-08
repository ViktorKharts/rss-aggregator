package main

import (
	"net/http"
	"strconv"

	"github.com/ViktorKharts/rss-aggregator/internal/database"
)

const defaultLimitValue = 10

func (c *apiConfig) postsGetHandler(w http.ResponseWriter, r *http.Request, u database.User){
	l, err := getLimit(r.URL.Query().Get("limit"))	
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	posts, err := c.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: u.ID,
		Limit: l,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	respondWithJson(w, http.StatusOK, posts)
}

func getLimit(l string) (int32, error) {
	if l == "" {
		return int32(defaultLimitValue), nil
	}
	lInt, err := strconv.Atoi(l)
	if err != nil {
		return 0, err
	}
	return int32(lInt), nil
} 
