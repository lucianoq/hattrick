package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AllianceDetailsAPIFile    = "alliancedetails"
	AllianceDetailsAPIVersion = "1.5"
)

// AllianceDetailsXML contains detailed information about a single
// alliance (federation), such as its roles or members, depending on the
// requested ActionType.
type AllianceDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Whether the fetching user is a Supporter.
	UserIsSupporter bool `xml:"UserIsSupporter"`

	// Indicates which Supporter package the fetching user has, or empty if not
	// a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Whether the fetching user has a clubhouse.
	UserHasClubhouse bool `xml:"UserHasClubhouse"`

	// The requested ActionType.
	ActionType string `xml:"ActionType"`

	// Container for the data of the federation.
	Alliance *Alliance `xml:"Alliance"`
}

// Alliance is a federation of teams, i.e. a Hattrick "alliance"/federation
// and its details (which fields are populated depends on the ActionType
// requested).
type Alliance struct {
	AllianceID          id.Alliance  `xml:"AllianceID"`
	AllianceName        string       `xml:"AllianceName"`
	AllianceDescription string       `xml:"AllianceDescription"`
	Abbreviation        string       `xml:"Abbreviation"`
	Description         string       `xml:"Description"`
	LogoURL             string       `xml:"LogoURL"`
	TopRole             string       `xml:"TopRole"`
	TopUserID           id.User      `xml:"TopUserID"`
	TopLoginName        string       `xml:"TopLoginname"`
	CreationDate        HattrickTime `xml:"CreationDate"`
	HomePageURL         string       `xml:"HomePageURL"`
	NumberOfMembers     uint         `xml:"NumberOfMembers"`
	AwaitingRequests    uint         `xml:"AwaitingRequests"`

	// Container for the list of languages the federation has declared as used
	// languages. An Attribute named Count
	// specifies how many languages it contains.
	Languages []struct {
		ID   uint   `xml:"LanguageID"`
		Name string `xml:"LanguageName"`
	} `xml:"Languages>Language"`
	Message string `xml:"Message"`

	// Internal message to members of the federation. Only present if the
	// input data ShowRules is 'true'.
	Rules string `xml:"Rules"`

	// Container for the role of the fetching user in the federation.
	UserRole struct {
		RoleID   uint   `xml:"RoleId"`
		RoleName string `xml:"RoleName"`
	} `xml:"UserRole"`

	// Container for the list of roles in the federation. An Attribute named Count
	// specifies how many roles it contains.
	Roles struct {
		// Container for a role. An Attribute named Index works as a counter.
		Roles []*AllianceRole `xml:"Role"`
	} `xml:"Roles"`

	ListSubset string `xml:"ListSubset"`

	// Container for the list of members, or a subset thereof, of the federation
	Members []*AllianceMember `xml:"Members>Member"`
}

// AllianceRole is a role that can be held by members of an alliance
// (federation), as returned for ActionType=roles.
type AllianceRole struct {
	ID          id.Role `xml:"RoleId"`
	Name        string  `xml:"RoleName"`
	Rank        uint    `xml:"RoleRank"`
	MemberCount uint    `xml:"RoleMemberCount"`
	MaxMember   uint    `xml:"RoleMaxMembers"`

	// How to become a member of this role: 1 = apply, 2 = free for all.
	RequestType uint   `xml:"RoleRequestType"`
	Description string `xml:"RoleDescription"`
}

// AllianceMember is a single member of an alliance (federation), as
// returned for ActionType=members or ActionType=membersSubset.
type AllianceMember struct {
	UserID    id.User `xml:"UserID"`
	LoginName string  `xml:"Loginname"`
	RoleID    uint    `xml:"RoleId"`
	RoleName  string  `xml:"RoleName"`
	IsOnline  bool    `xml:"IsOnline"`
	// The date when the member last changed role within the federation, not
	// the date they joined.
	MembershipDate HattrickTime `xml:"MembershipDate"`
}
