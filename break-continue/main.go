package main

import "fmt"

func main() {
	// break -> ketika sampai baris break, maka keluar dari perulangan
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("loop ke", i)
	}

	// continue -> ketika sampai continue, kode dibawahnya di-skip dan lanjut ke perulangan berikutnya
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println("loop ke", i)
	}

}