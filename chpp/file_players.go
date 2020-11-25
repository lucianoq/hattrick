package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	PlayersAPIFile    = "players"
	PlayersAPIVersion = "2.4"
)

// PlayersXML ...
type PlayersXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`
	IsYouth           bool          `xml:"IsYouth"`
	ActionType        string        `xml:"ActionType"`
	IsPlayingMatch    bool          `xml:"IsPlayingMatch"`
	Team              struct {
		TeamID     id.Team `xml:"TeamID"`
		TeamName   string  `xml:"TeamName"`
		PlayerList struct {
			Players []*Player `xml:"Player"`
		} `xml:"PlayerList"`
	} `xml:"Team"`
}

// Player ...
type Player struct {
	PlayerID           id.Player            `xml:"PlayerID"`
	FirstName          string               `xml:"FirstName"`
	NickName           string               `xml:"NickName"`
	LastName           string               `xml:"LastName"`
	PlayerNumber       uint                 `xml:"PlayerNumber"`
	Age                uint                 `xml:"Age"`
	AgeDays            uint                 `xml:"AgeDays"`
	ArrivalDate        HattrickTime         `xml:"ArrivalDate"`
	OwnerNotes         string               `xml:"OwnerNotes"`
	TSI                TSI                  `xml:"TSI"`
	PlayerForm         SkillLevel           `xml:"PlayerForm"`
	Statement          string               `xml:"Statement"`
	Experience         SkillLevel           `xml:"Experience"`
	Loyalty            SkillLevel           `xml:"Loyalty"`
	MotherClubBonus    bool                 `xml:"MotherClubBonus"`
	Leadership         SkillLevel           `xml:"Leadership"`
	Salary             Money                `xml:"Salary"`
	IsAbroad           bool                 `xml:"IsAbroad"`
	Agreeability       PlayerAgreeability   `xml:"Agreeability"`
	Aggressiveness     PlayerAggressiveness `xml:"Aggressiveness"`
	Honesty            PlayerHonesty        `xml:"Honesty"`
	LeagueGoals        uint                 `xml:"LeagueGoals"`
	CupGoals           uint                 `xml:"CupGoals"`
	FriendliesGoals    uint                 `xml:"FriendliesGoals"`
	CareerGoals        uint                 `xml:"CareerGoals"`
	CareerHattricks    uint                 `xml:"CareerHattricks"`
	MatchesCurrentTeam uint                 `xml:"MatchesCurrentTeam"`
	GoalsCurrentTeam   uint                 `xml:"GoalsCurrentTeam"`
	Specialty          SpecialtyID          `xml:"Specialty"`
	TransferListed     bool                 `xml:"TransferListed"`
	NationalTeamID     *id.NationalTeam     `xml:"NationalTeamID"`
	CountryID          id.Country           `xml:"CountryID"`
	Caps               uint                 `xml:"Caps"`
	CapsU20            uint                 `xml:"CapsU20"`
	Cards              uint                 `xml:"Cards"`
	InjuryLevel        PlayerInjuryLevel    `xml:"InjuryLevel"`
	StaminaSkill       SkillLevel           `xml:"StaminaSkill"`
	KeeperSkill        SkillLevel           `xml:"KeeperSkill"`
	PlaymakerSkill     SkillLevel           `xml:"PlaymakerSkill"`
	ScorerSkill        SkillLevel           `xml:"ScorerSkill"`
	PassingSkill       SkillLevel           `xml:"PassingSkill"`
	WingerSkill        SkillLevel           `xml:"WingerSkill"`
	DefenderSkill      SkillLevel           `xml:"DefenderSkill"`
	SetPiecesSkill     SkillLevel           `xml:"SetPiecesSkill"`
	PlayerCategoryID   PlayerCategoryID     `xml:"PlayerCategoryId"`

	// Container for the skills that relate to coach/trainer ability. Only
	// provided if the player has been made into a trainer, 'has gained trainer
	// license'.
	TrainerData struct {
		TrainerType  TrainerType `xml:"TrainerType"`
		TrainerSkill SkillLevel  `xml:"TrainerSkill"`
	} `xml:"TrainerData"`

	// Last played match. If MatchId = 0 it either means that the player has not
	// played a match yet, or that he's in a bot team.
	LastMatch struct {
		Date            HattrickTime `xml:"Date"`
		MatchID         id.Match     `xml:"MatchId"`
		PositionCode    MatchRole    `xml:"PositionCode"`
		PlayedMinutes   uint         `xml:"PlayedMinutes"`
		Rating          float64      `xml:"Rating"`
		RatingEndOfGame float64      `xml:"RatingEndOfGame"`
	} `xml:"LastMatch"`
}
