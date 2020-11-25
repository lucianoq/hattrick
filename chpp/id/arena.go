package id

import "strconv"

// Arena ...
type Arena uint

// String returns a string representation of the type.
func (a Arena) String() string {
	return strconv.FormatUint(uint64(a), 10)
}
