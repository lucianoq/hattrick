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

	// Container for a team in the group.
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
			// Container for data about a possible opponent.
			Opponent []*struct {
				// Whether this opponent is challengeable.
				IsChallengeable bool `xml:"IsChallengeable"`

				// The globally unique UserID of the opponent.
				UserID id.User `xml:"UserId"`

				// The live XML tag is "TeamId" (lowercase d), confirmed
				// against a live challenges 1.6 actionType=challengeable
				// response (2026-08-27) — the same doc-vs-live casing trap
				// as teamdetails' flag LeagueId. With "TeamID" this field
				// silently decoded to 0.
				TeamID id.Team `xml:"TeamId"`

				TeamName string `xml:"TeamName"`

				// URL for the opponent's logo. Live tag is "LogoURL"
				// (confirmed in the same 2026-08-27 response).
				LogoURL string `xml:"LogoURL"`
			} `xml:"Opponent"`
		} `xml:"ChallengeableResult"`
	} `xml:"Team"`
}

// ChallengeByMe is a friendly match challenge that the manager's team has
// made to an opponent.
type ChallengeByMe struct {
	// The globally unique identifier for the challenge/offer.
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`

	// The match start time. The doc claims "unsigned Integer" but the live
	// XML carries a datetime string like "2026-09-02 06:00:00" (confirmed
	// against a live challenges 1.6 response, 2026-08-27).
	MatchTime HattrickTime `xml:"MatchTime"`

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
	// The globally unique identifier for the challenge/offer.
	TrainingMatchID id.FriendlyMatch `xml:"TrainingMatchID"`

	// The match start time. The doc claims "unsigned Integer" but the live
	// XML carries a datetime string like "2026-09-02 06:00:00" (confirmed
	// against a live challenges 1.6 response, 2026-08-27).
	MatchTime HattrickTime `xml:"MatchTime"`

	// The match ID for the match created once the offer has been accepted.
	MatchID      id.Match     `xml:"MatchID"`
	FriendlyType FriendlyType `xml:"FriendlyType"`

	Opponent struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`

		// URL for the opponent's logo.
		LogoURL string `xml:"LogoURL"`

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
