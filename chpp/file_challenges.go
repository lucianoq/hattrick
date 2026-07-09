package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ChallengesAPIFile    = "challenges"
	ChallengesAPIVersion = "1.6"
)

// ChallengesXML contains the friendlies the manager has challenged others to
// and the friendly offers other teams have made to the manager's team.
type ChallengesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Team struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`

		// Container for the challenges that the user has done.
		ChallengesByMe []*ChallengeByMe `xml:"ChallengesByMe>Challenge"`

		// Container for the offers of friendlies that other teams have given to
		// the logged on user's team.
		OffersByOthers []*OffersByOthers `xml:"OffersByOthers>Offer"`

		// Only populated for actionType=challengeable requests, reporting
		// whether each of the queried teams can be challenged.
		ChallengeableResult *struct {
			Opponent []*struct {
				IsChallengeable bool    `xml:"IsChallengeable"`
				UserID          id.User `xml:"UserId"`
				TeamID          id.Team `xml:"TeamID"`
				TeamName        string  `xml:"TeamName"`
				LogoURL         string  `xml:"LogoUrl"`
			} `xml:"Opponent"`
		} `xml:"ChallengeableResult"`
	} `xml:"Team"`
}

// ChallengeByMe is a friendly match challenge that the manager's team has
// made to an opponent.
type ChallengeByMe struct {
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`
	MatchTime       uint             `xml:"MatchTime"`

	// The match ID for the match created once the challenge has been accepted.
	MatchID      id.Match     `xml:"MatchID"`
	FriendlyType FriendlyType `xml:"FriendlyType"`

	Opponent struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`

		// The url for the opponent's logo, if it has one.
		LogoURL string `xml:"LogoURL"`

		Arena struct {
			ID   id.Arena `xml:"ArenaID"`
			Name string   `xml:"ArenaName"`
		} `xml:"Arena"`

		League struct {
			ID   id.League `xml:"LeagueID"`
			Name string    `xml:"LeagueName"`
		} `xml:"League"`

		// Container for the data about the country that the friendly would be
		// played in.
		Country struct {
			ID   id.Country `xml:"CountryID"`
			Name string     `xml:"CountryName"`
		} `xml:"Country"`

		// Indicating whether the challenge/offer has been accepted
		// and the match arranged.
		IsAgreed bool `xml:"IsAgreed"`
	} `xml:"Opponent"`
}

// OffersByOthers is a friendly match offer that another team has made to
// the manager's team.
type OffersByOthers struct {
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`
	MatchTime       uint             `xml:"MatchTime"`

	// The match ID for the match created once the offer has been accepted.
	MatchID      id.Match     `xml:"MatchID"`
	FriendlyType FriendlyType `xml:"FriendlyType"`

	Opponent struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`
		LogoURL  string  `xml:"LogoURL"`

		Arena struct {
			ID   id.Arena `xml:"ArenaID"`
			Name string   `xml:"ArenaName"`
		} `xml:"Arena"`

		League struct {
			ID   id.League `xml:"LeagueID"`
			Name string    `xml:"LeagueName"`
		} `xml:"League"`

		// Indicating whether the challenge/offer has been accepted
		// and the match arranged.
		IsAgreed bool `xml:"IsAgreed"`
	} `xml:"Opponent"`
}
