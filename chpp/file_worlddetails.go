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

	LeagueList struct {
		Leagues []*Country `xml:"League"`
	} `xml:"LeagueList"`
}

type Country struct {
	LeagueID     id.League `xml:"LeagueID"`
	LeagueName   string    `xml:"LeagueName"`
	Season       uint      `xml:"Season"`
	MatchRound   uint      `xml:"MatchRound"`
	ShortName    string    `xml:"ShortName"`
	Continent    string    `xml:"Continent"`
	ZoneName     string    `xml:"ZoneName"`
	EnglishName  string    `xml:"EnglishName"`
	LanguageId   uint      `xml:"LanguageId"`
	LanguageName string    `xml:"LanguageName"`
	Country      struct {
		Available           string     `xml:"Available,attr"`
		CountryID           id.Country `xml:"CountryID"`
		CountryName         string     `xml:"CountryName"`
		CurrencyName        string     `xml:"CurrencyName"`
		CurrencyRate        string     `xml:"CurrencyRate"`
		CurrencyRateFloat64 float64
	} `xml:"Country"`
	Cup struct {
		CupID   id.Cup `xml:"CupID"`
		CupName string `xml:"CupName"`
	} `xml:"Cup"`
	ActiveUsers     uint         `xml:"ActiveUsers"`
	WaitingUsers    uint         `xml:"WaitingUsers"`
	TrainingDate    HattrickTime `xml:"TrainingDate"`
	EconomyDate     HattrickTime `xml:"EconomyDate"`
	CupMatchDate    HattrickTime `xml:"CupMatchDate"`
	SeriesMatchDate HattrickTime `xml:"SeriesMatchDate"`
	NumberOfLevels  uint         `xml:"NumberOfLevels"`
}
