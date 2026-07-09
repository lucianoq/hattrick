package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMySupportedTeams shows the teams the requesting user supports.
func (a *API) GetMySupportedTeams(pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	return a.supportedTeams(nil, pageIndex, pageSize)
}

// GetSupportedTeams shows the teams the given user supports.
func (a *API) GetSupportedTeams(userID id.User, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	return a.supportedTeams(&userID, pageIndex, pageSize)
}

func (a *API) supportedTeams(userID *id.User, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	values := map[string]string{
		"actionType": "supportedteams",
		"pageIndex":  strconv.FormatUint(uint64(pageIndex), 10),
	}
	if userID != nil {
		values["userId"] = userID.String()
	}
	if pageSize != 0 {
		values["pageSize"] = strconv.FormatUint(uint64(pageSize), 10)
	}

	res, err := a.parsed.GetSupportersXML(values)
	if err != nil {
		return nil, err
	}

	return res.SupportedTeams.Teams, nil
}

// GetMySupporters shows the teams supporting the requesting user's primary
// senior team.
func (a *API) GetMySupporters(pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	return a.mySupporters(nil, pageIndex, pageSize)
}

// GetSupportersForTeam shows the teams supporting the given team.
func (a *API) GetSupportersForTeam(teamID id.Team, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	return a.mySupporters(&teamID, pageIndex, pageSize)
}

func (a *API) mySupporters(teamID *id.Team, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error) {
	values := map[string]string{
		"actionType": "mysupporters",
		"pageIndex":  strconv.FormatUint(uint64(pageIndex), 10),
	}
	if teamID != nil {
		values["teamId"] = teamID.String()
	}
	if pageSize != 0 {
		values["pageSize"] = strconv.FormatUint(uint64(pageSize), 10)
	}

	res, err := a.parsed.GetSupportersXML(values)
	if err != nil {
		return nil, err
	}

	return res.MySupporters.Teams, nil
}
