package chpp

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// TSI is a player's Team Salary Index, a single number summarizing their
// overall market value/quality.
type TSI uint

// String returns the TSI humanized with thousands separators, right-padded
// to a width of 8 characters.
func (t TSI) String() string {
	return fmt.Sprintf("%8s", humanize.Comma(int64(t))) //nolint:gosec // TSI values are always far below int64's range
}
