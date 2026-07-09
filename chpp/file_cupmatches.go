package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	CupMatchesAPIFile    = "cupmatches"
	CupMatchesAPIVersion = "1.4"
)

// CupMatchesXML contains the matches for a given cup, season, and round.
type CupMatchesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for the cup and the matches.
	Cup *Cup `xml:"Cup"`
}

// Cup identifies a cup competition, season and round, and holds its matches.
type Cup struct {
	ID      id.Cup      `xml:"CupID"`
	Season  uint        `xml:"CupSeason"`
	Round   uint        `xml:"CupRound"`
	Name    string      `xml:"CupName"`
	Matches []*CupMatch `xml:"MatchList>Match"`
}

// CupMatch is a single match within a cup round.
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

	// The match result, only available once the match has been played.
	MatchResult struct {
		Available bool `xml:"Available,attr"`
		HomeGoals uint `xml:"HomeGoals"`
		AwayGoals uint `xml:"AwayGoals"`
	} `xml:"MatchResult"`

	// League data for the participating teams, only available for
	// international cups.
	LeagueInfo struct {
		Available      bool      `xml:"Available,attr"`
		HomeLeagueID   id.League `xml:"HomeLeagueID"`
		AwayLeagueID   id.League `xml:"AwayLeagueID"`
		HomeLeagueName string    `xml:"HomeLeagueName"`
		AwayLeagueName string    `xml:"AwayLeagueName"`
	} `xml:"LeagueInfo"`
}
