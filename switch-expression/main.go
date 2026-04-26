package main

import "fmt"

func main() {
	var name string = "joko"

	switch name {
	case "taylor":
		fmt.Println("hallo, taylor!")
	case "jean":
		fmt.Println("hallo, jean!")
	default:
		fmt.Println("hallo, user!")
	}

	// switch short expression

	// var length = len(name)

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama anda terlalu panjang!")
	case false:
		fmt.Println("nama anda sudah tepat!")
	}

	// switch tanpa condition

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("nama anda terlalu panjang!")
	case length > 5 && length < 10:
		fmt.Println("nama anda lumayan panjang!")
	default:
		fmt.Println("nama anda sudah tepat!")
	}
}