package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// GetCurrentBids shows all active, recently finished or hotlisted transfers for
// the specified team.
func (a *API) GetCurrentBids() ([]*chpp.BidItem, error) {
	values := map[string]string{
		"actionType": "view",
	}

	list, err := a.parsed.GetCurrentBidsXML(values)
	if err != nil {
		return nil, err
	}

	// degrouping bids embedding type in each item
	allBids := make([]*chpp.BidItem, 0)
	for _, bids := range list.BidItems {
		for _, subBid := range bids.BidItem {
			subBid.TrackingTypeID = bids.TrackingTypeID
			allBids = append(allBids, subBid)
		}
	}

	return allBids, nil
}

// DeleteAllFinishedBids removes the tracking of all finished transfers in one
// go, for Supporters only. (Requires "place_bid" scope.)
func (a *API) DeleteAllFinishedBids() error {
	values := map[string]string{
		"actionType": "deleteAllFinished",
	}

	_, err := a.parsed.GetCurrentBidsXML(values)
	return err
}

// IgnoreTransfer removes the tracking of a specific transfer, for Supporters
// only. Only possible for losing bids, hotlisted and finished.
// (Requires "place_bid" scope.)
//
//    transferId : unsigned Integer
//    Id of the transfer to be ignored.
//
//    trackingTypeId : trackingTypeId
//    Tracking type (category) of the transfer to be ignored. Only values 5, 8
//    and 9 are allowed.
func (a *API) IgnoreTransfer(transferID id.Transfer, category chpp.TrackingTypeID) error {
	values := map[string]string{
		"actionType":     "view",
		"transferId":     transferID.String(),
		"trackingTypeId": category.String(),
	}

	_, err := a.parsed.GetCurrentBidsXML(values)
	return err
}
