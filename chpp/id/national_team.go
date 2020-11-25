package id

import "strconv"

// NationalTeam ...
type NationalTeam uint

// String returns a string representation of the type.
func (n NationalTeam) String() string {
	return strconv.FormatUint(uint64(n), 10)
}
