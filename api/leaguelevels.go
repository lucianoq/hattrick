package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyLeagueLevels shows the division levels of the league of the
// requesting user's primary club.
func (a *API) GetMyLeagueLevels() ([]*chpp.LeagueLevelItem, error) {
	res, err := a.parsed.GetLeagueLevelsXML(nil)
	if err != nil {
		return nil, err
	}

	return res.Levels, nil
}

// GetLeagueLevels shows the division levels of the given league.
func (a *API) GetLeagueLevels(leagueID id.League) ([]*chpp.LeagueLevelItem, error) {
	values := map[string]string{
		"LeagueID": leagueID.String(),
	}

	res, err := a.parsed.GetLeagueLevelsXML(values)
	if err != nil {
		return nil, err
	}

	return res.Levels, nil
}
