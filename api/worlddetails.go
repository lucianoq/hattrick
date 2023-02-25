package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetWorldDetails ...
func (a *API) GetWorldDetails() ([]*chpp.Country, error) {
	wd, err := a.parsed.GetWorldDetailsXML(
		map[string]string{
			"includeRegions": "false",
		},
	)
	if err != nil {
		return nil, err
	}

	return wd.LeagueList.Leagues, nil
}

// GetCountryDetails ...
func (a *API) GetCountryDetails(country id.Country) (*chpp.Country, error) {
	e, err := a.parsed.GetWorldDetailsXML(
		map[string]string{
			"includeRegions": "false",
			"countryID":      country.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.LeagueList.Leagues[0], nil
}
