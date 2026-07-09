package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMe shows the requesting user's manager profile and all their teams.
func (a *API) GetMe() (*chpp.Manager, error) {
	details, err := a.parsed.GetManagerCompendiumXML(nil)
	if err != nil {
		return nil, err
	}

	return &details.Manager, nil
}

// GetManager shows the given user's manager profile and all their teams.
func (a *API) GetManager(uID id.User) (*chpp.Manager, error) {
	values := map[string]string{
		"userId": uID.String(),
	}
	details, err := a.parsed.GetManagerCompendiumXML(values)
	if err != nil {
		return nil, err
	}

	return &details.Manager, nil
}
