package chpp

// ErrorCode ...
type ErrorCode int

// List of ErrorCode constants.
const (
	NoInformation                            ErrorCode = -1
	NotLoggedIn                              ErrorCode = 0
	AccessDenied                             ErrorCode = 1
	FileNotSpecified                         ErrorCode = 2
	FileNotSupported                         ErrorCode = 3
	POSTNotSupported                         ErrorCode = 4
	MustUsePOSTForThisAction                 ErrorCode = 5
	OnlyForSupporters                        ErrorCode = 6
	NotSupportedVersion                      ErrorCode = 7
	InvalidParameter                         ErrorCode = 10
	UnknownTeamID                            ErrorCode = 50
	UnknownMatchID                           ErrorCode = 51
	UnknownActionType                        ErrorCode = 52
	MatchIDNotSubscribedTo                   ErrorCode = 53
	UnknownYouthTeamID                       ErrorCode = 54
	UnknownYouthPlayerID                     ErrorCode = 55
	UnknownPlayerID                          ErrorCode = 56
	UnknownLeagueID                          ErrorCode = 57
	UnknownLeagueLevelUnitID                 ErrorCode = 58
	AllowedOnlyOnYourTeam                    ErrorCode = 59
	UnknownTournamentID                      ErrorCode = 60
	UnknownLadderID                          ErrorCode = 61
	ChallengeFailed                          ErrorCode = 70
	BidFailed                                ErrorCode = 71
	HattrickDownForMaintenance               ErrorCode = 90
	HattrickYouthAcademyDownForMaintenance   ErrorCode = 91
	HattrickTransferMarketDownForMaintenance ErrorCode = 92
	HTOIntegratedPlatformDownForMaintenance  ErrorCode = 93
	UndefinedError                           ErrorCode = 99
)
