package id

import "strconv"

// Region identifies a geographic region within a country, e.g. where a
// team or arena is located.
type Region uint

// String returns a string representation of the type.
func (r Region) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
