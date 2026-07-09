package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyYouthLeagueFixtures shows the fixtures for the current season of the
// requesting user's primary youth team's league.
func (a *API) GetMyYouthLeagueFixtures() (*chpp.YouthLeagueFixtures, error) {
	res, err := a.parsed.GetYouthLeagueFixturesXML(nil)
	if err != nil {
		return nil, err
	}

	return res.YouthLeagueFixtures, nil
}

// GetYouthLeagueFixtures shows the fixtures for the given season of the
// given youth league. Pass 0 for season to use the league's current
// season.
func (a *API) GetYouthLeagueFixtures(youthLeagueID id.YouthLeague, season uint) (*chpp.YouthLeagueFixtures, error) {
	values := map[string]string{
		"youthleagueid": youthLeagueID.String(),
	}
	if season != 0 {
		values["season"] = strconv.FormatUint(uint64(season), 10)
	}

	res, err := a.parsed.GetYouthLeagueFixturesXML(values)
	if err != nil {
		return nil, err
	}

	return res.YouthLeagueFixtures, nil
}
