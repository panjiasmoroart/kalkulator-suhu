package main

import "fmt"

//  ketika kita tidak memberikan 0 makan akan dianggap int,
//  maka kita harus menambahkan .0 agar dianggap float pada saat melakukan perhitungan

type celsius struct {
	suhu float64
}

type fahrenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

// celsius to celsius
func (c celsius) toCelcius() float64 {
	return c.suhu
}

// input = 100 -> output = 212.00
func (c celsius) toFahrenheit() float64 {
	return (c.suhu * (9.0 / 5.0)) + 32
}

// input = 100 -> output = 373.15
func (c celsius) toKelvin() float64 {
	return c.suhu + 273.15
}

// fahrenheit to fahrenheit
func (f fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

// input = 212 -> output = 100.00
func (f fahrenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

// input = 212 -> output = 373.15
func (f fahrenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

// kelvin to kelvin
func (k kelvin) toKelvin() float64 {
	return k.suhu
}

// input = 373.15 -> output = 100.00
func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

// input = 373.15 -> output = 212.00
func (k kelvin) toFahrenheit() float64 {
	return (k.suhu * (9.0 / 5.0)) - 459.67
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Pilihan suhu awal : ")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukan suhu awal yang diinginkan : ")

	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 3 {
		fmt.Println("Suhu awal tidak valid, masukkan suhu awal yang diinginkan : ")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("Pilihan suhu akhir : ")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("Masukan suhu akhir yang diinginkan : ")

	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)
	for suhuAkhir < 1 || suhuAkhir > 3 {
		fmt.Println("Suhu akhir tidak valid, masukkan suhu akhir yang diinginkan : ")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Println("Masukkkan suhu : ")
	fmt.Scanf("%f", &suhu)

	var interfaceSuhu hitungSuhu
	switch suhuAwal {
	case 1:
		interfaceSuhu = celsius{suhu}
	case 2:
		interfaceSuhu = fahrenheit{suhu}
	case 3:
		interfaceSuhu = kelvin{suhu}
	}

	var suhuFinal float64
	switch suhuAkhir {
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
	case 2:
		suhuFinal = interfaceSuhu.toFahrenheit()
	case 3:
		suhuFinal = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Suhu akhir yang di dapa adalah : %.2f\n", suhuFinal)

}
