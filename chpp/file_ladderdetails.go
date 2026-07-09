package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LadderDetailsAPIFile    = "ladderdetails"
	LadderDetailsAPIVersion = "1.0"
)

// LadderDetailsXML contains the ranking table (teams and positions) of a
// single ladder.
type LadderDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Ladder LadderDetails `xml:"Ladder"`
}

// LadderDetails is the container for a ladder's ranking table.
type LadderDetails struct {
	ID   id.Ladder `xml:"LadderId"`
	Name string    `xml:"Name"`

	// The total number of teams in the ladder.
	NumOfTeams uint `xml:"NumOfTeams"`

	// The requested page size and index for the Team list below.
	PageSize  uint `xml:"PageSize"`
	PageIndex uint `xml:"PageIndex"`

	// The team currently on top of the ladder ("king of the hill").
	KingTeamID   id.Team      `xml:"KingTeamId"`
	KingTeamName string       `xml:"KingTeamName"`
	KingSince    HattrickTime `xml:"KingSince"`

	// The page of teams requested.
	Teams []*LadderDetailsTeam `xml:"Team"`
}

// LadderDetailsTeam is a single team's ranking entry in a ladder.
type LadderDetailsTeam struct {
	ID   id.Team `xml:"TeamId"`
	Name string  `xml:"TeamName"`

	Position uint `xml:"Position"`
	Wins     uint `xml:"Wins"`
	Lost     uint `xml:"Lost"`

	WinsInARow uint `xml:"WinsInARow"`
	LostInARow uint `xml:"LostInARow"`
}
