package api

import (
	"strconv"
	"time"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyMatches returns the requesting user's senior team's upcoming and
// recently played matches.
func (a *API) GetMyMatches() ([]*chpp.Match, error) {
	return a.matches(nil, false, time.Time{})
}

// GetMyYouthMatches returns the requesting user's youth team's upcoming
// and recently played matches.
func (a *API) GetMyYouthMatches() ([]*chpp.Match, error) {
	return a.matches(nil, true, time.Time{})
}

// GetMatches returns the upcoming and recently played matches for the given
// senior team.
func (a *API) GetMatches(teamID id.Team) ([]*chpp.Match, error) {
	return a.matches(&teamID, false, time.Time{})
}

// GetYouthMatches returns the upcoming and recently played matches for the
// given youth team.
func (a *API) GetYouthMatches(teamID id.Team) ([]*chpp.Match, error) {
	return a.matches(&teamID, true, time.Time{})
}

// GetMyMatchesUntil is like GetMyMatches, but only returns matches up to
// lastMatchDate (if more than 50 matches are affected, only the first 50
// are returned).
func (a *API) GetMyMatchesUntil(lastMatchDate time.Time) ([]*chpp.Match, error) {
	return a.matches(nil, false, lastMatchDate)
}

// GetMyYouthMatchesUntil is like GetMyYouthMatches, but only returns
// matches up to lastMatchDate (if more than 50 matches are affected, only
// the first 50 are returned).
func (a *API) GetMyYouthMatchesUntil(lastMatchDate time.Time) ([]*chpp.Match, error) {
	return a.matches(nil, true, lastMatchDate)
}

// GetMatchesUntil is like GetMatches, but only returns matches up to
// lastMatchDate (if more than 50 matches are affected, only the first 50
// are returned).
func (a *API) GetMatchesUntil(teamID id.Team, lastMatchDate time.Time) ([]*chpp.Match, error) {
	return a.matches(&teamID, false, lastMatchDate)
}

// GetYouthMatchesUntil is like GetYouthMatches, but only returns matches up
// to lastMatchDate (if more than 50 matches are affected, only the first 50
// are returned).
func (a *API) GetYouthMatchesUntil(teamID id.Team, lastMatchDate time.Time) ([]*chpp.Match, error) {
	return a.matches(&teamID, true, lastMatchDate)
}

func (a *API) matches(teamID *id.Team, isYouth bool, lastMatchDate time.Time) ([]*chpp.Match, error) {
	values := map[string]string{
		"isYouth": strconv.FormatBool(isYouth),
	}
	if teamID != nil {
		values["teamID"] = teamID.String()
	}
	if !lastMatchDate.IsZero() {
		values["LastMatchDate"] = chpp.FromTime(lastMatchDate).String()
	}

	res, err := a.parsed.GetMatchesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Team.Matches, nil
}
