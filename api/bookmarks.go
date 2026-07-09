package api

import (
	"github.com/lucianoq/hattrick/chpp"
)

// GetBookmarks shows the requesting user's bookmarks, optionally filtered
// to a single BookmarkType.
func (a *API) GetBookmarks(bookmarkType chpp.BookmarkType) ([]*chpp.Bookmark, error) {
	values := map[string]string{
		"BookmarkTypeID": bookmarkType.String(),
	}

	res, err := a.parsed.GetBookmarksXML(values)
	if err != nil {
		return nil, err
	}

	return res.Bookmarks, nil
}
