package main

import "fmt"

func main() {
	var months = [12]string{
		"janury",
		"february",
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

	// array[low:high] -> index low sampai satu index sebelum index high
	var slice1 = months[4:7]
	fmt.Println(slice1)
	// array[low:] -> index low sampai index terakhir array
	var slice2 = months[4:]
	fmt.Println(slice2)
	// array[:high] -> index pertama array sampai satu index sebelum index high
	var slice3 = months[:4]
	fmt.Println(slice3)
	// array[:] -> index pertama sampai index terakhir array (semua member array)
	var slice4 = months[:]
	fmt.Println(slice4)

	// built-in function: length of slice
	fmt.Println(len(slice1))

	// built-in function: cap() -> capacity of slice
	fmt.Println(cap(slice1))

	// mengubah array, akan mengubah data slice
	months[5] = "months-xxx"
	fmt.Println(slice1)
	// mengubah slice, maka mengubah data array
	slice1[1] = "months-yyy"
	fmt.Println(months)

	// built-in function: append() -> membuat slice baru dengan menambah data ke posisi terkahir
	var slice5 = append(slice1, "months-zzz")	// data baru akan me-repalce posisi selanjutnya dari data terakhir slice
	fmt.Println(slice5)
	fmt.Println(months)

	var slice6 = months[10:]
	var slice7 = append(slice6, "months-aaa")	// data baru tidak dapat dimuat referensi array dari slice6 karena capai batas kapasitas, sehingga dibuat array baru
	slice7[0] = "months-bbb"

	fmt.Println(slice7)
	fmt.Println(slice6)
	fmt.Println(months)

	// built-in function : make([]DataType, length, capacity) -> membuat slice baru
	newSlice := make([]string, 2, 5)

	newSlice[0] = "jean"
	newSlice[1] = "taylor"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// built-in function: copy(destination, source) -> menyalin slice dari source ke destination
	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// hati-hati membuat array! ada perbedaan dengan slice
	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int {1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}

// slice adalah tipe data yang mereferensi array, jadi slice tidak menyimpan data, tetapi mereferensi data yang disimpan di array. Slice memiliki 3 properti yaitu length, capacity, dan pointer ke array. Slice dapat dibuat dengan cara slicing array atau dengan menggunakan built-in function make(). Slice juga memiliki built-in function append() untuk menambah data ke slice dan copy() untuk menyalin slice.
// slice memiliki kelebihan dibandingkan array yaitu lebih fleksibel karena dapat berubah ukuran, tetapi juga memiliki kekurangan yaitu lebih lambat karena harus mereferensi data di array. Slice juga memiliki perbedaan dengan array yaitu slice tidak menyimpan data, tetapi mereferensi data yang disimpan di array. Slice juga memiliki perbedaan dengan array yaitu slice tidak memiliki ukuran tetap seperti array, sehingga dapat berubah ukuran sesuai kebutuhan.