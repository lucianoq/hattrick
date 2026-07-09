package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthTeamDetailsAPIFile    = "youthteamdetails"
	YouthTeamDetailsAPIVersion = "1.4"
)

// YouthTeamDetailsXML contains detailed information about a single youth
// team, and optionally its scouts.
type YouthTeamDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for the youth team to which the data applies.
	YouthTeam YouthTeamDetails `xml:"YouthTeam"`

	// Container for the list of scouts. Only present when the
	// showScouts input parameter is set to true.
	ScoutList []*YouthScout `xml:"ScoutList>Scout"`
}

// YouthTeamDetails is the full detailed information for a single youth
// team.
type YouthTeamDetails struct {
	ID        id.YouthTeam `xml:"YouthTeamID"`
	Name      string       `xml:"YouthTeamName"`
	ShortName string       `xml:"ShortTeamName"`

	// The date when the youth team was created.
	CreatedDate HattrickTime `xml:"CreatedDate"`

	// The globally unique UserID for the owner of the team.
	OwnerUserID id.User `xml:"UserID"`

	// The gender of the youth team.
	Gender GenderID `xml:"GenderID"`

	// The system type of the league for this team.
	LeagueSystemID LeagueSystemID `xml:"LeagueSystemID"`

	// Container for the country of the team.
	Country struct {
		ID   id.Country `xml:"CountryID"`
		Name string     `xml:"CountryName"`
	} `xml:"Country"`

	// Container for the region of the team.
	Region struct {
		ID   id.Region `xml:"RegionID"`
		Name string    `xml:"RegionName"`
	} `xml:"Region"`

	// Container for the youth arena.
	Arena struct {
		ID   id.Arena `xml:"YouthArenaID"`
		Name string   `xml:"YouthArenaName"`
	} `xml:"YouthArena"`

	// Container for the youth league the team is participating in. Empty
	// if the team is currently not participating in any league.
	League struct {
		ID     id.YouthLeague      `xml:"YouthLeagueID"`
		Name   string              `xml:"YouthLeagueName"`
		Status YouthLeagueStatusID `xml:"YouthLeagueStatus"`
	} `xml:"YouthLeague"`

	// Container for the mother team. Empty if the team has no owner.
	OwningTeam struct {
		MotherTeamID   id.Team `xml:"MotherTeamID"`
		MotherTeamName string  `xml:"MotherTeamName"`
	} `xml:"OwningTeam"`

	// The player that is the trainer ('coach') for the team.
	Trainer id.YouthPlayer `xml:"YouthTrainer>YouthPlayerID"`

	// The date that the next friendly match can be booked. (Will be less
	// than today if it can be booked now. If a friendly is already
	// booked it will be more than 3 weeks into the future.)
	NextTrainingMatchDate HattrickTime `xml:"NextTrainingMatchDate"`
}

// YouthScout is a scout hired by a youth team to search for and suggest new
// youth players.
type YouthScout struct {
	ID   uint   `xml:"YouthScoutID"`
	Name string `xml:"ScoutName"`
	Age  uint   `xml:"Age"`

	// The home country of the scout.
	Country struct {
		ID   id.Country `xml:"CountryID"`
		Name string     `xml:"CountryName"`
	} `xml:"Country"`

	// The home region of the scout.
	Region struct {
		ID   id.Region `xml:"RegionID"`
		Name string    `xml:"RegionName"`
	} `xml:"Region"`

	// The country the scout is currently in.
	InCountry struct {
		ID   id.Country `xml:"CountryID"`
		Name string     `xml:"CountryName"`
	} `xml:"InCountry"`

	// The region the scout is currently in.
	InRegion struct {
		ID   id.Region `xml:"RegionID"`
		Name string    `xml:"RegionName"`
	} `xml:"InRegion"`

	// The date the team hired the scout.
	HiredDate HattrickTime `xml:"HiredDate"`

	// The date when the scout was last called by the team.
	LastCalled HattrickTime `xml:"LastCalled"`

	// Which type of player the scout is currently searching for.
	PlayerTypeSearch ScoutSearchTypeID `xml:"PlayerTypeSearch"`

	// The id of the hall of fame player associated with this scout. 0 if
	// none.
	HOFPlayerID id.Player `xml:"HofPlayerId"`

	// Information about the scout's travel. Empty if the scout is
	// currently not traveling.
	Travel struct {
		StartDate HattrickTime `xml:"TravelStartDate"`

		// Travel distance in kilometers.
		Length uint            `xml:"TravelLength"`
		Type   ScoutTravelType `xml:"TravelType"`
	} `xml:"Travel"`
}
