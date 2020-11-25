package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// XML file name and version.
const (
	FansAPIFile    = "fans"
	FansAPIVersion = "1.3"
)

// FansXML ...
type FansXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"UserID"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Team Fans `xml:"Team"`
}

// Fans ...
type Fans struct {
	TeamID               id.Team              `xml:"TeamID"`
	FanclubID            id.Fanclub           `xml:"FanclubId"`
	FanclubName          string               `xml:"FanclubName"`
	FanMood              FanMood              `xml:"FanMood"`
	Members              uint                 `xml:"Members"`
	FanSeasonExpectation FanSeasonExpectation `xml:"FanSeasonExpectation"`
	PlayedMatches        struct {
		Match []struct {
			MatchID  id.Match `xml:"MatchID"`
			HomeTeam struct {
				HomeTeamID   id.Team `xml:"HomeTeamID"`
				HomeTeamName string  `xml:"HomeTeamName"`
			} `xml:"HomeTeam"`
			AwayTeam struct {
				AwayTeamID   id.Team `xml:"AwayTeamID"`
				AwayTeamName string  `xml:"AwayTeamName"`
			} `xml:"AwayTeam"`
			MatchDate           HattrickTime        `xml:"MatchDate"`
			MatchType           MatchType           `xml:"MatchType"`
			HomeGoals           uint                `xml:"HomeGoals"`
			AwayGoals           uint                `xml:"AwayGoals"`
			FanMatchExpectation FanMatchExpectation `xml:"FanMatchExpectation"`
			FanMoodAfterMatch   FanMood             `xml:"FanMoodAfterMatch"`
			Weather             Weather             `xml:"Weather"`
			SoldSeats           uint                `xml:"SoldSeats"`
		} `xml:"Match"`
	} `xml:"PlayedMatches"`
	UpcomingMatches struct {
		Match struct {
			MatchID  id.Match `xml:"MatchID"`
			HomeTeam struct {
				HomeTeamID   id.Team `xml:"HomeTeamID"`
				HomeTeamName string  `xml:"HomeTeamName"`
			} `xml:"HomeTeam"`
			AwayTeam struct {
				AwayTeamID   id.Team `xml:"AwayTeamID"`
				AwayTeamName string  `xml:"AwayTeamName"`
			} `xml:"AwayTeam"`
			MatchDate           HattrickTime        `xml:"MatchDate"`
			MatchType           MatchType           `xml:"MatchType"`
			FanMatchExpectation FanMatchExpectation `xml:"FanMatchExpectation"`
		} `xml:"Match"`
	} `xml:"UpcomingMatches"`
}
