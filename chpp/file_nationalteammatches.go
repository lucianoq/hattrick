package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	NationalTeamMatchesAPIFile    = "nationalteammatches"
	NationalTeamMatchesAPIVersion = "1.4"
)

// NationalTeamMatchesXML contains the matches played (or to be played) by
// national teams of a given office type.
type NationalTeamMatchesXML struct {
	Envelope

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// The requested LeagueOfficeTypeID (2 = National teams, 4 = U-20
	// teams).
	LeagueOfficeTypeID LeagueOfficeTypeID `xml:"LeagueOfficeTypeID"`

	Matches []*NationalTeamMatch `xml:"Matches>Match"`
}

// NationalTeamMatch is a single match played (or to be played) by a
// national team.
type NationalTeamMatch struct {
	MatchID   id.Match     `xml:"MatchID"`
	MatchDate HattrickTime `xml:"MatchDate"`
	MatchType MatchType    `xml:"MatchType"`

	// This will be either LeagueLevelUnitId (for League), CupId (Cup,
	// Hattrick Masters, World Cup and U-20 World Cup), LadderId,
	// TournamentId, or 0 for friendly, qualification, single matches and
	// preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

	HomeTeamID   id.NationalTeam `xml:"HomeTeamId"`
	HomeTeamName string          `xml:"HomeTeamName"`
	AwayTeamID   id.NationalTeam `xml:"AwayTeamId"`
	AwayTeamName string          `xml:"AwayTeamName"`

	// Empty (zero value) if the match result is not available yet.
	MatchResult struct {
		HomeGoals uint `xml:"HomeGoals"`
		AwayGoals uint `xml:"AwayGoals"`
	} `xml:"MatchResult"`
}
