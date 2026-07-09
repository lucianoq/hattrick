package chpp

import (
	"encoding/json"
	"errors"

	"github.com/lucianoq/hattrick/chpp/id"
)

// Hattrick JSON-formatted player slot (numeric id/behaviour), used in the
// positions, bench and kickers arrays.
type lineupPlayerSlot struct {
	ID        id.Player        `json:"id"`
	Behaviour MatchBehaviourID `json:"behaviour"`
}

// Hattrick string-based json
type lineup struct {
	Positions     [14]lineupPlayerSlot      `json:"positions"`
	Bench         [14]lineupPlayerSlot      `json:"bench"`
	Kickers       [11]lineupPlayerSlot      `json:"kickers"`
	Captain       string                    `json:"captain"`
	SetPieces     string                    `json:"setPieces"`
	Settings      lineupSettings            `json:"settings"`
	Substitutions []lineupSubstitutionOrder `json:"substitutions"`
}

// Hattrick string-based json
type lineupSettings struct {
	Tactic             string `json:"tactic"`
	SpeechLevel        string `json:"speechLevel"`
	NewLineup          string `json:"newLineup"`
	CoachModifier      string `json:"coachModifier"`
	ManMarkerPlayerID  string `json:"manMarkerPlayerId"`
	ManMarkingPlayerID string `json:"manMarkingPlayerId"`
}

// NewLineup returns a string containing the JSON representation of a Lineup,
// in the format required by the matchorders CHPP file from version 2.5
// onwards (MatchOrdersAPIVersion).
//
// Exported types and constants for the Go Library,
// strong typed struct, less chance of errors,
// leaving to the constructor the job of building the json
// as requested.
// (see org/Community/CHPP/NewDocs/File.aspx?name=matchorders#ref_Lineup_30)
func NewLineup(
	fsPlayers LineupFirstStringPlayers,
	bench LineupBenchPlayers,
	kickers [11]id.Player, // you can leave entries empty
	captain id.Player, // you can leave it empty
	setPiecesTaker id.Player, // you can leave it empty
	tactic MatchTacticType,
	speechLevel MatchTeamAttitude,
	coachModifier CoachModifier,
	manMarker id.Player, // your player man-marking an opponent; leave empty for none
	manMarking id.Player, // the opponent player being man-marked; leave empty for none
	subsOrders []LineupSubstitutionOrder, // you can leave it empty (0)
) (string, error) {

	if len(subsOrders) > 5 {
		return "", errors.New("too many substitution orders")
	}

	// The 14 field positions, starting with the goalkeeper (0), followed by
	// right defender (1), right central defender (2) and so on up to left
	// forward (13).
	positions := [14]lineupPlayerSlot{
		{ID: fsPlayers.Goalkeeper.Player, Behaviour: fsPlayers.Goalkeeper.Behaviour},
		{ID: fsPlayers.RightDefender.Player, Behaviour: fsPlayers.RightDefender.Behaviour},
		{ID: fsPlayers.RightCentralDefender.Player, Behaviour: fsPlayers.RightCentralDefender.Behaviour},
		{ID: fsPlayers.CentralDefender.Player, Behaviour: fsPlayers.CentralDefender.Behaviour},
		{ID: fsPlayers.LeftCentralDefender.Player, Behaviour: fsPlayers.LeftCentralDefender.Behaviour},
		{ID: fsPlayers.LeftDefender.Player, Behaviour: fsPlayers.LeftDefender.Behaviour},
		{ID: fsPlayers.RightWinger.Player, Behaviour: fsPlayers.RightWinger.Behaviour},
		{ID: fsPlayers.RightInnerMidfielder.Player, Behaviour: fsPlayers.RightInnerMidfielder.Behaviour},
		{ID: fsPlayers.CentralInnerMidfielder.Player, Behaviour: fsPlayers.CentralInnerMidfielder.Behaviour},
		{ID: fsPlayers.LeftInnerMidfielder.Player, Behaviour: fsPlayers.LeftInnerMidfielder.Behaviour},
		{ID: fsPlayers.LeftWinger.Player, Behaviour: fsPlayers.LeftWinger.Behaviour},
		{ID: fsPlayers.RightForward.Player, Behaviour: fsPlayers.RightForward.Behaviour},
		{ID: fsPlayers.CentralForward.Player, Behaviour: fsPlayers.CentralForward.Behaviour},
		{ID: fsPlayers.LeftForward.Player, Behaviour: fsPlayers.LeftForward.Behaviour},
	}

	// The 14 bench slots: 7 primary substitutes (goalkeeper, central
	// defender, wing back, inner midfielder, forward, winger, extra)
	// followed by 7 backup substitutes in the same order. Bench behaviour is
	// always ignored by Hattrick, so it's left at its zero value.
	benchSlot := func(p id.Player) lineupPlayerSlot { return lineupPlayerSlot{ID: p} }
	benchSlots := [14]lineupPlayerSlot{
		benchSlot(bench.Primary.Goalkeeper),
		benchSlot(bench.Primary.CentralDefender),
		benchSlot(bench.Primary.WingBack),
		benchSlot(bench.Primary.InnerMidfielder),
		benchSlot(bench.Primary.Forward),
		benchSlot(bench.Primary.Winger),
		benchSlot(bench.Primary.Extra),
		benchSlot(bench.Backup.Goalkeeper),
		benchSlot(bench.Backup.CentralDefender),
		benchSlot(bench.Backup.WingBack),
		benchSlot(bench.Backup.InnerMidfielder),
		benchSlot(bench.Backup.Forward),
		benchSlot(bench.Backup.Winger),
		benchSlot(bench.Backup.Extra),
	}

	// The 11 penalty takers. Kicker behaviour is always ignored by
	// Hattrick, so it's left at its zero value.
	var kickerSlots [11]lineupPlayerSlot
	for i, p := range kickers {
		kickerSlots[i] = benchSlot(p)
	}

	lso := make([]lineupSubstitutionOrder, 0, len(subsOrders))
	for _, s := range subsOrders {
		lso = append(lso, lineupSubstitutionOrder{
			PlayerIn:  s.PlayerIn.String(),
			PlayerOut: s.PlayerOut.String(),
			OrderType: s.OrderType.String(),
			Minute:    s.Minute.String(),
			Position:  s.Position.String(),
			Behaviour: s.Behaviour.String(),
			Card:      s.Card.String(),
			Standing:  s.Standing.String(),
		})
	}

	line := &lineup{
		Positions: positions,
		Bench:     benchSlots,
		Kickers:   kickerSlots,
		Captain:   captain.String(),
		SetPieces: setPiecesTaker.String(),
		Settings: lineupSettings{
			Tactic:             tactic.String(),
			SpeechLevel:        speechLevel.String(),
			NewLineup:          "", // should always be empty
			CoachModifier:      coachModifier.String(),
			ManMarkerPlayerID:  manMarker.String(),
			ManMarkingPlayerID: manMarking.String(),
		},
		Substitutions: lso,
	}

	buf, err := json.Marshal(line)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
