package id

import "strconv"

// Series ...
type Series uint

// String returns a string representation of the type.
func (l Series) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
