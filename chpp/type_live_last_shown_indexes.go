package chpp

import (
	"encoding/json"
	"strconv"

	"github.com/lucianoq/hattrick/chpp/id"
)

// LiveLastShownIndex is a single match's last shown event index, as used in
// the lastShownIndexes JSON payload for the live file's viewNew actionType.
type LiveLastShownIndex struct {
	MatchID id.Match

	// SourceSystem the match belongs to. Defaults to Hattrick if empty.
	SourceSystem SourceSystem

	// The last shown EventIndex for this match (from LiveMatchInfo's
	// LastShownEventIndex). -1 requests all events; omitting the match
	// entirely defaults to -1 as well.
	Index int
}

// Hattrick JSON-formatted last shown indexes entry.
type liveLastShownIndex struct {
	MatchID      string `json:"matchId"`
	SourceSystem string `json:"sourceSystem,omitempty"`
	Index        string `json:"index"`
}

// Hattrick JSON-formatted last shown indexes payload.
type liveLastShownIndexes struct {
	Matches []liveLastShownIndex `json:"matches"`
}

// NewLastShownIndexes returns a string containing the JSON representation
// of the given match indexes, in the format required by the live file's
// viewNew actionType's lastShownIndexes parameter.
func NewLastShownIndexes(indexes ...LiveLastShownIndex) (string, error) {
	matches := make([]liveLastShownIndex, 0, len(indexes))
	for _, idx := range indexes {
		matches = append(matches, liveLastShownIndex{
			MatchID:      idx.MatchID.String(),
			SourceSystem: string(idx.SourceSystem),
			Index:        strconv.Itoa(idx.Index),
		})
	}

	buf, err := json.Marshal(liveLastShownIndexes{Matches: matches})
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
