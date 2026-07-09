package api

import (
	"errors"
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetAlliancesNameStartsWith searches for alliances (federations) whose
// name starts with searchFor (at least 3 characters), optionally restricted
// to a language, returning one page of results (up to 25 per page).
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

// GetAlliancesAbbreviationIncludes searches for alliances (federations)
// whose abbreviation contains searchFor (at least 3 characters), optionally
// restricted to a language, returning one page of results (up to 25 per
// page).
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

// GetAlliancesDescriptionIncludes searches for alliances (federations)
// whose description contains searchFor (at least 3 characters), optionally
// restricted to a language, returning one page of results (up to 25 per
// page).
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

// GetAlliance looks up a single alliance (federation) by its ID.
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
	if len(as.Alliances.Alliances) == 0 {
		return nil, errors.New("alliance not found")
	}
	return as.Alliances.Alliances[0], nil
}

// GetMyAlliances lists the alliances (federations) the requesting user is a
// member of.
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
