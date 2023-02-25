package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyAchievements returns the Achievements of the current user.
func (a *API) GetMyAchievements() ([]*chpp.Achievement, error) {
	achievements, err := a.parsed.GetAchievementsXML(nil)
	if err != nil {
		return nil, err
	}
	return achievements.Achievements, nil
}

// GetAchievements returns the Achievements of a specific user.
func (a *API) GetAchievements(userID id.User) ([]*chpp.Achievement, error) {
	achievements, err := a.parsed.GetAchievementsXML(
		map[string]string{
			"userID": userID.String(),
		},
	)
	if err != nil {
		return nil, err
	}
	return achievements.Achievements, nil
}
