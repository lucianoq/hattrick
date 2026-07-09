package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	EconomyAPIFile    = "economy"
	EconomyAPIVersion = "1.4"
)

// EconomyXML contains a team's cash, budgeted income/costs for this week,
// and the actual income/costs from last week.
type EconomyXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Team Economy `xml:"Team"`
}

// Economy holds a team's finances: current cash plus the budgeted
// income/costs for this week and the actual income/costs from last week.
type Economy struct {
	TeamID   id.Team `xml:"TeamID"`
	TeamName string  `xml:"TeamName"`
	Cash     Money   `xml:"Cash"`

	// The budgeted cash for next week.
	ExpectedCash         Money             `xml:"ExpectedCash"`
	SponsorsPopularity   EconomyPopularity `xml:"SponsorsPopularity"`
	SupportersPopularity EconomyPopularity `xml:"SupportersPopularity"`
	FanClubSize          uint              `xml:"FanClubSize"`
	IncomeSpectators     Money             `xml:"IncomeSpectators"`
	IncomeSponsors       Money             `xml:"IncomeSponsors"`

	// The budgeted income for completing sponsor tasks.
	IncomeSponsorBonuses        Money `xml:"IncomeSponsorBonuses"`
	IncomeFinancial             Money `xml:"IncomeFinancial"`
	IncomeSoldPlayers           Money `xml:"IncomeSoldPlayers"`
	IncomeSoldPlayersCommission Money `xml:"IncomeSoldPlayersCommission"`
	IncomeTemporary             Money `xml:"IncomeTemporary"`

	// The budgeted total income for this week.
	IncomeSum          Money `xml:"IncomeSum"`
	CostsArena         Money `xml:"CostsArena"`
	CostsPlayer        Money `xml:"CostsPlayer"`
	CostsFinancial     Money `xml:"CostsFinancial"`
	CostsBoughtPlayers Money `xml:"CostsBoughtPlayers"`
	CostsArenaBuilding Money `xml:"CostsArenaBuilding"`
	CostsTemporary     Money `xml:"CostsTemporary"`
	CostsStaff         Money `xml:"CostsStaff"`
	CostsYouth         Money `xml:"CostsYouth"`

	// The budgeted total cost for this week.
	CostsSum Money `xml:"CostsSum"`

	// The budgeted economic result (income minus costs) for this week.
	ExpectedWeeksTotal              Money `xml:"ExpectedWeeksTotal"`
	LastIncomeSpectators            Money `xml:"LastIncomeSpectators"`
	LastIncomeSponsors              Money `xml:"LastIncomeSponsors"`
	LastIncomeFinancial             Money `xml:"LastIncomeFinancial"`
	LastIncomeSoldPlayers           Money `xml:"LastIncomeSoldPlayers"`
	LastIncomeSoldPlayersCommission Money `xml:"LastIncomeSoldPlayersCommission"`
	LastIncomeTemporary             Money `xml:"LastIncomeTemporary"`

	// The total income for last week.
	LastIncomeSum          Money `xml:"LastIncomeSum"`
	LastCostsArena         Money `xml:"LastCostsArena"`
	LastCostsPlayer        Money `xml:"LastCostsPlayer"`
	LastCostsFinancial     Money `xml:"LastCostsFinancial"`
	LastCostsBoughtPlayers Money `xml:"LastCostsBoughtPlayers"`
	LastCostsArenaBuilding Money `xml:"LastCostsArenaBuilding"`
	LastCostsTemporary     Money `xml:"LastCostsTemporary"`
	LastCostsStaff         Money `xml:"LastCostsStaff"`
	LastCostsYouth         Money `xml:"LastCostsYouth"`

	// The total cost for last week.
	LastCostsSum Money `xml:"LastCostsSum"`

	// The economic result (income minus costs) for last week.
	LastWeeksTotal Money `xml:"LastWeeksTotal"`
}

// EconomyPopularity is a popularity value (among sponsors or supporters)
// which may be unavailable while a match is running.
type EconomyPopularity struct {
	Value     uint `xml:",chardata"`
	Available bool `xml:"Available,attr"`
}
