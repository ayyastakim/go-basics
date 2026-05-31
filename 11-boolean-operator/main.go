package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	// var lulusUjian bool = ujian >= 80
	// var lulusAbsensi bool = absensi >= 80

	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// var result bool = lulusUjian && lulusAbsensi
	// fmt.Println(result)

	fmt.Println(ujian >= 80 && absensi >= 80)
}