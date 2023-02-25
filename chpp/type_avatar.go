package chpp

type Avatar struct {
	// The URL to the card background-image. This will show a silhouette for
	// non-supporter teams.
	BackgroundImage string `xml:"BackgroundImage"`

	// The container for each avatar bodypart item. Two attribute named X
	// and Y indicates where the item should be positioned. There are
	// several of this container for each player. This container will not be
	// provided for non-supporter team.
	Layers []AvatarLayer `xml:"Layer"`
}

// AvatarLayer
// The container for each avatar bodypart item.
// Two attribute named X and Y indicates where the item should be positioned.
// There are several of this container for each player.
// This container will not be provided for non-supporter team.
type AvatarLayer struct {
	// x-coordinate of image layer
	X uint `xml:"x,attr"`

	// y-coordinate of image layer
	Y uint `xml:"y,attr"`

	// The URL to the bodypart item.
	Image string `xml:"Image"`
}
