package main

import "fmt"

func main() {
	var name1 = "john"
	var name2 = "jean"

	var results = name1 == name2

	fmt.Println(results)

	var number1 = 4
	var number2 = 7

	fmt.Println(number1 < number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 <= number2)
	fmt.Println(number1 >= number2)
	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
}