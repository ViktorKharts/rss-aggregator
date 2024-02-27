package main

import (
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	type readinessResponse struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, readinessResponse{
		Status: "ok",
	})
}

