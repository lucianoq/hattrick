package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyYouthAvatars shows the avatars for all players of the requesting
// user's primary youth team.
func (a *API) GetMyYouthAvatars() ([]*chpp.YouthPlayerAvatars, error) {
	res, err := a.parsed.GetYouthAvatarsXML(nil)
	if err != nil {
		return nil, err
	}

	return res.YouthTeam.Players, nil
}

// GetYouthAvatars shows the avatars for all players of the given youth team.
func (a *API) GetYouthAvatars(youthTeamID id.YouthTeam) ([]*chpp.YouthPlayerAvatars, error) {
	values := map[string]string{
		"youthTeamId": youthTeamID.String(),
	}

	res, err := a.parsed.GetYouthAvatarsXML(values)
	if err != nil {
		return nil, err
	}

	return res.YouthTeam.Players, nil
}
