package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetCupMatchesLast shows the matches for the given cup's latest season
// and round.
func (a *API) GetCupMatchesLast(cup id.Cup) ([]*chpp.CupMatch, error) {
	values := map[string]string{
		"CupID": cup.String(),
	}

	list, err := a.getPaginatedCupMatches(values)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// GetCupMatches shows the matches for the given cup, season and round.
func (a *API) GetCupMatches(cup id.Cup, season, round uint) ([]*chpp.CupMatch, error) {
	values := map[string]string{
		"CupID":    cup.String(),
		"Season":   strconv.FormatUint(uint64(season), 10),
		"CupRound": strconv.FormatUint(uint64(round), 10),
	}

	list, err := a.getPaginatedCupMatches(values)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// cupMatchesPageSize is the number of matches returned per page by the
// cupmatches CHPP file; a full page means more results may be available.
const cupMatchesPageSize = 256

func (a *API) getPaginatedCupMatches(values map[string]string) ([]*chpp.CupMatch, error) {
	list := make([]*chpp.CupMatch, 0)

	for {
		res, err := a.parsed.GetCupMatchesXML(values)
		if err != nil {
			return nil, err
		}

		list = append(list, res.Cup.Matches...)

		if len(res.Cup.Matches) < cupMatchesPageSize {
			break
		}

		values["StartAfterMatchID"] = res.Cup.Matches[cupMatchesPageSize-1].MatchID.String()
	}

	return list, nil
}
