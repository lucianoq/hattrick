package id

import "strconv"

// Country ...
type Country uint

// String returns a string representation of the type.
func (c Country) String() string {
	return strconv.FormatUint(uint64(c), 10)
}
