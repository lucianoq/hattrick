package chpp

import (
	"encoding/xml"
	"errors"
	"strings"
)

// MatchStatus ...
type MatchStatus uint

// List of MatchStatus constants.
const (
	MatchStatusNotStarted MatchStatus = 0
	MatchStatusOngoing    MatchStatus = 1
	MatchStatusFinished   MatchStatus = 2
)

// UnmarshalXML ...
func (ms *MatchStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	switch strings.ToUpper(s) {
	case "0", "UPCOMING":
		*ms = MatchStatusNotStarted
	case "1", "ONGOING":
		*ms = MatchStatusOngoing
	case "", "2", "FINISHED":
		*ms = MatchStatusFinished
	default:
		return errors.New("failed on parsing MatchStatus")
	}

	return nil
}

// String returns a string representation of the type.
func (ms *MatchStatus) String() string {
	switch *ms {
	case MatchStatusNotStarted:
		return "UPCOMING"
	case MatchStatusOngoing:
		return "ONGOING"
	case MatchStatusFinished:
		return "FINISHED"
	default:
		return "UNKNOWN"
	}
}
