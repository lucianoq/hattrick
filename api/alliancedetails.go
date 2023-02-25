package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAllianceDetails ...
func (a *API) GetAllianceDetails(allianceID id.Alliance) (*chpp.Alliance, error) {
	values := map[string]string{
		"actionType": "view",
		"allianceID": allianceID.String(),
	}
	ad, err := a.parsed.GetAllianceDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return ad.Alliance, nil
}

// GetAllianceDetailsRoles ...
func (a *API) GetAllianceDetailsRoles(allianceID id.Alliance) ([]*chpp.AllianceRole, error) {
	values := map[string]string{
		"actionType": "roles",
		"allianceID": allianceID.String(),
	}
	allianceDetails, err := a.parsed.GetAllianceDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return allianceDetails.Alliance.Roles.Roles, nil
}

// GetAllianceDetailsMembers ...
func (a *API) GetAllianceDetailsMembers(allianceID id.Alliance) ([]*chpp.AllianceMember, error) {
	values := map[string]string{
		"actionType": "members",
		"allianceID": allianceID.String(),
	}
	allianceDetails, err := a.parsed.GetAllianceDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return allianceDetails.Alliance.Members, nil
}

// GetAllianceDetailsMembersSubset ...
func (a *API) GetAllianceDetailsMembersSubset(allianceID id.Alliance, subset uint) ([]*chpp.AllianceMember, error) {
	values := map[string]string{
		"actionType": "membersSubset",
		"allianceID": allianceID.String(),
	}
	if subset != 0 {
		values["Subset"] = strconv.FormatUint(uint64(subset), 10)
	}
	allianceDetails, err := a.parsed.GetAllianceDetailsXML(values)
	if err != nil {
		return nil, err
	}
	return allianceDetails.Alliance.Members, nil
}
