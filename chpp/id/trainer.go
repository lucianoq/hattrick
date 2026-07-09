package id

import "strconv"

// Trainer identifies a team's trainer (head coach), a specific kind of
// Staff member.
type Trainer uint

// String returns a string representation of the type.
func (t Trainer) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
