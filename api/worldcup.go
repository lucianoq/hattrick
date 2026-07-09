package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
)

// GetWorldCupGroups shows the scores/standings for the groups of the given
// cup, season and match round. If cupID is 0, defaults to the World Cup
// (new format). If matchRound is 0, defaults to the last initialized round.
func (a *API) GetWorldCupGroups(cupID chpp.WorldCupID, season, matchRound uint) ([]*chpp.WorldCupScore, error) {
	values := map[string]string{
		"actionType": "viewGroups",
		"season":     strconv.FormatUint(uint64(season), 10),
	}
	if cupID != 0 {
		values["cupID"] = strconv.FormatUint(uint64(cupID), 10)
	}
	if matchRound != 0 {
		values["matchRound"] = strconv.FormatUint(uint64(matchRound), 10)
	}

	res, err := a.parsed.GetWorldCupXML(values)
	if err != nil {
		return nil, err
	}

	return res.Scores, nil
}

// GetWorldCupMatches shows the matches of the given cup, season and group
// (CupSeriesUnit). If cupID is 0, defaults to the World Cup (new format).
func (a *API) GetWorldCupMatches(cupID chpp.WorldCupID, season uint, cupSeriesUnitID uint) ([]*chpp.WorldCupMatch, error) {
	values := map[string]string{
		"actionType":      "viewMatches",
		"season":          strconv.FormatUint(uint64(season), 10),
		"cupSeriesUnitID": strconv.FormatUint(uint64(cupSeriesUnitID), 10),
	}
	if cupID != 0 {
		values["cupID"] = strconv.FormatUint(uint64(cupID), 10)
	}

	res, err := a.parsed.GetWorldCupXML(values)
	if err != nil {
		return nil, err
	}

	return res.Matches, nil
}
