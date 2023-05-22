package main

type Address struct {
	Kabupaten string
	Provinsi  string
	Negara    string
}

// func main() {
// 	//Secara default dalam golang semua variable itu di passing melalui value bukan by reference
// 	//Artinya ketika kita mengirim sebuah variable ke dalam function, method, atau variable lain yang kita kirimkan adalah duplikasi valuenya

// 	//Pass by Value
// 	address1 := Address{"Badung", "Bali", "Indonesia"}
// 	address2 := address1
// 	address2.Kabupaten = "Malang"
// 	address2.Provinsi = "Jawa Timur"
// 	fmt.Println(address1)
// 	//Nilai dari address1 tidak akan berubah karena addres2 tidak me-reference ke address1
// 	fmt.Println(address2)

// 	//Namun dengan adanya pointer bisa membuat program menjadi pass by reference bukan menjadi pass by value lagi
// 	//Dengan menggunakan pass by reference maka lokasi di memorinya pun sama
// 	//'&' tanda untuk melakukan pointer
// 	address3 := &address1
// 	address3.Kabupaten = "Singaraja"
// 	address3.Provinsi = "Bali"
// 	//Data dari address1 berubah
// 	fmt.Println(address1)

// 	//Operator * (penanda menuju reference mana)
// 	//Jika ingin merubah semua data yang mengacu pada suatu reference yang sama kita bisa menggunakan operator *
// 	//Pointer digunakan untuk menyimpan alamat memori suatu nilai

// 	var address4 *Address = &address1
// 	//Jadi dengan operator * variabel apapun yang mereference ke lokasi yang sama akan berubah nilainya menjadi sesuai dengan yang baru di assign
// 	*address3 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

// 	fmt.Println(address1)
// 	fmt.Println(address2) //Karena tidak mereference ke address1
// 	fmt.Println(address4)
// 	fmt.Println(address3)

// 	//Apabila ingin membuat sebuah variabel baru dan kosong bisa menggunakan keyword new
// 	address5 := new(Address)
// 	address5.Kabupaten = "Jombang"
// 	fmt.Println(address5)
// }
