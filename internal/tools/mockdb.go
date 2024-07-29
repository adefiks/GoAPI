package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"John": {AuthToken: "1234", TeamMember: "John"},
	"Jane": {AuthToken: "5678", TeamMember: "Jane"},
	"Jack": {AuthToken: "91011", TeamMember: "Jack"},
	"Jill": {AuthToken: "121314", TeamMember: "Jill"},
}

var mockStoryPointsDetails = map[string]StoryPointsDetails{
	"John": {Points: 5, TeamMember: "John"},
	"Jane": {Points: 10, TeamMember: "Jane"},
	"Jack": {Points: 15, TeamMember: "Jack"},
	"Jill": {Points: 20, TeamMember: "Jill"},
}

func (m *mockDB) GetUserLoginDetails(TeamMember string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var TeamMemberData = LoginDetails{}
	TeamMemberData, ok := mockLoginDetails[TeamMember]
	if !ok {
		return nil
	}
	return &TeamMemberData
}

func (m *mockDB) GetUserStoryPoints(TeamMember string) *StoryPointsDetails {
	time.Sleep(1 * time.Second)
	var TeamMemberData = StoryPointsDetails{}
	TeamMemberData, ok := mockStoryPointsDetails[TeamMember]
	if !ok {
		return nil
	}
	return &TeamMemberData
}

func (m *mockDB) SetupDatabase() error {
	return nil
}
