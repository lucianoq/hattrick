package chpp

import (
	"encoding/xml"
	"testing"
)

func TestYouthScoutCallScout_UnmarshalXML(t *testing.T) {
	type wrapper struct {
		Scout YouthScoutCallScout `xml:"Scout"`
	}

	tests := []struct {
		name     string
		xml      string
		wantID   uint
		wantName string
	}{
		{
			name:     "details casing (ScoutId)",
			xml:      `<wrapper><Scout><ScoutId>5</ScoutId><ScoutName>Bob</ScoutName></Scout></wrapper>`,
			wantID:   5,
			wantName: "Bob",
		},
		{
			name:     "unlockskills casing (ScoutID)",
			xml:      `<wrapper><Scout><ScoutID>7</ScoutID><ScoutName>Alice</ScoutName></Scout></wrapper>`,
			wantID:   7,
			wantName: "Alice",
		},
		{
			name:     "neither casing present",
			xml:      `<wrapper><Scout><ScoutName>NoID</ScoutName></Scout></wrapper>`,
			wantID:   0,
			wantName: "NoID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w wrapper
			if err := xml.Unmarshal([]byte(tt.xml), &w); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if w.Scout.ID != tt.wantID {
				t.Errorf("ID = %d, want %d", w.Scout.ID, tt.wantID)
			}
			if w.Scout.Name != tt.wantName {
				t.Errorf("Name = %q, want %q", w.Scout.Name, tt.wantName)
			}
		})
	}
}
