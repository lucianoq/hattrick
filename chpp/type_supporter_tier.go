package chpp

// SupporterTier is a manager's paid CHPP/Hattrick Supporter subscription
// level (none, silver, gold, platinum, diamond).
type SupporterTier string

// List of SupporterTier constants.
const (
	SupporterTierEmpty    SupporterTier = ""
	SupporterTierNone     SupporterTier = "none"
	SupporterTierSilver   SupporterTier = "silver"
	SupporterTierGold     SupporterTier = "gold"
	SupporterTierPlatinum SupporterTier = "platinum"
	SupporterTierDiamond  SupporterTier = "diamond"
)

// String returns a string representation of the type.
func (t SupporterTier) String() string {
	switch t {

	case SupporterTierEmpty, SupporterTierNone:
		return "none"

	case SupporterTierSilver:
		return "silver"

	case SupporterTierGold:
		return "gold"

	case SupporterTierPlatinum:
		return "platinum"

	case SupporterTierDiamond:
		return "diamond"
	}

	return "unknown"
}
