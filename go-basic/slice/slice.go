package main

import "fmt"

func main() {
	//Slice adalah tipe data yang merupakan potongand dari array (di belakang slice ada array)
	//Slice sendiri hampir sama dengan array yang membedakan hanya ukurannya yang bisa berubah
	//Tipe data slice memiliki 3 data, yaitu pointer, length, capacity
	//Pointer menunjukkan data pertama di array para slice
	//Length menunjukkan panjang dari slice
	//Capacity menunujukkan kapasitas dari slice dimana capacity > length
	//Apabila data array diubah maka data slice juga ikut berubah
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[1:5]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	months[0] = "Tatarot"

	var slice2 = months[3:]

	var slice3 = append(slice2, "Holiday")
	fmt.Println(slice3)
	slice3[0] = "not April"
	fmt.Println(slice3)

	//Pengecekan slice 2
	fmt.Println(slice2)

	//Pengecekan months
	fmt.Println(months)

	//Apabila data dari slice yang ditambah element baru tidak cukup ditampung di dalam array maka secara otomatis akan membuat array baru
	//Apabila data dari append slice melebihi array maka apabila terdapat perubahan di slice tidak akan ber-impact dengan array

	//Slice Baru
	//2 Merupakan lengthnya dan 5 merupakan capacitynya
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Indra"
	newSlice[1] = "Adi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//Copyslice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//Array ditandai dengan adanya angka atau ... di dalam kotak
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

// func main() {
// 	array := [3]int{1, 2, 4}
// 	slice := array[0:]
// 	slice = append(slice, 7)
// 	slice[0] = 0
// 	fmt.Println(array)
// }
