package id

import "strconv"

// Player ...
type Player uint

const PlayerUnknown Player = 0

// String returns a string representation of the type.
func (p Player) String() string {
	return strconv.FormatUint(uint64(p), 10)
}
