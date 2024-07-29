package main

import (
	"fmt"
	"net/http"

	"github.com/adefiks/GoAPI/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API server on port 8000")
	fmt.Println(`
	_____ ____  _____ _____ _____ _____ _____    _____ _____ _____ 
	|  _  |    \|   __|   __|     |  |  |   __|  |  _  |  _  |     |
	|     |  |  |   __|   __|-   -|    -|__   |  |     |   __|-   -|
	|__|__|____/|_____|__|  |_____|__|__|_____|  |__|__|__|  |_____|
		`)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
