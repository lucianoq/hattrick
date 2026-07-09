package api

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

func liveValues(sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) map[string]string {
	values := map[string]string{}
	if sourceSystem != "" {
		values["sourceSystem"] = string(sourceSystem)
	}
	if includeStartingLineup {
		values["includeStartingLineup"] = "true"
	}
	if useLiveEventsAndTexts {
		values["useLiveEventsAndTexts"] = "true"
	}
	return values
}

// GetLiveMatch returns the latest event(s) for the given tracked match,
// with all events since the start of the game.
func (a *API) GetLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error) {
	values := liveValues(sourceSystem, includeStartingLineup, useLiveEventsAndTexts)
	values["matchID"] = matchID.String()

	res, err := a.parsed.GetLiveXML(values)
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}

// GetLiveMatches returns the latest event(s) for all tracked matches.
func (a *API) GetLiveMatches(sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error) {
	res, err := a.parsed.GetLiveXML(liveValues(sourceSystem, includeStartingLineup, useLiveEventsAndTexts))
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}

// GetAllLiveEvents returns all the events for all tracked matches. Meant to
// be used once per session; use GetNewLiveEvents for subsequent polling.
func (a *API) GetAllLiveEvents(includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error) {
	values := liveValues("", includeStartingLineup, useLiveEventsAndTexts)
	values["actionType"] = "viewAll"

	res, err := a.parsed.GetLiveXML(values)
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}

// GetNewLiveEvents returns the matches where new events have occurred since
// the last request. lastShownIndexes should be built with
// chpp.NewLastShownIndexes, from the LastShownEventIndex of each match
// returned by the previous call (to GetAllLiveEvents or GetNewLiveEvents).
// Note that StartingLineup is never provided for this actionType.
func (a *API) GetNewLiveEvents(lastShownIndexes string, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error) {
	values := liveValues("", includeStartingLineup, useLiveEventsAndTexts)
	values["actionType"] = "viewNew"
	if lastShownIndexes != "" {
		values["lastShownIndexes"] = lastShownIndexes
	}

	res, err := a.parsed.GetLiveXML(values)
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}

// ClearLiveMatches removes all tracked matches from the tracking session.
func (a *API) ClearLiveMatches() error {
	values := map[string]string{
		"actionType": "clearAll",
	}
	_, err := a.parsed.GetLiveXML(values)
	return err
}

// AddLiveMatch adds a new match to be tracked.
func (a *API) AddLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error) {
	values := liveValues(sourceSystem, includeStartingLineup, useLiveEventsAndTexts)
	values["actionType"] = "addMatch"
	values["matchID"] = matchID.String()

	res, err := a.parsed.GetLiveXML(values)
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}

// DeleteLiveMatch removes a particular match from the tracking session.
func (a *API) DeleteLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem) ([]*chpp.LiveMatchInfo, error) {
	values := map[string]string{
		"actionType": "deleteMatch",
		"matchID":    matchID.String(),
	}
	if sourceSystem != "" {
		values["sourceSystem"] = string(sourceSystem)
	}

	res, err := a.parsed.GetLiveXML(values)
	if err != nil {
		return nil, err
	}
	return res.Matches, nil
}
