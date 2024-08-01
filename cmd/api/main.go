package main

import (
	"fmt"
	"net/http"

	"github.com/adefiks/GoAPI/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // Enable reporting of caller information in log messages

	var r *chi.Mux = chi.NewRouter() // Create a new router using chi package
	handlers.Handler(r)              // Register the handlers for the router

	fmt.Println("Starting GO API server on port 8000")
	fmt.Println(`
	_____ ____  _____ _____ _____ _____ _____    _____ _____ _____ 
	|  _  |    \|   __|   __|     |  |  |   __|  |  _  |  _  |     |
	|     |  |  |   __|   __|-   -|    -|__   |  |     |   __|-   -|
	|__|__|____/|_____|__|  |_____|__|__|_____|  |__|__|__|  |_____|
		`)

	err := http.ListenAndServe(":8000", r) // Start the HTTP server on port 8000
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err) // Log a fatal error if the server fails to start
	}
}
