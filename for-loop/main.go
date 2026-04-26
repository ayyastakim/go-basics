package main

import "fmt"

func main() {

	// for loop conventional
	for counter := 1; counter <= 10; counter++{
		fmt.Println("loop ke-", counter)
	}

	// fo range -> diapakai untuk array, slice, map
	
	numbers := [5]int {1, 2, 3, 4, 5}

	// ambil data array dengan loop conventional

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// ambil data array dengan for range

	for i, number := range numbers {
		fmt.Println("loop ke", i, "=", number)
	}

	// ambil data slice dengan for range

	slice := make([]string, 3, 5)

	slice[0] = "viserys"
	slice[1] = "daenarys"
	slice[2] = "baelor"

	for i, value := range slice {
		fmt.Println("indeks-", i, "=", value)
	}

	// cara untuk tidak menggunakan variabel index
	for _, value := range slice {
		fmt.Println(value)
	}


	// ambil data map dengan for range

	names := make(map[int]string)

	names[1] = "rob stark"
	names[2] = "john snow"
	names[3] = "sansa stark"
	names[4] = "arya stark"
	names[5] = "brandon stark"
	names[6] = "rickon stark"

	for key, value := range names {
		fmt.Println("key: ", key, "|", "value: ", value)
	}

	// akses value tanpa tanpa key
	for _, value := range names {
		fmt.Println(value)
	}

}