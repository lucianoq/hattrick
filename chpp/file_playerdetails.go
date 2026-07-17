package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	PlayerDetailsAPIFile    = "playerdetails"
	PlayerDetailsAPIVersion = "3.2"
)

// PlayerDetailsXML contains detailed information about a single player,
// optionally including a bid placed on him if he is transfer listed.
type PlayerDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	Player PlayerDetails `xml:"Player"`
}

// PlayerDetails is the full detailed information for a single player.
type PlayerDetails struct {
	ID id.Player `xml:"PlayerID"`

	FirstName string `xml:"FirstName"`
	NickName  string `xml:"NickName"`
	LastName  string `xml:"LastName"`

	// The number (1 to 99) that the owner of the team has assigned this
	// player.
	Number uint `xml:"PlayerNumber"`

	// The category assigned by the owner. Only visible to the owner.
	CategoryID PlayerCategoryID `xml:"PlayerCategoryID"`

	// The notes the owner has set. Only visible to the owner.
	OwnerNotes string `xml:"OwnerNotes"`

	Age uint `xml:"Age"`

	// The number of days since the player's last birthday.
	AgeDays uint `xml:"AgeDays"`

	// The approximate date/time of the player's next birthday.
	NextBirthday HattrickTime `xml:"NextBirthDay"`

	Gender GenderID `xml:"GenderID"`

	// The date of arrival to the team (via transfer, youth promotion,
	// etc).
	ArrivalDate HattrickTime `xml:"ArrivalDate"`

	Form SkillLevel `xml:"PlayerForm"`

	// The number of currently accumulated cards. If the player is
	// suspended, this returns 3 regardless of whether he actually
	// accumulated 3 bookings or was sent off after 2 bookings in the
	// same game. Unavailable while a match is running.
	Cards uint `xml:"Cards"`

	// Signed: -1 = healthy, 0 = bruised, >0 = weeks predicted injured.
	// Unavailable while a match is running.
	InjuryLevel PlayerInjuryLevel `xml:"InjuryLevel"`

	// If the owner is a HT Supporter, he may have provided a statement
	// from the player.
	Statement string `xml:"Statement"`

	// The language the player speaks in his Statement, if any.
	Language string `xml:"PlayerLanguage"`

	// The languageID the player speaks in his Statement, if any.
	LanguageID uint `xml:"PlayerLanguageID"`

	// Container for the skills that relate to coach/trainer ability. Only
	// provided if the player has been made into a trainer.
	TrainerData struct {
		TrainerType TrainerType `xml:"TrainerType"`

		// The trainer skill level, from 1 (lowest) to 5 (highest).
		TrainerSkillLevel uint `xml:"TrainerSkillLevel"`
	} `xml:"TrainerData"`

	Agreeability   PlayerAgreeability   `xml:"Agreeability"`
	Aggressiveness PlayerAggressiveness `xml:"Aggressiveness"`
	Honesty        PlayerHonesty        `xml:"Honesty"`

	Experience SkillLevel `xml:"Experience"`

	// Until the loyalty features are activated, Loyalty is always 0 and
	// MotherClubBonus is always false.
	Loyalty         SkillLevel `xml:"Loyalty"`
	MotherClubBonus bool       `xml:"MotherClubBonus"`

	// The player's mother club, if any.
	MotherClub struct {
		TeamID   id.Team `xml:"TeamID"`
		TeamName string  `xml:"TeamName"`
	} `xml:"MotherClub"`

	Leadership SkillLevel  `xml:"Leadership"`
	Specialty  SpecialtyID `xml:"Specialty"`

	NativeCountryID  id.Country `xml:"NativeCountryID"`
	NativeLeagueID   id.League  `xml:"NativeLeagueID"`
	NativeLeagueName string     `xml:"NativeLeagueName"`

	TSI uint `xml:"TSI"`

	OwningTeam struct {
		TeamID   id.Team   `xml:"TeamID"`
		TeamName string    `xml:"TeamName"`
		LeagueID id.League `xml:"LeagueID"`
	} `xml:"OwningTeam"`

	// Empty if the player has no owner.
	Salary Money `xml:"Salary"`

	// Confusingly named: true means the player is in his home country,
	// false means he's with a team abroad. Empty if the player has no
	// owner.
	IsAbroad bool `xml:"IsAbroad"`

	// Only the owner can see these skills, except Stamina since v1.7.
	Skills struct {
		StaminaSkill   SkillLevel `xml:"StaminaSkill"`
		KeeperSkill    SkillLevel `xml:"KeeperSkill"`
		PlaymakerSkill SkillLevel `xml:"PlaymakerSkill"`
		ScorerSkill    SkillLevel `xml:"ScorerSkill"`
		PassingSkill   SkillLevel `xml:"PassingSkill"`
		WingerSkill    SkillLevel `xml:"WingerSkill"`
		DefenderSkill  SkillLevel `xml:"DefenderSkill"`
		SetPiecesSkill SkillLevel `xml:"SetPiecesSkill"`
	} `xml:"PlayerSkills"`

	Caps    uint `xml:"Caps"`
	CapsU20 uint `xml:"CapsU20"`

	// CareerGoals, CareerHattricks, LeagueGoals, FriendliesGoals, and
	// MatchesCurrentTeam are all unavailable while a match is running.
	CareerGoals     uint `xml:"CareerGoals"`
	CareerHattricks uint `xml:"CareerHattricks"`
	LeagueGoals     uint `xml:"LeagueGoals"`

	// If a match is running this value is not available.
	CupGoals struct {
		Available bool `xml:"Available,attr"`
		Value     uint `xml:",chardata"`
	} `xml:"CupGoals"`

	FriendliesGoals uint `xml:"FriendliesGoals"`

	MatchesCurrentTeam uint `xml:"MatchesCurrentTeam"`
	GoalsCurrentTeam   uint `xml:"GoalsCurrentTeam"`
	AssistsCurrentTeam uint `xml:"AssistsCurrentTeam"`
	CareerAssists      uint `xml:"CareerAssists"`

	// Empty (zero value) if the player is not part of a national team.
	NationalTeamID   id.NationalTeam `xml:"NationalTeamID"`
	NationalTeamName string          `xml:"NationalTeamName"`

	TransferListed bool `xml:"TransferListed"`

	// Empty if the player is not transfer listed.
	TransferDetails struct {
		AskingPrice Money `xml:"AskingPrice"`

		// The transfer deadline. Can be even a few hours in the past.
		Deadline HattrickTime `xml:"Deadline"`

		// 0 if no bids have been made.
		HighestBid Money `xml:"HighestBid"`

		// The maximum (autobid) bid amount, only visible to the bidder.
		// Not shown if not set.
		MaxBid Money `xml:"MaxBid"`

		// The team holding the highest bid. Empty if no bids have been
		// made.
		BidderTeam struct {
			TeamID   id.Team `xml:"TeamID"`
			TeamName string  `xml:"TeamName"`
		} `xml:"BidderTeam"`
	} `xml:"TransferDetails"`

	// Only present if the includeMatchInfo input parameter is true. If
	// MatchId = 0 it either means that the player has not played a match
	// yet, or that he's in a bot team.
	LastMatch struct {
		Date HattrickTime `xml:"Date"`

		// NOTE: lowercase "d" - deliberately not normalized to match
		// PlayerID/TeamID elsewhere in this same doc.
		ID id.Match `xml:"MatchId"`

		PositionCode    MatchRole `xml:"PositionCode"`
		PlayedMinutes   uint      `xml:"PlayedMinutes"`
		Rating          float64   `xml:"Rating"`
		RatingEndOfGame float64   `xml:"RatingEndOfGame"`
	} `xml:"LastMatch"`
}
