package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchesArchiveAPIFile    = "matchesArchive"
	MatchesArchiveAPIVersion = "1.5"
)

// MatchesArchiveXML contains a team's past matches within a date range
// (or season), most recent 50 matches only.
type MatchesArchiveXML struct {
	Envelope
	UserID id.User `xml:"User"`

	IsYouth bool `xml:"IsYouth"`
	Team    struct {
		TeamID id.Team `xml:"TeamID"`

		// The full name of the team.
		TeamName string `xml:"TeamName"`

		// The oldest date to show matches in the archive from.
		FirstMatchDate HattrickTime `xml:"FirstMatchDate"`

		// The latest date to show matches in the archive up to.
		LastMatchDate HattrickTime `xml:"LastMatchDate"`

		MatchList struct {
			// If more than 50 matches occurred between FirstMatchDate and
			// LastMatchDate, only the first 50 are returned.
			Matches []*Match `xml:"Match"`
		} `xml:"MatchList"`
	} `xml:"Team"`
}
