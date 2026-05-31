package main

import "fmt"

func main() {
	var names [3]string

	// built-in function: memasukan data array by index
	names[0] = "john"
	names[1] = "jean"
	names[2] = "taylor"

	// built-in function: mengakses/mendapatkan data array by index
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array secara langsung
	var results = [5]int {1, 2, 3, 4, 5}

	fmt.Println(results)

	// built-in function: mengubah data array
	results[0] = 0
	fmt.Println(results)

	// built-in function: mengakses panjang/ukuran data array
	fmt.Println(len(names))
	fmt.Println(len(results))


}