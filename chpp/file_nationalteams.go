package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	NationalTeamsAPIFile    = "nationalteams"
	NationalTeamsAPIVersion = "1.6"
)

// NationalTeamsXML contains the list of national teams of a given office
// type, and the cups they take part in.
type NationalTeamsXML struct {
	Envelope

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	*NationalTeamsList
}

// NationalTeamsList is the container for the list of national teams and
// the cups they take part in.
type NationalTeamsList struct {
	// The requested LeagueOfficeTypeID.
	LeagueOfficeTypeID LeagueOfficeTypeID `xml:"LeagueOfficeTypeID"`

	NationalTeams []*NationalTeamListEntry `xml:"NationalTeams>NationalTeam"`

	// The cups the national teams take part in, and which teams take
	// part in each. Replaces the deprecated singular Cup container.
	Cups []*NationalTeamCup `xml:"Cups>Cup"`
}

// NationalTeamListEntry is a single national team, as returned by the
// nationalteams file.
type NationalTeamListEntry struct {
	// The globally unique identifier of the national team.
	ID   id.NationalTeam `xml:"NationalTeamID"`
	Name string          `xml:"NationalTeamName"`

	// 7 character string representing the dress of the team: first
	// character is a letter indicating the type of dress, next two
	// numbers indicate the shirt colour, next two the pants colour and
	// final two the socks colour.
	Dress string `xml:"Dress"`

	RatingScore uint `xml:"RatingScore"`

	// The league the national team belongs to.
	LeagueID id.League `xml:"LeagueID"`
}

// NationalTeamCup is a single cup, and the national teams that take part in
// it.
type NationalTeamCup struct {
	// The teams in the cup. Unavailable while a match is running.
	Teams []*NationalTeamCupTeam `xml:"CupTeams>CupTeam"`
}

// NationalTeamCupTeam is a single national team taking part in a cup.
type NationalTeamCupTeam struct {
	// The globally unique identifier of the national team.
	ID   id.NationalTeam `xml:"CupNationalTeamID"`
	Name string          `xml:"CupNationalTeamName"`
}
