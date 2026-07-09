package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetNationalTeamPlayers shows the current squad of the given national
// team.
func (a *API) GetNationalTeamPlayers(teamID id.NationalTeam) ([]*chpp.NationalTeamPlayer, error) {
	values := map[string]string{
		"actionType": "view",
		"teamID":     teamID.String(),
	}

	res, err := a.parsed.GetNationalPlayersXML(values)
	if err != nil {
		return nil, err
	}

	return res.Players, nil
}

// GetNationalTeamPlayerStats shows the Supporter statistics of the number
// of matches players have played for the given national team.
// matchTypeCategory should be "NT" (all matches) or "WC" (World Cup matches
// only); an empty string defaults to "NT". If showAll is false, only the
// top 20 players by matches played are returned.
func (a *API) GetNationalTeamPlayerStats(teamID id.NationalTeam, matchTypeCategory string, showAll bool) (*chpp.NationalPlayersStats, error) {
	values := map[string]string{
		"actionType": "SupporterStats",
		"teamID":     teamID.String(),
		"ShowAll":    strconv.FormatBool(showAll),
	}
	if matchTypeCategory != "" {
		values["MatchTypeCategory"] = matchTypeCategory
	}

	res, err := a.parsed.GetNationalPlayersXML(values)
	if err != nil {
		return nil, err
	}

	return res.Stats, nil
}
