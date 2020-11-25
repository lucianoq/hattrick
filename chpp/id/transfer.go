package id

import "strconv"

// Transfer ...
type Transfer uint

// String returns a string representation of the type.
func (t Transfer) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
