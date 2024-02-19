package main

import "fmt"

// Interface merupakan tipe data yang abstract dia tidak memiliki implementasi langsung sebuah interface berisikan definisi method
// Biasanya interface digunakan sebagai kontrak
// Semua tipe data yang memiliki kesamaan dengan interface maka akan secara langsung menjadi bagian interface itu sendiri

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Persons struct {
	Name string
}

func (person Persons) GetName() string {
	return person.Name
}

func main() {
	Indra := Persons{"Indra"}
	sayHello(Indra)
	fmt.Println(Indra)
}
