package chpp

// SupportersPopularityID is how happy a team's supporters are with the
// manager, from murderous to sending you love poems.
type SupportersPopularityID uint

// List of SupportersPopularityID constants.
const (
	SupportersPopularityMurderous           SupportersPopularityID = 0
	SupportersPopularityFurious             SupportersPopularityID = 1
	SupportersPopularityIrritated           SupportersPopularityID = 2
	SupportersPopularityCalm                SupportersPopularityID = 3
	SupportersPopularityContent             SupportersPopularityID = 4
	SupportersPopularitySatisfied           SupportersPopularityID = 5
	SupportersPopularityDelirious           SupportersPopularityID = 6
	SupportersPopularityHighOnLife          SupportersPopularityID = 7
	SupportersPopularityDancingInTheStreets SupportersPopularityID = 8
	SupportersPopularitySendingYouLovePoems SupportersPopularityID = 9
)
