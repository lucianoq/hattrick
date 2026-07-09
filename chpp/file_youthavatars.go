package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthAvatarsAPIFile    = "youthavatars"
	YouthAvatarsAPIVersion = "1.2"
)

// YouthAvatarsXML contains data of Avatars for all players of a youth team.
type YouthAvatarsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	YouthTeam struct {
		ID      id.YouthTeam          `xml:"YouthTeamId"`
		Players []*YouthPlayerAvatars `xml:"YouthPlayers>YouthPlayer"`
	} `xml:"YouthTeam"`
}

// YouthPlayerAvatars is a container for a youth player's avatar.
type YouthPlayerAvatars struct {
	PlayerID id.YouthPlayer `xml:"YouthPlayerID"`
	Avatar   Avatar         `xml:"Avatar"`
}
