package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetFinancesForPrimaryTeam shows all active, recently finished or hotlisted
// transfers for the specified team.
func (a *API) GetFinancesForPrimaryTeam() (chpp.Economy, error) {
	e, err := a.parsed.GetEconomyXML(nil)
	if err != nil {
		return chpp.Economy{}, err
	}

	return e.Team, nil
}

// GetFinances shows all active, recently finished or hotlisted transfers
// for the specified team.
func (a *API) GetFinances(teamID id.Team) (chpp.Economy, error) {
	values := map[string]string{}
	values["teamId"] = teamID.String()
	e, err := a.parsed.GetEconomyXML(values)
	if err != nil {
		return chpp.Economy{}, err
	}

	return e.Team, nil
}
