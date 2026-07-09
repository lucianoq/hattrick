package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchLineupAPIFile    = "matchlineup"
	MatchLineupAPIVersion = "2.1"
)

// MatchLineupXML contains the lineup (starting eleven, substitutions, and
// end-of-match ratings) used by one team in a finished match.
type MatchLineupXML struct {
	Envelope
	UserID id.User `xml:"User"`

	*MatchLineup
}

// MatchLineup is the requested team's lineup for a specific match: the
// starting lineup, substitutions made during the match, and the final
// lineup with end-of-match player ratings.
type MatchLineup struct {
	MatchID id.Match `xml:"MatchID"`

	// Specifies which source system the match belongs to.
	SourceSystem SourceSystem `xml:"SourceSystem"`

	MatchType MatchType `xml:"MatchType"`

	// This will be either
	// LeagueLevelUnitId (for League),
	// CupId (Cup, Hattrick Masters, World Cup and U-20 World Cup),
	// LadderId, TournamentId, or 0 for friendly, qualification,
	// single matches and preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

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
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`
		// The aggregated experience level of the team.
		ExperienceLevel SkillLevel `xml:"ExperienceLevel"`
		// The style of play used in the match. Always 0 for youth matches.
		StyleOfPlay CoachModifier `xml:"StyleOfPlay"`

		StartingLineup struct {
			Player []struct {
				PlayerID  id.Player `xml:"PlayerID"`
				RoleID    MatchRole `xml:"RoleID"`
				FirstName string    `xml:"FirstName"`
				LastName  string    `xml:"LastName"`
				NickName  string    `xml:"NickName"`
				// Not provided for the Captain and the set pieces taker.
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
				// If b/ behaviour change: (once more) the player changing his behaviour.
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

		// The final lineup and end-of-match ratings, as opposed to
		// StartingLineup which only reflects the kick-off eleven.
		Lineup struct {
			Player []struct {
				PlayerID id.Player `xml:"PlayerID"`

				// An integer indicating which formal 'slot' (Role) the player
				// has filled in the match. All games are 553 based.
				RoleID    MatchRole `xml:"RoleID"`
				FirstName string    `xml:"FirstName"`
				LastName  string    `xml:"LastName"`
				NickName  string    `xml:"NickName"`

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
