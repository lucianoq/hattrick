package chpp

// MatchRating ...
type MatchRating uint

// List of MatchRating constants.
const (
	MatchRatingVeryLowDisastrous        MatchRating = 1
	MatchRatingLowDisastrous            MatchRating = 2
	MatchRatingHighDisastrous           MatchRating = 3
	MatchRatingVeryHighDisastrous       MatchRating = 4
	MatchRatingVeryLowWretched          MatchRating = 5
	MatchRatingLowWretched              MatchRating = 6
	MatchRatingHighWretched             MatchRating = 7
	MatchRatingVeryHighWretched         MatchRating = 8
	MatchRatingVeryLowPoor              MatchRating = 9
	MatchRatingLowPoor                  MatchRating = 10
	MatchRatingHighPoor                 MatchRating = 11
	MatchRatingVeryHighPoor             MatchRating = 12
	MatchRatingVeryLowWeak              MatchRating = 13
	MatchRatingLowWeak                  MatchRating = 14
	MatchRatingHighWeak                 MatchRating = 15
	MatchRatingVeryHighWeak             MatchRating = 16
	MatchRatingVeryLowInadequate        MatchRating = 17
	MatchRatingLowInadequate            MatchRating = 18
	MatchRatingHighInadequate           MatchRating = 19
	MatchRatingVeryHighInadequate       MatchRating = 20
	MatchRatingVeryLowPassable          MatchRating = 21
	MatchRatingLowPassable              MatchRating = 22
	MatchRatingHighPassable             MatchRating = 23
	MatchRatingVeryHighPassable         MatchRating = 24
	MatchRatingVeryLowSolid             MatchRating = 25
	MatchRatingLowSolid                 MatchRating = 26
	MatchRatingHighSolid                MatchRating = 27
	MatchRatingVeryHighSolid            MatchRating = 28
	MatchRatingVeryLowExcellent         MatchRating = 29
	MatchRatingLowExcellent             MatchRating = 30
	MatchRatingHighExcellent            MatchRating = 31
	MatchRatingVeryHighExcellent        MatchRating = 32
	MatchRatingVeryLowFormidable        MatchRating = 33
	MatchRatingLowFormidable            MatchRating = 34
	MatchRatingHighFormidable           MatchRating = 35
	MatchRatingVeryHighFormidable       MatchRating = 36
	MatchRatingVeryLowOutstanding       MatchRating = 37
	MatchRatingLowOutstanding           MatchRating = 38
	MatchRatingHighOutstanding          MatchRating = 39
	MatchRatingVeryHighOutstanding      MatchRating = 40
	MatchRatingVeryLowBrilliant         MatchRating = 41
	MatchRatingLowBrilliant             MatchRating = 42
	MatchRatingHighBrilliant            MatchRating = 43
	MatchRatingVeryHighBrilliant        MatchRating = 44
	MatchRatingVeryLowMagnificent       MatchRating = 45
	MatchRatingLowMagnificent           MatchRating = 46
	MatchRatingHighMagnificent          MatchRating = 47
	MatchRatingVeryHighMagnificent      MatchRating = 48
	MatchRatingVeryLowWorldClass        MatchRating = 49
	MatchRatingLowWorldClass            MatchRating = 50
	MatchRatingHighWorldClass           MatchRating = 51
	MatchRatingVeryHighWorldClass       MatchRating = 52
	MatchRatingVeryLowSupernatural      MatchRating = 53
	MatchRatingLowSupernatural          MatchRating = 54
	MatchRatingHighSupernatural         MatchRating = 55
	MatchRatingVeryHighSupernatural     MatchRating = 56
	MatchRatingVeryLowTitanic           MatchRating = 57
	MatchRatingLowTitanic               MatchRating = 58
	MatchRatingHighTitanic              MatchRating = 59
	MatchRatingVeryHighTitanic          MatchRating = 60
	MatchRatingVeryLowExtraTerrestrial  MatchRating = 61
	MatchRatingLowExtraTerrestrial      MatchRating = 62
	MatchRatingHighExtraTerrestrial     MatchRating = 63
	MatchRatingVeryHighExtraTerrestrial MatchRating = 64
	MatchRatingVeryLowMythical          MatchRating = 65
	MatchRatingLowMythical              MatchRating = 66
	MatchRatingHighMythical             MatchRating = 67
	MatchRatingVeryHighMythical         MatchRating = 68
	MatchRatingVeryLowMagical           MatchRating = 69
	MatchRatingLowMagical               MatchRating = 70
	MatchRatingHighMagical              MatchRating = 71
	MatchRatingVeryHighMagical          MatchRating = 72
	MatchRatingVeryLowUtopian           MatchRating = 73
	MatchRatingLowUtopian               MatchRating = 74
	MatchRatingHighUtopian              MatchRating = 75
	MatchRatingVeryHighUtopian          MatchRating = 76
	MatchRatingVeryLowDivine            MatchRating = 77
	MatchRatingLowDivine                MatchRating = 78
	MatchRatingHighDivine               MatchRating = 79
	MatchRatingVeryHighDivine           MatchRating = 80
)

// String returns a string representation of the type.
// nolint
func (m MatchRating) String() string {
	switch m {
	case MatchRatingVeryLowDisastrous:
		return "VeryLowDisastrous"
	case MatchRatingLowDisastrous:
		return "LowDisastrous"
	case MatchRatingHighDisastrous:
		return "HighDisastrous"
	case MatchRatingVeryHighDisastrous:
		return "VeryHighDisastrous"
	case MatchRatingVeryLowWretched:
		return "VeryLowWretched"
	case MatchRatingLowWretched:
		return "LowWretched"
	case MatchRatingHighWretched:
		return "HighWretched"
	case MatchRatingVeryHighWretched:
		return "VeryHighWretched"
	case MatchRatingVeryLowPoor:
		return "VeryLowPoor"
	case MatchRatingLowPoor:
		return "LowPoor"
	case MatchRatingHighPoor:
		return "HighPoor"
	case MatchRatingVeryHighPoor:
		return "VeryHighPoor"
	case MatchRatingVeryLowWeak:
		return "VeryLowWeak"
	case MatchRatingLowWeak:
		return "LowWeak"
	case MatchRatingHighWeak:
		return "HighWeak"
	case MatchRatingVeryHighWeak:
		return "VeryHighWeak"
	case MatchRatingVeryLowInadequate:
		return "VeryLowInadequate"
	case MatchRatingLowInadequate:
		return "LowInadequate"
	case MatchRatingHighInadequate:
		return "HighInadequate"
	case MatchRatingVeryHighInadequate:
		return "VeryHighInadequate"
	case MatchRatingVeryLowPassable:
		return "VeryLowPassable"
	case MatchRatingLowPassable:
		return "LowPassable"
	case MatchRatingHighPassable:
		return "HighPassable"
	case MatchRatingVeryHighPassable:
		return "VeryHighPassable"
	case MatchRatingVeryLowSolid:
		return "VeryLowSolid"
	case MatchRatingLowSolid:
		return "LowSolid"
	case MatchRatingHighSolid:
		return "HighSolid"
	case MatchRatingVeryHighSolid:
		return "VeryHighSolid"
	case MatchRatingVeryLowExcellent:
		return "VeryLowExcellent"
	case MatchRatingLowExcellent:
		return "LowExcellent"
	case MatchRatingHighExcellent:
		return "HighExcellent"
	case MatchRatingVeryHighExcellent:
		return "VeryHighExcellent"
	case MatchRatingVeryLowFormidable:
		return "VeryLowFormidable"
	case MatchRatingLowFormidable:
		return "LowFormidable"
	case MatchRatingHighFormidable:
		return "HighFormidable"
	case MatchRatingVeryHighFormidable:
		return "VeryHighFormidable"
	case MatchRatingVeryLowOutstanding:
		return "VeryLowOutstanding"
	case MatchRatingLowOutstanding:
		return "LowOutstanding"
	case MatchRatingHighOutstanding:
		return "HighOutstanding"
	case MatchRatingVeryHighOutstanding:
		return "VeryHighOutstanding"
	case MatchRatingVeryLowBrilliant:
		return "VeryLowBrilliant"
	case MatchRatingLowBrilliant:
		return "LowBrilliant"
	case MatchRatingHighBrilliant:
		return "HighBrilliant"
	case MatchRatingVeryHighBrilliant:
		return "VeryHighBrilliant"
	case MatchRatingVeryLowMagnificent:
		return "VeryLowMagnificent"
	case MatchRatingLowMagnificent:
		return "LowMagnificent"
	case MatchRatingHighMagnificent:
		return "HighMagnificent"
	case MatchRatingVeryHighMagnificent:
		return "VeryHighMagnificent"
	case MatchRatingVeryLowWorldClass:
		return "VeryLowWorldClass"
	case MatchRatingLowWorldClass:
		return "LowWorldClass"
	case MatchRatingHighWorldClass:
		return "HighWorldClass"
	case MatchRatingVeryHighWorldClass:
		return "VeryHighWorldClass"
	case MatchRatingVeryLowSupernatural:
		return "VeryLowSupernatural"
	case MatchRatingLowSupernatural:
		return "LowSupernatural"
	case MatchRatingHighSupernatural:
		return "HighSupernatural"
	case MatchRatingVeryHighSupernatural:
		return "VeryHighSupernatural"
	case MatchRatingVeryLowTitanic:
		return "VeryLowTitanic"
	case MatchRatingLowTitanic:
		return "LowTitanic"
	case MatchRatingHighTitanic:
		return "HighTitanic"
	case MatchRatingVeryHighTitanic:
		return "VeryHighTitanic"
	case MatchRatingVeryLowExtraTerrestrial:
		return "VeryLowExtraTerrestrial"
	case MatchRatingLowExtraTerrestrial:
		return "LowExtraTerrestrial"
	case MatchRatingHighExtraTerrestrial:
		return "HighExtraTerrestrial"
	case MatchRatingVeryHighExtraTerrestrial:
		return "VeryHighExtraTerrestrial "
	case MatchRatingVeryLowMythical:
		return "VeryLowMythical"
	case MatchRatingLowMythical:
		return "LowMythical"
	case MatchRatingHighMythical:
		return "HighMythical"
	case MatchRatingVeryHighMythical:
		return "VeryHighMythical"
	case MatchRatingVeryLowMagical:
		return "VeryLowMagical"
	case MatchRatingLowMagical:
		return "LowMagical"
	case MatchRatingHighMagical:
		return "HighMagical"
	case MatchRatingVeryHighMagical:
		return "VeryHighMagical"
	case MatchRatingVeryLowUtopian:
		return "VeryLowUtopian"
	case MatchRatingLowUtopian:
		return "LowUtopian"
	case MatchRatingHighUtopian:
		return "HighUtopian"
	case MatchRatingVeryHighUtopian:
		return "VeryHighUtopian"
	case MatchRatingVeryLowDivine:
		return "VeryLowDivine"
	case MatchRatingLowDivine:
		return "LowDivine"
	case MatchRatingHighDivine:
		return "HighDivine"
	case MatchRatingVeryHighDivine:
		return "VeryHighDivine"
	default:
		return "unknown"
	}
}

// Value ...
func (m MatchRating) Value() int {
	return int(m)
}
