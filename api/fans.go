package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetFansForPrimaryTeam shows fan data for the requesting user's primary
// senior team.
func (a *API) GetFansForPrimaryTeam() (chpp.Fans, error) {
	e, err := a.parsed.GetFansXML(nil)
	if err != nil {
		return chpp.Fans{}, err
	}

	return e.Team, nil
}

// GetFans shows fan data for the given senior team, which must be managed by
// the requesting user.
func (a *API) GetFans(teamID id.Team) (chpp.Fans, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}

	e, err := a.parsed.GetFansXML(values)
	if err != nil {
		return chpp.Fans{}, err
	}

	return e.Team, nil
}
