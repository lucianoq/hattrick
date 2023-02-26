package id

// CupLevelIndex ...
type CupLevelIndex uint

const (
	CupLevelIndexEmerald  CupLevelIndex = 1
	CupLevelIndexRuby     CupLevelIndex = 2
	CupLevelIndexSapphire CupLevelIndex = 3
)

// String returns a string representation of the type.
func (l CupLevelIndex) String() string {
	switch l {
	case CupLevelIndexEmerald:
		return "emerald"
	case CupLevelIndexRuby:
		return "ruby"
	case CupLevelIndexSapphire:
		return "sapphire"
	}
	return "unknown"
}
