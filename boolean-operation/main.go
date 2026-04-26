package main

import "fmt"

func main() {
	ujian := 80
	absensi := 75

	lulusUjian := ujian >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusUjian && lulusAbsensi

	fmt.Println(lulus)
}