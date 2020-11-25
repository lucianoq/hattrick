package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ClubAPIFile    = "club"
	ClubAPIVersion = "1.5"
)

// ClubXML ...
type ClubXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Team *Club `xml:"Team"`
}

// Club ...
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
		Investment  Money `xml:"Investment"`
		HasPromoted bool  `xml:"HasPromoted"`
		YouthLevel  uint  `xml:"YouthLevel"`
	} `xml:"YouthSquad"`
}
