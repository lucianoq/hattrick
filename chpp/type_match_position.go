package chpp

// MatchPosition is one of the 11 formal field positions in Hattrick's
// standard 5-5-3 lineup grid.
type MatchPosition uint

// List of MatchPosition constants.
const (
	MatchPositionKeeper           MatchPosition = 1
	MatchPositionRightBack        MatchPosition = 2
	MatchPositionCentralDefender1 MatchPosition = 3
	MatchPositionCentralDefender2 MatchPosition = 4
	MatchPositionLeftBack         MatchPosition = 5
	MatchPositionRightWinger      MatchPosition = 6
	MatchPositionInnerMidfield1   MatchPosition = 7
	MatchPositionInnerMidfield2   MatchPosition = 8
	MatchPositionLeftWinger       MatchPosition = 9
	MatchPositionForward1         MatchPosition = 10
	MatchPositionForward2         MatchPosition = 11
)
