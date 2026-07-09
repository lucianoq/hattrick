package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
)

// SearchFilters holds the optional filters for Search, beyond the primary
// search string or ID.
type SearchFilters struct {
	// A second search string, only used for SearchTypePlayers.
	SearchString2 string

	// The specific ID to search for, used instead of a search string for
	// some search types.
	SearchID uint

	// The league to search in. -1 searches in all leagues. 0 defaults to
	// the logged in user's league (only meaningful for
	// SearchTypePlayers/SearchTypeSeries).
	SearchLeagueID int

	PageIndex uint
}

// Search performs a CHPP search of the given type for the given search
// string (searchString must be at least 3 characters long, or empty if
// searching by SearchFilters.SearchID instead).
func (a *API) Search(searchType chpp.SearchType, searchString string, filters SearchFilters) ([]*chpp.SearchResult, error) {
	values := map[string]string{
		"searchType": strconv.FormatUint(uint64(searchType), 10),
		"pageIndex":  strconv.FormatUint(uint64(filters.PageIndex), 10),
	}
	if searchString != "" {
		values["searchString"] = searchString
	}
	if filters.SearchString2 != "" {
		values["searchString2"] = filters.SearchString2
	}
	if filters.SearchID != 0 {
		values["searchID"] = strconv.FormatUint(uint64(filters.SearchID), 10)
	}
	if filters.SearchLeagueID != 0 {
		values["searchLeagueID"] = strconv.Itoa(filters.SearchLeagueID)
	}

	res, err := a.parsed.GetSearchXML(values)
	if err != nil {
		return nil, err
	}

	return res.Results, nil
}
