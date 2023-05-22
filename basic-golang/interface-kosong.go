package main

//Interface kosong bisa dibilang sebagai kelas paling atas yang ada di golang seperti java.lang.object
//Interface kosong adalah interface yang tidak memiliki deklarasi apapun yang membuat bisa me-return berbagai tipe data\

func say(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

// func main() {
// 	var value interface{} = say(1)
// 	fmt.Println(value)
// }
