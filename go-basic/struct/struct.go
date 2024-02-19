// Struct digunakan sebagai template data yang menghubungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// Data di struct disimpan dalam field sederhananya struct adalah kumpulan field
// Struct hampir mirip dengan class di java

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	//Struct template data atau prototype data
	var Indra Customer
	Indra.Name = "Indra Adi"
	Indra.Address = "Indonesia"
	Indra.Age = 19

	Ryo := Customer{
		Name:    "Ryo",
		Address: "Indonesia",
		Age:     19,
	}
	Javier := Customer{"Javier", "Indonesia", 19}

	fmt.Println(Indra)
	fmt.Println(Ryo)
	fmt.Println(Javier)
}
