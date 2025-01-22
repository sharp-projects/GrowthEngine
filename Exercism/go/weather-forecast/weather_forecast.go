// Package weather provides the current forecast for a specific location.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the city.
var CurrentLocation string

// Forecast returns the current weather condition for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
