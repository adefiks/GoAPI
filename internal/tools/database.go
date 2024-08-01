package tools

import (
	log "github.com/sirupsen/logrus"
)

// LoginDetails represents the login details of a user.
type LoginDetails struct {
	AuthToken  string
	TeamMember string
}

// StoryPointsDetails represents the story points details of a user.
type StoryPointsDetails struct {
	Points     int64
	TeamMember string
}

// DatabaseInterface is an interface that defines the methods for interacting with the database.
type DatabaseInterface interface {
	GetUserLoginDetails(TeamMember string) *LoginDetails
	GetUserStoryPoints(TeamMember string) *StoryPointsDetails
	SetupDatabase() error
}

// NewDatabase creates a new instance of the DatabaseInterface.
func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDB{}
	err := db.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
