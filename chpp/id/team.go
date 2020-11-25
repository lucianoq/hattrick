package id

import "strconv"

// Team ...
type Team int

// String returns a string representation of the type.
func (t Team) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
