package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const (
	SERVER_PORT="PORT"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"HEAD","GET","PUT","POST","DELETE","OPTIONS"},
	}))

	v1 := chi.NewRouter()
	v1.Get("/readiness", readinessHandler)
	v1.Get("/err", errorHandler)

	r.Mount("/v1", v1)

	PORT := os.Getenv(SERVER_PORT)
	server := http.Server{
		Addr: fmt.Sprintf(":%s", PORT),
		Handler: r,
	}

	fmt.Printf("\nServer has started on PORT:%s\n", PORT)
	log.Fatal(server.ListenAndServe())
}

func respondWithJson (w http.ResponseWriter, sc int, p interface{}) {
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(sc)
	w.Write([]byte(j))
}

func respondWithError (w http.ResponseWriter, sc int, m string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, sc, errorResponse{
		Error: m,
	})
}

