package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPJohn NoKTP = "1234567890"
	var marriedStatus Married = true

	fmt.Println(noKTPJohn)
	fmt.Println(marriedStatus)
}

// type declaration adalah cara untuk membuat tipe data baru dengan nama yang berbeda dari tipe data yang sudah ada. Dalam contoh di atas, kita membuat tipe data baru bernama NoKTP yang merupakan alias dari tipe string, dan tipe data Married yang merupakan alias dari tipe bool. Dengan menggunakan type declaration, kita dapat memberikan nama yang lebih spesifik dan mudah dipahami untuk tipe data yang kita gunakan dalam program kita.
// Type declaration juga dapat digunakan untuk membuat tipe data yang lebih kompleks, seperti struct atau interface. Dengan menggunakan type declaration, kita dapat membuat tipe data yang lebih terstruktur dan mudah digunakan dalam program kita.