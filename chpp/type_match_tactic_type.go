package chpp

// MatchTacticType ...
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
	switch t {
	case MatchTacticTypeNormal:
		return "normal"
	case MatchTacticTypePressing:
		return "pressing"
	case MatchTacticTypeCounterAttacks:
		return "CA"
	case MatchTacticTypeAttackInTheMiddle:
		return "AIM"
	case MatchTacticTypeAttackInWings:
		return "AOW"
	case MatchTacticTypePlayCreatively:
		return "creatively"
	case MatchTacticTypeLongShots:
		return "longShots"
	default:
		return "unknown"
	}
}
