package chpp

import (
	"encoding/xml"

	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	YouthPlayerDetailsAPIFile    = "youthplayerdetails"
	YouthPlayerDetailsAPIVersion = "1.3"
)

// YouthPlayerDetailsXML contains detailed information about a single youth
// player.
type YouthPlayerDetailsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Container for the data about the youth player.
	YouthPlayer YouthPlayerDetail `xml:"YouthPlayer"`
}

// YouthPlayerDetail is the full detailed information for a youth player,
// as returned by youthplayerdetails and by youthplayerlist when
// actionType is "details" or "unlockskills". Fields that only appear in
// the "list" view of youthplayerlist are also present here since a
// leaner response simply leaves the extra fields at their zero value.
type YouthPlayerDetail struct {
	ID        id.YouthPlayer `xml:"YouthPlayerID"`
	FirstName string         `xml:"FirstName"`
	NickName  string         `xml:"NickName"`
	LastName  string         `xml:"LastName"`

	Age     uint     `xml:"Age"`
	AgeDays uint     `xml:"AgeDays"`
	Gender  GenderID `xml:"GenderID"`

	// The datetime when the player was pulled to the youth team.
	ArrivalDate HattrickTime `xml:"ArrivalDate"`

	// Number of days until the player can be promoted to the senior
	// squad.
	CanBePromotedIn uint `xml:"CanBePromotedIn"`

	// The number (1 to 99) assigned to this player. Only shown to
	// supporters.
	PlayerNumber uint `xml:"PlayerNumber"`

	// If the owner is a HT Supporter, they may have provided a statement
	// from the player.
	Statement string `xml:"Statement"`

	NativeCountryID   id.Country `xml:"NativeCountryID"`
	NativeCountryName string     `xml:"NativeCountryName"`

	// The notes the owner has set. Only visible to the owner.
	OwnerNotes string `xml:"OwnerNotes"`

	// The category assigned by the owner. Only visible to the owner.
	CategoryID PlayerCategoryID `xml:"PlayerCategoryID"`

	Cards       uint              `xml:"Cards"`
	InjuryLevel PlayerInjuryLevel `xml:"InjuryLevel"`

	// The specialty, if the player has any and it is already revealed.
	// 0 if none or unknown.
	Specialty SpecialtyID `xml:"Specialty"`

	CareerGoals     uint `xml:"CareerGoals"`
	CareerHattricks uint `xml:"CareerHattricks"`
	LeagueGoals     uint `xml:"LeagueGoals"`
	FriendlyGoals   uint `xml:"FriendlyGoals"`

	// The youth team that owns the player.
	OwningYouthTeam YouthPlayerOwningTeam `xml:"OwningYouthTeam"`

	// The skills of the player. Only visible to the owner.
	Skills YouthPlayerSkills `xml:"PlayerSkills"`

	// Information about the scout call. Only present if the
	// showScoutCall input parameter is true.
	ScoutCall *YouthScoutCall `xml:"ScoutCall"`

	// The last played match. Only present if the showLastMatch input
	// parameter is true.
	LastMatch *YouthPlayerLastMatch `xml:"LastMatch"`
}

// YouthPlayerOwningTeam identifies the youth team that owns a youth player,
// and that youth team's parent senior team.
type YouthPlayerOwningTeam struct {
	// The youth team owning the youth player.
	YouthTeamID   id.YouthTeam `xml:"YouthTeamID"`
	YouthTeamName string       `xml:"YouthTeamName"`

	// The league of the youth team owning the youth player.
	YouthTeamLeagueID id.YouthLeague `xml:"YouthTeamLeagueID"`

	// The senior team the youth team belongs to.
	SeniorTeam struct {
		ID   id.Team `xml:"SeniorTeamID"`
		Name string  `xml:"SeniorTeamName"`
	} `xml:"SeniorTeam"`
}

// YouthSkillValue is a skill value together with flags describing its
// visibility and trainability. IsMaxReached is only meaningful for the
// current-skill fields, not the skill-cap ("Max") fields. Value is the
// zero value if Available is false.
type YouthSkillValue struct {
	Available    bool `xml:"IsAvailable,attr"`
	IsMaxReached bool `xml:"IsMaxReached,attr"`
	MayUnlock    bool `xml:"MayUnlock,attr"`
	Value        uint `xml:",chardata"`
}

// YouthPlayerSkills is the container for the skills of a youth player.
// Only visible to the owner.
type YouthPlayerSkills struct {
	Keeper    YouthSkillValue `xml:"KeeperSkill"`
	KeeperMax YouthSkillValue `xml:"KeeperSkillMax"`

	Defender    YouthSkillValue `xml:"DefenderSkill"`
	DefenderMax YouthSkillValue `xml:"DefenderSkillMax"`

	Playmaker    YouthSkillValue `xml:"PlaymakerSkill"`
	PlaymakerMax YouthSkillValue `xml:"PlaymakerSkillMax"`

	Winger    YouthSkillValue `xml:"WingerSkill"`
	WingerMax YouthSkillValue `xml:"WingerSkillMax"`

	Scorer    YouthSkillValue `xml:"ScorerSkill"`
	ScorerMax YouthSkillValue `xml:"ScorerSkillMax"`

	Passing    YouthSkillValue `xml:"PassingSkill"`
	PassingMax YouthSkillValue `xml:"PassingSkillMax"`

	SetPieces    YouthSkillValue `xml:"SetPiecesSkill"`
	SetPiecesMax YouthSkillValue `xml:"SetPiecesSkillMax"`
}

// YouthScoutCall holds the information about the scout call that led to a
// youth player being pulled. Only present if the showScoutCall input
// parameter is true.
type YouthScoutCall struct {
	// The scout who suggested this youth player.
	Scout YouthScoutCallScout `xml:"Scout"`

	// The region the scout was looking for players in while he scouted
	// this youth player.
	ScoutingRegionID id.Region `xml:"ScoutingRegionID"`

	Comments []*YouthScoutComment `xml:"ScoutComments>ScoutComment"`
}

// YouthScoutCallScout identifies the scout who suggested a youth player. Its ID
// field is tagged ScoutId for actionType=details but ScoutID for
// actionType=unlockskills; a custom UnmarshalXML handles both casings.
type YouthScoutCallScout struct {
	ID   uint
	Name string
}

// UnmarshalXML implements custom unmarshaling to handle the ScoutId
// (details) / ScoutID (unlockskills) tag-casing inconsistency.
func (s *YouthScoutCallScout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw struct {
		IDDetails      uint   `xml:"ScoutId"`
		IDUnlockSkills uint   `xml:"ScoutID"`
		Name           string `xml:"ScoutName"`
	}
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	s.Name = raw.Name
	if raw.IDDetails != 0 {
		s.ID = raw.IDDetails
	} else {
		s.ID = raw.IDUnlockSkills
	}

	return nil
}

// YouthScoutComment is a single comment made by a scout about a youth
// player.
type YouthScoutComment struct {
	Text string           `xml:"CommentText"`
	Type ScoutCommentType `xml:"CommentType"`

	// The text variant of the comment for the given Type.
	Variation uint `xml:"CommentVariation"`

	// The skill named in the comment (0 if there is no skill). Will be
	// the youthPlayerId if Type=1.
	SkillType ScoutCommentSkillType `xml:"CommentSkillType"`

	// The skill level named in the comment (0 if there is no skill
	// level named). Will be the player age if Type=1.
	SkillLevel SkillLevel `xml:"CommentSkillLevel"`
}

// YouthPlayerLastMatch is the last played match of a youth player. Only
// present if the showLastMatch input parameter is true.
type YouthPlayerLastMatch struct {
	Date         HattrickTime `xml:"Date"`
	YouthMatchID uint         `xml:"YouthMatchID"`

	// The position played in the last match.
	PositionCode MatchRole `xml:"PositionCode"`

	PlayedMinutes uint    `xml:"PlayedMinutes"`
	Rating        float64 `xml:"Rating"`
}
