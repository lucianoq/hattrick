package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetTournament shows the details for the given tournament, for its current
// season.
func (a *API) GetTournament(tournamentID id.Tournament) (*chpp.Tournament, error) {
	return a.GetTournamentBySeason(tournamentID, 0)
}

// GetTournamentBySeason shows the details for the given tournament and
// season. Pass 0 for season to use the tournament's current season.
func (a *API) GetTournamentBySeason(tournamentID id.Tournament, season uint) (*chpp.Tournament, error) {
	values := map[string]string{
		"tournamentId": tournamentID.String(),
	}
	if season != 0 {
		values["season"] = strconv.FormatUint(uint64(season), 10)
	}

	res, err := a.parsed.GetTournamentDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.Tournament, nil
}
