package api

import (
	"github.com/lucianoq/hattrick/chpp"
	m "github.com/lucianoq/hattrick/chpp"
)

// GetBookmarks ...
func (a *API) GetBookmarks(bookmarkType m.BookmarkType) ([]*chpp.Bookmark, error) {
	values := map[string]string{
		"BookmarkType": bookmarkType.String(),
	}

	res, err := a.parsed.GetBookmarksXML(values)
	if err != nil {
		return nil, err
	}

	return res.Bookmarks, nil
}
