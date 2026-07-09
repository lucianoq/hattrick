package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// XML file name and version.
const (
	RegionDetailsAPIFile    = "regiondetails"
	RegionDetailsAPIVersion = "1.2"
)

// RegionDetailsXML contains detailed information about a single region and
// the league it belongs to.
type RegionDetailsXML struct {
	Envelope

	// Container for the data about the league that the region is
	// located in.
	League struct {
		ID   id.League `xml:"LeagueID"`
		Name string    `xml:"LeagueName"`

		Region RegionDetails `xml:"Region"`
	} `xml:"League"`
}

// RegionDetails describes a single Hattrick region, a geographic subdivision
// of a league used to group users by real-world-like locality.
type RegionDetails struct {
	ID   id.Region `xml:"RegionID"`
	Name string    `xml:"RegionName"`

	// The number of active users in the region.
	NumberOfUsers uint `xml:"NumberOfUsers"`

	// The number of users currently online from the region.
	NumberOfOnline uint `xml:"NumberOfOnline"`

	// The current weather.
	Weather Weather `xml:"WeatherID"`

	// The weather forecast for tomorrow.
	TomorrowWeather Weather `xml:"TomorrowWeatherID"`
}
