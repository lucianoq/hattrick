package parsed

import (
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/parsed/raw"
)

// Parsed ...
type Parsed interface {
	GetAchievementsXML(values map[string]string) (*chpp.AchievementsXML, error)
	GetArenaDetailsXML(values map[string]string) (*chpp.ArenaDetailsXML, error)
	GetAlliancesXML(values map[string]string) (*chpp.AlliancesXML, error)
	GetAllianceDetailsXML(values map[string]string) (*chpp.AllianceDetailsXML, error)
	GetAvatarsXML(values map[string]string) (*chpp.AvatarsXML, error)
	GetBookmarksXML(values map[string]string) (*chpp.BookmarksXML, error)
	GetChallengesXML(values map[string]string) (*chpp.ChallengesXML, error)
	GetClubXML(values map[string]string) (*chpp.ClubXML, error)
	GetCupMatchesXML(values map[string]string) (*chpp.CupMatchesXML, error)
	GetCurrentBidsXML(values map[string]string) (*chpp.CurrentBidsXML, error)
	GetEconomyXML(values map[string]string) (*chpp.EconomyXML, error)
	GetFansXML(values map[string]string) (*chpp.FansXML, error)
	GetHOFPlayersXML(values map[string]string) (*chpp.HOFPlayersXML, error)
	GetLadderDetailsXML(values map[string]string) (*chpp.LadderDetailsXML, error)
	GetLadderListXML(values map[string]string) (*chpp.LadderListXML, error)
	GetLeagueDetailsXML(values map[string]string) (*chpp.LeagueDetailsXML, error)
	GetLeagueFixturesXML(values map[string]string) (*chpp.LeagueFixturesXML, error)
	GetLiveXML(values map[string]string) (*chpp.LiveXML, error)
	GetManagerCompendiumXML(values map[string]string) (*chpp.ManagerCompendiumXML, error)
	GetMatchesXML(values map[string]string) (*chpp.MatchesXML, error)
	GetMatchesArchiveXML(values map[string]string) (*chpp.MatchesArchiveXML, error)
	GetMatchDetailsXML(values map[string]string) (*chpp.MatchDetailsXML, error)
	GetMatchLineupXML(values map[string]string) (*chpp.MatchLineupXML, error)
	GetMatchOrdersXML(values map[string]string) (*chpp.MatchOrdersXML, error)
	GetNationalTeamsXML(values map[string]string) (*chpp.NationalTeamsXML, error)
	GetNationalTeamDetailsXML(values map[string]string) (*chpp.NationalTeamDetailsXML, error)
	GetNationalTeamMatchesXML(values map[string]string) (*chpp.NationalTeamMatchesXML, error)
	GetNationalPlayersXML(values map[string]string) (*chpp.NationalPlayersXML, error)
	GetPlayersXML(values map[string]string) (*chpp.PlayersXML, error)
	GetPlayerDetailsXML(values map[string]string) (*chpp.PlayerDetailsXML, error)
	GetPlayerEventsXML(values map[string]string) (*chpp.PlayerEventsXML, error)
	GetRegionDetailsXML(values map[string]string) (*chpp.RegionDetailsXML, error)
	GetSearchXML(values map[string]string) (*chpp.SearchXML, error)
	GetStaffAvatarsXML(values map[string]string) (*chpp.StaffAvatarsXML, error)
	GetStaffListXML(values map[string]string) (*chpp.StaffListXML, error)
	GetSupportersXML(values map[string]string) (*chpp.SupportersXML, error)
	GetTeamDetailsXML(values map[string]string) (*chpp.TeamDetailsXML, error)
	GetTournamentDetailsXML(values map[string]string) (*chpp.TournamentDetailsXML, error)
	GetTournamentFixturesXML(values map[string]string) (*chpp.TournamentFixturesXML, error)
	GetTournamentLeagueTablesXML(values map[string]string) (*chpp.TournamentLeagueTablesXML, error)
	GetTournamentListXML(values map[string]string) (*chpp.TournamentListXML, error)
	GetTrainingXML(values map[string]string) (*chpp.TrainingXML, error)
	GetTrainingEventsXML(values map[string]string) (*chpp.TrainingEventsXML, error)
	GetTransferSearchXML(values map[string]string) (*chpp.TransferSearchXML, error)
	GetTransferSteamXML(values map[string]string) (*chpp.TransferSteamXML, error)
	GetTransfersPlayerXML(values map[string]string) (*chpp.TransfersPlayerXML, error)
	GetTranslationsXML(values map[string]string) (*chpp.TranslationsXML, error)
	GetWorldCupXML(values map[string]string) (*chpp.WorldCupXML, error)
	GetWorldDetailsXML(values map[string]string) (*chpp.WorldDetailsXML, error)
	GetWorldLanguagesXML(values map[string]string) (*chpp.WorldLanguagesXML, error)
	GetYouthAvatarsXML(values map[string]string) (*chpp.YouthAvatarsXML, error)
	GetYouthLeagueDetailsXML(values map[string]string) (*chpp.YouthLeagueDetailsXML, error)
	GetYouthLeagueFixturesXML(values map[string]string) (*chpp.YouthLeagueFixturesXML, error)
	GetYouthPlayerDetailsXML(values map[string]string) (*chpp.YouthPlayerDetailsXML, error)
	GetYouthPlayerListXML(values map[string]string) (*chpp.YouthPlayerListXML, error)
	GetYouthTeamDetailsXML(values map[string]string) (*chpp.YouthTeamDetailsXML, error)
}

type parsed struct {
	codec Codec
	raw   *raw.Raw
}

var _ Parsed = (*parsed)(nil)

// NewParsed ...
func NewParsed(consumerKey, consumerSecret, accessToken, accessSecret string, accessAdditionalData map[string]string) (Parsed, error) {
	rawData, err := raw.NewRaw(consumerKey, consumerSecret, accessToken, accessSecret, accessAdditionalData)
	if err != nil {
		return nil, err
	}
	return &parsed{
		codec: &codec{},
		raw:   rawData,
	}, nil
}

// GetAchievementsXML ...
func (p *parsed) GetAchievementsXML(values map[string]string) (*chpp.AchievementsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.AchievementsAPIFile, chpp.AchievementsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Achievements(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetArenaDetailsXML ...
func (p *parsed) GetArenaDetailsXML(values map[string]string) (*chpp.ArenaDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.ArenaDetailsAPIFile, chpp.ArenaDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.ArenaDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetAlliancesXML ...
func (p *parsed) GetAlliancesXML(values map[string]string) (*chpp.AlliancesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.AlliancesAPIFile, chpp.AlliancesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Alliances(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetAllianceDetailsXML ...
func (p *parsed) GetAllianceDetailsXML(values map[string]string) (*chpp.AllianceDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.AllianceDetailsAPIFile, chpp.AllianceDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.AllianceDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetAvatarsXML ...
func (p *parsed) GetAvatarsXML(values map[string]string) (*chpp.AvatarsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.AvatarsAPIFile, chpp.AvatarsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Avatars(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetBookmarksXML ...
func (p *parsed) GetBookmarksXML(values map[string]string) (*chpp.BookmarksXML, error) {
	buf, err := p.raw.GetRawXML(chpp.BookmarksAPIFile, chpp.BookmarksAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Bookmarks(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetChallengesXML ...
func (p *parsed) GetChallengesXML(values map[string]string) (*chpp.ChallengesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.ChallengesAPIFile, chpp.ChallengesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Challenges(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetClubXML ...
func (p *parsed) GetClubXML(values map[string]string) (*chpp.ClubXML, error) {
	buf, err := p.raw.GetRawXML(chpp.ClubAPIFile, chpp.ClubAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Club(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetCupMatchesXML ...
func (p *parsed) GetCupMatchesXML(values map[string]string) (*chpp.CupMatchesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.CupMatchesAPIFile, chpp.CupMatchesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.CupMatches(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetCurrentBidsXML ...
func (p *parsed) GetCurrentBidsXML(values map[string]string) (*chpp.CurrentBidsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.CurrentBidsAPIFile, chpp.CurrentBidsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.CurrentBids(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetEconomyXML ...
func (p *parsed) GetEconomyXML(values map[string]string) (*chpp.EconomyXML, error) {
	buf, err := p.raw.GetRawXML(chpp.EconomyAPIFile, chpp.EconomyAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Economy(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetFansXML ...
func (p *parsed) GetFansXML(values map[string]string) (*chpp.FansXML, error) {
	buf, err := p.raw.GetRawXML(chpp.FansAPIFile, chpp.FansAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Fans(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetHOFPlayersXML ...
func (p *parsed) GetHOFPlayersXML(values map[string]string) (*chpp.HOFPlayersXML, error) {
	buf, err := p.raw.GetRawXML(chpp.HOFPlayersAPIFile, chpp.HOFPlayersAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.HOFPlayers(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetLadderDetailsXML ...
func (p *parsed) GetLadderDetailsXML(values map[string]string) (*chpp.LadderDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.LadderDetailsAPIFile, chpp.LadderDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.LadderDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetLadderListXML ...
func (p *parsed) GetLadderListXML(values map[string]string) (*chpp.LadderListXML, error) {
	buf, err := p.raw.GetRawXML(chpp.LadderListAPIFile, chpp.LadderListAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.LadderList(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetLeagueDetailsXML ...
func (p *parsed) GetLeagueDetailsXML(values map[string]string) (*chpp.LeagueDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.LeagueDetailsAPIFile, chpp.LeagueDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.LeagueDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetLeagueFixturesXML ...
func (p *parsed) GetLeagueFixturesXML(values map[string]string) (*chpp.LeagueFixturesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.LeagueFixturesAPIFile, chpp.LeagueFixturesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.LeagueFixtures(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetLiveXML ...
func (p *parsed) GetLiveXML(values map[string]string) (*chpp.LiveXML, error) {
	buf, err := p.raw.GetRawXML(chpp.LiveAPIFile, chpp.LiveAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Live(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetManagerCompendiumXML ...
func (p *parsed) GetManagerCompendiumXML(values map[string]string) (*chpp.ManagerCompendiumXML, error) {
	buf, err := p.raw.GetRawXML(chpp.ManagerCompendiumAPIFile, chpp.ManagerCompendiumAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.ManagerCompendium(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetMatchesXML ...
func (p *parsed) GetMatchesXML(values map[string]string) (*chpp.MatchesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.MatchesAPIFile, chpp.MatchesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Matches(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetMatchesArchiveXML ...
func (p *parsed) GetMatchesArchiveXML(values map[string]string) (*chpp.MatchesArchiveXML, error) {
	buf, err := p.raw.GetRawXML(chpp.MatchesArchiveAPIFile, chpp.MatchesArchiveAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.MatchesArchive(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetMatchDetailsXML ...
func (p *parsed) GetMatchDetailsXML(values map[string]string) (*chpp.MatchDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.MatchDetailsAPIFile, chpp.MatchDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.MatchDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetMatchLineupXML ...
func (p *parsed) GetMatchLineupXML(values map[string]string) (*chpp.MatchLineupXML, error) {
	buf, err := p.raw.GetRawXML(chpp.MatchLineupAPIFile, chpp.MatchLineupAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.MatchLineup(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetMatchOrdersXML ...
func (p *parsed) GetMatchOrdersXML(values map[string]string) (*chpp.MatchOrdersXML, error) {
	buf, err := p.raw.GetRawXML(chpp.MatchOrdersAPIFile, chpp.MatchOrdersAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.MatchOrders(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetNationalTeamsXML ...
func (p *parsed) GetNationalTeamsXML(values map[string]string) (*chpp.NationalTeamsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.NationalTeamsAPIFile, chpp.NationalTeamsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.NationalTeams(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetNationalTeamDetailsXML ...
func (p *parsed) GetNationalTeamDetailsXML(values map[string]string) (*chpp.NationalTeamDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.NationalTeamDetailsAPIFile, chpp.NationalTeamDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.NationalTeamDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetNationalTeamMatchesXML ...
func (p *parsed) GetNationalTeamMatchesXML(values map[string]string) (*chpp.NationalTeamMatchesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.NationalTeamMatchesAPIFile, chpp.NationalTeamMatchesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.NationalTeamMatches(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetNationalPlayersXML ...
func (p *parsed) GetNationalPlayersXML(values map[string]string) (*chpp.NationalPlayersXML, error) {
	buf, err := p.raw.GetRawXML(chpp.NationalPlayersAPIFile, chpp.NationalPlayersAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.NationalPlayers(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetPlayersXML ...
func (p *parsed) GetPlayersXML(values map[string]string) (*chpp.PlayersXML, error) {
	buf, err := p.raw.GetRawXML(chpp.PlayersAPIFile, chpp.PlayersAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Players(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetPlayerDetailsXML ...
func (p *parsed) GetPlayerDetailsXML(values map[string]string) (*chpp.PlayerDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.PlayerDetailsAPIFile, chpp.PlayerDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.PlayerDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetPlayerEventsXML ...
func (p *parsed) GetPlayerEventsXML(values map[string]string) (*chpp.PlayerEventsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.PlayerEventsAPIFile, chpp.PlayerEventsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.PlayerEvents(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetRegionDetailsXML ...
func (p *parsed) GetRegionDetailsXML(values map[string]string) (*chpp.RegionDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.RegionDetailsAPIFile, chpp.RegionDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.RegionDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetSearchXML ...
func (p *parsed) GetSearchXML(values map[string]string) (*chpp.SearchXML, error) {
	buf, err := p.raw.GetRawXML(chpp.SearchAPIFile, chpp.SearchAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Search(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetStaffAvatarsXML ...
func (p *parsed) GetStaffAvatarsXML(values map[string]string) (*chpp.StaffAvatarsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.StaffAvatarsAPIFile, chpp.StaffAvatarsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.StaffAvatars(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetStaffListXML ...
func (p *parsed) GetStaffListXML(values map[string]string) (*chpp.StaffListXML, error) {
	buf, err := p.raw.GetRawXML(chpp.StaffListAPIFile, chpp.StaffListAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.StaffList(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetSupportersXML ...
func (p *parsed) GetSupportersXML(values map[string]string) (*chpp.SupportersXML, error) {
	buf, err := p.raw.GetRawXML(chpp.SupportersAPIFile, chpp.SupportersAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Supporters(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTeamDetailsXML ...
func (p *parsed) GetTeamDetailsXML(values map[string]string) (*chpp.TeamDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TeamDetailsAPIFile, chpp.TeamDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TeamDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTournamentDetailsXML ...
func (p *parsed) GetTournamentDetailsXML(values map[string]string) (*chpp.TournamentDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TournamentDetailsAPIFile, chpp.TournamentDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TournamentDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTournamentFixturesXML ...
func (p *parsed) GetTournamentFixturesXML(values map[string]string) (*chpp.TournamentFixturesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TournamentFixturesAPIFile, chpp.TournamentFixturesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TournamentFixtures(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTournamentLeagueTablesXML ...
func (p *parsed) GetTournamentLeagueTablesXML(values map[string]string) (*chpp.TournamentLeagueTablesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TournamentLeagueTablesAPIFile, chpp.TournamentLeagueTablesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TournamentLeagueTables(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTournamentListXML ...
func (p *parsed) GetTournamentListXML(values map[string]string) (*chpp.TournamentListXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TournamentListAPIFile, chpp.TournamentListAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TournamentList(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTrainingXML ...
func (p *parsed) GetTrainingXML(values map[string]string) (*chpp.TrainingXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TrainingAPIFile, chpp.TrainingAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Training(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTrainingEventsXML ...
func (p *parsed) GetTrainingEventsXML(values map[string]string) (*chpp.TrainingEventsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TrainingEventsAPIFile, chpp.TrainingEventsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TrainingEvents(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTransferSearchXML ...
func (p *parsed) GetTransferSearchXML(values map[string]string) (*chpp.TransferSearchXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TransferSearchAPIFile, chpp.TransferSearchAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TransferSearch(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTransferSteamXML ...
func (p *parsed) GetTransferSteamXML(values map[string]string) (*chpp.TransferSteamXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TransferSteamAPIFile, chpp.TransferSteamAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TransferSteam(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTransfersPlayerXML ...
func (p *parsed) GetTransfersPlayerXML(values map[string]string) (*chpp.TransfersPlayerXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TransfersPlayerAPIFile, chpp.TransfersPlayerAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.TransfersPlayer(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetTranslationsXML ...
func (p *parsed) GetTranslationsXML(values map[string]string) (*chpp.TranslationsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.TranslationsAPIFile, chpp.TranslationsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.Translations(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetWorldCupXML ...
func (p *parsed) GetWorldCupXML(values map[string]string) (*chpp.WorldCupXML, error) {
	buf, err := p.raw.GetRawXML(chpp.WorldCupAPIFile, chpp.WorldCupAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.WorldCup(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetWorldDetailsXML ...
func (p *parsed) GetWorldDetailsXML(values map[string]string) (*chpp.WorldDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.WorldDetailsAPIFile, chpp.WorldDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.WorldDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetWorldLanguagesXML ...
func (p *parsed) GetWorldLanguagesXML(values map[string]string) (*chpp.WorldLanguagesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.WorldLanguagesAPIFile, chpp.WorldLanguagesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.WorldLanguages(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthAvatarsXML ...
func (p *parsed) GetYouthAvatarsXML(values map[string]string) (*chpp.YouthAvatarsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthAvatarsAPIFile, chpp.YouthAvatarsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthAvatars(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthLeagueDetailsXML ...
func (p *parsed) GetYouthLeagueDetailsXML(values map[string]string) (*chpp.YouthLeagueDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthLeagueDetailsAPIFile, chpp.YouthLeagueDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthLeagueDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthLeagueFixturesXML ...
func (p *parsed) GetYouthLeagueFixturesXML(values map[string]string) (*chpp.YouthLeagueFixturesXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthLeagueFixturesAPIFile, chpp.YouthLeagueFixturesAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthLeagueFixtures(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthPlayerDetailsXML ...
func (p *parsed) GetYouthPlayerDetailsXML(values map[string]string) (*chpp.YouthPlayerDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthPlayerDetailsAPIFile, chpp.YouthPlayerDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthPlayerDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthPlayerListXML ...
func (p *parsed) GetYouthPlayerListXML(values map[string]string) (*chpp.YouthPlayerListXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthPlayerListAPIFile, chpp.YouthPlayerListAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthPlayerList(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetYouthTeamDetailsXML ...
func (p *parsed) GetYouthTeamDetailsXML(values map[string]string) (*chpp.YouthTeamDetailsXML, error) {
	buf, err := p.raw.GetRawXML(chpp.YouthTeamDetailsAPIFile, chpp.YouthTeamDetailsAPIVersion, values)
	if err != nil {
		return nil, err
	}

	ret, err := p.codec.YouthTeamDetails(buf)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
