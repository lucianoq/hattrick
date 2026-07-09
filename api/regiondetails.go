package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetMyRegion returns detailed information about the logged on user's
// region.
func (a *API) GetMyRegion() (*chpp.RegionDetails, error) {
	res, err := a.parsed.GetRegionDetailsXML(nil)
	if err != nil {
		return nil, err
	}

	return &res.League.Region, nil
}

// GetRegion returns detailed information about the given region.
func (a *API) GetRegion(regionID id.Region) (*chpp.RegionDetails, error) {
	values := map[string]string{
		"regionID": regionID.String(),
	}
	res, err := a.parsed.GetRegionDetailsXML(values)
	if err != nil {
		return nil, err
	}

	return &res.League.Region, nil
}
