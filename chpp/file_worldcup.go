package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	WorldCupAPIFile    = "worldcup"
	WorldCupAPIVersion = "1.1"
)

// WorldCupXML contains the groups, matches, or match rounds of a
// national-team cup (e.g. the World Cup), depending on the requested
// actionType.
type WorldCupXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The scores/standings for each CupSeriesUnit ("group") in the round.
	// Only sent for actionType=viewGroups.
	Scores []*WorldCupScore `xml:"WorldCupScores>Team"`

	// The matches of the requested CupSeriesUnit ("group"). Only sent for
	// actionType=viewMatches.
	Matches []*WorldCupMatch `xml:"Matches>Match"`

	// All the available match rounds for the chosen cup/season.
	Rounds []*WorldCupRound `xml:"Rounds>Round"`
}

// WorldCupScore is a single team's standing within a CupSeriesUnit
// ("group"). Only sent for actionType=viewGroups.
type WorldCupScore struct {
	TeamID   id.Team `xml:"TeamID"`
	TeamName string  `xml:"TeamName"`

	// The team's place within the group.
	Place uint `xml:"Place"`

	CupSeriesUnitID   uint   `xml:"CupSeriesUnitID"`
	CupSeriesUnitName string `xml:"CupSeriesUnitName"`

	MatchesPlayed uint `xml:"MatchesPlayed"`
	GoalsFor      uint `xml:"GoalsFor"`
	GoalsAgainst  uint `xml:"GoalsAgainst"`
	Points        uint `xml:"Points"`
}

// WorldCupMatch is a single match within a CupSeriesUnit ("group"). Only
// sent for actionType=viewMatches.
type WorldCupMatch struct {
	// The globally unique matchID. May be -1 if the match has not been
	// arranged yet.
	MatchID id.Match `xml:"MatchID"`

	HomeTeam struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`
	} `xml:"AwayTeam"`

	MatchDate HattrickTime `xml:"MatchDate"`

	// Empty if the match is not finished.
	FinishedDate HattrickTime `xml:"FinishedDate"`

	// Empty (zero value) if the match is not finished.
	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`
}

// WorldCupRound identifies a single match round of a cup/season.
type WorldCupRound struct {
	// A key that indicates a certain round (which typically runs for
	// either 10 weeks or 9 days). Qualification is always matchRound 1.
	MatchRound uint         `xml:"MatchRound"`
	StartDate  HattrickTime `xml:"StartDate"`
}
