package chpp

// ScoutCommentSkillType is the skill a scout's comment about a player
// refers to (e.g. defending, scoring, passing).
type ScoutCommentSkillType uint

// List of ScoutCommentSkillType constants.
const (
	ScoutCommentSkillTypeKeeper ScoutCommentSkillType = 1
	// not in use = 2
	ScoutCommentSkillTypeDefending  ScoutCommentSkillType = 3
	ScoutCommentSkillTypePlaymaking ScoutCommentSkillType = 4
	ScoutCommentSkillTypeWinger     ScoutCommentSkillType = 5
	ScoutCommentSkillTypeScorer     ScoutCommentSkillType = 6
	ScoutCommentSkillTypeSetPieces  ScoutCommentSkillType = 7
	ScoutCommentSkillTypePassing    ScoutCommentSkillType = 8
)
