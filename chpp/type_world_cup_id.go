package chpp

// WorldCupID identifies which national-team World Cup tournament a match
// or competition belongs to (regular or U-20, old or new format).
type WorldCupID uint

// List of WorldCupID constants.
const (
	WorldCupOldFormat    WorldCupID = 39
	WorldCupU20OldFormat WorldCupID = 45
	WorldCupNewFormat    WorldCupID = 137
	WorldCupU20NewFormat WorldCupID = 149
)
