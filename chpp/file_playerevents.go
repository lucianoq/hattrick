package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	PlayerEventsAPIFile    = "playerevents"
	PlayerEventsAPIVersion = "1.3"
)

// PlayerEventsXML contains the log of notable events (injuries, cards,
// transfers, etc.) recorded for a single player.
type PlayerEventsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Indicates which Supporter package the fetching user has, or empty if
	// not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	Player struct {
		PlayerID id.Player `xml:"PlayerID"`

		// Container for the player's events.
		PlayerEvents struct {
			// If a match is running this value is not available.
			Available bool               `xml:"Available,attr"`
			Events    []*PlayerEventItem `xml:"PlayerEvent"`
		} `xml:"PlayerEvents"`
	} `xml:"Player"`
}

// PlayerEventItem is a single event that happened to a player. No fixed
// list of event types is documented; each CHPP has to collect its own
// event list by observing PlayerEventTypeID/EventText over time.
type PlayerEventItem struct {
	// The date and time when the event was recorded.
	EventDate HattrickTime `xml:"EventDate"`

	// An identifier to show which type of event it is.
	PlayerEventTypeID uint `xml:"PlayerEventTypeID"`

	// String describing the event. May contain HTML tags.
	EventText string `xml:"EventText"`
}
