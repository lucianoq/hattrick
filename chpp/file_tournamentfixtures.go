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

	// 0 means this is a playoff match. Use the order of those matches to
	// display the upcoming playoff rounds: the winner of the first match
	// faces the winner of the second, and so on, with the winner of the
	// first match becoming the home team.
	Group  uint        `xml:"Group"`
	Status MatchStatus `xml:"Status"`

	// 0 until the match is finished.
	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`

	HomeStatement string `xml:"HomeStatement"`
	AwayStatement string `xml:"AwayStatement"`
}
