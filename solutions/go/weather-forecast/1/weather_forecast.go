// Package weather provides tools to forcast weather.
package weather

// CurrentCondition represent the actual weather at location x.
var CurrentCondition string
// CurrentLocation represent a city of Goblinocus.
var CurrentLocation string

// Forecast function show the current weather of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
