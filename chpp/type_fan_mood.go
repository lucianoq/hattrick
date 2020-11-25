package chpp

// FanMood ...
type FanMood uint

// List of FanMood constants.
const (
	FanMoodMurderous             = 0
	FanMoodFurious               = 1
	FanMoodAngry                 = 2
	FanMoodIrritated             = 3
	FanMoodDisappointed          = 4
	FanMoodCalm                  = 5
	FanMoodContent               = 6
	FanMoodSatisfied             = 7
	FanMoodDelirious             = 8
	FanMoodHighOnLife            = 9
	FanMoodDancingInTheStreet    = 10
	FanMoodSendingLovePoemsToYou = 11
)

// String returns a string representation of the type.
func (s FanMood) String() string {
	switch s {
	case FanMoodMurderous:
		return "murderous"
	case FanMoodFurious:
		return "furious"
	case FanMoodAngry:
		return "angry"
	case FanMoodIrritated:
		return "irritated"
	case FanMoodDisappointed:
		return "disappointed"
	case FanMoodCalm:
		return "calm"
	case FanMoodContent:
		return "content"
	case FanMoodSatisfied:
		return "satisfied"
	case FanMoodDelirious:
		return "delirious"
	case FanMoodHighOnLife:
		return "high on life"
	case FanMoodDancingInTheStreet:
		return "dancing in the streets"
	case FanMoodSendingLovePoemsToYou:
		return "sending love poems to you"
	default:
		return ""
	}
}
