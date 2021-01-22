package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchLineupAPIFile    = "matchlineup"
	MatchLineupAPIVersion = "2.0"
)

// MatchLineupXML ...
type MatchLineupXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	*MatchLineup
}

// MatchLineup ...
type MatchLineup struct {
	MatchID   id.Match  `xml:"MatchID"`
	IsYouth   bool      `xml:"IsYouth"`
	MatchType MatchType `xml:"MatchType"`

	// This will be either
	// LeagueLevelUnitId (for League),
	// CupId (Cup, Hattrick Masters, World Cup and U-20 World Cup),
	// LadderId, TournamentId, or 0 for friendly, qualification,
	// single matches and preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

	MatchDate HattrickTime `xml:"MatchDate"`

	HomeTeam struct {
		HomeTeamID   id.Team `xml:"HomeTeamID"`
		HomeTeamName string  `xml:"HomeTeamName"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		AwayTeamID   id.Team `xml:"AwayTeamID"`
		AwayTeamName string  `xml:"AwayTeamName"`
	} `xml:"AwayTeam"`

	Arena struct {
		ArenaID   id.Arena `xml:"ArenaID"`
		ArenaName string   `xml:"ArenaName"`
	} `xml:"Arena"`

	Team struct {
		TeamID          id.Team       `xml:"TeamID"`
		TeamName        string        `xml:"TeamName"`
		ExperienceLevel SkillLevel    `xml:"ExperienceLevel"`
		StyleOfPlay     CoachModifier `xml:"StyleOfPlay"`

		StartingLineup struct {
			Player []struct {
				PlayerID  id.Player        `xml:"PlayerID"`
				RoleID    id.Role          `xml:"RoleID"`
				FirstName string           `xml:"FirstName"`
				LastName  string           `xml:"LastName"`
				NickName  string           `xml:"NickName"`
				Behaviour MatchBehaviourID `xml:"Behaviour"`
			} `xml:"Player"`
		} `xml:"StartingLineup"`

		Substitutions struct {
			Substitution []struct {
				// The globally unique TeamID for which the orders were sent.
				TeamID id.Team `xml:"TeamID"`

				// If a/ substitution: The player leaving.
				// If b/ behaviour change: The player changing his behaviour.
				// If c/ position swap: the first player that will change his position.
				SubjectPlayerID id.Player `xml:"SubjectPlayerID"`

				// If a/ substitution: The player entering
				// If b/ behaviour change: (once more) The player changing behaviour).
				// If c/ position swap: The player to swap with.
				ObjectPlayerID id.Player `xml:"ObjectPlayerID"`

				// The type of the order.
				OrderType MatchOrderType `xml:"OrderType"`

				// An integer representing the new position for the (entering or staying) player.
				NewPositionID MatchRole `xml:"NewPositionId"`

				// An integer representing the new behaviour for the (entering or staying) player.
				NewPositionBehaviour MatchBehaviourID `xml:"NewPositionBehaviour"`

				// An integer indicating in which minute the order was executed.
				MatchMinute uint `xml:"MatchMinute"`

				// Which part of the match the order was executed.
				MatchPart MatchPart `xml:"MatchPart"`
			}
		} `xml:"Substitutions"`

		Lineup struct {
			Player []struct {
				PlayerID id.Player `xml:"PlayerID"`

				// An integer indicating which formal 'slot' (Role) the player has filled in the match. All games are 553 based.
				RoleID    id.Role `xml:"RoleID"`
				FirstName string  `xml:"FirstName"`
				LastName  string  `xml:"LastName"`
				NickName  string  `xml:"NickName"`

				// The number of stars the player is rated for.
				RatingStars float64 `xml:"RatingStars"`

				// The number of stars the player is rated for at the end of the game. This element is not provided for youth.
				RatingStarsEndOfMatch float64 `xml:"RatingStarsEndOfMatch"`

				// An integer indicating the individual order or repositioning that the player has played with.
				Behaviour MatchBehaviourID `xml:"Behaviour"`
			} `xml:"Player"`
		} `xml:"Lineup"`
	} `xml:"Team"`
}
