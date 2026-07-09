package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LeagueLevelsAPIFile    = "leaguelevels"
	LeagueLevelsAPIVersion = "1.0"
)

// LeagueLevelsXML contains the division/level structure (series counts,
// team counts, promotion/demotion slots) of a league for a season.
type LeagueLevelsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	LeagueID id.League `xml:"LeagueID"`
	Season   uint      `xml:"Season"`

	NrOfLeagueLevels uint `xml:"NrOfLeagueLevels"`

	Levels []*LeagueLevelItem `xml:"LeagueLevelList>LeagueLevelItem"`
}

// LeagueLevelItem is a single division level of a league (e.g. all of a
// country's top division, all its second-division series, etc).
type LeagueLevelItem struct {
	// 1 = top division, 2 = second division, and so on.
	LeagueLevel uint `xml:"LeagueLevel"`

	// NOTE: documented as a string, not a uint, despite being a count.
	NrOfLeagueLevelUnits string `xml:"NrOfLeagueLevelUnits"`

	NrOfTeams uint `xml:"NrOfTeams"`

	// A comma-separated string of LeagueLevelUnitIDs, not a nested XML
	// list.
	LeagueLevelUnitIDList string `xml:"LeagueLevelUnitIdList"`

	NrOfSharedPromotionSlotsPerSeries        uint `xml:"NrOfSharedPromotionSlotsPerSeries"`
	NrOfDirectPromotionSlotsPerSeries        uint `xml:"NrOfDirectPromotionSlotsPerSeries"`
	NrOfQualificationPromotionSlotsPerSeries uint `xml:"NrOfQualificationPromotionSlotsPerSeries"`
	NrOfDirectDemotionSlotsPerSeries         uint `xml:"NrOfDirectDemotionSlotsPerSeries"`
	NrOfQualificationDemotionSlotsPerSeries  uint `xml:"NrOfQualificationDemotionSlotsPerSeries"`
}
