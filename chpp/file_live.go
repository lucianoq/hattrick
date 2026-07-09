package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	LiveAPIFile    = "live"
	LiveAPIVersion = "2.3"
)

// LiveXML contains the live match ticker: tracked matches with their
// current score, lineups, substitutions, and events.
type LiveXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for all the tracked matches.
	Matches []*LiveMatchInfo `xml:"MatchList>Match"`
}

// LiveMatchInfo is a single tracked match, with its teams, lineups,
// substitutions, events, and current score.
type LiveMatchInfo struct {
	// The number of the match in the list.
	Index uint `xml:"Index,attr"`

	SourceSystem SourceSystem `xml:"SourceSystem"`
	MatchID      id.Match     `xml:"MatchID"`
	MatchType    MatchType    `xml:"MatchType"`
	MatchDate    HattrickTime `xml:"MatchDate"`

	HomeTeam struct {
		ID             id.Team        `xml:"HomeTeamID"`
		Name           string         `xml:"HomeTeamName"`
		ShortName      string         `xml:"HomeTeamShortName"`
		StartingLineup StartingLineUp `xml:"StartingLineup"`
	} `xml:"HomeTeam"`

	AwayTeam struct {
		ID             id.Team        `xml:"AwayTeamID"`
		Name           string         `xml:"AwayTeamName"`
		ShortName      string         `xml:"AwayTeamShortName"`
		StartingLineup StartingLineUp `xml:"StartingLineup"`
	} `xml:"AwayTeam"`

	// The starting lineup for the match, as a single combined list of both
	// teams' players. Only sent for actionType=viewAll/addMatch, as a
	// sibling of HomeTeam/AwayTeam rather than nested within them (unlike
	// the per-team StartingLineup sent for actionType=view).
	StartingLineup StartingLineUp `xml:"StartingLineup"`

	Substitutions struct {
		Substitutions []*struct {
			TeamID id.Team `xml:"TeamID"`
			// If substitution: The player leaving.
			// If behaviour change: The player changing his behaviour.
			// If position swap: the first player that will change his position.
			SubjectPlayerID id.Player `xml:"SubjectPlayerID"`

			// If substitution: The player entering
			// If behaviour change: The player changing behaviour.
			// If position swap: The player to swap with.
			ObjectPlayerID       id.Player        `xml:"ObjectPlayerID"`
			OrderType            MatchOrderType   `xml:"OrderType"`
			NewPositionID        MatchRole        `xml:"NewPositionId"`
			NewPositionBehaviour MatchBehaviourID `xml:"NewPositionBehaviour"`
			MatchMinute          uint             `xml:"MatchMinute"`
			MatchPart            MatchPart        `xml:"MatchPart"`
		} `xml:"Substitution"`
	} `xml:"Substitutions"`

	// Container for all the events for this match.
	EventList struct {
		// The index of the last event in this list. Only sent for
		// actionType=viewAll/clearAll/addMatch/deleteMatch.
		Index  uint         `xml:"Index,attr"`
		Events []*LiveEvent `xml:"Event"`
	} `xml:"EventList"`

	HomeGoals uint `xml:"HomeGoals"`
	AwayGoals uint `xml:"AwayGoals"`

	// The last shown event. Use this as LastShownIndexes for
	// actionType=viewNew.
	LastShownEventIndex uint `xml:"LastShownEventIndex"`

	// The minute the next matchevent will happen. Only available while
	// the match is ongoing. Should only be used by the application, not
	// shown to users.
	NextEventMinute uint `xml:"NextEventMinute"`

	// The matchPart in which the next matchevent will happen. Only
	// available while the match is ongoing. Should only be used by the
	// application, not shown to users.
	NextEventMatchPart MatchPart `xml:"NextEventMatchPart"`
}

// LiveEvent is a single event that happened during a live match.
type LiveEvent struct {
	// The order of the event relative to the other events sent in the
	// same batch for this match.
	Index uint `xml:"Index,attr"`

	Minute    uint      `xml:"Minute"`
	MatchPart MatchPart `xml:"MatchPart"`

	// A unique key defining what type of event it was.
	EventKey string `xml:"EventKey"`

	// String describing the event as it would appear in the match
	// report.
	EventText string `xml:"EventText"`

	// 'Unsupported' fields, used at your own risk: for goals/chances
	// these indicate the attacking team, scorer/failed scorer, and
	// defending goalkeeper (or assisting player) respectively, but for
	// other events they may be used freely to carry other data.
	SubjectTeamID   id.Team   `xml:"SubjectTeamID"`
	SubjectPlayerID id.Player `xml:"SubjectPlayerID"`
	ObjectPlayerID  id.Player `xml:"ObjectPlayerID"`
}

// StartingLineUp is the list of players who started a match for a team.
type StartingLineUp struct {
	Players []*LiveMatchPlayer `xml:"Player"`
}

// LiveMatchPlayer is a player in a match's starting lineup.
type LiveMatchPlayer struct {
	PlayerID   id.Player `xml:"PlayerID"`
	RoleID     MatchRole `xml:"RoleID"`
	PlayerName string    `xml:"PlayerName"`
	// The individual order or repositioning the player played with. Not
	// provided for the Captain and the set pieces taker.
	Behaviour MatchBehaviourID `xml:"Behaviour"`
}
