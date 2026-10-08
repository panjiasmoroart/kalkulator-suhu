package main

import "fmt"

func main() {
	var suhu float64
	fmt.Print("Masukkan suhu: ")
	fmt.Scanf("%f", &suhu)
	fmt.Printf("Suhu yang di input adalah : %f\n", suhu)

	// uint64 bisa positif dan negatif
	var uang uint64
	fmt.Print("Masukkan uang: ")
	fmt.Scanf("%d", &uang)
	fmt.Printf("Uang yang di input adalah : %d\n", uang)

	var nama string
	fmt.Println("Masukan nama :")
	fmt.Scanf("%s", &nama)
	fmt.Printf("Nama yang di input adalah : %s\n", nama)
}
