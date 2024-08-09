package handlers

import (
	"errors"
	"net/http"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

// Handler is a function that sets up the routes and middleware for the API.
func Handler(r *chi.Mux) {
	// Use the StripSlashes middleware to remove trailing slashes from the URL.
	r.Use(chimiddleware.StripSlashes)

	// Create a sub-route for the "/account" path.
	r.Route("/account", func(router chi.Router) {
		// Use the Authorization middleware for all routes under "/account".
		router.Use(middleware.Authorization)

		// Use the AddTeamMemberMiddleware for the "/account/team" route.
		router.With(middleware.AddTeamMemberMiddleware).Post("/team", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// This handler will only be called if the middleware does not handle the request.
			err := errors.New("not Implemented")
			log.Error(err)
			api.InternalErrorHandler(w)
		}))

		// Handle GET requests to "/account/points" with the GetStoryPoints function.
		router.Get("/points", GetStoryPoints)
	})
}
