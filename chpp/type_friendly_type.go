package chpp

import "strconv"

// FriendlyType ...
type FriendlyType uint

// List of FriendlyType constants.
const (
	FriendlyTypeNormalRules          FriendlyType = 0
	FriendlyTypeCupRules             FriendlyType = 1
	FriendlyTypeNationalTeamFriendly FriendlyType = 12
)

// String returns a string representation of the type.
func (m FriendlyType) String() string {
	return strconv.FormatUint(uint64(m), 10)
}
