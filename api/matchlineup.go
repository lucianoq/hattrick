package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyMatchLineup ...
func (a *API) GetMyMatchLineup(matchID id.Match) (*chpp.MatchLineup, error) {
	ps, err := a.parsed.GetMatchLineupXML(
		map[string]string{
			"matchID": matchID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.MatchLineup, nil
}

// GetMatchLineup ...
func (a *API) GetMatchLineup(matchID id.Match, teamID id.Team) (*chpp.MatchLineup, error) {
	ps, err := a.parsed.GetMatchLineupXML(
		map[string]string{
			"matchID": matchID.String(),
			"teamID":  teamID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return ps.MatchLineup, nil
}
