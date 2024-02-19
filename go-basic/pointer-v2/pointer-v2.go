package main

import "fmt"

type Student struct {
	Name  string
	Email string
	Grade int
	Age   int
}

func (s *Student) getEmail1() string {
	return s.Email
}

func getEmail2(s Student) string {
	return s.Email
}

func (s *Student) ChangeEmail1(email string) {
	s.Email = email
}

//Ketika kita tidak menggunakan pointer kita tidak akan bisa melakukan pembaruan data pada suatu tahap

func main() {
	std := Student{"John Wick", "jw@foo.com", 11, 17}

	std.ChangeEmail1("foofoo@bar.com")

	fmt.Println(std.getEmail1())
	fmt.Println(getEmail2(std))
}
