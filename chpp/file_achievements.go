package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AchievementsAPIFile    = "achievements"
	AchievementsAPIVersion = "1.2"
)

// AchievementsXML ...
type AchievementsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// The maximum number of points a user currently can achieve
	MaxPoints uint `xml:"MaxPoints"`

	// Container for the data about a particular Achievement.
	Achievements []*Achievement `xml:"AchievementList>Achievement"`
}

// Achievement ...
type Achievement struct {
	// The AchievementTypeID (integer). NOTE. We do not provide a list of
	// available achievements.
	ID id.AchievementID `xml:"AchievementTypeID"`

	// The title for the achievement
	Title string `xml:"AchievementTitle"`

	// The text describing the achievement
	Text string `xml:"AchievementText"`

	// The CategoryID (integer) the achievement belongs to
	Category AchievementCategory `xml:"CategoryID"`

	// The date when achievement was awarded
	EventDate HattrickTime `xml:"EventDate"`

	// Points awarded for achievement
	Points uint `xml:"Points"`

	// True if AchievementTypeID has multiple levels.
	MultiLevel bool `xml:"MultiLevel"`

	// The user's current rank in this achievement, where 1 = best rank possible
	Rank uint `xml:"Rank"`

	// Global number of users been awarded this achievement
	NumberOfEvents uint `xml:"NumberOfEvents"`
}
