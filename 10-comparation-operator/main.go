package main

import "fmt"

func main() {
	
	var name1 = "john"
	var name2 = "jean"

	var result bool = name1 == name2

	fmt.Println(result)

	var number1 = 2
	var number2 = 5

	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
	fmt.Println(number1 >= number2)
	fmt.Println(number1 <= number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 < number2)

}