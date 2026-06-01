package main

import "fmt"

// named return values
func getFullName() (firstName, middleName, lastName string) {	// cara singkat: ketiga return value bertipe string
	firstName = "jack"
	middleName = "kahuna"
	lastName = "laguna"

	// return firstName, middleName, lastName

	// bisa melakukan return tanpa mencantumkan variable
	return
}

func main() {
	a, b, c := getFullName()
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}