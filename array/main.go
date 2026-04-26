package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "john"
	names[1] = "taylor"
	names[2] = "luke"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var numbers = [3]int {
		70, 
		75, 
		80,
	}

	fmt.Println(numbers)

	// built-in function -> mendapatkan panjang array
	fmt.Println(len(names))
	fmt.Println(len(numbers))
}