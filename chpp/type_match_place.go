package chpp

import "strconv"

// MatchPlace ...
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
