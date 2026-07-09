package id

import "strconv"

// Role identifies a role that can be held by a member of an alliance
// (federation).
type Role int

// String returns a string representation of the type.
func (r Role) String() string {
	return strconv.Itoa(int(r))
}
