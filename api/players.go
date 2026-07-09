package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyPlayers shows the requesting user's senior team's players, sorted
// by shirt number.
func (a *API) GetMyPlayers() ([]*chpp.Player, error) {
	return a.GetMyPlayersOrderedBy("")
}

// GetPlayers shows the given team's players, sorted by shirt number.
func (a *API) GetPlayers(teamID id.Team) ([]*chpp.Player, error) {
	return a.GetPlayersOrderedBy(teamID, "")
}

// GetMyPlayersOrderedBy shows the requesting user's players, sorted by the
// given field (defaults to "PlayerNumber" if empty).
func (a *API) GetMyPlayersOrderedBy(orderBy string) ([]*chpp.Player, error) {
	values := map[string]string{
		"actionType":       "view",
		"includeMatchInfo": "true",
	}
	if orderBy != "" {
		values["orderBy"] = orderBy
	}

	ps, err := a.parsed.GetPlayersXML(values)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}

// GetPlayersOrderedBy shows the given team's players, sorted by the given
// field (defaults to "PlayerNumber" if empty).
func (a *API) GetPlayersOrderedBy(teamID id.Team, orderBy string) ([]*chpp.Player, error) {
	values := map[string]string{
		"actionType":       "view",
		"includeMatchInfo": "true",
		"teamID":           teamID.String(),
	}
	if orderBy != "" {
		values["orderBy"] = orderBy
	}

	ps, err := a.parsed.GetPlayersXML(values)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}

// GetMyOldPlayers shows the players who no longer belong to the requesting
// user's club but were promoted from its youth squad, or were part of its
// start-up squad.
func (a *API) GetMyOldPlayers() ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType": "viewOldies",
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}

// GetOldPlayers shows the players who no longer belong to the given club but
// were promoted from its youth squad, or were part of its start-up squad.
func (a *API) GetOldPlayers(teamID id.Team) ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType": "viewOldies",
			"teamID":     teamID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}

// GetMyOldCoaches shows the requesting user's former players who are now
// coaches.
func (a *API) GetMyOldCoaches() ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType": "viewOldCoaches",
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}

// GetOldCoaches shows the given club's former players who are now coaches.
func (a *API) GetOldCoaches(teamID id.Team) ([]*chpp.Player, error) {
	ps, err := a.parsed.GetPlayersXML(
		map[string]string{
			"actionType": "viewOldCoaches",
			"teamID":     teamID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.Team.Players, nil
}
