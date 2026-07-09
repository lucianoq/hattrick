package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyYouthLeague returns the youth league details for the youth team of
// the currently logged on user's primary club.
func (a *API) GetMyYouthLeague() (*chpp.YouthLeagueDetails, error) {
	details, err := a.parsed.GetYouthLeagueDetailsXML(nil)
	if err != nil {
		return nil, err
	}

	return &details.YouthLeagueDetails, nil
}

// GetYouthLeague returns the details for the given youth league.
func (a *API) GetYouthLeague(youthLeagueID id.YouthLeague) (*chpp.YouthLeagueDetails, error) {
	values := map[string]string{
		"youthleagueid": youthLeagueID.String(),
	}
	details, err := a.parsed.GetYouthLeagueDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &details.YouthLeagueDetails, nil
}
