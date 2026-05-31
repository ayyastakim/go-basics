package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke-", counter)
		counter++
	}

	// for loop conventional:
	// init statement: kode yang pertama kali dijalankan sebelum kondisi (inisialisai)
	// condition: mengontrol berapa banyak perulangan dilakukan (boolean operation)
	// post statement: dieksekusi setiap satu kali perulangan dijalankan (incremental/decremental)
	for i := 1; i <= 10; i++ {
		fmt.Println("looping", i)
	}

	// loop mennampilkan data slice
	 slice := []string {"john", "jean", "taylor", "jack", "adam"}

	 for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	 }

	 // for range : cocok digunakan untuk array, slice, map
	 for i, value := range slice {
		fmt.Println("index", i, "=", value)
	 }

	 // jika tidak menggunakan variable index/key, gunakan underscore (_)
	 for _, value := range slice {
		fmt.Println(value)
	 }

	 // for range untuk map
	 person := make(map[string]string)

	 person["name"] = "taylor"
	 person["gender"] = "female"
	 person["country"] = "united states"
	 person["title"] = "singer"
	 person["hobby"] = "cooking"

	 for key, value := range person {
		fmt.Println(key, "=", value)
	 }
}