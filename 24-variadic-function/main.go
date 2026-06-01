package main

import "fmt"

// variadic function: membuat parameter fungsi dengan data yang dinamis menggunakan varargs
// varargs: variable yang bisa menerima lebih dari satu input
// keunggulan dibanding menggunakan array atau slice:
// jika menggunakan array atau slice, kita wajib mendeklarasikan array atau slice tersebut sebelum dikirim ke parameter fungsi
// sedangkan dengan varargs, kita bisa langsung memasukan variable tersebut
// jika sebuah variable varargs menjadi parameter fungsi yang lebih, maka wajib berada pada posisi paling kanan atau final

// menggunakan varargs (variadic function)
func sumAll(numbers ...int16) int16 {
	var total int16 = 0

	for _, value := range numbers {	// ignore indexing
		total += value
	}

	return total
}


// menggunakan array
func sumAllArray(numbers [5]int) int {
	var total int = 0

	for _, value := range numbers {	// ignore indexing
		total += value
	}

	return total
}

func main() {

	// memanggil fungsi dengan parameter slice/array: perlu mempassing array / slice baru
	arrayOfNumbers := [...]int {10, 30, 50, 70, 90}

	var values = sumAllArray(arrayOfNumbers)

	fmt.Println(values)

	// memanggil variadic function: tidak perlu membuat array/slice baru
	var numbers int16 = sumAll(10, 20, 30, 40, 50)

	fmt.Println(numbers)

	// kondisi ketika mempunyai data slice atau array dan passing sebagai argumen dari fungsi variadic function sumAll()
	slices := []int16 {10, 20, 30, 40, 50}

	fmt.Println(sumAll(slices...))	// gunakan '...' di akhir nama varibel slice
}