package chpp

import "strconv"

// MatchOrderType ...
type MatchOrderType uint

// List of MatchOrderType constants.
const (
	MatchOrderTypeSubOrBehaviouralChange MatchOrderType = 1
	MatchOrderTypePlayerSwap             MatchOrderType = 3
)

// String returns a string representation of the type.
func (o MatchOrderType) String() string {
	return strconv.FormatUint(uint64(o), 10)
}
