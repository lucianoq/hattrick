package api

import (
	"errors"
	"strings"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetChallenges shows the requesting user's team's outgoing friendly
// challenges and incoming offers from other teams.
func (a *API) GetChallenges(weekend bool) ([]*chpp.ChallengeByMe, []*chpp.OffersByOthers, error) {
	values := map[string]string{
		"actionType": "view",
	}

	if weekend {
		values["isWeekendFriendly"] = "1"
	}

	res, err := a.parsed.GetChallengesXML(values)
	if err != nil {
		return nil, nil, err
	}
	return res.Team.ChallengesByMe, res.Team.OffersByOthers, nil
}

// IsChallengeable reports whether the given team can currently be
// challenged to a friendly.
func (a *API) IsChallengeable(weekend bool, teamID id.Team) (bool, error) {
	values := map[string]string{
		"actionType":       "challengeable",
		"suggestedTeamIds": teamID.String(),
	}

	if weekend {
		values["isWeekendFriendly"] = "1"
	}

	ch, err := a.parsed.GetChallengesXML(values)
	if err != nil {
		return false, err
	}

	if len(ch.Team.ChallengeableResult.Opponent) == 0 {
		return false, errors.New("error in challenges file returned")
	}

	return ch.Team.ChallengeableResult.Opponent[0].IsChallengeable, nil
}

// AreChallengeable reports, for each given team in order, whether it can
// currently be challenged to a friendly.
func (a *API) AreChallengeable(weekend bool, teamIDs ...id.Team) ([]bool, error) {
	strIDs := make([]string, len(teamIDs))
	for i, t := range teamIDs {
		strIDs[i] = t.String()
	}

	values := map[string]string{
		"actionType":       "challengeable",
		"suggestedTeamIds": strings.Join(strIDs, ","),
	}

	if weekend {
		values["isWeekendFriendly"] = "1"
	}

	ch, err := a.parsed.GetChallengesXML(values)
	if err != nil {
		return nil, err
	}

	if len(ch.Team.ChallengeableResult.Opponent) != len(teamIDs) {
		return nil, errors.New("error in challenges file returned")
	}

	byTeamID := make(map[id.Team]bool, len(ch.Team.ChallengeableResult.Opponent))
	for _, op := range ch.Team.ChallengeableResult.Opponent {
		byTeamID[op.TeamID] = op.IsChallengeable
	}

	results := make([]bool, len(teamIDs))
	for i, t := range teamIDs {
		isChallengeable, ok := byTeamID[t]
		if !ok {
			return nil, errors.New("error in challenges file returned")
		}
		results[i] = isChallengeable
	}

	return results, nil
}

// Challenge sends a friendly match challenge to the given opponent team.
func (a *API) Challenge(opponentTeam id.Team, friendlyType chpp.FriendlyType, matchPlace chpp.MatchPlace, otherArena id.Arena, weekend bool) error {
	values := map[string]string{
		"actionType":     "challenge",
		"opponentTeamId": opponentTeam.String(),
		"matchType":      friendlyType.String(),
		"matchPlace":     matchPlace.String(),
	}

	if matchPlace == chpp.MatchPlaceNeutral {
		values["neutralArenaId"] = otherArena.String()
	}

	if weekend {
		values["isWeekendFriendly"] = "1"
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}

// AcceptChallenge accepts an incoming friendly match offer.
func (a *API) AcceptChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "accept",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}

// DeclineChallenge declines an incoming friendly match offer.
func (a *API) DeclineChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "decline",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}

// WithdrawChallenge withdraws an outgoing friendly match challenge that
// hasn't been accepted yet.
func (a *API) WithdrawChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "withdraw",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}
