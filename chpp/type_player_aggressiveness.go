package chpp

// PlayerAggressiveness ...
type PlayerAggressiveness uint

// List of PlayerAggressiveness constants.
const (
	PlayerAggressivenessTranquil      PlayerAggressiveness = 0
	PlayerAggressivenessCalm          PlayerAggressiveness = 1
	PlayerAggressivenessBalanced      PlayerAggressiveness = 2
	PlayerAggressivenessTemperamental PlayerAggressiveness = 3
	PlayerAggressivenessFiery         PlayerAggressiveness = 4
	PlayerAggressivenessUnstable      PlayerAggressiveness = 5
)

// String returns a string representation of the type.
func (a PlayerAggressiveness) String() string {
	switch a {
	case PlayerAggressivenessTranquil:
		return "tranquil"
	case PlayerAggressivenessCalm:
		return "calm"
	case PlayerAggressivenessBalanced:
		return "balanced"
	case PlayerAggressivenessTemperamental:
		return "temperamental"
	case PlayerAggressivenessFiery:
		return "fiery"
	case PlayerAggressivenessUnstable:
		return "unstable"
	}

	return ""
}
