package id

import "strconv"

// Transfer identifies a player transfer-market listing.
type Transfer uint

// String returns a string representation of the type.
func (t Transfer) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
