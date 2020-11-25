package id

import "strconv"

// Bookmark ...
type Bookmark uint

// String returns a string representation of the type.
func (b Bookmark) String() string {
	return strconv.FormatUint(uint64(b), 10)
}
