package chpp

// XML file name and version.
const (
	RegionDetailsAPIFile    = "regiondetails"
	RegionDetailsAPIVersion = "1.3"
)

// RegionDetailsXML ...
type RegionDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`
}
