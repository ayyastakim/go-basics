package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "taylor",
		"country": "united state",
	}

	fmt.Println(person)

	// mengambil data map dengan key
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	// mengubah/menambah data map (dinamis)
	person["tittle"] = "programmer"

	fmt.Println(person)

	// built-in function : len() -> mendapatkan jumlah data map
	fmt.Println(len(person))
	
	// built-in function : make(map[TypeKey]TypeValue) -> membuat map baru
	book := make(map[string]string)
	
	book["title"] = "go-lang basics"
	book["author"] = "john doe"
	book["penerbit"] = "umbrella company"
	
	fmt.Println(book)
	fmt.Println(len(book))
	
	// built-in function : delete(map, key) -> menghapus data di map dengan key
	delete(book, "penerbit")
	
	fmt.Println(book)
	fmt.Println(len(book))

}