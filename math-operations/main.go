package main

import "fmt"

func main() {
	// operasi aritmatika
	var result = 10 + 10

	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b

	fmt.Println(c)

	// augmented assigments
	var d = 10
	d += 10	// d = d + 10

	fmt.Println(d)

	// unary operator
	var i = 10

	i++	// increment
	fmt.Println(i)
	i--	// decrement
	fmt.Println(i)

	var negative = -12
	var psoitive = +12	// default, tidak perlu tanda + untuk bilangan positif

	fmt.Println(negative)
	fmt.Println(psoitive)

}