package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	// konversi ke int64
	var nilai64 int64 = int64(nilai32)
	// konversi ke int8
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi byte -> string
	var nama string = "john"
	var j byte = nama[0]
	var jString = string(j)

	fmt.Println(nama)
	fmt.Println(j)
	fmt.Println(jString)
}

// konversi tipe data adalah proses mengubah nilai dari satu tipe data ke tipe data lain.
// konversi tipe data dapat dilakukan secara eksplisit dengan menggunakan fungsi konversi seperti int(), float(), string(), dll.
// konversi tipe data dapat menyebabkan kehilangan data jika nilai yang dikonversi tidak sesuai dengan tipe data tujuan.
// konversi tipe data dapat dilakukan secara implisit oleh compiler jika nilai yang dikonversi sesuai dengan tipe data tujuan.