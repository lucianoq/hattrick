package chpp

// MatchTeamAttitude ...
type MatchTeamAttitude int

// List of MatchTeamAttitude constants.
const (
	MatchTeamAttitudePlayItCool       MatchTeamAttitude = -1
	MatchTeamAttitudeNormal           MatchTeamAttitude = 0
	MatchTeamAttitudeMatchOfTheSeason MatchTeamAttitude = 1
)

// String returns a string representation of the type.
func (a MatchTeamAttitude) String() string {
	switch a {
	case MatchTeamAttitudePlayItCool:
		return "PIC"
	case MatchTeamAttitudeNormal:
		return "normal"
	case MatchTeamAttitudeMatchOfTheSeason:
		return "MOTS"
	default:
		return "unknown"
	}
}
