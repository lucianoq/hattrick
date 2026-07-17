package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AlliancesAPIFile    = "alliances"
	AlliancesAPIVersion = "1.4"
)

// AlliancesXML contains the (paginated) results of searching for alliances
// (federations), e.g. by name, abbreviation, description, ID, or the
// fetching user's own federations.
type AlliancesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Number of result pages.
	Pages uint `xml:"Pages"`

	// The zero-based index of this result page; each page holds at most 25
	// rows.
	PageIndex uint `xml:"PageIndex"`

	// Indicates which Supporter package the fetching user has, or empty if not
	// a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Container for the list of federations. An Attribute named Count specifies
	// how many federations it contains.
	Alliances struct {
		Alliances []*Alliance `xml:"Alliance"`
	} `xml:"Alliances"`
}
