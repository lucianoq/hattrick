package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyArena ...
func (a *API) GetMyArena() (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArena ...
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

// GetArenaByTeamID ...
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
