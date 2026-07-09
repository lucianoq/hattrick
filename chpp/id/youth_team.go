package id

import "strconv"

// YouthTeam identifies a club's youth academy team.
type YouthTeam uint

// String returns a string representation of the type.
func (y YouthTeam) String() string {
	return strconv.FormatUint(uint64(y), 10)
}
