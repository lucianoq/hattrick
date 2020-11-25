package chpp

// ScoutCommentType ...
type ScoutCommentType uint

// List of ScoutCommentType constants.
const (
	ScoutCommentTypeHello               ScoutCommentType = 0
	ScoutCommentTypeFoundPlayer         ScoutCommentType = 1
	ScoutCommentTypeTalentedInOneSkill  ScoutCommentType = 3
	ScoutCommentTypeCurrentSkillLevel   ScoutCommentType = 4
	ScoutCommentTypePotentialSkillLevel ScoutCommentType = 5
	ScoutCommentTypeAverageSkillLevel   ScoutCommentType = 6
	ScoutCommentTypeSignPlayerToTeam    ScoutCommentType = 7
	ScoutCommentTypePlayerHasSpeciality ScoutCommentType = 9
)
