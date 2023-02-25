package chpp

// AchievementCategory ...
type AchievementCategory uint

// List of AchievementCategory constants.
const (
	AchievementCategoryRanking       AchievementCategory = 1
	AchievementCategoryTeam          AchievementCategory = 2
	AchievementCategoryMatches       AchievementCategory = 3
	AchievementCategoryManager       AchievementCategory = 4
	AchievementCategorySpecialAwards AchievementCategory = 5
	AchievementCategorySupporter     AchievementCategory = 6
)

func (c AchievementCategory) String() string {
	switch c {
	case AchievementCategoryRanking:
		return "ranking"
	case AchievementCategoryTeam:
		return "team"
	case AchievementCategoryMatches:
		return "matches"
	case AchievementCategoryManager:
		return "manager"
	case AchievementCategorySpecialAwards:
		return "special_awards"
	case AchievementCategorySupporter:
		return "supporter"
	}
	return "unknown achievement"
}
