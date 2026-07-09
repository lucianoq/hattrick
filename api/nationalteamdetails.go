package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetNationalTeamDetails shows the details of the given national team.
func (a *API) GetNationalTeamDetails(teamID id.NationalTeam) (*chpp.NationalTeamDetails, error) {
	values := map[string]string{
		"teamID": teamID.String(),
	}

	res, err := a.parsed.GetNationalTeamDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Team, nil
}
