package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TransferSteamAPIFile    = "transfersteam"
	TransferSteamAPIVersion = "1.2"
)

// TransferSteamXML contains a team's transfer history (a page of buys and
// sales) plus aggregate buy/sell statistics.
type TransferSteamXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Team struct {
		ID   id.Team `xml:"TeamID"`
		Name string  `xml:"TeamName"`

		// The date when the user got this team.
		ActivatedDate HattrickTime `xml:"ActivatedDate"`
	} `xml:"Team"`

	// TotalSumOfBuys and TotalSumOfSales are always expressed in SEK,
	// regardless of the requesting team's local currency.
	Stats struct {
		TotalSumOfBuys  Money `xml:"TotalSumOfBuys"`
		TotalSumOfSales Money `xml:"TotalSumOfSales"`
		NumberOfBuys    uint  `xml:"NumberOfBuys"`
		NumberOfSales   uint  `xml:"NumberOfSales"`
	} `xml:"Stats"`

	Transfers TeamTransferHistory `xml:"Transfers"`
}

// TeamTransferHistory is the container for a page of a team's transfer
// history.
type TeamTransferHistory struct {
	// The page returned. Note that requesting pageIndex=0 returns the
	// last page, not literally page 0 - PageIndex reflects the actual
	// page returned.
	PageIndex uint `xml:"PageIndex"`

	// The total number of pages available.
	Pages uint `xml:"Pages"`

	// The oldest date of the selected transfers list.
	StartDate HattrickTime `xml:"StartDate"`

	// The latest date of the selected transfers list.
	EndDate HattrickTime `xml:"EndDate"`

	Transfers []*TeamTransfer `xml:"Transfer"`
}

// TeamTransfer is a single historical transfer involving a team.
type TeamTransfer struct {
	ID id.Transfer `xml:"TransferID"`

	// The date when bidding closed for the transfer.
	Deadline HattrickTime `xml:"Deadline"`

	Player struct {
		ID   id.Player `xml:"PlayerID"`
		Name string    `xml:"PlayerName"`

		// The player's TSI at the time he was sent to the transfer list.
		TSI uint `xml:"TSI"`

		// Either "S" or "B", representing whether the transfer was a
		// sale or a buy, seen from the requested team's point of view.
		TransferType string `xml:"TransferType"`
	} `xml:"Player"`

	Buyer struct {
		TeamID   id.Team `xml:"BuyerTeamID"`
		TeamName string  `xml:"BuyerTeamName"`
	} `xml:"Buyer"`

	Seller struct {
		TeamID   id.Team `xml:"SellerTeamID"`
		TeamName string  `xml:"SellerTeamName"`
	} `xml:"Seller"`

	Price Money `xml:"Price"`
}
