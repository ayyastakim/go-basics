package main

import "fmt"

func main() {
	var name = "taylor swift"

	switch name {
	case "john":
		fmt.Println("hello john!")
	case "taylor":
		fmt.Println("hello taylor!")
	default:
		fmt.Println("hello stranger!")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("your name is too long!")
	case false:
		fmt.Println("your name is correct!")
	}

	// simple expression : kondisi / variable dicek di dalam case
	var length = len(name)

	switch {
	case length > 10:
		fmt.Println("nama anda terlalu panjang!")
	case length > 5 && length < 10:
		fmt.Println("nama anda lumayan panjang!")
	default:
		fmt.Println("nama anda sudah benar!")
	}
}