package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetCupMatchesLast ...
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

// GetCupMatches ...
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

func (a *API) getPaginatedCupMatches(values map[string]string) ([]*chpp.CupMatch, error) {
	list := make([]*chpp.CupMatch, 0)

	for {
		res, err := a.parsed.GetCupMatchesXML(values)
		if err != nil {
			return nil, err
		}

		list = append(list, res.Cup.MatchList.Match...)

		if len(res.Cup.MatchList.Match) < 256 {
			break
		}

		values["StartAfterMatchID"] = res.Cup.MatchList.Match[255].MatchID.String()
	}

	return list, nil
}
