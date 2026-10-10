package main

import "fmt"

func main() {
	fmt.Println("Masukkan kalkulator suhu yang ingin dipakai")
	fmt.Println("1. Celsius ke Fahrenheit")
	fmt.Println("2. Celsius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celsius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celsius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Println("Masukan pilihan anda : ")

	var pilihan int
	fmt.Scanf("%d", &pilihan)
	for pilihan < 1 || pilihan > 6 {
		fmt.Println("Pilihan tidak valid. Silakan masukkan pilihan antara 1 hingga 6.")
		fmt.Scanf("%d", &pilihan)
	}

	var suhu float64
	fmt.Println("Masukkan suhu awal : ")
	fmt.Scanf("%f", &suhu)

	var suhuAkhir float64
	if pilihan == 1 {
		suhuAkhir = CelciusToFahrenheit(suhu)
	} else if pilihan == 2 {
		suhuAkhir = CelciusToKelvin(suhu)
	} else if pilihan == 3 {
		suhuAkhir = FahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		suhuAkhir = FahrenheitToKelvin(suhu)
	} else if pilihan == 5 {
		suhuAkhir = KelvinToCelcius(suhu)
	} else if pilihan == 6 {
		suhuAkhir = KelvinToFahrenheit(suhu)
	}
	fmt.Printf("Suhu akhir: %.2f\n", suhuAkhir)
}

// input = 100 -> output = 212.00
func CelciusToFahrenheit(suhuCelcius float64) float64 {
	return (suhuCelcius * (9.0 / 5.0)) + 32
}

// input = 100 -> output = 373.15
func CelciusToKelvin(suhuCelcius float64) float64 {
	return suhuCelcius + 273.15
}

// rumus
// konversi suhu dari Fahrenheit ke Celcius:
// c = (F - 32) * 5 / 9

// input = 212 -> output = 100.00
func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit - 32) * (5.0 / 9.0)
}

// konversi suhu fahrenheit ke Kelvin:
// k = (F + 459.67) * 5 / 9

// input = 212 -> output = 373.15
func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit + 459.67) * (5.0 / 9.0)
}

// konversi suhu dari Kelvin ke Celcius:
// c = K - 273.15

// input = 373.15 -> output = 100.00
func KelvinToCelcius(suhuKelvin float64) float64 {
	return suhuKelvin - 273.15
}

// konversi suhu kelvin ke Fahrenheit:
// F = (K * 9 / 5) - 459.67

// input = 373.15 -> output = 212.00
func KelvinToFahrenheit(suhuKelvin float64) float64 {
	return (suhuKelvin * (9.0 / 5.0)) - 459.67
}

// karena ketika kita tidak memberikan 0 makan akan dianggap int, maka kita harus menambahkan .0 agar dianggap float
