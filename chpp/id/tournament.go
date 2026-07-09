package id

import "strconv"

// Tournament identifies a Hattrick tournament, a competition distinct from
// the regular league and cup structure.
type Tournament uint

// String returns a string representation of the type.
func (t Tournament) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
