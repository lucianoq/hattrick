package chpp

import (
	"encoding/json"
	"errors"

	"github.com/lucianoq/hattrick/chpp/id"
)

// Hattrick string-based json
type lineup struct {
	Positions     [30]lineupPlayerPosition  `json:"positions"`
	Settings      lineupSettings            `json:"settings"`
	Substitutions []lineupSubstitutionOrder `json:"substitutions"`
}

// Hattrick string-based json
type lineupSettings struct {
	Tactic        string `json:"tactic"`
	SpeechLevel   string `json:"speechLevel"`
	NewLineup     string `json:"newLineup"`
	CoachModifier string `json:"coachModifier"`
}

// NewLineup returns a string containing the JSON representation of a Lineup.
//
// Exported types and constants for the Go Library,
// strong typed struct, less chance of errors,
// leaving to the constructor the job of building the json
// as requested.
// (see org/Community/CHPP/NewDocs/File.aspx?name=matchorders#ref_Lineup_18)
func NewLineup(
	fsPlayers LineupFirstStringPlayers,
	subs LineupSubPlayers,
	captain id.Player, // you can leave it empty
	setPiecesTaker id.Player, // you can leave it empty
	penalty [11]id.Player, // you can leave it empty
	tactic MatchTacticType,
	speechLevel MatchTeamAttitude,
	coachModifier CoachModifier,
	subsOrders []LineupSubstitutionOrder, // you can leave it empty (0)
) (string, error) {

	if len(subsOrders) > 5 {
		return "", errors.New("too many substitution orders")
	}

	lu := [30]lineupPlayerPosition{}

	// First 14 position slots (0-13) represent the players on the field,
	// starting with the goalkeeper (0), followed by right defender (1),
	// right central defender (2) and so on up to left forward (13).
	lu[0].ID = fsPlayers.Goalkeeper.Player.String()
	lu[0].Behaviour = fsPlayers.Goalkeeper.Behaviour.String()
	lu[1].ID = fsPlayers.RightDefender.Player.String()
	lu[1].Behaviour = fsPlayers.RightDefender.Behaviour.String()
	lu[2].ID = fsPlayers.RightCentralDefender.Player.String()
	lu[2].Behaviour = fsPlayers.RightCentralDefender.Behaviour.String()
	lu[3].ID = fsPlayers.CentralDefender.Player.String()
	lu[3].Behaviour = fsPlayers.CentralDefender.Behaviour.String()
	lu[4].ID = fsPlayers.LeftCentralDefender.Player.String()
	lu[4].Behaviour = fsPlayers.LeftCentralDefender.Behaviour.String()
	lu[5].ID = fsPlayers.LeftDefender.Player.String()
	lu[5].Behaviour = fsPlayers.LeftDefender.Behaviour.String()
	lu[6].ID = fsPlayers.RightWinger.Player.String()
	lu[6].Behaviour = fsPlayers.RightWinger.Behaviour.String()
	lu[7].ID = fsPlayers.RightInnerMidfielder.Player.String()
	lu[7].Behaviour = fsPlayers.RightInnerMidfielder.Behaviour.String()
	lu[8].ID = fsPlayers.CentralInnerMidfielder.Player.String()
	lu[8].Behaviour = fsPlayers.CentralInnerMidfielder.Behaviour.String()
	lu[9].ID = fsPlayers.LeftInnerMidfielder.Player.String()
	lu[9].Behaviour = fsPlayers.LeftInnerMidfielder.Behaviour.String()
	lu[10].ID = fsPlayers.LeftWinger.Player.String()
	lu[10].Behaviour = fsPlayers.LeftWinger.Behaviour.String()
	lu[11].ID = fsPlayers.RightForward.Player.String()
	lu[11].Behaviour = fsPlayers.RightForward.Behaviour.String()
	lu[12].ID = fsPlayers.CentralForward.Player.String()
	lu[12].Behaviour = fsPlayers.CentralForward.Behaviour.String()
	lu[13].ID = fsPlayers.LeftForward.Player.String()
	lu[13].Behaviour = fsPlayers.LeftForward.Behaviour.String()

	// The next 5 slots are substitutes in the following order:
	// goalkeeper, defender, midfielder, forward, winger.
	lu[14].ID = subs.Goalkeeper.String()
	lu[15].ID = subs.Defender.String()
	lu[16].ID = subs.Midfielder.String()
	lu[17].ID = subs.Forward.String()
	lu[18].ID = subs.Winger.String()

	// Line 20 and 21 is captain and set pieces taker.
	lu[19].ID = captain.String()
	lu[20].ID = setPiecesTaker.String()

	// The last 11 slots are penalty takers. All 11 of these should always be
	// provided, sending {"id":"0","behaviour":"0"} where not provided by the
	// user.
	for i := 0; i < 11; i++ {
		lu[i+21].ID = penalty[i].String()
	}

	lso := make([]lineupSubstitutionOrder, 0)

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
		Positions: lu,
		Settings: lineupSettings{
			Tactic:        tactic.String(),
			SpeechLevel:   speechLevel.String(),
			NewLineup:     "", // should always be empty
			CoachModifier: coachModifier.String(),
		},
		Substitutions: lso,
	}

	buf, err := json.Marshal(line)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
