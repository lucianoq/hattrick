package id

import "strconv"

// Region ...
type Region uint

// String returns a string representation of the type.
func (r Region) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
