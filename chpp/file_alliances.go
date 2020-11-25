package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AlliancesAPIFile    = "alliances"
	AlliancesAPIVersion = "1.4"
)

// AlliancesXML ...
type AlliancesXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// Number of result pages.
	Pages uint `xml:"Pages"`

	// Indicates which Supporter package the fetching user has, or empty if not
	// a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Container for the list of federation. An Attribute named Count specifies
	// how many federations it contains.
	Alliances struct {
		Alliances []*Alliance
	} `xml:"Alliances"`
}
