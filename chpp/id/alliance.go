package id

import "strconv"

// Alliance identifies a Hattrick alliance (federation), a group of teams.
type Alliance uint

// String returns a string representation of the type.
func (a Alliance) String() string {
	return strconv.FormatUint(uint64(a), 10)
}
