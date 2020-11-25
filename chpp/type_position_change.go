package chpp

// PositionChange ...
type PositionChange uint

// List of PositionChange constants.
const (
	NoChange  PositionChange = 0
	MovingUp  PositionChange = 1
	MovingDow PositionChange = 2
)

// Icon ...
func (p PositionChange) Icon() string {
	switch p {
	case NoChange:
		return " "
	case MovingUp:
		// return "⮝"
		// return "⬆️"
		return "^"
	case MovingDow:
		// return "⮟️️"
		// return "⬇️"
		return "v"
	default:
		return ""
	}
}
