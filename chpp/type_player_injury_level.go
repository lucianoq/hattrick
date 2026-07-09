package chpp

import "strconv"

// PlayerInjuryLevel is the number of weeks the player is predicted to be
// injured. If the player is bruised, the value is 0. If the player is
// healthy, the value is -1.
type PlayerInjuryLevel int

// List of special PlayerInjuryLevel values.
const (
	InjuryHealthy PlayerInjuryLevel = -1
	InjuryBruised PlayerInjuryLevel = 0
)

// String returns a string representation of the type.
func (i PlayerInjuryLevel) String() string {
	switch i {
	case InjuryHealthy:
		return ""
	case InjuryBruised:
		return "bruised"
	default:
		return "+" + strconv.Itoa(int(i))
	}
}

// Short returns a compact emoji-based representation of the injury level,
// for space-constrained display.
func (i PlayerInjuryLevel) Short() string {
	switch i {
	case InjuryHealthy:
		return ""
	case InjuryBruised:
		return "🤕"
	default:
		return "🏥" + strconv.Itoa(int(i))
	}
}
