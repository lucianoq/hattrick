package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetTournamentFixtures shows the fixtures for the given tournament.
func (a *API) GetTournamentFixtures(tournamentID id.Tournament) ([]*chpp.TournamentFixture, error) {
	return a.tournamentFixtures(tournamentID, 0, 0)
}

// GetTournamentFixturesBySeason is like GetTournamentFixtures, but for the
// given season. Only available for National Teams cups.
func (a *API) GetTournamentFixturesBySeason(tournamentID id.Tournament, season uint) ([]*chpp.TournamentFixture, error) {
	return a.tournamentFixtures(tournamentID, season, 0)
}

// GetTournamentFixturesByMatchRound is like GetTournamentFixtures, but for
// the given match round. If matchRound is greater than the tournament's
// last round, the last round is returned instead.
func (a *API) GetTournamentFixturesByMatchRound(tournamentID id.Tournament, matchRound uint) ([]*chpp.TournamentFixture, error) {
	return a.tournamentFixtures(tournamentID, 0, matchRound)
}

func (a *API) tournamentFixtures(tournamentID id.Tournament, season, matchRound uint) ([]*chpp.TournamentFixture, error) {
	values := map[string]string{
		"tournamentId": tournamentID.String(),
	}
	if season != 0 {
		values["season"] = strconv.FormatUint(uint64(season), 10)
	}
	if matchRound != 0 {
		values["matchRound"] = strconv.FormatUint(uint64(matchRound), 10)
	}

	res, err := a.parsed.GetTournamentFixturesXML(values)
	if err != nil {
		return nil, err
	}

	return res.Matches, nil
}
