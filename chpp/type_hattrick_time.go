package chpp

import (
	"encoding/xml"
	"time"
)

const hattrickTimeLayout = "2006-01-02 15:04:05"

// HattrickTime ...
type HattrickTime time.Time

// UnmarshalXML ...
func (h *HattrickTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	t, err := time.Parse(hattrickTimeLayout, s)
	if err != nil {
		t = time.Time{}
	}

	// TODO take timezone from preferences
	location, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		return err
	}

	*h = HattrickTime(t.In(location))

	return nil
}

// Time ...
func (h *HattrickTime) Time() time.Time {
	return time.Time(*h)
}

// String returns a string representation of the type.
func (h *HattrickTime) String() string {
	return h.Time().Format(hattrickTimeLayout)
}

// FromTime ...
func FromTime(t time.Time) HattrickTime {
	return HattrickTime(t)
}
