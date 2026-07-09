package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// PlayerIDBehaviour pairs a player with the behaviour/order they start the
// match with, for one of the 14 starting lineup slots.
type PlayerIDBehaviour struct {
	Player    id.Player
	Behaviour MatchBehaviourID
}

// LineupFirstStringPlayers is the 14 starting-lineup slots (11 field
// players plus reserved slots), for building the matchorders lineup JSON
// payload with NewLineup.
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

// LineupBenchSlots represents one tier (primary or backup) of the 7
// substitute slots, in the order required by the CHPP matchorders "bench"
// array: goalkeeper, central defender, wing back, inner midfielder,
// forward, winger, extra.
type LineupBenchSlots struct {
	Goalkeeper      id.Player
	CentralDefender id.Player
	WingBack        id.Player
	InnerMidfielder id.Player
	Forward         id.Player
	Winger          id.Player
	Extra           id.Player
}

// LineupBenchPlayers represents the 14 bench slots: 7 primary substitutes
// followed by 7 backup substitutes. A player must be a primary substitute
// to qualify as a backup.
type LineupBenchPlayers struct {
	Primary LineupBenchSlots
	Backup  LineupBenchSlots
}
