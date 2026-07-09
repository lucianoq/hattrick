package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTransfersTeam shows the requesting user's primary team's transfer
// history and stats, for the given page. pageIndex=0 means "last page", not
// literally page 0 (per the CHPP doc); each page has 25 records except
// possibly the last.
func (a *API) GetMyTransfersTeam(pageIndex uint) (*chpp.TransferSteamXML, error) {
	values := map[string]string{
		"pageIndex": strconv.FormatUint(uint64(pageIndex), 10),
	}

	return a.parsed.GetTransferSteamXML(values)
}

// GetTransfersTeam shows the given team's transfer history and stats, for
// the given page. pageIndex=0 means "last page", not literally page 0 (per
// the CHPP doc); each page has 25 records except possibly the last.
func (a *API) GetTransfersTeam(teamID id.Team, pageIndex uint) (*chpp.TransferSteamXML, error) {
	values := map[string]string{
		"teamID":    teamID.String(),
		"pageIndex": strconv.FormatUint(uint64(pageIndex), 10),
	}

	return a.parsed.GetTransferSteamXML(values)
}
