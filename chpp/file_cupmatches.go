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
	CupID     id.Cup `xml:"CupID"`
	CupSeason uint   `xml:"CupSeason"`
	CupRound  uint   `xml:"CupRound"`
	CupName   string `xml:"CupName"`
	MatchList struct {
		Match []*CupMatch `xml:"Match"`
	} `xml:"MatchList"`
}

// CupMatch ...
type CupMatch struct {
	MatchID   id.Match     `xml:"MatchID"`
	MatchDate HattrickTime `xml:"MatchDate"`
	HomeTeam  struct {
		TeamID   id.Team `xml:"TeamId"`
		TeamName string  `xml:"TeamName"`
	} `xml:"HomeTeam"`
	AwayTeam struct {
		TeamID   id.Team `xml:"TeamId"`
		TeamName string  `xml:"TeamName"`
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
