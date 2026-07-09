package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetTournamentLeagueTables shows the league tables for the given
// tournament.
func (a *API) GetTournamentLeagueTables(tournamentID id.Tournament) ([]*chpp.TournamentLeagueTable, error) {
	return a.tournamentLeagueTables(tournamentID, 0, 0)
}

// GetTournamentLeagueTablesBySeason is like GetTournamentLeagueTables, but
// for the given season.
func (a *API) GetTournamentLeagueTablesBySeason(tournamentID id.Tournament, season uint) ([]*chpp.TournamentLeagueTable, error) {
	return a.tournamentLeagueTables(tournamentID, season, 0)
}

// GetTournamentLeagueTablesByWorldCupRound is like
// GetTournamentLeagueTables, but for the given World Cup stage. Only used
// for World Cup tournaments.
func (a *API) GetTournamentLeagueTablesByWorldCupRound(tournamentID id.Tournament, worldCupRound uint) ([]*chpp.TournamentLeagueTable, error) {
	return a.tournamentLeagueTables(tournamentID, 0, worldCupRound)
}

func (a *API) tournamentLeagueTables(tournamentID id.Tournament, season, worldCupRound uint) ([]*chpp.TournamentLeagueTable, error) {
	values := map[string]string{
		"tournamentId": tournamentID.String(),
	}
	if season != 0 {
		values["season"] = strconv.FormatUint(uint64(season), 10)
	}
	if worldCupRound != 0 {
		values["worldCupRound"] = strconv.FormatUint(uint64(worldCupRound), 10)
	}

	res, err := a.parsed.GetTournamentLeagueTablesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Tables, nil
}
