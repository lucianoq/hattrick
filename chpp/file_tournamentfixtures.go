package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TournamentFixturesAPIFile    = "tournamentfixtures"
	TournamentFixturesAPIVersion = "1.1"
)

// TournamentFixturesXML contains the list of matches for a tournament,
// for the current season (and up to 2 seasons after it finished, for
// non-restarted tournaments).
type TournamentFixturesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Matches []*TournamentFixture `xml:"Matches>Match"`
}

// TournamentFixture is a single match in a tournament's fixture list.
type TournamentFixture struct {
	MatchID id.Match `xml:"MatchId"`

	HomeTeamID        id.Team `xml:"HomeTeamId"`
	HomeTeamName      string  `xml:"HomeTeamName"`
	HomeShortTeamName string  `xml:"HomeShortTeamName"`

	AwayTeamID        id.Team `xml:"AwayTeamId"`
	AwayTeamName      string  `xml:"AwayTeamName"`
	AwayShortTeamName string  `xml:"AwayShortTeamName"`

	MatchDate  HattrickTime `xml:"MatchDate"`
	MatchType  MatchType    `xml:"MatchType"`
	MatchRound uint         `xml:"MatchRound"`
	Group      uint         `xml:"Group"`
	Status     MatchStatus  `xml:"Status"`

	// Only sent for matches that are finished.
	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`

	HomeStatement string `xml:"HomeStatement"`
	AwayStatement string `xml:"AwayStatement"`
}
