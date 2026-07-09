package chpp

// ScoutSearchTypeID is the player position a scout is asked to look for
// (e.g. keeper, wingback, forward), or any position.
type ScoutSearchTypeID uint

// List of ScoutSearchTypeID constants.
const (
	ScoutSearchTypeAny        ScoutSearchTypeID = 0
	ScoutSearchTypeKeeper     ScoutSearchTypeID = 1
	ScoutSearchTypeDefender   ScoutSearchTypeID = 2
	ScoutSearchTypeWingback   ScoutSearchTypeID = 3
	ScoutSearchTypeMidfielder ScoutSearchTypeID = 4
	ScoutSearchTypeWinger     ScoutSearchTypeID = 5
	ScoutSearchTypeForward    ScoutSearchTypeID = 6
)
