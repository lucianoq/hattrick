package chpp

import (
	"encoding/xml"
	"testing"
)

func TestHattrickLoginTime_UnmarshalXML(t *testing.T) {
	type wrapper struct {
		L HattrickLoginTime `xml:"L"`
	}

	tests := []struct {
		name    string
		xml     string
		wantIP  string
		wantErr bool
	}{
		{
			name:   "valid datetime and IP",
			xml:    `<wrapper><L>2026-07-09 17:41:45 : 1.2.3.4</L></wrapper>`,
			wantIP: "1.2.3.4",
		},
		{
			name:    "missing separator",
			xml:     `<wrapper><L>2026-07-09 17:41:45 1.2.3.4</L></wrapper>`,
			wantErr: true,
		},
		{
			name:    "invalid date part",
			xml:     `<wrapper><L>not-a-date : 1.2.3.4</L></wrapper>`,
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

			if w.L.IP != tt.wantIP {
				t.Errorf("IP = %q, want %q", w.L.IP, tt.wantIP)
			}
			if w.L.Time.IsZero() {
				t.Error("Time should not be zero")
			}
		})
	}
}
