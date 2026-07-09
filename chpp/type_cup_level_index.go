package chpp

// CupLevelIndex identifies which of the parallel cups at the same CupLevel
// a team is in (Emerald, Ruby, or Sapphire).
type CupLevelIndex uint

// List of CupLevelIndex constants.
const (
	CupLevelIndexNotCup   CupLevelIndex = 0
	CupLevelIndexEmerald  CupLevelIndex = 1
	CupLevelIndexRuby     CupLevelIndex = 2
	CupLevelIndexSapphire CupLevelIndex = 3
)
