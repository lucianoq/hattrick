package api

import (
	"errors"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetWorld shows every league in Hattrick.
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

// GetLeague shows the details of the given league.
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
	if len(e.Leagues) == 0 {
		return nil, errors.New("league not found")
	}

	return e.Leagues[0], nil
}

// GetCountry shows the details of the league associated with the given
// country.
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
	if len(e.Leagues) == 0 {
		return nil, errors.New("country not found")
	}

	return e.Leagues[0], nil
}
