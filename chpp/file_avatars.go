package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AvatarsAPIFile    = "avatars"
	AvatarsAPIVersion = "1.1"
)

// AvatarsXML contains avatar data for all players on the user's team.
type AvatarsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Team struct {
		ID      id.Team          `xml:"TeamId"`
		Players []*PlayerAvatars `xml:"Players>Player"`
	} `xml:"Team"`
}

// PlayerAvatars pairs a player with the avatar image data used to render
// their player card.
type PlayerAvatars struct {
	PlayerID id.Player `xml:"PlayerID"`
	Avatar   Avatar    `xml:"Avatar"`
}
