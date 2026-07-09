package chpp

// FanSeasonExpectation is what a team's fans expect from the current
// season, from fearing relegation to expecting to dominate the division.
type FanSeasonExpectation uint

// List of FanSeasonExpectation constants.
const (
	FanSeasonExpectationWeAreNotWorthyOfThisDivision      FanSeasonExpectation = 0
	FanSeasonExpectationEveryDayInThisDivisionIsABonus    FanSeasonExpectation = 1
	FanSeasonExpectationWeWillHaveToFightToStayUp         FanSeasonExpectation = 2
	FanSeasonExpectationAMidTableFinishIsNice             FanSeasonExpectation = 3
	FanSeasonExpectationWeBelongInTheTop4                 FanSeasonExpectation = 4
	FanSeasonExpectationAimForTheTitle                    FanSeasonExpectation = 5
	FanSeasonExpectationWeHaveToWinThisSeason             FanSeasonExpectation = 6
	FanSeasonExpectationWeAreSoMuchBetterThanThisDivision FanSeasonExpectation = 7
)
