// Package hattrick is the top-level entry point for this CHPP client
// library: NewClient authenticates with OAuth1 credentials and returns an
// *api.API exposing one method per Hattrick CHPP operation.
package hattrick

import (
	"time"

	"github.com/lucianoq/hattrick/api"
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// Client is the full public interface of this CHPP client library. Each
// method corresponds to one operation of the Hattrick CHPP API (fetching a
// resource such as a team, player, or match, or performing an action such as
// placing a bid or setting training), and returns already-decoded, typed
// results. Method names mirror the CHPP file/action they wrap, so the
// signature is generally self-explanatory; see the CHPP documentation
// (https://www.hattrick.org/en/chpp/) for the semantics of individual calls.
type Client interface {
	// achievements
	GetMyAchievements() ([]*chpp.Achievement, error)
	GetAchievements(userID id.User) ([]*chpp.Achievement, error)

	// alliancedetails
	GetAllianceDetails(allianceID id.Alliance) (*chpp.Alliance, error)
	GetAllianceDetailsRoles(allianceID id.Alliance) ([]*chpp.AllianceRole, error)
	GetAllianceDetailsMembers(allianceID id.Alliance) ([]*chpp.AllianceMember, error)
	GetAllianceDetailsMembersSubset(allianceID id.Alliance, subset uint) ([]*chpp.AllianceMember, error)

	// alliances
	GetAlliancesNameStartsWith(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliancesAbbreviationIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliancesDescriptionIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliance(searchFor id.Alliance, searchLanguageID id.Language, pageIndex uint) (*chpp.Alliance, error)
	GetMyAlliances() ([]*chpp.Alliance, error)

	// arenadetails
	GetMyArena() (*chpp.Arena, error)
	GetArena(arenaID id.Arena) (*chpp.Arena, error)
	GetArenaByTeamID(teamID id.Team) (*chpp.Arena, error)
	GetArenaStats(leagueID id.League) (*chpp.LeagueArenaStats, error)
	GetMyArenaSupporterStats() (*chpp.MyArenaStats, error)

	// avatars
	GetAvatarsMyPlayers() ([]*chpp.PlayerAvatars, error)
	GetAvatarsPlayers(teamID id.Team) ([]*chpp.PlayerAvatars, error)
	GetAvatarsMyHallOfFame() ([]*chpp.PlayerAvatars, error)
	GetAvatarsHallOfFame(teamID id.Team) ([]*chpp.PlayerAvatars, error)

	// bookmarks
	GetBookmarks(bookmarkType chpp.BookmarkType) ([]*chpp.Bookmark, error)

	// challenges
	GetChallenges(weekend bool) ([]*chpp.ChallengeByMe, []*chpp.OffersByOthers, error)
	IsChallengeable(weekend bool, teamID id.Team) (bool, error)
	AreChallengeable(weekend bool, teamIDs ...id.Team) ([]bool, error)
	Challenge(opponentTeam id.Team, friendlyType chpp.FriendlyType, matchPlace chpp.MatchPlace, otherArena id.Arena, weekend bool) error
	AcceptChallenge(friendlyMatchID id.FriendlyMatch) error
	DeclineChallenge(friendlyMatchID id.FriendlyMatch) error
	WithdrawChallenge(friendlyMatchID id.FriendlyMatch) error

	// club
	GetMyClub() (*chpp.Club, error)
	GetClubByID(teamID id.Team) (*chpp.Club, error)

	// cupmatches
	GetCupMatchesLast(cup id.Cup) ([]*chpp.CupMatch, error)
	GetCupMatches(cup id.Cup, season, round uint) ([]*chpp.CupMatch, error)

	// currentbids
	GetCurrentBids() ([]*chpp.BidItem, error)
	IgnoreTransfer(transferID id.Transfer, category chpp.TrackingTypeID) error
	DeleteAllFinishedBids() error

	// economy
	GetFinancesForPrimaryTeam() (chpp.Economy, error)
	GetFinances(teamID id.Team) (chpp.Economy, error)

	// fans
	GetFansForPrimaryTeam() (chpp.Fans, error)
	GetFans(teamID id.Team) (chpp.Fans, error)

	// hofplayers
	GetHOFPlayersForPrimaryTeam() ([]*chpp.HOFPlayer, error)
	GetHOFPlayers(teamID id.Team) ([]*chpp.HOFPlayer, error)

	// ladderdetails
	GetLadderDetails(ladderID id.Ladder) (*chpp.LadderDetails, error)
	GetLadderDetailsForTeam(ladderID id.Ladder, teamID id.Team) (*chpp.LadderDetails, error)
	GetLadderDetailsPage(ladderID id.Ladder, pageIndex, pageSize uint) (*chpp.LadderDetails, error)

	// ladderlist
	GetMyLadders() ([]*chpp.LadderListEntry, error)
	GetLadders(teamID id.Team) ([]*chpp.LadderListEntry, error)

	// leaguedetails
	GetMySeries() (*chpp.Series, error)
	GetSeries(seriesID id.Series) (*chpp.Series, error)

	// leaguefixtures
	GetMySeriesFixtures() (*chpp.SeriesFixtures, error)
	GetMySeriesFixturesBySeason(season uint) (*chpp.SeriesFixtures, error)
	GetSeriesFixtures(seriesID id.Series) (*chpp.SeriesFixtures, error)
	GetSeriesFixturesBySeason(seriesID id.Series, season uint) (*chpp.SeriesFixtures, error)

	// leaguelevels
	GetMyLeagueLevels() ([]*chpp.LeagueLevelItem, error)
	GetLeagueLevels(leagueID id.League) ([]*chpp.LeagueLevelItem, error)

	// live
	GetLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error)
	GetLiveMatches(sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error)
	GetAllLiveEvents(includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error)
	GetNewLiveEvents(lastShownIndexes string, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error)
	ClearLiveMatches() error
	AddLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem, includeStartingLineup, useLiveEventsAndTexts bool) ([]*chpp.LiveMatchInfo, error)
	DeleteLiveMatch(matchID id.Match, sourceSystem chpp.SourceSystem) ([]*chpp.LiveMatchInfo, error)

	// managercompendium
	GetMe() (*chpp.Manager, error)
	GetManager(uID id.User) (*chpp.Manager, error)

	// matchdetails
	GetMatch(matchID id.Match) (*chpp.MatchDetails, error)
	GetMatchDetails(matchID id.Match, sourceSystem chpp.SourceSystem, matchEvents bool) (*chpp.MatchDetails, error)

	// matches
	GetMyMatches() ([]*chpp.Match, error)
	GetMyYouthMatches() ([]*chpp.Match, error)
	GetMatches(teamID id.Team) ([]*chpp.Match, error)
	GetYouthMatches(teamID id.Team) ([]*chpp.Match, error)
	GetMyMatchesUntil(lastMatchDate time.Time) ([]*chpp.Match, error)
	GetMyYouthMatchesUntil(lastMatchDate time.Time) ([]*chpp.Match, error)
	GetMatchesUntil(teamID id.Team, lastMatchDate time.Time) ([]*chpp.Match, error)
	GetYouthMatchesUntil(teamID id.Team, lastMatchDate time.Time) ([]*chpp.Match, error)

	// matchesArchive
	GetMatchesArchive(teamID id.Team, start, end time.Time) ([]*chpp.Match, error)
	GetMyMatchesArchive(start, end time.Time) ([]*chpp.Match, error)
	GetYouthMatchesArchive(youthTeamID id.Team, start, end time.Time) ([]*chpp.Match, error)
	GetMyYouthMatchesArchive(start, end time.Time) ([]*chpp.Match, error)
	GetMatchesArchiveIncludeHTO(teamID id.Team, start, end time.Time) ([]*chpp.Match, error)
	GetMyMatchesArchiveIncludeHTO(start, end time.Time) ([]*chpp.Match, error)
	GetMatchesArchiveBySeason(teamID id.Team, season uint) ([]*chpp.Match, error)
	GetMyMatchesArchiveBySeason(season uint) ([]*chpp.Match, error)

	// matchlineup
	GetMatchLineup(matchID id.Match, teamID id.Team) (*chpp.MatchLineup, error)
	GetMyMatchLineup(matchID id.Match) (*chpp.MatchLineup, error)
	GetMatchLineupBySourceSystem(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem) (*chpp.MatchLineup, error)

	// matchorders
	GetMatchOrders(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem) (*chpp.MatchOrdersMatchData, error)
	SetMatchOrders(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem, lineup string) (*chpp.MatchOrdersMatchData, error)
	PredictRatings(matchID id.Match, teamID id.Team, sourceSystem chpp.SourceSystem, lineup string) (*chpp.MatchOrdersMatchData, error)

	// nationalplayers
	GetNationalTeamPlayers(teamID id.NationalTeam) ([]*chpp.NationalTeamPlayer, error)
	GetNationalTeamPlayerStats(teamID id.NationalTeam, matchTypeCategory string, showAll bool) (*chpp.NationalPlayersStats, error)

	// nationalteamdetails
	GetNationalTeamDetails(teamID id.NationalTeam) (*chpp.NationalTeamDetails, error)

	// nationalteammatches
	GetNationalTeamMatches(leagueOfficeTypeID chpp.LeagueOfficeTypeID) ([]*chpp.NationalTeamMatch, error)

	// nationalteams
	GetNationalTeams(leagueOfficeTypeID chpp.LeagueOfficeTypeID) (*chpp.NationalTeamsList, error)

	// playerdetails
	GetPlayerDetails(playerID id.Player, includeMatchInfo bool) (*chpp.PlayerDetails, error)
	PlaceBidOnPlayer(playerID id.Player, teamID id.Team, bidAmount, maxBidAmount chpp.Money) (*chpp.PlayerDetails, error)

	// playerevents
	GetPlayerEvents(playerID id.Player) ([]*chpp.PlayerEventItem, error)

	// players
	GetMyPlayers() ([]*chpp.Player, error)
	GetPlayers(teamID id.Team) ([]*chpp.Player, error)
	GetMyPlayersOrderedBy(orderBy string) ([]*chpp.Player, error)
	GetPlayersOrderedBy(teamID id.Team, orderBy string) ([]*chpp.Player, error)
	GetMyOldPlayers() ([]*chpp.Player, error)
	GetOldPlayers(teamID id.Team) ([]*chpp.Player, error)
	GetMyOldCoaches() ([]*chpp.Player, error)
	GetOldCoaches(teamID id.Team) ([]*chpp.Player, error)

	// regiondetails
	GetMyRegion() (*chpp.RegionDetails, error)
	GetRegion(regionID id.Region) (*chpp.RegionDetails, error)

	// search
	Search(searchType chpp.SearchType, searchString string, filters api.SearchFilters) ([]*chpp.SearchResult, error)

	// staffavatars
	GetMyStaffAvatars() (*chpp.StaffAvatars, error)
	GetStaffAvatars(teamID id.Team) (*chpp.StaffAvatars, error)

	// stafflist
	GetMyStaff() (*chpp.StaffList, error)
	GetStaff(teamID id.Team) (*chpp.StaffList, error)

	// supporters
	GetMySupportedTeams(pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error)
	GetSupportedTeams(userID id.User, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error)
	GetMySupporters(pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error)
	GetSupportersForTeam(teamID id.Team, pageIndex, pageSize uint) ([]*chpp.SupporterTeam, error)

	// teamdetails
	GetMyTeams() ([]*chpp.Team, error)
	GetMyPrimaryTeam() (*chpp.Team, error)
	GetTeam(teamID id.Team) (*chpp.Team, error)
	GetPrimaryTeamByUser(userID id.User) (*chpp.Team, error)
	GetTeamsByUser(userID id.User) ([]*chpp.Team, error)

	// tournamentdetails
	GetTournament(tournamentID id.Tournament) (*chpp.Tournament, error)
	GetTournamentBySeason(tournamentID id.Tournament, season uint) (*chpp.Tournament, error)

	// tournamentfixtures
	GetTournamentFixtures(tournamentID id.Tournament) ([]*chpp.TournamentFixture, error)
	GetTournamentFixturesBySeason(tournamentID id.Tournament, season uint) ([]*chpp.TournamentFixture, error)
	GetTournamentFixturesByMatchRound(tournamentID id.Tournament, matchRound uint) ([]*chpp.TournamentFixture, error)

	// tournamentleaguetables
	GetTournamentLeagueTables(tournamentID id.Tournament) ([]*chpp.TournamentLeagueTable, error)
	GetTournamentLeagueTablesBySeason(tournamentID id.Tournament, season uint) ([]*chpp.TournamentLeagueTable, error)
	GetTournamentLeagueTablesByWorldCupRound(tournamentID id.Tournament, worldCupRound uint) ([]*chpp.TournamentLeagueTable, error)

	// tournamentlist
	GetMyTournaments() ([]*chpp.Tournament, error)
	GetTournaments(teamID id.Team) ([]*chpp.Tournament, error)

	// training
	GetMyTraining() (*chpp.TrainingTeam, error)
	GetTraining(teamID id.Team) (*chpp.TrainingTeam, error)
	GetTrainingStats(leagueID id.League) (*chpp.TrainingLeagueStats, error)
	SetTraining(teamID id.Team, trainingType chpp.TrainingType, trainingLevel, trainingLevelStamina uint) (bool, *chpp.TrainingTeam, error)

	// trainingevents
	GetTrainingEvents(playerID id.Player) ([]*chpp.TrainingEvent, error)

	// transfersearch
	GetTransferSearch(ageMin, ageMax uint, skillType1 chpp.SkillID, minSkillValue1, maxSkillValue1 chpp.SkillLevel, filters api.TransferSearchFilters) (*chpp.TransferSearchResults, error)

	// transfersplayer
	GetPlayerTransferHistory(playerID id.Player) (*chpp.PlayerTransferHistory, error)

	// transfersteam
	GetMyTransfersTeam(pageIndex uint) (*chpp.TransferSteamXML, error)
	GetTransfersTeam(teamID id.Team, pageIndex uint) (*chpp.TransferSteamXML, error)

	// translations
	GetTranslations(languageID id.Language) (*chpp.TranslationsXML, error)

	// worldcup
	GetWorldCupGroups(cupID chpp.WorldCupID, season, matchRound uint) ([]*chpp.WorldCupScore, error)
	GetWorldCupMatches(cupID chpp.WorldCupID, season, cupSeriesUnitID uint) ([]*chpp.WorldCupMatch, error)

	// worlddetails
	GetWorld() ([]*chpp.League, error)
	GetLeague(leagueID id.League) (*chpp.League, error)
	GetCountry(countryID id.Country) (*chpp.League, error)

	// worldlanguages
	GetWorldLanguages() ([]*chpp.WorldLanguage, error)

	// youthavatars
	GetMyYouthAvatars() ([]*chpp.YouthPlayerAvatars, error)
	GetYouthAvatars(youthTeamID id.YouthTeam) ([]*chpp.YouthPlayerAvatars, error)

	// youthleaguedetails
	GetMyYouthLeague() (*chpp.YouthLeagueDetails, error)
	GetYouthLeague(youthLeagueID id.YouthLeague) (*chpp.YouthLeagueDetails, error)

	// youthleaguefixtures
	GetMyYouthLeagueFixtures() (*chpp.YouthLeagueFixtures, error)
	GetYouthLeagueFixtures(youthLeagueID id.YouthLeague, season uint) (*chpp.YouthLeagueFixtures, error)

	// youthplayerdetails
	GetYouthPlayer(youthPlayerID id.YouthPlayer, showScoutCall, showLastMatch bool) (*chpp.YouthPlayerDetail, error)
	UnlockYouthPlayerSkills(youthPlayerID id.YouthPlayer, showScoutCall, showLastMatch bool) (*chpp.YouthPlayerDetail, error)

	// youthplayerlist
	GetMyYouthPlayers() ([]*chpp.YouthPlayerDetail, error)
	GetYouthPlayers(youthTeamID id.YouthTeam) ([]*chpp.YouthPlayerDetail, error)
	GetMyYouthPlayersDetails(showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error)
	GetYouthPlayersDetails(youthTeamID id.YouthTeam, showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error)
	UnlockYouthTeamSkills(youthTeamID id.YouthTeam, showScoutCall, showLastMatch bool) ([]*chpp.YouthPlayerDetail, error)

	// youthteamdetails
	GetMyYouthTeam(showScouts bool) (*chpp.YouthTeamDetails, []*chpp.YouthScout, error)
	GetYouthTeam(youthTeamID id.YouthTeam, showScouts bool) (*chpp.YouthTeamDetails, []*chpp.YouthScout, error)
}

// var _ Client = (*api.API)(nil) is a compile-time assertion that *api.API
// implements Client - the two are otherwise only linked implicitly via
// NewClient's return type, so this is what actually guarantees they stay
// in sync as either one changes.
var _ Client = (*api.API)(nil)

// NewClient builds a Client authenticated against the CHPP API using OAuth1
// credentials: the application's consumer key/secret (issued when the CHPP
// application is registered) and the access token/secret obtained after the
// user completes the OAuth authorization flow (see package login). Some CHPP
// applications also require the AdditionalData values returned alongside the
// access token during that flow.
func NewClient(cfg AuthConfig) (Client, error) {
	return api.NewAPI(
		cfg.ConsumerKey,
		cfg.ConsumerSecret,
		cfg.AccessToken,
		cfg.AccessSecret,
		cfg.AccessAdditionalData,
	)
}

// AuthConfig holds the OAuth1 credentials required by NewClient: the
// consumer key/secret identifying the registered CHPP application, and the
// access token/secret (plus any AccessAdditionalData) obtained for a
// specific user through the OAuth1 login/authorization flow.
type AuthConfig struct {
	ConsumerKey          string
	ConsumerSecret       string
	AccessToken          string
	AccessSecret         string
	AccessAdditionalData map[string]string
}
