package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func main() {

// 	blacklist := func(name string) bool {
// 		return name == "admin"
// 	}
// 	registerUser("Indra", blacklist)
// 	registerUser("Admin", blacklist)

// }
