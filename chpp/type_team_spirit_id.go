package chpp

// TeamSpiritID is the morale/spirit level of a team's squad, from like
// the cold war to paradise on earth.
type TeamSpiritID uint

// List of TeamSpiritID constants.
const (
	TeamSpiritLikeTheColdWar  TeamSpiritID = 0
	TeamSpiritMurderous       TeamSpiritID = 1
	TeamSpiritFurious         TeamSpiritID = 2
	TeamSpiritIrritated       TeamSpiritID = 3
	TeamSpiritComposed        TeamSpiritID = 4
	TeamSpiritCalm            TeamSpiritID = 5
	TeamSpiritContent         TeamSpiritID = 6
	TeamSpiritSatisfied       TeamSpiritID = 7
	TeamSpiritDelirious       TeamSpiritID = 8
	TeamSpiritWalkingOnClouds TeamSpiritID = 9
	TeamSpiritParadiseOnEarth TeamSpiritID = 10
)
