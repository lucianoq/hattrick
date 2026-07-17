package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TrainingAPIFile    = "training"
	TrainingAPIVersion = "2.2"
)

// TrainingXML contains a team's training settings, league-wide training
// statistics, or the result of setting a new training, depending on the
// requested actionType.
type TrainingXML struct {
	Envelope
	UserID id.User `xml:"User"`

	// Indicates which Supporter package the fetching user has, or empty
	// if not a Supporter. Only sent for actionType=view.
	UserSupporterTier SupporterTier `xml:"UserSupporterTier"`

	// The requested team's training data. Only sent for actionType=view
	// (and, nested under View, for actionType=setTraining).
	Team *TrainingTeam `xml:"Team"`

	// The league's training-type distribution statistics. Only sent for
	// actionType=stats.
	League *TrainingLeagueStats `xml:"League"`

	// True if setting the training was successful. Only sent for
	// actionType=setTraining.
	TrainingSet bool `xml:"TrainingSet"`

	// The team's training data after setting it. Only sent for
	// actionType=setTraining.
	View *TrainingTeam `xml:"View"`
}

// TrainingTeam is a team's current training settings and experience.
type TrainingTeam struct {
	TeamID   id.Team `xml:"TeamID"`
	TeamName string  `xml:"TeamName"`

	// The current training level for the team, in percent.
	TrainingLevel uint `xml:"TrainingLevel"`

	// The goal training level for the team in percent. If the team is
	// not about to change training level, Available is false and the
	// value is empty.
	NewTrainingLevel struct {
		Available bool `xml:"Available,attr"`
		Value     uint `xml:",chardata"`
	} `xml:"NewTrainingLevel"`

	TrainingType TrainingType `xml:"TrainingType"`

	// The current part of the training level, in percent, spent on
	// stamina training.
	StaminaTrainingPart uint `xml:"StaminaTrainingPart"`

	LastTrainingTrainingType        TrainingType `xml:"LastTrainingTrainingType"`
	LastTrainingTrainingLevel       uint         `xml:"LastTrainingTrainingLevel"`
	LastTrainingStaminaTrainingPart uint         `xml:"LastTrainingStaminaTrainingPart"`

	Trainer struct {
		ID          uint         `xml:"TrainerID"`
		Name        string       `xml:"TrainerName"`
		ArrivalDate HattrickTime `xml:"ArrivalDate"`
	} `xml:"Trainer"`

	// Container for special training instructions. Empty if none are
	// given. Contains at most 3 Player nodes.
	SpecialTraining struct {
		Players []*struct {
			PlayerID id.Player `xml:"PlayerID"`

			// Always 1 = Stamina.
			SpecialTrainingTypeID uint `xml:"SpecialTrainingTypeID"`
		} `xml:"Player"`

		// NOTE: per the CHPP doc, these are documented as direct
		// children of SpecialTraining (not of a Player) - kept
		// verbatim even though it duplicates Trainer's fields above.
		TrainerName string       `xml:"TrainerName"`
		ArrivalDate HattrickTime `xml:"ArrivalDate"`
	} `xml:"SpecialTraining"`

	// If the team is playing a match, Available is false and the value
	// is empty; otherwise Available is true.
	Morale struct {
		Available bool         `xml:"Available,attr"`
		Value     TeamSpiritID `xml:",chardata"`
	} `xml:"Morale"`

	// If the team is playing a match, Available is false and the value
	// is empty; otherwise Available is true.
	SelfConfidence struct {
		Available bool           `xml:"Available,attr"`
		Value     SelfConfidence `xml:",chardata"`
	} `xml:"SelfConfidence"`

	// The team's tactical experience with each formation.
	Experience442 SkillLevel `xml:"Experience442"`
	Experience433 SkillLevel `xml:"Experience433"`
	Experience451 SkillLevel `xml:"Experience451"`
	Experience352 SkillLevel `xml:"Experience352"`
	Experience532 SkillLevel `xml:"Experience532"`
	Experience343 SkillLevel `xml:"Experience343"`
	Experience541 SkillLevel `xml:"Experience541"`
	Experience523 SkillLevel `xml:"Experience523"`
	Experience550 SkillLevel `xml:"Experience550"`
	Experience253 SkillLevel `xml:"Experience253"`
}

// TrainingLeagueStats is the training-type distribution statistics for a
// league (or globally, if no league was specified).
type TrainingLeagueStats struct {
	LeagueID id.League `xml:"LeagueID"`

	// Empty for global (all-leagues) stats.
	LeagueName string `xml:"LeagueName"`

	// NOTE: the CHPP doc's own JSON/markdown dumps are ambiguous here -
	// they list TrainingStat/NumberOfTeams/FractionOfTeams as if all
	// three were direct children of TrainingStatList, but also describe
	// TrainingStat itself as a "container for a particular training
	// stat". This models TrainingStat as the repeating item, wrapping
	// NumberOfTeams/FractionOfTeams as its own children with the
	// trainingType as its chardata - this has not been verified against
	// a live response.
	Stats []*TrainingStat `xml:"TrainingStatList>TrainingStat"`
}

// TrainingStat is the distribution of teams using a particular training
// type within a league.
type TrainingStat struct {
	Type            TrainingType `xml:",chardata"`
	NumberOfTeams   uint         `xml:"NumberOfTeams"`
	FractionOfTeams uint         `xml:"FractionOfTeams"`
}
