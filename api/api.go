package api

import (
	"encoding/json"
	"net/http"
)

// StoryPointsHandler is a handler function for the /storypoints route
type StoryPointsParams struct {
	TeamMember string
}

// StoryPointsHandler is a handler function for the /storypoints route
type StoryPointsResponse struct {
	Code   int
	Points int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Error{Code: code, Message: message})
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "Internal server error")
	}
)
