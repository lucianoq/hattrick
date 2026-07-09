package id

import "strconv"

// Cup identifies a Hattrick cup competition.
type Cup uint

// HattrickMasters is the Cup ID of the Hattrick Masters competition.
const HattrickMasters Cup = 183

// String returns a string representation of the type.
func (c Cup) String() string {
	return strconv.FormatUint(uint64(c), 10)
}
