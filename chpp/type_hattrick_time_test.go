package chpp

import (
	"encoding/xml"
	"testing"
)

func TestHattrickTime_UnmarshalXML(t *testing.T) {
	type wrapper struct {
		T HattrickTime `xml:"T"`
	}

	tests := []struct {
		name    string
		xml     string
		want    string // expected h.String(), skipped if wantZero
		wantErr bool
	}{
		{
			name: "valid datetime",
			xml:  `<wrapper><T>2026-07-09 17:41:45</T></wrapper>`,
			want: "2026-07-09 17:41:45",
		},
		{
			// From a live club.xml FetchedDate, cross-checked against a
			// known BST wall-clock time at the same instant: confirms the
			// wire value is Hattrick Time (HTT, = CET/CEST), not UTC, and
			// that the round trip must not shift it.
			name: "live-fetched FetchedDate round-trips unshifted",
			xml:  `<wrapper><T>2026-07-09 18:21:37</T></wrapper>`,
			want: "2026-07-09 18:21:37",
		},
		{
			name: "empty string is treated as zero value",
			xml:  `<wrapper><T></T></wrapper>`,
		},
		{
			name:    "invalid format",
			xml:     `<wrapper><T>not-a-date</T></wrapper>`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w wrapper
			err := xml.Unmarshal([]byte(tt.xml), &w)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.want == "" {
				if !w.T.Time().IsZero() {
					t.Errorf("T = %v, want the zero value", w.T.Time())
				}
				return
			}

			if got := w.T.String(); got != tt.want {
				t.Errorf("T.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFromTime_RoundTrip(t *testing.T) {
	orig := HattrickTime{}
	if !FromTime(orig.Time()).Time().IsZero() {
		t.Error("FromTime(zero value).Time() should be zero")
	}
}
