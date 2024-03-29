package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	ChallengesAPIFile    = "challenges"
	ChallengesAPIVersion = "1.6"
)

// ChallengesXML ...
type ChallengesXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Team struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`

		// Container for the challenges that the user has done.
		ChallengesByMe []*ChallengeByMe `xml:"ChallengesByMe>Challenge"`

		// Container for the offers of friendlies that other teams have given to
		// the logged on user's team.
		OffersByOthers []*OffersByOthers `xml:"OffersByOthers>Offer"`

		ChallengeableResult *struct {
			Opponent []*struct {
				IsChallengeable bool    `xml:"IsChallengeable"`
				UserID          id.User `xml:"UserId"`
				TeamID          id.Team `xml:"TeamId"`
				TeamName        string  `xml:"TeamName"`
				LogoURL         string  `xml:"LogoURL"`
			} `xml:"Opponent"`
		} `xml:"ChallengeableResult"`
	} `xml:"Team"`
}

// ChallengeByMe ...
type ChallengeByMe struct {
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`
	MatchTime       HattrickTime     `xml:"MatchTime"`
	MatchID         id.Match         `xml:"MatchID"`
	FriendlyType    FriendlyType     `xml:"FriendlyType"`

	Opponent struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`

		// The url for the opponent's logo, if it has one.
		LogoURL string `xml:"LogoURL"`

		Arena struct {
			ID   id.Arena `xml:"ArenaID"`
			Name string   `xml:"ArenaName"`
		} `xml:"Arena"`

		// TODO check League or Country. Inconsistency.
		Country struct {
			ID   id.Country `xml:"CountryID"`
			Name string     `xml:"CountryName"`
		} `xml:"Country"`

		// Indicating whether the challenge/offer has been accepted
		// and the match arranged.
		IsAgreed bool `xml:"IsAgreed"`
	} `xml:"Opponent"`
}

// OffersByOthers ...
type OffersByOthers struct {
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`
	MatchTime       HattrickTime     `xml:"MatchTime"`
	MatchID         id.Match         `xml:"MatchID"`
	FriendlyType    FriendlyType     `xml:"FriendlyType"`

	Opponent struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`
		LogoURL  string  `xml:"LogoURL"`

		Arena struct {
			ID   id.Arena `xml:"ArenaID"`
			Name string   `xml:"ArenaName"`
		} `xml:"Arena"`

		// TODO check League or Country. Inconsistency.
		League struct {
			ID   id.League `xml:"LeagueID"`
			Name string    `xml:"LeagueName"`
		} `xml:"League"`

		// Indicating whether the challenge/offer has been accepted
		// and the match arranged.
		IsAgreed bool `xml:"IsAgreed"`
	} `xml:"Opponent"`
}
