package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	HOFPlayersAPIFile    = "hofplayers"
	HOFPlayersAPIVersion = "1.2"
)

// HOFPlayersXML contains the players a team has inducted into its Hall of
// Fame.
type HOFPlayersXML struct {
	Envelope
	UserID id.User `xml:"UserID"`

	PlayerList struct {
		Players []*HOFPlayer `xml:"Player"`
	} `xml:"PlayerList"`
}

// HOFPlayer is a single player inducted into a team's Hall of Fame.
type HOFPlayer struct {
	PlayerID  id.Player `xml:"PlayerID"`
	FirstName string    `xml:"FirstName"`
	NickName  string    `xml:"NickName"`
	LastName  string    `xml:"LastName"`
	Age       uint      `xml:"Age"`

	// The approximate date/time of the player's next birthday.
	NextBirthday HattrickTime `xml:"NextBirthday"`

	CountryID id.Country `xml:"CountryID"`

	// The date the player arrived at the team.
	ArrivalDate HattrickTime `xml:"ArrivalDate"`

	// The type of "life after football" job the player now holds.
	ExpertType HOFExpertType `xml:"ExpertType"`

	// The date the player was inducted into the Hall of Fame.
	HofDate HattrickTime `xml:"HofDate"`

	// The player's age, in years, when inducted into the Hall of Fame.
	HofAge uint `xml:"HofAge"`
}
