package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTeamDetails ...
func (a API) GetMyTeamDetails() (*chpp.TeamDetails, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			"includeDomesticFlags": "true",
			"includeFlags":         "true",
			"includeSupporters":    "true",
		},
	)
	if err != nil {
		return nil, err
	}

	return &e.TeamDetails, nil
}

// GetTeamDetails ...
func (a API) GetTeamDetails(teamID id.Team) (*chpp.TeamDetails, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			// What team/user to show the data for. teamID and userID generates
			// the same result, except that ownerless teams can only be accessed
			// if submitting a teamID and that users without a team can only be
			// accessed if userID is submitted.
			"teamID": teamID.String(),

			"includeDomesticFlags": "true",
			"includeFlags":         "true",
			"includeSupporters":    "true",
		},
	)
	if err != nil {
		return nil, err
	}

	return &e.TeamDetails, nil
}

// GetTeamDetailsByUser ...
func (a API) GetTeamDetailsByUser(userID id.User) (*chpp.TeamDetails, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			// What team/user to show the data for. teamID and userID generates
			// the same result, except that ownerless teams can only be accessed
			// if submitting a teamID and that users without a team can only be
			// accessed if userID is submitted. If neither userID or teamID is
			// supplied, userID defaults to the logged on user's userID.
			"userID": userID.String(),

			"includeDomesticFlags": "true",
			"includeFlags":         "true",
			"includeSupporters":    "true",
		},
	)
	if err != nil {
		return nil, err
	}

	return &e.TeamDetails, nil
}
