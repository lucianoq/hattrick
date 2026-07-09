package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyYouthPlayers returns the list of players (id and name only) for the
// currently logged on user's youth team.
func (a *API) GetMyYouthPlayers() ([]*chpp.YouthPlayerDetail, error) {
	return a.getYouthPlayerList(nil)
}

// GetYouthPlayers returns the list of players (id and name only) for the
// given youth team.
func (a *API) GetYouthPlayers(youthTeamID id.YouthTeam) ([]*chpp.YouthPlayerDetail, error) {
	return a.getYouthPlayerList(map[string]string{
		"youthTeamID": youthTeamID.String(),
	})
}

// GetMyYouthPlayersDetails returns full details for all of the currently
// logged on user's youth players.
func (a *API) GetMyYouthPlayersDetails(showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error) {
	values := youthPlayerDetailsValues(showScoutCall, showLastMatch)
	values["actionType"] = "details"
	return a.getYouthPlayerList(values)
}

// GetYouthPlayersDetails returns full details for all youth players of the
// given youth team.
func (a *API) GetYouthPlayersDetails(youthTeamID id.YouthTeam, showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error) {
	values := youthPlayerDetailsValues(showScoutCall, showLastMatch)
	values["actionType"] = "details"
	values["youthTeamID"] = youthTeamID.String()
	return a.getYouthPlayerList(values)
}

// UnlockYouthTeamSkills unlocks the skills of the entire youth team and
// returns full details for all players. (Requires "manage_youthplayers"
// scope.)
func (a *API) UnlockYouthTeamSkills(youthTeamID id.YouthTeam, showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error) {
	values := youthPlayerDetailsValues(showScoutCall, showLastMatch)
	values["actionType"] = "unlockskills"
	values["youthTeamID"] = youthTeamID.String()
	return a.getYouthPlayerList(values)
}

func (a *API) getYouthPlayerList(values map[string]string) ([]*chpp.YouthPlayerDetail, error) {
	res, err := a.parsed.GetYouthPlayerListXML(values)
	if err != nil {
		return nil, err
	}

	return res.Players, nil
}

func youthPlayerDetailsValues(showScoutCall, showLastMatch bool) map[string]string {
	values := map[string]string{}
	if showScoutCall {
		values["showScoutCall"] = "true"
	}
	if showLastMatch {
		values["showLastMatch"] = "true"
	}
	return values
}
