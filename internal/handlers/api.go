package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/middleware"
	log "github.com/sirupsen/logrus"
)

// Handler is a function that sets up the routes and middleware for the API.
func Handler(mux *http.ServeMux) {
	mux.Handle("/", stripSlashesMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/account") {
			authorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch {
				case strings.HasPrefix(r.URL.Path, "/account/team"):
					addTeamMemberMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						err := errors.New("not Implemented")
						log.Error(err)
						api.InternalErrorHandler(w)
					})).ServeHTTP(w, r)
				case strings.HasPrefix(r.URL.Path, "/account/points"):
					GetStoryPoints(w, r)
				default:
					http.NotFound(w, r)
				}
			})).ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})))
}

// Middleware to strip trailing slashes
func stripSlashesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.StripSlashes(next).ServeHTTP(w, r)
	})
}

// Middleware for authorization
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.Authorization(next).ServeHTTP(w, r)
	})
}

// Middleware to add team member
func addTeamMemberMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.AddTeamMemberMiddleware(next).ServeHTTP(w, r)
	})
}
