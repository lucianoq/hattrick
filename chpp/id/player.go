package id

import "strconv"

// Player identifies a senior Hattrick player.
type Player uint

// PlayerUnknown is the zero value, used where a player ID is optional or
// not yet known.
const PlayerUnknown Player = 0

// String returns a string representation of the type.
func (p Player) String() string {
	return strconv.FormatUint(uint64(p), 10)
}
