package parsed

import (
	"encoding/xml"
	"errors"

	"github.com/lucianoq/hattrick/chpp"
)

type codec struct{}

var _ Codec = (*codec)(nil)

// Achievements ...
func (c *codec) Achievements(buf []byte) (*chpp.AchievementsXML, error) {
	var x *chpp.AchievementsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.AchievementsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right achievements type")
	}

	return x, nil
}

// ArenaDetails ...
func (c *codec) ArenaDetails(buf []byte) (*chpp.ArenaDetailsXML, error) {
	var x *chpp.ArenaDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.ArenaDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right arenadetails type")
	}

	return x, nil
}

// Alliances ...
func (c *codec) Alliances(buf []byte) (*chpp.AlliancesXML, error) {
	var x *chpp.AlliancesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.AlliancesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right alliances type")
	}

	return x, nil
}

// AllianceDetails ...
func (c *codec) AllianceDetails(buf []byte) (*chpp.AllianceDetailsXML, error) {
	var x *chpp.AllianceDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.AllianceDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right alliancedetails type")
	}

	return x, nil
}

// Avatars ...
func (c *codec) Avatars(buf []byte) (*chpp.AvatarsXML, error) {
	var x *chpp.AvatarsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.AvatarsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right avatars type")
	}

	return x, nil
}

// Bookmarks ...
func (c *codec) Bookmarks(buf []byte) (*chpp.BookmarksXML, error) {
	var x *chpp.BookmarksXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.BookmarksAPIFile+".xml" {
		return nil, errors.New("failed to parse the right bookmarks type")
	}

	return x, nil
}

// Challenges ...
func (c *codec) Challenges(buf []byte) (*chpp.ChallengesXML, error) {
	var x *chpp.ChallengesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.ChallengesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right challenges type")
	}

	return x, nil
}

// Club ...
func (c *codec) Club(buf []byte) (*chpp.ClubXML, error) {
	var x *chpp.ClubXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.ClubAPIFile+".xml" {
		return nil, errors.New("failed to parse the right club type")
	}

	return x, nil
}

// CupMatches ...
func (c *codec) CupMatches(buf []byte) (*chpp.CupMatchesXML, error) {
	var x *chpp.CupMatchesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.CupMatchesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right cupmatches type")
	}

	return x, nil
}

// CurrentBids ...
func (c *codec) CurrentBids(buf []byte) (*chpp.CurrentBidsXML, error) {
	var x *chpp.CurrentBidsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.CurrentBidsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right currentbids type")
	}

	return x, nil
}

// Economy ...
func (c *codec) Economy(buf []byte) (*chpp.EconomyXML, error) {
	var x *chpp.EconomyXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.EconomyAPIFile+".xml" {
		return nil, errors.New("failed to parse the right economy type")
	}

	return x, nil
}

// Fans ...
func (c *codec) Fans(buf []byte) (*chpp.FansXML, error) {
	var x *chpp.FansXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.FansAPIFile+".xml" {
		return nil, errors.New("failed to parse the right fans type")
	}

	return x, nil
}

// HOFPlayers ...
func (c *codec) HOFPlayers(buf []byte) (*chpp.HOFPlayersXML, error) {
	var x *chpp.HOFPlayersXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.HOFPlayersAPIFile+".xml" {
		return nil, errors.New("failed to parse the right hofplayers type")
	}

	return x, nil
}

// LadderDetails ...
func (c *codec) LadderDetails(buf []byte) (*chpp.LadderDetailsXML, error) {
	var x *chpp.LadderDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.LadderDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right ladderdetails type")
	}

	return x, nil
}

// LadderList ...
func (c *codec) LadderList(buf []byte) (*chpp.LadderListXML, error) {
	var x *chpp.LadderListXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.LadderListAPIFile+".xml" {
		return nil, errors.New("failed to parse the right ladderlist type")
	}

	return x, nil
}

// LeagueDetails ...
func (c *codec) LeagueDetails(buf []byte) (*chpp.LeagueDetailsXML, error) {
	var x *chpp.LeagueDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.LeagueDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right leaguedetails type")
	}

	return x, nil
}

// LeagueFixtures ...
func (c *codec) LeagueFixtures(buf []byte) (*chpp.LeagueFixturesXML, error) {
	var x *chpp.LeagueFixturesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.LeagueFixturesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right leaguefixtures type")
	}

	return x, nil
}

// Live ...
func (c *codec) Live(buf []byte) (*chpp.LiveXML, error) {
	var x *chpp.LiveXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.LiveAPIFile+".xml" {
		return nil, errors.New("failed to parse the right live type")
	}

	return x, nil
}

// ManagerCompendium ...
func (c *codec) ManagerCompendium(buf []byte) (*chpp.ManagerCompendiumXML, error) {
	var x *chpp.ManagerCompendiumXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.ManagerCompendiumAPIFile+".xml" {
		return nil, errors.New("failed to parse the right managercompendium type")
	}

	return x, nil
}

// Matches ...
func (c *codec) Matches(buf []byte) (*chpp.MatchesXML, error) {
	var x *chpp.MatchesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.MatchesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right matches type")
	}

	return x, nil
}

// MatchesArchive ...
func (c *codec) MatchesArchive(buf []byte) (*chpp.MatchesArchiveXML, error) {
	var x *chpp.MatchesArchiveXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.MatchesArchiveAPIFile+".xml" {
		return nil, errors.New("failed to parse the right matchesarchive type")
	}

	// We are using the same `Match` type as `matches` file.
	// In `archiveMatches`, all the matches are finished, so just write down the
	// value instead of leaving it empty.
	for _, m := range x.Team.MatchList.Matches {
		m.Status = chpp.MatchStatusFinished
	}

	return x, nil
}

// MatchDetails ...
func (c *codec) MatchDetails(buf []byte) (*chpp.MatchDetailsXML, error) {
	var x *chpp.MatchDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.MatchDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right matchdetails type")
	}

	return x, nil
}

// MatchLineup ...
func (c *codec) MatchLineup(buf []byte) (*chpp.MatchLineupXML, error) {
	var x *chpp.MatchLineupXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.MatchLineupAPIFile+".xml" {
		return nil, errors.New("failed to parse the right matchlineup type")
	}

	return x, nil
}

// MatchOrders ...
func (c *codec) MatchOrders(buf []byte) (*chpp.MatchOrdersXML, error) {
	var x *chpp.MatchOrdersXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.MatchOrdersAPIFile+".xml" {
		return nil, errors.New("failed to parse the right matchorders type")
	}

	return x, nil
}

// NationalTeams ...
func (c *codec) NationalTeams(buf []byte) (*chpp.NationalTeamsXML, error) {
	var x *chpp.NationalTeamsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.NationalTeamsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right nationalteams type")
	}

	return x, nil
}

// NationalTeamDetails ...
func (c *codec) NationalTeamDetails(buf []byte) (*chpp.NationalTeamDetailsXML, error) {
	var x *chpp.NationalTeamDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.NationalTeamDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right nationalteamdetails type")
	}

	return x, nil
}

// NationalTeamMatches ...
func (c *codec) NationalTeamMatches(buf []byte) (*chpp.NationalTeamMatchesXML, error) {
	var x *chpp.NationalTeamMatchesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.NationalTeamMatchesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right nationalteammatches type")
	}

	return x, nil
}

// NationalPlayers ...
func (c *codec) NationalPlayers(buf []byte) (*chpp.NationalPlayersXML, error) {
	var x *chpp.NationalPlayersXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.NationalPlayersAPIFile+".xml" {
		return nil, errors.New("failed to parse the right nationalplayers type")
	}

	return x, nil
}

// Players ...
func (c *codec) Players(buf []byte) (*chpp.PlayersXML, error) {
	var x *chpp.PlayersXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.PlayersAPIFile+".xml" {
		return nil, errors.New("failed to parse the right players type")
	}

	return x, nil
}

// PlayerDetails ...
func (c *codec) PlayerDetails(buf []byte) (*chpp.PlayerDetailsXML, error) {
	var x *chpp.PlayerDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.PlayerDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right playerdetails type")
	}

	return x, nil
}

// PlayerEvents ...
func (c *codec) PlayerEvents(buf []byte) (*chpp.PlayerEventsXML, error) {
	var x *chpp.PlayerEventsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.PlayerEventsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right playerevents type")
	}

	return x, nil
}

// RegionDetails ...
func (c *codec) RegionDetails(buf []byte) (*chpp.RegionDetailsXML, error) {
	var x *chpp.RegionDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.RegionDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right regiondetails type")
	}

	return x, nil
}

// Search ...
func (c *codec) Search(buf []byte) (*chpp.SearchXML, error) {
	var x *chpp.SearchXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.SearchAPIFile+".xml" {
		return nil, errors.New("failed to parse the right search type")
	}

	return x, nil
}

// StaffAvatars ...
func (c *codec) StaffAvatars(buf []byte) (*chpp.StaffAvatarsXML, error) {
	var x *chpp.StaffAvatarsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.StaffAvatarsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right staffavatars type")
	}

	return x, nil
}

// StaffList ...
func (c *codec) StaffList(buf []byte) (*chpp.StaffListXML, error) {
	var x *chpp.StaffListXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.StaffListAPIFile+".xml" {
		return nil, errors.New("failed to parse the right stafflist type")
	}

	return x, nil
}

// Supporters ...
func (c *codec) Supporters(buf []byte) (*chpp.SupportersXML, error) {
	var x *chpp.SupportersXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.SupportersAPIFile+".xml" {
		return nil, errors.New("failed to parse the right supporters type")
	}

	return x, nil
}

// TeamDetails ...
func (c *codec) TeamDetails(buf []byte) (*chpp.TeamDetailsXML, error) {
	var x *chpp.TeamDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TeamDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right teamdetails type")
	}

	return x, nil
}

// TournamentDetails ...
func (c *codec) TournamentDetails(buf []byte) (*chpp.TournamentDetailsXML, error) {
	var x *chpp.TournamentDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TournamentDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right tournamentdetails type")
	}

	return x, nil
}

// TournamentFixtures ...
func (c *codec) TournamentFixtures(buf []byte) (*chpp.TournamentFixturesXML, error) {
	var x *chpp.TournamentFixturesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TournamentFixturesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right tournamentfixtures type")
	}

	return x, nil
}

// TournamentLeagueTables ...
func (c *codec) TournamentLeagueTables(buf []byte) (*chpp.TournamentLeagueTablesXML, error) {
	var x *chpp.TournamentLeagueTablesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TournamentLeagueTablesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right tournamentleaguetables type")
	}

	return x, nil
}

// TournamentList ...
func (c *codec) TournamentList(buf []byte) (*chpp.TournamentListXML, error) {
	var x *chpp.TournamentListXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TournamentListAPIFile+".xml" {
		return nil, errors.New("failed to parse the right tournamentlist type")
	}

	return x, nil
}

// Training ...
func (c *codec) Training(buf []byte) (*chpp.TrainingXML, error) {
	var x *chpp.TrainingXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TrainingAPIFile+".xml" {
		return nil, errors.New("failed to parse the right training type")
	}

	return x, nil
}

// TrainingEvents ...
func (c *codec) TrainingEvents(buf []byte) (*chpp.TrainingEventsXML, error) {
	var x *chpp.TrainingEventsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TrainingEventsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right trainingevents type")
	}

	return x, nil
}

// TransferSearch ...
func (c *codec) TransferSearch(buf []byte) (*chpp.TransferSearchXML, error) {
	var x *chpp.TransferSearchXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TransferSearchAPIFile+".xml" {
		return nil, errors.New("failed to parse the right transfersearch type")
	}

	return x, nil
}

// TransferSteam ...
func (c *codec) TransferSteam(buf []byte) (*chpp.TransferSteamXML, error) {
	var x *chpp.TransferSteamXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TransferSteamAPIFile+".xml" {
		return nil, errors.New("failed to parse the right transfersteam type")
	}

	return x, nil
}

// TransfersPlayer ...
func (c *codec) TransfersPlayer(buf []byte) (*chpp.TransfersPlayerXML, error) {
	var x *chpp.TransfersPlayerXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TransfersPlayerAPIFile+".xml" {
		return nil, errors.New("failed to parse the right transfersplayer type")
	}

	return x, nil
}

// Translations ...
func (c *codec) Translations(buf []byte) (*chpp.TranslationsXML, error) {
	var x *chpp.TranslationsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.TranslationsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right translations type")
	}

	return x, nil
}

// WorldCup ...
func (c *codec) WorldCup(buf []byte) (*chpp.WorldCupXML, error) {
	var x *chpp.WorldCupXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.WorldCupAPIFile+".xml" {
		return nil, errors.New("failed to parse the right worldcup type")
	}

	return x, nil
}

// WorldDetails ...
func (c *codec) WorldDetails(buf []byte) (*chpp.WorldDetailsXML, error) {
	var x *chpp.WorldDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.WorldDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right worlddetails type")
	}

	return x, nil
}

// WorldLanguages ...
func (c *codec) WorldLanguages(buf []byte) (*chpp.WorldLanguagesXML, error) {
	var x *chpp.WorldLanguagesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.WorldLanguagesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right worldlanguages type")
	}

	return x, nil
}

// YouthAvatars ...
func (c *codec) YouthAvatars(buf []byte) (*chpp.YouthAvatarsXML, error) {
	var x *chpp.YouthAvatarsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthAvatarsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthavatars type")
	}

	return x, nil
}

// YouthLeagueDetails ...
func (c *codec) YouthLeagueDetails(buf []byte) (*chpp.YouthLeagueDetailsXML, error) {
	var x *chpp.YouthLeagueDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthLeagueDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthleaguedetails type")
	}

	return x, nil
}

// YouthLeagueFixtures ...
func (c *codec) YouthLeagueFixtures(buf []byte) (*chpp.YouthLeagueFixturesXML, error) {
	var x *chpp.YouthLeagueFixturesXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthLeagueFixturesAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthleaguefixtures type")
	}

	return x, nil
}

// YouthPlayerDetails ...
func (c *codec) YouthPlayerDetails(buf []byte) (*chpp.YouthPlayerDetailsXML, error) {
	var x *chpp.YouthPlayerDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthPlayerDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthplayerdetails type")
	}

	return x, nil
}

// YouthPlayerList ...
func (c *codec) YouthPlayerList(buf []byte) (*chpp.YouthPlayerListXML, error) {
	var x *chpp.YouthPlayerListXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthPlayerListAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthplayerlist type")
	}

	return x, nil
}

// YouthTeamDetails ...
func (c *codec) YouthTeamDetails(buf []byte) (*chpp.YouthTeamDetailsXML, error) {
	var x *chpp.YouthTeamDetailsXML

	// Unmarshal
	err := xml.Unmarshal(buf, &x)
	if err != nil {
		return nil, err
	}

	// Check
	if x.Error != "" {
		return nil, errors.New(x.Error)
	}

	if x.FileName != chpp.YouthTeamDetailsAPIFile+".xml" {
		return nil, errors.New("failed to parse the right youthteamdetails type")
	}

	return x, nil
}
