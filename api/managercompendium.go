package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMe ...
func (a *API) GetMe() (*chpp.Manager, error) {
	values := map[string]string{
		"sourceSystem": "hattrick",
	}
	details, err := a.parsed.GetManagerCompendiumXML(values)
	if err != nil {
		return nil, err
	}

	return &details.Manager, nil
}

// GetManager ...
func (a *API) GetManager(uID id.User) (*chpp.Manager, error) {
	values := map[string]string{
		"userId":       uID.String(),
		"sourceSystem": "hattrick",
	}
	details, err := a.parsed.GetManagerCompendiumXML(values)
	if err != nil {
		return nil, err
	}

	return &details.Manager, nil
}
