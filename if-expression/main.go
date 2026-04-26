package main

import "fmt"

func main() {
	var name string = "taylor"

	if name == "taylor" {
		fmt.Println("hallo, taylor!")
	} else if name == "jean" {
		fmt.Println("hallo, jean!")
	} else if name == "sansa" {
		fmt.Println("hallo, sansa!")
	} else {
		fmt.Println("hallo, user!")
	}

	// if short statement
	
	// var length = len(name)

	if length := len(name); length > 5 {
		fmt.Println("nama anda terlalu panjang!")
	} else {
		fmt.Println("nama anda sudah tepat!")
	}
}