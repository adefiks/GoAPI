package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

// AddTeamMemberMiddleware is middleware that processes the addition of a new team member.
func AddTeamMemberMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/account/team" {
			var member tools.StoryPointsDetails

			if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
				// http.Error(w, err.Error(), http.StatusBadRequest)
				log.Error(err)
				api.InternalErrorHandler(w)
				return
			}

			// add team member (TBD - add to database)
			log.Info("Adding team member: ", member.TeamMember)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(member)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
