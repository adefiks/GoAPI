package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken  string
	TeamMember string
}

type StoryPointsDetails struct {
	Points     int64
	TeamMember string
}

type DatabaseInterface interface {
	GetUserLoginDetails(TeamMember string) *LoginDetails
	GetUserStoryPoints(TeamMember string) *StoryPointsDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDB{}
	err := db.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
