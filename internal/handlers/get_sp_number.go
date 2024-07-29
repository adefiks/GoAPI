package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adefiks/GoAPI/api"
	"github.com/adefiks/GoAPI/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetStoryPoints(w http.ResponseWriter, r *http.Request) {
	var err error
	var params api.StoryPointsParams
	var decoder *schema.Decoder = schema.NewDecoder()

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	tokenDetails := (*database).GetUserStoryPoints(params.TeamMember)
	if tokenDetails == nil {
		api.InternalErrorHandler(w)
		return
	}

	var response = api.StoryPointsResponse{
		Code:   http.StatusOK,
		Points: (*tokenDetails).Points,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
