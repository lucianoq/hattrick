package id

import "strconv"

// YouthLeague identifies a youth-team competition, the youth-academy
// equivalent of a League.
type YouthLeague uint

// String returns a string representation of the type.
func (y YouthLeague) String() string {
	return strconv.FormatUint(uint64(y), 10)
}
