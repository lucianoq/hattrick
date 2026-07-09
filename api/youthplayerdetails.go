package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetYouthPlayer returns the full details for the given youth player.
func (a *API) GetYouthPlayer(youthPlayerID id.YouthPlayer, showScoutCall, showLastMatch bool) (*chpp.YouthPlayerDetail, error) {
	values := youthPlayerDetailsValues(showScoutCall, showLastMatch)
	values["youthPlayerId"] = youthPlayerID.String()

	details, err := a.parsed.GetYouthPlayerDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &details.YouthPlayer, nil
}

// UnlockYouthPlayerSkills unlocks the skills of the given youth player and
// returns their (now unlocked) details. (Requires "manage_youthplayers"
// scope.)
func (a *API) UnlockYouthPlayerSkills(youthPlayerID id.YouthPlayer, showScoutCall, showLastMatch bool) (*chpp.YouthPlayerDetail, error) {
	values := youthPlayerDetailsValues(showScoutCall, showLastMatch)
	values["actionType"] = "unlockskills"
	values["youthPlayerId"] = youthPlayerID.String()

	details, err := a.parsed.GetYouthPlayerDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &details.YouthPlayer, nil
}
