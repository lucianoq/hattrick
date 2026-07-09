package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthLeagueDetailsAPIFile    = "youthleaguedetails"
	YouthLeagueDetailsAPIVersion = "1.1"
)

// YouthLeagueDetailsXML contains the league table for a single youth
// league.
type YouthLeagueDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	YouthLeagueDetails
}

// YouthLeagueDetails describes a youth league (series) and its current
// league table.
type YouthLeagueDetails struct {
	// The ID number of the Youth League (series).
	ID id.YouthLeague `xml:"YouthLeagueID"`

	// The name of the Youth League.
	Name string `xml:"YouthLeagueName"`

	// The type of the league: regional, national or international.
	Type YouthLeagueType `xml:"YouthLeagueType"`

	// The season of the league.
	Season uint `xml:"Season"`

	// The last match round played for the league.
	LastMatchRound uint `xml:"LastMatchRound"`

	// Number of teams in the league.
	NrOfTeamsInLeague uint `xml:"NrOfTeamsInLeague"`

	// The system type of the league.
	LeagueSystemID LeagueSystemID `xml:"LeagueSystemID"`

	// Container for the teams.
	Teams []*YouthLeagueTeam `xml:"Teams>Team"`
}

// YouthLeagueTeam is a team in a youth league table.
type YouthLeagueTeam struct {
	ID   id.YouthTeam `xml:"TeamID"`
	Name string       `xml:"TeamName"`

	// The position in the league.
	Position uint `xml:"Position"`

	// How the team is moving in the league.
	PositionChange PositionChange `xml:"PositionChange"`

	Matches      uint `xml:"Matches"`
	GoalsFor     uint `xml:"GoalsFor"`
	GoalsAgainst uint `xml:"GoalsAgainst"`
	Points       uint `xml:"Points"`
	Won          uint `xml:"Won"`
	Draws        uint `xml:"Draws"`
	Lost         uint `xml:"Lost"`
}
