package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TournamentDetailsAPIFile    = "tournamentdetails"
	TournamentDetailsAPIVersion = "1.0"
)

// TournamentDetailsXML contains the details of a single HTO tournament,
// for the current season only.
type TournamentDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Tournament Tournament `xml:"Tournament"`
}
