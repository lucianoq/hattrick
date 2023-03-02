package chpp

import "github.com/lucianoq/hattrick/chpp/id"

// XML file name and version.
const (
	TeamDetailsAPIFile    = "teamdetails"
	TeamDetailsAPIVersion = "3.6"
)

// TeamDetailsXML ...
type TeamDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"UserID"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	User User `xml:"User"`

	Teams []*Team `xml:"Teams>Team"`
}

// User ...
type User struct {
	UserID id.User `xml:"UserID"`

	Language struct {
		ID   id.Language `xml:"LanguageID"`
		Name string      `xml:"LanguageName"`
	} `xml:"Language"`

	// The current level of Hattrick Supporter that the user has (including
	// Diamond). Empty if not a Supporter.
	SupporterTier SupporterTier `xml:"SupporterTier"`

	// The 'username' or 'nickname' used in Forums and around the site. Not
	// supplied if the team is ownerless.
	LoginName string `xml:"Loginname"`

	// The user's personal name or the value 'HIDDEN' to indicate that the
	// user has chosen to hide it. Not supplied if the team is ownerless.
	Name string `xml:"Name"`

	// The user's ICQ number (if any). Not supplied if the team is
	// ownerless.
	ICQ string `xml:"ICQ"`

	// The date and time when the user signed up for Hattrick. Not supplied
	// if the team is ownerless.
	SignupDate HattrickTime `xml:"SignupDate"`

	// The date and time when the user gained control of this team. Not
	// supplied if the team is ownerless.
	ActivationDate HattrickTime `xml:"ActivationDate"`

	// The date and time when the user last logged on. Not supplied if the
	// team is ownerless.
	LastLoginDate HattrickTime `xml:"LastLoginDate"`

	// Whether the user has received the manager license
	HasManagerLicense bool `xml:"HasManagerLicense"`

	// Container for the national teams.
	NationalTeams []*struct {
		// The globally unique identifier of the national team.
		ID id.NationalTeam `xml:"NationalTeamID"`

		// The name of the national team.
		Name string `xml:"NationalTeamName"`

		// The staff type id from the position the user has in the national team
		StaffType NationalTeamStaffType `xml:"NationalTeamStaffType"`
	} `xml:"NationalTeams>NationalTeam"`
}

// Team ...
type Team struct {
	ID            id.Team      `xml:"TeamID"`
	Name          string       `xml:"TeamName"`
	ShortName     string       `xml:"ShortTeamName"`
	IsPrimaryClub bool         `xml:"IsPrimaryClub"`
	FoundedDate   HattrickTime `xml:"FoundedDate"`

	// Container for the data about the team's home ground.
	Arena struct {
		ID   id.Arena `xml:"ArenaID"`
		Name string   `xml:"ArenaName"`
	} `xml:"Arena"`

	// Container for the data about the league ('Sverige', 'England' and
	// so on) the team is part of.
	League struct {
		ID   id.League `xml:"LeagueID"`
		Name string    `xml:"LeagueName"`
	} `xml:"League"`

	// Container for the data about the country ('Sverige', 'England'
	// and so on) the team is part of.
	Country struct {
		ID   id.Country `xml:"CountryID"`
		Name string     `xml:"CountryName"`
	} `xml:"Country"`

	// Container for the data about the region the team is located in.
	Region struct {
		ID   id.Region `xml:"RegionID"`
		Name string    `xml:"RegionName"`
	} `xml:"Region"`

	// The player that is trainer ('coach') for the team. (Not to be
	// confused with the 'owner'/'user').
	Trainer id.Player `xml:"Trainer>PlayerID"`

	// The home page URL that the team has specified.
	HomePage string `xml:"HomePage"`

	// Data about the team's cup status. If the team is playing a match
	// the container is empty.
	Cup struct {

		// Boolean value stating if the team is still in the cup.
		StillInCup bool `xml:"StillInCup"`

		// The globally unique CupID of the cup that the team is still
		// in. Only provided if team is still in cup.
		ID id.Cup `xml:"CupID"`

		// The name of the cup that the team is still in. Only provided
		// if team is still in cup.
		Name string `xml:"CupName"`

		// LeagueLevel for the cup. 0 = National (LeagueLevel 1-6),
		// 7-9 = Divisional. Only provided if team is still in cup.
		LeagueLevel uint `xml:"CupLeagueLevel"`

		// Level of the cup. 1 = National/Divisional, 2 = Challenger,
		// 3 = Consolation. Only provided if team is still in cup.
		Level uint `xml:"CupLevel"`

		// Index of the cup. Always 1 for National and Consolation cups,
		// for Challenger cup: 1 = Emerald, 2 = Ruby, 3 = Sapphire. Only
		// provided if team is still in cup.
		LevelIndex uint `xml:"CupLevelIndex"`

		// Next round (or current if matches are ongoing). (Note that
		// this is different from WorldDetails!). Only provided if team
		// is still in cup.
		MatchRound uint `xml:"MatchRound"`

		// How many rounds are left of this cup. Only provided if team
		// is still in cup.
		MatchRoundsLeft uint `xml:"MatchRoundsLeft"`
	} `xml:"Cup"`

	// Data about the team's Power Rating
	PowerRating struct {
		// The Power Rating team ranking in Global.
		GlobalRanking uint `xml:"GlobalRanking"`

		// The Power Rating team ranking in the same league.
		LeagueRanking uint `xml:"LeagueRanking"`

		// The Power Rating team ranking in the same region.
		RegionRanking uint `xml:"RegionRanking"`

		// The Power Rating value of the team.
		PowerRating uint `xml:"PowerRating"`
	} `xml:"PowerRating"`

	// The team's status about if a friendly is booked. The value is the
	// opponents globally unique teamID if the team has booked a
	// friendly, 0 if not. If the team is playing a match the tag will
	// be empty.
	FriendlyTeamID id.Team `xml:"FriendlyTeamID"`

	// Data about the team's LeagueLevelUnit ('series'). If the league
	// is playing qualification matches the tag will be empty.
	Series struct {
		// The globally unique LeagueLevelUnitID.
		ID id.Series `xml:"LeagueLevelUnitID"`

		// The name of the LeagueLevelUnit.
		Name string `xml:"LeagueLevelUnitName"`

		// Integer value stating the relative level of the
		// LeagueLevelUnit, with 1 indicating the league's top series.
		Level uint `xml:"LeagueLevel"`
	} `xml:"LeagueLevelUnit"`

	// If the team is in a winning streak of at least 2 matches, this
	// integer indicates how many victories in a row it currently has.
	// If the team is playing a match the tag will be empty.
	NumberOfVictories uint `xml:"NumberOfVictories"`

	// If the team is in a streak of at least 2 matches as undefeated,
	// this integer indicates how many matches in a row it is
	// undefeated. If the team is playing a match the tag will be empty.
	NumberOfUndefeated uint `xml:"NumberOfUndefeated"`

	// Container for the data about the team's fanclub.
	Fanclub struct {
		// The globally unique FanclubID.
		ID id.Fanclub `xml:"FanclubID"`

		// The name of the fanclub.
		Name string `xml:"FanclubName"`

		// The amount of members in the fanclub.
		Size uint `xml:"FanclubSize"`
	} `xml:"Fanclub"`

	// The logo URL that the team has specified.
	LogoURL string `xml:"LogoURL"`

	// This container and its elements are only shown if the user has
	// supporter
	// Container for the data about the team's guestbook.
	Guestbook *struct {
		NumberOfGuestbookItems uint `xml:"NumberOfGuestbookItems"`
	} `xml:"Guestbook"`

	// This container and its elements are only shown if the user has
	// supporter
	// Container for the data about the team's most recent
	// PressAnnouncement.
	PressAnnouncement []struct {

		// The subject text specified for the PressAnnouncement.
		Subject string `xml:"Subject"`

		// The body of the PressAnnouncement.
		Body string `xml:"Body"`

		// The time and date when the PressAnnouncement was submitted.
		SendDate HattrickTime `xml:"SendDate"`
	} `xml:"PressAnnouncement"`

	// This container and its elements are only shown if the user has
	// supporter
	// Container for the team's club theme colors. Empty if no theme has
	// been set for the club.
	TeamColors []struct {
		// The defined background color from the club theme.
		BackgroundColor string `xml:"BackgroundColor"`

		// The matching text color to the defined background color.
		Color string `xml:"Color"`
	} `xml:"TeamColors"`

	// URI to an image representing the dress of the team.
	DressURI string `xml:"DressURI"`

	// Same as DressURI, except that it shows the alternate dress.
	DressAlternateURI string `xml:"DressAlternateURI"`

	// Container with data indicating if a team is a 'bot'. If yes the
	// team has default and random lineup and players.
	BotStatus struct {
		// Specifies whether a team is currently a bot or not.
		IsBot bool `xml:"IsBot"`

		// The date when the team was made a bot, only available if team
		// is a bot.
		BotSince HattrickTime `xml:"BotSince"`
	} `xml:"BotStatus"`

	// If the team has an owner, last is the League Rank, a number based
	// on LeagueLevel, position, etc., only counting teams with an owner.
	// If the team is playing a match the tag will be empty.
	TeamRank uint `xml:"TeamRank"`

	// If the Team has a Youth academy the ID is represented here. If
	// not it will be 0.
	YouthTeamID id.YouthTeam `xml:"YouthTeamID"`

	// If the Team has a Youth academy the team name is represented
	// here. If not it will be an empty string.
	YouthTeamName string `xml:"YouthTeamName"`

	// The number of visits the team had in the latest day with at least one visit.
	NumberOfVisits uint `xml:"NumberOfVisits"`

	Flags struct {
		Home []flag `xml:"HomeFlags>Flag"`
		Away []flag `xml:"AwayFlags>Flag"`
	} `xml:"Flags"`

	Trophies []*struct {
		TypeID TrophyID `xml:"TrophyTypeId"`
		Season uint     `xml:"TrophySeason"`

		// The level the league the user won. For tournament trophies, this act as a "type".
		LeagueLevel uint `xml:"LeagueLevel"`

		// The LeagueLevelUnitId of the leaguelevelunit that the user won.
		// For tournaments this is the id of the tournament.
		SeriesID   id.Series `xml:"LeagueLevelUnitID"`
		SeriesName string    `xml:"LeagueLevelUnitName"`

		GainedDate HattrickTime `xml:"GainedDate"`
		ImageURL   string       `xml:"ImageUrl"`

		CupLeagueLevel id.CupLeagueLevel `xml:"CupLeagueLevel"`
		CupLevel       id.CupLevel       `xml:"CupLevel"`
		CupLevelIndex  id.CupLevelIndex  `xml:"CupLevelIndex"`
	} `xml:"TrophyList>Trophy"`

	// Container for the supported teams. It will be empty if the
	// user have no supporters and not present if the user are not a
	// supporter. Only available, when includeSupporters=true.
	SupportedTeams *struct {
		TotalItems *uint `xml:"TotalItems,attr"`
		MaxItems   *uint `xml:"MaxItems,attr"`

		SupportedTeam []*struct {
			supportedTeam

			// Container for the last match of the supported team.
			LastMatch struct {
				ID           id.Match     `xml:"LastMatchId"`
				Date         HattrickTime `xml:"LastMatchDate"`
				HomeTeamID   id.Team      `xml:"LastMatchHomeTeamId"`
				HomeTeamName string       `xml:"LastMatchHomeTeamName"`
				HomeGoals    uint         `xml:"LastMatchHomeGoals"`
				AwayTeamID   id.Team      `xml:"LastMatchAwayTeamId"`
				AwayTeamName string       `xml:"LastMatchAwayTeamName"`
				AwayGoals    uint         `xml:"LastMatchAwayGoals"`
			} `xml:"LastMatch"`

			// Container for the last match of the supported team.
			NextMatch struct {
				ID           id.Match     `xml:"NextMatchId"`
				Date         HattrickTime `xml:"NextMatchDate"`
				HomeTeamID   id.Team      `xml:"NextMatchHomeTeamId"`
				HomeTeamName string       `xml:"NextMatchHomeTeamName"`
				AwayTeamID   id.Team      `xml:"NextMatchAwayTeamId"`
				AwayTeamName string       `xml:"NextMatchAwayTeamName"`
			} `xml:"NextMatch"`

			PressAnnouncement struct {
				SendDate HattrickTime `xml:"PressAnnouncementSendDate"`
				Subject  string       `xml:"PressAnnouncementSubject"`
				Body     string       `xml:"PressAnnouncementBody"`
			} `xml:"PressAnnouncement"`
		} `xml:"SupportedTeam"`
	} `xml:"SupportedTeams"`

	// Attribute: 'TotalItems' : unsigned Integer
	//   Shows how many teams supports the team in total
	// Attribute: 'MaxItems' : unsigned Integer
	//   Shows how many teams that can maximum be outputted in this list
	// Container for the supported teams. It will be empty if the
	// user have no supporters and not present if the user are not a
	// supporter. Only available, when includeSupporters=true.
	MySupporters *struct {
		TotalItems *uint `xml:"TotalItems,attr"`
		MaxItems   *uint `xml:"MaxItems,attr"`

		// Container for a supported team. Not supplied if the user doesn't
		// have supporters.
		SupporterTeam []*supportedTeam `xml:"SupporterTeam"`
	} `xml:"MySupporters"`

	// Whether the team is possible to challenge for a mid-week friendly
	PossibleToChallengeMidweek bool `xml:"PossibleToChallengeMidweek"`

	// Whether the team is possible to challenge for a weekend friendly
	PossibleToChallengeWeekend bool `xml:"PossibleToChallengeWeekend"`
}

// flag ...
type flag struct {
	LeagueID    id.League `xml:"LeagueId"`
	LeagueName  string    `xml:"LeagueName"`
	CountryCode string    `xml:"CountryCode"`
}

// supportedTeam ...
type supportedTeam struct {
	UserID     id.User   `xml:"UserId"`
	LoginName  string    `xml:"LoginName"`
	ID         id.Team   `xml:"TeamId"`
	Name       string    `xml:"TeamName"`
	LeagueID   id.League `xml:"LeagueID"`
	LeagueName string    `xml:"LeagueName"`
	SeriesID   id.Series `xml:"LeagueLevelUnitID"`
	SeriesName string    `xml:"LeagueLevelUnitName"`
}
