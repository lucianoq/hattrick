package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LeagueFixturesAPIFile    = "leaguefixtures"
	LeagueFixturesAPIVersion = "1.2"
)

// LeagueFixturesXML contains the match fixtures for a League Level Unit
// (series).
type LeagueFixturesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	*SeriesFixtures
}

// SeriesFixtures is the full match schedule of a League Level Unit (series)
// for a given season.
type SeriesFixtures struct {
	SeriesID   id.Series `xml:"LeagueLevelUnitID"`
	SeriesName string    `xml:"LeagueLevelUnitName"`

	// The season the returned fixtures belong to.
	Season  uint           `xml:"Season"`
	Matches []*SeriesMatch `xml:"Match"`
}

// SeriesMatch is a single league match within a series' fixture list.
type SeriesMatch struct {
	// The globally unique match identifier.
	ID id.Match `xml:"MatchID"`

	// MatchRound that the match is part of.
	Round uint `xml:"MatchRound"`

	// The home team container.
	HomeTeam struct {
		ID   id.Team `xml:"HomeTeamID"`
		Name string  `xml:"HomeTeamName"`
	} `xml:"HomeTeam"`

	// The away team container.
	AwayTeam struct {
		ID   id.Team `xml:"AwayTeamID"`
		Name string  `xml:"AwayTeamName"`
	} `xml:"AwayTeam"`

	// The date indicating when kick-off takes/took place.
	Date HattrickTime `xml:"MatchDate"`

	// The final number of goals that the home team scored. This parameter is
	// only sent for matches that are finished.
	HomeGoals uint `xml:"HomeGoals"`

	// The final number of goals that the away team scored. This parameter is
	// only sent for matches that are finished.
	AwayGoals uint `xml:"AwayGoals"`
}
