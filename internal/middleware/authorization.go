package middleware

import (
	"errors"
	"net/http"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

var errUnauthorizedUserOrToken = errors.New("unauthorized User or Token")

// Authorization is a middleware function that performs authorization checks
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the value of the "TeamMember" query parameter from the request URL
		var TeamMember string = r.URL.Query().Get("TeamMember")
		var err error
		// Get the value of the "Authorization" header from the request
		token := r.Header.Get("Authorization")

		// Check if either the "TeamMember" or "Authorization" values are empty
		if TeamMember == "" || token == "" {
			log.Error(errUnauthorizedUserOrToken)
			api.RequestErrorHandler(w, errUnauthorizedUserOrToken)
			return
		}

		// Create a new instance of the DatabaseInterface
		database, err := tools.NewDatabase()

		// Check if there was an error creating the database instance
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		// Get the login details for the specified TeamMember from the database
		loginDetails := (*database).GetUserLoginDetails(TeamMember)

		// Check if the login details are nil or if the token doesn't match the stored token
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(errUnauthorizedUserOrToken)
			api.RequestErrorHandler(w, errUnauthorizedUserOrToken)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
