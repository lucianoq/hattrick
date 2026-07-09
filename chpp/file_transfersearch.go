package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TransferSearchAPIFile    = "transfersearch"
	TransferSearchAPIVersion = "1.1"
)

// TransferSearchXML contains a page of results for a transfer market search
// against the given filters.
type TransferSearchXML struct {
	Envelope
	UserID id.User `xml:"User"`

	TransferSearch TransferSearchResults `xml:"TransferSearch"`
}

// TransferSearchResults is the container for a page of transfer search
// results.
type TransferSearchResults struct {
	// The total number of items matching the search, across all pages.
	// If over 100, Hattrick returns -1 (wraps around as a very large value
	// since this field is unsigned) as an analogy to "many".
	ItemCount uint `xml:"ItemCount"`

	PageSize  uint `xml:"PageSize"`
	PageIndex uint `xml:"PageIndex"`

	Results []*TransferSearchResult `xml:"TransferResults>TransferResult"`
}

// TransferSearchResult is a single player currently listed for transfer,
// matching the search filters.
type TransferSearchResult struct {
	PlayerID id.Player `xml:"PlayerId"`

	FirstName string `xml:"FirstName"`
	NickName  string `xml:"NickName"`
	LastName  string `xml:"LastName"`

	NativeCountryID id.Country `xml:"NativeCountryID"`

	AskingPrice Money        `xml:"AskingPrice"`
	Deadline    HattrickTime `xml:"Deadline"`
	HighestBid  Money        `xml:"HighestBid"`

	BidderTeam struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`
	} `xml:"BidderTeam"`

	Details TransferSearchResultDetails `xml:"Details"`
}

// TransferSearchResultDetails holds the detailed skills/attributes of a
// player listed for transfer.
type TransferSearchResultDetails struct {
	Age     uint `xml:"Age"`
	AgeDays uint `xml:"AgeDays"`

	// Only sent since API version 1.1.
	Salary Money `xml:"Salary"`

	TSI uint `xml:"TSI"`

	Form       SkillLevel `xml:"PlayerForm"`
	Experience SkillLevel `xml:"Experience"`
	Leadership SkillLevel `xml:"Leadership"`

	Specialty SpecialtyID `xml:"Specialty"`

	Cards       uint              `xml:"Cards"`
	InjuryLevel PlayerInjuryLevel `xml:"InjuryLevel"`

	StaminaSkill   SkillLevel `xml:"StaminaSkill"`
	KeeperSkill    SkillLevel `xml:"KeeperSkill"`
	PlaymakerSkill SkillLevel `xml:"PlaymakerSkill"`
	ScorerSkill    SkillLevel `xml:"ScorerSkill"`
	PassingSkill   SkillLevel `xml:"PassingSkill"`
	WingerSkill    SkillLevel `xml:"WingerSkill"`
	DefenderSkill  SkillLevel `xml:"DefenderSkill"`
	SetPiecesSkill SkillLevel `xml:"SetPiecesSkill"`

	SellerTeam struct {
		ID       id.Team   `xml:"TeamID"`
		Name     string    `xml:"TeamName"`
		LeagueID id.League `xml:"LeagueId"`
	} `xml:"SellerTeam"`
}
