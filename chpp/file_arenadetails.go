package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ArenaDetailsAPIFile    = "arenadetails"
	ArenaDetailsAPIVersion = "1.5"
)

// ArenaDetailsXML contains Information about specific arenas, supporter
// statistics and the largest arenas in Hattrick.
type ArenaDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// Container for the data about the arena.
	Arena *Arena `xml:"Arena"`
}

// Arena is a container for the data about the arena.
type Arena struct {
	ArenaID   id.Arena `xml:"ArenaID"`
	ArenaName string   `xml:"ArenaName"`

	// Container for the data about the team owning this arena.
	Team struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`
	} `xml:"Team"`

	// Container for the data about the league of the arena's team.
	League struct {
		ID   id.League `xml:"LeagueID"`
		Name string    `xml:"LeagueName"`
	} `xml:"League"`

	// Container for the data about the region of the arena.
	Region struct {
		ID   id.Region `xml:"RegionID"`
		Name string    `xml:"RegionName"`
	} `xml:"Region"`

	// Container for the data about the current capacity of the arena.
	CurrentCapacity struct {
		RebuiltDate HattrickTime `xml:"RebuiltDate"`
		Terraces    uint         `xml:"Terraces"`
		Basic       uint         `xml:"Basic"`
		Roof        uint         `xml:"Roof"`
		VIP         uint         `xml:"VIP"`
		Total       uint         `xml:"Total"`
	} `xml:"CurrentCapacity"`

	// Container for the data about the expanded capacity of the arena during
	// construction. If the arena is under construction an attribute named
	// Available is set to true, otherwise to false and the container is then
	// empty.
	ExpandedCapacity struct {
		Available     bool         `xml:"Available,attr"`
		ExpansionDate HattrickTime `xml:"ExpansionDate"`
		Terraces      uint         `xml:"Terraces"`
		Basic         uint         `xml:"Basic"`
		Roof          uint         `xml:"Roof"`
		VIP           uint         `xml:"VIP"`
		Total         uint         `xml:"Total"`
	} `xml:"ExpandedCapacity"`
}
