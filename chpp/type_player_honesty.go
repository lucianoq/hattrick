package chpp

// PlayerHonesty ...
type PlayerHonesty uint

// List of PlayerHonesty constants.
const (
	PlayerHonestyInfamous  PlayerHonesty = 0
	PlayerHonestyDishonest PlayerHonesty = 1
	PlayerHonestyHonest    PlayerHonesty = 2
	PlayerHonestyUpright   PlayerHonesty = 3
	PlayerHonestyRighteous PlayerHonesty = 4
	PlayerHonestySaintly   PlayerHonesty = 5
)

// String returns a string representation of the type.
func (h PlayerHonesty) String() string {
	switch h {
	case PlayerHonestyInfamous:
		return "infamous"
	case PlayerHonestyDishonest:
		return "dishonest"
	case PlayerHonestyHonest:
		return "honest"
	case PlayerHonestyUpright:
		return "upright"
	case PlayerHonestyRighteous:
		return "righteous"
	case PlayerHonestySaintly:
		return "saintly"
	}

	return ""
}
