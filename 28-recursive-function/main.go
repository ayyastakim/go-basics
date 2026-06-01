package main

import "fmt"

// membuat factorial dengan recursive
func getFactorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * getFactorialRecursive(number-1)
	}
}

// // membuat factorial dengan for loop
// func getFactorialLoop(number int) int {
// 	total := number
	
// 	for i := 1; i < number ; i++ {
// 		total *= (number-i)
// 	}
// 	return total
// }

func main() {
	value := getFactorialRecursive(5)
	fmt.Println(value)

	// fmt.Println(getFactorialLoop(5))

}