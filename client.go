package hattrick

import (
	"time"

	"github.com/lucianoq/hattrick/api"
	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// Client ...
type Client interface {
	GetMyAchievements() ([]*chpp.Achievement, error)
	GetAchievements(userID id.User) ([]*chpp.Achievement, error)

	GetAllianceDetails(allianceID id.Alliance) (*chpp.Alliance, error)
	GetAllianceDetailsRoles(allianceID id.Alliance) ([]*chpp.AllianceRole, error)
	GetAllianceDetailsMembers(allianceID id.Alliance) ([]*chpp.AllianceMember, error)
	GetAllianceDetailsMembersSubset(allianceID id.Alliance, subset uint) ([]*chpp.AllianceMember, error)
	GetAlliancesNameStartsWith(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliancesAbbreviationIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliancesDescriptionIncludes(searchFor string, searchLanguageID id.Language, pageIndex uint) ([]*chpp.Alliance, error)
	GetAlliance(searchFor id.Alliance, searchLanguageID id.Language, pageIndex uint) (*chpp.Alliance, error)
	GetMyAlliances() ([]*chpp.Alliance, error)

	GetMyArena() (*chpp.Arena, error)
	GetArena(arenaID id.Arena) (*chpp.Arena, error)
	GetArenaByTeamID(teamID id.Team) (*chpp.Arena, error)

	GetAvatarsMyPlayers() ([]*chpp.PlayerAvatars, error)
	GetAvatarsPlayers(teamID id.Team) ([]*chpp.PlayerAvatars, error)
	GetAvatarsMyHallOfFame() ([]*chpp.PlayerAvatars, error)
	GetAvatarsHallOfFame(teamID id.Team) ([]*chpp.PlayerAvatars, error)

	GetBookmarks(bookmarkType chpp.BookmarkType) ([]*chpp.Bookmark, error)

	GetChallenges(weekend bool) ([]*chpp.ChallengeByMe, []*chpp.OffersByOthers, error)
	IsChallengeable(weekend bool, teamID id.Team) (bool, error)
	AreChallengeable(weekend bool, teamIDs ...id.Team) ([]bool, error)
	Challenge(opponentTeam id.Team, friendlyType chpp.FriendlyType, matchPlace chpp.MatchPlace, otherArena id.Arena, weekend bool) error
	AcceptChallenge(friendlyMatchID id.FriendlyMatch) error
	DeclineChallenge(friendlyMatchID id.FriendlyMatch) error
	WithdrawChallenge(friendlyMatchID id.FriendlyMatch) error

	GetMyClub() (*chpp.Club, error)
	GetClubByID(teamID id.Team) (*chpp.Club, error)

	GetCupMatchesLast(cup id.Cup) ([]*chpp.CupMatch, error)
	GetCupMatches(cup id.Cup, season, round uint) ([]*chpp.CupMatch, error)

	GetMatch(matchID id.Match) (*chpp.MatchDetails, error)
	GetMatchesArchive(teamID id.Team, start, end time.Time) ([]*chpp.Match, error)
	GetMyMatchesArchive(start, end time.Time) ([]*chpp.Match, error)
	GetMatchLineup(matchID id.Match, teamID id.Team) (*chpp.MatchLineup, error)
	GetMyMatchLineup(matchID id.Match) (*chpp.MatchLineup, error)

	GetCurrentBids() ([]*chpp.BidItem, error)
	IgnoreTransfer(transferID id.Transfer, category chpp.TrackingTypeID) error
	DeleteAllFinishedBids() error

	GetFinancesForPrimaryTeam() (chpp.Economy, error)
	GetFinances(teamID id.Team) (chpp.Economy, error)

	GetHOFPlayers() ([]*chpp.HOFPlayer, error)

	GetMySeries() (*chpp.Series, error)
	GetSeries(leagueLevelUnitID id.Series) (*chpp.Series, error)

	GetMySeriesFixtures() (*chpp.SeriesFixtures, error)
	GetMySeriesFixturesBySeason(season uint) (*chpp.SeriesFixtures, error)
	GetSeriesFixtures(leagueLevelUnitID id.Series) (*chpp.SeriesFixtures, error)
	GetSeriesFixturesBySeason(leagueLevelUnitID id.Series, season uint) (*chpp.SeriesFixtures, error)

	GetMyPlayers() ([]*chpp.Player, error)
	GetPlayers(teamID id.Team) ([]*chpp.Player, error)

	GetMyTeams() ([]*chpp.Team, error)
	GetMyPrimaryTeam() (*chpp.Team, error)
	GetTeam(teamID id.Team) (*chpp.Team, error)
	GetPrimaryTeamByUser(userID id.User) (*chpp.Team, error)
	GetTeamsByUser(userID id.User) ([]*chpp.Team, error)

	GetMe() (*chpp.Manager, error)
	GetManager(uID id.User) (*chpp.Manager, error)

	GetMyMatches() ([]*chpp.Match, error)
	GetMyYouthMatches() ([]*chpp.Match, error)

	GetWorld() ([]*chpp.League, error)
	GetLeague(leagueID id.League) (*chpp.League, error)
	GetCountry(countryID id.Country) (*chpp.League, error)
}

// NewClient ...
func NewClient(cfg AuthConfig) (Client, error) {
	return api.NewAPI(
		cfg.ConsumerKey,
		cfg.ConsumerSecret,
		cfg.AccessToken,
		cfg.AccessSecret,
		cfg.AccessAdditionalData,
	)
}

// AuthConfig ...
type AuthConfig struct {
	ConsumerKey          string
	ConsumerSecret       string
	AccessToken          string
	AccessSecret         string
	AccessAdditionalData map[string]string
}
