package id

import "strconv"

// FriendlyMatch ...
type FriendlyMatch uint

// String returns a string representation of the type.
func (t FriendlyMatch) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
