package id

import "strconv"

// Ladder identifies one of Hattrick's global team-ranking ladders.
type Ladder uint

// String returns a string representation of the type.
func (l Ladder) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
