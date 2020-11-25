package id

import "strconv"

// Referee ...
type Referee uint

// String returns a string representation of the type.
func (r Referee) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
