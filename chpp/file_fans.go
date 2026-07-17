package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// XML file name and version.
const (
	FansAPIFile    = "fans"
	FansAPIVersion = "1.3"
)

// FansXML contains a team's fan club information, current fan mood/season
// expectations, and the three most recent and three upcoming matches.
type FansXML struct {
	Envelope
	UserID id.User `xml:"UserID"`

	Team Fans `xml:"Team"`
}

// Fans holds a team's fan club data: mood, season expectations, and the
// last three and next three series/qualification/cup/masters matches.
type Fans struct {
	TeamID    id.Team    `xml:"TeamID"`
	FanClubID id.Fanclub `xml:"FanClubId"`

	// The name of the fan club. Empty if the user hasn't set one.
	FanClubName string `xml:"FanClubName"`

	// The popularity among supporters. Unavailable while a match is running.
	FanMood FanMood `xml:"FanMood"`
	Members uint    `xml:"Members"`

	// The fans' expectations for the current season.
	FanSeasonExpectation FanSeasonExpectation `xml:"FanSeasonExpectation"`

	// The three most recently played matches.
	PlayedMatches struct {
		Match []struct {
			MatchID  id.Match `xml:"MatchID"`
			HomeTeam struct {
				// The teamID of the home team. Negative for street teams.
				HomeTeamID   id.Team `xml:"HomeTeamID"`
				HomeTeamName string  `xml:"HomeTeamName"`
			} `xml:"HomeTeam"`
			AwayTeam struct {
				// The teamID of the away team. Negative for street teams.
				AwayTeamID   id.Team `xml:"AwayTeamID"`
				AwayTeamName string  `xml:"AwayTeamName"`
			} `xml:"AwayTeam"`
			MatchDate HattrickTime `xml:"MatchDate"`
			MatchType MatchType    `xml:"MatchType"`
			HomeGoals uint         `xml:"HomeGoals"`
			AwayGoals uint         `xml:"AwayGoals"`

			// The fans' expectation for this match, before it was played.
			FanMatchExpectation FanMatchExpectation `xml:"FanMatchExpectation"`

			// The fan mood after this match was played.
			FanMoodAfterMatch FanMood `xml:"FanMoodAfterMatch"`
			Weather           Weather `xml:"Weather"`
			SoldSeats         uint    `xml:"SoldSeats"`
		} `xml:"Match"`
	} `xml:"PlayedMatches"`

	// The three upcoming matches.
	UpcomingMatches struct {
		Match []struct {
			MatchID  id.Match `xml:"MatchID"`
			HomeTeam struct {
				// The teamID of the home team. Negative for street teams.
				HomeTeamID   id.Team `xml:"HomeTeamID"`
				HomeTeamName string  `xml:"HomeTeamName"`
			} `xml:"HomeTeam"`
			AwayTeam struct {
				// The teamID of the away team. Negative for street teams.
				AwayTeamID   id.Team `xml:"AwayTeamID"`
				AwayTeamName string  `xml:"AwayTeamName"`
			} `xml:"AwayTeam"`
			MatchDate HattrickTime `xml:"MatchDate"`
			MatchType MatchType    `xml:"MatchType"`

			// The fans' expectation for this upcoming match.
			FanMatchExpectation FanMatchExpectation `xml:"FanMatchExpectation"`
		} `xml:"Match"`
	} `xml:"UpcomingMatches"`
}
