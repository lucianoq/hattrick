package id

import "strconv"

// Match ...
type Match int

// String returns a string representation of the type.
func (m Match) String() string {
	return strconv.FormatInt(int64(m), 10)
}
