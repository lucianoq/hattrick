package chpp

// TrainingType is which skill (or combination of skills/positions) a
// team's training session focuses on.
type TrainingType uint

// List of TrainingType constants. The "who benefits" notes below come from
// the official Hattrick manual (hattrick.org/Help/Rules/Training.aspx) and
// describe, per type, which fielded positions get the full training effect
// versus a reduced or very-small ("osmosis") one for that type's trained
// skill.
const (
	TrainingTypeGeneral TrainingType = 0 // Deprecated
	TrainingTypeStamina TrainingType = 1 // Deprecated

	// "Set pieces": full effect for every player who played that week,
	// with an extra 25% bonus for the goalkeeper and the set pieces
	// taker specifically.
	TrainingTypeSetPieces TrainingType = 2

	// "Defending": full effect for Defenders and Wing backs, very small
	// (osmosis) for everyone else on the pitch.
	TrainingTypeDefending TrainingType = 3

	// "Scoring": full effect for Forwards, very small for everyone else.
	TrainingTypeScoring TrainingType = 4

	// "Crossing (Winger)": full effect for Wingers, reduced for Wing
	// backs, very small for everyone else.
	TrainingTypeWinger TrainingType = 5

	// "Shooting": Scoring skill at a reduced rate, and Set Pieces skill
	// at a very small rate, for every player who played that week (no
	// position gets the full effect for either skill with this type).
	TrainingTypeScoringAndSetPieces TrainingType = 6

	// "Short passes": full effect for Inner Midfielders, Wingers and
	// Forwards, very small for everyone else.
	TrainingTypePassing TrainingType = 7

	// "Playmaking": full effect for Inner Midfielders, reduced for
	// Wingers, very small for everyone else.
	TrainingTypePlaymaking TrainingType = 8

	// "Goalkeeping": full effect for Goalkeepers only.
	TrainingTypeKeeper TrainingType = 9

	// "Through passes" ("Passing (Defenders + Midfielders)"): full
	// effect for Defenders, Inner Midfielders and Wingers, very small
	// for everyone else.
	TrainingTypePassingDefendersAndMidfielders TrainingType = 10

	// "Defensive positions" ("Defending (Defenders + Midfielders)"):
	// reduced effect (not full) for Goalkeepers, Defenders, Inner
	// Midfielders and Wingers, very small (osmosis) for Forwards. This
	// is the one type in the table where the whole benefiting group is
	// parenthesized as reduced rather than listed plainly as full -
	// confirmed against both the English and Italian versions of the
	// official manual, which agree on this despite otherwise matching
	// the pattern of every other type in the table.
	TrainingTypeDefendingDefendersAndMidfielders TrainingType = 11

	// "Wing attacks" ("Winger (Winger + Attackers)"): full effect for
	// Forwards and Wingers, very small for everyone else.
	TrainingTypeWingerAndAttackers TrainingType = 12
)
