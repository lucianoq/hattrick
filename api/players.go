package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyPlayers ...
func (a *API) GetMyPlayers() ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType":       "view",
			"includeMatchInfo": "true",
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.PlayerList.Players, nil
}

// GetPlayers ...
func (a *API) GetPlayers(teamID id.Team) ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType":       "view",
			"includeMatchInfo": "true",
			"teamID":           teamID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.PlayerList.Players, nil
}
