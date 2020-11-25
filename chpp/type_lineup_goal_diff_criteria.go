package chpp

import "strconv"

// LineupGoalDiffCriteria ...
type LineupGoalDiffCriteria int

// List of LineupGoalDiffCriteria constants.
const (
	LineupGoalDiffAny             LineupGoalDiffCriteria = -1
	LineupGoalDiffTied            LineupGoalDiffCriteria = 0
	LineupGoalDiffLead            LineupGoalDiffCriteria = 1
	LineupGoalDiffDown            LineupGoalDiffCriteria = 2
	LineupGoalDiffLeadByMoreThan1 LineupGoalDiffCriteria = 3
	LineupGoalDiffDownByMoreThan1 LineupGoalDiffCriteria = 4
	LineupGoalDiffNotDown         LineupGoalDiffCriteria = 5
	LineupGoalDiffNotLead         LineupGoalDiffCriteria = 6
	LineupGoalDiffLeadByMoreThan2 LineupGoalDiffCriteria = 7
	LineupGoalDiffDownByMoreThan2 LineupGoalDiffCriteria = 8
	LineupGoalDiffNotTied         LineupGoalDiffCriteria = 9
)

// String returns a string representation of the type.
func (x LineupGoalDiffCriteria) String() string {
	return strconv.Itoa(int(x))
}
