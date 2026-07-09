package chpp

import (
	"encoding/xml"
	"time"
)

const hattrickTimeLayout = "2006-01-02 15:04:05"

// Hattrick Time (HTT) is documented as always equal to Swedish time
// (https://www.hattrick.org/en/Help/Rules/): CET (UTC+1) from November to
// March, CEST (UTC+2) from April to October, for every manager worldwide
// regardless of their own country. Europe/Stockholm is the canonical
// location for that rule (confirmed against a live response from an
// Italian account, which happens to share the same CET/CEST offset).
var hattrickLocation = mustLoadLocation("Europe/Stockholm")

func mustLoadLocation(name string) *time.Location {
	location, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return location
}

// HattrickTime is a CHPP date/time value, always in Hattrick Time (HTT) -
// Sweden's own CET/CEST timezone, used globally regardless of the
// requesting manager's own country.
type HattrickTime time.Time

// UnmarshalXML parses a CHPP date/time string into Hattrick Time (HTT),
// treating an empty string as the zero value rather than an error (many
// CHPP fields are documented as empty when not yet applicable).
func (h *HattrickTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	// Many CHPP files document a date field as "empty" when not yet
	// applicable (e.g. FinishedDate before a match ends). Treat that as
	// the zero value rather than a parse error.
	if s == "" {
		*h = HattrickTime{}
		return nil
	}

	// The wall-clock digits CHPP sends are already Hattrick Time (HTT),
	// not UTC - so this must parse directly into that location rather
	// than parse-as-UTC-then-convert, which would double-shift the value
	// by the CET/CEST offset.
	t, err := time.ParseInLocation(hattrickTimeLayout, s, hattrickLocation)
	if err != nil {
		return err
	}

	*h = HattrickTime(t)

	return nil
}

// Time returns the value as a standard time.Time.
func (h HattrickTime) Time() time.Time {
	return time.Time(h)
}

// String returns a string representation of the type.
func (h HattrickTime) String() string {
	return h.Time().Format(hattrickTimeLayout)
}

// FromTime wraps a standard time.Time as a HattrickTime, e.g. to build a
// query parameter value.
func FromTime(t time.Time) HattrickTime {
	return HattrickTime(t)
}
