package api

import (
	"errors"
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTeams ...
func (a *API) GetMyTeams() ([]*chpp.Team, error) {
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

	return e.Teams, nil
}

// GetMyMainTeam ...
func (a *API) GetMyPrimaryTeam() (*chpp.Team, error) {
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

	for _, team := range e.Teams {
		if team.IsPrimaryClub {
			return team, nil
		}
	}

	return nil, errors.New("user without a team")
}

// GetTeam ...
func (a *API) GetTeam(teamID id.Team) (*chpp.Team, error) {
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

	for _, team := range e.Teams {
		if team.ID == teamID {
			return team, nil
		}
	}

	return nil, errors.New("team not found")
}

// GetPrimaryTeamByUser ...
func (a *API) GetPrimaryTeamByUser(userID id.User) (*chpp.Team, error) {
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

	for _, team := range e.Teams {
		if team.IsPrimaryClub {
			return team, nil
		}
	}

	return nil, errors.New("team not found")
}

// GetTeamsByUser ...
func (a *API) GetTeamsByUser(userID id.User) ([]*chpp.Team, error) {
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

	return e.Teams, nil
}
