package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ClubAPIFile    = "club"
	ClubAPIVersion = "1.5"
)

// ClubXML contains information about a team's staff (specialists) and
// youth squad.
type ClubXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Team *Club `xml:"Team"`
}

// Club holds a team's specialist staff levels and youth squad status.
type Club struct {
	TeamID   id.Team `xml:"TeamID"`
	TeamName string  `xml:"TeamName"`
	Staff    struct {
		AssistantTrainerLevels  uint `xml:"AssistantTrainerLevels"`
		FinancialDirectorLevels uint `xml:"FinancialDirectorLevels"`
		FormCoachLevels         uint `xml:"FormCoachLevels"`
		MedicLevels             uint `xml:"MedicLevels"`
		SpokespersonLevels      uint `xml:"SpokespersonLevels"`
		SportPsychologistLevels uint `xml:"SportPsychologistLevels"`
		TacticalAssistantLevels uint `xml:"TacticalAssistantLevels"`
	} `xml:"Staff"`
	YouthSquad struct {
		// The weekly youth squad investment.
		Investment Money `xml:"Investment"`

		// Whether the team has promoted a youth player to the senior squad
		// this week.
		HasPromoted bool `xml:"HasPromoted"`
		YouthLevel  uint `xml:"YouthLevel"`
	} `xml:"YouthSquad"`
}
