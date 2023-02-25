package chpp

// SupporterTier ...
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
