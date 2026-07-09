package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	AchievementsAPIFile    = "achievements"
	AchievementsAPIVersion = "1.2"
)

// AchievementsXML contains the achievements a user has been awarded.
type AchievementsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The maximum number of points a user can currently achieve
	MaxPoints uint `xml:"MaxPoints"`

	// Container for the data about a particular Achievement.
	Achievements []*Achievement `xml:"AchievementList>Achievement"`
}

// Achievement is a single award granted to a user, along with the user's
// progress/rank within it.
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

	// The date when the achievement was awarded
	EventDate HattrickTime `xml:"EventDate"`

	// Points awarded for the achievement
	Points uint `xml:"Points"`

	// True if AchievementTypeID has multiple levels.
	MultiLevel bool `xml:"Multilevel"`

	// The user's current rank in this achievement, where 1 = best rank possible
	Rank uint `xml:"Rank"`

	// Global number of users who have been awarded this achievement
	NumberOfEvents uint `xml:"NumberOfEvents"`
}
