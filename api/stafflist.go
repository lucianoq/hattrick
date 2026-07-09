package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyStaff shows the requesting user's primary senior team's trainer and
// other staff members.
func (a *API) GetMyStaff() (*chpp.StaffList, error) {
	res, err := a.parsed.GetStaffListXML(nil)
	if err != nil {
		return nil, err
	}

	return &res.StaffList, nil
}

// GetStaff shows the given team's trainer and other staff members.
func (a *API) GetStaff(teamID id.Team) (*chpp.StaffList, error) {
	values := map[string]string{
		"teamId": teamID.String(),
	}

	res, err := a.parsed.GetStaffListXML(values)
	if err != nil {
		return nil, err
	}

	return &res.StaffList, nil
}
