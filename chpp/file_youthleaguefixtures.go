package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthLeagueFixturesAPIFile    = "youthleaguefixtures"
	YouthLeagueFixturesAPIVersion = "1.0"
)

// YouthLeagueFixturesXML contains the fixture list (past and upcoming
// matches) for a single youth league.
type YouthLeagueFixturesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	*YouthLeagueFixtures
}

// YouthLeagueFixtures is the container for a youth league's fixture list.
type YouthLeagueFixtures struct {
	YouthLeagueID   id.YouthLeague  `xml:"YouthLeagueID"`
	YouthLeagueName string          `xml:"YouthLeagueName"`
	YouthLeagueType YouthLeagueType `xml:"YouthLeagueType"`

	Season uint `xml:"Season"`

	// The last played matchround (or currently played if matches are
	// ongoing).
	LastMatchRound uint `xml:"LastMatchRound"`

	NrOfTeamsInLeague uint `xml:"NrOfTeamsInLeague"`

	Matches []*YouthLeagueFixture `xml:"Matches>Match"`
}

// YouthLeagueFixture is a single match in a youth league's fixture list.
// Deliberately distinct from Match: no ShortName fields, and HomeGoals/
// AwayGoals are only sent once the match is finished.
type YouthLeagueFixture struct {
	MatchID    id.Match    `xml:"MatchID"`
	MatchRound uint        `xml:"MatchRound"`
	Status     MatchStatus `xml:"Status"`

	HomeTeam struct {
		ID   id.Team `xml:"HomeTeamID"`
		Name string  `xml:"HomeTeamName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID   id.Team `xml:"AwayTeamID"`
		Name string  `xml:"AwayTeamName"`
	} `xml:"AwayTeam"`

	MatchDate HattrickTime `xml:"MatchDate"`

	// Only sent for matches that are finished.
	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`
}
