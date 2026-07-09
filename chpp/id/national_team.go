package id

import "strconv"

// NationalTeam identifies a national team (including U-20 squads), as
// opposed to a regular club Team.
type NationalTeam uint

// String returns a string representation of the type.
func (n NationalTeam) String() string {
	return strconv.FormatUint(uint64(n), 10)
}
