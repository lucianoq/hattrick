package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMySeries ...
func (a *API) GetMySeries() (*chpp.Series, error) {
	e, err := a.parsed.GetLeagueDetailsXML(nil)
	if err != nil {
		return nil, err
	}

	return e.Series, nil
}

// GetSeries ...
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
