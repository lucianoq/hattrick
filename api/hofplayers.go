package api

import "github.com/lucianoq/hattrick/chpp"

// GetHOFPlayers ...
func (a *API) GetHOFPlayers() ([]*chpp.HOFPlayer, error) {
	e, err := a.parsed.GetHOFPlayersXML(nil)
	if err != nil {
		return nil, err
	}

	return e.PlayerList.Players, nil
}
