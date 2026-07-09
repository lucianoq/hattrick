package chpp

// TrainingType is which skill (or combination of skills/positions) a
// team's training session focuses on.
type TrainingType uint

// List of TrainingType constants.
const (
	TrainingTypeGeneral                          TrainingType = 0 // Deprecated
	TrainingTypeStamina                          TrainingType = 1 // Deprecated
	TrainingTypeSetPieces                        TrainingType = 2
	TrainingTypeDefending                        TrainingType = 3
	TrainingTypeScoring                          TrainingType = 4
	TrainingTypeWinger                           TrainingType = 5
	TrainingTypeScoringAndSetPieces              TrainingType = 6
	TrainingTypePassing                          TrainingType = 7
	TrainingTypePlaymaking                       TrainingType = 8
	TrainingTypeKeeper                           TrainingType = 9
	TrainingTypePassingDefendersAndMidfielders   TrainingType = 10
	TrainingTypeDefendingDefendersAndMidfielders TrainingType = 11
	TrainingTypeWingerAndAttackers               TrainingType = 12
)
