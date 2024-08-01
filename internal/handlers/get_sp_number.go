package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// GetStoryPoints is a handler function for retrieving story points.
func GetStoryPoints(w http.ResponseWriter, r *http.Request) {
	var err error
	var params api.StoryPointsParams
	var decoder *schema.Decoder = schema.NewDecoder()

	// Decode the query parameters into the params struct
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Create a new instance of the database interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Get the user's story points from the database
	tokenDetails := (*database).GetUserStoryPoints(params.TeamMember)
	if tokenDetails == nil {
		api.InternalErrorHandler(w)
		return
	}

	// Create the response object
	var response = api.StoryPointsResponse{
		Code:   http.StatusOK,
		Points: (*tokenDetails).Points,
	}

	// Encode the response object as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
