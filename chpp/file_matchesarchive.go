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
		TeamID         id.Team      `xml:"TeamID"`
		TeamName       string       `xml:"TeamName"`
		FirstMatchDate HattrickTime `xml:"FirstMatchDate"`
		LastMatchDate  HattrickTime `xml:"LastMatchDate"`
		MatchList      struct {
			Matches []*Match `xml:"Match"`
		} `xml:"MatchList"`
	} `xml:"Team"`
}
