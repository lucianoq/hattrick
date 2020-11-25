package id

import "strconv"

// LeagueLevelUnit ...
type LeagueLevelUnit uint

// String returns a string representation of the type.
func (l LeagueLevelUnit) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
