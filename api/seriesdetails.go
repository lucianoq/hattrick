package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMySeries shows the details of the requesting user's own team's
// series (league table).
func (a *API) GetMySeries() (*chpp.Series, error) {
	e, err := a.parsed.GetLeagueDetailsXML(nil)
	if err != nil {
		return nil, err
	}

	return e.Series, nil
}

// GetSeries shows the details of the given series (league table).
func (a *API) GetSeries(leagueLevelUnitID id.Series) (*chpp.Series, error) {
	e, err := a.parsed.GetLeagueDetailsXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.Series, nil
}
