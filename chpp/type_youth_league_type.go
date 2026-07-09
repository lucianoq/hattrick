package chpp

// YouthLeagueType is the geographic scope of a youth league (regional,
// national, or international).
type YouthLeagueType uint

// List of YouthLeagueType constants.
const (
	YouthLeagueTypeRegional      YouthLeagueType = 1
	YouthLeagueTypeNational      YouthLeagueType = 2
	YouthLeagueTypeInternational YouthLeagueType = 3
)
