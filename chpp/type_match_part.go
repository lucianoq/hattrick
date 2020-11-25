package chpp

// MatchPart ...
type MatchPart uint

// List of MatchPart constants.
const (
	MatchPartBeforeStart MatchPart = 0
	MatchPartFirstHalf   MatchPart = 1
	MatchPartSecondHalf  MatchPart = 2
	MatchPartOvertime    MatchPart = 3
	MatchPartPenalty     MatchPart = 4
)
