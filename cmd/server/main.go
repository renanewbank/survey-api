// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"github.com/renanewbank/survey-api/internal/handlers"
	"github.com/renanewbank/survey-api/internal/repository"
)

func main() {
	repo := repository.NewQuestionRepository()
	handler := handlers.NewQuestionHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/questions", handler.HandleListOrCreate)
	mux.HandleFunc("/questions/", handler.HandleByID) // para GET, PUT, DELETE

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
