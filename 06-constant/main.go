package main

import "fmt"

func main() {
	// constant: value tidak dapat diubah/re-assign
	const name string = "john doe"

	fmt.Println(name)

	// error: tidak dapat diubah nilainya
	// name = "jean doe"

	// error: harus diinisialisasi nilainya secara langsung
	// const age int
	// age = 100

	// constant: tidak wajib mendeklarasikan tipe data
	const country = "singapore"

	fmt.Println(country)

	// multiple constant
	const (
		role = "backend developer"
		company = "pt. freport indonesia"
	)

	fmt.Println(role)
	fmt.Println(company)

}