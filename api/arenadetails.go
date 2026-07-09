package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyArena shows the details of the requesting user's own arena.
func (a *API) GetMyArena() (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArena shows the details of the given arena.
func (a *API) GetArena(arenaID id.Arena) (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	values["arenaID"] = arenaID.String()
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArenaByTeamID shows the details of the given team's arena.
func (a *API) GetArenaByTeamID(teamID id.Team) (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	values["teamId"] = teamID.String()
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArenaStats returns supporter statistics for the biggest arenas in the
// given league, or globally if leagueID is 0.
func (a *API) GetArenaStats(leagueID id.League) (*chpp.LeagueArenaStats, error) {
	values := map[string]string{
		"StatsType":     "OtherArenas",
		"StatsLeagueID": leagueID.String(),
	}
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.LeagueArenaStats, nil
}

// GetMyArenaSupporterStats returns supporter statistics for the requesting
// user's own arena, for all match types over its full history.
func (a *API) GetMyArenaSupporterStats() (*chpp.MyArenaStats, error) {
	values := map[string]string{
		"StatsType": "MyArena",
	}
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.MyArena, nil
}
