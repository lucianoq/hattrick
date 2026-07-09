package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyYouthTeam returns the details of the currently logged on user's
// youth team, and its scouts if showScouts is true.
func (a *API) GetMyYouthTeam(showScouts bool) (*chpp.YouthTeamDetails, []*chpp.YouthScout, error) {
	values := map[string]string{}
	if showScouts {
		values["showScouts"] = "true"
	}

	details, err := a.parsed.GetYouthTeamDetailsXML(values)
	if err != nil {
		return nil, nil, err
	}

	return &details.YouthTeam, details.ScoutList, nil
}

// GetYouthTeam returns the details of the given youth team, and its
// scouts if showScouts is true.
func (a *API) GetYouthTeam(youthTeamID id.YouthTeam, showScouts bool) (*chpp.YouthTeamDetails, []*chpp.YouthScout, error) {
	values := map[string]string{
		"youthTeamId": youthTeamID.String(),
	}
	if showScouts {
		values["showScouts"] = "true"
	}

	details, err := a.parsed.GetYouthTeamDetailsXML(values)
	if err != nil {
		return nil, nil, err
	}

	return &details.YouthTeam, details.ScoutList, nil
}
