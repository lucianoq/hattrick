package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	MatchDetailsAPIFile    = "matchdetails"
	MatchDetailsAPIVersion = "3.1"
)

// MatchDetailsXML contains detailed information about a single match:
// teams, ratings, tactics, arena, officials, scorers, bookings, injuries,
// possession, and (optionally) the play-by-play events.
type MatchDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`
	SourceSystem      SourceSystem  `xml:"SourceSystem"`
	Match             MatchDetails  `xml:"Match"`
}

// MatchDetails is the detailed report of a single match.
type MatchDetails struct {
	MatchID   id.Match  `xml:"MatchID"`
	MatchType MatchType `xml:"MatchType"`

	// This will be either
	// LeagueLevelUnitId (for League),
	// CupId (Cup, Hattrick Masters, World Cup and U-20 World Cup),
	// LadderId, TournamentId, or 0 for friendly, qualification,
	// single matches and preparation matches.
	MatchContextID uint `xml:"MatchContextId"`

	// The rules in place for the ladder or tournament in question, if any.
	// Defaults to 0 = no rules.
	MatchRuleID MatchRule `xml:"MatchRuleId"`

	// 1 = National/Divisional cup, 2 = Challenger cup, 3 = Consolation cup.
	// 0 if MatchType is not a cup match.
	CupLevel uint `xml:"CupLevel"`

	// In Challenger cups: 1 = Emerald (start week 2), 2 = Ruby (start
	// week 3), 3 = Sapphire (start week 4). Always 1 for
	// National/Divisional (main cups) and Consolation cups. 0 if
	// MatchType is not a cup match.
	CupLevelIndex uint `xml:"CupLevelIndex"`

	// The start date and time (kick-off) of the match.
	MatchDate HattrickTime `xml:"MatchDate"`

	// Not sent until the match is finished.
	FinishedDate HattrickTime `xml:"FinishedDate"`

	// How many minutes of added (injury/stoppage) time the match had.
	// Not sent until the match is finished.
	AddedMinutes int `xml:"AddedMinutes"`

	HomeTeam struct {
		// The teamID of the home team. Negative for street teams.
		HomeTeamID   id.Team `xml:"HomeTeamID"`
		HomeTeamName string  `xml:"HomeTeamName"`
		// URI to an image of the team's dress. Only sent for senior teams.
		DressURI string `xml:"DressURI"`
		// Character string (x-x-x) representing the formation the team started with.
		Formation                  string          `xml:"Formation"`
		HomeGoals                  uint            `xml:"HomeGoals"`
		TacticType                 MatchTacticType `xml:"TacticType"`
		TacticSkill                SkillLevel      `xml:"TacticSkill"`
		RatingMidfield             MatchRating     `xml:"RatingMidfield"`
		RatingRightDef             MatchRating     `xml:"RatingRightDef"`
		RatingMidDef               MatchRating     `xml:"RatingMidDef"`
		RatingLeftDef              MatchRating     `xml:"RatingLeftDef"`
		RatingRightAtt             MatchRating     `xml:"RatingRightAtt"`
		RatingMidAtt               MatchRating     `xml:"RatingMidAtt"`
		RatingLeftAtt              MatchRating     `xml:"RatingLeftAtt"`
		RatingIndirectSetPiecesDef MatchRating     `xml:"RatingIndirectSetPiecesDef"`
		RatingIndirectSetPiecesAtt MatchRating     `xml:"RatingIndirectSetPiecesAtt"`
		// The team attitude set for the match. Only shown to the owner
		// of the team, omitted otherwise (hence the pointer).
		TeamAttitude      *MatchTeamAttitude `xml:"TeamAttitude"`
		NrOfChancesLeft   uint               `xml:"NrOfChancesLeft"`
		NrOfChancesCenter uint               `xml:"NrOfChancesCenter"`
		NrOfChancesRight  uint               `xml:"NrOfChancesRight"`

		// The number of special-event chances.
		NrOfChancesSpecialEvents uint `xml:"NrOfChancesSpecialEvents"`
		// Chances that don't fit any of the other NrOfChances categories.
		NrOfChancesOther uint `xml:"NrOfChancesOther"`
	} `xml:"HomeTeam"`
	AwayTeam struct {
		// The teamID of the away team. Negative for street teams.
		AwayTeamID   id.Team `xml:"AwayTeamID"`
		AwayTeamName string  `xml:"AwayTeamName"`
		// URI to an image of the team's dress. Only sent for senior teams.
		DressURI string `xml:"DressURI"`
		// Character string (x-x-x) representing the formation the team started with.
		Formation                  string          `xml:"Formation"`
		AwayGoals                  uint            `xml:"AwayGoals"`
		TacticType                 MatchTacticType `xml:"TacticType"`
		TacticSkill                SkillLevel      `xml:"TacticSkill"`
		RatingMidfield             MatchRating     `xml:"RatingMidfield"`
		RatingRightDef             MatchRating     `xml:"RatingRightDef"`
		RatingMidDef               MatchRating     `xml:"RatingMidDef"`
		RatingLeftDef              MatchRating     `xml:"RatingLeftDef"`
		RatingRightAtt             MatchRating     `xml:"RatingRightAtt"`
		RatingMidAtt               MatchRating     `xml:"RatingMidAtt"`
		RatingLeftAtt              MatchRating     `xml:"RatingLeftAtt"`
		RatingIndirectSetPiecesDef MatchRating     `xml:"RatingIndirectSetPiecesDef"`
		RatingIndirectSetPiecesAtt MatchRating     `xml:"RatingIndirectSetPiecesAtt"`
		// The team attitude set for the match. Only shown to the owner
		// of the team, omitted otherwise (hence the pointer).
		TeamAttitude      *MatchTeamAttitude `xml:"TeamAttitude"`
		NrOfChancesLeft   uint               `xml:"NrOfChancesLeft"`
		NrOfChancesCenter uint               `xml:"NrOfChancesCenter"`
		NrOfChancesRight  uint               `xml:"NrOfChancesRight"`

		// The number of special-event chances.
		NrOfChancesSpecialEvents uint `xml:"NrOfChancesSpecialEvents"`
		// Chances that don't fit any of the other NrOfChances categories.
		NrOfChancesOther uint `xml:"NrOfChancesOther"`
	} `xml:"AwayTeam"`
	Arena struct {
		ArenaID   id.Arena `xml:"ArenaID"`
		ArenaName string   `xml:"ArenaName"`
		Weather   Weather  `xml:"WeatherID"`
		SoldTotal uint     `xml:"SoldTotal"`

		// The number of sold tickets for the terraces/basic/roof/VIP
		// sections, respectively.
		SoldTerraces uint `xml:"SoldTerraces"`
		SoldBasic    uint `xml:"SoldBasic"`
		SoldRoof     uint `xml:"SoldRoof"`
		SoldVIP      uint `xml:"SoldVIP"`
	} `xml:"Arena"`
	// Only sent for matches with SourceSystem=Hattrick.
	MatchOfficials struct {
		Referee struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			// The team the referee is a Hall of Fame member of.
			RefereeTeamID   id.Team `xml:"RefereeTeamId"`
			RefereeTeamName string  `xml:"RefereeTeamname"`
		} `xml:"Referee"`
		RefereeAssistant1 struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			// The team the referee is a Hall of Fame member of.
			RefereeTeamID   id.Team `xml:"RefereeTeamId"`
			RefereeTeamName string  `xml:"RefereeTeamname"`
		} `xml:"RefereeAssistant1"`
		RefereeAssistant2 struct {
			RefereeID          id.Referee `xml:"RefereeId"`
			RefereeName        string     `xml:"RefereeName"`
			RefereeCountryID   id.Country `xml:"RefereeCountryId"`
			RefereeCountryName string     `xml:"RefereeCountryName"`
			// The team the referee is a Hall of Fame member of.
			RefereeTeamID   id.Team `xml:"RefereeTeamId"`
			RefereeTeamName string  `xml:"RefereeTeamname"`
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
			Index            uint      `xml:"Index,attr"`
			InjuryPlayerID   id.Player `xml:"InjuryPlayerID"`
			InjuryPlayerName string    `xml:"InjuryPlayerName"`
			InjuryTeamID     id.Team   `xml:"InjuryTeamID"`
			// Reuses the BookingType XML shape, but the values mean
			// something different here: 1 = bruise, 2 = injury.
			InjuryType   BookingType `xml:"InjuryType"`
			InjuryMinute uint        `xml:"InjuryMinute"`
			MatchPart    MatchPart   `xml:"MatchPart"`
		} `xml:"Injury"`
	} `xml:"Injuries"`

	// Ball possession in % for each team/half.
	PossessionFirstHalfHome  uint `xml:"PossessionFirstHalfHome"`
	PossessionFirstHalfAway  uint `xml:"PossessionFirstHalfAway"`
	PossessionSecondHalfHome uint `xml:"PossessionSecondHalfHome"`
	PossessionSecondHalfAway uint `xml:"PossessionSecondHalfAway"`

	// Container for the play-by-play match events. Only sent when
	// matchEvents=true.
	EventList struct {
		Events []*MatchDetailsEvent `xml:"Event"`
	} `xml:"EventList"`
}

// MatchDetailsEvent is a single play-by-play match event.
type MatchDetailsEvent struct {
	Minute    uint      `xml:"Minute"`
	MatchPart MatchPart `xml:"MatchPart"`

	// A unique key defining what type of event it was.
	EventTypeID uint `xml:"EventTypeID"`

	// The variation defines which text is used to describe the event.
	EventVariation uint `xml:"EventVariation"`

	// String describing the event as it would appear in the match report.
	EventText string `xml:"EventText"`

	// For goals and chances, indicates the attacking team. For other
	// events, usually indicates the team that is doing something.
	SubjectTeamID id.Team `xml:"SubjectTeamID"`

	// For goals and chances, indicates the scorer or the player who
	// failed to score. For other events, usually indicates the
	// primarily active player.
	SubjectPlayerID id.Player `xml:"SubjectPlayerID"`

	// For regular goals and chances, indicates the defending team's
	// goalkeeper. For special-event chances/goals, often indicates the
	// assisting player of the attacking team.
	ObjectPlayerID id.Player `xml:"ObjectPlayerID"`
}
