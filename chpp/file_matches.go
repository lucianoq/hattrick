package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchesAPIFile    = "matches"
	MatchesAPIVersion = "2.9"
)

// MatchesXML contains the most recent and upcoming matches for a
// particular team (senior or youth).
type MatchesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	IsYouth bool `xml:"IsYouth"`

	Team struct {
		ID        id.Team `xml:"TeamID"`
		Name      string  `xml:"TeamName"`
		ShortName string  `xml:"ShortTeamName"`

		League struct {
			ID   id.League `xml:"LeagueID"`
			Name string    `xml:"LeagueName"`
		} `xml:"League"`

		Series struct {
			ID    id.Series `xml:"LeagueLevelUnitID"`
			Name  string    `xml:"LeagueLevelUnitName"`
			Level uint      `xml:"LeagueLevel"`
		} `xml:"LeagueLevelUnit"`

		Matches []*Match `xml:"MatchList>Match"`
	} `xml:"Team"`
}

// Match is a single (past, ongoing, or upcoming) match for a team, as
// returned by the matches and matchesArchive files.
type Match struct {
	// The globally unique identifier of the match.
	MatchID id.Match `xml:"MatchID"`

	HomeTeam struct {
		ID        id.Team `xml:"HomeTeamID"`
		Name      string  `xml:"HomeTeamName"`
		ShortName string  `xml:"HomeTeamShortName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID        id.Team `xml:"AwayTeamID"`
		Name      string  `xml:"AwayTeamName"`
		ShortName string  `xml:"AwayTeamShortName"`
	} `xml:"AwayTeam"`

	MatchDate HattrickTime `xml:"MatchDate"`

	// SourceSystem tells from which system the match is, e.g. hattrick, youth,
	// htointegrated
	SourceSystem SourceSystem `xml:"SourceSystem"`

	// Integer defining the type of match.
	MatchType MatchType `xml:"MatchType"`

	// This will be either
	// LeagueLevelUnitId (for League),
	// CupId (Cup, Hattrick Masters, World Cup and U-20 World Cup),
	// LadderId, TournamentId, or 0 for friendly, qualification,
	// single matches and preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

	// The rules in place for the ladder or tournament in question, if any.
	// Only sent by matchesArchive. Defaults to 0 = no rules.
	MatchRuleID MatchRule `xml:"MatchRuleId"`

	// The id of the cup the match belongs to, if any. Only sent by
	// matchesArchive.
	CupID id.Cup `xml:"CupId"`

	// 1 = National/Divisional cup
	// 2 = Challenger cup
	// 3 = Consolation cup
	// 0 if MatchType is not a cup match.
	CupLevel CupLevel `xml:"CupLevel"`

	// In Challenger cups:
	// 1 = Emerald (start week 2),
	// 2 = Ruby (start week 3),
	// 3 = Sapphire (start week 4).
	// Always 1 for National/Divisional (main cups) and Consolation cups.
	// 0 if MatchType is not a cup match.
	CupLevelIndex CupLevelIndex `xml:"CupLevelIndex"`

	// Not sent by matchesArchive if the match is still upcoming or ongoing.
	HomeGoals uint `xml:"HomeGoals"`
	// Not sent by matchesArchive if the match is still upcoming or ongoing.
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
