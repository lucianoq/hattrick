package chpp

import (
	"encoding/json"
	"testing"

	"github.com/lucianoq/hattrick/chpp/id"
)

func TestNewLineup(t *testing.T) {
	var kickers [11]id.Player

	buf, err := NewLineup(
		LineupFirstStringPlayers{},
		LineupBenchPlayers{},
		kickers,
		1,
		2,
		MatchTacticType(0),
		MatchTeamAttitude(0),
		CoachModifier(0),
		3,
		4,
		[]LineupSubstitutionOrder{
			{PlayerIn: 5, PlayerOut: 6},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if buf == "" {
		t.Fatal("expected non-empty JSON output")
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal([]byte(buf), &decoded); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	positions, ok := decoded["positions"].([]interface{})
	if !ok || len(positions) != 14 {
		t.Fatalf("expected 14 positions, got %v", decoded["positions"])
	}

	bench, ok := decoded["bench"].([]interface{})
	if !ok || len(bench) != 14 {
		t.Fatalf("expected 14 bench slots, got %v", decoded["bench"])
	}

	kickersOut, ok := decoded["kickers"].([]interface{})
	if !ok || len(kickersOut) != 11 {
		t.Fatalf("expected 11 kickers, got %v", decoded["kickers"])
	}
}

func TestNewLineupTooManySubstitutionOrders(t *testing.T) {
	var kickers [11]id.Player

	_, err := NewLineup(
		LineupFirstStringPlayers{},
		LineupBenchPlayers{},
		kickers,
		0,
		0,
		MatchTacticType(0),
		MatchTeamAttitude(0),
		CoachModifier(0),
		0,
		0,
		make([]LineupSubstitutionOrder, 6),
	)
	if err == nil {
		t.Fatal("expected error for too many substitution orders")
	}
}
