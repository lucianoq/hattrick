package id

import "strconv"

// Cup ...
type Cup uint

// List of Cup constants.
const (
	HattrickMasters Cup = 183
)

// String returns a string representation of the type.
func (c Cup) String() string {
	return strconv.FormatUint(uint64(c), 10)
}
