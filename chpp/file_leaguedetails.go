package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LeagueDetailsAPIFile    = "leaguedetails"
	LeagueDetailsAPIVersion = "1.6"
)

// LeagueDetailsXML contains data about a League Level Unit (series)
type LeagueDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	*Series
}

// Series is a League Level Unit: one table/division series within a league,
// including its standings.
type Series struct {
	ID   id.Series `xml:"LeagueLevelUnitID"`
	Name string    `xml:"LeagueLevelUnitName"`

	LeagueID   id.League `xml:"LeagueID"`
	LeagueName string    `xml:"LeagueName"`

	// 1 = top division, 2 = second division, and so on.
	LeagueLevel uint `xml:"LeagueLevel"`

	// The total number of divisions/levels in this league.
	MaxLevel uint `xml:"MaxLevel"`

	CurrentMatchRound string `xml:"CurrentMatchRound"`

	// The series' ranking relative to other series on the same level.
	Rank  uint `xml:"Rank"`
	Teams []*struct {
		UserID id.User `xml:"UserId"`
		ID     id.Team `xml:"TeamID"`
		Name   string  `xml:"TeamName"`

		// The team's current position in the table.
		Position uint `xml:"Position"`

		// How the team's position has changed since the last round.
		PositionChange PositionChange `xml:"PositionChange"`

		Matches      uint `xml:"Matches"`
		GoalsFor     uint `xml:"GoalsFor"`
		GoalsAgainst uint `xml:"GoalsAgainst"`
		Points       uint `xml:"Points"`
		Won          uint `xml:"Won"`
		Draws        uint `xml:"Draws"`
		Lost         uint `xml:"Lost"`
	} `xml:"Team"`
}
