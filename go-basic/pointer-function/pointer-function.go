package main

import "fmt"

type Addresses struct {
	Kabupaten string
	Provinsi  string
	Negara    string
}

type Man struct {
	Name string
}

// Apabila dilakukan seperti cara dibawah maka namanya tidak akan berubah menjadi Mr. Indra maka perlu ditambahkan operator *
func (man *Man) changeName() {
	man.Name = "Mr. " + man.Name
}

// Dengan menambahkan operator * kita bisa mengakses data yang dikirimkan dan mengubah isi dari datanya.
// Tanpa operator * kita hanya akan merubah duplikat nilainya saja
func changeNegara(addreses *Addresses) {
	addreses.Negara = "Indonesia"
}

func main() {
	address := Addresses{"Malang", "Jawa Timur", ""}
	changeNegara(&address)
	fmt.Println(address)

	Indra := Man{"Indra"}
	Indra.changeName()
	fmt.Println(Indra.Name)
}
