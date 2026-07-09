package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyLadders shows the ladders the requesting user's primary senior team
// takes part in.
func (a *API) GetMyLadders() ([]*chpp.LadderListEntry, error) {
	res, err := a.parsed.GetLadderListXML(nil)
	if err != nil {
		return nil, err
	}

	return res.Ladders, nil
}

// GetLadders shows the ladders the given team takes part in.
func (a *API) GetLadders(teamID id.Team) ([]*chpp.LadderListEntry, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}

	res, err := a.parsed.GetLadderListXML(values)
	if err != nil {
		return nil, err
	}

	return res.Ladders, nil
}
