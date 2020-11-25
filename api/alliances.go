package api

import (
	"errors"
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAlliancesNameStartsWith ...
func (a *API) GetAlliancesNameStartsWith(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error) {
	if len(searchFor) < 3 {
		return nil, errors.New("searchFor must be >= 3 chars")
	}

	values := map[string]string{
		"SearchType":       "1",
		"SearchFor":        searchFor,
		"SearchLanguageID": searchLanguageID.String(),
		"PageIndex":        strconv.FormatUint(uint64(pageIndex), 10),
	}
	as, err := a.parsed.GetAlliancesXML(values)
	if err != nil {
		return nil, err
	}
	return as.Alliances.Alliances, nil
}

// GetAlliancesAbbreviationIncludes ...
func (a *API) GetAlliancesAbbreviationIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error) {
	if len(searchFor) < 3 {
		return nil, errors.New("searchFor must be >= 3 chars")
	}

	values := map[string]string{
		"SearchType":       "2",
		"SearchFor":        searchFor,
		"SearchLanguageID": searchLanguageID.String(),
		"PageIndex":        strconv.FormatUint(uint64(pageIndex), 10),
	}
	as, err := a.parsed.GetAlliancesXML(values)
	if err != nil {
		return nil, err
	}
	return as.Alliances.Alliances, nil
}

// GetAlliancesDescriptionIncludes ...
func (a *API) GetAlliancesDescriptionIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error) {
	if len(searchFor) < 3 {
		return nil, errors.New("searchFor must be >= 3 chars")
	}

	values := map[string]string{
		"SearchType":       "3",
		"SearchFor":        searchFor,
		"SearchLanguageID": searchLanguageID.String(),
		"PageIndex":        strconv.FormatUint(uint64(pageIndex), 10),
	}
	as, err := a.parsed.GetAlliancesXML(values)
	if err != nil {
		return nil, err
	}
	return as.Alliances.Alliances, nil
}

// GetAlliance ...
func (a *API) GetAlliance(searchFor id.Alliance, searchLanguageID id.Language, pageIndex uint) (*chpp.Alliance, error) {
	values := map[string]string{
		"SearchType":       "4",
		"SearchFor":        searchFor.String(),
		"SearchLanguageID": searchLanguageID.String(),
		"PageIndex":        strconv.FormatUint(uint64(pageIndex), 10),
	}
	as, err := a.parsed.GetAlliancesXML(values)
	if err != nil {
		return nil, err
	}
	return as.Alliances.Alliances[0], nil
}

// GetMyAlliances ...
func (a *API) GetMyAlliances() ([]*chpp.Alliance, error) {
	values := map[string]string{
		"SearchType": "5",
	}
	as, err := a.parsed.GetAlliancesXML(values)
	if err != nil {
		return nil, err
	}
	return as.Alliances.Alliances, nil
}
