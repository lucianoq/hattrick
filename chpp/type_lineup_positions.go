package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// Hattrick string-based json
type lineupPlayerPosition struct {
	ID        string `json:"id"`
	Behaviour string `json:"behaviour"`
}

// PlayerIDBehaviour ...
type PlayerIDBehaviour struct {
	Player    id.Player
	Behaviour MatchBehaviourID
}

// LineupFirstStringPlayers ...
type LineupFirstStringPlayers struct {
	Goalkeeper             PlayerIDBehaviour
	RightDefender          PlayerIDBehaviour
	RightCentralDefender   PlayerIDBehaviour
	CentralDefender        PlayerIDBehaviour
	LeftCentralDefender    PlayerIDBehaviour
	LeftDefender           PlayerIDBehaviour
	RightWinger            PlayerIDBehaviour
	RightInnerMidfielder   PlayerIDBehaviour
	CentralInnerMidfielder PlayerIDBehaviour
	LeftInnerMidfielder    PlayerIDBehaviour
	LeftWinger             PlayerIDBehaviour
	RightForward           PlayerIDBehaviour
	CentralForward         PlayerIDBehaviour
	LeftForward            PlayerIDBehaviour
}

// LineupSubPlayers ...
type LineupSubPlayers struct {
	Goalkeeper id.Player
	Defender   id.Player
	Midfielder id.Player
	Forward    id.Player
	Winger     id.Player
}
