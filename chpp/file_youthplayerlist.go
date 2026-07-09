package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthPlayerListAPIFile    = "youthplayerlist"
	YouthPlayerListAPIVersion = "1.3"
)

// YouthPlayerListXML contains the list of players belonging to a youth
// team.
type YouthPlayerListXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for all players. For the default actionType=list, only
	// the id and name fields (and Gender) are populated; the remaining
	// fields on YouthPlayerDetail are left at their zero value. For
	// actionType=details/unlockskills, the full player detail is
	// populated.
	Players []*YouthPlayerDetail `xml:"PlayerList>YouthPlayer"`
}
