package main

import "fmt"

func main() {
	//Type declaration ini digunakan untuk membuat alias atau nama lain dari type data yang sudah ada agar lebih mudah dimengerti
	type NIM string
	type status bool
	//Inisialisasi di atas menunjukan kita membuat type data baru yang namanya NIM dengan string
	var mahasiswa NIM = "225150200111006"
	var ActiveStatus status = true
	fmt.Println(mahasiswa)
	fmt.Println(NIM("225150200111004"))
	fmt.Println(ActiveStatus)
}
