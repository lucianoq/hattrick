package chpp

// ScoutTravelType is the means of transport a scout uses to travel abroad
// (plane or car), which affects how long the search takes.
type ScoutTravelType uint

// List of ScoutTravelType constants.
const (
	ScoutTravelTypePlane ScoutTravelType = 1
	ScoutTravelTypeCar   ScoutTravelType = 2
)
