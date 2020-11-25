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

// LineupSubstitutionOrder ...
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
