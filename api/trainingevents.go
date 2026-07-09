package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetTrainingEvents shows the skill-level change events for the given
// player.
func (a *API) GetTrainingEvents(playerID id.Player) ([]*chpp.TrainingEvent, error) {
	values := map[string]string{
		"playerID": playerID.String(),
	}

	res, err := a.parsed.GetTrainingEventsXML(values)
	if err != nil {
		return nil, err
	}

	return res.Player.TrainingEvents.Events, nil
}
