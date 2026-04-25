package main

import "fmt"

func main ()  {
	// deklarasi variable
	var nama string

	nama = "lorem ipsum"
	fmt.Println(nama)

	nama = "dolor sit"
	fmt.Println(nama)

	// tanpa eksplisit tipe data, langsung initialize
	var anotherNama = "lorem ipsum"
	fmt.Println(anotherNama)

	var age = 30	// int default minimal int32, sesuai sistem operasi yang digunakan (32bit/64bit)
	fmt.Println(age)

	// var anotherAge int8 = 30	// bisa juga ekspilisit tipe datanya ketika initialize
	// fmt.Println(anotherAge)

	// tanpa keyword var, gunakan := untuk initialize
	country := "indonesia"

	fmt.Println(country)

	// multiple variable declaration
	var (
		firstName = "lorem"
		lastName = "ipsum"
	)

	fmt.Println(firstName, lastName)
}

// variable adalah tempat untuk menyimpan data, variable memiliki nama dan tipe data. Di golang, variable harus dideklarasikan sebelum digunakan, dan dapat diubah nilainya setelah dideklarasikan. Variable juga memiliki scope, dimana variable hanya dapat diakses dalam blok kode tertentu. Variable dapat dideklarasikan dengan menggunakan keyword var atau dengan menggunakan := untuk langsung menginisialisasi nilai.
// tipe data yang umum digunakan untuk variable di golang antara lain string, int, float64, bool, dan lain-lain. Variable juga dapat dideklarasikan dengan tipe data yang lebih kompleks seperti array, slice, map, struct, dan lain-lain.
