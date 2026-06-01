package main

import "fmt"

// function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

func main() {
	firstName := "taylor"

	sayHelloTo(firstName, "swift")
}