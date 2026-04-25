package main

import "fmt"

func main() {
	fmt.Println("lorem")
	fmt.Println("lorem ipsum")
	fmt.Println("lorem ipsum dolor")

	fmt.Println(len("lorem"))
	fmt.Println("lorem ipsum"[0])
	fmt.Println("lorem ipsum dolor"[1])
}

// string adalah tipe data untuk menyimpan karakter atau teks, string di golang immutable (tidak bisa diubah setelah dibuat) dan menggunakan UTF-8 encoding, sehingga bisa menyimpan karakter dari berbagai bahasa. String dapat diakses menggunakan indeks, dimana indeks dimulai dari 0.
// string juga memiliki banyak fungsi built-in yang dapat digunakan untuk memanipulasi string, seperti len() untuk mendapatkan panjang string, dan operator + untuk menggabungkan string.