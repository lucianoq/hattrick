package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TeamDetailsAPIFile    = "teamdetails"
	TeamDetailsAPIVersion = "3.5"
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

	TeamDetails
}

// TeamDetails ...
type TeamDetails struct {
	User TeamDetailsUser `xml:"User"`

	Teams struct {
		Team []*TeamDetailsTeam `xml:"Team"`
	} `xml:"Teams"`
}

// TeamDetailsUser ...
type TeamDetailsUser struct {
	UserID id.User `xml:"UserID"`

	Language struct {
		LanguageID   id.Language `xml:"LanguageID"`
		LanguageName string      `xml:"LanguageName"`
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

	// Whether or not the user has received the manager license
	HasManagerLicense bool `xml:"HasManagerLicense"`

	// Container for the national teams.
	NationalTeams struct {
		NationalTeam []*struct {
			// The globally unique identifier of the national team.
			NationalTeamID id.NationalTeam `xml:"NationalTeamID"`
			// The name of the national team.
			NationalTeamName string `xml:"NationalTeamName"`
			// The staff type id from the position the user has in the national team
			NationalTeamStaffType NationalTeamStaffType `xml:"NationalTeamStaffType"`
		}
	} `xml:"NationalTeams"`
}

// TeamDetailsTeam ...
type TeamDetailsTeam struct {
	TeamID        id.Team      `xml:"TeamID"`
	TeamName      string       `xml:"TeamName"`
	ShortTeamName string       `xml:"ShortTeamName"`
	IsPrimaryClub bool         `xml:"IsPrimaryClub"`
	FoundedDate   HattrickTime `xml:"FoundedDate"`

	// Container for the data about the team's home ground.
	Arena struct {
		ArenaID   id.Arena `xml:"ArenaID"`
		ArenaName string   `xml:"ArenaName"`
	} `xml:"Arena"`

	// Container for the data about the league ('Sverige', 'England' and
	// so on) the team is part of.
	League struct {
		LeagueID   id.League `xml:"LeagueID"`
		LeagueName string    `xml:"LeagueName"`
	} `xml:"League"`

	// Container for the data about the country ('Sverige', 'England'
	// and so on) the team is part of.
	Country struct {
		CountryID   id.Country `xml:"CountryID"`
		CountryName string     `xml:"CountryName"`
	} `xml:"Country"`

	// Container for the data about the region theteam is located in.
	Region struct {
		RegionID   id.Region `xml:"RegionID"`
		RegionName string    `xml:"RegionName"`
	} `xml:"Region"`

	// The player that is trainer ('coach') for the team. (Not to be
	// confused with the 'owner'/'user').
	Trainer struct {
		PlayerID id.Player `xml:"PlayerID"`
	} `xml:"Trainer"`

	// The home page URL that the team has specified.
	HomePage string `xml:"HomePage"`

	// Data about the team's cup status. If the team is playing a match
	// the container is empty.
	Cup struct {

		// Boolean value stating if the team is still in the cup.
		StillInCup bool `xml:"StillInCup"`

		// The globally unique CupID of the cup that the team is still
		// in. Only provided if team is still in cup.
		CupID id.Cup `xml:"CupID"`

		// The name of the cup that the team is still in. Only provided
		// if team is still in cup.
		CupName string `xml:"CupName"`

		// LeagueLevel for the cup. 0 = National (LeagueLevel 1-6),
		// 7-9 = Divisional. Only provided if team is still in cup.
		CupLeagueLevel uint `xml:"CupLeagueLevel"`

		// Level of the cup. 1 = National/Divisional, 2 = Challenger,
		// 3 = Consolation. Only provided if team is still in cup.
		CupLevel uint `xml:"CupLevel"`

		// Index of the cup. Always 1 for National and Consolation cups,
		// for Challenger cup: 1 = Emerald, 2 = Ruby, 3 = Sapphire. Only
		// provided if team is still in cup.
		CupLevelIndex uint `xml:"CupLevelIndex"`

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
	LeagueLevelUnit struct {
		// The globally unique LeagueLevelUnitID.
		LeagueLevelUnitID id.LeagueLevelUnit `xml:"LeagueLevelUnitID"`

		// The name of the LeagueLevelUnit.
		LeagueLevelUnitName string `xml:"LeagueLevelUnitName"`

		// Integer value stating the relative level of the
		// LeagueLevelUnit, with 1 indicating the league's top series.
		LeagueLevel uint `xml:"LeagueLevel"`
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
		FanclubID id.Fanclub `xml:"FanclubID"`

		// The name of the fanclub.
		FanclubName string `xml:"FanclubName"`

		// The amount of members in the fanclub.
		FanclubSize uint `xml:"FanclubSize"`
	} `xml:"Fanclub"`

	// The logo URL that the team has specified.
	LogoURL string `xml:"LogoURL"`

	// This container and its elements is only shown if the user has
	// supporter
	// Container for the data about the team's guestbook.
	Guestbook *struct {
		NumberOfGuestbookItems uint `xml:"NumberOfGuestbookItems"`
	} `xml:"Guestbook"`

	// This container and its elements is only shown if the user has
	// supporter
	// Container for the data about the team's most recent
	// PressAnnouncement.
	PressAnnouncement struct {

		// The subject text specified for the PressAnnouncement.
		Subject string `xml:"Subject"`

		// The body of the PressAnnouncement.
		Body string `xml:"Body"`

		// The time and date when the PressAnnouncement was submitted.
		SendDate HattrickTime `xml:"SendDate"`
	} `xml:"PressAnnouncement"`

	// This container and its elements is only shown if the user has
	// supporter
	// Container for the team's club theme colors. Empty if no theme has
	// been set for the club.
	TeamColors struct {
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
		BotSinceDate HattrickTime `xml:"BotSince"`
	} `xml:"BotStatus"`

	// If the team has an owner, last is the League Rank, a number based
	// on LeagueLevel, position etc, only counting teams with an owner.
	// If the team is playing a match the tag will be empty.
	TeamRank uint `xml:"TeamRank"`

	// If the Team has a Youth academy the ID is represented here. If
	// not it will be 0.
	YouthTeamID id.YouthTeam `xml:"YouthTeamID"`

	// If the Team has a Youth academy the team name is represented
	// here. If not it will be an empty string.
	YouthTeamName string `xml:"YouthTeamName"`

	Flags struct {
		AwayFlags struct {
			Flag []*Flag `xml:"Flag"`
		} `xml:"AwayFlags"`
		HomeFlags struct {
			Flag []*Flag `xml:"Flag"`
		} `xml:"HomeFlags"`
	} `xml:"Flags"`

	TrophyList struct {
		Trophies []*struct {
			TrophyTypeID        TrophyID     `xml:"TrophyTypeId"`
			TrophySeason        uint         `xml:"TrophySeason"`
			LeagueLevel         uint         `xml:"LeagueLevel"`
			LeagueLevelUnitName string       `xml:"LeagueLevelUnitName"`
			GainedDate          HattrickTime `xml:"GainedDate"`
			ImageURL            string       `xml:"ImageUrl"`
		} `xml:"Trophy"`
	} `xml:"TrophyList"`

	// Container for the the supported teams. It will be empty if the
	// user have no supporters and not present if the user are not a
	// supporter. Only available, when includeSupporters=true.
	SupportedTeams *struct {
		TotalItems *uint `xml:"TotalItems,attr"`
		MaxItems   *uint `xml:"MaxItems,attr"`

		SupportedTeam []*struct {
			SupportedTeam

			// Container for the last match of the supported team.
			LastMatch struct {
				LastMatchID           id.Match     `xml:"LastMatchId"`
				LastMatchDate         HattrickTime `xml:"LastMatchDate"`
				LastMatchHomeTeamID   id.Team      `xml:"LastMatchHomeTeamId"`
				LastMatchHomeTeamName string       `xml:"LastMatchHomeTeamName"`
				LastMatchHomeGoals    uint         `xml:"LastMatchHomeGoals"`
				LastMatchAwayTeamID   id.Team      `xml:"LastMatchAwayTeamId"`
				LastMatchAwayTeamName string       `xml:"LastMatchAwayTeamName"`
				LastMatchAwayGoals    uint         `xml:"LastMatchAwayGoals"`
			} `xml:"LastMatch"`

			// Container for the last match of the supported team.
			NextMatch struct {
				NextMatchID           id.Match     `xml:"NextMatchId"`
				NextMatchDate         HattrickTime `xml:"NextMatchDate"`
				NextMatchHomeTeamID   id.Team      `xml:"NextMatchHomeTeamId"`
				NextMatchHomeTeamName string       `xml:"NextMatchHomeTeamName"`
				NextMatchAwayTeamID   id.Team      `xml:"NextMatchAwayTeamId"`
				NextMatchAwayTeamName string       `xml:"NextMatchAwayTeamName"`
			} `xml:"NextMatch"`

			PressAnnouncement struct {
				PressAnnouncementSendDate HattrickTime `xml:"PressAnnouncementSendDate"`
				PressAnnouncementSubject  string       `xml:"PressAnnouncementSubject"`
				PressAnnouncementBody     string       `xml:"PressAnnouncementBody"`
			} `xml:"PressAnnouncement"`
		} `xml:"SupportedTeam"`
	} `xml:"SupportedTeams"`

	// Attribute: 'TotalItems' : unsigned Integer
	//   Shows how many teams supports the team in total
	// Attribute: 'MaxItems' : unsigned Integer
	//   Shows how many teams that can maximum be outputted in this list
	// Container for the the supported teams. It will be empty if the
	// user have no supporters and not present if the user are not a
	// supporter. Only available, when includeSupporters=true.
	MySupporters *struct {
		TotalItems *uint `xml:"TotalItems,attr"`
		MaxItems   *uint `xml:"MaxItems,attr"`

		// Container for a supported team. Not supplied if the user dont
		// have no supporters.
		SupporterTeam []*SupportedTeam `xml:"SupporterTeam"`
	} `xml:"MySupporters"`

	// Undocumented
	NumberOfVisits uint `xml:"NumberOfVisits"`
}

// Flag ...
type Flag struct {
	LeagueID    id.League `xml:"LeagueId"`
	LeagueName  string    `xml:"LeagueName"`
	CountryCode string    `xml:"CountryCode"`
}

// SupportedTeam ...
type SupportedTeam struct {
	UserID              id.User            `xml:"UserId"`
	LoginName           string             `xml:"LoginName"`
	TeamID              id.Team            `xml:"TeamId"`
	TeamName            string             `xml:"TeamName"`
	LeagueID            id.League          `xml:"LeagueID"`
	LeagueName          string             `xml:"LeagueName"`
	LeagueLevelUnitID   id.LeagueLevelUnit `xml:"LeagueLevelUnitID"`
	LeagueLevelUnitName string             `xml:"LeagueLevelUnitName"`
}
