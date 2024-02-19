package main

import "fmt"

func main() {

	go print("go routine")
	print("main")
	var str string
	fmt.Scanln(&str)
}

func print(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, str)
	}
}
