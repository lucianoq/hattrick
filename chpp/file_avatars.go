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

	// Container for the team to which the data applies.
	Team struct {
		// The globally unique teamID.
		ID      id.Team          `xml:"TeamId"`
		Players []*PlayerAvatars `xml:"Players>Player"`
	} `xml:"Team"`
}

// PlayerAvatars pairs a player with the avatar image data used to render
// their player card.
type PlayerAvatars struct {
	// The globally unique PlayerID.
	PlayerID id.Player `xml:"PlayerID"`

	// Container for the elements to build the avatar.
	Avatar Avatar `xml:"Avatar"`
}
