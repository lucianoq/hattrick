package id

import "strconv"

// FriendlyMatch identifies a friendly (training) match challenge, as
// opposed to a competitive Match.
type FriendlyMatch uint

// String returns a string representation of the type.
func (t FriendlyMatch) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
