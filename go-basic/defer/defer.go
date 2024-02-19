package main

import "fmt"

func endApp() {

	message := recover()

	if message != nil {

		fmt.Println("Terjadi Error", message)
	}
	fmt.Println("App End")
}

func errorApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	//Defer berfungsi untuk mengeksekusi suatu function apabila function lain telah selesai dijalankan baik itu error atau berhasil
	//Function defer ini hampir mirip dengan async await

	//Panic function berguna untuk menghentikan program yang sedang dijalankan
	//Panic function dipanggil ketika terjadi error pada saat program kita dijalankan
	//Meskipun panic dijalankan namun defer function tetap dilakukan

	//Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	//Dengan recover proses panic akan berhenti dan function akan tetap berjalan
	//Recover sendiri diletakkan di bagian function lain sebelum terdapat panic
	runApp()
	//Recover biasanya digunakan diletakkan di defer function
}
