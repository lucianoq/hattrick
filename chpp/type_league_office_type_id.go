package chpp

import "strconv"

// LeagueOfficeTypeID identifies which national team office a league office
// represents (senior national team or U20 team).
type LeagueOfficeTypeID uint

// List of LeagueOfficeTypeID constants.
const (
	LeagueOfficeTypeNationalTeams LeagueOfficeTypeID = 2
	LeagueOfficeTypeU20Teams      LeagueOfficeTypeID = 4
)

// String returns a string representation of the type.
func (l LeagueOfficeTypeID) String() string {
	return strconv.FormatUint(uint64(l), 10)
}
