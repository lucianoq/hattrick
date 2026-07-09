package id

import "strconv"

// Staff identifies a member of a team's staff.
type Staff uint

// String returns a string representation of the type.
func (s Staff) String() string {
	return strconv.FormatUint(uint64(s), 10)
}
