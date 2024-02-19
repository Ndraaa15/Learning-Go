package main

import "fmt"

func main() {
	//Array adalah kumpulan data dengan tipe data yang sama (daya tampung array harus di set dari awal dan tidak bisa ditambahkan)
	var names [3]string
	names[0] = "Indra"
	names[1] = "Ryo"
	names[2] = "Zaky"

	fmt.Println(names[0])

	var nums = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(nums)
	fmt.Println(len(names))
}
