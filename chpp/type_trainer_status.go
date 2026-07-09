package chpp

// TrainerStatus is whether a player also acts as the team's trainer, and
// if so whether they still play (playing trainer, only trainer, HoF trainer).
type TrainerStatus uint

// List of TrainerStatus constants.
const (
	TrainerStatusPlayingTrainer TrainerStatus = 1
	TrainerStatusOnlyTrainer    TrainerStatus = 2
	TrainerStatusHoFTrainer     TrainerStatus = 3
)
