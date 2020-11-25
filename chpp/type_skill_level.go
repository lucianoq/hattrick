package chpp

// SkillLevel ...
type SkillLevel uint

// List of SkillLevel constants.
const (
	SkillLevelNonExistent      SkillLevel = 0
	SkillLevelDisastrous       SkillLevel = 1
	SkillLevelWretched         SkillLevel = 2
	SkillLevelPoor             SkillLevel = 3
	SkillLevelWeak             SkillLevel = 4
	SkillLevelInadequate       SkillLevel = 5
	SkillLevelPassable         SkillLevel = 6
	SkillLevelSolid            SkillLevel = 7
	SkillLevelExcellent        SkillLevel = 8
	SkillLevelFormidable       SkillLevel = 9
	SkillLevelOutstanding      SkillLevel = 10
	SkillLevelBrilliant        SkillLevel = 11
	SkillLevelMagnificent      SkillLevel = 12
	SkillLevelWorldClass       SkillLevel = 13
	SkillLevelSupernatural     SkillLevel = 14
	SkillLevelTitanic          SkillLevel = 15
	SkillLevelExtraTerrestrial SkillLevel = 16
	SkillLevelMythical         SkillLevel = 17
	SkillLevelMagical          SkillLevel = 18
	SkillLevelUtopian          SkillLevel = 19
	SkillLevelDivine           SkillLevel = 20
)

// String returns a string representation of the type.
func (s SkillLevel) String() string {
	switch s {
	case SkillLevelNonExistent:
		return "non-existent"
	case SkillLevelDisastrous:
		return "disastrous"
	case SkillLevelWretched:
		return "wretched"
	case SkillLevelPoor:
		return "poor"
	case SkillLevelWeak:
		return "weak"
	case SkillLevelInadequate:
		return "inadequate"
	case SkillLevelPassable:
		return "passable"
	case SkillLevelSolid:
		return "solid"
	case SkillLevelExcellent:
		return "excellent"
	case SkillLevelFormidable:
		return "formidable"
	case SkillLevelOutstanding:
		return "outstanding"
	case SkillLevelBrilliant:
		return "brilliant"
	case SkillLevelMagnificent:
		return "magnificent"
	case SkillLevelWorldClass:
		return "world class"
	case SkillLevelSupernatural:
		return "supernatural"
	case SkillLevelTitanic:
		return "titanic"
	case SkillLevelExtraTerrestrial:
		return "extra-terrestrial"
	case SkillLevelMythical:
		return "mythical"
	case SkillLevelMagical:
		return "magical"
	case SkillLevelUtopian:
		return "utopian"
	case SkillLevelDivine:
		return "divine"
	}
	return ""
}

// Value ...
func (s SkillLevel) Value() uint {
	return uint(s)
}
