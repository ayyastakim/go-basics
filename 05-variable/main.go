package main

import (
	"fmt"
)

func main() {
	// deklarasi variable
	var name string

	name = "john doe"
	fmt.Println(name)

	// mengubah value
	name = "jean doe"
	fmt.Println(name)

	// inisialisasi variable (tidak wajib menggunakan keyword tipe data)
	var age = 20	// default data type: int (minimal int32)
	fmt.Println(age)

	var country = "indonesia" // default data type: string
	fmt.Println(country)

	// inisialisasi variable tanpa keyword 'var'
	isMarried := true		// gunakan ':=' ketika inisialisasi
	fmt.Println(isMarried)

	isMarried = false
	fmt.Println(isMarried)

	// multiple variabel
	var (
		role = "backend developer"
		company = "pt. freport indonesia"
	)

	fmt.Println(role)
	fmt.Println(company)
}