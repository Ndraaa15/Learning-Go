package main

// import "fmt"

// func sumAll(numbers ...int) int {
// 	//Argument atau varags ini hanya bisa digunakand di final parameter
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func identity() (name string, college string, age int) {
// 	//Named return value
// 	name = "Gede Indra Adi Brata"
// 	college = "Braiwjaya University"
// 	age = 19
// 	return
// }

// func getFullNama() (string, string, string) {
// 	//Multiple return value
// 	return "Indra", "Adi", "Brata"
// }

// func sayHello(firstName string, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

// //Apabila menggunakan return function maka diperlukan tipe data dari function tersebut/
// func sayGoodBye(firstName string) string {
// 	return "Good Bye " + firstName
// }

// //Returning Multiple Value dan apabila

// func main() {

// 	name, college, age := identity()
// 	fmt.Println(name, college, age)

// 	firstName, _, lastName := getFullNama()
// 	fmt.Println(firstName, lastName)

// 	sayHello("Indra", "Brata")
// 	result := sayGoodBye("Indra")
// 	fmt.Println(result)

// 	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
// 	fmt.Println(total)

// 	slice1 := []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8}
// 	total = sumAll(slice1...)
// 	fmt.Println(total)
// 	//Function sendiri bisa digunakan sebagai value dalam sebuah variabel
// 	goodbye := sayGoodBye
// 	fmt.Println(goodbye("Indra"))
// }
