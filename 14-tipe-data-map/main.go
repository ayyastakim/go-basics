package main

import "fmt"

func main() {
	
	var person = map[string]string {
		"name" : "taylor",
		"country" : "singapore",
	}

	fmt.Println(person)

	// built-in function: mapVariableName[key] -> mengakses data map
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	// built-in function: mapVariableName[key] = value -> menambahkan data ke map
	person["company"] = "amazon"

	fmt.Println(person)

	// built-in function: len(mapVariableName) -> mengakses ukuran data map
	fmt.Println(len(person))

	// built-in function: make(map[keyTipe]valueType) -> membuat map baru
	books := make(map[string]string)

	books["title"] = "learning basics go languange"
	books["author"] = "john doe"
	books["random"] = "ups!!"

	fmt.Println(books)

	// built-in function: delete(mapVariableName, key) -> menghapus data map dengan key
	delete(books, "random")
	fmt.Println(books)

}