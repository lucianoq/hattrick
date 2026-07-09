package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyClub shows the requesting user's club's staff skill levels and
// youth squad investment/promotion status.
func (a *API) GetMyClub() (*chpp.Club, error) {
	team, err := a.parsed.GetClubXML(nil)
	if err != nil {
		return nil, err
	}
	return team.Team, nil
}

// GetClubByID shows the given team's club's staff skill levels and youth
// squad investment/promotion status.
func (a *API) GetClubByID(teamID id.Team) (*chpp.Club, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}
	team, err := a.parsed.GetClubXML(values)
	if err != nil {
		return nil, err
	}
	return team.Team, nil
}
