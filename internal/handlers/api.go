package handlers

import (
	"github.com/adefiks/GoAPI/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

// Handler is a function that sets up the routes and middleware for the API.
func Handler(r *chi.Mux) {
	// Use the StripSlashes middleware to remove trailing slashes from the URL.
	r.Use(chimiddleware.StripSlashes)

	// Create a sub-route for the "/account" path.
	r.Route("/account", func(router chi.Router) {
		// Use the Authorization middleware for all routes under "/account".
		router.Use(middleware.Authorization)

		// Handle GET requests to "/account/points" with the GetStoryPoints function.
		router.Get("/points", GetStoryPoints)
	})
}
