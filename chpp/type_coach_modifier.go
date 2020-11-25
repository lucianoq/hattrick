package chpp

import "strconv"

// CoachModifier ...
type CoachModifier int

// List of CoachModifier constants.
const (
	CoachModifierDefensive100 CoachModifier = -10
	CoachModifierDefensive90  CoachModifier = -9
	CoachModifierDefensive80  CoachModifier = -8
	CoachModifierDefensive70  CoachModifier = -7
	CoachModifierDefensive60  CoachModifier = -6
	CoachModifierDefensive50  CoachModifier = -5
	CoachModifierDefensive40  CoachModifier = -4
	CoachModifierDefensive30  CoachModifier = -3
	CoachModifierDefensive20  CoachModifier = -2
	CoachModifierDefensive10  CoachModifier = -1
	CoachModifierNeutral      CoachModifier = 0
	CoachModifierOffensive10  CoachModifier = 1
	CoachModifierOffensive20  CoachModifier = 2
	CoachModifierOffensive30  CoachModifier = 3
	CoachModifierOffensive40  CoachModifier = 4
	CoachModifierOffensive50  CoachModifier = 5
	CoachModifierOffensive60  CoachModifier = 6
	CoachModifierOffensive70  CoachModifier = 7
	CoachModifierOffensive80  CoachModifier = 8
	CoachModifierOffensive90  CoachModifier = 9
	CoachModifierOffensive100 CoachModifier = 10
)

// String returns a string representation of the type.
func (x CoachModifier) String() string {
	return strconv.Itoa(int(x))
}
