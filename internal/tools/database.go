package tools

import (
	"sync"

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

// instance is a singleton instance of the DatabaseInterface.
var (
	instance *DatabaseInterface
	once     sync.Once
)

// NewDatabase creates a Singleton of the DatabaseInterface.
func NewDatabase() (*DatabaseInterface, error) {
	var err error
	once.Do(func() {
		var db DatabaseInterface = &mockDB{}
		err = db.SetupDatabase()
		if err != nil {
			log.Error(err)
			instance = nil
		} else {
			instance = &db
		}
	})
	return instance, err
}
