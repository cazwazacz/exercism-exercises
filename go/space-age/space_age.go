package space

const earthOrbitalPeriodInSeconds = 31557600

// Age gives you your planetary ages
func Age(seconds float64, planet Planet) float64 {
	return seconds / (earthOrbitalPeriodInSeconds * planet.orbitalPeriodInEarthYears)
}

// Planet holds information needed to calculate age in years
type Planet struct {
	name                      string
	orbitalPeriodInEarthYears float64
}
