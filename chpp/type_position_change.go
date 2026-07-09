package chpp

// PositionChange indicates whether a team moved up, down, or stayed put
// in a league/ladder ranking since the last update.
type PositionChange uint

// List of PositionChange constants.
const (
	NoChange   PositionChange = 0
	MovingUp   PositionChange = 1
	MovingDown PositionChange = 2
)

// Icon returns a single-character glyph representing the change ("^" up,
// "v" down, " " unchanged).
func (p PositionChange) Icon() string {
	switch p {
	case NoChange:
		return " "
	case MovingUp:
		// return "⮝"
		// return "⬆️"
		return "^"
	case MovingDown:
		// return "⮟️️"
		// return "⬇️"
		return "v"
	default:
		return ""
	}
}
