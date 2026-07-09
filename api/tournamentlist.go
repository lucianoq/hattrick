package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTournaments shows the tournaments the requesting user's primary
// senior team takes part in.
func (a *API) GetMyTournaments() ([]*chpp.Tournament, error) {
	res, err := a.parsed.GetTournamentListXML(nil)
	if err != nil {
		return nil, err
	}

	return res.Tournaments, nil
}

// GetTournaments shows the tournaments the given team takes part in.
func (a *API) GetTournaments(teamID id.Team) ([]*chpp.Tournament, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}

	res, err := a.parsed.GetTournamentListXML(values)
	if err != nil {
		return nil, err
	}

	return res.Tournaments, nil
}
