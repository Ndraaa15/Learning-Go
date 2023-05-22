package main

import "fmt"

type Person struct {
	Name    string
	Address string
	Age     int
}

//Struct method berfungsi untuk menambahkan komponen komponenm yang ada di dalam struct ke dalam function
//Secara singkatnya dengan struct method ini kita bisa mengkakses atribut di struct
//Seperti OOP
func (person Person) sayHello(name string) {
	fmt.Println("Hello", name, "My name is ", person.Name)
}

// func main() {
// 	Indra := Person{"Indra", "Borobudur", 19}
// 	Indra.sayHello("Bayu")
// }
