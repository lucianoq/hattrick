package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TrainingAPIFile    = "training"
	TrainingAPIVersion = "2.2"
)

// TrainingXML ...
type TrainingXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`
}
