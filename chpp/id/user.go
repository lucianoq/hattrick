package id

import "strconv"

// User identifies a Hattrick manager (user account).
type User uint

// String returns a string representation of the type.
func (u User) String() string {
	return strconv.FormatUint(uint64(u), 10)
}
