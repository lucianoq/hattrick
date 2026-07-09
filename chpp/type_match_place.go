package chpp

import "strconv"

// MatchPlace is where a match is played, relative to a given team: at
// home, away, or at a neutral venue.
type MatchPlace uint

// List of MatchPlace constants.
const (
	MatchPlaceHome    MatchPlace = 0
	MatchPlaceAway    MatchPlace = 1
	MatchPlaceNeutral MatchPlace = 2
)

// String returns a string representation of the type.
func (m MatchPlace) String() string {
	return strconv.FormatUint(uint64(m), 10)
}
