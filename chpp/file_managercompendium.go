package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// XML file name and version.
const (
	ManagerCompendiumAPIFile    = "managercompendium"
	ManagerCompendiumAPIVersion = "1.5"
)

// ManagerCompendiumXML ...
type ManagerCompendiumXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	// Container wrapping the info about the manager
	Manager Manager `xml:"Manager"`
}

// Manager Container wrapping the info about the manager
type Manager struct {

	// The globally unique UserID.
	ID id.User `xml:"UserId"`

	// The 'username' or 'nickname' used in Forums and around the site.
	LoginName string `xml:"Loginname"`

	// The current level of Hattrick Supporter that the user has. Empty if not a Supporter.
	SupporterTier SupporterTier `xml:"SupporterTier"`

	// Container for the last logins.
	LastLogins []HattrickLoginTime `xml:"LastLogins>LoginTime"`

	// Container for the language.
	Language struct {
		ID   id.Language `xml:"LanguageId"`
		Name string      `xml:"LanguageName"`
	}

	// Container for the country.
	Country struct {
		ID   id.Country `xml:"CountryId"`
		Name string     `xml:"CountryName"`
	}

	// Container for the currency selected by the user.
	Currency struct {
		Name string  `xml:"CurrencyName"`
		Rate float64 `xml:"CurrencyRate"`
	}

	// Container for the list of Teams.
	// Container for the data about a particular Team.
	Teams []*ManagerTeam `xml:"Teams>Team"`

	// Container for national teams the manager is a coach of.
	// Empty if user is not a national team coach.
	NationalTeamCoach struct {
		NationalTeamID   id.NationalTeam `xml:"NationalTeam>NationalTeamId"`
		NationalTeamName string          `xml:"NationalTeam>NationalTeamName"`
	} `xml:"NationalTeamCoach"`

	// Container for national teams the manager is an assistant coach of.
	// Empty if user is not a national team assistant coach.
	NationalTeamAssistant struct {
		NationalTeamID   id.NationalTeam `xml:"NationalTeam>NationalTeamId"`
		NationalTeamName string          `xml:"NationalTeam>NationalTeamName"`
	} `xml:"NationalTeamAssistant"`

	// Container for the elements to build the avatar.
	Avatar Avatar `xml:"Avatar"`
}

// ManagerTeam Container for the data about a particular Team.
type ManagerTeam struct {
	// The globally unique TeamID.
	ID id.Team `xml:"TeamId"`

	// The full team name
	Name string `xml:"TeamName"`

	// Container for the data about the team's home ground.
	Arena struct {
		ID   id.Arena `xml:"ArenaId"`
		Name string   `xml:"ArenaName"`
	} `xml:"Arena"`

	// Container for the data about the league ('Sverige', 'England' and
	// so on) the team is part of.
	League struct {
		ID   id.League `xml:"LeagueId"`
		Name string    `xml:"LeagueName"`
	} `xml:"League"`

	// Container for the data about the country ('Sverige', 'England'
	// and so on) the team is part of.
	Country struct {
		ID   id.Country `xml:"CountryId"`
		Name string     `xml:"CountryName"`
	} `xml:"Country"`

	// Data about the team's LeagueLevelUnit ('series'). If the league
	// is playing qualification matches the tag will be empty.
	Series struct {
		// The globally unique LeagueLevelUnitID.
		ID id.Series `xml:"LeagueLevelUnitId"`

		// The name of the LeagueLevelUnit.
		Name string `xml:"LeagueLevelUnitName"`

		// Integer value stating the relative level of the
		// LeagueLevelUnit, with 1 indicating the league's top series.
		Level uint `xml:"LeagueLevel"`
	} `xml:"LeagueLevelUnit"`

	// Container for the data about the region theteam is located in.
	Region struct {
		ID   id.Region `xml:"RegionId"`
		Name string    `xml:"RegionName"`
	} `xml:"Region"`

	// Container for the youthTeam. Will be empty if the team has no Youth academy.
	YouthTeam struct {

		// If the Team has a Youth academy the ID is represented here.
		ID id.YouthTeam `xml:"YouthTeamId"`

		// If the Team has a Youth academy the team name is represented here.
		Name string `xml:"YouthTeamName"`

		// Container for the league of the youthTeam.
		// Will be empty if the youthTeam is currently not playing in a league.
		League struct {
			ID   id.YouthLeague `xml:"YouthLeagueId"`
			Name string         `xml:"YouthLeagueName"`
		} `xml:"YouthLeague"`
	} `xml:"YouthTeam"`
}
