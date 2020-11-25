package chpp

// XML file name and version.
const (
	NationalTeamDetailsAPIFile    = "nationalteamdetails"
	NationalTeamDetailsAPIVersion = "1.9"
)

// NationalTeamDetailsXML ...
type NationalTeamDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`
}
