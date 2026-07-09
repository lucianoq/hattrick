package chpp

// TrophyID identifies an award a team or player has won (e.g. league,
// cup, World Cup, Hattrick Masters).
type TrophyID uint

// List of TrophyID constants.
const (
	TrophyCupWinner                TrophyID = 16
	TrophySeriesWinner             TrophyID = 17
	TrophyLeagueWinner             TrophyID = 18
	TrophyWorldCupGold             TrophyID = 78
	TrophyWorldCupSilver           TrophyID = 79
	TrophyWorldCupBronze           TrophyID = 80
	TrophyHattrickMastersWinner    TrophyID = 91
	TrophyHattrickMastersTopScorer TrophyID = 93
	TrophyTournamentWinner         TrophyID = 103
	TrophyTutorialTournamentWinner TrophyID = 203
)

// String returns a string representation of the type.
func (t TrophyID) String() string {
	switch t {
	case TrophyCupWinner:
		return "cup_winner"
	case TrophySeriesWinner:
		return "series_winner"
	case TrophyLeagueWinner:
		return "league_winner"

	case TrophyWorldCupGold:
		return "world_cup_gold"
	case TrophyWorldCupSilver:
		return "world_cup_silver"
	case TrophyWorldCupBronze:
		return "world_cup_bronze"

	case TrophyHattrickMastersWinner:
		return "hattrick_masters_winner"
	case TrophyHattrickMastersTopScorer:
		return "hattrick_masters_top_scorer"

	case TrophyTournamentWinner:
		return "tournament_winner"

	case TrophyTutorialTournamentWinner:
		return "tutorial_tournament_winner"
	}
	return "unknown"
}
