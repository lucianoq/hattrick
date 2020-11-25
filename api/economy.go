package api

import "github.com/lucianoq/hattrick/chpp"

// GetEconomy shows all active, recently finished or hotlisted transfers for the
// specified team.
func (a *API) GetEconomy() (chpp.Economy, error) {
	e, err := a.parsed.GetEconomyXML(nil)
	if err != nil {
		return chpp.Economy{}, err
	}

	return e.Team, nil
}
