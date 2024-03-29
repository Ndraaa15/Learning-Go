package main

import "fmt"

//Type declarations untuk function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFiltered(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Eko", spamFiltered)

}
