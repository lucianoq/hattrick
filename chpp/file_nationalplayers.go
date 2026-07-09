package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	NationalPlayersAPIFile    = "nationalplayers"
	NationalPlayersAPIVersion = "1.5"
)

// NationalPlayersXML contains a national team's current squad, or the
// Supporter statistics of how many matches players have played for it.
type NationalPlayersXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter. Only sent for actionType=view.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Whether the fetching user is a Supporter. Only sent for
	// actionType=SupporterStats.
	UserIsSupporter bool `xml:"UserIsSupporter"`

	// Whether the fetching user has a clubhouse. Only sent for
	// actionType=SupporterStats.
	UserHasClubhouse bool `xml:"UserHasClubhouse"`

	// The requested ActionType.
	ActionType string `xml:"ActionType"`

	TeamID   id.NationalTeam `xml:"TeamID"`
	TeamName string          `xml:"TeamName"`

	// The national team's players. Only sent for actionType=view.
	Players []*NationalTeamPlayer `xml:"Players>Player"`

	// The Supporter statistics for the national team. Only sent for
	// actionType=SupporterStats.
	Stats *NationalPlayersStats `xml:"Stats"`
}

// NationalTeamPlayer is a single player currently in a national team's
// squad, as returned for actionType=view.
type NationalTeamPlayer struct {
	PlayerID id.Player `xml:"PlayerID"`
	Name     string    `xml:"PlayerName"`

	// Number of national team matches for which the player is suspended
	// due to cards. Only sent since API version 1.4.
	Cards uint `xml:"Cards"`

	// The player's specialty, if revealed. 0 if none or unknown. Only
	// sent since API version 1.5.
	Specialty SpecialtyID `xml:"Specialty"`

	// Only sent since API version 1.5.
	Avatar Avatar `xml:"Avatar"`
}

// NationalPlayersStats is the Supporter statistics of players who have
// performed for a national team, as returned for actionType=SupporterStats.
type NationalPlayersStats struct {
	// The requested MatchTypeCategory ("NT" or "WC").
	MatchTypeCategory string `xml:"MatchTypeCategory"`

	// The requested ShowAll value.
	ShowAll bool `xml:"ShowAll"`

	// True if there are more players available than were returned
	// (ShowAll=false caps the list at 20 players).
	MoreRecordsAvailable bool `xml:"MoreRecordsAvailable"`

	Players []*NationalPlayerStat `xml:"Players>Player"`
}

// NationalPlayerStat is a single player's number of matches played for a
// national team.
type NationalPlayerStat struct {
	PlayerID    id.Player `xml:"PlayerID"`
	Name        string    `xml:"PlayerName"`
	NrOfMatches uint      `xml:"NrOfMatches"`
}
