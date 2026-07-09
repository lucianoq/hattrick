package chpp

// FanMatchExpectation is how confident a team's fans are about the outcome
// of an upcoming match, from expecting a heavy loss to expecting a win.
type FanMatchExpectation uint

// List of FanMatchExpectation constants.
const (
	FanMatchExpectationBetterNotShowUp      FanMatchExpectation = 0
	FanMatchExpectationWeAreOutclassed      FanMatchExpectation = 1
	FanMatchExpectationWeWillLose           FanMatchExpectation = 2
	FanMatchExpectationTheyAreFavourites    FanMatchExpectation = 3
	FanMatchExpectationTheyHaveTheEdge      FanMatchExpectation = 4
	FanMatchExpectationItWillBeACloseAffair FanMatchExpectation = 5
	FanMatchExpectationWeHaveTheEdge        FanMatchExpectation = 6
	FanMatchExpectationWeAreFavourites      FanMatchExpectation = 7
	FanMatchExpectationWeWillWin            FanMatchExpectation = 8
)
