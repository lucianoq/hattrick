package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMatch ...
func (a *API) GetMatch(mID id.Match) (*chpp.MatchDetails, error) {
	values := map[string]string{
		"matchEvents":  "false",
		"matchID":      mID.String(),
		"sourceSystem": "hattrick",
	}
	details, err := a.parsed.GetMatchDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &details.Match, nil
}
