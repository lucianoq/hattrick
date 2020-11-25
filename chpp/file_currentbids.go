package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	CurrentBidsAPIFile    = "currentbids"
	CurrentBidsAPIVersion = "1.0"
)

// CurrentBidsXML ...
type CurrentBidsXML struct {
	FileName    string       `xml:"FileName"`
	Version     string       `xml:"Version"`
	UserID      id.User      `xml:"UserID"`
	FetchedDate HattrickTime `xml:"FetchedDate"`

	Error     string    `xml:"Error"`
	ErrorCode ErrorCode `xml:"ErrorCode"`
	ErrorGUID string    `xml:"ErrorGUID"`
	Server    string    `xml:"Server"`
	Request   string    `xml:"Request"`

	TeamID id.Team `xml:"TeamId"`

	BidItems []*struct {
		TrackingTypeID TrackingTypeID `xml:"TrackingTypeID,attr"`
		BidItem        []*BidItem     `xml:"BidItem"`
	} `xml:"BidItems"`
}

// BidItem ...
type BidItem struct {
	TransferID     id.Transfer `xml:"TransferId"`
	TrackingTypeID TrackingTypeID
	PlayerID       id.Player `xml:"PlayerId"`
	PlayerName     string    `xml:"PlayerName"`

	HighestBid struct {
		Amount   Money   `xml:"Amount"`
		TeamID   id.Team `xml:"TeamId"`
		TeamName string  `xml:"TeamName"`
	} `xml:"HighestBid"`

	Deadline HattrickTime `xml:"Deadline"`
}
