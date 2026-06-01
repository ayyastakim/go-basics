package main

import "fmt"

type Filter func(string) string	// memudahkan pembuatan tipe data fungsi sebagai parameter

// fungsi sebagai parameter
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("hello", filter(name))
}


func spamFilter(name string) string {
	if name == "bangsat" {
		return "..."
	} else	{
		return name
	}
}

// func sayHello(name string) {
// 	nameFilter := name

// 	if nameFilter == "bangsat" {
// 		nameFilter = "..."
// 	}

// 	fmt.Println("hello", nameFilter)
// }

func main() {
	sayHelloWithFilter("john", spamFilter)
	sayHelloWithFilter("bangsat", spamFilter)

	// sayHello("taylor")
	// sayHello("bangsat")
}