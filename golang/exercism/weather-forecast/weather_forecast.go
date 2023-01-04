// Package weather - forecast the weather based on city and condition parameters.
package weather


// CurrentCondition - should be the weather condition.
var CurrentCondition string

// CurrentLocation - should be the city.
var CurrentLocation string


// Forecasts the weather based on city and condition inputs.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
