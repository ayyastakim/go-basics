package main

import "fmt"

func registerUser(name string, blacklist func(string) bool) {
	if blacklist(name) {
		fmt.Println("you're blocked", name)
	} else {
		fmt.Println("you're verified", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {

	// anonymous function -> fungsi diinisialisasi sebagai value/nilai variable
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("john", blacklist)
	registerUser("admin", blacklist)

	//anonymous function -> fungsi diinisialisasi sebagai argument
	registerUser("taylor", func(name string) bool {
		return name == "root"
	})

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}