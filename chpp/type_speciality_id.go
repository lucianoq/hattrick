package chpp

// SpecialtyID ...
type SpecialtyID uint

// List of SpecialtyID constants.
const (
	SpecialtyNoSpecialty    SpecialtyID = 0
	SpecialtyTechnical      SpecialtyID = 1
	SpecialtyQuick          SpecialtyID = 2
	SpecialtyPowerful       SpecialtyID = 3
	SpecialtyUnpredictable  SpecialtyID = 4
	SpecialtyHeadSpecialist SpecialtyID = 5
	SpecialtyResilient      SpecialtyID = 6
	SpecialtySupport        SpecialtyID = 8
)

// String returns a string representation of the type.
func (s SpecialtyID) String() string {
	switch s {
	case SpecialtyNoSpecialty:
		return ""
	case SpecialtyTechnical:
		return "Technical"
	case SpecialtyQuick:
		return "Quick"
	case SpecialtyPowerful:
		return "Powerful"
	case SpecialtyUnpredictable:
		return "Unpredictable"
	case SpecialtyHeadSpecialist:
		return "Head specialist"
	case SpecialtyResilient:
		return "Resilient"
	case SpecialtySupport:
		return "Support"
	}

	return ""
}

// StringTag ...
func (s SpecialtyID) StringTag() string {
	switch s {
	case SpecialtyNoSpecialty:
		return ""
	case SpecialtyTechnical:
		return "[Technical]"
	case SpecialtyQuick:
		return "[Quick]"
	case SpecialtyPowerful:
		return "[Powerful]"
	case SpecialtyUnpredictable:
		return "[Unpredict]"
	case SpecialtyHeadSpecialist:
		return "[Head]"
	case SpecialtyResilient:
		return "[Resilient]"
	case SpecialtySupport:
		return "[Supporter]"
	}

	return ""
}

// Emoji ...
func (s SpecialtyID) Emoji() string {
	switch s {
	case SpecialtyNoSpecialty:
		return ""
	case SpecialtyTechnical:
		return "⚽"
	case SpecialtyQuick:
		return "⚡"
	case SpecialtyPowerful:
		return "💪"
	case SpecialtyUnpredictable:
		return "🎩"
	case SpecialtyHeadSpecialist:
		return "👨"
	case SpecialtyResilient:
		return "" // TODO
	case SpecialtySupport:
		return "" // TODO
	}

	return ""
}
