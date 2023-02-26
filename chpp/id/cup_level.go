package id

// CupLevel ...
type CupLevel uint

const (
	CupLevelNational    CupLevel = 1
	CupLevelChallenger  CupLevel = 2
	CupLevelConsolation CupLevel = 3
)

// String returns a string representation of the type.
func (l CupLevel) String() string {
	switch l {
	case CupLevelNational:
		return "national"
	case CupLevelChallenger:
		return "challenger"
	case CupLevelConsolation:
		return "consolation"
	}
	return "unknown"
}
