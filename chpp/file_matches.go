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
		ID        id.Team `xml:"TeamID"`
		Name      string  `xml:"TeamName"`
		ShortName string  `xml:"ShortTeamName"`

		League struct {
			ID   id.League `xml:"LeagueID"`
			Name string    `xml:"LeagueName"`
		} `xml:"League"`

		LeagueLevelUnit struct {
			ID    id.LeagueLevelUnit `xml:"LeagueLevelUnitID"`
			Name  string             `xml:"LeagueLevelUnitName"`
			Level uint               `xml:"LeagueLevel"`
		} `xml:"LeagueLevelUnit"`

		Matches []*Match `xml:"MatchList>Match"`
	} `xml:"Team"`
}

// Match ...
type Match struct {
	// The globally unique identifier of the match.
	MatchID id.Match `xml:"MatchID"`

	HomeTeam struct {
		ID        id.Team `xml:"HomeTeamID"`
		Name      string  `xml:"HomeTeamName"`
		ShortName string  `xml:"HomeTeamNameShortName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID        id.Team `xml:"AwayTeamID"`
		Name      string  `xml:"AwayTeamName"`
		ShortName string  `xml:"AwayTeamNameShortName"`
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
