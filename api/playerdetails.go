package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetPlayerDetails shows detailed information for the given player.
func (a *API) GetPlayerDetails(playerID id.Player, includeMatchInfo bool) (*chpp.PlayerDetails, error) {
	values := map[string]string{
		"actionType":       "view",
		"playerID":         playerID.String(),
		"includeMatchInfo": strconv.FormatBool(includeMatchInfo),
	}

	res, err := a.parsed.GetPlayerDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Player, nil
}

// PlaceBidOnPlayer places a bid on the given player, for the given team,
// and returns the player's updated details. For Supporters only. (Requires
// "place_bid" scope.) maxBidAmount is the optional autobid ceiling; pass 0
// to omit it.
func (a *API) PlaceBidOnPlayer(playerID id.Player, teamID id.Team, bidAmount, maxBidAmount chpp.Money) (*chpp.PlayerDetails, error) {
	values := map[string]string{
		"actionType": "placeBid",
		"playerID":   playerID.String(),
		"teamId":     teamID.String(),
		"bidAmount":  strconv.FormatInt(int64(bidAmount), 10),
	}
	if maxBidAmount != 0 {
		values["maxBidAmount"] = strconv.FormatInt(int64(maxBidAmount), 10)
	}

	res, err := a.parsed.GetPlayerDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Player, nil
}
