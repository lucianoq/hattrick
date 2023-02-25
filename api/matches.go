package api

import (
	"github.com/lucianoq/hattrick/chpp"
)

// GetMyMatches ...
func (a *API) GetMyMatches() ([]*chpp.Match, error) {
	values := map[string]string{}
	values["isYouth"] = "false"

	res, err := a.parsed.GetMatchesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Team.Matches, nil
}

// GetMyYouthMatches ...
func (a *API) GetMyYouthMatches() ([]*chpp.Match, error) {
	values := map[string]string{}
	values["isYouth"] = "true"

	res, err := a.parsed.GetMatchesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Team.Matches, nil
}
