package chpp

// NationalTeamStaffType ...
type NationalTeamStaffType uint

// List of NationalTeamStaffType constants.
const (
	NationalTeamCoach NationalTeamStaffType = 0
	AssistantCoach    NationalTeamStaffType = 1
	Scout             NationalTeamStaffType = 2
)

// String returns a string representation of the type.
func (ntst NationalTeamStaffType) String() string {
	switch ntst {
	case NationalTeamCoach:
		return "national team coach"
	case AssistantCoach:
		return "assistant coach"
	case Scout:
		return "scout"
	default:
		return ""
	}
}
