package api

import (
	"errors"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTeams shows all of the requesting user's teams.
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

// GetMyPrimaryTeam shows the requesting user's primary club senior team.
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

// GetTeam shows the details of the given team.
func (a *API) GetTeam(teamID id.Team) (*chpp.Team, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			// What team/user to show the data for. teamID and userID generate
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

// GetPrimaryTeamByUser shows the given user's primary club senior team.
func (a *API) GetPrimaryTeamByUser(userID id.User) (*chpp.Team, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			// What team/user to show the data for. teamID and userID generate
			// the same result, except that ownerless teams can only be accessed
			// if submitting a teamID and that users without a team can only be
			// accessed if userID is submitted. If neither userID nor teamID is
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

// GetTeamsByUser shows all of the given user's teams.
func (a *API) GetTeamsByUser(userID id.User) ([]*chpp.Team, error) {
	e, err := a.parsed.GetTeamDetailsXML(
		map[string]string{
			// What team/user to show the data for. teamID and userID generate
			// the same result, except that ownerless teams can only be accessed
			// if submitting a teamID and that users without a team can only be
			// accessed if userID is submitted. If neither userID nor teamID is
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
