package api

import (
	"errors"
	"time"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyMatchesArchive ...
func (a *API) GetMyMatchesArchive(start, end time.Time) ([]*chpp.Match, error) {
	if start.After(end) {
		return nil, errors.New("start is after the end date")
	}

	matches := make([]*chpp.Match, 0, 50)

	for _, intv := range createBins(start, end) {
		values := map[string]string{}
		values["isYouth"] = "false"
		values["FirstMatchDate"] = chpp.FromTime(intv.from).String()
		values["LastMatchDate"] = chpp.FromTime(intv.to).String()

		res, err := a.parsed.GetMatchesArchiveXML(values)
		if err != nil {
			return nil, err
		}

		matches = append(matches, res.Team.MatchList.Matches...)
	}

	return matches, nil
}

// GetMatchesArchive ...
func (a *API) GetMatchesArchive(teamID id.Team, start, end time.Time) ([]*chpp.Match, error) {
	if start.After(end) {
		return nil, errors.New("start is after the end date")
	}

	matches := make([]*chpp.Match, 0, 50)

	for _, intv := range createBins(start, end) {
		values := map[string]string{}
		values["teamID"] = teamID.String()
		values["isYouth"] = "false"
		values["FirstMatchDate"] = chpp.FromTime(intv.from).String()
		values["LastMatchDate"] = chpp.FromTime(intv.to).String()

		res, err := a.parsed.GetMatchesArchiveXML(values)
		if err != nil {
			return nil, err
		}

		matches = append(matches, res.Team.MatchList.Matches...)
	}

	return matches, nil
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
//   firstMatchDate = DateTime.Now.AddMonths(-3).Date
//   lastMatchDate = DateTime.Now.Date
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
