package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchesAPIFile    = "matches"
	MatchesAPIVersion = "2.8"
)

// MatchesXML ...
type MatchesXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	IsYouth bool `xml:"IsYouth"`

	Team struct {
		TeamID        id.Team `xml:"TeamID"`
		TeamName      string  `xml:"TeamName"`
		ShortTeamName string  `xml:"ShortTeamName"`

		League struct {
			LeagueID   id.League `xml:"LeagueID"`
			LeagueName string    `xml:"LeagueName"`
		} `xml:"League"`

		LeagueLevelUnit struct {
			LeagueLevelUnitID   id.LeagueLevelUnit `xml:"LeagueLevelUnitID"`
			LeagueLevelUnitName string             `xml:"LeagueLevelUnitName"`
			LeagueLevel         uint               `xml:"LeagueLevel"`
		} `xml:"LeagueLevelUnit"`

		MatchList struct {
			Match []*Match `xml:"Match"`
		} `xml:"MatchList"`
	} `xml:"Team"`
}

// Match ...
type Match struct {
	// The globally unique identifier of the match.
	MatchID id.Match `xml:"MatchID"`

	HomeTeam struct {
		HomeTeamID            id.Team `xml:"HomeTeamID"`
		HomeTeamName          string  `xml:"HomeTeamName"`
		HomeTeamNameShortName string  `xml:"HomeTeamNameShortName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		AwayTeamID            id.Team `xml:"AwayTeamID"`
		AwayTeamName          string  `xml:"AwayTeamName"`
		AwayTeamNameShortName string  `xml:"AwayTeamNameShortName"`
	} `xml:"AwayTeam"`

	MatchDate HattrickTime `xml:"MatchDate"`

	// SourceSystem tells from witch system the match is ex: hattrick, youth,
	// htointegrated
	SourceSystem string `xml:"SourceSystem"`

	// Integer defining the type of match.
	MatchType MatchType `xml:"MatchType"`

	// This will be either
	// LeagueLevelUnitId (for League),
	// CupId (Cup, Hattrick Masters, World Cup and U-20 World Cup),
	// LadderId, TournamentId, or 0 for friendly, qualification,
	// single matches and preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

	// 1 = National/Divisional cup
	// 2 = Challenger cup
	// 3 = Consolation cup
	CupLevel CupLevel `xml:"CupLevel"`

	// In Challenger cups:
	// 1 = Emerald (start week 2),
	// 2 = Ruby (start week 3),
	// 3 = Sapphire (start week 4).
	// Always 1 for National/Divisional (main cups) and Consolation cups.
	CupLevelIndex CupLevelIndex `xml:"CupLevelIndex"`

	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`

	// Specifying whether the match is FINISHED, ONGOING or UPCOMING.
	Status MatchStatus `xml:"Status"`

	// A boolean value only supplied for upcoming matches (haven't
	// started yet) of your own team that signifies whether you have
	// given orders or not. If the request is for another team than
	// your own (even if it is for your opponent), this data is not
	// sent.
	OrdersGiven bool `xml:"OrdersGiven"`
}
