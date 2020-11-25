package id

import "strconv"

// League ...
type League uint

// String returns a string representation of the type.
func (l League) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
