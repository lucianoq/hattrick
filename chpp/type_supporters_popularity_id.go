package chpp

// SupportersPopularityID ...
type SupportersPopularityID uint

// List of SupportersPopularityID constants.
const (
	SupportersPopularityMurderous           = 0
	SupportersPopularityFurious             = 1
	SupportersPopularityIrritated           = 2
	SupportersPopularityCalm                = 3
	SupportersPopularityContent             = 4
	SupportersPopularitySatisfied           = 5
	SupportersPopularityDelirious           = 6
	SupportersPopularityHighOnLife          = 7
	SupportersPopularityDancingInTheStreet  = 8
	SupportersPopularitySendingYouLovePoems = 9
)

// String returns a string representation of the type.
func (s SupportersPopularityID) String() string {
	switch s {
	case SupportersPopularityMurderous:
		return "murderous"
	case SupportersPopularityFurious:
		return "furious"
	case SupportersPopularityIrritated:
		return "irritated"
	case SupportersPopularityCalm:
		return "calm"
	case SupportersPopularityContent:
		return "content"
	case SupportersPopularitySatisfied:
		return "satisfied"
	case SupportersPopularityDelirious:
		return "delirious"
	case SupportersPopularityHighOnLife:
		return "high on life"
	case SupportersPopularityDancingInTheStreet:
		return "dancing in the streets"
	case SupportersPopularitySendingYouLovePoems:
		return "sending you love poems"
	default:
		return ""
	}
}
