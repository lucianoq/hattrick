package chpp

import "testing"

func TestAchievementCategory_String(t *testing.T) {
	tests := []struct {
		c    AchievementCategory
		want string
	}{
		{AchievementCategoryRanking, "ranking"},
		{AchievementCategoryTeam, "team"},
		{AchievementCategoryMatches, "matches"},
		{AchievementCategoryManager, "manager"},
		{AchievementCategorySpecialAwards, "special_awards"},
		{AchievementCategorySupporter, "supporter"},
		{AchievementCategory(99), "unknown achievement"},
	}
	for _, tt := range tests {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("AchievementCategory(%d).String() = %q, want %q", tt.c, got, tt.want)
		}
	}
}

func TestBookmarkType_String(t *testing.T) {
	if got := BookmarkTypePlayers.String(); got != "2" {
		t.Errorf("BookmarkTypePlayers.String() = %q, want %q", got, "2")
	}
}

func TestFanMood_String(t *testing.T) {
	tests := []struct {
		s    FanMood
		want string
	}{
		{FanMoodMurderous, "murderous"},
		{FanMoodFurious, "furious"},
		{FanMoodAngry, "angry"},
		{FanMoodIrritated, "irritated"},
		{FanMoodDisappointed, "disappointed"},
		{FanMoodCalm, "calm"},
		{FanMoodContent, "content"},
		{FanMoodSatisfied, "satisfied"},
		{FanMoodDelirious, "delirious"},
		{FanMoodHighOnLife, "high on life"},
		{FanMoodDancingInTheStreet, "dancing in the streets"},
		{FanMoodSendingLovePoemsToYou, "sending love poems to you"},
		{FanMood(99), ""},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("FanMood(%d).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestFriendlyType_String(t *testing.T) {
	if got := FriendlyTypeNationalTeamFriendly.String(); got != "12" {
		t.Errorf("FriendlyTypeNationalTeamFriendly.String() = %q, want %q", got, "12")
	}
}

func TestLeagueOfficeTypeID_String(t *testing.T) {
	if got := LeagueOfficeTypeU20Teams.String(); got != "4" {
		t.Errorf("LeagueOfficeTypeU20Teams.String() = %q, want %q", got, "4")
	}
}

func TestMatchPlace_String(t *testing.T) {
	if got := MatchPlaceAway.String(); got != "1" {
		t.Errorf("MatchPlaceAway.String() = %q, want %q", got, "1")
	}
}

func TestMatchRating_String(t *testing.T) {
	tests := []struct {
		m    MatchRating
		want string
	}{
		{MatchRatingVeryLowDisastrous, "VeryLowDisastrous"},
		{MatchRatingVeryHighExtraTerrestrial, "VeryHighExtraTerrestrial"},
		{MatchRatingVeryHighDivine, "VeryHighDivine"},
		{MatchRating(0), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.m.String(); got != tt.want {
			t.Errorf("MatchRating(%d).String() = %q, want %q", tt.m, got, tt.want)
		}
	}
}

func TestMatchRating_Value(t *testing.T) {
	if got := MatchRatingVeryHighDivine.Value(); got != 80 {
		t.Errorf("MatchRatingVeryHighDivine.Value() = %d, want 80", got)
	}
}

func TestMatchType_String(t *testing.T) {
	tests := []struct {
		m    MatchType
		want string
	}{
		{MatchTypeLeagueMatch, "league"},
		{MatchTypeYouthInternationalFriendlyMatchNormalRules, "youth_friendly_international_normal_rules"},
		{MatchTypeYouthInternationalFriendlyMatchCupRules, "youth_friendly_international_cup_rules"},
		{MatchType(0), ""},
	}
	for _, tt := range tests {
		if got := tt.m.String(); got != tt.want {
			t.Errorf("MatchType(%d).String() = %q, want %q", tt.m, got, tt.want)
		}
	}
}

func TestNationalTeamStaffType_String(t *testing.T) {
	tests := []struct {
		s    NationalTeamStaffType
		want string
	}{
		{NationalTeamCoach, "national team coach"},
		{AssistantCoach, "assistant coach"},
		{Scout, "scout"},
		{NationalTeamStaffType(99), ""},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("NationalTeamStaffType(%d).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestPlayerAggressiveness_String(t *testing.T) {
	tests := []struct {
		a    PlayerAggressiveness
		want string
	}{
		{PlayerAggressivenessTranquil, "tranquil"},
		{PlayerAggressivenessUnstable, "unstable"},
		{PlayerAggressiveness(99), ""},
	}
	for _, tt := range tests {
		if got := tt.a.String(); got != tt.want {
			t.Errorf("PlayerAggressiveness(%d).String() = %q, want %q", tt.a, got, tt.want)
		}
	}
}

func TestPlayerAgreeability_String(t *testing.T) {
	tests := []struct {
		a    PlayerAgreeability
		want string
	}{
		{PlayerAgreeabilityNastyFellow, "nasty fellow"},
		{PlayerAgreeabilityBelovedTeamMember, "beloved team member"},
		{PlayerAgreeability(99), ""},
	}
	for _, tt := range tests {
		if got := tt.a.String(); got != tt.want {
			t.Errorf("PlayerAgreeability(%d).String() = %q, want %q", tt.a, got, tt.want)
		}
	}
}

func TestPlayerHonesty_String(t *testing.T) {
	tests := []struct {
		h    PlayerHonesty
		want string
	}{
		{PlayerHonestyInfamous, "infamous"},
		{PlayerHonestySaintly, "saintly"},
		{PlayerHonesty(99), ""},
	}
	for _, tt := range tests {
		if got := tt.h.String(); got != tt.want {
			t.Errorf("PlayerHonesty(%d).String() = %q, want %q", tt.h, got, tt.want)
		}
	}
}

func TestPlayerInjuryLevel(t *testing.T) {
	tests := []struct {
		i          PlayerInjuryLevel
		wantString string
		wantShort  string
	}{
		{InjuryHealthy, "", ""},
		{InjuryBruised, "bruised", "🤕"},
		{PlayerInjuryLevel(3), "+3", "🏥3"},
	}
	for _, tt := range tests {
		if got := tt.i.String(); got != tt.wantString {
			t.Errorf("PlayerInjuryLevel(%d).String() = %q, want %q", tt.i, got, tt.wantString)
		}
		if got := tt.i.Short(); got != tt.wantShort {
			t.Errorf("PlayerInjuryLevel(%d).Short() = %q, want %q", tt.i, got, tt.wantShort)
		}
	}
}

func TestPositionChange_Icon(t *testing.T) {
	tests := []struct {
		p    PositionChange
		want string
	}{
		{NoChange, " "},
		{MovingUp, "^"},
		{MovingDown, "v"},
		{PositionChange(99), ""},
	}
	for _, tt := range tests {
		if got := tt.p.Icon(); got != tt.want {
			t.Errorf("PositionChange(%d).Icon() = %q, want %q", tt.p, got, tt.want)
		}
	}
}

func TestSkillLevel(t *testing.T) {
	if got := SkillLevelDivine.String(); got != "divine" {
		t.Errorf("SkillLevelDivine.String() = %q, want %q", got, "divine")
	}
	if got := SkillLevel(99).String(); got != "" {
		t.Errorf("SkillLevel(99).String() = %q, want empty", got)
	}
	if got := SkillLevelExcellent.Value(); got != 8 {
		t.Errorf("SkillLevelExcellent.Value() = %d, want 8", got)
	}
}

func TestSpecialtyID(t *testing.T) {
	tests := []struct {
		s             SpecialtyID
		wantString    string
		wantStringTag string
		wantEmoji     string
	}{
		{SpecialtyNoSpecialty, "", "", ""},
		{SpecialtyTechnical, "Technical", "[Technical]", "⚽"},
		{SpecialtyQuick, "Quick", "[Quick]", "⚡"},
		{SpecialtyResilient, "Resilient", "[Resilient]", ""},
		{SpecialtySupport, "Support", "[Supporter]", ""},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.wantString {
			t.Errorf("SpecialtyID(%d).String() = %q, want %q", tt.s, got, tt.wantString)
		}
		if got := tt.s.StringTag(); got != tt.wantStringTag {
			t.Errorf("SpecialtyID(%d).StringTag() = %q, want %q", tt.s, got, tt.wantStringTag)
		}
		if got := tt.s.Emoji(); got != tt.wantEmoji {
			t.Errorf("SpecialtyID(%d).Emoji() = %q, want %q", tt.s, got, tt.wantEmoji)
		}
	}
}

func TestSupporterTier_String(t *testing.T) {
	tests := []struct {
		s    SupporterTier
		want string
	}{
		{SupporterTierEmpty, "none"},
		{SupporterTierNone, "none"},
		{SupporterTierSilver, "silver"},
		{SupporterTierGold, "gold"},
		{SupporterTierPlatinum, "platinum"},
		{SupporterTierDiamond, "diamond"},
		{SupporterTier("garbage"), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("SupporterTier(%q).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestTrackingTypeID_String(t *testing.T) {
	if got := TrackingTypeHotListed.String(); got != "5" {
		t.Errorf("TrackingTypeHotListed.String() = %q, want %q", got, "5")
	}
}

func TestTrophyID_String(t *testing.T) {
	tests := []struct {
		tr   TrophyID
		want string
	}{
		{TrophyCupWinner, "cup_winner"},
		{TrophyWorldCupGold, "world_cup_gold"},
		{TrophyID(1), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.tr.String(); got != tt.want {
			t.Errorf("TrophyID(%d).String() = %q, want %q", tt.tr, got, tt.want)
		}
	}
}

func TestTSI_String(t *testing.T) {
	if got := TSI(1234567).String(); got != "1,234,567" {
		t.Errorf("TSI(1234567).String() = %q, want %q", got, "1,234,567")
	}
}

func TestWeather_String(t *testing.T) {
	tests := []struct {
		w    Weather
		want string
	}{
		{WeatherRain, "rain"},
		{WeatherOvercast, "overcast"},
		{WeatherPartiallyCloudy, "partially_cloudy"},
		{WeatherSunny, "sunny"},
		{Weather(99), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.w.String(); got != tt.want {
			t.Errorf("Weather(%d).String() = %q, want %q", tt.w, got, tt.want)
		}
	}
}
