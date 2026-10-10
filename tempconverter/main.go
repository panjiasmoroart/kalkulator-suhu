package main

func main() {}

func CelciusToFahrenheit(suhuCelcius float64) float64 {
	return (suhuCelcius * 9 / 5) + 32
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	return suhuCelcius + 273.15
}

// rumus
// konversi suhu dari Fahrenheit ke Celcius:
// c = (F - 32) * 5 / 9

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit - 32) * 5 / 9
}

// konversi suhu fahrenheit ke Kelvin:
// k = (F + 459.67) * 5 / 9

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit + 459.67) * 5 / 9
}

// konversi suhu dari Kelvin ke Celcius:
// c = K - 273.15

func KelvinToCelcius(suhuKelvin float64) float64 {
	return suhuKelvin - 273.15
}

// konversi suhu kelvin ke Fahrenheit:
// F = (K * 9 / 5) - 459.67

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	return (suhuKelvin * 9 / 5) - 459.67
}
