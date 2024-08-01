package tools

import "time"

type mockDB struct{}

// mockLoginDetails is a map that stores mock login details for each team member.
var mockLoginDetails = map[string]LoginDetails{
	"John": {AuthToken: "1234", TeamMember: "John"},
	"Jane": {AuthToken: "5678", TeamMember: "Jane"},
	"Jack": {AuthToken: "91011", TeamMember: "Jack"},
	"Jill": {AuthToken: "121314", TeamMember: "Jill"},
}

// mockStoryPointsDetails is a map that stores mock story points details for each team member.
var mockStoryPointsDetails = map[string]StoryPointsDetails{
	"John": {Points: 5, TeamMember: "John"},
	"Jane": {Points: 10, TeamMember: "Jane"},
	"Jack": {Points: 15, TeamMember: "Jack"},
	"Jill": {Points: 20, TeamMember: "Jill"},
}

// GetUserLoginDetails retrieves the mock login details for a given team member.
func (m *mockDB) GetUserLoginDetails(TeamMember string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var TeamMemberData = LoginDetails{}
	TeamMemberData, ok := mockLoginDetails[TeamMember]
	if !ok {
		return nil
	}
	return &TeamMemberData
}

// GetUserStoryPoints retrieves mock the story points details for a given team member.
func (m *mockDB) GetUserStoryPoints(TeamMember string) *StoryPointsDetails {
	time.Sleep(1 * time.Second)
	var TeamMemberData = StoryPointsDetails{}
	TeamMemberData, ok := mockStoryPointsDetails[TeamMember]
	if !ok {
		return nil
	}
	return &TeamMemberData
}

// SetupDatabase sets up the mock database.
func (m *mockDB) SetupDatabase() error {
	return nil
}
