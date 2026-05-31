package main

import "fmt"

func main() {
	var nilai32 int32 = 1000

	//konversi ke int64
	var nilai64 int64 = int64(nilai32)

	//konversi ke int8
	var nilai8 int8 = int8(nilai32) // nilainya berubah karena: overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversi string
	var name = "john doe"

	//mengambil karakter pertama dari string
	var j = name[0]

	// konversi tipe data j (uint8/byte) ke string
	var jString string =string(j)

	fmt.Println(name)
	fmt.Println(j)
	fmt.Println(jString)
}