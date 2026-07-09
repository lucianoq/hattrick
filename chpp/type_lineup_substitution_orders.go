package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// Hattrick string-based json
type lineupSubstitutionOrder struct {
	PlayerIn  string `json:"playerin"`
	PlayerOut string `json:"playerout"`
	OrderType string `json:"orderType"`
	Minute    string `json:"min"`
	Position  string `json:"pos"`
	Behaviour string `json:"beh"`
	Card      string `json:"card"`
	Standing  string `json:"standing"`
}

// LineupSubstitutionOrder is a single player order (substitution,
// behaviour change, or position swap) to include in the matchorders
// lineup JSON payload built by NewLineup.
type LineupSubstitutionOrder struct {
	PlayerIn  id.Player
	PlayerOut id.Player // same as In if behaviour change
	OrderType MatchOrderType
	Minute    LineupMatchMinuteCriteria
	Position  LineupPosition
	Behaviour MatchBehaviourID
	Card      LineupRedCardCriteria
	Standing  LineupGoalDiffCriteria
}
