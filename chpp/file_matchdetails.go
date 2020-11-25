package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchDetailsAPIFile    = "matchdetails"
	MatchDetailsAPIVersion = "3.0"
)

// MatchDetailsXML ...
type MatchDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	UserSupporterTier string       `xml:"UserSupporterTier"`
	SourceSystem      string       `xml:"SourceSystem"`
	Match             MatchDetails `xml:"Match"`
}

// MatchDetails ...
type MatchDetails struct {
	MatchID        id.Match     `xml:"MatchID"`
	MatchType      MatchType    `xml:"MatchType"`
	MatchContextID string       `xml:"MatchContextId"`
	MatchRuleID    MatchRule    `xml:"MatchRuleId"`
	CupLevel       uint         `xml:"CupLevel"`
	CupLevelIndex  uint         `xml:"CupLevelIndex"`
	MatchDate      HattrickTime `xml:"MatchDate"`
	FinishedDate   HattrickTime `xml:"FinishedDate"`
	AddedMinutes   int          `xml:"AddedMinutes"`
	HomeTeam       struct {
		HomeTeamID   id.Team `xml:"HomeTeamID"`
		HomeTeamName string  `xml:"HomeTeamName"`
		DressURI     string  `xml:"DressURI"`
		// Character string (x-x-x) representing the formation the team started with.
		Formation                  string             `xml:"Formation"`
		HomeGoals                  uint               `xml:"HomeGoals"`
		TacticType                 MatchTacticType    `xml:"TacticType"`
		TacticSkill                SkillLevel         `xml:"TacticSkill"`
		RatingMidfield             MatchRating        `xml:"RatingMidfield"`
		RatingRightDef             MatchRating        `xml:"RatingRightDef"`
		RatingMidDef               MatchRating        `xml:"RatingMidDef"`
		RatingLeftDef              MatchRating        `xml:"RatingLeftDef"`
		RatingRightAtt             MatchRating        `xml:"RatingRightAtt"`
		RatingMidAtt               MatchRating        `xml:"RatingMidAtt"`
		RatingLeftAtt              MatchRating        `xml:"RatingLeftAtt"`
		RatingIndirectSetPiecesDef MatchRating        `xml:"RatingIndirectSetPiecesDef"`
		RatingIndirectSetPiecesAtt MatchRating        `xml:"RatingIndirectSetPiecesAtt"`
		TeamAttitude               *MatchTeamAttitude `xml:"TeamAttitude"`
	} `xml:"HomeTeam"`
	AwayTeam struct {
		AwayTeamID   id.Team `xml:"AwayTeamID"`
		AwayTeamName string  `xml:"AwayTeamName"`
		DressURI     string  `xml:"DressURI"`
		// Character string (x-x-x) representing the formation the team started with.
		Formation                  string             `xml:"Formation"`
		AwayGoals                  uint               `xml:"AwayGoals"`
		TacticType                 MatchTacticType    `xml:"TacticType"`
		TacticSkill                SkillLevel         `xml:"TacticSkill"`
		RatingMidfield             MatchRating        `xml:"RatingMidfield"`
		RatingRightDef             MatchRating        `xml:"RatingRightDef"`
		RatingMidDef               MatchRating        `xml:"RatingMidDef"`
		RatingLeftDef              MatchRating        `xml:"RatingLeftDef"`
		RatingRightAtt             MatchRating        `xml:"RatingRightAtt"`
		RatingMidAtt               MatchRating        `xml:"RatingMidAtt"`
		RatingLeftAtt              MatchRating        `xml:"RatingLeftAtt"`
		RatingIndirectSetPiecesDef MatchRating        `xml:"RatingIndirectSetPiecesDef"`
		RatingIndirectSetPiecesAtt MatchRating        `xml:"RatingIndirectSetPiecesAtt"`
		TeamAttitude               *MatchTeamAttitude `xml:"TeamAttitude"`
	} `xml:"AwayTeam"`
	Arena struct {
		ArenaID      id.Arena `xml:"ArenaID"`
		ArenaName    string   `xml:"ArenaName"`
		Weather      Weather  `xml:"WeatherID"`
		SoldTotal    uint     `xml:"SoldTotal"`
		SoldTerraces uint     `xml:"SoldTerraces"`
		SoldBasic    uint     `xml:"SoldBasic"`
		SoldRoof     uint     `xml:"SoldRoof"`
		SoldVIP      uint     `xml:"SoldVIP"`
	} `xml:"Arena"`
	MatchOfficials struct {
		Referee struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			RefereeTeamID      id.Team    `xml:"RefereeTeamId"`
			RefereeTeamName    string     `xml:"RefereeTeamname"`
		} `xml:"Referee"`
		RefereeAssistant1 struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			RefereeTeamID      id.Team    `xml:"RefereeTeamId"`
			RefereeTeamName    string     `xml:"RefereeTeamname"`
		} `xml:"RefereeAssistant1"`
		RefereeAssistant2 struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			RefereeTeamID      id.Team    `xml:"RefereeTeamId"`
			RefereeTeamName    string     `xml:"RefereeTeamname"`
		} `xml:"RefereeAssistant2"`
	} `xml:"MatchOfficials"`
	Scorers struct {
		Goal []struct {
			Index            uint      `xml:"Index,attr"`
			ScorerPlayerID   id.Player `xml:"ScorerPlayerID"`
			ScorerPlayerName string    `xml:"ScorerPlayerName"`
			ScorerTeamID     id.Team   `xml:"ScorerTeamID"`
			ScorerHomeGoals  uint      `xml:"ScorerHomeGoals"`
			ScorerAwayGoals  uint      `xml:"ScorerAwayGoals"`
			ScorerMinute     uint      `xml:"ScorerMinute"`
			MatchPart        MatchPart `xml:"MatchPart"`
		} `xml:"Goal"`
	} `xml:"Scorers"`
	Bookings struct {
		Booking []struct {
			Index             uint        `xml:"Index,attr"`
			BookingPlayerID   id.Player   `xml:"BookingPlayerID"`
			BookingPlayerName string      `xml:"BookingPlayerName"`
			BookingTeamID     id.Team     `xml:"BookingTeamID"`
			BookingType       BookingType `xml:"BookingType"`
			BookingMinute     uint        `xml:"BookingMinute"`
			MatchPart         MatchPart   `xml:"MatchPart"`
		} `xml:"Booking"`
	} `xml:"Bookings"`
	Injuries struct {
		Injury []struct {
			Index            uint        `xml:"Index,attr"`
			InjuryPlayerID   id.Player   `xml:"InjuryPlayerID"`
			InjuryPlayerName string      `xml:"InjuryPlayerName"`
			InjuryTeamID     id.Team     `xml:"InjuryTeamID"`
			InjuryType       BookingType `xml:"InjuryType"`
			InjuryMinute     uint        `xml:"InjuryMinute"`
			MatchPart        MatchPart   `xml:"MatchPart"`
		} `xml:"Injury"`
	} `xml:"Injuries"`
	PossessionFirstHalfHome  uint `xml:"PossessionFirstHalfHome"`
	PossessionFirstHalfAway  uint `xml:"PossessionFirstHalfAway"`
	PossessionSecondHalfHome uint `xml:"PossessionSecondHalfHome"`
	PossessionSecondHalfAway uint `xml:"PossessionSecondHalfAway"`
}
