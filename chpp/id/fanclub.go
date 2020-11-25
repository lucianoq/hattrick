package id

import "strconv"

// Fanclub ...
type Fanclub uint

// String returns a string representation of the type.
func (f Fanclub) String() string {
	return strconv.FormatUint(uint64(f), 10)
}
