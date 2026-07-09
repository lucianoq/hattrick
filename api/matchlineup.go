package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyMatchLineup shows the lineup for the given match and the requesting
// user's own team, assuming it belongs to the Hattrick main system.
func (a *API) GetMyMatchLineup(matchID id.Match) (*chpp.MatchLineup, error) {
	return a.GetMatchLineupBySourceSystem(matchID, 0, chpp.SourceSystemHattrickMainSystem)
}

// GetMatchLineup shows the lineup for the given match and team, assuming it
// belongs to the Hattrick main system.
func (a *API) GetMatchLineup(matchID id.Match, teamID id.Team) (*chpp.MatchLineup, error) {
	return a.GetMatchLineupBySourceSystem(matchID, teamID, chpp.SourceSystemHattrickMainSystem)
}

// GetMatchLineupBySourceSystem shows the lineup for the given match and
// team, in the given source system (e.g. youth). If teamID is 0, the
// logged-on user's own team is used.
func (a *API) GetMatchLineupBySourceSystem(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem) (*chpp.MatchLineup, error) {
	values := map[string]string{
		"matchID":      matchID.String(),
		"sourceSystem": string(sourceSystem),
	}
	if teamID != 0 {
		values["teamID"] = teamID.String()
	}

	ps, err := a.parsed.GetMatchLineupXML(values)
	if err != nil {
		return nil, err
	}

	return ps.MatchLineup, nil
}
