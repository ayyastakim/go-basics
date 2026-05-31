package main

import "fmt"

func main() {
	
	var months = [12]string {
		"january", 
		"februari",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december",
	}

	var slice1 = months[4:]	// slice dari index-4 sampai index terakhir dari array
	fmt.Println(slice1)

	var slice2 = months[:3]	// slice dari index-0 sampai sebelum index-3 dari array
	fmt.Println(slice2)

	var slice3 = months[:]	// slice dengan mengambil semua data array (index-0 sampai index terakhir)
	fmt.Println(slice3)	

	var slice4 = months[4:7] // slice dari index-4 hingga sebelum index-7, pointer: data awal slice, length: ukuran data slice, capacity: ukuran data maksimal slice (pointer hingga index terakhir array)
	
	// built-in function: len() -> length
	fmt.Println(len(slice4))
	//built-in function: cap() -> capacity
	fmt.Println(cap(slice4))

	// slice reference ke array: mengubah data array maka data slice juga berubah dan sebaliknya
	// months[5] = "new-month"
	// fmt.Println(slice4)
	// fmt.Println(months)
	
	// slice4[2] ="another-month"
	// fmt.Println(slice4)
	// fmt.Println(months)
	
	// built-in function: append
	var slice5 = months[10:]
	fmt.Println(slice5)

	var slice6 = append(slice5, "new-month")	// data akhir slice5 sudah mencapai maks capacity, maka dibuat array baru untk menampung slice6
	fmt.Println(slice6)

	// tidak mempengaruhi slice5 maupun array months
	slice6[1] = "not-december"
	fmt.Println(slice6)
	fmt.Println(slice5)
	fmt.Println(months)

	var slice7 = append(slice4, "new-month2")	// data akhir slice4 belum mencapai capacity, maka data baru ditampung ke array aslinya
	fmt.Println(slice7)
	fmt.Println(months)

	// mempengaruhi slice4 dan array months
	slice7[1] = "not-month2"
	fmt.Println(slice7)
	fmt.Println(slice4)
	fmt.Println(months)

	// built-in function: make -> membuat slice baru
	newSlice := make([]string, 2, 5)	// membuat slice: dengan array tipe string, length 2, capacity 5

	newSlice[0] = "john"
	newSlice[1] = "taylor"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// built-in function: copy -> men-copas data slice dari source ke destination -> copy(destination, source)

	copySlice := make([]string, len(newSlice), cap(newSlice))	// pastikan length/ukuran destination sama dengan source-nya

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// perbedaan deklarasi aray dan slice!
	thisArray := [5]int8 {1, 2, 3, 4, 5}
	thisAnotherArray := [...]int8 {1, 2, 3, 4, 5}

	thisSlice := []int8 {1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisAnotherArray)
	fmt.Println(thisSlice)



}