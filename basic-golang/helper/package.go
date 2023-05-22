package helper

import "fmt"

//Package adalah suatu tempat yang digunakan untuk mengorganisir kode program yang kita buat di golang
//Simplenya package adalah folder
//Dalam 1 package kita tidak boleh menggunakan nama method yang sama
//Tapi apabila kita membuat sebuah method dengan nama sama namun di beda package itu masih bisa
func sayHello() {
	fmt.Print("Hallooooo!")
}
