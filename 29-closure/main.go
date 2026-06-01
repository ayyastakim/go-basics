package main

import "fmt"

// closure: kemampuan suatu inner function untuk berinteraksi dengan data-data pada block outer function
func main() {
	counter := 0	// dapat diakses pada inner main function
	name := "taylor"	// dapat diakses pada pada inner main function

	increment := func()  {	// fungsi ini dapat mengakses data di atasnya pada block outer function
		name := "john"	// hanya dapat diakses dalam scope increment function (inner function)
		counter++

		fmt.Println(name)
	}

	fmt.Println(name)

	increment()
	fmt.Println(counter)
}