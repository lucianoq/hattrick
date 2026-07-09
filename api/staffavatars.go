package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyStaffAvatars shows the avatars for the requesting user's primary
// senior team's trainer and other staff members.
func (a *API) GetMyStaffAvatars() (*chpp.StaffAvatars, error) {
	res, err := a.parsed.GetStaffAvatarsXML(nil)
	if err != nil {
		return nil, err
	}

	return res.StaffAvatars, nil
}

// GetStaffAvatars shows the avatars for the given team's trainer and other
// staff members.
func (a *API) GetStaffAvatars(teamID id.Team) (*chpp.StaffAvatars, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}

	res, err := a.parsed.GetStaffAvatarsXML(values)
	if err != nil {
		return nil, err
	}

	return res.StaffAvatars, nil
}
