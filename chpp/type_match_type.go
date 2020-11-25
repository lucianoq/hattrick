package chpp

// MatchType ...
type MatchType uint

// List of MatchType constants.
const (
	MatchTypeLeagueMatch                                MatchType = 1
	MatchTypeQualificationMatch                         MatchType = 2
	MatchTypeCupMatch                                   MatchType = 3
	MatchTypeFriendlyNormalRules                        MatchType = 4
	MatchTypeFriendlyCupRules                           MatchType = 5
	MatchTypeHattrickMasters                            MatchType = 7
	MatchTypeInternationalFriendlyNormalRules           MatchType = 8
	MatchTypeInternationalFriendlyCupRules              MatchType = 9
	MatchTypeNationalTeamCompetitionMatchNormalRules    MatchType = 10
	MatchTypeNationalTeamCompetitionMatchCupRules       MatchType = 11
	MatchTypeNationalTeamsFriendly                      MatchType = 12
	MatchTypeTournamentLeagueMatch                      MatchType = 50
	MatchTypeTournamentPlayoffMatch                     MatchType = 51
	MatchTypeSingleMatch                                MatchType = 61
	MatchTypeLadderMatch                                MatchType = 62
	MatchTypePreparationMatch                           MatchType = 80
	MatchTypeYouthLeagueMatch                           MatchType = 100
	MatchTypeYouthFriendlyMatch                         MatchType = 101
	MatchTypeYouthFriendlyMatchCupRules                 MatchType = 103
	MatchTypeYouthInternationalFriendlyMatchNormalRules MatchType = 105
	MatchTypeYouthInternationalFriendlyMatchCupRules    MatchType = 106
)

// String returns a string representation of the type.
func (m MatchType) String() string {
	switch m {
	case MatchTypeLeagueMatch:
		return "league"
	case MatchTypeQualificationMatch:
		return "qualification"
	case MatchTypeCupMatch:
		return "cup"
	case MatchTypeFriendlyNormalRules:
		return "friendly_normal_rules"
	case MatchTypeFriendlyCupRules:
		return "friendly_cup_rules"
	case MatchTypeHattrickMasters:
		return "hattrick_masters"
	case MatchTypeInternationalFriendlyNormalRules:
		return "friendly_international_normal_rules"
	case MatchTypeInternationalFriendlyCupRules:
		return "friendly_international_cup_rules"
	case MatchTypeNationalTeamCompetitionMatchNormalRules:
		return "national_team_competition_normal_rules"
	case MatchTypeNationalTeamCompetitionMatchCupRules:
		return "national_team_competition_cup_rules"
	case MatchTypeNationalTeamsFriendly:
		return "national_team_friendly"
	case MatchTypeTournamentLeagueMatch:
		return "tournament_league"
	case MatchTypeTournamentPlayoffMatch:
		return "tournament_playoff"
	case MatchTypeSingleMatch:
		return "single_match"
	case MatchTypeLadderMatch:
		return "ladder_match"
	case MatchTypePreparationMatch:
		return "preparation_match"
	case MatchTypeYouthLeagueMatch:
		return "youth_league"
	case MatchTypeYouthFriendlyMatch:
		return "youth_friendly_normal_rules"
	case MatchTypeYouthFriendlyMatchCupRules:
		return "youth_friendly_cup_rules"
	case MatchTypeYouthInternationalFriendlyMatchNormalRules:
		return "youth_friendly_international_normal_rules"
	case MatchTypeYouthInternationalFriendlyMatchCupRules:
		return "youth_friendly_international_normal_rules"
	default:
		return ""
	}
}
