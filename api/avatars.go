package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAvatarsMyPlayers ...
func (a *API) GetAvatarsMyPlayers() ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "players"
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players.Player, nil

}

// GetAvatarsPlayers ...
func (a *API) GetAvatarsPlayers(teamID id.Team) ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "players"
	values["teamId"] = teamID.String()
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players.Player, nil
}

// GetAvatarsMyHallOfFame ...
func (a *API) GetAvatarsMyHallOfFame() ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "hof"
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players.Player, nil
}

// GetAvatarsHallOfFame ...
func (a *API) GetAvatarsHallOfFame(teamID id.Team) ([]*chpp.PlayerAvatars, error) {
	values := map[string]string{}
	values["actionType"] = "hof"
	values["teamId"] = teamID.String()
	avatars, err := a.parsed.GetAvatarsXML(values)
	if err != nil {
		return nil, err
	}
	return avatars.Team.Players.Player, nil
}
