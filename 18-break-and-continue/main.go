package main

import "fmt"

func main() {
	// break: menghentikan perulangan
	for i := 0; i < 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println("loop", i)
	}

	// continue: menghentikan perulangan saat itu, dan beralih ke perulangan selanjutnya
	for i := 0; i < 10; i++ {

		if i % 2 == 0 {		// jika genap, skip!, lanjut peulangan berikut
			continue
		}

		fmt.Println("perulangan", i)
	}
}