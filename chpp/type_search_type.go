package chpp

// SearchType is the kind of entity a Hattrick search query looks for
// (players, teams, arenas, managers, etc).
type SearchType uint

// List of SearchType constants.
const (
	SearchTypePlayers  SearchType = 0
	SearchTypeArenas   SearchType = 1
	SearchTypeManagers SearchType = 2
	SearchTypeSeries   SearchType = 3
	SearchTypeTeams    SearchType = 4
	SearchTypeRegions  SearchType = 5
	SearchTypeMatch    SearchType = 6
)
