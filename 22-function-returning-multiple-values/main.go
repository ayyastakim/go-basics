package main

import "fmt"

// returning multiple values in function
func getFullName() (string, string) {
	return "taylor", "swift"
}

func main() {

	fmt.Println(getFullName())

	// function call dengan memasukan argument dalam variable pada fungsi dengan multiple return values
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// ignore argmument pada fungsi dengan multiple return values
	anotherFirstName, _ := getFullName()
	fmt.Println(anotherFirstName)
}