package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LeagueFixturesAPIFile    = "leaguefixtures"
	LeagueFixturesAPIVersion = "1.2"
)

// LeagueFixturesXML ...
type LeagueFixturesXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	*LeagueFixtures
}

// LeagueFixtures ...
type LeagueFixtures struct {
	LeagueLevelUnitID   id.LeagueLevelUnit     `xml:"LeagueLevelUnitID"`
	LeagueLevelUnitName string                 `xml:"LeagueLevelUnitName"`
	Season              uint                   `xml:"Season"`
	Matches             []*LeagueFixturesMatch `xml:"Match"`
}

// LeagueFixturesMatch ...
type LeagueFixturesMatch struct {
	// The globally unique match identifier.
	MatchID id.Match `xml:"MatchID"`

	// MatchRound that the match is part of.
	MatchRound uint `xml:"MatchRound"`

	// The home team container.
	HomeTeam struct {
		HomeTeamID   id.Team `xml:"HomeTeamID"`
		HomeTeamName string  `xml:"HomeTeamName"`
	} `xml:"HomeTeam"`

	// The away team container.
	AwayTeam struct {
		AwayTeamID   id.Team `xml:"AwayTeamID"`
		AwayTeamName string  `xml:"AwayTeamName"`
	} `xml:"AwayTeam"`

	// The date indicating when kick-off takes/took place.
	MatchDate HattrickTime `xml:"MatchDate"`

	// The final number of goals that the home team scored. This parameter is
	// only sent for matches that are finished.
	HomeGoals uint `xml:"HomeGoals"`

	// The final number of goals that the away team scored. This parameter is
	// only sent for matches that are finished.
	AwayGoals uint `xml:"AwayGoals"`
}
