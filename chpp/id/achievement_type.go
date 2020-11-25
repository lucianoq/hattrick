package id

import "strconv"

// AchievementID ...
type AchievementID uint

// String returns a string representation of the type.
func (a AchievementID) String() string {
	return strconv.FormatUint(uint64(a), 10)
}
