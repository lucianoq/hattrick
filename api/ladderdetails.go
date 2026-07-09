package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetLadderDetails shows the ranking table for the given ladder, on the
// page containing the requesting user's primary senior team.
func (a *API) GetLadderDetails(ladderID id.Ladder) (*chpp.LadderDetails, error) {
	return a.ladderDetails(ladderID, nil, -1, 0)
}

// GetLadderDetailsForTeam shows the ranking table for the given ladder, on
// the page containing the given team.
func (a *API) GetLadderDetailsForTeam(ladderID id.Ladder, teamID id.Team) (*chpp.LadderDetails, error) {
	return a.ladderDetails(ladderID, &teamID, -1, 0)
}

// GetLadderDetailsPage shows the given page (0-indexed) of the ranking
// table for the given ladder. pageSize is capped at 100 by the CHPP; pass 0
// to use the default of 25.
func (a *API) GetLadderDetailsPage(ladderID id.Ladder, pageIndex, pageSize uint) (*chpp.LadderDetails, error) {
	return a.ladderDetails(ladderID, nil, int(pageIndex), pageSize) //nolint:gosec // page indexes are small; ladderDetails uses -1 as a sentinel, hence the signed int
}

func (a *API) ladderDetails(ladderID id.Ladder, teamID *id.Team, pageIndex int, pageSize uint) (*chpp.LadderDetails, error) {
	values := map[string]string{
		"ladderid":  ladderID.String(),
		"pageindex": strconv.Itoa(pageIndex),
	}
	if teamID != nil {
		values["teamid"] = teamID.String()
	}
	if pageSize != 0 {
		values["pagesize"] = strconv.FormatUint(uint64(pageSize), 10)
	}

	res, err := a.parsed.GetLadderDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Ladder, nil
}
