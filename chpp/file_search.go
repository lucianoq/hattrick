package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	SearchAPIFile    = "search"
	SearchAPIVersion = "1.2"
)

// SearchXML contains the results of a search against Hattrick's global
// database of players, teams, managers, arenas, leagues/regions or matches.
type SearchXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The search parameters as echoed back by the server. NOTE: the
	// CHPP doc's own field casing is inconsistent here (SearchType vs.
	// searchString/searchID/searchLeagueID) - kept verbatim.
	SearchParams struct {
		SearchType     SearchType `xml:"SearchType"`
		SearchString   string     `xml:"searchString"`
		SearchString2  string     `xml:"searchString2"`
		SearchID       uint       `xml:"searchID"`
		SearchLeagueID int        `xml:"searchLeagueID"`
	} `xml:"SearchParams"`

	// The search results. Each Result is one row of the resultset.
	Results []*SearchResult `xml:"Result"`
}

// SearchResult is a single row of a search resultset. The meaning of ID
// and Name depends on the requested SearchType (player, arena, manager,
// series, team, region or match).
type SearchResult struct {
	ID   uint   `xml:"ResultID"`
	Name string `xml:"ResultName"`
}
