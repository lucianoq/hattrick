// Package id holds strongly-typed identifiers for the various entities
// exposed by the CHPP API (teams, players, leagues, and so on), each a
// distinct type wrapping an underlying integer so they can't be mixed up
// with one another or with plain numeric parameters.
package id

import "strconv"

// AchievementID identifies a type of achievement (badge) that Hattrick can
// award to a manager.
type AchievementID uint

// String returns a string representation of the type.
func (a AchievementID) String() string {
	return strconv.FormatUint(uint64(a), 10)
}
