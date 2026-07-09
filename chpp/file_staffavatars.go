package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	StaffAvatarsAPIFile    = "staffavatars"
	StaffAvatarsAPIVersion = "1.1"
)

// StaffAvatarsXML contains the avatars of a team's trainer and other staff
// members.
type StaffAvatarsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	*StaffAvatars
}

// StaffAvatars is the container for a team's trainer and other staff
// members' avatars.
type StaffAvatars struct {
	// The team's trainer's avatar.
	Trainer struct {
		ID     id.Trainer `xml:"TrainerId"`
		Avatar Avatar     `xml:"Avatar"`
	} `xml:"Trainer"`

	// The team's other staff members' avatars.
	StaffMembers []*StaffAvatar `xml:"StaffMembers>Staff"`
}

// StaffAvatar is a container for a staff member's avatar.
type StaffAvatar struct {
	ID     id.Staff `xml:"StaffId"`
	Avatar Avatar   `xml:"Avatar"`
}
