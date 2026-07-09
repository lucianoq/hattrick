package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetPlayerTransferHistory shows the transfer history for the given
// player.
func (a *API) GetPlayerTransferHistory(playerID id.Player) (*chpp.PlayerTransferHistory, error) {
	values := map[string]string{
		"playerID": playerID.String(),
	}

	res, err := a.parsed.GetTransfersPlayerXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Transfers, nil
}
