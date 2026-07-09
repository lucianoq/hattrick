package id

import "strconv"

// Team identifies a Hattrick team.
type Team int

// String returns a string representation of the type.
func (t Team) String() string {
	return strconv.Itoa(int(t))
}
