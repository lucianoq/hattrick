package chpp

import "strconv"

// LineupRedCardCriteria ...
type LineupRedCardCriteria int

// List of LineupRedCardCriteria constants.
const (
	LineupRedCardIgnore                  LineupRedCardCriteria = -1
	LineupRedCardMyPlayer                LineupRedCardCriteria = 1
	LineupRedCardOpponent                LineupRedCardCriteria = 2
	LineupRedCardMyCentralDefender       LineupRedCardCriteria = 11
	LineupRedCardMyMidfielder            LineupRedCardCriteria = 12
	LineupRedCardMyForward               LineupRedCardCriteria = 13
	LineupRedCardMyWingBack              LineupRedCardCriteria = 14
	LineupRedCardMyWinger                LineupRedCardCriteria = 15
	LineupRedCardOpponentCentralDefender LineupRedCardCriteria = 21
	LineupRedCardOpponentMidfielder      LineupRedCardCriteria = 22
	LineupRedCardOpponentForward         LineupRedCardCriteria = 23
	LineupRedCardOpponentWingBack        LineupRedCardCriteria = 24
	LineupRedCardOpponentWinger          LineupRedCardCriteria = 25
)

// String returns a string representation of the type.
func (x LineupRedCardCriteria) String() string {
	return strconv.Itoa(int(x))
}
