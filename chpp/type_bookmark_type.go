package chpp

import "strconv"

// BookmarkType ...
type BookmarkType uint

// List of BookmarkType constants.
const (
	BookmarkTypeAll           BookmarkType = 0
	BookmarkTypeTeamsAndUsers BookmarkType = 1
	BookmarkTypePlayers       BookmarkType = 2
	BookmarkTypeMatches       BookmarkType = 3
	BookmarkTypeForumUsers    BookmarkType = 4
	BookmarkTypeLeagueUnits   BookmarkType = 5
	BookmarkTypeYouthTeams    BookmarkType = 6
	BookmarkTypeYouthPlayers  BookmarkType = 7
	BookmarkTypeYouthMatches  BookmarkType = 8
	BookmarkTypeYouthLeagues  BookmarkType = 9
	BookmarkTypeForumPosts    BookmarkType = 10
	BookmarkTypeForumThreads  BookmarkType = 11
)

// String returns a string representation of the type.
func (b BookmarkType) String() string {
	return strconv.FormatUint(uint64(b), 10)
}
