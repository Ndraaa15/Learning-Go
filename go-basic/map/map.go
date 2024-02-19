package main

import "fmt"

func main() {
	//Map adalah tipe data yang lain yang berisikan data dengan tipe sama namun kita bisa menentukan jenis tipe data index yang akan digunakan
	//Map adalah kumpulan key-value dimana kata kunci tidak boleh sama harus bersifat unik
	//Map tidak memiliki batasan jumlah asal data yang dimasukkan memiliki kata kunci yang berbeda dan apabila memasukkan kata kunci yang sama maka nilai yang akan digunakan adalah nilai yang baru
	person := map[string]string{
		"Name":    "Indra",
		"College": "Brawijaya University",
	}
	//Cara menginisiasi data baru
	person["address"] = "Borobudur Street"
	fmt.Println(person)
	fmt.Println(person["Name"])
	fmt.Println(person["College"])
	fmt.Println(len(person))

	//---

	var book map[string]string = make(map[string]string)
	book["title"] = "Atomic Habits"
	book["author"] = "James Clear"
	book["lol"] = "Hi!"

	delete(book, "lol")

}
