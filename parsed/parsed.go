// Package parsed orchestrates, per CHPP file, the fetch-then-decode
// pipeline: it issues a signed HTTP request to retrieve a CHPP file's raw
// XML bytes, then decodes and validates them into the corresponding typed
// XML struct via chpp.DecodeXML.
package parsed

import (
	"github.com/lucianoq/hattrick/chpp"
)

// Parsed is one method per CHPP file/action; the values map carries that
// call's query-string parameters (e.g. team or league IDs).
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
	GetLeagueLevelsXML(values map[string]string) (*chpp.LeagueLevelsXML, error)
	GetLeagueFixturesXML(values map[string]string) (*chpp.LeagueFixturesXML, error)
	GetLiveXML(values map[string]string) (*chpp.LiveXML, error)
	GetManagerCompendiumXML(values map[string]string) (*chpp.ManagerCompendiumXML, error)
	GetMatchesXML(values map[string]string) (*chpp.MatchesXML, error)
	GetMatchesArchiveXML(values map[string]string) (*chpp.MatchesArchiveXML, error)
	GetMatchDetailsXML(values map[string]string) (*chpp.MatchDetailsXML, error)
	GetMatchLineupXML(values map[string]string) (*chpp.MatchLineupXML, error)
	GetMatchOrdersXML(values map[string]string) (*chpp.MatchOrdersXML, error)
	SetMatchOrdersXML(values map[string]string, lineup string) (*chpp.MatchOrdersXML, error)
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
	raw *rawClient
}

var _ Parsed = (*parsed)(nil)

// NewParsed builds a Parsed backed by a signed HTTP client. It takes the
// application's OAuth1 consumer key/secret and the user's access
// token/secret (plus any additional data returned during authorization).
func NewParsed(consumerKey, consumerSecret, accessToken, accessSecret string, accessAdditionalData map[string]string) (Parsed, error) {
	rawData, err := newRawClient(consumerKey, consumerSecret, accessToken, accessSecret, accessAdditionalData)
	if err != nil {
		return nil, err
	}
	return &parsed{
		raw: rawData,
	}, nil
}

// fetchXML issues a signed GET request for the given CHPP file/version and
// decodes+validates the response into a *T via chpp.DecodeXML.
func fetchXML[T chpp.XMLEnvelope](p *parsed, file, version string, values map[string]string) (*T, error) {
	buf, err := p.raw.getRawXML(file, version, values)
	if err != nil {
		return nil, err
	}
	return chpp.DecodeXML[T](buf, file)
}

// GetAchievementsXML fetches the raw achievements CHPP file over HTTP and decodes it into an AchievementsXML.
func (p *parsed) GetAchievementsXML(values map[string]string) (*chpp.AchievementsXML, error) {
	return fetchXML[chpp.AchievementsXML](p, chpp.AchievementsAPIFile, chpp.AchievementsAPIVersion, values)
}

// GetArenaDetailsXML fetches the raw arenaDetails CHPP file over HTTP and decodes it into an ArenaDetailsXML.
func (p *parsed) GetArenaDetailsXML(values map[string]string) (*chpp.ArenaDetailsXML, error) {
	return fetchXML[chpp.ArenaDetailsXML](p, chpp.ArenaDetailsAPIFile, chpp.ArenaDetailsAPIVersion, values)
}

// GetAlliancesXML fetches the raw alliances CHPP file over HTTP and decodes it into an AlliancesXML.
func (p *parsed) GetAlliancesXML(values map[string]string) (*chpp.AlliancesXML, error) {
	return fetchXML[chpp.AlliancesXML](p, chpp.AlliancesAPIFile, chpp.AlliancesAPIVersion, values)
}

// GetAllianceDetailsXML fetches the raw allianceDetails CHPP file over HTTP and decodes it into an AllianceDetailsXML.
func (p *parsed) GetAllianceDetailsXML(values map[string]string) (*chpp.AllianceDetailsXML, error) {
	return fetchXML[chpp.AllianceDetailsXML](p, chpp.AllianceDetailsAPIFile, chpp.AllianceDetailsAPIVersion, values)
}

// GetAvatarsXML fetches the raw avatars CHPP file over HTTP and decodes it into an AvatarsXML.
func (p *parsed) GetAvatarsXML(values map[string]string) (*chpp.AvatarsXML, error) {
	return fetchXML[chpp.AvatarsXML](p, chpp.AvatarsAPIFile, chpp.AvatarsAPIVersion, values)
}

// GetBookmarksXML fetches the raw bookmarks CHPP file over HTTP and decodes it into a BookmarksXML.
func (p *parsed) GetBookmarksXML(values map[string]string) (*chpp.BookmarksXML, error) {
	return fetchXML[chpp.BookmarksXML](p, chpp.BookmarksAPIFile, chpp.BookmarksAPIVersion, values)
}

// GetChallengesXML fetches the raw challenges CHPP file over HTTP and decodes it into a ChallengesXML.
func (p *parsed) GetChallengesXML(values map[string]string) (*chpp.ChallengesXML, error) {
	return fetchXML[chpp.ChallengesXML](p, chpp.ChallengesAPIFile, chpp.ChallengesAPIVersion, values)
}

// GetClubXML fetches the raw club CHPP file over HTTP and decodes it into a ClubXML.
func (p *parsed) GetClubXML(values map[string]string) (*chpp.ClubXML, error) {
	return fetchXML[chpp.ClubXML](p, chpp.ClubAPIFile, chpp.ClubAPIVersion, values)
}

// GetCupMatchesXML fetches the raw cupMatches CHPP file over HTTP and decodes it into a CupMatchesXML.
func (p *parsed) GetCupMatchesXML(values map[string]string) (*chpp.CupMatchesXML, error) {
	return fetchXML[chpp.CupMatchesXML](p, chpp.CupMatchesAPIFile, chpp.CupMatchesAPIVersion, values)
}

// GetCurrentBidsXML fetches the raw currentBids CHPP file over HTTP and decodes it into a CurrentBidsXML.
func (p *parsed) GetCurrentBidsXML(values map[string]string) (*chpp.CurrentBidsXML, error) {
	return fetchXML[chpp.CurrentBidsXML](p, chpp.CurrentBidsAPIFile, chpp.CurrentBidsAPIVersion, values)
}

// GetEconomyXML fetches the raw economy CHPP file over HTTP and decodes it into an EconomyXML.
func (p *parsed) GetEconomyXML(values map[string]string) (*chpp.EconomyXML, error) {
	return fetchXML[chpp.EconomyXML](p, chpp.EconomyAPIFile, chpp.EconomyAPIVersion, values)
}

// GetFansXML fetches the raw fans CHPP file over HTTP and decodes it into a FansXML.
func (p *parsed) GetFansXML(values map[string]string) (*chpp.FansXML, error) {
	return fetchXML[chpp.FansXML](p, chpp.FansAPIFile, chpp.FansAPIVersion, values)
}

// GetHOFPlayersXML fetches the raw hOFPlayers CHPP file over HTTP and decodes it into an HOFPlayersXML.
func (p *parsed) GetHOFPlayersXML(values map[string]string) (*chpp.HOFPlayersXML, error) {
	return fetchXML[chpp.HOFPlayersXML](p, chpp.HOFPlayersAPIFile, chpp.HOFPlayersAPIVersion, values)
}

// GetLadderDetailsXML fetches the raw ladderDetails CHPP file over HTTP and decodes it into a LadderDetailsXML.
func (p *parsed) GetLadderDetailsXML(values map[string]string) (*chpp.LadderDetailsXML, error) {
	return fetchXML[chpp.LadderDetailsXML](p, chpp.LadderDetailsAPIFile, chpp.LadderDetailsAPIVersion, values)
}

// GetLadderListXML fetches the raw ladderList CHPP file over HTTP and decodes it into a LadderListXML.
func (p *parsed) GetLadderListXML(values map[string]string) (*chpp.LadderListXML, error) {
	return fetchXML[chpp.LadderListXML](p, chpp.LadderListAPIFile, chpp.LadderListAPIVersion, values)
}

// GetLeagueDetailsXML fetches the raw leagueDetails CHPP file over HTTP and decodes it into a LeagueDetailsXML.
func (p *parsed) GetLeagueDetailsXML(values map[string]string) (*chpp.LeagueDetailsXML, error) {
	return fetchXML[chpp.LeagueDetailsXML](p, chpp.LeagueDetailsAPIFile, chpp.LeagueDetailsAPIVersion, values)
}

// GetLeagueLevelsXML fetches the raw leagueLevels CHPP file over HTTP and decodes it into a LeagueLevelsXML.
func (p *parsed) GetLeagueLevelsXML(values map[string]string) (*chpp.LeagueLevelsXML, error) {
	return fetchXML[chpp.LeagueLevelsXML](p, chpp.LeagueLevelsAPIFile, chpp.LeagueLevelsAPIVersion, values)
}

// GetLeagueFixturesXML fetches the raw leagueFixtures CHPP file over HTTP and decodes it into a LeagueFixturesXML.
func (p *parsed) GetLeagueFixturesXML(values map[string]string) (*chpp.LeagueFixturesXML, error) {
	return fetchXML[chpp.LeagueFixturesXML](p, chpp.LeagueFixturesAPIFile, chpp.LeagueFixturesAPIVersion, values)
}

// GetLiveXML fetches the raw live CHPP file over HTTP and decodes it into a LiveXML.
func (p *parsed) GetLiveXML(values map[string]string) (*chpp.LiveXML, error) {
	return fetchXML[chpp.LiveXML](p, chpp.LiveAPIFile, chpp.LiveAPIVersion, values)
}

// GetManagerCompendiumXML fetches the raw managerCompendium CHPP file over HTTP and decodes it into a ManagerCompendiumXML.
func (p *parsed) GetManagerCompendiumXML(values map[string]string) (*chpp.ManagerCompendiumXML, error) {
	return fetchXML[chpp.ManagerCompendiumXML](p, chpp.ManagerCompendiumAPIFile, chpp.ManagerCompendiumAPIVersion, values)
}

// GetMatchesXML fetches the raw matches CHPP file over HTTP and decodes it into a MatchesXML.
func (p *parsed) GetMatchesXML(values map[string]string) (*chpp.MatchesXML, error) {
	return fetchXML[chpp.MatchesXML](p, chpp.MatchesAPIFile, chpp.MatchesAPIVersion, values)
}

// GetMatchesArchiveXML fetches the raw matchesArchive CHPP file over HTTP and decodes it into a MatchesArchiveXML.
func (p *parsed) GetMatchesArchiveXML(values map[string]string) (*chpp.MatchesArchiveXML, error) {
	return fetchXML[chpp.MatchesArchiveXML](p, chpp.MatchesArchiveAPIFile, chpp.MatchesArchiveAPIVersion, values)
}

// GetMatchDetailsXML fetches the raw matchDetails CHPP file over HTTP and decodes it into a MatchDetailsXML.
func (p *parsed) GetMatchDetailsXML(values map[string]string) (*chpp.MatchDetailsXML, error) {
	return fetchXML[chpp.MatchDetailsXML](p, chpp.MatchDetailsAPIFile, chpp.MatchDetailsAPIVersion, values)
}

// GetMatchLineupXML fetches the raw matchLineup CHPP file over HTTP and decodes it into a MatchLineupXML.
func (p *parsed) GetMatchLineupXML(values map[string]string) (*chpp.MatchLineupXML, error) {
	return fetchXML[chpp.MatchLineupXML](p, chpp.MatchLineupAPIFile, chpp.MatchLineupAPIVersion, values)
}

// GetMatchOrdersXML fetches the raw matchOrders CHPP file over HTTP and decodes it into a MatchOrdersXML.
func (p *parsed) GetMatchOrdersXML(values map[string]string) (*chpp.MatchOrdersXML, error) {
	return fetchXML[chpp.MatchOrdersXML](p, chpp.MatchOrdersAPIFile, chpp.MatchOrdersAPIVersion, values)
}

// SetMatchOrdersXML posts a new lineup to the matchorders CHPP file over
// HTTP and decodes the resulting confirmation into a MatchOrdersXML.
func (p *parsed) SetMatchOrdersXML(values map[string]string, lineup string) (*chpp.MatchOrdersXML, error) {
	buf, err := p.raw.postRawXML(chpp.MatchOrdersAPIFile, chpp.MatchOrdersAPIVersion, values, lineup)
	if err != nil {
		return nil, err
	}
	return chpp.DecodeXML[chpp.MatchOrdersXML](buf, chpp.MatchOrdersAPIFile)
}

// GetNationalTeamsXML fetches the raw nationalTeams CHPP file over HTTP and decodes it into a NationalTeamsXML.
func (p *parsed) GetNationalTeamsXML(values map[string]string) (*chpp.NationalTeamsXML, error) {
	return fetchXML[chpp.NationalTeamsXML](p, chpp.NationalTeamsAPIFile, chpp.NationalTeamsAPIVersion, values)
}

// GetNationalTeamDetailsXML fetches the raw nationalTeamDetails CHPP file over HTTP and decodes it into a NationalTeamDetailsXML.
func (p *parsed) GetNationalTeamDetailsXML(values map[string]string) (*chpp.NationalTeamDetailsXML, error) {
	return fetchXML[chpp.NationalTeamDetailsXML](p, chpp.NationalTeamDetailsAPIFile, chpp.NationalTeamDetailsAPIVersion, values)
}

// GetNationalTeamMatchesXML fetches the raw nationalTeamMatches CHPP file over HTTP and decodes it into a NationalTeamMatchesXML.
func (p *parsed) GetNationalTeamMatchesXML(values map[string]string) (*chpp.NationalTeamMatchesXML, error) {
	return fetchXML[chpp.NationalTeamMatchesXML](p, chpp.NationalTeamMatchesAPIFile, chpp.NationalTeamMatchesAPIVersion, values)
}

// GetNationalPlayersXML fetches the raw nationalPlayers CHPP file over HTTP and decodes it into a NationalPlayersXML.
func (p *parsed) GetNationalPlayersXML(values map[string]string) (*chpp.NationalPlayersXML, error) {
	return fetchXML[chpp.NationalPlayersXML](p, chpp.NationalPlayersAPIFile, chpp.NationalPlayersAPIVersion, values)
}

// GetPlayersXML fetches the raw players CHPP file over HTTP and decodes it into a PlayersXML.
func (p *parsed) GetPlayersXML(values map[string]string) (*chpp.PlayersXML, error) {
	return fetchXML[chpp.PlayersXML](p, chpp.PlayersAPIFile, chpp.PlayersAPIVersion, values)
}

// GetPlayerDetailsXML fetches the raw playerDetails CHPP file over HTTP and decodes it into a PlayerDetailsXML.
func (p *parsed) GetPlayerDetailsXML(values map[string]string) (*chpp.PlayerDetailsXML, error) {
	return fetchXML[chpp.PlayerDetailsXML](p, chpp.PlayerDetailsAPIFile, chpp.PlayerDetailsAPIVersion, values)
}

// GetPlayerEventsXML fetches the raw playerEvents CHPP file over HTTP and decodes it into a PlayerEventsXML.
func (p *parsed) GetPlayerEventsXML(values map[string]string) (*chpp.PlayerEventsXML, error) {
	return fetchXML[chpp.PlayerEventsXML](p, chpp.PlayerEventsAPIFile, chpp.PlayerEventsAPIVersion, values)
}

// GetRegionDetailsXML fetches the raw regionDetails CHPP file over HTTP and decodes it into a RegionDetailsXML.
func (p *parsed) GetRegionDetailsXML(values map[string]string) (*chpp.RegionDetailsXML, error) {
	return fetchXML[chpp.RegionDetailsXML](p, chpp.RegionDetailsAPIFile, chpp.RegionDetailsAPIVersion, values)
}

// GetSearchXML fetches the raw search CHPP file over HTTP and decodes it into a SearchXML.
func (p *parsed) GetSearchXML(values map[string]string) (*chpp.SearchXML, error) {
	return fetchXML[chpp.SearchXML](p, chpp.SearchAPIFile, chpp.SearchAPIVersion, values)
}

// GetStaffAvatarsXML fetches the raw staffAvatars CHPP file over HTTP and decodes it into a StaffAvatarsXML.
func (p *parsed) GetStaffAvatarsXML(values map[string]string) (*chpp.StaffAvatarsXML, error) {
	return fetchXML[chpp.StaffAvatarsXML](p, chpp.StaffAvatarsAPIFile, chpp.StaffAvatarsAPIVersion, values)
}

// GetStaffListXML fetches the raw staffList CHPP file over HTTP and decodes it into a StaffListXML.
func (p *parsed) GetStaffListXML(values map[string]string) (*chpp.StaffListXML, error) {
	return fetchXML[chpp.StaffListXML](p, chpp.StaffListAPIFile, chpp.StaffListAPIVersion, values)
}

// GetSupportersXML fetches the raw supporters CHPP file over HTTP and decodes it into a SupportersXML.
func (p *parsed) GetSupportersXML(values map[string]string) (*chpp.SupportersXML, error) {
	return fetchXML[chpp.SupportersXML](p, chpp.SupportersAPIFile, chpp.SupportersAPIVersion, values)
}

// GetTeamDetailsXML fetches the raw teamDetails CHPP file over HTTP and decodes it into a TeamDetailsXML.
func (p *parsed) GetTeamDetailsXML(values map[string]string) (*chpp.TeamDetailsXML, error) {
	return fetchXML[chpp.TeamDetailsXML](p, chpp.TeamDetailsAPIFile, chpp.TeamDetailsAPIVersion, values)
}

// GetTournamentDetailsXML fetches the raw tournamentDetails CHPP file over HTTP and decodes it into a TournamentDetailsXML.
func (p *parsed) GetTournamentDetailsXML(values map[string]string) (*chpp.TournamentDetailsXML, error) {
	return fetchXML[chpp.TournamentDetailsXML](p, chpp.TournamentDetailsAPIFile, chpp.TournamentDetailsAPIVersion, values)
}

// GetTournamentFixturesXML fetches the raw tournamentFixtures CHPP file over HTTP and decodes it into a TournamentFixturesXML.
func (p *parsed) GetTournamentFixturesXML(values map[string]string) (*chpp.TournamentFixturesXML, error) {
	return fetchXML[chpp.TournamentFixturesXML](p, chpp.TournamentFixturesAPIFile, chpp.TournamentFixturesAPIVersion, values)
}

// GetTournamentLeagueTablesXML fetches the raw tournamentLeagueTables CHPP file over HTTP and decodes it into a TournamentLeagueTablesXML.
func (p *parsed) GetTournamentLeagueTablesXML(values map[string]string) (*chpp.TournamentLeagueTablesXML, error) {
	return fetchXML[chpp.TournamentLeagueTablesXML](p, chpp.TournamentLeagueTablesAPIFile, chpp.TournamentLeagueTablesAPIVersion, values)
}

// GetTournamentListXML fetches the raw tournamentList CHPP file over HTTP and decodes it into a TournamentListXML.
func (p *parsed) GetTournamentListXML(values map[string]string) (*chpp.TournamentListXML, error) {
	return fetchXML[chpp.TournamentListXML](p, chpp.TournamentListAPIFile, chpp.TournamentListAPIVersion, values)
}

// GetTrainingXML fetches the raw training CHPP file over HTTP and decodes it into a TrainingXML.
func (p *parsed) GetTrainingXML(values map[string]string) (*chpp.TrainingXML, error) {
	return fetchXML[chpp.TrainingXML](p, chpp.TrainingAPIFile, chpp.TrainingAPIVersion, values)
}

// GetTrainingEventsXML fetches the raw trainingEvents CHPP file over HTTP and decodes it into a TrainingEventsXML.
func (p *parsed) GetTrainingEventsXML(values map[string]string) (*chpp.TrainingEventsXML, error) {
	return fetchXML[chpp.TrainingEventsXML](p, chpp.TrainingEventsAPIFile, chpp.TrainingEventsAPIVersion, values)
}

// GetTransferSearchXML fetches the raw transferSearch CHPP file over HTTP and decodes it into a TransferSearchXML.
func (p *parsed) GetTransferSearchXML(values map[string]string) (*chpp.TransferSearchXML, error) {
	return fetchXML[chpp.TransferSearchXML](p, chpp.TransferSearchAPIFile, chpp.TransferSearchAPIVersion, values)
}

// GetTransferSteamXML fetches the raw transferSteam CHPP file over HTTP and decodes it into a TransferSteamXML.
func (p *parsed) GetTransferSteamXML(values map[string]string) (*chpp.TransferSteamXML, error) {
	return fetchXML[chpp.TransferSteamXML](p, chpp.TransferSteamAPIFile, chpp.TransferSteamAPIVersion, values)
}

// GetTransfersPlayerXML fetches the raw transfersPlayer CHPP file over HTTP and decodes it into a TransfersPlayerXML.
func (p *parsed) GetTransfersPlayerXML(values map[string]string) (*chpp.TransfersPlayerXML, error) {
	return fetchXML[chpp.TransfersPlayerXML](p, chpp.TransfersPlayerAPIFile, chpp.TransfersPlayerAPIVersion, values)
}

// GetTranslationsXML fetches the raw translations CHPP file over HTTP and decodes it into a TranslationsXML.
func (p *parsed) GetTranslationsXML(values map[string]string) (*chpp.TranslationsXML, error) {
	return fetchXML[chpp.TranslationsXML](p, chpp.TranslationsAPIFile, chpp.TranslationsAPIVersion, values)
}

// GetWorldCupXML fetches the raw worldCup CHPP file over HTTP and decodes it into a WorldCupXML.
func (p *parsed) GetWorldCupXML(values map[string]string) (*chpp.WorldCupXML, error) {
	return fetchXML[chpp.WorldCupXML](p, chpp.WorldCupAPIFile, chpp.WorldCupAPIVersion, values)
}

// GetWorldDetailsXML fetches the raw worldDetails CHPP file over HTTP and decodes it into a WorldDetailsXML.
func (p *parsed) GetWorldDetailsXML(values map[string]string) (*chpp.WorldDetailsXML, error) {
	return fetchXML[chpp.WorldDetailsXML](p, chpp.WorldDetailsAPIFile, chpp.WorldDetailsAPIVersion, values)
}

// GetWorldLanguagesXML fetches the raw worldLanguages CHPP file over HTTP and decodes it into a WorldLanguagesXML.
func (p *parsed) GetWorldLanguagesXML(values map[string]string) (*chpp.WorldLanguagesXML, error) {
	return fetchXML[chpp.WorldLanguagesXML](p, chpp.WorldLanguagesAPIFile, chpp.WorldLanguagesAPIVersion, values)
}

// GetYouthAvatarsXML fetches the raw youthAvatars CHPP file over HTTP and decodes it into a YouthAvatarsXML.
func (p *parsed) GetYouthAvatarsXML(values map[string]string) (*chpp.YouthAvatarsXML, error) {
	return fetchXML[chpp.YouthAvatarsXML](p, chpp.YouthAvatarsAPIFile, chpp.YouthAvatarsAPIVersion, values)
}

// GetYouthLeagueDetailsXML fetches the raw youthLeagueDetails CHPP file over HTTP and decodes it into a YouthLeagueDetailsXML.
func (p *parsed) GetYouthLeagueDetailsXML(values map[string]string) (*chpp.YouthLeagueDetailsXML, error) {
	return fetchXML[chpp.YouthLeagueDetailsXML](p, chpp.YouthLeagueDetailsAPIFile, chpp.YouthLeagueDetailsAPIVersion, values)
}

// GetYouthLeagueFixturesXML fetches the raw youthLeagueFixtures CHPP file over HTTP and decodes it into a YouthLeagueFixturesXML.
func (p *parsed) GetYouthLeagueFixturesXML(values map[string]string) (*chpp.YouthLeagueFixturesXML, error) {
	return fetchXML[chpp.YouthLeagueFixturesXML](p, chpp.YouthLeagueFixturesAPIFile, chpp.YouthLeagueFixturesAPIVersion, values)
}

// GetYouthPlayerDetailsXML fetches the raw youthPlayerDetails CHPP file over HTTP and decodes it into a YouthPlayerDetailsXML.
func (p *parsed) GetYouthPlayerDetailsXML(values map[string]string) (*chpp.YouthPlayerDetailsXML, error) {
	return fetchXML[chpp.YouthPlayerDetailsXML](p, chpp.YouthPlayerDetailsAPIFile, chpp.YouthPlayerDetailsAPIVersion, values)
}

// GetYouthPlayerListXML fetches the raw youthPlayerList CHPP file over HTTP and decodes it into a YouthPlayerListXML.
func (p *parsed) GetYouthPlayerListXML(values map[string]string) (*chpp.YouthPlayerListXML, error) {
	return fetchXML[chpp.YouthPlayerListXML](p, chpp.YouthPlayerListAPIFile, chpp.YouthPlayerListAPIVersion, values)
}

// GetYouthTeamDetailsXML fetches the raw youthTeamDetails CHPP file over HTTP and decodes it into a YouthTeamDetailsXML.
func (p *parsed) GetYouthTeamDetailsXML(values map[string]string) (*chpp.YouthTeamDetailsXML, error) {
	return fetchXML[chpp.YouthTeamDetailsXML](p, chpp.YouthTeamDetailsAPIFile, chpp.YouthTeamDetailsAPIVersion, values)
}
