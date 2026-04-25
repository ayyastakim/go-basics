package main

import "fmt"

func main() {
	// dengan keyword tipe data
	const name string = "john"
	// tanpa keyword tipe data
	const age = 30
	const country = "netherland"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// multiple constant variable
	const (
		anotherName = "jean"
		anotherAge = 25
		anotherCountry = "united state"
	)

	fmt.Println(anotherName)
	fmt.Println(anotherAge)
	fmt.Println(anotherCountry)
}

// constant variable adalah variable yang nilainya tidak dapat diubah setelah dideklarasikan
// constant variable harus diinisialisasi pada saat deklarasi
// constant variable dapat dideklarasikan dengan atau tanpa tipe data
// constant variable dapat dideklarasikan dengan keyword const
// constant variable dapat dideklarasikan dengan multiple constant variable
// constant variable dapat digunakan untuk menyimpan nilai yang tidak berubah seperti nama, umur, negara, dll
// constant variable dapat digunakan untuk meningkatkan readability dan maintainability kode program
// constant variable dapat digunakan untuk menghindari magic number dalam kode program
// constant variable dapat digunakan untuk menghindari hardcode dalam kode program
// constant variable dapat digunakan untuk meningkatkan performa kode program karena nilai constant variable disimpan di dalam memory yang lebih cepat diakses daripada variable biasa