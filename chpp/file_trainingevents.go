package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TrainingEventsAPIFile    = "trainingevents"
	TrainingEventsAPIVersion = "1.3"
)

// TrainingEventsXML contains the history of skill-level changes (training
// events) recorded for a single player.
type TrainingEventsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	Player struct {
		PlayerID id.Player `xml:"PlayerID"`

		TrainingEvents struct {
			// If a match is running this value is not available.
			Available bool             `xml:"Available,attr"`
			Events    []*TrainingEvent `xml:"TrainingEvent"`
		} `xml:"TrainingEvents"`
	} `xml:"Player"`
}

// TrainingEvent is a single skill-level change event for a player.
type TrainingEvent struct {
	// The order of the event relative to the other events for this
	// player.
	Index uint `xml:"Index,attr"`

	SkillID  SkillID    `xml:"SkillID"`
	OldLevel SkillLevel `xml:"OldLevel"`
	NewLevel SkillLevel `xml:"NewLevel"`

	Season     uint `xml:"Season"`
	MatchRound uint `xml:"MatchRound"`

	// Ranges from 1 to 7.
	DayNumber uint `xml:"DayNumber"`
}
