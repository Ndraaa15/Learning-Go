package main

import "fmt"

func main() {
	//Switch expression lebih sederhana dari if condition
	//Switch condition biasanya dilakukan untuk pengecekan ke kondisi dalam satu variable
	//Sama dengan if expression switch statement juga bisa menggunakan short statement
	name := "Indra"

	switch name {
	case "Eko":
		fmt.Println("Hai Eko")
	case "Nami":
		fmt.Println("Hai Nami")
	case "Indra":
		fmt.Println("Hai Indra")
	default:
		fmt.Println("Hai, Boleh kenalan ga?")
	}
	//Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama panjang")
	case false:
		fmt.Println("Nama pendek")
	}
	//Switch statement dengan kondisi
	//Apabila menggunakan switch dalam kondisi maka tidak memerlukan parameter di switchnya
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length < 5:
		fmt.Println("Nama terlalu pendek")
	}

}
