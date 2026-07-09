package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMatch shows the default match details (no play-by-play events) for the
// given match, assuming it belongs to the Hattrick main system.
func (a *API) GetMatch(mID id.Match) (*chpp.MatchDetails, error) {
	return a.GetMatchDetails(mID, chpp.SourceSystemHattrickMainSystem, false)
}

// GetMatchDetails shows match details for the given match and source system.
// If matchEvents is true, the returned MatchDetails.EventList is populated
// with the play-by-play match events.
func (a *API) GetMatchDetails(mID id.Match, sourceSystem chpp.SourceSystem, matchEvents bool) (*chpp.MatchDetails, error) {
	values := map[string]string{
		"matchEvents":  strconv.FormatBool(matchEvents),
		"matchID":      mID.String(),
		"sourceSystem": string(sourceSystem),
	}
	details, err := a.parsed.GetMatchDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &details.Match, nil
}
