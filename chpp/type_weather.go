package chpp

// Weather ...
type Weather uint

// List of Weather constants.
const (
	WeatherRain            Weather = 0
	WeatherOvercast        Weather = 1
	WeatherPartiallyCloudy Weather = 2
	WeatherSunny           Weather = 3
)

// String returns a string representation of the type.
func (w Weather) String() string {
	switch w {
	case WeatherRain:
		return "rain"
	case WeatherOvercast:
		return "overcast"
	case WeatherPartiallyCloudy:
		return "partially_cloudy"
	case WeatherSunny:
		return "sunny"
	default:
		return "unknown"
	}
}
