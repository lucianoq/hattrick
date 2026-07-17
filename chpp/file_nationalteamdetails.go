package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	NationalTeamDetailsAPIFile    = "nationalteamdetails"
	NationalTeamDetailsAPIVersion = "1.9"
)

// NationalTeamDetailsXML contains the details of a single national team.
type NationalTeamDetailsXML struct {
	Envelope

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// Whether the team is currently playing a match. While true, some
	// fields below (e.g. Morale, SelfConfidence, SupportersPopularity)
	// are unavailable.
	IsPlayingMatch bool `xml:"IsPlayingMatch"`

	Team NationalTeamDetails `xml:"Team"`
}

// NationalTeamDetails is the container for a national team's details.
type NationalTeamDetails struct {
	ID   id.NationalTeam `xml:"TeamID"`
	Name string          `xml:"TeamName"`

	NationalCoach struct {
		UserID id.User `xml:"NationalCoachUserID"`

		// The "username" or "nickname" used in forums.
		LoginName string `xml:"NationalCoachLoginname"`
	} `xml:"NationalCoach"`

	League struct {
		ID   id.League `xml:"LeagueID"`
		Name string    `xml:"LeagueName"`
	} `xml:"League"`

	HomePage string `xml:"HomePage"`
	Logo     string `xml:"Logo"`

	DressURI          string `xml:"DressURI"`
	DressAlternateURI string `xml:"DressAlternateURI"`

	// The team's tactical experience with each formation.
	Experience442 SkillLevel `xml:"Experience442"`
	Experience433 SkillLevel `xml:"Experience433"`
	Experience451 SkillLevel `xml:"Experience451"`
	Experience352 SkillLevel `xml:"Experience352"`
	Experience532 SkillLevel `xml:"Experience532"`
	Experience343 SkillLevel `xml:"Experience343"`
	Experience541 SkillLevel `xml:"Experience541"`
	Experience523 SkillLevel `xml:"Experience523"`
	Experience550 SkillLevel `xml:"Experience550"`
	Experience253 SkillLevel `xml:"Experience253"`

	// Hidden (empty) starting 4 hours before every match the team plays,
	// and until the match ends.
	Morale TeamSpiritID `xml:"Morale"`

	// Hidden (empty) starting 4 hours before every match the team plays,
	// and until the match ends.
	SelfConfidence SelfConfidence `xml:"SelfConfidence"`

	// The popularity among the supporters of the national team. If a
	// match is running this value is not available.
	SupportersPopularity struct {
		Available bool                   `xml:"Available,attr"`
		Value     SupportersPopularityID `xml:",chardata"`
	} `xml:"SupportersPopularity"`

	RatingScore uint `xml:"RatingScore"`
	FanClubSize uint `xml:"FanClubSize"`

	// The national team's rank.
	Rank uint `xml:"Rank"`
}
