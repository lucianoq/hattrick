package chpp

// TrainerType is a trainer's tactical leaning (defensive, offensive, or
// balanced), which affects the team's training bonus.
type TrainerType uint

// List of TrainerType constants.
const (
	TrainerTypeDefensive TrainerType = 0
	TrainerTypeOffensive TrainerType = 1
	TrainerTypeBalanced  TrainerType = 2
)
