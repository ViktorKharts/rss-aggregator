package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *apiConfig) feedFollowsDeleteHandler (w http.ResponseWriter, r *http.Request) {
	ffId := chi.URLParam(r, "feedFollowsID")

	ffID, err := uuid.Parse(ffId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = c.DB.RemoveFeedFollow(r.Context(), ffID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to remove a feed follow")
	}
	
	respondWithJson(w, http.StatusOK, "Feed follow removed")
}

