package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	WorldDetailsAPIFile    = "worlddetails"
	WorldDetailsAPIVersion = "1.9"
)

// WorldDetailsXML ...
type WorldDetailsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Leagues []*League `xml:"LeagueList>League"`
}

type League struct {
	ID           id.League `xml:"LeagueID"`
	Name         string    `xml:"LeagueName"`
	Season       uint      `xml:"Season"`
	MatchRound   uint      `xml:"MatchRound"`
	ShortName    string    `xml:"ShortName"`
	Continent    string    `xml:"Continent"`
	ZoneName     string    `xml:"ZoneName"`
	EnglishName  string    `xml:"EnglishName"`
	LanguageId   uint      `xml:"LanguageId"`
	LanguageName string    `xml:"LanguageName"`

	// Container for the country that is related to the league.
	// It may seem confusing what the difference is between a league
	// and a country. Well, there isn't much, except that in some cases
	// we refer to CountryID instead of LeagueID. The easiest way to deal
	// with it is perhaps to just say: "It's the way it is. Just learn to live
	// with it.". Some things are formally tied to the Country instead of
	// the League. We have chosen to maintain that relationship in the
	// XML output, hence the sub-containers of Country.
	Country struct {
		Available    bool       `xml:"Available,attr"`
		ID           id.Country `xml:"CountryID"`
		Name         string     `xml:"CountryName"`
		CurrencyName string     `xml:"CurrencyName"`

		// Decimal value specifying the relative currency rate to SEK (swedish krona).
		CurrencyRate        string `xml:"CurrencyRate"`
		CurrencyRateFloat64 float64

		// The date format for users of this country using ISO_8601
		DateFormat string `xml:"DateFormat"`

		// The time format for users of this country using ISO_8601
		TimeFormat string `xml:"TimeFormat"`
	} `xml:"Country"`

	// Data about the league's cups.
	Cups []struct {
		ID   id.Cup `xml:"CupID"`
		Name string `xml:"CupName"`

		// LeagueLevel for the cup. 0 = National (LeagueLevel 1-6), 7-9 = Divisional.
		CupLeagueLevel id.CupLeagueLevel `xml:"CupLeagueLevel"`

		// Level of the cup. 1 = National/Divisional, 2 = Challenger, 3 = Consolation.
		CupLevel id.CupLevel `xml:"CupLevel"`

		// Index of the cup. Always 1 for National and Consolation cups.
		// For Challenger cup: 1 = Emerald, 2 = Ruby, 3 = Sapphire.
		CupLevelIndex id.CupLevelIndex `xml:"CupLevelIndex"`

		// The last played matchround (or currently played if matches are ongoing).
		// (Note that this is different from TeamDetails!).
		MatchRound uint `xml:"MatchRound"`

		// How many rounds are left of this cup.
		MatchRoundLeft uint `xml:"MatchRoundLeft"`
	} `xml:"Cups>Cup"`

	NationalTeamID id.NationalTeam `xml:"NationalTeamID"`
	U20TeamID      id.NationalTeam `xml:"U20TeamID"`
	ActiveTeams    uint            `xml:"ActiveTeams"`
	ActiveUsers    uint            `xml:"ActiveUsers"`
	WaitingUsers   uint            `xml:"WaitingUsers"`

	// The date and time when the next training update will start in this league.
	TrainingDate HattrickTime `xml:"TrainingDate"`

	// The date and time when the next economy update will start in this league.
	EconomyDate HattrickTime `xml:"EconomyDate"`

	// The date and time when the next cup round is being checked for (that is,
	// there are not necessarily any matches to play, if the cup has been finished).
	CupMatchDate HattrickTime `xml:"CupMatchDate"`

	// The date and time when the next series match round starts. Some leagues may
	// have several series match dates - then this data gives the date when
	// the first batch for the league starts.
	SeriesMatchDate HattrickTime `xml:"SeriesMatchDate"`

	// Time and date for an upcoming daily update.
	Sequence1 HattrickTime `xml:"Sequence1"`
	Sequence2 HattrickTime `xml:"Sequence2"`
	Sequence3 HattrickTime `xml:"Sequence3"`
	Sequence5 HattrickTime `xml:"Sequence5"`
	Sequence7 HattrickTime `xml:"Sequence7"`

	// Integer representing the number of LeagueLevels (divisions) active for this league.
	NumberOfLevels uint `xml:"NumberOfLevels"`
}
