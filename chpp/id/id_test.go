package id

import (
	"fmt"
	"testing"
)

func TestStringers(t *testing.T) {
	tests := []struct {
		name string
		val  fmt.Stringer
		want string
	}{
		{"AchievementID", AchievementID(123), "123"},
		{"Alliance", Alliance(123), "123"},
		{"Arena", Arena(123), "123"},
		{"Country", Country(123), "123"},
		{"Cup", Cup(123), "123"},
		{"CupLeagueLevel", CupLeagueLevel(123), "123"},
		{"Fanclub", Fanclub(123), "123"},
		{"FriendlyMatch", FriendlyMatch(123), "123"},
		{"Ladder", Ladder(123), "123"},
		{"Language", Language(123), "123"},
		{"League", League(123), "123"},
		{"Match", Match(123), "123"},
		{"Match negative sentinel", Match(-1), "-1"},
		{"NationalTeam", NationalTeam(123), "123"},
		{"Player", Player(123), "123"},
		{"Referee", Referee(123), "123"},
		{"Region", Region(123), "123"},
		{"Role", Role(123), "123"},
		{"Series", Series(123), "123"},
		{"Staff", Staff(123), "123"},
		{"Team", Team(123), "123"},
		{"Tournament", Tournament(123), "123"},
		{"Trainer", Trainer(123), "123"},
		{"Transfer", Transfer(123), "123"},
		{"User", User(123), "123"},
		{"YouthLeague", YouthLeague(123), "123"},
		{"YouthPlayer", YouthPlayer(123), "123"},
		{"YouthTeam", YouthTeam(123), "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCountry_Flag(t *testing.T) {
	if got := Country(4).Flag(); got != "🇮🇹" {
		t.Errorf("Country(4).Flag() = %q, want 🇮🇹", got)
	}
	if got := Country(999999).Flag(); got != "unknown" {
		t.Errorf("Country(999999).Flag() = %q, want %q", got, "unknown")
	}
}

func TestLeague_Flag(t *testing.T) {
	if got := League(4).Flag(); got != "🇮🇹" {
		t.Errorf("League(4).Flag() = %q, want 🇮🇹", got)
	}
	if got := League(999999).Flag(); got != "unknown" {
		t.Errorf("League(999999).Flag() = %q, want %q", got, "unknown")
	}
}

func TestCupLevel_String(t *testing.T) {
	tests := []struct {
		l    CupLevel
		want string
	}{
		{CupLevelNational, "national"},
		{CupLevelChallenger, "challenger"},
		{CupLevelConsolation, "consolation"},
		{CupLevel(99), "unknown"},
	}

	for _, tt := range tests {
		if got := tt.l.String(); got != tt.want {
			t.Errorf("CupLevel(%d).String() = %q, want %q", tt.l, got, tt.want)
		}
	}
}

func TestCupLevelIndex_String(t *testing.T) {
	tests := []struct {
		l    CupLevelIndex
		want string
	}{
		{CupLevelIndexEmerald, "emerald"},
		{CupLevelIndexRuby, "ruby"},
		{CupLevelIndexSapphire, "sapphire"},
		{CupLevelIndex(99), "unknown"},
	}

	for _, tt := range tests {
		if got := tt.l.String(); got != tt.want {
			t.Errorf("CupLevelIndex(%d).String() = %q, want %q", tt.l, got, tt.want)
		}
	}
}

func TestCupLeagueLevel_Type(t *testing.T) {
	tests := []struct {
		l    CupLeagueLevel
		want string
	}{
		{1, "national"},
		{6, "national"},
		{7, "divisional"},
		{9, "divisional"},
		{0, "unknown"},
		{10, "unknown"},
	}

	for _, tt := range tests {
		if got := tt.l.Type(); got != tt.want {
			t.Errorf("CupLeagueLevel(%d).Type() = %q, want %q", tt.l, got, tt.want)
		}
	}
}
