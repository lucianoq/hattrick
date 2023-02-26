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
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	*Series
}

// Series ...
type Series struct {
	ID                id.Series `xml:"LeagueLevelUnitID"`
	Name              string    `xml:"LeagueLevelUnitName"`
	LeagueID          id.League `xml:"LeagueID"`
	LeagueName        string    `xml:"LeagueName"`
	LeagueLevel       uint      `xml:"LeagueLevel"`
	MaxLevel          uint      `xml:"MaxLevel"`
	CurrentMatchRound string    `xml:"CurrentMatchRound"`
	Rank              uint      `xml:"Rank"`
	Teams             []*struct {
		UserID         id.User        `xml:"UserId"`
		ID             id.Team        `xml:"TeamID"`
		Name           string         `xml:"TeamName"`
		Position       uint           `xml:"Position"`
		PositionChange PositionChange `xml:"PositionChange"`
		Matches        uint           `xml:"Matches"`
		GoalsFor       uint           `xml:"GoalsFor"`
		GoalsAgainst   uint           `xml:"GoalsAgainst"`
		Points         uint           `xml:"Points"`
		Won            uint           `xml:"Won"`
		Draws          uint           `xml:"Draws"`
		Lost           uint           `xml:"Lost"`
	} `xml:"Team"`
}
