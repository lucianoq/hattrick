package chpp

import (
	"encoding/json"
	"testing"
)

func TestNewLastShownIndexes(t *testing.T) {
	got, err := NewLastShownIndexes(
		LiveLastShownIndex{MatchID: 123, Index: 5},
		LiveLastShownIndex{MatchID: 456, SourceSystem: SourceSystemYouthSystem, Index: -1},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded struct {
		Matches []struct {
			MatchID      string `json:"matchId"`
			SourceSystem string `json:"sourceSystem"`
			Index        string `json:"index"`
		} `json:"matches"`
	}
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	if len(decoded.Matches) != 2 {
		t.Fatalf("got %d matches, want 2", len(decoded.Matches))
	}

	m0 := decoded.Matches[0]
	if m0.MatchID != "123" || m0.Index != "5" || m0.SourceSystem != "" {
		t.Errorf("matches[0] = %+v, want MatchID=123 Index=5 SourceSystem=\"\"", m0)
	}

	m1 := decoded.Matches[1]
	if m1.MatchID != "456" || m1.Index != "-1" || m1.SourceSystem != string(SourceSystemYouthSystem) {
		t.Errorf("matches[1] = %+v, want MatchID=456 Index=-1 SourceSystem=%q", m1, SourceSystemYouthSystem)
	}
}

func TestNewLastShownIndexes_Empty(t *testing.T) {
	got, err := NewLastShownIndexes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `{"matches":[]}`
	if got != want {
		t.Errorf("NewLastShownIndexes() = %q, want %q", got, want)
	}
}
