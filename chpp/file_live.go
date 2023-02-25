package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LiveAPIFile    = "live"
	LiveAPIVersion = "1.9"
)

// LiveXML ...
type LiveXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Matches *LiveMatchInfo `xml:"MatchList>Match"`
}

// LiveMatchInfo ...
type LiveMatchInfo struct {
	SourceSystem string       `xml:"SourceSystem"`
	MatchID      id.Match     `xml:"MatchID"`
	MatchDate    HattrickTime `xml:"MatchDate"`

	HomeTeam struct {
		ID             id.Team        `xml:"HomeTeamID"`
		Name           string         `xml:"HomeTeamName"`
		ShortName      string         `xml:"HomeTeamShortName"`
		StartingLineup StartingLineUp `xml:"StartingLineup"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID             id.Team        `xml:"AwayTeamID"`
		Name           string         `xml:"AwayTeamName"`
		ShortName      string         `xml:"AwayTeamShortName"`
		StartingLineup StartingLineUp `xml:"StartingLineup"`
	} `xml:"AwayTeam"`

	Substitutions struct {
		Substitutions []*struct {
			TeamID id.Team `xml:"TeamID"`
			// If substitution: The player leaving.
			// If behaviour change: The player changing his behaviour.
			// If position swap: the first player that will change his position.
			SubjectPlayerID id.Player `xml:"SubjectPlayerID"`

			// If substitution: The player entering
			// If behaviour change: The player changing behaviour.
			// If position swap: The player to swap with.
			ObjectPlayerID       id.Player        `xml:"ObjectPlayerID"`
			OrderType            MatchOrderType   `xml:"OrderType"`
			NewPositionID        MatchRole        `xml:"NewPositionId"`
			NewPositionBehaviour MatchBehaviourID `xml:"NewPositionBehaviour"`
			MatchMinute          uint             `xml:"MatchMinute"`
		} `xml:"Substitution"`
	} `xml:"Substitutions"`

	EventList struct {
		HomeGoals uint `xml:"HomeGoals"`
		AwayGoals uint `xml:"AwayGoals"`
	} `xml:"EventList"`
}

// StartingLineUp ...
type StartingLineUp struct {
	Players []*LiveMatchPlayer `xml:"Player"`
}

// LiveMatchPlayer ...
type LiveMatchPlayer struct {
	PlayerID   id.Player        `xml:"PlayerID"`
	RoleID     MatchRole        `xml:"RoleID"`
	PlayerName string           `xml:"PlayerName"`
	Behaviour  MatchBehaviourID `xml:"Behaviour"`
}
