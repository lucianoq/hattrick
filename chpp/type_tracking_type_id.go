package chpp

import "strconv"

// TrackingTypeID ...
type TrackingTypeID uint

// List of TrackingTypeID constants.
const (
	TrackingTypeSelling           TrackingTypeID = 1
	TrackingTypeBuying            TrackingTypeID = 2
	TrackingTypeMotherClub        TrackingTypeID = 3
	TrackingTypePreviousTeam      TrackingTypeID = 4
	TrackingTypeHotListed         TrackingTypeID = 5
	TrackingTypeLosingBids        TrackingTypeID = 8
	TrackingTypeFinished          TrackingTypeID = 9
	TrackingTypeTransferProspects TrackingTypeID = 10
)

// String returns a string representation of the type.
func (t TrackingTypeID) String() string {
	return strconv.FormatUint(uint64(t), 10)
}
