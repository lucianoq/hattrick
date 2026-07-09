package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAllianceDetails shows general information about the given alliance
// (federation): name, description, logo, top user, and the requesting
// user's own role in it.
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

// GetAllianceDetailsRoles shows the roles defined within the given alliance
// (federation), including each role's rank, member count and description.
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

// GetAllianceDetailsMembers lists all the members of the given alliance
// (federation), with each member's role and online status.
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

// GetAllianceDetailsMembersSubset lists the members of the given alliance
// (federation) whose name starts with the given character, per subset (use
// 65-90 for 'A'-'Z', or 0 for non-English starting characters).
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
