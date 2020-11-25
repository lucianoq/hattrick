package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	HOFPlayersAPIFile    = "hofplayers"
	HOFPlayersAPIVersion = "1.2"
)

// HOFPlayersXML ...
type HOFPlayersXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"UserID"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	PlayerList struct {
		Players []*HOFPlayer `xml:"Player"`
	} `xml:"PlayerList"`
}

// HOFPlayer ...
type HOFPlayer struct {
	PlayerID     id.Player     `xml:"PlayerId"`
	FirstName    string        `xml:"FirstName"`
	NickName     string        `xml:"NickName"`
	LastName     string        `xml:"LastName"`
	Age          uint          `xml:"Age"`
	NextBirthday HattrickTime  `xml:"NextBirthday"`
	CountryID    id.Country    `xml:"CountryID"`
	ArrivalDate  HattrickTime  `xml:"ArrivalDate"`
	ExpertType   HOFExpertType `xml:"ExpertType"`
	HofDate      HattrickTime  `xml:"HofDate"`
	HofAge       uint          `xml:"HofAge"`
}
