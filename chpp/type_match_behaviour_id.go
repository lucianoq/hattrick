package chpp

import "strconv"

// MatchBehaviourID ...
type MatchBehaviourID int

// List of MatchBehaviourID constants.
const (
	MatchBehaviourNoChange           MatchBehaviourID = -1
	MatchBehaviourNormal             MatchBehaviourID = 0
	MatchBehaviourOffensive          MatchBehaviourID = 1
	MatchBehaviourDefensive          MatchBehaviourID = 2
	MatchBehaviourTowardsMiddle      MatchBehaviourID = 3
	MatchBehaviourTowardsWing        MatchBehaviourID = 4
	MatchBehaviourExtraForward       MatchBehaviourID = 5
	MatchBehaviourExtraInnerMidfield MatchBehaviourID = 6
	MatchBehaviourExtraDefender      MatchBehaviourID = 7
)

// String returns a string representation of the type.
func (b MatchBehaviourID) String() string {
	return strconv.Itoa(int(b))
}
