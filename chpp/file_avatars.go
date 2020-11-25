package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AvatarsAPIFile    = "avatars"
	AvatarsAPIVersion = "1.1"
)

// AvatarsXML contains data of Avatars for all players of user's team.
type AvatarsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Team struct {
		TeamID  id.Team `xml:"TeamId"`
		Players struct {
			Player []*PlayerAvatars `xml:"Player"`
		} `xml:"Players"`
	} `xml:"Team"`
}

// PlayerAvatars is a container for a player.
type PlayerAvatars struct {
	PlayerID id.Player `xml:"PlayerID"`
	Avatar   struct {

		// The URL to the card background-image. This will show a silhouette for
		// non-supporter teams.
		BackgroundImage string `xml:"BackgroundImage"`

		// The container for each avatar bodypart item. Two attribute named X
		// and Y indicates where the item should be positioned. There are
		// several of this container for each player. This container will not be
		// provided for non-supporter team.
		Layers []*struct {
			X uint `xml:"x"`
			Y uint `xml:"y"`

			// The URL to the bodypart item.
			Image string `xml:"Image"`
		} `xml:"Layer"`
	} `xml:"Avatar"`
}
