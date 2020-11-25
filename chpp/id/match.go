package id

import "strconv"

// Match ...
type Match uint

// String returns a string representation of the type.
func (m Match) String() string {
	return strconv.FormatUint(uint64(m), 10)
}
