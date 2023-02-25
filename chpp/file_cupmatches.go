package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	CupMatchesAPIFile    = "cupmatches"
	CupMatchesAPIVersion = "1.4"
)

// CupMatchesXML ...
type CupMatchesXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// Container for the cup and the matches.
	Cup *Cup `xml:"Cup"`
}

// Cup ...
type Cup struct {
	ID      id.Cup      `xml:"CupID"`
	Season  uint        `xml:"CupSeason"`
	Round   uint        `xml:"CupRound"`
	Name    string      `xml:"CupName"`
	Matches []*CupMatch `xml:"MatchList>Match"`
}

// CupMatch ...
type CupMatch struct {
	MatchID   id.Match     `xml:"MatchID"`
	MatchDate HattrickTime `xml:"MatchDate"`
	HomeTeam  struct {
		ID   id.Team `xml:"TeamId"`
		Name string  `xml:"TeamName"`
	} `xml:"HomeTeam"`
	AwayTeam struct {
		ID   id.Team `xml:"TeamId"`
		Name string  `xml:"TeamName"`
	} `xml:"AwayTeam"`
	MatchResult struct {
		Available bool `xml:"Available,attr"`
		HomeGoals uint `xml:"HomeGoals"`
		AwayGoals uint `xml:"AwayGoals"`
	} `xml:"MatchResult"`
	LeagueInfo struct {
		Available      bool      `xml:"Available,attr"`
		HomeLeagueID   id.League `xml:"HomeLeagueID"`
		AwayLeagueID   id.League `xml:"AwayLeagueID"`
		HomeLeagueName string    `xml:"HomeLeagueName"`
		AwayLeagueName string    `xml:"AwayLeagueName"`
	} `xml:"LeagueInfo"`
}
