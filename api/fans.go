package api

import "github.com/lucianoq/hattrick/chpp"

// GetFans ...
func (a *API) GetFans() (chpp.Fans, error) {
	e, err := a.parsed.GetFansXML(nil)
	if err != nil {
		return chpp.Fans{}, err
	}

	return e.Team, nil
}
