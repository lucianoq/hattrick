package api

import (
	"github.com/lucianoq/hattrick/chpp"
)

// GetNationalTeams shows the national teams of the given office type
// (national teams or U-20 teams), and the cups they take part in.
func (a *API) GetNationalTeams(leagueOfficeTypeID chpp.LeagueOfficeTypeID) (*chpp.NationalTeamsList, error) {
	values := map[string]string{
		"LeagueOfficeTypeID": leagueOfficeTypeID.String(),
	}

	res, err := a.parsed.GetNationalTeamsXML(values)
	if err != nil {
		return nil, err
	}

	return res.NationalTeamsList, nil
}
