package main

import (
	"net/http"
	
	"github.com/ViktorKharts/rss-aggregator/internal/database"
	"github.com/ViktorKharts/rss-aggregator/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header) 	
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		
		user, err := c.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		handler(w, r, user)
	} 
}

