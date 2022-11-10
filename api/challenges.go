package api

import (
	"errors"
	"strings"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetChallenges ...
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
	return res.Team.ChallengesByMe.Challenge, res.Team.OffersByOthers.Offer, nil
}

// IsChallengeable ...
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

// AreChallengeable ...
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

	results := make([]bool, len(teamIDs))
	for i, op := range ch.Team.ChallengeableResult.Opponent {
		results[i] = op.IsChallengeable
	}

	return results, nil
}

// Challenge ...
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

// AcceptChallenge ...
func (a *API) AcceptChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "accept",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}

// DeclineChallenge ...
func (a *API) DeclineChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "decline",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}

// WithdrawChallenge ...
func (a *API) WithdrawChallenge(friendlyMatchID id.FriendlyMatch) error {
	values := map[string]string{
		"actionType":      "withdraw",
		"trainingMatchId": friendlyMatchID.String(),
	}

	_, err := a.parsed.GetChallengesXML(values)
	return err
}
