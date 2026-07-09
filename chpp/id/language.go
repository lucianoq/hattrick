package id

import "strconv"

// Language identifies one of the languages supported by Hattrick.
type Language uint

// AllLanguages is a sentinel Language value meaning no language filter is
// applied (all languages).
const AllLanguages Language = 0

// String returns a string representation of the type.
func (l Language) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
