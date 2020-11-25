package chpp

// YouthLeagueStatusID ...
type YouthLeagueStatusID uint

// List of YouthLeagueStatusID constants.
const (
	YouthLeagueStatusNotFull              YouthLeagueStatusID = 0
	YouthLeagueStatusAboutToCreateMatches YouthLeagueStatusID = 1
	YouthLeagueStatusLeagueRunning        YouthLeagueStatusID = 3
	YouthLeagueStatusLeagueFinished       YouthLeagueStatusID = 10
)
