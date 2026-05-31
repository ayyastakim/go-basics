package main

import "fmt"

func main() {
	// membuat type declarations, mirip alias
	type noKTP string
	type isMarried bool
	
	var myID noKTP = "1234567890"
	var marriedStatus isMarried = false

	fmt.Println(myID)
	fmt.Println(marriedStatus)
}