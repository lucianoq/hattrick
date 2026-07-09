package chpp

import (
	"github.com/lucianoq/hattrick/chpp/id"
)

// XML file name and version.
const (
	TournamentLeagueTablesAPIFile    = "tournamentleaguetables"
	TournamentLeagueTablesAPIVersion = "1.1"
)

// TournamentLeagueTablesXML contains the league tables (one per group) for
// a tournament, for the current season (and up to 2 seasons after it
// finished, for non-restarted tournaments).
type TournamentLeagueTablesXML struct {
	Envelope
	UserID id.User `xml:"User"`

	TournamentID id.Tournament `xml:"TournamentId"`
	Season       uint          `xml:"Season"`

	// Only sent for World Cup tournaments.
	WorldCupRound uint `xml:"WorldCupRound"`

	Tables []*TournamentLeagueTable `xml:"TournamentLeagueTables>TournamentLeagueTable"`
}

// TournamentLeagueTable is a single group's league table within a
// tournament.
type TournamentLeagueTable struct {
	// The tournament group, 1=A, 2=B and so on.
	GroupID uint `xml:"GroupId"`

	// Confirmed against a live response: Teams is a single wrapper
	// element per group (self-closing "<Teams />" when empty), not
	// itself repeated per team. The exact tag name of the repeated
	// per-team item inside it ("Team") is inferred from this API's
	// otherwise-universal Container>Item naming convention (e.g.
	// Matches>Match, CupTeams>CupTeam) - the live example available had
	// no populated teams to confirm the item's own tag name.
	Teams []*TournamentLeagueTableTeam `xml:"Teams>Team"`
}

// TournamentLeagueTableTeam is a single team's row in a tournament league
// table.
type TournamentLeagueTableTeam struct {
	TeamID   id.Team `xml:"TeamID"`
	TeamName string  `xml:"TeamName"`

	Position uint `xml:"Position"`

	// Whether the team has moved up or down from the previous match
	// round.
	PositionChange int `xml:"PositionChange"`

	Matches      uint `xml:"Matches"`
	GoalsFor     uint `xml:"GoalsFor"`
	GoalsAgainst uint `xml:"GoalsAgainst"`
	Points       uint `xml:"Points"`
	Won          uint `xml:"Won"`
	Draws        uint `xml:"Draws"`
	Lost         uint `xml:"Lost"`
}
