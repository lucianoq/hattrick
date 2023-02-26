package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMySeriesFixtures ...
func (a *API) GetMySeriesFixtures() (*chpp.SeriesFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(nil)
	if err != nil {
		return nil, err
	}

	return e.SeriesFixtures, nil
}

// GetMySeriesFixturesBySeason ...
func (a *API) GetMySeriesFixturesBySeason(season uint) (*chpp.SeriesFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"season": strconv.FormatUint(uint64(season), 10),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.SeriesFixtures, nil
}

// GetSeriesFixtures ...
func (a *API) GetSeriesFixtures(leagueLevelUnitID id.Series) (*chpp.SeriesFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.SeriesFixtures, nil
}

// GetSeriesFixturesBySeason ...
func (a *API) GetSeriesFixturesBySeason(leagueLevelUnitID id.Series, season uint) (*chpp.SeriesFixtures, error) {
	e, err := a.parsed.GetLeagueFixturesXML(
		map[string]string{
			"leagueLevelUnitID": leagueLevelUnitID.String(),
			"season":            strconv.FormatUint(uint64(season), 10),
		},
	)
	if err != nil {
		return nil, err
	}

	return e.SeriesFixtures, nil
}
