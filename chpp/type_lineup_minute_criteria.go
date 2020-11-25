package chpp

import "strconv"

// LineupMatchMinuteCriteria ...
type LineupMatchMinuteCriteria int

// List of LineupMatchMinuteCriteria constants.
const (
	LineupMatchMinuteAnytime         LineupMatchMinuteCriteria = -1
	LineupMatchMinuteHalftime        LineupMatchMinuteCriteria = 46
	LineupMatchMinuteBeforeExtratime LineupMatchMinuteCriteria = 91
)

// String returns a string representation of the type.
func (x LineupMatchMinuteCriteria) String() string {
	return strconv.Itoa(int(x))
}
