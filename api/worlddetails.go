package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetWorld ...
func (a *API) GetWorld() ([]*chpp.League, error) {
	wd, err := a.parsed.GetWorldDetailsXML(
		map[string]string{
			"includeRegions": "false",
		},
	)
	if err != nil {
		return nil, err
	}

	return wd.Leagues, nil
}

// GetLeague ...
func (a *API) GetLeague(league id.League) (*chpp.League, error) {
	e, err := a.parsed.GetWorldDetailsXML(
		map[string]string{
			"includeRegions": "false",
			"leagueID":       league.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.Leagues[0], nil
}

// GetCountry ...
func (a *API) GetCountry(country id.Country) (*chpp.League, error) {
	e, err := a.parsed.GetWorldDetailsXML(
		map[string]string{
			"includeRegions": "false",
			"countryID":      country.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.Leagues[0], nil
}
