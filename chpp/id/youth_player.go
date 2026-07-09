package id

import "strconv"

// YouthPlayer identifies a player in a team's youth academy.
type YouthPlayer uint

// String returns a string representation of the type.
func (y YouthPlayer) String() string {
	return strconv.FormatUint(uint64(y), 10)
}
