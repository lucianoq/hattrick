package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchesArchiveAPIFile    = "matchesArchive"
	MatchesArchiveAPIVersion = "1.4"
)

// MatchesArchiveXML ...
type MatchesArchiveXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	IsYouth bool `xml:"IsYouth"`
	Team    struct {
		TeamID        id.Team      `xml:"TeamID"`
		TeamName      string       `xml:"TeamName"`
		FistMatchDate HattrickTime `xml:"FirstMatchDate"`
		LastMatchDate HattrickTime `xml:"LastMatchDate"`
		MatchList     struct {
			Matches []*Match `xml:"Match"`
		} `xml:"MatchList"`
	} `xml:"Team"`
}
