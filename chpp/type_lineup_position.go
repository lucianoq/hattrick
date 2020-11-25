package chpp

import "strconv"

// LineupPosition ...
type LineupPosition int

// List of LineupPosition constants.
const (
	LineupPositionNoChange             LineupPosition = -1
	LineupPositionGoalkeeper           LineupPosition = 0
	LineupPositionRightDefender        LineupPosition = 1
	LineupPositionRightCentralDefender LineupPosition = 2
	LineupPositionCentralDefender      LineupPosition = 3
	LineupPositionLeftCentralDefender  LineupPosition = 4
	LineupPositionLeftDefender         LineupPosition = 5
	LineupPositionRightWinger          LineupPosition = 6
	LineupPositionRightInnerMidfielder LineupPosition = 7
	LineupPositionCentralMidfielder    LineupPosition = 8
	LineupPositionLeftInnerMidfielder  LineupPosition = 9
	LineupPositionLeftWinger           LineupPosition = 10
	LineupPositionRightForward         LineupPosition = 11
	LineupPositionCentralForward       LineupPosition = 12
	LineupPositionLeftForward          LineupPosition = 13
)

// String returns a string representation of the type.
func (x LineupPosition) String() string {
	return strconv.Itoa(int(x))
}
