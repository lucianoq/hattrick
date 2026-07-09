package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LadderListAPIFile    = "ladderlist"
	LadderListAPIVersion = "1.0"
)

// LadderListXML contains the list of ladders that a team currently takes
// part in.
type LadderListXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The list of ladders the given team takes part in.
	Ladders []*LadderListEntry `xml:"Ladders>Ladder"`
}

// LadderListEntry is a single ladder a team takes part in.
type LadderListEntry struct {
	ID   id.Ladder `xml:"LadderId"`
	Name string    `xml:"Name"`

	// The team's current position in the ladder. Note: the doc's own
	// element name is misspelled "Posistion" - keep it verbatim, it is not
	// a typo to fix here.
	Position uint `xml:"Posistion"`

	NextMatchDate HattrickTime `xml:"NextMatchDate"`
	Wins          uint         `xml:"Wins"`
	Lost          uint         `xml:"Lost"`
}
