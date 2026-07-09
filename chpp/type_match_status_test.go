package chpp

import (
	"encoding/xml"
	"testing"
)

func TestMatchStatus_UnmarshalXML(t *testing.T) {
	type wrapper struct {
		S MatchStatus `xml:"S"`
	}

	tests := []struct {
		name    string
		value   string
		want    MatchStatus
		wantErr bool
	}{
		{name: "numeric not started", value: "0", want: MatchStatusNotStarted},
		{name: "word not started", value: "UPCOMING", want: MatchStatusNotStarted},
		{name: "numeric ongoing", value: "1", want: MatchStatusOngoing},
		{name: "word ongoing", value: "ONGOING", want: MatchStatusOngoing},
		{name: "lowercase is case-insensitive", value: "ongoing", want: MatchStatusOngoing},
		{name: "numeric finished", value: "2", want: MatchStatusFinished},
		{name: "word finished", value: "FINISHED", want: MatchStatusFinished},
		{name: "empty string defaults to finished", value: "", want: MatchStatusFinished},
		{name: "unrecognized value", value: "garbage", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := "<wrapper><S>" + tt.value + "</S></wrapper>"
			var w wrapper
			err := xml.Unmarshal([]byte(x), &w)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if w.S != tt.want {
				t.Errorf("S = %v, want %v", w.S, tt.want)
			}
		})
	}
}

func TestMatchStatus_String(t *testing.T) {
	tests := []struct {
		ms   MatchStatus
		want string
	}{
		{MatchStatusNotStarted, "UPCOMING"},
		{MatchStatusOngoing, "ONGOING"},
		{MatchStatusFinished, "FINISHED"},
		{MatchStatus(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if got := tt.ms.String(); got != tt.want {
			t.Errorf("MatchStatus(%d).String() = %q, want %q", tt.ms, got, tt.want)
		}
	}
}
