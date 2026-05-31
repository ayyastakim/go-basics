package main

import "fmt"

func main() {
	// operasi artimatik
	var number1 int8 = 3
	var number2 int8 = 5

	// operasi penjumlahan
	fmt.Println(number1 + number2)
	// operasi pengurangan
	fmt.Println(number1 - number2)
	// operasi perkalian
	fmt.Println(number1 * number2)
	// operasi pembagian
	fmt.Println(number1 / number2)
	// operasi sisa hasil bagi (modulus)
	fmt.Println(number1 % number2)


	// augmented assignment
	var number3 = 10

	number3 += 10	// number3 = number3 + 10
	fmt.Println(number3)
	number3 -= 10	// number3 = number3 - 10
	fmt.Println(number3)

	// unary operator

	number4 := 10

	number4++ // increment
	fmt.Println(number4)

	number4--	// decrement
	fmt.Println(number4)

	// nilai negatif dan positif
	var negative = -2
	fmt.Println(negative)
	var positive = +2	// default nilai positive tidak perlu tanda '+'
	fmt.Println(positive)

	// negasi (!) atau kebalikan (boolean type)
	var isMarried = !true
	fmt.Println(isMarried)

}