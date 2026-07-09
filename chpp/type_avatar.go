package chpp

// Avatar is a player's or staff member's avatar, made up of a background
// image plus a stack of positioned bodypart layers.
type Avatar struct {
	// The URL to the card background-image. This will show a silhouette for
	// non-supporter teams.
	BackgroundImage string `xml:"BackgroundImage"`

	// The container for each avatar bodypart item. Two attributes named X
	// and Y indicate where the item should be positioned. There are
	// several of these containers for each player. This container will not be
	// provided for non-supporter teams.
	Layers []AvatarLayer `xml:"Layer"`
}

// AvatarLayer is the container for a single avatar bodypart item. Its X and
// Y attributes indicate where the item should be positioned; there are
// several of these per player. Not provided for non-supporter teams.
type AvatarLayer struct {
	// x-coordinate of image layer
	X uint `xml:"x,attr"`

	// y-coordinate of image layer
	Y uint `xml:"y,attr"`

	// The URL to the bodypart item.
	Image string `xml:"Image"`
}
