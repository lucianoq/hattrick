package chpp

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// TSI ...
type TSI uint

// String returns a string representation of the type.
func (t TSI) String() string {
	return fmt.Sprintf("%8s", humanize.Comma(int64(t)))
}
