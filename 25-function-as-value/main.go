package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name + "!"
}

func main() {

	// fungsi diset sebagai value/nilai variabel
	var sayGoodBye func(name string) string = getGoodBye	// bentuk simple -> sayGoodBye := getGoodBye
	
	// value := sayGoodBye("john")
	// fmt.Println(value)

	fmt.Println(sayGoodBye("john"))
}