package id

import "strconv"

// YouthLeague ...
type YouthLeague uint

// String returns a string representation of the type.
func (y YouthLeague) String() string {
	return strconv.FormatUint(uint64(y), 10)
}
