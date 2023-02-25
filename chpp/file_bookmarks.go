package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	BookmarksAPIFile    = "bookmarks"
	BookmarksAPIVersion = "1.0"
)

// BookmarksXML contains user bookmarks.
type BookmarksXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Bookmarks []*Bookmark `xml:"BookmarkList>Bookmark"`
}

// Bookmark ...
type Bookmark struct {
	// The unique ID of the bookmark. The ID is prefixed with a 'y' for youth
	// data with BookmarkType of 6, 7, 8 and 9, and prefixed with a 'c' for
	// forum data with BookmarkType of 4, 10 and 11.
	BookmarkID id.Bookmark `xml:"BookmarkID"`

	// The type of bookmark.
	BookmarkType BookmarkType `xml:"BookmarkTypeID"`

	// BookmarkTypeReturnValues
	//  BookmarkTypeID  Text             Text2      ObjectID           ObjectID2  ObjectID3
	//  --------------  ---------------  ---------  -----------------  ---------  ---------
	//  TeamsAndUser    TeamName         UserAlias  TeamID             n/a        n/a
	//  Players         PlayerName       Deadline   PlayerID           n/a        n/a
	//  Matches         Teams            MatchDate  MatchID            n/a        n/a
	//  ForumUsers      Alias            n/a        UserID             n/a        n/a
	//  LeagueUnits     LeagueUnitName   n/a        LeagueLevelUnitID  LeagueID   n/a
	//  YouthTeams      YouthTeamName    n/a        YouthTeamID        n/a        n/a
	//  YouthPlayers    YouthPlayerName  n/a        YouthPlayerID      n/a        n/a
	//  YouthMatches    YouthTeams       MatchDate  YouthMatchID       n/a        n/a
	//  YouthLeagues    YouthLeague      n/a        YouthLeagueID      n/a        n/a
	//  ForumPosts      ThreadName       ByUser     ThreadID           MessageID  User
	//  ForumThreads    ThreadSubject    ByUser     ThreadID           UserID     n/a
	Text      string `xml:"Text"`
	Text2     string `xml:"Text2"`
	ObjectID  uint   `xml:"ObjectID"`
	ObjectID2 uint   `xml:"ObjectID2"`
	ObjectID3 uint   `xml:"ObjectID3"`

	// The user supplied comment about a bookmark item.
	Comment string `xml:"Comment"`
}
