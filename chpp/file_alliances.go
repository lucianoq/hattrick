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

	// The current page index of the result.
	PageIndex uint `xml:"PageIndex"`

	// Whether the fetching user is a Supporter.
	UserIsSupporter bool `xml:"UserIsSupporter"`

	// Indicates which Supporter package the fetching user has, or empty if not
	// a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Whether the fetching user has a clubhouse.
	UserHasClubHouse bool `xml:"UserHasClubHouse"`

	// Container for the list of federations. An Attribute named Count specifies
	// how many federations it contains.
	Alliances struct {
		Alliances []*Alliance `xml:"Alliance"`
	} `xml:"Alliances"`
}
