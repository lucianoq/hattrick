package chpp

// SourceSystem identifies which Hattrick sub-system a team/player belongs
// to (main game, youth system, or the ex-tournament integrated system).
type SourceSystem string

// List of SourceSystem constants.
const (
	SourceSystemHattrickMainSystem         SourceSystem = "hattrick"
	SourceSystemYouthSystem                SourceSystem = "youth"
	SourceSystemIntegratedForExTournaments SourceSystem = "htointegrated"
)
