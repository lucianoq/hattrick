package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetPlayerEvents shows the events that happened to the given player.
func (a *API) GetPlayerEvents(playerID id.Player) ([]*chpp.PlayerEventItem, error) {
	values := map[string]string{
		"playerID": playerID.String(),
	}

	res, err := a.parsed.GetPlayerEventsXML(values)
	if err != nil {
		return nil, err
	}

	return res.Player.PlayerEvents.Events, nil
}
