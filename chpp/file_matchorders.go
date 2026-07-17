package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchOrdersAPIFile    = "matchorders"
	MatchOrdersAPIVersion = "3.1"
)

// MatchOrdersXML contains the match orders (lineup, tactics and player
// orders) for a given match and team.
type MatchOrdersXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The globally unique identifier of the match.
	MatchID id.Match `xml:"MatchID"`

	// Specifies which source system the match belongs to.
	SourceSystem SourceSystem `xml:"sourceSystem"`

	MatchData MatchOrdersMatchData `xml:"MatchData"`
}

// MatchOrdersMatchData holds the match orders for one team, or diagnostic
// information about a set/predict operation.
type MatchOrdersMatchData struct {
	// True if the user is allowed to see the matchorder for this match,
	// false if too close to the match or the user is not involved in the
	// match. Only meaningful for actionType=view.
	Available bool `xml:"Available,attr"`

	// True if setting the matchorder was successful. Only present for
	// actionType=setmatchorder.
	OrdersSet bool `xml:"OrdersSet,attr"`

	// The reason for failing to set the matchorder. Only present for
	// actionType=setmatchorder and when OrdersSet is false.
	Reason string `xml:"Reason"`

	// The predicted tactic skill of the team. Only present for
	// actionType=predictratings.
	TacticSkill SkillLevel `xml:"TacticSkill"`

	// The predicted sector ratings of the team. Only present for
	// actionType=predictratings.
	RatingMidfield MatchRating `xml:"RatingMidfield"`
	RatingRightDef MatchRating `xml:"RatingRightDef"`
	RatingMidDef   MatchRating `xml:"RatingMidDef"`
	RatingLeftDef  MatchRating `xml:"RatingLeftDef"`
	RatingRightAtt MatchRating `xml:"RatingRightAtt"`
	RatingMidAtt   MatchRating `xml:"RatingMidAtt"`
	RatingLeftAtt  MatchRating `xml:"RatingLeftAtt"`

	// The home (formal, not taking arena location into account) team of
	// the match.
	HomeTeam struct {
		// The teamID of the Home team in the match. (Negative for street
		// teams.)
		ID   id.Team `xml:"HomeTeamID"`
		Name string  `xml:"HomeTeamName"`
	} `xml:"HomeTeam"`

	// The away (formal, not taking arena location into account) team of
	// the match.
	AwayTeam struct {
		// The teamID of the Away team in the match. (Negative for street
		// teams.)
		ID   id.Team `xml:"AwayTeamID"`
		Name string  `xml:"AwayTeamName"`
	} `xml:"AwayTeam"`

	Arena struct {
		ID   id.Arena `xml:"ArenaID"`
		Name string   `xml:"ArenaName"`
	} `xml:"Arena"`

	// The kick-off date and time of the match.
	MatchDate HattrickTime `xml:"MatchDate"`

	MatchType MatchType `xml:"MatchType"`

	// The team attitude specified for the match. Only shown to the owner
	// of the team - otherwise omitted.
	Attitude MatchTeamAttitude `xml:"Attitude"`

	TacticType MatchTacticType `xml:"TacticType"`

	// The style of play the team will use in the match, ranging from -10
	// (100% defensive) to 10 (100% offensive). CoachModifier.Available
	// specifies whether the value has been set or not.
	CoachModifier MatchOrdersCoachModifier `xml:"CoachModifier"`

	Lineup MatchOrdersLineup `xml:"Lineup"`

	// The submitted playerorders (substitutions, behaviour changes and
	// position swaps) of the team.
	PlayerOrders []*PlayerOrder `xml:"PlayerOrders>PlayerOrder"`
}

// MatchOrdersCoachModifier holds the CoachModifier value together with an
// Available flag indicating whether the value has actually been set.
type MatchOrdersCoachModifier struct {
	Available bool          `xml:"Available,attr"`
	Value     CoachModifier `xml:",chardata"`
}

// MatchOrdersLineup is the submitted lineup for a match: the players on
// the field, on the bench, taking penalties, plus set pieces and captain.
type MatchOrdersLineup struct {
	// The players on the field.
	Positions []*MatchOrdersFieldPlayer `xml:"Positions>Player"`

	// The players on the bench.
	Bench []*MatchOrdersPlayer `xml:"Bench>Player"`

	// The penalty takers.
	Kickers []*MatchOrdersPlayer `xml:"Kickers>Player"`

	SetPieces *MatchOrdersNamedPlayer `xml:"SetPieces"`
	Captain   *MatchOrdersNamedPlayer `xml:"Captain"`
}

// MatchOrdersPlayer is a player listed in the lineup (Bench or Kickers),
// together with the formal role slot they filled.
type MatchOrdersPlayer struct {
	PlayerID id.Player `xml:"PlayerID"`

	// An integer indicating which formal 'slot' (Role) the player has
	// filled in the match, before the repositionings take effect. There
	// can only be one (or zero) player with a particular RoleID,
	// contrary to how PositionCode works.
	RoleID MatchRole `xml:"RoleID"`

	FirstName string `xml:"FirstName"`
	NickName  string `xml:"NickName"`
	LastName  string `xml:"LastName"`
}

// MatchOrdersFieldPlayer is a MatchOrdersPlayer on the field, together with
// the individual order/repositioning they played with.
type MatchOrdersFieldPlayer struct {
	MatchOrdersPlayer

	// The individual order or repositioning that the player has played
	// with.
	Behaviour MatchBehaviourID `xml:"Behaviour"`
}

// MatchOrdersNamedPlayer is a player identified without a formal role slot
// (used for SetPieces and Captain).
type MatchOrdersNamedPlayer struct {
	PlayerID  id.Player `xml:"PlayerID"`
	FirstName string    `xml:"FirstName"`
	NickName  string    `xml:"NickName"`
	LastName  string    `xml:"LastName"`
}

// PlayerOrder describes a single submitted player order: a substitution,
// behaviour change or position swap.
type PlayerOrder struct {
	// The globally unique PlayerOrderID.
	PlayerOrderID uint `xml:"PlayerOrderID"`

	// The minute the order should be executed at the earliest.
	MatchMinuteCriteria LineupMatchMinuteCriteria `xml:"MatchMinuteCriteria"`

	// The standing also required to trigger the order.
	GoalDiffCriteria LineupGoalDiffCriteria `xml:"GoalDiffCriteria"`

	// The red card event also required to trigger the order.
	RedCardCriteria LineupRedCardCriteria `xml:"RedCardCriteria"`

	// If a/ substitution: the player leaving. If b/ behaviour change: the
	// player changing behaviour. If c/ position swap: the first player
	// that will change position.
	SubjectPlayerID id.Player `xml:"SubjectPlayerID"`

	// The player entering (or, once more, the player changing behaviour).
	// If position swap, the player to swap with.
	ObjectPlayerID id.Player `xml:"ObjectPlayerID"`

	OrderType MatchOrderType `xml:"OrderType"`

	// The new position for the (entering or staying) player.
	NewPositionID MatchRole `xml:"NewPositionId"`

	// The new behaviour for the (entering or staying) player.
	NewPositionBehaviour MatchBehaviourID `xml:"NewPositionBehaviour"`

	// Always 0; not used by the current match engine.
	PlayerOrderExtraInteger int `xml:"PlayerOrderExtraInteger"`
}
