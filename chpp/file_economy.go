package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	EconomyAPIFile    = "economy"
	EconomyAPIVersion = "1.4"
)

// EconomyXML ...
type EconomyXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"User"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	Team Economy `xml:"Team"`
}

// Economy ...
type Economy struct {
	TeamID                          id.Team                `xml:"TeamID"`
	TeamName                        string                 `xml:"TeamName"`
	Cash                            Money                  `xml:"Cash"`
	ExpectedCash                    Money                  `xml:"ExpectedCash"`
	SponsorsPopularity              SupportersPopularityID `xml:"SponsorsPopularity"`
	SupportersPopularity            FanMood                `xml:"SupportersPopularity"`
	FanClubSize                     uint                   `xml:"FanClubSize"`
	IncomeSpectators                Money                  `xml:"IncomeSpectators"`
	IncomeSponsors                  Money                  `xml:"IncomeSponsors"`
	IncomeSponsorsBonuses           Money                  `xml:"IncomeSponsorsBonuses"`
	IncomeFinancial                 Money                  `xml:"IncomeFinancial"`
	IncomeSoldPlayers               Money                  `xml:"IncomeSoldPlayers"`
	IncomeSoldPlayersCommission     Money                  `xml:"IncomeSoldPlayersCommission"`
	IncomeSum                       Money                  `xml:"IncomeSum"`
	CostsArena                      Money                  `xml:"CostsArena"`
	CostsPlayers                    Money                  `xml:"CostsPlayers"`
	CostsFinancial                  Money                  `xml:"CostsFinancial"`
	CostsBoughtPlayers              Money                  `xml:"CostsBoughtPlayers"`
	CostsArenaBuilding              Money                  `xml:"CostsArenaBuilding"`
	CostsStaff                      Money                  `xml:"CostsStaff"`
	CostsYouth                      Money                  `xml:"CostsYouth"`
	CostsSum                        Money                  `xml:"CostsSum"`
	ExpectedWeeksTotal              Money                  `xml:"ExpectedWeeksTotal"`
	LastIncomeSpectators            Money                  `xml:"LastIncomeSpectators"`
	LastIncomeSponsors              Money                  `xml:"LastIncomeSponsors"`
	LastIncomeFinancial             Money                  `xml:"LastIncomeFinancial"`
	LastIncomeSoldPlayers           Money                  `xml:"LastIncomeSoldPlayers"`
	LastIncomeSoldPlayersCommission Money                  `xml:"LastIncomeSoldPlayersCommission"`
	LastIncomeSum                   Money                  `xml:"LastIncomeSum"`
	LastCostsArena                  Money                  `xml:"LastCostsArena"`
	LastCostsPlayers                Money                  `xml:"LastCostsPlayers"`
	LastCostsFinancial              Money                  `xml:"LastCostsFinancial"`
	LastCostsBoughtPlayers          Money                  `xml:"LastCostsBoughtPlayers"`
	LastCostsArenaBuilding          Money                  `xml:"LastCostsArenaBuilding"`
	LastCostsStaff                  Money                  `xml:"LastCostsStaff"`
	LastCostsYouth                  Money                  `xml:"LastCostsYouth"`
	LastCostsSum                    Money                  `xml:"LastCostsSum"`
	LastWeeksTotal                  Money                  `xml:"LastWeeksTotal"`
}
