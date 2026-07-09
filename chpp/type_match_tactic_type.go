package chpp

import "strconv"

// MatchTacticType is the special tactic order set for a team in a match
// (e.g. pressing, counter-attacks, long shots).
type MatchTacticType uint

// List of MatchTacticType constants.
const (
	MatchTacticTypeNormal            MatchTacticType = 0
	MatchTacticTypePressing          MatchTacticType = 1
	MatchTacticTypeCounterAttacks    MatchTacticType = 2
	MatchTacticTypeAttackInTheMiddle MatchTacticType = 3
	MatchTacticTypeAttackInWings     MatchTacticType = 4
	MatchTacticTypePlayCreatively    MatchTacticType = 7
	MatchTacticTypeLongShots         MatchTacticType = 8
)

// String returns a string representation of the type.
func (t MatchTacticType) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
