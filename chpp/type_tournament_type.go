package chpp

// TournamentType is the competition format of a tournament (e.g. league,
// cup, ladder, Swiss tournament, World Cup).
type TournamentType uint

// List of TournamentType constants.
const (
	TournamentTypeSingleMatch          TournamentType = 1
	TournamentTypeLeague               TournamentType = 2
	TournamentTypeLeagueWithPlayoffs   TournamentType = 3
	TournamentTypeCup                  TournamentType = 4
	TournamentTypeDoubleEliminationCup TournamentType = 5
	TournamentTypeLadder               TournamentType = 6
	TournamentTypeSwissTournament      TournamentType = 8
	TournamentTypeDivisionBattle       TournamentType = 10
	TournamentTypeWildcards            TournamentType = 11
	TournamentTypeWorldCup             TournamentType = 12
)
