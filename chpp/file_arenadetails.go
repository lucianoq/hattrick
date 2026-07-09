package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ArenaDetailsAPIFile    = "arenadetails"
	ArenaDetailsAPIVersion = "1.7"
)

// ArenaDetailsXML contains information about specific arenas, supporter
// statistics and the largest arenas in Hattrick.
type ArenaDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for the data about the arena. Only provided for the
	// default StatsType.
	Arena *Arena `xml:"Arena"`

	// Container for the biggest arena statistics. Only provided for
	// StatsType=OtherArenas.
	LeagueArenaStats *LeagueArenaStats `xml:"LeagueArenaStats"`

	// Container for supporter statistics for the user's own arena. Only
	// provided for StatsType=MyArena.
	MyArena *MyArenaStats `xml:"MyArena"`
}

// MyArenaStats is the container for supporter statistics for the requesting
// user's own arena, as returned when StatsType=MyArena.
type MyArenaStats struct {
	ID            id.Arena `xml:"ArenaID"`
	Name          string   `xml:"ArenaName"`
	Image         string   `xml:"ArenaImage"`
	FallbackImage string   `xml:"ArenaFallbackImage"`

	// The type of matches included in the statistics.
	MatchTypes ArenaMatchType `xml:"MatchTypes"`

	// The period the statistics cover.
	FirstDate HattrickTime `xml:"FirstDate"`
	LastDate  HattrickTime `xml:"LastDate"`

	// Number of matches of the specified type within the given period.
	NumberOfMatches uint `xml:"NumberOfMatches"`

	VisitorsAverage seats `xml:"VisitorsAverage"`
	VisitorsMost    seats `xml:"VisitorsMost"`
	VisitorsLeast   seats `xml:"VisitorsLeast"`
}

// LeagueArenaStats is the container for the biggest arena statistics for a
// league (or globally, if LeagueID is 0).
type LeagueArenaStats struct {
	// The globally unique LeagueID. 0 for global statistics.
	LeagueID id.League `xml:"LeagueID"`

	// The name of the league. Not provided for global statistics.
	LeagueName string `xml:"LeagueName"`

	// The date when the statistics were created/updated.
	CreatedDate HattrickTime `xml:"CreatedDate"`

	// Container for biggest arena statistics for the league.
	Stats []*ArenaStat `xml:"OtherArenasStatList>ArenaStat"`
}

// ArenaStat is a particular arena's statistics, as returned when
// StatsType=OtherArenas.
type ArenaStat struct {
	ID            id.Arena `xml:"ArenaID"`
	Name          string   `xml:"ArenaName"`
	Image         string   `xml:"ArenaImage"`
	FallbackImage string   `xml:"ArenaFallbackImage"`

	// The size of the arena.
	Size uint `xml:"ArenaSize"`

	// The league for where the arena is used.
	LeagueID   id.League `xml:"ArenaLeagueID"`
	LeagueName string    `xml:"ArenaLeagueName"`

	// The region for where the arena is placed.
	RegionID   id.Region `xml:"ArenaRegionID"`
	RegionName string    `xml:"ArenaRegionName"`
}

// Arena is a container for the data about the arena.
type Arena struct {
	ID            id.Arena `xml:"ArenaID"`
	Name          string   `xml:"ArenaName"`
	Image         string   `xml:"ArenaImage"`
	FallbackImage string   `xml:"ArenaFallbackImage"`

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

	// Container for the data about the current capacity of the arena. If
	// the arena is under construction, Available is false and the rest
	// of the container is empty.
	CurrentCapacity struct {
		Available   bool         `xml:"Available,attr"`
		RebuiltDate HattrickTime `xml:"RebuiltDate"`
		seats
	} `xml:"CurrentCapacity"`

	// Container for the data about the expanded capacity of the arena during
	// construction. If the arena is under construction an attribute named
	// Available is set to true, otherwise to false and the container is then
	// empty.
	ExpandedCapacity struct {
		Available     bool         `xml:"Available,attr"`
		ExpansionDate HattrickTime `xml:"ExpansionDate"`
		seats
	} `xml:"ExpandedCapacity"`
}

type seats struct {
	Terraces uint `xml:"Terraces"`
	Basic    uint `xml:"Basic"`
	Roof     uint `xml:"Roof"`
	VIP      uint `xml:"VIP"`
	Total    uint `xml:"Total"`
}
