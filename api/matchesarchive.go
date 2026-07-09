package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyMatchesArchive returns the requesting user's senior team's archived
// matches between start and end.
func (a *API) GetMyMatchesArchive(start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(nil, start, end, false, false)
}

// GetMatchesArchive returns the given senior team's archived matches
// between start and end.
func (a *API) GetMatchesArchive(teamID id.Team, start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(&teamID, start, end, false, false)
}

// GetMyYouthMatchesArchive returns the requesting user's youth team's
// archived matches between start and end.
func (a *API) GetMyYouthMatchesArchive(start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(nil, start, end, true, false)
}

// GetYouthMatchesArchive returns the given youth team's archived matches
// between start and end.
func (a *API) GetYouthMatchesArchive(youthTeamID id.Team, start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(&youthTeamID, start, end, true, false)
}

// GetMyMatchesArchiveIncludeHTO is like GetMyMatchesArchive, but also
// includes HTO matches (tournaments, ladders, preparation and single
// matches).
func (a *API) GetMyMatchesArchiveIncludeHTO(start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(nil, start, end, false, true)
}

// GetMatchesArchiveIncludeHTO is like GetMatchesArchive, but also includes
// HTO matches (tournaments, ladders, preparation and single matches).
func (a *API) GetMatchesArchiveIncludeHTO(teamID id.Team, start, end time.Time) ([]*chpp.Match, error) {
	return a.matchesArchive(&teamID, start, end, false, true)
}

func (a *API) matchesArchive(teamID *id.Team, start, end time.Time, isYouth, includeHTO bool) ([]*chpp.Match, error) {
	if start.After(end) {
		return nil, errors.New("start is after the end date")
	}

	matches := make([]*chpp.Match, 0, 50)

	for _, intv := range createBins(start, end) {
		values := map[string]string{
			"isYouth":        strconv.FormatBool(isYouth),
			"includeHTO":     strconv.FormatBool(includeHTO),
			"FirstMatchDate": chpp.FromTime(intv.from).String(),
			"LastMatchDate":  chpp.FromTime(intv.to).String(),
		}
		if teamID != nil {
			values["teamID"] = teamID.String()
		}

		res, err := a.parsed.GetMatchesArchiveXML(values)
		if err != nil {
			return nil, err
		}

		matches = append(matches, res.Team.MatchList.Matches...)
	}

	return matches, nil
}

// GetMyMatchesArchiveBySeason returns the requesting user's senior team's
// archived matches for the given season. Only valid for senior teams, not
// youth. Overrides any date range. Per the CHPP doc, if more than 50
// matches occurred in the season, only the first 50 are returned.
func (a *API) GetMyMatchesArchiveBySeason(season uint) ([]*chpp.Match, error) {
	return a.matchesArchiveBySeason(nil, season)
}

// GetMatchesArchiveBySeason returns the given senior team's archived matches
// for the given season. Only valid for senior teams, not youth. Overrides
// any date range. Per the CHPP doc, if more than 50 matches occurred in the
// season, only the first 50 are returned.
func (a *API) GetMatchesArchiveBySeason(teamID id.Team, season uint) ([]*chpp.Match, error) {
	return a.matchesArchiveBySeason(&teamID, season)
}

func (a *API) matchesArchiveBySeason(teamID *id.Team, season uint) ([]*chpp.Match, error) {
	values := map[string]string{
		"isYouth": "false",
		"season":  strconv.FormatUint(uint64(season), 10),
	}
	if teamID != nil {
		values["teamID"] = teamID.String()
	}

	res, err := a.parsed.GetMatchesArchiveXML(values)
	if err != nil {
		return nil, err
	}

	return res.Team.MatchList.Matches, nil
}

type interval struct {
	from time.Time
	to   time.Time
}

// * If more than 50 matches have occurred between FirstMatchDate and
// LastMatchDate, only the 50 first will be returned.
// * For performance reasons you may only specify an interval of 2 seasons back in
// time. If you specify a larger interval we'll automatically adjust it to the
// default which is:
//
//	firstMatchDate = DateTime.Now.AddMonths(-3).Date
//	lastMatchDate = DateTime.Now.Date
const binMaxSize = 50 * 24 * time.Hour // 50 days

func createBins(start, end time.Time) []interval {
	intervals := make([]interval, 0, end.Sub(start)/binMaxSize+1)

	left, right := start, start.Add(binMaxSize)
	for right.Before(end) {
		intervals = append(intervals, interval{left, right})
		left, right = right.Add(time.Nanosecond), right.Add(binMaxSize)
	}

	// last bin
	intervals = append(intervals, interval{left, end})

	return intervals
}
