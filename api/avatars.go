package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAvatarsMyPlayers shows the avatars of the requesting user's own
// team's players.
func (a *API) GetAvatarsMyPlayers() ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "players"
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players, nil

}

// GetAvatarsPlayers shows the avatars of the given team's players.
func (a *API) GetAvatarsPlayers(teamID id.Team) ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "players"
	values["teamId"] = teamID.String()
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players, nil
}

// GetAvatarsMyHallOfFame shows the avatars of the requesting user's own
// team's Hall of Fame players.
func (a *API) GetAvatarsMyHallOfFame() ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "hof"
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players, nil
}

// GetAvatarsHallOfFame shows the avatars of the given team's Hall of Fame
// players.
func (a *API) GetAvatarsHallOfFame(teamID id.Team) ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "hof"
	values["teamId"] = teamID.String()
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players, nil
}
