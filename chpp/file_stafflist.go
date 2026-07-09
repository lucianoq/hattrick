package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	StaffListAPIFile    = "stafflist"
	StaffListAPIVersion = "1.2"
)

// StaffListXML contains the list of a team's trainer and other staff
// members.
type StaffListXML struct {
	Envelope
	UserID id.User `xml:"User"`

	StaffList StaffList `xml:"StaffList"`
}

// StaffList is the container for a team's trainer and other staff members.
type StaffList struct {
	// The team's trainer. Only present since API version 1.2.
	Trainer struct {
		ID           id.Trainer   `xml:"TrainerId"`
		Name         string       `xml:"Name"`
		Age          uint         `xml:"Age"`
		AgeDays      uint         `xml:"AgeDays"`
		ContractDate HattrickTime `xml:"ContractDate"`
		Cost         uint         `xml:"Cost"`
		CountryID    id.Country   `xml:"CountryID"`
		TrainerType  TrainerType  `xml:"TrainerType"`
		Leadership   SkillLevel   `xml:"Leadership"`

		// The trainer skill level, from 1 (lowest) to 5 (highest).
		TrainerSkillLevel uint `xml:"TrainerSkillLevel"`

		TrainerStatus TrainerStatus `xml:"TrainerStatus"`
	} `xml:"Trainer"`

	// The team's other staff members.
	StaffMembers []*StaffListEntry `xml:"StaffMembers>Staff"`

	// Total number of staff members (including the trainer).
	TotalStaffMembers uint `xml:"TotalStaffMembers"`

	// Total cost per week for all the staff members.
	TotalCost uint `xml:"TotalCost"`
}

// StaffListEntry is a single (non-trainer) staff member.
type StaffListEntry struct {
	ID   id.Staff  `xml:"StaffId"`
	Name string    `xml:"Name"`
	Type StaffType `xml:"StaffType"`

	// The level of the staff member.
	Level uint `xml:"StaffLevel"`

	HiredDate HattrickTime `xml:"HiredDate"`

	// The cost per week for the staff member.
	Cost uint `xml:"Cost"`

	// The id of the Hall of Fame player this staff member used to be, if
	// any (0 otherwise).
	HofPlayerID uint `xml:"HofPlayerId"`
}
