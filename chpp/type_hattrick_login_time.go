package chpp

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// HattrickLoginTime is a login timestamp paired with the IP address it was
// made from, as reported by the loginhistory CHPP file's combined
// "yyyy-MM-dd HH:mm:ss : x.x.x.x" text format.
type HattrickLoginTime struct {
	Time time.Time
	IP   string
}

// UnmarshalXML parses the combined "<date/time> : <IP>" text format into
// its Time and IP fields.
func (h *HattrickLoginTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	fields := strings.Split(s, " : ")
	if len(fields) != 2 {
		return fmt.Errorf("unexpected HattrickLoginTime format: %q", s)
	}

	// See HattrickTime.UnmarshalXML: the date part is already Hattrick
	// Time (HTT), so it must be parsed directly into that location.
	t, err := time.ParseInLocation(hattrickTimeLayout, fields[0], hattrickLocation)
	if err != nil {
		return err
	}

	*h = HattrickLoginTime{
		Time: t,
		IP:   fields[1],
	}

	return nil
}
