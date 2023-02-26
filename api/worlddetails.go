package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetWorldDetails ...
func (a *API) GetWorldDetails() ([]*chpp.League, error) {
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

// GetLeagueDetails ...
func (a *API) GetLeagueDetails(league id.League) (*chpp.League, error) {
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

// GetCountryDetails ...
func (a *API) GetCountryDetails(country id.Country) (*chpp.League, error) {
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
