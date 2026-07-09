package id

import "strconv"

// Match identifies a single Hattrick match. It is signed rather than
// unsigned because Hattrick sometimes returns -1 as a match ID.
type Match int

// String returns a string representation of the type.
func (m Match) String() string {
	return strconv.FormatInt(int64(m), 10)
}
