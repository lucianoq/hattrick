package id

import "strconv"

// Role ...
type Role int

// String returns a string representation of the type.
func (r Role) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
