package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyLeague ...
func (a *API) GetMyLeague() (*chpp.League, error) {
	e, err := a.parsed.GetLeagueDetailsXML(nil)
	if err != nil {
		return nil, err
	}

	return e.League, nil
}

// GetLeague ...
func (a *API) GetLeague(leagueLevelUnitID id.LeagueLevelUnit) (*chpp.League, error) {
	e, err := a.parsed.GetLeagueDetailsXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.League, nil
}
