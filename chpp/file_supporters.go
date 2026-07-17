package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	SupportersAPIFile    = "supporters"
	SupportersAPIVersion = "1.0"
)

// SupportersXML contains the teams a user supports, or the teams that
// support a given team, depending on the requested actionType.
type SupportersXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The teams the requested user supports. Only sent for
	// actionType=supportedteams.
	SupportedTeams struct {
		// How many teams the user supports in total.
		TotalItems uint             `xml:"TotalItems,attr"`
		Teams      []*SupporterTeam `xml:"SupportedTeam"`
	} `xml:"SupportedTeams"`

	// The teams supporting the requested team. Only sent for
	// actionType=mysupporters.
	MySupporters struct {
		// How many teams support this team in total.
		TotalItems uint             `xml:"TotalItems,attr"`
		Teams      []*SupporterTeam `xml:"SupporterTeam"`
	} `xml:"MySupporters"`
}

// SupporterTeam is a single supporter/supported team relationship. The
// LastMatch, NextMatch and PressAnnouncement containers are only sent for
// actionType=supportedteams.
type SupporterTeam struct {
	UserID id.User `xml:"UserId"`

	// The "username" or "nickname" used in forums and around the site;
	// not a login credential.
	LoginName string `xml:"LoginName"`

	TeamID   id.Team `xml:"TeamId"`
	TeamName string  `xml:"TeamName"`

	LeagueID   id.League `xml:"LeagueID"`
	LeagueName string    `xml:"LeagueName"`

	SeriesID   id.Series `xml:"LeagueLevelUnitID"`
	SeriesName string    `xml:"LeagueLevelUnitName"`

	LastMatch *SupporterTeamLastMatch `xml:"LastMatch"`
	NextMatch *SupporterTeamNextMatch `xml:"NextMatch"`

	// The team's most recent press announcement.
	PressAnnouncement *struct {
		SendDate HattrickTime `xml:"PressAnnouncementSendDate"`
		Subject  string       `xml:"PressAnnouncementSubject"`
		Body     string       `xml:"PressAnnouncementBody"`
	} `xml:"PressAnnouncement"`
}

// SupporterTeamLastMatch is a supported team's last played match.
type SupporterTeamLastMatch struct {
	ID   id.Match     `xml:"LastMatchId"`
	Date HattrickTime `xml:"LastMatchDate"`

	HomeTeamID   id.Team `xml:"LastMatchHomeTeamId"`
	HomeTeamName string  `xml:"LastMatchHomeTeamName"`
	HomeGoals    uint    `xml:"LastMatchHomeGoals"`

	AwayTeamID   id.Team `xml:"LastMatchAwayTeamId"`
	AwayTeamName string  `xml:"LastMatchAwayTeamName"`
	AwayGoals    uint    `xml:"LastMatchAwayGoals"`
}

// SupporterTeamNextMatch is a supported team's next upcoming match.
type SupporterTeamNextMatch struct {
	ID   id.Match     `xml:"NextMatchId"`
	Date HattrickTime `xml:"NextMatchDate"`

	HomeTeamID   id.Team `xml:"NextMatchHomeTeamId"`
	HomeTeamName string  `xml:"NextMatchHomeTeamName"`

	AwayTeamID   id.Team `xml:"NextMatchAwayTeamId"`
	AwayTeamName string  `xml:"NextMatchAwayTeamName"`
}
