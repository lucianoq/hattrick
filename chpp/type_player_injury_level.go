package chpp

// PlayerInjuryLevel ...
type PlayerInjuryLevel string

// List of PlayerInjuryLevel constants.
const (
	InjuryHealthy      PlayerInjuryLevel = "-1"
	InjuryBruised      PlayerInjuryLevel = "0"
	InjuryNotAvailable PlayerInjuryLevel = "NOT AVAILABLE"
	InjuryPlus1        PlayerInjuryLevel = "+1"
	InjuryPlus2        PlayerInjuryLevel = "+2"
	InjuryPlus3        PlayerInjuryLevel = "+3"
	InjuryPlus4        PlayerInjuryLevel = "+4"
	InjuryPlus5        PlayerInjuryLevel = "+5"
	InjuryPlus6        PlayerInjuryLevel = "+6"
	InjuryPlus7        PlayerInjuryLevel = "+7"
	InjuryPlus8        PlayerInjuryLevel = "+8"
	InjuryPlus9        PlayerInjuryLevel = "+9"
	InjuryPlus10       PlayerInjuryLevel = "+10"
)

// String returns a string representation of the type.
func (i PlayerInjuryLevel) String() string {
	switch i {
	case InjuryHealthy:
		return ""
	case InjuryBruised:
		return "bruised"
	case InjuryNotAvailable:
		return "NA"
	case InjuryPlus1:
		return "+1"
	case InjuryPlus2:
		return "+2"
	case InjuryPlus3:
		return "+3"
	case InjuryPlus4:
		return "+4"
	case InjuryPlus5:
		return "+5"
	case InjuryPlus6:
		return "+6"
	case InjuryPlus7:
		return "+7"
	case InjuryPlus8:
		return "+8"
	case InjuryPlus9:
		return "+9"
	case InjuryPlus10:
		return "+10"
	default:
		return ""
	}
}

// Short ...
func (i PlayerInjuryLevel) Short() string {
	switch i {
	case InjuryHealthy:
		return ""
	case InjuryBruised:
		return "🤕"
	case InjuryNotAvailable:
		return "NA"
	case InjuryPlus1:
		return "🏥1"
	case InjuryPlus2:
		return "🏥2"
	case InjuryPlus3:
		return "🏥3"
	case InjuryPlus4:
		return "🏥4"
	case InjuryPlus5:
		return "🏥5"
	case InjuryPlus6:
		return "🏥6"
	case InjuryPlus7:
		return "🏥7"
	case InjuryPlus8:
		return "🏥8"
	case InjuryPlus9:
		return "🏥9"
	case InjuryPlus10:
		return "🏥10"
	default:
		return ""
	}
}
