package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyArenaDetails ...
func (a *API) GetMyArenaDetails() (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArenaDetailsByArenaID ...
func (a *API) GetArenaDetailsByArenaID(arenaID id.Arena) (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	values["arenaID"] = arenaID.String()
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}

// GetArenaDetailsByTeamID ...
func (a *API) GetArenaDetailsByTeamID(teamID id.Team) (*chpp.Arena, error) {
	values := map[string]string{}
	values["StatsType"] = ""
	values["teamId"] = teamID.String()
	arenaDetails, err := a.parsed.GetArenaDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return arenaDetails.Arena, nil
}
