package chpp

// SkillID identifies one of a player's trainable skills (e.g. defending,
// scoring, passing) or attributes (e.g. stamina, experience, leadership).
type SkillID uint

// List of SkillID constants.
const (
	SkillIDGoaltending SkillID = 1
	SkillIDStamina     SkillID = 2
	SkillIDSetPieces   SkillID = 3
	SkillIDDefending   SkillID = 4
	SkillIDScoring     SkillID = 5
	SkillIDWinger      SkillID = 6
	SkillIDPassing     SkillID = 7
	SkillIDPlaymaking  SkillID = 8
	SkillIDTrainer     SkillID = 9
	SkillIDLeadership  SkillID = 10
	SkillIDExperience  SkillID = 11
)
