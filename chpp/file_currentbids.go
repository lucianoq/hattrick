package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	CurrentBidsAPIFile    = "currentbids"
	CurrentBidsAPIVersion = "1.0"
)

// CurrentBidsXML contains the active, recently finished, and hotlisted
// player transfers being tracked for a team.
type CurrentBidsXML struct {
	Envelope
	UserID id.User `xml:"UserID"`

	TeamID id.Team `xml:"TeamId"`

	// Transfers grouped by tracking category (e.g. selling, buying,
	// hotlisted, finished).
	BidItems []*struct {
		TrackingTypeID TrackingTypeID `xml:"TrackingTypeId,attr"`
		BidItem        []*BidItem     `xml:"BidItem"`
	} `xml:"BidItems"`
}

// BidItem is a single tracked player transfer, along with its current
// highest bid.
type BidItem struct {
	// Unique ID for the transfer. Only needed for the ignore functionality.
	TransferID     id.Transfer `xml:"TransferId"`
	TrackingTypeID TrackingTypeID
	PlayerID       id.Player `xml:"PlayerId"`
	PlayerName     string    `xml:"PlayerName"`

	// The current highest bid for the player. Empty if there are no bids.
	HighestBid struct {
		Amount   Money   `xml:"Amount"`
		TeamID   id.Team `xml:"TeamId"`
		TeamName string  `xml:"TeamName"`
	} `xml:"HighestBid"`

	Deadline HattrickTime `xml:"Deadline"`
}
