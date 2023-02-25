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
		ID      id.Team          `xml:"TeamId"`
		Players []*PlayerAvatars `xml:"Players>Player"`
	} `xml:"Team"`
}

// PlayerAvatars is a container for a player.
type PlayerAvatars struct {
	PlayerID id.Player `xml:"PlayerID"`
	Avatar   Avatar    `xml:"Avatar"`
}
