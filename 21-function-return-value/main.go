package main

import "fmt"

// function dengan return value (tipe data string)
func getHello(name string) string {
	return "hello " + name
}

func main() {

	// inisialisasi function call ke dalam variable
	value := getHello("john doe")

	fmt.Println(value)

	// function call secara langsung dengan fungsi Println()
	fmt.Println(getHello("jean doe"))
}