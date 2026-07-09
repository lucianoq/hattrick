package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMatchOrders returns the match orders (lineup, tactics and player
// orders) currently saved for the given match and team. If too close to
// kickoff or the user is not involved in the match, the returned
// MatchOrdersMatchData.Available will be false.
func (a *API) GetMatchOrders(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem) (*chpp.MatchOrdersMatchData, error) {
	values := map[string]string{
		"actionType":   "view",
		"matchID":      matchID.String(),
		"sourceSystem": string(sourceSystem),
	}
	if teamID != 0 {
		values["teamId"] = teamID.String()
	}

	res, err := a.parsed.GetMatchOrdersXML(values)
	if err != nil {
		return nil, err
	}

	return &res.MatchData, nil
}

// SetMatchOrders submits a new lineup for the given match and team. Build
// the lineup JSON payload with chpp.NewLineup. (Requires "set_matchorder"
// scope.) The returned MatchOrdersMatchData's OrdersSet indicates whether
// the submission succeeded; if not, Reason explains why.
func (a *API) SetMatchOrders(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem, lineup string) (*chpp.MatchOrdersMatchData, error) {
	values := map[string]string{
		"actionType":   "setmatchorder",
		"matchID":      matchID.String(),
		"sourceSystem": string(sourceSystem),
	}
	if teamID != 0 {
		values["teamId"] = teamID.String()
	}

	res, err := a.parsed.SetMatchOrdersXML(values, lineup)
	if err != nil {
		return nil, err
	}

	return &res.MatchData, nil
}

// PredictRatings predicts the sector and tactic ratings for the given match,
// according to the sent lineup JSON payload (built with chpp.NewLineup). If
// lineup is empty, the predictions are for the currently saved match orders.
func (a *API) PredictRatings(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem, lineup string) (*chpp.MatchOrdersMatchData, error) {
	values := map[string]string{
		"actionType":   "predictratings",
		"matchID":      matchID.String(),
		"sourceSystem": string(sourceSystem),
	}
	if teamID != 0 {
		values["teamId"] = teamID.String()
	}

	res, err := a.parsed.SetMatchOrdersXML(values, lineup)
	if err != nil {
		return nil, err
	}

	return &res.MatchData, nil
}
