// Type asertions merupakan kemampuan golang untuk merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering digunakan apabila berhubungan dengan interface kosong

package main

func random() interface{} {
	return "Hi!"
}

// func main() {
// 	result := random()
// 	resultString := result.(string)

// 	fmt.Println(resultString)

// 	//Pembuatan type asertions dengan menggunakan metode switch
// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("String : ", value)
// 	case int:
// 		fmt.Println("Integer : ", value)
// 	default:
// 		fmt.Println("Unknown")
// 	}
// }
