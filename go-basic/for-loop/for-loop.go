package main

import (
	"fmt"
)

func main() {
	//Dalam for loop bisa ditambahkan statement

	// i := 0
	// for i <= 10 {
	// 	fmt.Println("Perulangan ke-", i)
	// 	i++
	// }

	names := []string{"Indra", "Adi", "Brata", "Gede"}

	for i := 0; i <= len(names); i++ {
		fmt.Println(names[i])
	}

	//For range hampir sama dengan for each
	//Penambahan index itu optional apabila kita tidak ingin indexnya cukup berikan tanda _
	for index, name := range names {
		fmt.Println(name, "Mahasiswa ke ", index)
	}

	person := make(map[string]string)
	person["Name"] = "Indra"
	person["College"] = "Brawijaya University"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
