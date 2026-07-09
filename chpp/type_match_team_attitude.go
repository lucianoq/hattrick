package chpp

import "strconv"

// MatchTeamAttitude is the team's motivational speech level for a match,
// ranging from playing it cool (-1) to going for match of the season (1).
type MatchTeamAttitude int

// List of MatchTeamAttitude constants.
const (
	MatchTeamAttitudePlayItCool       MatchTeamAttitude = -1
	MatchTeamAttitudeNormal           MatchTeamAttitude = 0
	MatchTeamAttitudeMatchOfTheSeason MatchTeamAttitude = 1
)

// String returns a string representation of the type.
func (a MatchTeamAttitude) String() string {
	return strconv.Itoa(int(a))
}
