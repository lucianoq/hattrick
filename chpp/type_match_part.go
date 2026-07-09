package chpp

// MatchPart identifies which part of a match an event happened in (before
// kickoff, first/second half, overtime, or penalty shootout).
type MatchPart uint

// List of MatchPart constants.
const (
	MatchPartBeforeStart MatchPart = 0
	MatchPartFirstHalf   MatchPart = 1
	MatchPartSecondHalf  MatchPart = 2
	MatchPartOvertime    MatchPart = 3
	MatchPartPenalty     MatchPart = 4
)
