package id

import "strconv"

// Language ...
type Language uint

// List of Language constants.
const (
	AllLanguages Language = 0
)

// String returns a string representation of the type.
func (l Language) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
