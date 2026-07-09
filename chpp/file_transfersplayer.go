package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TransfersPlayerAPIFile    = "transfersplayer"
	TransfersPlayerAPIVersion = "1.1"
)

// TransfersPlayerXML contains the transfer history for a single player.
type TransfersPlayerXML struct {
	Envelope
	UserID id.User `xml:"User"`

	Transfers PlayerTransferHistory `xml:"Transfers"`
}

// PlayerTransferHistory is the container for the given player's transfer
// history.
type PlayerTransferHistory struct {
	StartDate HattrickTime `xml:"StartDate"`
	EndDate   HattrickTime `xml:"EndDate"`

	Player struct {
		ID   id.Player `xml:"PlayerID"`
		Name string    `xml:"PlayerName"`
	} `xml:"Player"`

	Transfers []*PlayerTransfer `xml:"Transfer"`
}

// PlayerTransfer is a single historical transfer of a player.
type PlayerTransfer struct {
	ID       id.Transfer  `xml:"TransferID"`
	Deadline HattrickTime `xml:"Deadline"`

	Buyer struct {
		TeamID   id.Team `xml:"BuyerTeamID"`
		TeamName string  `xml:"BuyerTeamName"`
	} `xml:"Buyer"`

	Seller struct {
		TeamID   id.Team `xml:"SellerTeamID"`
		TeamName string  `xml:"SellerTeamName"`
	} `xml:"Seller"`

	Price Money `xml:"Price"`
	TSI   uint  `xml:"TSI"`
}
