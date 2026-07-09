package chpp

// StaffType is a non-player staff role at a club (e.g. assistant trainer,
// medic, financial director).
type StaffType uint

// List of StaffType constants.
const (
	StaffTypeAssistantTrainer  StaffType = 1
	StaffTypeMedic             StaffType = 2
	StaffTypeSpokesperson      StaffType = 3
	StaffTypeSportPsychologist StaffType = 4
	StaffTypeFormCoach         StaffType = 5
	StaffTypeFinancialDirector StaffType = 6
	StaffTypeTacticalAssistant StaffType = 7
)
