package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TranslationsAPIFile = "translations"
	// Confirmed against a live v1.2 response: the changelog's "Added
	// form, indirect set pieces, attack and defence" did NOT add new
	// top-level containers (as the doc's own JSON/markdown dumps had
	// implied) - it added new items within the pre-existing SkillNames
	// ("Form") and RatingSectors ("Indirect set pieces", "Defense",
	// "Attack") lists, which this package's slice-based modeling already
	// handles without any struct changes.
	TranslationsAPIVersion = "1.2"
)

// TranslationsXML contains the localized text tables (skill names, levels,
// specialties, tactic types, etc.) for a single requested language.
type TranslationsXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// The language sent in.
	Language struct {
		ID   uint   `xml:"Id,attr"`
		Name string `xml:",chardata"`
	} `xml:"Language"`

	Texts TranslationTexts `xml:"Texts"`
}

// TranslationTexts is the container for all the translated text tables.
type TranslationTexts struct {
	// No Label attribute on these three, confirmed against a live
	// response.
	SkillNames []*TranslationTypeValue `xml:"SkillNames>Skill"`

	SkillLevels []*TranslationValue `xml:"SkillLevels>Level"`

	// NOTE: the doc's own JSON dump lists this element's id as
	// "SubLevel " with a trailing space, which is not a legal XML tag
	// name and can never appear on the wire - almost certainly a
	// doc-extraction artifact. Tagged as "SubLevel" (no trailing space),
	// confirmed against a live response.
	SkillSubLevels []*TranslationValue `xml:"SkillSubLevels>SubLevel"`

	PlayerSpecialties TranslationItemList `xml:"PlayerSpecialties"`

	PlayerAgreeability TranslationLevelList `xml:"PlayerAgreeability"`

	// NOTE: "PlayerAgressiveness" (single "g") is the file's actual
	// documented spelling - differs from data-types.md's
	// "PlayerAggressiveness". Kept verbatim, confirmed against a live
	// response.
	PlayerAgressiveness TranslationLevelList `xml:"PlayerAgressiveness"`

	PlayerHonesty TranslationLevelList `xml:"PlayerHonesty"`

	TacticTypes TranslationItemList `xml:"TacticTypes"`

	// NOTE: unlike every other container here, MatchPositions and
	// RatingSectors key their items by a "Type" attribute (a string
	// identifier), not "Value" - confirmed against a live response.
	MatchPositions TranslationTypeItemList `xml:"MatchPositions"`
	RatingSectors  TranslationTypeItemList `xml:"RatingSectors"`

	TeamAttitude TranslationLevelList `xml:"TeamAttitude"`
	TeamSpirit   TranslationLevelList `xml:"TeamSpirit"`
	Confidence   TranslationLevelList `xml:"Confidence"`

	TrainingTypes TranslationItemList `xml:"TrainingTypes"`

	Sponsors TranslationLevelList `xml:"Sponsors"`

	FanMood               TranslationLevelList `xml:"FanMood"`
	FanMatchExpectations  TranslationLevelList `xml:"FanMatchExpectations"`
	FanSeasonExpectations TranslationLevelList `xml:"FanSeasonExpectations"`

	// Only sent since API version 1.1. No Label attribute, confirmed
	// against a live response.
	LeagueNames []*TranslationLeagueName `xml:"LeagueNames>League"`
}

// TranslationItemList is a labeled list of translated text entries, keyed
// by a "Value" attribute, whose repeated child element is named "Item".
type TranslationItemList struct {
	// A human-readable label for this table. Only sent since API
	// version 1.2.
	Label string              `xml:"Label,attr"`
	Items []*TranslationValue `xml:"Item"`
}

// TranslationLevelList is a labeled list of translated text entries, keyed
// by a "Value" attribute, whose repeated child element is named "Level".
type TranslationLevelList struct {
	// A human-readable label for this table. Only sent since API
	// version 1.2.
	Label string              `xml:"Label,attr"`
	Items []*TranslationValue `xml:"Level"`
}

// TranslationTypeItemList is a labeled list of translated text entries,
// keyed by a "Type" attribute (a string identifier, not "Value"), whose
// repeated child element is named "Item".
type TranslationTypeItemList struct {
	// A human-readable label for this table. Only sent since API
	// version 1.2.
	Label string                  `xml:"Label,attr"`
	Items []*TranslationTypeValue `xml:"Item"`
}

// TranslationTypeValue is a single translated text entry keyed by a "Type"
// string identifier (not a numeric id).
type TranslationTypeValue struct {
	Type string `xml:"Type,attr"`
	Text string `xml:",chardata"`
}

// TranslationValue is a single translated text entry for an enum value.
// Value is documented as a plain string identifier in every container it
// appears in (not consistently a numeric id), so it is kept untyped.
type TranslationValue struct {
	Value string `xml:"Value,attr"`
	Text  string `xml:",chardata"`
}

// TranslationLeagueName is the local and chosen-language names of a
// league.
type TranslationLeagueName struct {
	LeagueID           id.League `xml:"LeagueId"`
	LocalLeagueName    string    `xml:"LocalLeagueName"`
	LanguageLeagueName string    `xml:"LanguageLeagueName"`
}
