package api

import (
	"encoding/json"
	"net/http"
)

// StoryPointsParams is a struct that represents the parameters for the /storypoints route
type StoryPointsParams struct {
	TeamMember string
}

// StoryPointsResponse is a struct that represents the response for the /storypoints route
type StoryPointsResponse struct {
	Code   int
	Points int64
}

// Error is a struct that represents an error response
type Error struct {
	Code    int
	Message string
}

// writeError is a helper function to write an error response
func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Error{Code: code, Message: message})
}

// RequestErrorHandler is a function that handles errors in HTTP requests
var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, http.StatusInternalServerError, err.Error())
}

// InternalErrorHandler is a function that handles internal server errors
var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, http.StatusInternalServerError, "Internal server error")
}
