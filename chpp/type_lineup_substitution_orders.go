package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// Hattrick matchorders JSON substitution order. ids and criteria are JSON
// numbers; only orderType is a string, matching the live CHPP endpoint.
type lineupSubstitutionOrder struct {
	PlayerIn  id.Player                 `json:"playerin"`
	PlayerOut id.Player                 `json:"playerout"`
	OrderType string                    `json:"orderType"`
	Minute    LineupMatchMinuteCriteria `json:"min"`
	Position  LineupPosition            `json:"pos"`
	Behaviour MatchBehaviourID          `json:"beh"`
	Card      LineupRedCardCriteria     `json:"card"`
	Standing  LineupGoalDiffCriteria    `json:"standing"`
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
