package id

import "strconv"

// Series identifies a League Level Unit: a single table/division within a
// League.
type Series uint

// String returns a string representation of the type.
func (l Series) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
