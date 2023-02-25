package chpp

import (
	"encoding/xml"
	"strings"
	"time"
)

// HattrickLoginTime ...
type HattrickLoginTime struct {
	Time time.Time
	IP   string
}

// UnmarshalXML ...
func (h *HattrickLoginTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	fields := strings.Split(s, " : ")

	t, err := time.Parse(hattrickTimeLayout, fields[0])
	if err != nil {
		return err
	}

	// TODO take timezone from preferences
	location, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		return err
	}

	*h = HattrickLoginTime{
		Time: t.In(location),
		IP:   fields[1],
	}

	return nil
}
