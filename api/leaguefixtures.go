package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyLeagueFixtures ...
func (a *API) GetMyLeagueFixtures() (*chpp.LeagueFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(nil)
	if err != nil {
		return nil, err
	}

	return e.LeagueFixtures, nil
}

// GetMyLeagueFixturesBySeason ...
func (a *API) GetMyLeagueFixturesBySeason(season uint) (*chpp.LeagueFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"season": strconv.FormatUint(uint64(season), 10),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.LeagueFixtures, nil
}

// GetLeagueFixtures ...
func (a *API) GetLeagueFixtures(leagueLevelUnitID id.LeagueLevelUnit) (*chpp.LeagueFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.LeagueFixtures, nil
}

// GetLeagueFixturesBySeason ...
func (a *API) GetLeagueFixturesBySeason(leagueLevelUnitID id.LeagueLevelUnit, season uint) (*chpp.LeagueFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
			"season":            strconv.FormatUint(uint64(season), 10),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.LeagueFixtures, nil
}
