package main

import "fmt"

func main() {
	var name = "brandon the builder"

	if name == "john" {
		fmt.Println("hello john!")
	} else if name == "taylor" {
		fmt.Println("hello taylor!")
	} else {
		fmt.Println("hello stranger!")
	}

	// short statement : pembuatan variable di dalam if, best practice: jika variable tesebut hanya digunakan sebagai parameter di kondisi if tersebut!

	if length := len(name); length > 5 {
		fmt.Println("your name is too long!")
	}
}