package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetHOFPlayersForPrimaryTeam shows the Hall of Fame players for the
// requesting user's own team.
func (a *API) GetHOFPlayersForPrimaryTeam() ([]*chpp.HOFPlayer, error) {
	e, err := a.parsed.GetHOFPlayersXML(nil)
	if err != nil {
		return nil, err
	}

	return e.PlayerList.Players, nil
}

// GetHOFPlayers shows the Hall of Fame players for the given team.
func (a *API) GetHOFPlayers(teamID id.Team) ([]*chpp.HOFPlayer, error) {
	values := map[string]string{
		"teamID": teamID.String(),
	}

	e, err := a.parsed.GetHOFPlayersXML(values)
	if err != nil {
		return nil, err
	}

	return e.PlayerList.Players, nil
}
