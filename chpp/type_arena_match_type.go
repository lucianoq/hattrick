package chpp

// ArenaMatchType filters which kind of games arena statistics are shown for
// (all games, competitive only, league only, or friendlies only).
type ArenaMatchType string

// List of ArenaMatchType constants.
const (
	ArenaMatchTypeAll          ArenaMatchType = "All"
	ArenaMatchTypeCompOnly     ArenaMatchType = "CompOnly"
	ArenaMatchTypeLeagueOnly   ArenaMatchType = "LeagueOnly"
	ArenaMatchTypeFriendlyOnly ArenaMatchType = "FriendlyOnly"
)
