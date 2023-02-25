package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AllianceDetailsAPIFile    = "alliancedetails"
	AllianceDetailsAPIVersion = "1.5"
)

// AllianceDetailsXML ...
type AllianceDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// Indicates which Supporter package the fetching user has, or empty if not
	// a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// The requested ActionType.
	ActionType string `xml:"ActionType"`

	// Container for the data of the federation.
	Alliance *Alliance `xml:"Alliance"`
}

// Alliance ...
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

	// Container for the list of languages the federation has declared As used
	// languages. An Attribute named Count
	// specifies how many federations it contains.
	Languages []struct {
		ID   uint   `xml:"LanguageID"`
		Name string `xml:"LanguageName"`
	} `xml:"Languages>Language"`
	Message string `xml:"Message"`

	// Internal message to members of federation. Only present if indata
	// ShowRules is 'true'.
	Rules string `xml:"Rules"`

	// Container for role of fetching user in federation.
	UserRole struct {
		RoleID   uint   `xml:"RoleId"`
		RoleName string `xml:"RoleName"`
	} `xml:"UserRole"`

	// Container for list of roles in federation. An Attribute named Count
	// specifies how many federations it contains.
	Roles struct {
		// Container for a role. An Attribute named Index works as a counter.
		Roles []*AllianceRole `xml:"Role"`
	} `xml:"Roles"`

	ListSubset string `xml:"ListSubset"`

	// Container for the list of members, or a subset thereof, of the federation
	Members []*AllianceMember `xml:"Members>Member"`
}

// AllianceRole ...
type AllianceRole struct {
	ID          id.Role `xml:"RoleId"`
	Name        string  `xml:"RoleName"`
	Rank        uint    `xml:"RoleRank"`
	MemberCount uint    `xml:"RoleMemberCount"`
	MaxMember   uint    `xml:"RoleMaxMember"`
	RequestType uint    `xml:"RoleRequestType"`
	Description string  `xml:"RoleDescription"`
}

// AllianceMember ...
type AllianceMember struct {
	UserID         id.User      `xml:"User"`
	LoginName      string       `xml:"Loginname"`
	RoleID         uint         `xml:"RoleId"`
	RoleName       string       `xml:"RoleName"`
	IsOnline       bool         `xml:"IsOnline"`
	MembershipDate HattrickTime `xml:"MemberShipDate"`
}
