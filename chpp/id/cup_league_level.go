package id

import "strconv"

// CupLeagueLevel ...
type CupLeagueLevel uint

// String returns a string representation of the type.
func (l CupLeagueLevel) String() string {
	return strconv.FormatUint(uint64(l), 10)
}

// Type returns a string representation of the type.
func (l CupLeagueLevel) Type() string {
	switch l {
	case 1, 2, 3, 4, 5, 6:
		return "national"
	case 7, 8, 9:
		return "divisional"
	}
	return "unknown"
}
