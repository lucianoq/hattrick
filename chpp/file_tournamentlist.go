package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TournamentListAPIFile    = "tournamentlist"
	TournamentListAPIVersion = "1.0"
)

// TournamentListXML contains the list of tournaments a team is currently
// playing in.
type TournamentListXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The list of tournaments the given team takes part in.
	Tournaments []*Tournament `xml:"Tournaments>Tournament"`
}

// Tournament is the full information about a single HTO tournament, as
// returned both by tournamentlist and tournamentdetails.
type Tournament struct {
	ID   id.Tournament  `xml:"TournamentId"`
	Name string         `xml:"Name"`
	Type TournamentType `xml:"TournamentType"`

	Season uint `xml:"Season"`

	LogoURL string `xml:"LogoUrl"`

	// The type of trophy awarded, if any.
	TrophyType uint `xml:"TrophyType"`

	NumberOfTeams  uint `xml:"NumberOfTeams"`
	NumberOfGroups uint `xml:"NumberOfGroups"`

	LastMatchRound      uint         `xml:"LastMatchRound"`
	FirstMatchRoundDate HattrickTime `xml:"FirstMatchRoundDate"`
	NextMatchRoundDate  HattrickTime `xml:"NextMatchRoundDate"`
	IsMatchesOngoing    bool         `xml:"IsMatchesOngoing"`

	// The user who created the tournament.
	Creator struct {
		UserID    id.User `xml:"UserId"`
		LoginName string  `xml:"Loginname"`
	} `xml:"Creator"`
}
