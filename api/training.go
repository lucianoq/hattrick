package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyTraining shows the requesting user's primary senior team's training
// information.
func (a *API) GetMyTraining() (*chpp.TrainingTeam, error) {
	res, err := a.parsed.GetTrainingXML(map[string]string{"actionType": "view"})
	if err != nil {
		return nil, err
	}

	return res.Team, nil
}

// GetTraining shows the given senior team's training information.
func (a *API) GetTraining(teamID id.Team) (*chpp.TrainingTeam, error) {
	values := map[string]string{
		"actionType": "view",
		"teamId":     teamID.String(),
	}

	res, err := a.parsed.GetTrainingXML(values)
	if err != nil {
		return nil, err
	}

	return res.Team, nil
}

// GetTrainingStats shows the distribution among training types for all
// leagues (leagueID=0) or a specific league. For Supporters only.
func (a *API) GetTrainingStats(leagueID id.League) (*chpp.TrainingLeagueStats, error) {
	values := map[string]string{
		"actionType": "stats",
	}
	if leagueID != 0 {
		values["leagueID"] = leagueID.String()
	}

	res, err := a.parsed.GetTrainingXML(values)
	if err != nil {
		return nil, err
	}

	return res.League, nil
}

// SetTraining sets the training settings for the given senior team.
// trainingLevel ranges 0-100; trainingLevelStamina ranges 5-100. Returns
// whether setting the training succeeded, and the team's training data
// after the change. (Requires "set_training" scope.)
func (a *API) SetTraining(teamID id.Team, trainingType chpp.TrainingType, trainingLevel, trainingLevelStamina uint) (bool, *chpp.TrainingTeam, error) {
	values := map[string]string{
		"actionType":           "setTraining",
		"teamId":               teamID.String(),
		"trainingType":         strconv.FormatUint(uint64(trainingType), 10),
		"trainingLevel":        strconv.FormatUint(uint64(trainingLevel), 10),
		"trainingLevelStamina": strconv.FormatUint(uint64(trainingLevelStamina), 10),
	}

	res, err := a.parsed.GetTrainingXML(values)
	if err != nil {
		return false, nil, err
	}

	return res.TrainingSet, res.View, nil
}
