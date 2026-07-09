package api

import (
	"github.com/lucianoq/hattrick/chpp"
)

// GetNationalTeamMatches shows the matches played (or to be played) by
// national teams of the given office type (national teams or U-20 teams).
func (a *API) GetNationalTeamMatches(leagueOfficeTypeID chpp.LeagueOfficeTypeID) ([]*chpp.NationalTeamMatch, error) {
	values := map[string]string{
		"LeagueOfficeTypeID": leagueOfficeTypeID.String(),
	}

	res, err := a.parsed.GetNationalTeamMatchesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Matches, nil
}
